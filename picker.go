package psd

import (
	"image"
	"image/color"
	"math"

	psdColor "github.com/oov/psd/color"
)

type Picker interface {
	SetSource(rect image.Rectangle, src ...[]byte)
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	Bounds() image.Rectangle
	// At returns the color of the pixel at (x, y).
	At(x, y int) color.Color
}

func findPicker(depth int, colorMode ColorMode, hasAlpha bool) Picker {
	switch colorMode {
	case ColorModeBitmap, ColorModeGrayscale:
		return findGrayPicker(depth, hasAlpha)
	case ColorModeRGB:
		return findNRGBAPicker(depth, hasAlpha)
	case ColorModeCMYK:
		return findNCMYKAPicker(depth, hasAlpha)
	}
	return nil
}

func findGrayPicker(depth int, hasAlpha bool) Picker {
	// TODO(oov): alpha support
	switch depth {
	case 1:
		return &pickerGray1{}
	case 8:
		return &pickerGray8{}
	case 16:
		return &pickerGray16{}
	case 32:
		return &pickerGray32{}
	}
	return nil
}

func findNRGBAPicker(depth int, hasAlpha bool) Picker {
	switch depth {
	case 8:
		if hasAlpha {
			return &pickerNRGBA8{}
		}
		return &pickerNRGB8{}
	case 16:
		if hasAlpha {
			return &pickerNRGBA16{}
		}
		return &pickerNRGB16{}
	}
	return nil
}

func findNCMYKAPicker(depth int, hasAlpha bool) Picker {
	switch depth {
	case 8:
		if hasAlpha {
			return &pickerNCMYKA8{}
		}
		return &pickerNCMYK8{}
	}
	return nil
}

type pickerPalette struct {
	Rect    image.Rectangle
	Src     []byte
	Palette color.Palette
}

func (p *pickerPalette) SetSource(rect image.Rectangle, src ...[]byte) { p.Rect, p.Src = rect, src[0] }
func (p *pickerPalette) ColorModel() color.Model                       { return p.Palette }
func (p *pickerPalette) Bounds() image.Rectangle                       { return p.Rect }
func (p *pickerPalette) At(x, y int) color.Color {
	pos := (y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X
	return p.Palette[p.Src[pos]]
}

type pickerGray1 struct {
	Rect image.Rectangle
	Src  []byte
}

func (p *pickerGray1) SetSource(rect image.Rectangle, src ...[]byte) { p.Rect, p.Src = rect, src[0] }
func (p *pickerGray1) ColorModel() color.Model                       { return psdColor.Gray1Model }
func (p *pickerGray1) Bounds() image.Rectangle                       { return p.Rect }
func (p *pickerGray1) At(x, y int) color.Color {
	xx := x - p.Rect.Min.X
	pos := (p.Rect.Dx()+7)>>3*(y-p.Rect.Min.Y) + xx>>3
	return psdColor.Gray1{p.Src[pos]&(1<<uint(^xx&7)) == 0}
}

type pickerGray8 struct {
	Rect image.Rectangle
	Src  []byte
}

func (p *pickerGray8) SetSource(rect image.Rectangle, src ...[]byte) { p.Rect, p.Src = rect, src[0] }
func (p *pickerGray8) ColorModel() color.Model                       { return color.GrayModel }
func (p *pickerGray8) Bounds() image.Rectangle                       { return p.Rect }
func (p *pickerGray8) At(x, y int) color.Color {
	pos := (y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X
	return color.Gray{Y: p.Src[pos]}
}

type pickerGray16 struct {
	Rect image.Rectangle
	Src  []byte
}

func (p *pickerGray16) SetSource(rect image.Rectangle, src ...[]byte) { p.Rect, p.Src = rect, src[0] }
func (p *pickerGray16) ColorModel() color.Model                       { return color.Gray16Model }
func (p *pickerGray16) Bounds() image.Rectangle                       { return p.Rect }
func (p *pickerGray16) At(x, y int) color.Color {
	pos := ((y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X) << 1
	return color.Gray16{Y: readUint16(p.Src, pos)}
}

type pickerGray32 struct {
	Rect image.Rectangle
	Src  []byte
}

func (p *pickerGray32) SetSource(rect image.Rectangle, src ...[]byte) { p.Rect, p.Src = rect, src[0] }
func (p *pickerGray32) ColorModel() color.Model                       { return psdColor.Gray32Model }
func (p *pickerGray32) Bounds() image.Rectangle                       { return p.Rect }
func (p *pickerGray32) At(x, y int) color.Color {
	pos := ((y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X) << 2
	return psdColor.Gray32{Y: math.Float32frombits(readUint32(p.Src, pos))}
}

type pickerNRGB8 struct {
	Rect    image.Rectangle
	R, G, B []byte
}

func (p *pickerNRGB8) SetSource(rect image.Rectangle, src ...[]byte) {
	p.Rect, p.R, p.G, p.B = rect, src[0], src[1], src[2]
}
func (p *pickerNRGB8) ColorModel() color.Model { return color.RGBAModel }
func (p *pickerNRGB8) Bounds() image.Rectangle { return p.Rect }
func (p *pickerNRGB8) At(x, y int) color.Color {
	pos := (y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X
	return color.NRGBA{
		p.R[pos],
		p.G[pos],
		p.B[pos],
		0xff,
	}
}

type pickerNRGBA8 struct {
	Rect       image.Rectangle
	R, G, B, A []byte
}

func (p *pickerNRGBA8) SetSource(rect image.Rectangle, src ...[]byte) {
	p.Rect, p.R, p.G, p.B, p.A = rect, src[0], src[1], src[2], src[3]
}
func (p *pickerNRGBA8) ColorModel() color.Model { return color.NRGBAModel }
func (p *pickerNRGBA8) Bounds() image.Rectangle { return p.Rect }
func (p *pickerNRGBA8) At(x, y int) color.Color {
	pos := (y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X
	return color.NRGBA{p.R[pos], p.G[pos], p.B[pos], p.A[pos]}
}

type pickerNRGB16 struct {
	Rect    image.Rectangle
	R, G, B []byte
}

func (p *pickerNRGB16) SetSource(rect image.Rectangle, src ...[]byte) {
	p.Rect, p.R, p.G, p.B = rect, src[0], src[1], src[2]
}
func (p *pickerNRGB16) ColorModel() color.Model { return color.NRGBA64Model }
func (p *pickerNRGB16) Bounds() image.Rectangle { return p.Rect }
func (p *pickerNRGB16) At(x, y int) color.Color {
	pos := ((y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X) << 1
	return color.NRGBA64{
		readUint16(p.R, pos),
		readUint16(p.G, pos),
		readUint16(p.B, pos),
		0xffff,
	}
}

type pickerNRGBA16 struct {
	Rect       image.Rectangle
	R, G, B, A []byte
}

func (p *pickerNRGBA16) SetSource(rect image.Rectangle, src ...[]byte) {
	p.Rect, p.R, p.G, p.B, p.A = rect, src[0], src[1], src[2], src[3]
}
func (p *pickerNRGBA16) ColorModel() color.Model { return color.NRGBA64Model }
func (p *pickerNRGBA16) Bounds() image.Rectangle { return p.Rect }
func (p *pickerNRGBA16) At(x, y int) color.Color {
	pos := ((y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X) << 1
	return color.NRGBA64{
		readUint16(p.R, pos),
		readUint16(p.G, pos),
		readUint16(p.B, pos),
		readUint16(p.A, pos),
	}
}

type pickerNCMYK8 struct {
	Rect       image.Rectangle
	C, M, Y, K []byte
}

func (p *pickerNCMYK8) SetSource(rect image.Rectangle, src ...[]byte) {
	p.Rect, p.C, p.M, p.Y, p.K = rect, src[0], src[1], src[2], src[3]
}
func (p *pickerNCMYK8) ColorModel() color.Model { return psdColor.NCMYKAModel }
func (p *pickerNCMYK8) Bounds() image.Rectangle { return p.Rect }
func (p *pickerNCMYK8) At(x, y int) color.Color {
	pos := (y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X
	return psdColor.NCMYKA{
		0xff - p.C[pos],
		0xff - p.M[pos],
		0xff - p.Y[pos],
		0xff - p.K[pos],
		0xff,
	}
}

type pickerNCMYKA8 struct {
	Rect          image.Rectangle
	C, M, Y, K, A []byte
}

func (p *pickerNCMYKA8) SetSource(rect image.Rectangle, src ...[]byte) {
	p.Rect, p.C, p.M, p.Y, p.K, p.A = rect, src[0], src[1], src[2], src[3], src[4]
}
func (p *pickerNCMYKA8) ColorModel() color.Model { return psdColor.NCMYKAModel }
func (p *pickerNCMYKA8) Bounds() image.Rectangle { return p.Rect }
func (p *pickerNCMYKA8) At(x, y int) color.Color {
	pos := (y-p.Rect.Min.Y)*p.Rect.Dx() + x - p.Rect.Min.X
	return psdColor.NCMYKA{
		0xff - p.C[pos],
		0xff - p.M[pos],
		0xff - p.Y[pos],
		0xff - p.K[pos],
		p.A[pos],
	}
}
