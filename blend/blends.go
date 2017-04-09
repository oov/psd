// DO NOT EDIT.
// Generate with: go generate

package blend

import (
	"image"
	stdcolor "image/color"
	"image/draw"
)

// blend modes
var (
	Normal       Drawer = normal{}
	Darken       Drawer = darken{}
	Multiply     Drawer = multiply{}
	ColorBurn    Drawer = colorBurn{}
	LinearBurn   Drawer = linearBurn{}
	DarkerColor  Drawer = darkerColor{}
	Lighten      Drawer = lighten{}
	Screen       Drawer = screen{}
	ColorDodge   Drawer = colorDodge{}
	LinearDodge  Drawer = linearDodge{}
	LighterColor Drawer = lighterColor{}
	Add          Drawer = add{}
	Overlay      Drawer = overlay{}
	SoftLight    Drawer = softLight{}
	HardLight    Drawer = hardLight{}
	LinearLight  Drawer = linearLight{}
	VividLight   Drawer = vividLight{}
	PinLight     Drawer = pinLight{}
	HardMix      Drawer = hardMix{}
	Difference   Drawer = difference{}
	Exclusion    Drawer = exclusion{}
	Subtract     Drawer = subtract{}
	Divide       Drawer = divide{}
	Hue          Drawer = hue{}
	Saturation   Drawer = saturation{}
	Color        Drawer = color{}
	Luminosity   Drawer = luminosity{}
)

// normal implements the normal blend mode.
type normal struct{}

// String implemenets fmt.Stringer interface.
func (d normal) String() string {
	return "Normal"
}

// Draw implements image.Drawer interface.
func (d normal) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d normal) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d normal) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawNormalNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d normal) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawNormalRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d normal) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawNormalNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d normal) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawNormalRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawNormalNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawNormalRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawNormalNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawNormalRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawNormalFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d normal) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawNormalFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// darken implements the darken blend mode.
type darken struct{}

// String implemenets fmt.Stringer interface.
func (d darken) String() string {
	return "Darken"
}

// Draw implements image.Drawer interface.
func (d darken) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d darken) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d darken) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkenNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d darken) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkenRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d darken) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkenNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d darken) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkenRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawDarkenNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < dr {
				r = sr
			} else {
				r = dr
			}

			if sg < dg {
				g = sg
			} else {
				g = dg
			}

			if sb < db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkenRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < dr {
				r = sr
			} else {
				r = dr
			}

			if sg < dg {
				g = sg
			} else {
				g = dg
			}

			if sb < db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkenNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < dr {
				r = sr
			} else {
				r = dr
			}

			if sg < dg {
				g = sg
			} else {
				g = dg
			}

			if sb < db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkenRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < dr {
				r = sr
			} else {
				r = dr
			}

			if sg < dg {
				g = sg
			} else {
				g = dg
			}

			if sb < db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkenFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < dr {
				r = sr
			} else {
				r = dr
			}

			if sg < dg {
				g = sg
			} else {
				g = dg
			}

			if sb < db {
				b = sb
			} else {
				b = db
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d darken) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawDarkenFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// multiply implements the multiply blend mode.
type multiply struct{}

// String implemenets fmt.Stringer interface.
func (d multiply) String() string {
	return "Multiply"
}

// Draw implements image.Drawer interface.
func (d multiply) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d multiply) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d multiply) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawMultiplyNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d multiply) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawMultiplyRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d multiply) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawMultiplyNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d multiply) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawMultiplyRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawMultiplyNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawMultiplyRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawMultiplyNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawMultiplyRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawMultiplyFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r = sr * dr / 0xffff

			g = sg * dg / 0xffff

			b = sb * db / 0xffff

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d multiply) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawMultiplyFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// colorBurn implements the colorBurn blend mode.
type colorBurn struct{}

// String implemenets fmt.Stringer interface.
func (d colorBurn) String() string {
	return "ColorBurn"
}

// Draw implements image.Drawer interface.
func (d colorBurn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d colorBurn) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d colorBurn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorBurnNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d colorBurn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorBurnRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d colorBurn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorBurnNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d colorBurn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorBurnRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawColorBurnNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if dr == 0xff {
				r = 0xff
			} else if sr == 0 {
				r = 0
			} else {
				r = 0xff - clip8((0xff-dr)*0xff/sr)
			}

			if dg == 0xff {
				g = 0xff
			} else if sg == 0 {
				g = 0
			} else {
				g = 0xff - clip8((0xff-dg)*0xff/sg)
			}

			if db == 0xff {
				b = 0xff
			} else if sb == 0 {
				b = 0
			} else {
				b = 0xff - clip8((0xff-db)*0xff/sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorBurnRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if dr == 0xff {
				r = 0xff
			} else if sr == 0 {
				r = 0
			} else {
				r = 0xff - clip8((0xff-dr)*0xff/sr)
			}

			if dg == 0xff {
				g = 0xff
			} else if sg == 0 {
				g = 0
			} else {
				g = 0xff - clip8((0xff-dg)*0xff/sg)
			}

			if db == 0xff {
				b = 0xff
			} else if sb == 0 {
				b = 0
			} else {
				b = 0xff - clip8((0xff-db)*0xff/sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorBurnNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr == 0xff {
				r = 0xff
			} else if sr == 0 {
				r = 0
			} else {
				r = 0xff - clip8((0xff-dr)*0xff/sr)
			}

			if dg == 0xff {
				g = 0xff
			} else if sg == 0 {
				g = 0
			} else {
				g = 0xff - clip8((0xff-dg)*0xff/sg)
			}

			if db == 0xff {
				b = 0xff
			} else if sb == 0 {
				b = 0
			} else {
				b = 0xff - clip8((0xff-db)*0xff/sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorBurnRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr == 0xff {
				r = 0xff
			} else if sr == 0 {
				r = 0
			} else {
				r = 0xff - clip8((0xff-dr)*0xff/sr)
			}

			if dg == 0xff {
				g = 0xff
			} else if sg == 0 {
				g = 0
			} else {
				g = 0xff - clip8((0xff-dg)*0xff/sg)
			}

			if db == 0xff {
				b = 0xff
			} else if sb == 0 {
				b = 0
			} else {
				b = 0xff - clip8((0xff-db)*0xff/sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorBurnFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if dr == 0xffff {
				r = 0xffff
			} else if sr == 0 {
				r = 0
			} else {
				r = 0xffff - clip16((0xffff-dr)*0xffff/sr)
			}

			if dg == 0xffff {
				g = 0xffff
			} else if sg == 0 {
				g = 0
			} else {
				g = 0xffff - clip16((0xffff-dg)*0xffff/sg)
			}

			if db == 0xffff {
				b = 0xffff
			} else if sb == 0 {
				b = 0
			} else {
				b = 0xffff - clip16((0xffff-db)*0xffff/sb)
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d colorBurn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawColorBurnFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// linearBurn implements the linearBurn blend mode.
type linearBurn struct{}

// String implemenets fmt.Stringer interface.
func (d linearBurn) String() string {
	return "LinearBurn"
}

// Draw implements image.Drawer interface.
func (d linearBurn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d linearBurn) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d linearBurn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearBurnNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearBurn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearBurnRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearBurn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearBurnNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearBurn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearBurnRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawLinearBurnNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			tmp = dr + sr
			if tmp > 0xff {
				r = tmp - 0xff
			} else {
				r = 0
			}

			tmp = dg + sg
			if tmp > 0xff {
				g = tmp - 0xff
			} else {
				g = 0
			}

			tmp = db + sb
			if tmp > 0xff {
				b = tmp - 0xff
			} else {
				b = 0
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearBurnRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			tmp = dr + sr
			if tmp > 0xff {
				r = tmp - 0xff
			} else {
				r = 0
			}

			tmp = dg + sg
			if tmp > 0xff {
				g = tmp - 0xff
			} else {
				g = 0
			}

			tmp = db + sb
			if tmp > 0xff {
				b = tmp - 0xff
			} else {
				b = 0
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearBurnNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			tmp = dr + sr
			if tmp > 0xff {
				r = tmp - 0xff
			} else {
				r = 0
			}

			tmp = dg + sg
			if tmp > 0xff {
				g = tmp - 0xff
			} else {
				g = 0
			}

			tmp = db + sb
			if tmp > 0xff {
				b = tmp - 0xff
			} else {
				b = 0
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearBurnRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			tmp = dr + sr
			if tmp > 0xff {
				r = tmp - 0xff
			} else {
				r = 0
			}

			tmp = dg + sg
			if tmp > 0xff {
				g = tmp - 0xff
			} else {
				g = 0
			}

			tmp = db + sb
			if tmp > 0xff {
				b = tmp - 0xff
			} else {
				b = 0
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearBurnFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			tmp = dr + sr
			if tmp > 0xffff {
				r = tmp - 0xffff
			} else {
				r = 0
			}

			tmp = dg + sg
			if tmp > 0xffff {
				g = tmp - 0xffff
			} else {
				g = 0
			}

			tmp = db + sb
			if tmp > 0xffff {
				b = tmp - 0xffff
			} else {
				b = 0
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d linearBurn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawLinearBurnFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// darkerColor implements the darkerColor blend mode.
type darkerColor struct{}

// String implemenets fmt.Stringer interface.
func (d darkerColor) String() string {
	return "DarkerColor"
}

// Draw implements image.Drawer interface.
func (d darkerColor) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d darkerColor) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d darkerColor) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkerColorNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d darkerColor) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkerColorRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d darkerColor) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkerColorNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d darkerColor) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDarkerColorRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawDarkerColorNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkerColorRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkerColorNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkerColorRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDarkerColorFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if lum16(sr, sg, sb) < lum16(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d darkerColor) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawDarkerColorFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// lighten implements the lighten blend mode.
type lighten struct{}

// String implemenets fmt.Stringer interface.
func (d lighten) String() string {
	return "Lighten"
}

// Draw implements image.Drawer interface.
func (d lighten) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d lighten) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d lighten) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLightenNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d lighten) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLightenRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d lighten) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLightenNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d lighten) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLightenRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawLightenNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr > dr {
				r = sr
			} else {
				r = dr
			}

			if sg > dg {
				g = sg
			} else {
				g = dg
			}

			if sb > db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLightenRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr > dr {
				r = sr
			} else {
				r = dr
			}

			if sg > dg {
				g = sg
			} else {
				g = dg
			}

			if sb > db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLightenNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr > dr {
				r = sr
			} else {
				r = dr
			}

			if sg > dg {
				g = sg
			} else {
				g = dg
			}

			if sb > db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLightenRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr > dr {
				r = sr
			} else {
				r = dr
			}

			if sg > dg {
				g = sg
			} else {
				g = dg
			}

			if sb > db {
				b = sb
			} else {
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLightenFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr > dr {
				r = sr
			} else {
				r = dr
			}

			if sg > dg {
				g = sg
			} else {
				g = dg
			}

			if sb > db {
				b = sb
			} else {
				b = db
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d lighten) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawLightenFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// screen implements the screen blend mode.
type screen struct{}

// String implemenets fmt.Stringer interface.
func (d screen) String() string {
	return "Screen"
}

// Draw implements image.Drawer interface.
func (d screen) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d screen) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d screen) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawScreenNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d screen) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawScreenRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d screen) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawScreenNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d screen) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawScreenRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawScreenNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawScreenRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawScreenNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawScreenRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawScreenFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr / 0xffff)

			g = sg + dg - (sg * dg / 0xffff)

			b = sb + db - (sb * db / 0xffff)

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d screen) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawScreenFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// colorDodge implements the colorDodge blend mode.
type colorDodge struct{}

// String implemenets fmt.Stringer interface.
func (d colorDodge) String() string {
	return "ColorDodge"
}

// Draw implements image.Drawer interface.
func (d colorDodge) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d colorDodge) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d colorDodge) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorDodgeNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d colorDodge) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorDodgeRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d colorDodge) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorDodgeNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d colorDodge) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorDodgeRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawColorDodgeNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0xff {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / (0xff - sr))
			}

			if dg == 0 {
				g = 0
			} else if sg == 0xff {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / (0xff - sg))
			}

			if db == 0 {
				b = 0
			} else if sb == 0xff {
				b = 0xff
			} else {
				b = clip8(db * 0xff / (0xff - sb))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorDodgeRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0xff {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / (0xff - sr))
			}

			if dg == 0 {
				g = 0
			} else if sg == 0xff {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / (0xff - sg))
			}

			if db == 0 {
				b = 0
			} else if sb == 0xff {
				b = 0xff
			} else {
				b = clip8(db * 0xff / (0xff - sb))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorDodgeNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0xff {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / (0xff - sr))
			}

			if dg == 0 {
				g = 0
			} else if sg == 0xff {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / (0xff - sg))
			}

			if db == 0 {
				b = 0
			} else if sb == 0xff {
				b = 0xff
			} else {
				b = clip8(db * 0xff / (0xff - sb))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorDodgeRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0xff {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / (0xff - sr))
			}

			if dg == 0 {
				g = 0
			} else if sg == 0xff {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / (0xff - sg))
			}

			if db == 0 {
				b = 0
			} else if sb == 0xff {
				b = 0xff
			} else {
				b = clip8(db * 0xff / (0xff - sb))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorDodgeFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0xffff {
				r = 0xffff
			} else {
				r = clip16(dr * 0xffff / (0xffff - sr))
			}

			if dg == 0 {
				g = 0
			} else if sg == 0xffff {
				g = 0xffff
			} else {
				g = clip16(dg * 0xffff / (0xffff - sg))
			}

			if db == 0 {
				b = 0
			} else if sb == 0xffff {
				b = 0xffff
			} else {
				b = clip16(db * 0xffff / (0xffff - sb))
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d colorDodge) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawColorDodgeFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// linearDodge implements the linearDodge blend mode.
type linearDodge struct{}

// String implemenets fmt.Stringer interface.
func (d linearDodge) String() string {
	return "LinearDodge"
}

// Draw implements image.Drawer interface.
func (d linearDodge) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d linearDodge) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d linearDodge) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearDodgeNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearDodge) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearDodgeRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearDodge) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearDodgeNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearDodge) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearDodgeRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawLinearDodgeNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearDodgeRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearDodgeNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearDodgeRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearDodgeFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r = clip16(sr + dr)

			g = clip16(sg + dg)

			b = clip16(sb + db)

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d linearDodge) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawLinearDodgeFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// lighterColor implements the lighterColor blend mode.
type lighterColor struct{}

// String implemenets fmt.Stringer interface.
func (d lighterColor) String() string {
	return "LighterColor"
}

// Draw implements image.Drawer interface.
func (d lighterColor) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d lighterColor) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d lighterColor) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLighterColorNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d lighterColor) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLighterColorRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d lighterColor) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLighterColorNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d lighterColor) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLighterColorRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawLighterColorNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLighterColorRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLighterColorNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLighterColorRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLighterColorFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if lum16(sr, sg, sb) > lum16(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d lighterColor) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawLighterColorFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// add implements the add blend mode.
type add struct{}

// String implemenets fmt.Stringer interface.
func (d add) String() string {
	return "Add"
}

// Draw implements image.Drawer interface.
func (d add) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d add) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d add) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawAddNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d add) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawAddRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d add) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawAddNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d add) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawAddRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawAddNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[j+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[j+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawAddRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[j+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[j+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawAddNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(clip8((b*a1+sb*a2+db*a3)/a) * a * 32897 >> 23)
			dpix[j+1] = uint8(clip8((g*a1+sg*a2+dg*a3)/a) * a * 32897 >> 23)
			dpix[j+0] = uint8(clip8((r*a1+sr*a2+dr*a3)/a) * a * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawAddRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(clip8((b*a1+sb*a2+db*a3)/a) * a * 32897 >> 23)
			dpix[j+1] = uint8(clip8((g*a1+sg*a2+dg*a3)/a) * a * 32897 >> 23)
			dpix[j+0] = uint8(clip8((r*a1+sr*a2+dr*a3)/a) * a * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawAddFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			out.R = uint16(clip6416(uint64(sr*a2+dr*a3)/0xffff+uint64(r)*uint64(a1)/uint64(a)) * a / 0xffff)
			out.G = uint16(clip6416(uint64(sg*a2+dg*a3)/0xffff+uint64(g)*uint64(a1)/uint64(a)) * a / 0xffff)
			out.B = uint16(clip6416(uint64(sb*a2+db*a3)/0xffff+uint64(b)*uint64(a1)/uint64(a)) * a / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d add) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawAddFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// overlay implements the overlay blend mode.
type overlay struct{}

// String implemenets fmt.Stringer interface.
func (d overlay) String() string {
	return "Overlay"
}

// Draw implements image.Drawer interface.
func (d overlay) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d overlay) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d overlay) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawOverlayNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d overlay) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawOverlayRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d overlay) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawOverlayNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d overlay) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawOverlayRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawOverlayNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if dr < 0x80 {
				r = sr * dr * 32897 >> 22
			} else {
				r = 0xff - ((0xff - ((dr - 0x80) << 1)) * (0xff - sr) * 32897 >> 23)
			}

			if dg < 0x80 {
				g = sg * dg * 32897 >> 22
			} else {
				g = 0xff - ((0xff - ((dg - 0x80) << 1)) * (0xff - sg) * 32897 >> 23)
			}

			if db < 0x80 {
				b = sb * db * 32897 >> 22
			} else {
				b = 0xff - ((0xff - ((db - 0x80) << 1)) * (0xff - sb) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawOverlayRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if dr < 0x80 {
				r = sr * dr * 32897 >> 22
			} else {
				r = 0xff - ((0xff - ((dr - 0x80) << 1)) * (0xff - sr) * 32897 >> 23)
			}

			if dg < 0x80 {
				g = sg * dg * 32897 >> 22
			} else {
				g = 0xff - ((0xff - ((dg - 0x80) << 1)) * (0xff - sg) * 32897 >> 23)
			}

			if db < 0x80 {
				b = sb * db * 32897 >> 22
			} else {
				b = 0xff - ((0xff - ((db - 0x80) << 1)) * (0xff - sb) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawOverlayNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr < 0x80 {
				r = sr * dr * 32897 >> 22
			} else {
				r = 0xff - ((0xff - ((dr - 0x80) << 1)) * (0xff - sr) * 32897 >> 23)
			}

			if dg < 0x80 {
				g = sg * dg * 32897 >> 22
			} else {
				g = 0xff - ((0xff - ((dg - 0x80) << 1)) * (0xff - sg) * 32897 >> 23)
			}

			if db < 0x80 {
				b = sb * db * 32897 >> 22
			} else {
				b = 0xff - ((0xff - ((db - 0x80) << 1)) * (0xff - sb) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawOverlayRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr < 0x80 {
				r = sr * dr * 32897 >> 22
			} else {
				r = 0xff - ((0xff - ((dr - 0x80) << 1)) * (0xff - sr) * 32897 >> 23)
			}

			if dg < 0x80 {
				g = sg * dg * 32897 >> 22
			} else {
				g = 0xff - ((0xff - ((dg - 0x80) << 1)) * (0xff - sg) * 32897 >> 23)
			}

			if db < 0x80 {
				b = sb * db * 32897 >> 22
			} else {
				b = 0xff - ((0xff - ((db - 0x80) << 1)) * (0xff - sb) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawOverlayFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if dr < 0x8000 {
				r = sr * dr / 0x8000
			} else {
				r = 0xffff - ((0xffff - ((dr - 0x8000) << 1)) * (0xffff - sr) / 0xffff)
			}

			if dg < 0x8000 {
				g = sg * dg / 0x8000
			} else {
				g = 0xffff - ((0xffff - ((dg - 0x8000) << 1)) * (0xffff - sg) / 0xffff)
			}

			if db < 0x8000 {
				b = sb * db / 0x8000
			} else {
				b = 0xffff - ((0xffff - ((db - 0x8000) << 1)) * (0xffff - sb) / 0xffff)
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d overlay) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawOverlayFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// softLight implements the softLight blend mode.
type softLight struct{}

// String implemenets fmt.Stringer interface.
func (d softLight) String() string {
	return "SoftLight"
}

// Draw implements image.Drawer interface.
func (d softLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d softLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d softLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSoftLightNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d softLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSoftLightRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d softLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSoftLightNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d softLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSoftLightRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawSoftLightNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr - (((0xff - (sr << 1)) * dr * 32897 >> 23) * (0xff - dr) * 32897 >> 23)
			} else {
				if dr < 0x40 {
					tmp = uint32((((int32(dr)<<4-0xff*12)*32897>>23)*int32(dr) + 0xff*4) * int32(dr) * 32897 >> 23)
				} else {
					tmp = sqrt(dr, 0xff)
				}
				r = dr + (((sr << 1) - 0xff) * (tmp - dr) * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg - (((0xff - (sg << 1)) * dg * 32897 >> 23) * (0xff - dg) * 32897 >> 23)
			} else {
				if dg < 0x40 {
					tmp = uint32((((int32(dg)<<4-0xff*12)*32897>>23)*int32(dg) + 0xff*4) * int32(dg) * 32897 >> 23)
				} else {
					tmp = sqrt(dg, 0xff)
				}
				g = dg + (((sg << 1) - 0xff) * (tmp - dg) * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db - (((0xff - (sb << 1)) * db * 32897 >> 23) * (0xff - db) * 32897 >> 23)
			} else {
				if db < 0x40 {
					tmp = uint32((((int32(db)<<4-0xff*12)*32897>>23)*int32(db) + 0xff*4) * int32(db) * 32897 >> 23)
				} else {
					tmp = sqrt(db, 0xff)
				}
				b = db + (((sb << 1) - 0xff) * (tmp - db) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSoftLightRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr - (((0xff - (sr << 1)) * dr * 32897 >> 23) * (0xff - dr) * 32897 >> 23)
			} else {
				if dr < 0x40 {
					tmp = uint32((((int32(dr)<<4-0xff*12)*32897>>23)*int32(dr) + 0xff*4) * int32(dr) * 32897 >> 23)
				} else {
					tmp = sqrt(dr, 0xff)
				}
				r = dr + (((sr << 1) - 0xff) * (tmp - dr) * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg - (((0xff - (sg << 1)) * dg * 32897 >> 23) * (0xff - dg) * 32897 >> 23)
			} else {
				if dg < 0x40 {
					tmp = uint32((((int32(dg)<<4-0xff*12)*32897>>23)*int32(dg) + 0xff*4) * int32(dg) * 32897 >> 23)
				} else {
					tmp = sqrt(dg, 0xff)
				}
				g = dg + (((sg << 1) - 0xff) * (tmp - dg) * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db - (((0xff - (sb << 1)) * db * 32897 >> 23) * (0xff - db) * 32897 >> 23)
			} else {
				if db < 0x40 {
					tmp = uint32((((int32(db)<<4-0xff*12)*32897>>23)*int32(db) + 0xff*4) * int32(db) * 32897 >> 23)
				} else {
					tmp = sqrt(db, 0xff)
				}
				b = db + (((sb << 1) - 0xff) * (tmp - db) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSoftLightNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr - (((0xff - (sr << 1)) * dr * 32897 >> 23) * (0xff - dr) * 32897 >> 23)
			} else {
				if dr < 0x40 {
					tmp = uint32((((int32(dr)<<4-0xff*12)*32897>>23)*int32(dr) + 0xff*4) * int32(dr) * 32897 >> 23)
				} else {
					tmp = sqrt(dr, 0xff)
				}
				r = dr + (((sr << 1) - 0xff) * (tmp - dr) * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg - (((0xff - (sg << 1)) * dg * 32897 >> 23) * (0xff - dg) * 32897 >> 23)
			} else {
				if dg < 0x40 {
					tmp = uint32((((int32(dg)<<4-0xff*12)*32897>>23)*int32(dg) + 0xff*4) * int32(dg) * 32897 >> 23)
				} else {
					tmp = sqrt(dg, 0xff)
				}
				g = dg + (((sg << 1) - 0xff) * (tmp - dg) * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db - (((0xff - (sb << 1)) * db * 32897 >> 23) * (0xff - db) * 32897 >> 23)
			} else {
				if db < 0x40 {
					tmp = uint32((((int32(db)<<4-0xff*12)*32897>>23)*int32(db) + 0xff*4) * int32(db) * 32897 >> 23)
				} else {
					tmp = sqrt(db, 0xff)
				}
				b = db + (((sb << 1) - 0xff) * (tmp - db) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSoftLightRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr - (((0xff - (sr << 1)) * dr * 32897 >> 23) * (0xff - dr) * 32897 >> 23)
			} else {
				if dr < 0x40 {
					tmp = uint32((((int32(dr)<<4-0xff*12)*32897>>23)*int32(dr) + 0xff*4) * int32(dr) * 32897 >> 23)
				} else {
					tmp = sqrt(dr, 0xff)
				}
				r = dr + (((sr << 1) - 0xff) * (tmp - dr) * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg - (((0xff - (sg << 1)) * dg * 32897 >> 23) * (0xff - dg) * 32897 >> 23)
			} else {
				if dg < 0x40 {
					tmp = uint32((((int32(dg)<<4-0xff*12)*32897>>23)*int32(dg) + 0xff*4) * int32(dg) * 32897 >> 23)
				} else {
					tmp = sqrt(dg, 0xff)
				}
				g = dg + (((sg << 1) - 0xff) * (tmp - dg) * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db - (((0xff - (sb << 1)) * db * 32897 >> 23) * (0xff - db) * 32897 >> 23)
			} else {
				if db < 0x40 {
					tmp = uint32((((int32(db)<<4-0xff*12)*32897>>23)*int32(db) + 0xff*4) * int32(db) * 32897 >> 23)
				} else {
					tmp = sqrt(db, 0xff)
				}
				b = db + (((sb << 1) - 0xff) * (tmp - db) * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSoftLightFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < 0x8000 {
				r = dr - (((0xffff - (sr << 1)) * dr / 0xffff) * (0xffff - dr) / 0xffff)
			} else {
				if dr < 0x4000 {
					tmp = uint32((((int32(dr)<<4-0xffff*12)/0xffff)*int32(dr) + 0xffff*4) * int32(dr) / 0xffff)
				} else {
					tmp = sqrt(dr, 0xffff)
				}
				r = dr + (((sr << 1) - 0xffff) * (tmp - dr) / 0xffff)
			}

			if sg < 0x8000 {
				g = dg - (((0xffff - (sg << 1)) * dg / 0xffff) * (0xffff - dg) / 0xffff)
			} else {
				if dg < 0x4000 {
					tmp = uint32((((int32(dg)<<4-0xffff*12)/0xffff)*int32(dg) + 0xffff*4) * int32(dg) / 0xffff)
				} else {
					tmp = sqrt(dg, 0xffff)
				}
				g = dg + (((sg << 1) - 0xffff) * (tmp - dg) / 0xffff)
			}

			if sb < 0x8000 {
				b = db - (((0xffff - (sb << 1)) * db / 0xffff) * (0xffff - db) / 0xffff)
			} else {
				if db < 0x4000 {
					tmp = uint32((((int32(db)<<4-0xffff*12)/0xffff)*int32(db) + 0xffff*4) * int32(db) / 0xffff)
				} else {
					tmp = sqrt(db, 0xffff)
				}
				b = db + (((sb << 1) - 0xffff) * (tmp - db) / 0xffff)
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d softLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawSoftLightFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// hardLight implements the hardLight blend mode.
type hardLight struct{}

// String implemenets fmt.Stringer interface.
func (d hardLight) String() string {
	return "HardLight"
}

// Draw implements image.Drawer interface.
func (d hardLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d hardLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d hardLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardLightNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hardLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardLightRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hardLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardLightNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hardLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardLightRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawHardLightNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr * sr * 32897 >> 22
			} else {
				tmp = (sr << 1) - 0xff
				r = dr + tmp - (dr * tmp * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg * sg * 32897 >> 22
			} else {
				tmp = (sg << 1) - 0xff
				g = dg + tmp - (dg * tmp * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db * sb * 32897 >> 22
			} else {
				tmp = (sb << 1) - 0xff
				b = db + tmp - (db * tmp * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardLightRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr * sr * 32897 >> 22
			} else {
				tmp = (sr << 1) - 0xff
				r = dr + tmp - (dr * tmp * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg * sg * 32897 >> 22
			} else {
				tmp = (sg << 1) - 0xff
				g = dg + tmp - (dg * tmp * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db * sb * 32897 >> 22
			} else {
				tmp = (sb << 1) - 0xff
				b = db + tmp - (db * tmp * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardLightNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr * sr * 32897 >> 22
			} else {
				tmp = (sr << 1) - 0xff
				r = dr + tmp - (dr * tmp * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg * sg * 32897 >> 22
			} else {
				tmp = (sg << 1) - 0xff
				g = dg + tmp - (dg * tmp * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db * sb * 32897 >> 22
			} else {
				tmp = (sb << 1) - 0xff
				b = db + tmp - (db * tmp * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardLightRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = dr * sr * 32897 >> 22
			} else {
				tmp = (sr << 1) - 0xff
				r = dr + tmp - (dr * tmp * 32897 >> 23)
			}

			if sg < 0x80 {
				g = dg * sg * 32897 >> 22
			} else {
				tmp = (sg << 1) - 0xff
				g = dg + tmp - (dg * tmp * 32897 >> 23)
			}

			if sb < 0x80 {
				b = db * sb * 32897 >> 22
			} else {
				tmp = (sb << 1) - 0xff
				b = db + tmp - (db * tmp * 32897 >> 23)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardLightFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < 0x8000 {
				r = dr * sr / 0x8000
			} else {
				tmp = (sr << 1) - 0xffff
				r = dr + tmp - (dr * tmp / 0xffff)
			}

			if sg < 0x8000 {
				g = dg * sg / 0x8000
			} else {
				tmp = (sg << 1) - 0xffff
				g = dg + tmp - (dg * tmp / 0xffff)
			}

			if sb < 0x8000 {
				b = db * sb / 0x8000
			} else {
				tmp = (sb << 1) - 0xffff
				b = db + tmp - (db * tmp / 0xffff)
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d hardLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawHardLightFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// linearLight implements the linearLight blend mode.
type linearLight struct{}

// String implemenets fmt.Stringer interface.
func (d linearLight) String() string {
	return "LinearLight"
}

// Draw implements image.Drawer interface.
func (d linearLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d linearLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d linearLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearLightNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearLightRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearLightNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d linearLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLinearLightRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawLinearLightNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = clip0(dr + (sr << 1) - 0xff)
			} else {
				r = clip8(dr + ((sr - 0x80) << 1))
			}

			if sg < 0x80 {
				g = clip0(dg + (sg << 1) - 0xff)
			} else {
				g = clip8(dg + ((sg - 0x80) << 1))
			}

			if sb < 0x80 {
				b = clip0(db + (sb << 1) - 0xff)
			} else {
				b = clip8(db + ((sb - 0x80) << 1))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearLightRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = clip0(dr + (sr << 1) - 0xff)
			} else {
				r = clip8(dr + ((sr - 0x80) << 1))
			}

			if sg < 0x80 {
				g = clip0(dg + (sg << 1) - 0xff)
			} else {
				g = clip8(dg + ((sg - 0x80) << 1))
			}

			if sb < 0x80 {
				b = clip0(db + (sb << 1) - 0xff)
			} else {
				b = clip8(db + ((sb - 0x80) << 1))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearLightNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = clip0(dr + (sr << 1) - 0xff)
			} else {
				r = clip8(dr + ((sr - 0x80) << 1))
			}

			if sg < 0x80 {
				g = clip0(dg + (sg << 1) - 0xff)
			} else {
				g = clip8(dg + ((sg - 0x80) << 1))
			}

			if sb < 0x80 {
				b = clip0(db + (sb << 1) - 0xff)
			} else {
				b = clip8(db + ((sb - 0x80) << 1))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearLightRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				r = clip0(dr + (sr << 1) - 0xff)
			} else {
				r = clip8(dr + ((sr - 0x80) << 1))
			}

			if sg < 0x80 {
				g = clip0(dg + (sg << 1) - 0xff)
			} else {
				g = clip8(dg + ((sg - 0x80) << 1))
			}

			if sb < 0x80 {
				b = clip0(db + (sb << 1) - 0xff)
			} else {
				b = clip8(db + ((sb - 0x80) << 1))
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLinearLightFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < 0x8000 {
				r = clip0(dr + (sr << 1) - 0xffff)
			} else {
				r = clip16(dr + ((sr - 0x8000) << 1))
			}

			if sg < 0x8000 {
				g = clip0(dg + (sg << 1) - 0xffff)
			} else {
				g = clip16(dg + ((sg - 0x8000) << 1))
			}

			if sb < 0x8000 {
				b = clip0(db + (sb << 1) - 0xffff)
			} else {
				b = clip16(db + ((sb - 0x8000) << 1))
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d linearLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawLinearLightFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// vividLight implements the vividLight blend mode.
type vividLight struct{}

// String implemenets fmt.Stringer interface.
func (d vividLight) String() string {
	return "VividLight"
}

// Draw implements image.Drawer interface.
func (d vividLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d vividLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d vividLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawVividLightNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d vividLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawVividLightRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d vividLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawVividLightNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d vividLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawVividLightRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawVividLightNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp == 0 {
					r = 0
				} else {
					r = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = (sr << 1) - 0xff
				if tmp == 0xff {
					r = 0xff
				} else {
					r = clip8((dr * 0xff) / (0xff - tmp))
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp == 0 {
					g = 0
				} else {
					g = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = (sg << 1) - 0xff
				if tmp == 0xff {
					g = 0xff
				} else {
					g = clip8((dg * 0xff) / (0xff - tmp))
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp == 0 {
					b = 0
				} else {
					b = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = (sb << 1) - 0xff
				if tmp == 0xff {
					b = 0xff
				} else {
					b = clip8((db * 0xff) / (0xff - tmp))
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawVividLightRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp == 0 {
					r = 0
				} else {
					r = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = (sr << 1) - 0xff
				if tmp == 0xff {
					r = 0xff
				} else {
					r = clip8((dr * 0xff) / (0xff - tmp))
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp == 0 {
					g = 0
				} else {
					g = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = (sg << 1) - 0xff
				if tmp == 0xff {
					g = 0xff
				} else {
					g = clip8((dg * 0xff) / (0xff - tmp))
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp == 0 {
					b = 0
				} else {
					b = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = (sb << 1) - 0xff
				if tmp == 0xff {
					b = 0xff
				} else {
					b = clip8((db * 0xff) / (0xff - tmp))
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawVividLightNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp == 0 {
					r = 0
				} else {
					r = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = (sr << 1) - 0xff
				if tmp == 0xff {
					r = 0xff
				} else {
					r = clip8((dr * 0xff) / (0xff - tmp))
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp == 0 {
					g = 0
				} else {
					g = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = (sg << 1) - 0xff
				if tmp == 0xff {
					g = 0xff
				} else {
					g = clip8((dg * 0xff) / (0xff - tmp))
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp == 0 {
					b = 0
				} else {
					b = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = (sb << 1) - 0xff
				if tmp == 0xff {
					b = 0xff
				} else {
					b = clip8((db * 0xff) / (0xff - tmp))
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawVividLightRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp == 0 {
					r = 0
				} else {
					r = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = (sr << 1) - 0xff
				if tmp == 0xff {
					r = 0xff
				} else {
					r = clip8((dr * 0xff) / (0xff - tmp))
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp == 0 {
					g = 0
				} else {
					g = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = (sg << 1) - 0xff
				if tmp == 0xff {
					g = 0xff
				} else {
					g = clip8((dg * 0xff) / (0xff - tmp))
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp == 0 {
					b = 0
				} else {
					b = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = (sb << 1) - 0xff
				if tmp == 0xff {
					b = 0xff
				} else {
					b = clip8((db * 0xff) / (0xff - tmp))
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawVividLightFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < 0x8000 {
				tmp = sr << 1
				if tmp == 0 {
					r = 0
				} else {
					r = 0xffff - clip16((0xffff-dr)*0xffff/tmp)
				}
			} else {
				tmp = (sr << 1) - 0xffff
				if tmp == 0xffff {
					r = 0xffff
				} else {
					r = clip16((dr * 0xffff) / (0xffff - tmp))
				}
			}

			if sg < 0x8000 {
				tmp = sg << 1
				if tmp == 0 {
					g = 0
				} else {
					g = 0xffff - clip16((0xffff-dg)*0xffff/tmp)
				}
			} else {
				tmp = (sg << 1) - 0xffff
				if tmp == 0xffff {
					g = 0xffff
				} else {
					g = clip16((dg * 0xffff) / (0xffff - tmp))
				}
			}

			if sb < 0x8000 {
				tmp = sb << 1
				if tmp == 0 {
					b = 0
				} else {
					b = 0xffff - clip16((0xffff-db)*0xffff/tmp)
				}
			} else {
				tmp = (sb << 1) - 0xffff
				if tmp == 0xffff {
					b = 0xffff
				} else {
					b = clip16((db * 0xffff) / (0xffff - tmp))
				}
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d vividLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawVividLightFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// pinLight implements the pinLight blend mode.
type pinLight struct{}

// String implemenets fmt.Stringer interface.
func (d pinLight) String() string {
	return "PinLight"
}

// Draw implements image.Drawer interface.
func (d pinLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d pinLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d pinLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawPinLightNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d pinLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawPinLightRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d pinLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawPinLightNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d pinLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawPinLightRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawPinLightNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp < dr {
					r = tmp
				} else {
					r = dr
				}
			} else {
				tmp = (sr - 0x80) << 1
				if tmp > dr {
					r = tmp
				} else {
					r = dr
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp < dg {
					g = tmp
				} else {
					g = dg
				}
			} else {
				tmp = (sg - 0x80) << 1
				if tmp > dg {
					g = tmp
				} else {
					g = dg
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp < db {
					b = tmp
				} else {
					b = db
				}
			} else {
				tmp = (sb - 0x80) << 1
				if tmp > db {
					b = tmp
				} else {
					b = db
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawPinLightRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp < dr {
					r = tmp
				} else {
					r = dr
				}
			} else {
				tmp = (sr - 0x80) << 1
				if tmp > dr {
					r = tmp
				} else {
					r = dr
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp < dg {
					g = tmp
				} else {
					g = dg
				}
			} else {
				tmp = (sg - 0x80) << 1
				if tmp > dg {
					g = tmp
				} else {
					g = dg
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp < db {
					b = tmp
				} else {
					b = db
				}
			} else {
				tmp = (sb - 0x80) << 1
				if tmp > db {
					b = tmp
				} else {
					b = db
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawPinLightNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp < dr {
					r = tmp
				} else {
					r = dr
				}
			} else {
				tmp = (sr - 0x80) << 1
				if tmp > dr {
					r = tmp
				} else {
					r = dr
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp < dg {
					g = tmp
				} else {
					g = dg
				}
			} else {
				tmp = (sg - 0x80) << 1
				if tmp > dg {
					g = tmp
				} else {
					g = dg
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp < db {
					b = tmp
				} else {
					b = db
				}
			} else {
				tmp = (sb - 0x80) << 1
				if tmp > db {
					b = tmp
				} else {
					b = db
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawPinLightRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if tmp < dr {
					r = tmp
				} else {
					r = dr
				}
			} else {
				tmp = (sr - 0x80) << 1
				if tmp > dr {
					r = tmp
				} else {
					r = dr
				}
			}

			if sg < 0x80 {
				tmp = sg << 1
				if tmp < dg {
					g = tmp
				} else {
					g = dg
				}
			} else {
				tmp = (sg - 0x80) << 1
				if tmp > dg {
					g = tmp
				} else {
					g = dg
				}
			}

			if sb < 0x80 {
				tmp = sb << 1
				if tmp < db {
					b = tmp
				} else {
					b = db
				}
			} else {
				tmp = (sb - 0x80) << 1
				if tmp > db {
					b = tmp
				} else {
					b = db
				}
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawPinLightFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < 0x8000 {
				tmp = sr << 1
				if tmp < dr {
					r = tmp
				} else {
					r = dr
				}
			} else {
				tmp = (sr - 0x8000) << 1
				if tmp > dr {
					r = tmp
				} else {
					r = dr
				}
			}

			if sg < 0x8000 {
				tmp = sg << 1
				if tmp < dg {
					g = tmp
				} else {
					g = dg
				}
			} else {
				tmp = (sg - 0x8000) << 1
				if tmp > dg {
					g = tmp
				} else {
					g = dg
				}
			}

			if sb < 0x8000 {
				tmp = sb << 1
				if tmp < db {
					b = tmp
				} else {
					b = db
				}
			} else {
				tmp = (sb - 0x8000) << 1
				if tmp > db {
					b = tmp
				} else {
					b = db
				}
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d pinLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawPinLightFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// hardMix implements the hardMix blend mode.
type hardMix struct{}

// String implemenets fmt.Stringer interface.
func (d hardMix) String() string {
	return "HardMix"
}

// Draw implements image.Drawer interface.
func (d hardMix) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d hardMix) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d hardMix) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardMixNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hardMix) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardMixRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hardMix) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardMixNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hardMix) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHardMixRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawHardMixNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if dr == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = ((sr - 0x80) << 1)
				if dr == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dr * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				r = 0
			} else {
				r = 0xff
			}

			if sg < 0x80 {
				tmp = sg << 1
				if dg == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = ((sg - 0x80) << 1)
				if dg == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dg * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				g = 0
			} else {
				g = 0xff
			}

			if sb < 0x80 {
				tmp = sb << 1
				if db == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = ((sb - 0x80) << 1)
				if db == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(db * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				b = 0
			} else {
				b = 0xff
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardMixRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if dr == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = ((sr - 0x80) << 1)
				if dr == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dr * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				r = 0
			} else {
				r = 0xff
			}

			if sg < 0x80 {
				tmp = sg << 1
				if dg == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = ((sg - 0x80) << 1)
				if dg == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dg * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				g = 0
			} else {
				g = 0xff
			}

			if sb < 0x80 {
				tmp = sb << 1
				if db == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = ((sb - 0x80) << 1)
				if db == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(db * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				b = 0
			} else {
				b = 0xff
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardMixNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if dr == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = ((sr - 0x80) << 1)
				if dr == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dr * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				r = 0
			} else {
				r = 0xff
			}

			if sg < 0x80 {
				tmp = sg << 1
				if dg == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = ((sg - 0x80) << 1)
				if dg == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dg * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				g = 0
			} else {
				g = 0xff
			}

			if sb < 0x80 {
				tmp = sb << 1
				if db == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = ((sb - 0x80) << 1)
				if db == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(db * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				b = 0
			} else {
				b = 0xff
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardMixRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if sr < 0x80 {
				tmp = sr << 1
				if dr == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dr)*0xff/tmp)
				}
			} else {
				tmp = ((sr - 0x80) << 1)
				if dr == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dr * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				r = 0
			} else {
				r = 0xff
			}

			if sg < 0x80 {
				tmp = sg << 1
				if dg == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-dg)*0xff/tmp)
				}
			} else {
				tmp = ((sg - 0x80) << 1)
				if dg == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(dg * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				g = 0
			} else {
				g = 0xff
			}

			if sb < 0x80 {
				tmp = sb << 1
				if db == 0xff {
					tmp = 0xff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xff - clip8((0xff-db)*0xff/tmp)
				}
			} else {
				tmp = ((sb - 0x80) << 1)
				if db == 0 {
					tmp = 0
				} else if tmp == 0xff {
					tmp = 0xff
				} else {
					tmp = clip8(db * 0xff / (0xff - tmp))
				}
			}
			if tmp <= 0x80 {
				b = 0
			} else {
				b = 0xff
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHardMixFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if sr < 0x8000 {
				tmp = sr << 1
				if dr == 0xffff {
					tmp = 0xffff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xffff - clip16((0xffff-dr)*0xffff/tmp)
				}
			} else {
				tmp = ((sr - 0x8000) << 1)
				if dr == 0 {
					tmp = 0
				} else if tmp == 0xffff {
					tmp = 0xffff
				} else {
					tmp = clip16(dr * 0xffff / (0xffff - tmp))
				}
			}
			if tmp <= 0x8000 {
				r = 0
			} else {
				r = 0xffff
			}

			if sg < 0x8000 {
				tmp = sg << 1
				if dg == 0xffff {
					tmp = 0xffff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xffff - clip16((0xffff-dg)*0xffff/tmp)
				}
			} else {
				tmp = ((sg - 0x8000) << 1)
				if dg == 0 {
					tmp = 0
				} else if tmp == 0xffff {
					tmp = 0xffff
				} else {
					tmp = clip16(dg * 0xffff / (0xffff - tmp))
				}
			}
			if tmp <= 0x8000 {
				g = 0
			} else {
				g = 0xffff
			}

			if sb < 0x8000 {
				tmp = sb << 1
				if db == 0xffff {
					tmp = 0xffff
				} else if tmp == 0 {
					tmp = 0
				} else {
					tmp = 0xffff - clip16((0xffff-db)*0xffff/tmp)
				}
			} else {
				tmp = ((sb - 0x8000) << 1)
				if db == 0 {
					tmp = 0
				} else if tmp == 0xffff {
					tmp = 0xffff
				} else {
					tmp = clip16(db * 0xffff / (0xffff - tmp))
				}
			}
			if tmp <= 0x8000 {
				b = 0
			} else {
				b = 0xffff
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d hardMix) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawHardMixFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// difference implements the difference blend mode.
type difference struct{}

// String implemenets fmt.Stringer interface.
func (d difference) String() string {
	return "Difference"
}

// Draw implements image.Drawer interface.
func (d difference) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d difference) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d difference) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDifferenceNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d difference) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDifferenceRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d difference) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDifferenceNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d difference) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDifferenceRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawDifferenceNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if dr < sr {
				r = sr - dr
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = sg - dg
			} else {
				g = dg - sg
			}

			if db < sb {
				b = sb - db
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDifferenceRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if dr < sr {
				r = sr - dr
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = sg - dg
			} else {
				g = dg - sg
			}

			if db < sb {
				b = sb - db
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDifferenceNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr < sr {
				r = sr - dr
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = sg - dg
			} else {
				g = dg - sg
			}

			if db < sb {
				b = sb - db
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDifferenceRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr < sr {
				r = sr - dr
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = sg - dg
			} else {
				g = dg - sg
			}

			if db < sb {
				b = sb - db
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDifferenceFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if dr < sr {
				r = sr - dr
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = sg - dg
			} else {
				g = dg - sg
			}

			if db < sb {
				b = sb - db
			} else {
				b = db - sb
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d difference) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawDifferenceFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// exclusion implements the exclusion blend mode.
type exclusion struct{}

// String implemenets fmt.Stringer interface.
func (d exclusion) String() string {
	return "Exclusion"
}

// Draw implements image.Drawer interface.
func (d exclusion) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d exclusion) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d exclusion) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawExclusionNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d exclusion) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawExclusionRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d exclusion) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawExclusionNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d exclusion) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawExclusionRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawExclusionNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawExclusionRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawExclusionNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawExclusionRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawExclusionFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr / 0x8000)

			g = dg + sg - (dg * sg / 0x8000)

			b = db + sb - (db * sb / 0x8000)

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d exclusion) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawExclusionFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// subtract implements the subtract blend mode.
type subtract struct{}

// String implemenets fmt.Stringer interface.
func (d subtract) String() string {
	return "Subtract"
}

// Draw implements image.Drawer interface.
func (d subtract) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d subtract) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d subtract) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSubtractNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d subtract) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSubtractRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d subtract) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSubtractNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d subtract) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSubtractRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawSubtractNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if dr < sr {
				r = 0
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = 0
			} else {
				g = dg - sg
			}

			if db < sb {
				b = 0
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSubtractRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if dr < sr {
				r = 0
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = 0
			} else {
				g = dg - sg
			}

			if db < sb {
				b = 0
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSubtractNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr < sr {
				r = 0
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = 0
			} else {
				g = dg - sg
			}

			if db < sb {
				b = 0
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSubtractRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr < sr {
				r = 0
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = 0
			} else {
				g = dg - sg
			}

			if db < sb {
				b = 0
			} else {
				b = db - sb
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSubtractFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if dr < sr {
				r = 0
			} else {
				r = dr - sr
			}

			if dg < sg {
				g = 0
			} else {
				g = dg - sg
			}

			if db < sb {
				b = 0
			} else {
				b = db - sb
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d subtract) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawSubtractFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// divide implements the divide blend mode.
type divide struct{}

// String implemenets fmt.Stringer interface.
func (d divide) String() string {
	return "Divide"
}

// Draw implements image.Drawer interface.
func (d divide) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d divide) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d divide) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDivideNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d divide) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDivideRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d divide) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDivideNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d divide) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawDivideRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawDivideNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0 {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / sr)
			}

			if dg == 0 {
				g = 0
			} else if sg == 0 {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / sg)
			}

			if db == 0 {
				b = 0
			} else if sb == 0 {
				b = 0xff
			} else {
				b = clip8(db * 0xff / sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDivideRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0 {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / sr)
			}

			if dg == 0 {
				g = 0
			} else if sg == 0 {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / sg)
			}

			if db == 0 {
				b = 0
			} else if sb == 0 {
				b = 0xff
			} else {
				b = clip8(db * 0xff / sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDivideNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0 {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / sr)
			}

			if dg == 0 {
				g = 0
			} else if sg == 0 {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / sg)
			}

			if db == 0 {
				b = 0
			} else if sb == 0 {
				b = 0xff
			} else {
				b = clip8(db * 0xff / sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDivideRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0 {
				r = 0xff
			} else {
				r = clip8(dr * 0xff / sr)
			}

			if dg == 0 {
				g = 0
			} else if sg == 0 {
				g = 0xff
			} else {
				g = clip8(dg * 0xff / sg)
			}

			if db == 0 {
				b = 0
			} else if sb == 0 {
				b = 0xff
			} else {
				b = clip8(db * 0xff / sb)
			}

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawDivideFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			if dr == 0 {
				r = 0
			} else if sr == 0 {
				r = 0xffff
			} else {
				r = clip16(dr * 0xffff / sr)
			}

			if dg == 0 {
				g = 0
			} else if sg == 0 {
				g = 0xffff
			} else {
				g = clip16(dg * 0xffff / sg)
			}

			if db == 0 {
				b = 0
			} else if sb == 0 {
				b = 0xffff
			} else {
				b = clip16(db * 0xffff / sb)
			}

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d divide) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawDivideFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// hue implements the hue blend mode.
type hue struct{}

// String implemenets fmt.Stringer interface.
func (d hue) String() string {
	return "Hue"
}

// Draw implements image.Drawer interface.
func (d hue) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d hue) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d hue) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHueNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hue) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHueRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hue) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHueNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d hue) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawHueRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawHueNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHueRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHueNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHueRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawHueFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum16(r, g, b, lum16(dr, dg, db))

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d hue) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawHueFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// saturation implements the saturation blend mode.
type saturation struct{}

// String implemenets fmt.Stringer interface.
func (d saturation) String() string {
	return "Saturation"
}

// Draw implements image.Drawer interface.
func (d saturation) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d saturation) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d saturation) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSaturationNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d saturation) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSaturationRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d saturation) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSaturationNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d saturation) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawSaturationRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawSaturationNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSaturationRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSaturationNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSaturationRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawSaturationFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum16(r, g, b, lum16(dr, dg, db))

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d saturation) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawSaturationFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// color implements the color blend mode.
type color struct{}

// String implemenets fmt.Stringer interface.
func (d color) String() string {
	return "Color"
}

// Draw implements image.Drawer interface.
func (d color) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d color) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d color) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d color) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d color) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d color) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawColorRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawColorNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawColorFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r, g, b = setLum16(sr, sg, sb, lum16(dr, dg, db))

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d color) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawColorFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}

// luminosity implements the luminosity blend mode.
type luminosity struct{}

// String implemenets fmt.Stringer interface.
func (d luminosity) String() string {
	return "Luminosity"
}

// Draw implements image.Drawer interface.
func (d luminosity) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d luminosity) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d luminosity) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLuminosityNRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d luminosity) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLuminosityRGBAToNRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d luminosity) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLuminosityNRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

func (d luminosity) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {

	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
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
	drawLuminosityRGBAToRGBA.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)

}

var drawLuminosityNRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLuminosityRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLuminosityNRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLuminosityRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}

			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dyDelta
		sPos += syDelta
	}

}

var drawLuminosityFallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

			var r, g, b uint32

			r, g, b = setLum16(dr, dg, db, lum16(sr, sg, sb))

			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)

			dst.Set(x, y, &out)
		}
	}
}

func (d luminosity) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX-r.Min.X, pY-r.Min.Y
	drawLuminosityFallback.Parallel(
		dst, pX, pY,
		src, sp.X+ofsX, sp.Y+ofsY,
		mask, mp.X+ofsX, mp.Y+ofsY,
		endX, endY, delta, delta,
	)
}
