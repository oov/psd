package color

import "image/color"

// Gray32 represents a 32-bit float grayscale color.
type Gray32 struct {
	Y float32
}

// RGBA implements color.Color interface's method.
func (c Gray32) RGBA() (r, g, b, a uint32) {
	var y uint32
	switch {
	case c.Y > 1:
		y = 0xffff
	case c.Y < 0:
		y = 0
	default:
		y = uint32(c.Y * 0xffff)
	}
	return y, y, y, 0xffff
}

func gray32Model(c color.Color) color.Color {
	if _, ok := c.(Gray32); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 1000
	return Gray32{float32(y) / 0xffff}
}

// Gray1 represents an 1-bit monochrome bitmap color.
type Gray1 struct {
	B bool
}

// RGBA implements color.Color interface's method.
func (c Gray1) RGBA() (r, g, b, a uint32) {
	if c.B {
		return 0xffff, 0xffff, 0xffff, 0xffff
	}
	return 0, 0, 0, 0xffff
}

func gray1Model(c color.Color) color.Color {
	if _, ok := c.(Gray1); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := (299*r + 587*g + 114*b + 500) / 1000
	return Gray1{y >= 0x8000}
}

// NCMYKA represents a non-alpha-premultiplied CMYK color, having 8 bits for each of cyan,
// magenta, yellow, black and alpha.
// NCMYKA is different from color.CMYK, CMYK is inverted value.
//
// It is not associated with any particular color profile.
type NCMYKA struct {
	C, M, Y, K, A uint8
}

// RGBA implements color.Color interface's method.
func (c NCMYKA) RGBA() (uint32, uint32, uint32, uint32) {
	w := uint32(c.K) * 0x10201
	r := uint32(c.C) * w / 0xffff
	g := uint32(c.M) * w / 0xffff
	b := uint32(c.Y) * w / 0xffff
	if c.A == 0xff {
		return r, g, b, 0xffff
	}

	a := uint32(c.A) * 0x101
	r = r * a / 0xffff
	g = g * a / 0xffff
	b = b * a / 0xffff
	return r, g, b, a
}

func ncmykaModel(c color.Color) color.Color {
	if _, ok := c.(NCMYKA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	cc, mm, yy, kk := color.RGBToCMYK(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	cc = uint8((uint32(cc) * 0xffff) / a)
	mm = uint8((uint32(mm) * 0xffff) / a)
	yy = uint8((uint32(yy) * 0xffff) / a)
	kk = uint8((uint32(kk) * 0xffff) / a)
	return NCMYKA{255 - cc, 255 - mm, 255 - yy, 255 - kk, uint8(a >> 8)}
}

// These are color model.
var (
	Gray1Model  = color.ModelFunc(gray1Model)
	Gray32Model = color.ModelFunc(gray32Model)
	NCMYKAModel = color.ModelFunc(ncmykaModel)
)
