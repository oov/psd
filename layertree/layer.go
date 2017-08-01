// Package layertree implements PSD image drawing function.
//
// This package supports only RGBA color mode.
package layertree

import (
	"context"
	"image"
	"io"

	"github.com/oov/psd"

	"golang.org/x/text/encoding"
)

// Root represents root of the layer tree.
type Root struct {
	Renderer *Renderer

	Rect       image.Rectangle
	CanvasRect image.Rectangle

	Children []Layer
}

// Clone creates copy of r.
func (r *Root) Clone() *Root {
	return cloner{}.Clone(r)
}

// Layer represents the layer image.
type Layer struct {
	SeqID int

	Name        string
	DisplayName string

	Folder     bool
	FolderOpen bool

	Visible   bool
	BlendMode psd.BlendMode
	Opacity   int // 0-255
	Clipping  bool

	BlendClippedElements bool

	Rect image.Rectangle

	MaskEnabled      bool
	MaskRect         image.Rectangle
	MaskDefaultColor int // 0 or 255

	Parent   *Layer
	Children []Layer

	ClippedBy *Layer
	Clip      []*Layer
}

type Options struct {
	TileSize                  int
	LayerNameEncodingDetector func([]byte) encoding.Encoding
}

// New creates layer tree from the psdFile.
//
// New can cancel processing through ctx.
// If you pass 0 to tileSize, it will be defaultTileSize.
func New(ctx context.Context, psdFile io.Reader, opt *Options) (*Root, error) {
	if opt == nil {
		opt = &Options{
			TileSize:                  defaultTileSize,
			LayerNameEncodingDetector: func([]byte) encoding.Encoding { return encoding.Nop },
		}
	} else {
		if opt.TileSize == 0 {
			opt.TileSize = defaultTileSize
		}
	}
	renderer, img, err := newRenderer(ctx, psdFile, opt.TileSize)
	if err != nil {
		return nil, err
	}

	r := &Root{
		Renderer:   renderer,
		CanvasRect: img.Config.Rect,
	}
	renderer.layertree = r

	b := &builder{
		Img: img,
		LayerNameEncodingDetector: opt.LayerNameEncodingDetector,
	}
	for i := range img.Layer {
		r.Children = append(r.Children, Layer{})
		b.buildLayer(&r.Children[i], &img.Layer[i])
	}
	b.registerClippingGroup(r.Children)
	r.Rect = b.Rect.Intersect(r.CanvasRect)
	return r, nil
}
