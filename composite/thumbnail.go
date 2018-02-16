package composite

import (
	"context"
	"image"
	"runtime"

	"github.com/pkg/errors"

	"github.com/oov/downscale"
	"github.com/oov/psd/blend"
)

var (
	errLayerNotFound = errors.New("composite: layer not found")
	errNoCanvas      = errors.New("composite: no canvas")
)

func trimUnusedArea(src *image.RGBA, threshold byte) {
	w, h := src.Rect.Dx(), src.Rect.Dy()
	stride := src.Stride
	pix := src.Pix[:h*stride]
	rect := image.Rectangle{Min: image.Point{X: w, Y: h}, Max: image.Point{X: -1, Y: -1}}
	for y := 0; y < h; y++ {
		for l := 0; l < w; l++ {
			if pix[l*4+3] > threshold {
				if l < rect.Min.X {
					rect.Min.X = l
				}
				for r := w - 1; r >= l; r-- {
					if pix[r*4+3] > threshold {
						if r > rect.Max.X {
							rect.Max.X = r
						}
						break
					}
				}
				if y < rect.Min.Y {
					rect.Min.Y = y
				}
				rect.Max.Y = y
				break
			}
		}
		pix = pix[stride:]
	}

	if rect.Dx()*rect.Dy() == w*h {
		return
	}

	// trim
	d := src.Pix[:h*stride]
	s := d[stride*rect.Min.Y+rect.Min.X*4:]
	w, h = rect.Dx()*4, rect.Dy()
	for y := 0; y < h; y++ {
		copy(d, s[:w])
		d = d[w:]
		s = s[stride:]
	}
	src.Rect = rect
	src.Stride = w
	src.Pix = src.Pix[:rect.Dy()*w]
}

// Thumbnail creates thumbnail.
func (t *Tree) Thumbnail(seqID int, size int, tempBuffer []byte) (*image.RGBA, error) {
	ld, ok := t.layerImage[seqID]
	if !ok {
		return nil, errors.Wrap(errLayerNotFound, "composite: cannot create thumbnail")
	}
	if ld.Canvas == nil || ld.Canvas.Rect().Empty() {
		return nil, errors.Wrap(errNoCanvas, "composite: cannot create thumbnail")
	}
	rect := ld.Canvas.Rect()
	src := &image.RGBA{
		Rect:   rect,
		Stride: rect.Dx() * 4,
	}
	if src.Stride*src.Rect.Dy() > len(tempBuffer) {
		src.Pix = make([]byte, src.Stride*src.Rect.Dy())
	} else {
		src.Pix = tempBuffer[:src.Stride*src.Rect.Dy()]
	}
	if err := ld.Canvas.Render(context.Background(), src); err != nil {
		return nil, errors.Wrap(err, "composite: failed to create thumbnail")
	}

	trimUnusedArea(src, 0)
	rect = src.Rect

	var dest *image.RGBA
	sw, sh := rect.Dx(), rect.Dy()
	if sw > sh {
		dest = image.NewRGBA(image.Rect(0, 0, size, sh*size/sw))
	} else {
		dest = image.NewRGBA(image.Rect(0, 0, sw*size/sh, size))
	}
	if err := downscale.RGBAFast(context.Background(), dest, src); err != nil {
		return nil, errors.Wrap(err, "composite: failed to create thumbnail")
	}
	return dest, nil
}

func gatherLayer(layers *[]*Layer, l *Layer) {
	*layers = append(*layers, l)
	for i := range l.Children {
		gatherLayer(layers, &l.Children[i])
	}
}

func (t *Tree) Thumbnails(ctx context.Context, size int) (map[int]*image.RGBA, error) {
	var layers []*Layer
	gatherLayer(&layers, &t.Root)

	nLayers := len(layers)
	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > nLayers {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := nLayers / n

	m := make(map[int]*image.RGBA)
	idx := 0
	for i := 1; i < n; i++ {
		go t.thumbnailsInner(pc, m, layers, size, idx, idx+step)
		idx += step
	}
	go t.thumbnailsInner(pc, m, layers, size, idx, nLayers)
	if err := pc.Wait(ctx); err != nil {
		return nil, err
	}
	return m, nil

}

func (t *Tree) thumbnailsInner(pc *parallelContext, m map[int]*image.RGBA, layers []*Layer, size, sIdx, eIdx int) {
	defer pc.Done()

	bufLen := 0
	for i := sIdx; i < eIdx; i++ {
		ld, ok := t.layerImage[layers[i].SeqID]
		if !ok || ld.Canvas == nil {
			continue
		}
		rect := ld.Canvas.Rect()
		l := rect.Dx() * 4 * rect.Dy()
		if bufLen < l {
			bufLen = l
		}
	}
	var buf []byte
	if bufLen > 0 {
		buf = make([]byte, bufLen)
	}

	for i := sIdx; i < eIdx; i++ {
		t, err := t.Thumbnail(layers[i].SeqID, size, buf)
		if err != nil {
			continue
		}
		pc.M.Lock()
		m[layers[i].SeqID] = t
		pc.M.Unlock()
	}
}

func (t *Tree) ThumbnailSheet(ctx context.Context, size int) (*image.RGBA, map[int]image.Rectangle, error) {
	m, err := t.Thumbnails(ctx, size)
	if err != nil {
		return nil, nil, err
	}

	var indices []int
	for i := range m {
		indices = append(indices, i)
	}

	textureSize := 1
	nIndices := len(indices)
	for {
		n := (textureSize / size) * (textureSize / size)
		if n >= nIndices {
			break
		}
		textureSize += textureSize
	}
	if textureSize > 4096 {
		return nil, nil, errors.New("composite: could not create thumbnail sheet because too many layers")
	}
	img := image.NewRGBA(image.Rect(0, 0, textureSize, textureSize))
	mpt := make(map[int]image.Rectangle)

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > nIndices {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := nIndices / n

	idx := 0
	for i := 1; i < n; i++ {
		t.thumbnailSheetInner(pc, img, mpt, m, indices, idx, idx+step, size)
		idx += step
	}
	t.thumbnailSheetInner(pc, img, mpt, m, indices, idx, nIndices, size)
	if err := pc.Wait(ctx); err != nil {
		return nil, nil, err
	}
	return img, mpt, nil

}

func (t *Tree) thumbnailSheetInner(pc *parallelContext, img *image.RGBA, mpt map[int]image.Rectangle, m map[int]*image.RGBA, indices []int, sIdx, eIdx, size int) {
	defer pc.Done()
	iw := img.Rect.Dx() / size
	for i := sIdx; i < eIdx; i++ {
		if pc.Aborted() {
			return
		}
		idx := indices[i]
		thumb := m[idx]
		n := i / iw
		x, y := (i-n*iw)*size, n*size
		ox, oy := x+(size-thumb.Rect.Dx())/2, y+(size-thumb.Rect.Dy())/2
		rect := image.Rect(x, y, x+size, y+size)
		blend.Copy.Draw(img, image.Rect(ox, oy, ox+thumb.Rect.Dx(), oy+thumb.Rect.Dy()), thumb, image.Point{})
		pc.M.Lock()
		mpt[idx] = rect
		pc.M.Unlock()
	}
}
