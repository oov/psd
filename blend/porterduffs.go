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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d clear) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawClearNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d clear) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawClearRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d clear) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawClearNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d clear) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawClearRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d clear) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawClearAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d clear) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawClearAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawClearNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawClearRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawClearNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawClearRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawClearAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawClearAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			var r, g, b, a, tmp uint32
			_ = tmp

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d clear) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d copy) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawCopyNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d copy) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawCopyRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d copy) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawCopyNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d copy) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawCopyRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d copy) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawCopyAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d copy) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawCopyAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawCopyNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawCopyRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawCopyNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawCopyRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawCopyAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawCopyAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa

			r = sr

			g = sg

			b = sb

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d copy) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d dest) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d dest) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d dest) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d dest) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d dest) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d dest) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawDestNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da

			r = dr

			g = dg

			b = db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d dest) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d srcOver) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOverNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOver) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOverRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOver) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOverNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOver) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOverRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOver) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOverAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d srcOver) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOverAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawSrcOverNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOverRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOverNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOverRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = sa + (tmp*da*32768)>>23

			r = sr + (tmp*dr*32768)>>23

			g = sg + (tmp*dg*32768)>>23

			b = sb + (tmp*db*32768)>>23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOverAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOverAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = sa + (tmp*da*32768)>>23

			r = sr + (tmp*dr*32768)>>23

			g = sg + (tmp*dg*32768)>>23

			b = sb + (tmp*db*32768)>>23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d srcOver) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d destOver) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOverNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOver) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOverRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOver) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOverNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOver) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOverRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOver) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestOverAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d destOver) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestOverAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawDestOverNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOverRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOverNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOverRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = da + (tmp*sa*32768)>>23

			r = dr + (tmp*sr*32768)>>23

			g = dg + (tmp*sg*32768)>>23

			b = db + (tmp*sb*32768)>>23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOverAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOverAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = da + (tmp*sa*32768)>>23

			r = dr + (tmp*sr*32768)>>23

			g = dg + (tmp*sg*32768)>>23

			b = db + (tmp*sb*32768)>>23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d destOver) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d srcIn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcInNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcIn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcInRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcIn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcInNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcIn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcInRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcIn) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcInAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d srcIn) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcInAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawSrcInNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcInRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcInNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcInRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcInAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcInAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (da * sr * 32768) >> 23

			g = (da * sg * 32768) >> 23

			b = (da * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d srcIn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d destIn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestInNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destIn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestInRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destIn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestInNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destIn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestInRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destIn) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestInAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d destIn) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestInAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawDestInNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestInRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestInNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestInRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestInAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestInAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (da * sa * 32768) >> 23

			r = (sa * dr * 32768) >> 23

			g = (sa * dg * 32768) >> 23

			b = (sa * db * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d destIn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d srcOut) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOutNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOut) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOutRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOut) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOutNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOut) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOutRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcOut) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOutAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d srcOut) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcOutAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawSrcOutNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOutRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOutNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOutRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOutAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcOutAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - da
			a = (tmp * sa * 32768) >> 23

			r = (tmp * sr * 32768) >> 23

			g = (tmp * sg * 32768) >> 23

			b = (tmp * sb * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d srcOut) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d destOut) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOutNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOut) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOutRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOut) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOutNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOut) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestOutRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destOut) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestOutAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d destOut) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestOutAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawDestOutNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOutRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOutNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOutRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOutAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestOutAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			tmp = 0xff - sa
			a = (tmp * da * 32768) >> 23

			r = (tmp * dr * 32768) >> 23

			g = (tmp * dg * 32768) >> 23

			b = (tmp * db * 32768) >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d destOut) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d srcAtop) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcAtopNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcAtop) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcAtopRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcAtop) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcAtopNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcAtop) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawSrcAtopRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d srcAtop) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcAtopAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d srcAtop) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawSrcAtopAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawSrcAtopNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcAtopRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcAtopNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcAtopRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da
			tmp = 0xff - sa

			r = (sr*da + dr*tmp) * 32768 >> 23

			g = (sg*da + dg*tmp) * 32768 >> 23

			b = (sb*da + db*tmp) * 32768 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcAtopAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSrcAtopAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = da
			tmp = 0xff - sa

			r = (sr*da + dr*tmp) * 32768 >> 23

			g = (sg*da + dg*tmp) * 32768 >> 23

			b = (sb*da + db*tmp) * 32768 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d srcAtop) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d destAtop) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestAtopNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destAtop) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestAtopRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destAtop) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestAtopNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destAtop) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawDestAtopRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d destAtop) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestAtopAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d destAtop) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawDestAtopAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawDestAtopNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAtopRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAtopNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAtopRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xff - da

			r = (sr*tmp + dr*sa) * 32768 >> 23

			g = (sg*tmp + dg*sa) * 32768 >> 23

			b = (sb*tmp + db*sa) * 32768 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAtopAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDestAtopAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = sa
			tmp = 0xff - da

			r = (sr*tmp + dr*sa) * 32768 >> 23

			g = (sg*tmp + dg*sa) * 32768 >> 23

			b = (sb*tmp + db*sa) * 32768 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d destAtop) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
	// d.drawFallback(dst, r, src, sp, nil, image.Point{})
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

func (d xOR) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawXORNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d xOR) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawXORRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d xOR) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawXORNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d xOR) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

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
		dyDelta, syDelta int
		x0, x1, xDelta   int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	drawXORRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d xOR) drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawXORAlphaToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

func (d xOR) drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform) {

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
		sx0, sx1, sxDelta, syDelta int
		dx0, dx1, dxDelta, dyDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		sx0, sx1, sxDelta = 0, dx, +1
		dx0, dx1, dxDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		sx0, sx1, sxDelta = (dx - 1), -1, -1
		dx0, dx1, dxDelta = (dx-1)<<2, -4, -4
	}
	drawXORAlphaToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)

}

var drawXORNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawXORRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawXORNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawXORRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (sa*(0xff-da) + da*(0xff-sa)) * 32768 >> 23

			r = (sr*(0xff-da) + dr*(0xff-sa)) * 32768 >> 23

			g = (sg*(0xff-da) + dg*(0xff-sa)) * 32768 >> 23

			b = (sb*(0xff-da) + db*(0xff-sa)) * 32768 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawXORAlphaToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

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

			dpix[j+3] = uint8(a)
			if a == 255 {
				dpix[j+2] = uint8(b)
				dpix[j+1] = uint8(g)
				dpix[j+0] = uint8(r)
			} else if a == 0 {
				dpix[j+2] = 0
				dpix[j+1] = 0
				dpix[j+0] = 0
			} else {
				dpix[j+2] = uint8(b * 0xff / a)
				dpix[j+1] = uint8(g * 0xff / a)
				dpix[j+0] = uint8(r * 0xff / a)
			}

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawXORAlphaToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	const (
		sb = 0
		sg = 0
		sr = 0
	)
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]

		spix := src[sPos:]

		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {

			sa := (uint32(spix[i]) * alpha) >> 23

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			var r, g, b, a, tmp uint32
			_ = tmp

			a = (sa*(0xff-da) + da*(0xff-sa)) * 32768 >> 23

			r = (sr*(0xff-da) + dr*(0xff-sa)) * 32768 >> 23

			g = (sg*(0xff-da) + dg*(0xff-sa)) * 32768 >> 23

			b = (sb*(0xff-da) + db*(0xff-sa)) * 32768 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(b)
			dpix[j+1] = uint8(g)
			dpix[j+0] = uint8(r)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

func (d xOR) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
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
