// DO NOT EDIT.
// Generate with: go generate

package blend

import (
	"image"
	stdcolor "image/color"
	"image/draw"
)

// porter/duff compositing modes
var (
	Clear    draw.Drawer = clear{}
	Copy     draw.Drawer = copy{}
	Dest     draw.Drawer = dest{}
	SrcOver  draw.Drawer = srcOver{}
	DestOver draw.Drawer = destOver{}
	SrcIn    draw.Drawer = srcIn{}
	DestIn   draw.Drawer = destIn{}
	SrcOut   draw.Drawer = srcOut{}
	DestOut  draw.Drawer = destOut{}
	SrcAtop  draw.Drawer = srcAtop{}
	DestAtop draw.Drawer = destAtop{}
	XOR      draw.Drawer = xOR{}
)

// clear implements the clear porter-duff mode.
type clear struct{}

// String implemenets fmt.Stringer interface.
func (d clear) String() string {
	return "Clear"
}

// Draw implements image.Drawer interface.
func (d clear) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d clear) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawClearNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d clear) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawClearRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d clear) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawClearNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d clear) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawClearRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawClearNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawClearRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawClearNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawClearRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d clear) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			var a, r, g, b, tmp uint32
			_ = tmp

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// copy implements the copy porter-duff mode.
type copy struct{}

// String implemenets fmt.Stringer interface.
func (d copy) String() string {
	return "Copy"
}

// Draw implements image.Drawer interface.
func (d copy) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d copy) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawCopyNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d copy) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawCopyRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d copy) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawCopyNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d copy) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawCopyRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawCopyNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawCopyRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawCopyNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawCopyRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d copy) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// dest implements the dest porter-duff mode.
type dest struct{}

// String implemenets fmt.Stringer interface.
func (d dest) String() string {
	return "Dest"
}

// Draw implements image.Drawer interface.
func (d dest) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d dest) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d dest) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d dest) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d dest) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDestNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i := xMin; i != xMax; i += xDelta {

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d dest) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// srcOver implements the srcOver porter-duff mode.
type srcOver struct{}

// String implemenets fmt.Stringer interface.
func (d srcOver) String() string {
	return "SrcOver"
}

// Draw implements image.Drawer interface.
func (d srcOver) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d srcOver) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOverNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcOver) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOverRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcOver) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOverNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcOver) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOverRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSrcOverNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = sa + (tmp*da*32768)>>23

			r = sr + (tmp*dr*32768)>>23

			g = sg + (tmp*dg*32768)>>23

			b = sb + (tmp*db*32768)>>23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcOverRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = sa + (tmp*da*32768)>>23

			r = sr + (tmp*dr*32768)>>23

			g = sg + (tmp*dg*32768)>>23

			b = sb + (tmp*db*32768)>>23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcOverNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = sa + (tmp*da*32768)>>23

			r = sr + (tmp*dr*32768)>>23

			g = sg + (tmp*dg*32768)>>23

			b = sb + (tmp*db*32768)>>23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcOverRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = sa + (tmp*da*32768)>>23

			r = sr + (tmp*dr*32768)>>23

			g = sg + (tmp*dg*32768)>>23

			b = sb + (tmp*db*32768)>>23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d srcOver) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			tmp = 0xffff - sa
			a = sa + (tmp*da)/0xffff

			r = sr + (tmp*dr)/0xffff

			g = sg + (tmp*dg)/0xffff

			b = sb + (tmp*db)/0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// destOver implements the destOver porter-duff mode.
type destOver struct{}

// String implemenets fmt.Stringer interface.
func (d destOver) String() string {
	return "DestOver"
}

// Draw implements image.Drawer interface.
func (d destOver) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d destOver) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOverNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destOver) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOverRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destOver) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOverNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destOver) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOverRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDestOverNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = da + (tmp*sa*32768)>>23

			r = dr + (tmp*sr*32768)>>23

			g = dg + (tmp*sg*32768)>>23

			b = db + (tmp*sb*32768)>>23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestOverRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = da + (tmp*sa*32768)>>23

			r = dr + (tmp*sr*32768)>>23

			g = dg + (tmp*sg*32768)>>23

			b = db + (tmp*sb*32768)>>23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestOverNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = da + (tmp*sa*32768)>>23

			r = dr + (tmp*sr*32768)>>23

			g = dg + (tmp*sg*32768)>>23

			b = db + (tmp*sb*32768)>>23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestOverRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = da + (tmp*sa*32768)>>23

			r = dr + (tmp*sr*32768)>>23

			g = dg + (tmp*sg*32768)>>23

			b = db + (tmp*sb*32768)>>23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d destOver) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			tmp = 0xffff - da
			a = da + (tmp*sa)/0xffff

			r = dr + (tmp*sr)/0xffff

			g = dg + (tmp*sg)/0xffff

			b = db + (tmp*sb)/0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// srcIn implements the srcIn porter-duff mode.
type srcIn struct{}

// String implemenets fmt.Stringer interface.
func (d srcIn) String() string {
	return "SrcIn"
}

// Draw implements image.Drawer interface.
func (d srcIn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d srcIn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcInNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcIn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcInRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcIn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcInNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcIn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcInRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSrcInNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcInRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcInNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcInRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d srcIn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			_, _, _, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = (da * sa) / 0xffff

			r = (da * sr) / 0xffff

			g = (da * sg) / 0xffff

			b = (da * sb) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// destIn implements the destIn porter-duff mode.
type destIn struct{}

// String implemenets fmt.Stringer interface.
func (d destIn) String() string {
	return "DestIn"
}

// Draw implements image.Drawer interface.
func (d destIn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d destIn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestInNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destIn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestInRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destIn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestInNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destIn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestInRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDestInNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestInRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestInNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestInRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d destIn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			_, _, _, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = (da * sa) / 0xffff

			r = (sa * dr) / 0xffff

			g = (sa * dg) / 0xffff

			b = (sa * db) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// srcOut implements the srcOut porter-duff mode.
type srcOut struct{}

// String implemenets fmt.Stringer interface.
func (d srcOut) String() string {
	return "SrcOut"
}

// Draw implements image.Drawer interface.
func (d srcOut) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d srcOut) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOutNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcOut) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOutRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcOut) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOutNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcOut) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcOutRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSrcOutNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcOutRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcOutNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcOutRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d srcOut) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			_, _, _, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			tmp = 0xffff - da
			a = (tmp * sa) / 0xffff

			r = (tmp * sr) / 0xffff

			g = (tmp * sg) / 0xffff

			b = (tmp * sb) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// destOut implements the destOut porter-duff mode.
type destOut struct{}

// String implemenets fmt.Stringer interface.
func (d destOut) String() string {
	return "DestOut"
}

// Draw implements image.Drawer interface.
func (d destOut) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d destOut) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOutNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destOut) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOutRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destOut) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOutNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destOut) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestOutRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDestOutNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestOutRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestOutNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestOutRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d destOut) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			_, _, _, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			tmp = 0xffff - sa
			a = (tmp * da) / 0xffff

			r = (tmp * dr) / 0xffff

			g = (tmp * dg) / 0xffff

			b = (tmp * db) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// srcAtop implements the srcAtop porter-duff mode.
type srcAtop struct{}

// String implemenets fmt.Stringer interface.
func (d srcAtop) String() string {
	return "SrcAtop"
}

// Draw implements image.Drawer interface.
func (d srcAtop) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d srcAtop) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcAtopNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcAtop) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcAtopRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcAtop) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcAtopNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d srcAtop) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawSrcAtopRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSrcAtopNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da
			tmp = 0xff - sa

			r = (sr*da + dr*tmp) * 32768 >> 23

			g = (sg*da + dg*tmp) * 32768 >> 23

			b = (sb*da + db*tmp) * 32768 >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcAtopRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da
			tmp = 0xff - sa

			r = (sr*da + dr*tmp) * 32768 >> 23

			g = (sg*da + dg*tmp) * 32768 >> 23

			b = (sb*da + db*tmp) * 32768 >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcAtopNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da
			tmp = 0xff - sa

			r = (sr*da + dr*tmp) * 32768 >> 23

			g = (sg*da + dg*tmp) * 32768 >> 23

			b = (sb*da + db*tmp) * 32768 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSrcAtopRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da
			tmp = 0xff - sa

			r = (sr*da + dr*tmp) * 32768 >> 23

			g = (sg*da + dg*tmp) * 32768 >> 23

			b = (sb*da + db*tmp) * 32768 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d srcAtop) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = da
			tmp = 0xffff - sa

			r = (sr*da + dr*tmp) / 0xffff

			g = (sg*da + dg*tmp) / 0xffff

			b = (sb*da + db*tmp) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// destAtop implements the destAtop porter-duff mode.
type destAtop struct{}

// String implemenets fmt.Stringer interface.
func (d destAtop) String() string {
	return "DestAtop"
}

// Draw implements image.Drawer interface.
func (d destAtop) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d destAtop) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestAtopNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destAtop) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestAtopRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destAtop) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestAtopNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d destAtop) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawDestAtopRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDestAtopNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xff - da

			r = (sr*tmp + dr*sa) * 32768 >> 23

			g = (sg*tmp + dg*sa) * 32768 >> 23

			b = (sb*tmp + db*sa) * 32768 >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestAtopRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xff - da

			r = (sr*tmp + dr*sa) * 32768 >> 23

			g = (sg*tmp + dg*sa) * 32768 >> 23

			b = (sb*tmp + db*sa) * 32768 >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestAtopNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xff - da

			r = (sr*tmp + dr*sa) * 32768 >> 23

			g = (sg*tmp + dg*sa) * 32768 >> 23

			b = (sb*tmp + db*sa) * 32768 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDestAtopRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xff - da

			r = (sr*tmp + dr*sa) * 32768 >> 23

			g = (sg*tmp + dg*sa) * 32768 >> 23

			b = (sb*tmp + db*sa) * 32768 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d destAtop) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xffff - da

			r = (sr*tmp + dr*sa) / 0xffff

			g = (sg*tmp + dg*sa) / 0xffff

			b = (sb*tmp + db*sa) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}

// xOR implements the xOR porter-duff mode.
type xOR struct{}

// String implemenets fmt.Stringer interface.
func (d xOR) String() string {
	return "XOR"
}

// Draw implements image.Drawer interface.
func (d xOR) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d xOR) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawXORNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d xOR) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawXORRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d xOR) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawXORNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d xOR) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	drawXORRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawXORNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (sa*(0xff-da) + da*(0xff-sa)) * 32768 >> 23

			r = (sr*(0xff-da) + dr*(0xff-sa)) * 32768 >> 23

			g = (sg*(0xff-da) + dg*(0xff-sa)) * 32768 >> 23

			b = (sb*(0xff-da) + db*(0xff-sa)) * 32768 >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawXORRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (sa*(0xff-da) + da*(0xff-sa)) * 32768 >> 23

			r = (sr*(0xff-da) + dr*(0xff-sa)) * 32768 >> 23

			g = (sg*(0xff-da) + dg*(0xff-sa)) * 32768 >> 23

			b = (sb*(0xff-da) + db*(0xff-sa)) * 32768 >> 23

			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawXORNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (sa*(0xff-da) + da*(0xff-sa)) * 32768 >> 23

			r = (sr*(0xff-da) + dr*(0xff-sa)) * 32768 >> 23

			g = (sg*(0xff-da) + dg*(0xff-sa)) * 32768 >> 23

			b = (sb*(0xff-da) + db*(0xff-sa)) * 32768 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawXORRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i := xMin; i != xMax; i += xDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (sa*(0xff-da) + da*(0xff-sa)) * 32768 >> 23

			r = (sr*(0xff-da) + dr*(0xff-sa)) * 32768 >> 23

			g = (sg*(0xff-da) + dg*(0xff-sa)) * 32768 >> 23

			b = (sb*(0xff-da) + db*(0xff-sa)) * 32768 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d xOR) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()

			dr, dg, db, da := dst.At(x, y).RGBA()

			var a, r, g, b, tmp uint32
			_ = tmp

			a = (sa*(0xffff-da) + da*(0xffff-sa)) / 0xffff

			r = (sr*(0xffff-da) + dr*(0xffff-sa)) / 0xffff

			g = (sg*(0xffff-da) + dg*(0xffff-sa)) / 0xffff

			b = (sb*(0xffff-da) + db*(0xffff-sa)) / 0xffff

			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}
