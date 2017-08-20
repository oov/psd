// Package layertree implements PSD image drawing function.
//
// This package supports only RGBA color mode.
package layertree

import (
	"context"
	"errors"
	"image"
	"io"
	"math"
	"runtime"

	"golang.org/x/image/math/f64"
	"golang.org/x/text/encoding"

	"github.com/oov/psd"
)

// Root represents root of the layer tree.
type Root struct {
	Renderer *Renderer

	tileSize int

	layerImage map[int]layerImage

	Rect       image.Rectangle
	CanvasRect image.Rectangle

	Children []Layer
}

// Clone creates copy of r.
//
// Required memory is not very large because layer images are shared between them.
func (r *Root) Clone() *Root {
	return cloner{}.Clone(r)
}

// Transform creates copy of r that transformed by m.
//
// This takes time because it applies transformations to all layers.
func (r *Root) Transform(ctx context.Context, m f64.Aff3) (*Root, error) {
	rr := r.Clone()

	// flatten layerImage
	nImages := len(r.layerImage)
	seqIDs := make([]int, nImages)
	images := make([]layerImage, nImages)
	i := 0
	for seqID, li := range r.layerImage {
		seqIDs[i] = seqID
		images[i] = li
		i++
	}

	rr.layerImage = map[int]layerImage{}

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > nImages {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := nImages / n
	idx := 0
	for i := 1; i < n; i++ {
		go rr.transformInner(pc, m, seqIDs, images, idx, idx+step)
		idx += step
	}
	go rr.transformInner(pc, m, seqIDs, images, idx, nImages)
	if err := pc.Wait(ctx); err != nil {
		return nil, err
	}
	return rr, nil

}

func (r *Root) transformInner(pc *parallelContext, m f64.Aff3, seqIDs []int, images []layerImage, sIdx, eIdx int) {
	defer pc.Done()
	for i := sIdx; i < eIdx; i++ {
		if pc.Aborted() {
			return
		}

		li := images[i]
		var ti tiledImage
		var tm tiledMask
		var err error
		if li.Canvas != nil {
			if ti, err = li.Canvas.Transform(context.Background(), m); err != nil {
				return
			}
		}
		if li.Mask != nil {
			if tm, err = li.Mask.Transform(context.Background(), m); err != nil {
				return
			}
		}
		pc.M.Lock()
		r.layerImage[seqIDs[i]] = layerImage{ti, tm}
		pc.M.Unlock()
	}
}

// Layer represents the layer image.
type Layer struct {
	SeqID int

	Name string

	Folder     bool
	FolderOpen bool

	Visible   bool
	BlendMode psd.BlendMode
	Opacity   int // 0-255
	Clipping  bool

	BlendClippedElements bool

	MaskEnabled      bool
	MaskDefaultColor int // 0 or 255

	Parent   *Layer
	Children []Layer

	ClippedBy *Layer
	Clip      []*Layer
}

const DefaultTileSize = 64

type Options struct {
	TileSize        int
	TransformMatrix f64.Aff3
	// It will used to detect character encoding of a variable-width encoding layer name.
	LayerNameEncodingDetector func([]byte) encoding.Encoding
}

func transformRect(r image.Rectangle, m f64.Aff3) image.Rectangle {
	pts := []float64{
		float64(r.Min.X), float64(r.Min.Y),
		float64(r.Max.X), float64(r.Min.Y),
		float64(r.Max.X), float64(r.Max.Y),
		float64(r.Min.X), float64(r.Max.Y),
	}
	var xMin, yMin, xMax, yMax float64
	for i := 0; i < len(pts); i += 2 {
		sx, sy := pts[i], pts[i+1]
		dx, dy := sx*m[0]+sy*m[1]+m[2], sx*m[3]+sy*m[4]+m[5]
		if i == 0 {
			xMin, yMin = dx, dy
			xMax, yMax = dx+1, dy+1
			continue
		}
		if xMin > dx {
			xMin = dx
		}
		if yMin > dy {
			yMin = dy
		}
		dx++
		dy++
		if xMax < dx {
			xMax = dx
		}
		if yMax < dy {
			yMax = dy
		}
	}
	return image.Rectangle{
		Min: image.Point{X: int(math.Floor(xMin)), Y: int(math.Floor(yMin))},
		Max: image.Point{X: int(math.Ceil(xMax)), Y: int(math.Ceil(yMax))},
	}
}

// New creates layer tree from the psdFile.
//
// New can cancel processing through ctx.
// If you pass 0 to opt.TileSize, it will be DefaultTileSize.
func New(ctx context.Context, psdFile io.Reader, opt *Options) (*Root, error) {
	if opt == nil {
		opt = &Options{}
	}
	if opt.TileSize == 0 {
		opt.TileSize = DefaultTileSize
	}
	if opt.TransformMatrix[0] == 0 {
		opt.TransformMatrix[0] = 1
	}
	if opt.TransformMatrix[4] == 0 {
		opt.TransformMatrix[4] = 1
	}
	if opt.LayerNameEncodingDetector == nil {
		opt.LayerNameEncodingDetector = func([]byte) encoding.Encoding { return encoding.Nop }
	}

	layerImages, img, err := createCanvas(ctx, psdFile, opt.TileSize, opt.TransformMatrix)
	if err != nil {
		return nil, err
	}

	b := &builder{
		Img: img,
		LayerNameEncodingDetector: opt.LayerNameEncodingDetector,
	}
	var layers []Layer
	for i := range img.Layer {
		layers = append(layers, Layer{})
		b.buildLayer(&layers[i], &img.Layer[i])
	}
	b.registerClippingGroup(layers)

	r := &Root{
		Renderer: &Renderer{
			cache: map[int]*cache{},
		},
		tileSize:   opt.TileSize,
		layerImage: layerImages,

		CanvasRect: img.Config.Rect,
		Rect:       b.Rect.Intersect(img.Config.Rect),

		Children: layers,
	}
	r.Renderer.pool.New = r.Renderer.allocate
	r.Renderer.layertree = r
	return r, nil
}

func createCanvas(ctx context.Context, psdFile io.Reader, tileSize int, m f64.Aff3) (map[int]layerImage, *psd.PSD, error) {
	n := runtime.GOMAXPROCS(0)

	ch := make(chan *psd.Layer)
	layerImages := map[int]layerImage{}

	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pc := &parallelContext{}
	pc.Wg.Add(n)
	for i := 0; i < n; i++ {
		go createCanvasInner(cctx, pc, ch, tileSize, m, layerImages)
	}
	img, _, err := psd.Decode(psdFile, &psd.DecodeOptions{
		SkipMergedImage: true,
		ConfigLoaded: func(c psd.Config) error {
			if c.ColorMode != psd.ColorModeRGB {
				return errors.New("Unsupported color mode")
			}
			return nil
		},
		LayerImageLoaded: func(l *psd.Layer, index int, total int) { ch <- l },
	})
	close(ch)
	if err != nil {
		return nil, nil, err
	}
	if err = pc.Wait(ctx); err != nil {
		return nil, nil, err
	}
	img.Config.Rect = transformRect(img.Config.Rect, m)
	return layerImages, img, nil
}

func createCanvasInner(ctx context.Context, pc *parallelContext, ch <-chan *psd.Layer, tileSize int, m f64.Aff3, layerImages map[int]layerImage) {
	defer pc.Done()
	for l := range ch {
		var ld layerImage
		if l.HasImage() && !l.Rect.Empty() {
			r, g, b := l.Channel[0].Data, l.Channel[1].Data, l.Channel[2].Data
			var a []byte
			if ach, ok := l.Channel[-1]; ok {
				a = ach.Data
			}
			ti, err := newScaledTiledImage(ctx, tileSize, l.Rect, r, g, b, a, 1, m)
			if err != nil {
				return
			}
			ld.Canvas = ti
			l.Rect = ti.Rect()
		}
		if !l.Mask.Rect.Empty() {
			if a, ok := l.Channel[-2]; ok {
				m, err := newScaledTiledMask(ctx, tileSize, l.Mask.Rect, a.Data, l.Mask.DefaultColor, m)
				if err != nil {
					return
				}
				ld.Mask = m
				l.Mask.Rect = m.Rect()
			}
		}

		pc.M.Lock()
		layerImages[l.SeqID] = ld
		pc.M.Unlock()
	}
}
