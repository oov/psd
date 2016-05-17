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
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d normal) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d normal) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawNormalNRGBAToNRGBAProtectAlpha
	} else {
		f = drawNormalNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d normal) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawNormalRGBAToNRGBAProtectAlpha
	} else {
		f = drawNormalRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d normal) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawNormalNRGBAToRGBAProtectAlpha
	} else {
		f = drawNormalNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d normal) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawNormalRGBAToRGBAProtectAlpha
	} else {
		f = drawNormalRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawNormalNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr

			g = sg

			b = sb

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d normal) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r = sr

				g = sg

				b = sb

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// darken implements the darken blend mode.
type darken struct{}

// String implemenets fmt.Stringer interface.
func (d darken) String() string {
	return "Darken"
}

// Draw implements image.Drawer interface.
func (d darken) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d darken) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d darken) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkenNRGBAToNRGBAProtectAlpha
	} else {
		f = drawDarkenNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d darken) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkenRGBAToNRGBAProtectAlpha
	} else {
		f = drawDarkenRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d darken) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkenNRGBAToRGBAProtectAlpha
	} else {
		f = drawDarkenNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d darken) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkenRGBAToRGBAProtectAlpha
	} else {
		f = drawDarkenRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDarkenNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d darken) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// multiply implements the multiply blend mode.
type multiply struct{}

// String implemenets fmt.Stringer interface.
func (d multiply) String() string {
	return "Multiply"
}

// Draw implements image.Drawer interface.
func (d multiply) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d multiply) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d multiply) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawMultiplyNRGBAToNRGBAProtectAlpha
	} else {
		f = drawMultiplyNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d multiply) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawMultiplyRGBAToNRGBAProtectAlpha
	} else {
		f = drawMultiplyRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d multiply) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawMultiplyNRGBAToRGBAProtectAlpha
	} else {
		f = drawMultiplyNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d multiply) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawMultiplyRGBAToRGBAProtectAlpha
	} else {
		f = drawMultiplyRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawMultiplyNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr * dr * 32897 >> 23

			g = sg * dg * 32897 >> 23

			b = sb * db * 32897 >> 23

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d multiply) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r = sr * dr / 0xffff

				g = sg * dg / 0xffff

				b = sb * db / 0xffff

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// colorBurn implements the colorBurn blend mode.
type colorBurn struct{}

// String implemenets fmt.Stringer interface.
func (d colorBurn) String() string {
	return "ColorBurn"
}

// Draw implements image.Drawer interface.
func (d colorBurn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d colorBurn) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d colorBurn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorBurnNRGBAToNRGBAProtectAlpha
	} else {
		f = drawColorBurnNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d colorBurn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorBurnRGBAToNRGBAProtectAlpha
	} else {
		f = drawColorBurnRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d colorBurn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorBurnNRGBAToRGBAProtectAlpha
	} else {
		f = drawColorBurnNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d colorBurn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorBurnRGBAToRGBAProtectAlpha
	} else {
		f = drawColorBurnRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawColorBurnNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d colorBurn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// linearBurn implements the linearBurn blend mode.
type linearBurn struct{}

// String implemenets fmt.Stringer interface.
func (d linearBurn) String() string {
	return "LinearBurn"
}

// Draw implements image.Drawer interface.
func (d linearBurn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d linearBurn) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d linearBurn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearBurnNRGBAToNRGBAProtectAlpha
	} else {
		f = drawLinearBurnNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearBurn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearBurnRGBAToNRGBAProtectAlpha
	} else {
		f = drawLinearBurnRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearBurn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearBurnNRGBAToRGBAProtectAlpha
	} else {
		f = drawLinearBurnNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearBurn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearBurnRGBAToRGBAProtectAlpha
	} else {
		f = drawLinearBurnRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawLinearBurnNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d linearBurn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// darkerColor implements the darkerColor blend mode.
type darkerColor struct{}

// String implemenets fmt.Stringer interface.
func (d darkerColor) String() string {
	return "DarkerColor"
}

// Draw implements image.Drawer interface.
func (d darkerColor) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d darkerColor) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d darkerColor) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkerColorNRGBAToNRGBAProtectAlpha
	} else {
		f = drawDarkerColorNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d darkerColor) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkerColorRGBAToNRGBAProtectAlpha
	} else {
		f = drawDarkerColorRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d darkerColor) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkerColorNRGBAToRGBAProtectAlpha
	} else {
		f = drawDarkerColorNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d darkerColor) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDarkerColorRGBAToRGBAProtectAlpha
	} else {
		f = drawDarkerColorRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDarkerColorNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) < lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d darkerColor) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				if lum16(sr, sg, sb) < lum16(dr, dg, db) {
					r = sr
					g = sg
					b = sb
				} else {
					r = dr
					g = dg
					b = db
				}

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// lighten implements the lighten blend mode.
type lighten struct{}

// String implemenets fmt.Stringer interface.
func (d lighten) String() string {
	return "Lighten"
}

// Draw implements image.Drawer interface.
func (d lighten) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d lighten) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d lighten) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLightenNRGBAToNRGBAProtectAlpha
	} else {
		f = drawLightenNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d lighten) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLightenRGBAToNRGBAProtectAlpha
	} else {
		f = drawLightenRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d lighten) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLightenNRGBAToRGBAProtectAlpha
	} else {
		f = drawLightenNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d lighten) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLightenRGBAToRGBAProtectAlpha
	} else {
		f = drawLightenRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawLightenNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d lighten) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// screen implements the screen blend mode.
type screen struct{}

// String implemenets fmt.Stringer interface.
func (d screen) String() string {
	return "Screen"
}

// Draw implements image.Drawer interface.
func (d screen) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d screen) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d screen) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawScreenNRGBAToNRGBAProtectAlpha
	} else {
		f = drawScreenNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d screen) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawScreenRGBAToNRGBAProtectAlpha
	} else {
		f = drawScreenRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d screen) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawScreenNRGBAToRGBAProtectAlpha
	} else {
		f = drawScreenNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d screen) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawScreenRGBAToRGBAProtectAlpha
	} else {
		f = drawScreenRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawScreenNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr - (sr * dr * 32897 >> 23)

			g = sg + dg - (sg * dg * 32897 >> 23)

			b = sb + db - (sb * db * 32897 >> 23)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d screen) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r = sr + dr - (sr * dr / 0xffff)

				g = sg + dg - (sg * dg / 0xffff)

				b = sb + db - (sb * db / 0xffff)

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// colorDodge implements the colorDodge blend mode.
type colorDodge struct{}

// String implemenets fmt.Stringer interface.
func (d colorDodge) String() string {
	return "ColorDodge"
}

// Draw implements image.Drawer interface.
func (d colorDodge) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d colorDodge) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d colorDodge) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorDodgeNRGBAToNRGBAProtectAlpha
	} else {
		f = drawColorDodgeNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d colorDodge) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorDodgeRGBAToNRGBAProtectAlpha
	} else {
		f = drawColorDodgeRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d colorDodge) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorDodgeNRGBAToRGBAProtectAlpha
	} else {
		f = drawColorDodgeNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d colorDodge) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorDodgeRGBAToRGBAProtectAlpha
	} else {
		f = drawColorDodgeRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawColorDodgeNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d colorDodge) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// linearDodge implements the linearDodge blend mode.
type linearDodge struct{}

// String implemenets fmt.Stringer interface.
func (d linearDodge) String() string {
	return "LinearDodge"
}

// Draw implements image.Drawer interface.
func (d linearDodge) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d linearDodge) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d linearDodge) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearDodgeNRGBAToNRGBAProtectAlpha
	} else {
		f = drawLinearDodgeNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearDodge) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearDodgeRGBAToNRGBAProtectAlpha
	} else {
		f = drawLinearDodgeRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearDodge) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearDodgeNRGBAToRGBAProtectAlpha
	} else {
		f = drawLinearDodgeNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearDodge) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearDodgeRGBAToRGBAProtectAlpha
	} else {
		f = drawLinearDodgeRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawLinearDodgeNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = clip8(sr + dr)

			g = clip8(sg + dg)

			b = clip8(sb + db)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d linearDodge) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r = clip16(sr + dr)

				g = clip16(sg + dg)

				b = clip16(sb + db)

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// lighterColor implements the lighterColor blend mode.
type lighterColor struct{}

// String implemenets fmt.Stringer interface.
func (d lighterColor) String() string {
	return "LighterColor"
}

// Draw implements image.Drawer interface.
func (d lighterColor) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d lighterColor) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d lighterColor) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLighterColorNRGBAToNRGBAProtectAlpha
	} else {
		f = drawLighterColorNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d lighterColor) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLighterColorRGBAToNRGBAProtectAlpha
	} else {
		f = drawLighterColorRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d lighterColor) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLighterColorNRGBAToRGBAProtectAlpha
	} else {
		f = drawLighterColorNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d lighterColor) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLighterColorRGBAToRGBAProtectAlpha
	} else {
		f = drawLighterColorRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawLighterColorNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			if lum8(sr, sg, sb) > lum8(dr, dg, db) {
				r = sr
				g = sg
				b = sb
			} else {
				r = dr
				g = dg
				b = db
			}

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d lighterColor) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				if lum16(sr, sg, sb) > lum16(dr, dg, db) {
					r = sr
					g = sg
					b = sb
				} else {
					r = dr
					g = dg
					b = db
				}

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// add implements the add blend mode.
type add struct{}

// String implemenets fmt.Stringer interface.
func (d add) String() string {
	return "Add"
}

// Draw implements image.Drawer interface.
func (d add) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d add) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d add) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawAddNRGBAToNRGBAProtectAlpha
	} else {
		f = drawAddNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d add) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawAddRGBAToNRGBAProtectAlpha
	} else {
		f = drawAddRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d add) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawAddNRGBAToRGBAProtectAlpha
	} else {
		f = drawAddNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d add) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawAddRGBAToRGBAProtectAlpha
	} else {
		f = drawAddRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawAddNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawAddRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = sr + dr

			g = sg + dg

			b = sb + db

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d add) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r = sr + dr

				g = sg + dg

				b = sb + db

				out.R = uint16(clip16(dr*a3/0xffff + uint32(uint64(r)*uint64(a1)/0xffff)))
				out.G = uint16(clip16(dg*a3/0xffff + uint32(uint64(g)*uint64(a1)/0xffff)))
				out.B = uint16(clip16(db*a3/0xffff + uint32(uint64(b)*uint64(a1)/0xffff)))
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var r, g, b uint32

				r = sr + dr

				g = sg + dg

				b = sb + db

				out.R = uint16(clip16((sr*a2+dr*a3)/0xffff + uint32(uint64(r)*uint64(a1)/0xffff)))
				out.G = uint16(clip16((sg*a2+dg*a3)/0xffff + uint32(uint64(g)*uint64(a1)/0xffff)))
				out.B = uint16(clip16((sb*a2+db*a3)/0xffff + uint32(uint64(b)*uint64(a1)/0xffff)))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// overlay implements the overlay blend mode.
type overlay struct{}

// String implemenets fmt.Stringer interface.
func (d overlay) String() string {
	return "Overlay"
}

// Draw implements image.Drawer interface.
func (d overlay) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d overlay) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d overlay) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawOverlayNRGBAToNRGBAProtectAlpha
	} else {
		f = drawOverlayNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d overlay) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawOverlayRGBAToNRGBAProtectAlpha
	} else {
		f = drawOverlayRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d overlay) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawOverlayNRGBAToRGBAProtectAlpha
	} else {
		f = drawOverlayNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d overlay) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawOverlayRGBAToRGBAProtectAlpha
	} else {
		f = drawOverlayRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawOverlayNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d overlay) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// softLight implements the softLight blend mode.
type softLight struct{}

// String implemenets fmt.Stringer interface.
func (d softLight) String() string {
	return "SoftLight"
}

// Draw implements image.Drawer interface.
func (d softLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d softLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d softLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSoftLightNRGBAToNRGBAProtectAlpha
	} else {
		f = drawSoftLightNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d softLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSoftLightRGBAToNRGBAProtectAlpha
	} else {
		f = drawSoftLightRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d softLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSoftLightNRGBAToRGBAProtectAlpha
	} else {
		f = drawSoftLightNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d softLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSoftLightRGBAToRGBAProtectAlpha
	} else {
		f = drawSoftLightRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSoftLightNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d softLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// hardLight implements the hardLight blend mode.
type hardLight struct{}

// String implemenets fmt.Stringer interface.
func (d hardLight) String() string {
	return "HardLight"
}

// Draw implements image.Drawer interface.
func (d hardLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d hardLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d hardLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardLightNRGBAToNRGBAProtectAlpha
	} else {
		f = drawHardLightNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hardLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardLightRGBAToNRGBAProtectAlpha
	} else {
		f = drawHardLightRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hardLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardLightNRGBAToRGBAProtectAlpha
	} else {
		f = drawHardLightNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hardLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardLightRGBAToRGBAProtectAlpha
	} else {
		f = drawHardLightRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawHardLightNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d hardLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// linearLight implements the linearLight blend mode.
type linearLight struct{}

// String implemenets fmt.Stringer interface.
func (d linearLight) String() string {
	return "LinearLight"
}

// Draw implements image.Drawer interface.
func (d linearLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d linearLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d linearLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearLightNRGBAToNRGBAProtectAlpha
	} else {
		f = drawLinearLightNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearLightRGBAToNRGBAProtectAlpha
	} else {
		f = drawLinearLightRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearLightNRGBAToRGBAProtectAlpha
	} else {
		f = drawLinearLightNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d linearLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLinearLightRGBAToRGBAProtectAlpha
	} else {
		f = drawLinearLightRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawLinearLightNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d linearLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// vividLight implements the vividLight blend mode.
type vividLight struct{}

// String implemenets fmt.Stringer interface.
func (d vividLight) String() string {
	return "VividLight"
}

// Draw implements image.Drawer interface.
func (d vividLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d vividLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d vividLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawVividLightNRGBAToNRGBAProtectAlpha
	} else {
		f = drawVividLightNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d vividLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawVividLightRGBAToNRGBAProtectAlpha
	} else {
		f = drawVividLightRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d vividLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawVividLightNRGBAToRGBAProtectAlpha
	} else {
		f = drawVividLightNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d vividLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawVividLightRGBAToRGBAProtectAlpha
	} else {
		f = drawVividLightRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawVividLightNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d vividLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// pinLight implements the pinLight blend mode.
type pinLight struct{}

// String implemenets fmt.Stringer interface.
func (d pinLight) String() string {
	return "PinLight"
}

// Draw implements image.Drawer interface.
func (d pinLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d pinLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d pinLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawPinLightNRGBAToNRGBAProtectAlpha
	} else {
		f = drawPinLightNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d pinLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawPinLightRGBAToNRGBAProtectAlpha
	} else {
		f = drawPinLightRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d pinLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawPinLightNRGBAToRGBAProtectAlpha
	} else {
		f = drawPinLightNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d pinLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawPinLightRGBAToRGBAProtectAlpha
	} else {
		f = drawPinLightRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawPinLightNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d pinLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// hardMix implements the hardMix blend mode.
type hardMix struct{}

// String implemenets fmt.Stringer interface.
func (d hardMix) String() string {
	return "HardMix"
}

// Draw implements image.Drawer interface.
func (d hardMix) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d hardMix) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d hardMix) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardMixNRGBAToNRGBAProtectAlpha
	} else {
		f = drawHardMixNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hardMix) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardMixRGBAToNRGBAProtectAlpha
	} else {
		f = drawHardMixRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hardMix) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardMixNRGBAToRGBAProtectAlpha
	} else {
		f = drawHardMixNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hardMix) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHardMixRGBAToRGBAProtectAlpha
	} else {
		f = drawHardMixRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawHardMixNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d hardMix) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// difference implements the difference blend mode.
type difference struct{}

// String implemenets fmt.Stringer interface.
func (d difference) String() string {
	return "Difference"
}

// Draw implements image.Drawer interface.
func (d difference) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d difference) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d difference) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDifferenceNRGBAToNRGBAProtectAlpha
	} else {
		f = drawDifferenceNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d difference) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDifferenceRGBAToNRGBAProtectAlpha
	} else {
		f = drawDifferenceRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d difference) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDifferenceNRGBAToRGBAProtectAlpha
	} else {
		f = drawDifferenceNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d difference) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDifferenceRGBAToRGBAProtectAlpha
	} else {
		f = drawDifferenceRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDifferenceNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d difference) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// exclusion implements the exclusion blend mode.
type exclusion struct{}

// String implemenets fmt.Stringer interface.
func (d exclusion) String() string {
	return "Exclusion"
}

// Draw implements image.Drawer interface.
func (d exclusion) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d exclusion) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d exclusion) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawExclusionNRGBAToNRGBAProtectAlpha
	} else {
		f = drawExclusionNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d exclusion) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawExclusionRGBAToNRGBAProtectAlpha
	} else {
		f = drawExclusionRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d exclusion) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawExclusionNRGBAToRGBAProtectAlpha
	} else {
		f = drawExclusionNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d exclusion) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawExclusionRGBAToRGBAProtectAlpha
	} else {
		f = drawExclusionRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawExclusionNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r = dr + sr - (dr * sr * 32897 >> 22)

			g = dg + sg - (dg * sg * 32897 >> 22)

			b = db + sb - (db * sb * 32897 >> 22)

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d exclusion) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r = dr + sr - (dr * sr / 0x8000)

				g = dg + sg - (dg * sg / 0x8000)

				b = db + sb - (db * sb / 0x8000)

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// subtract implements the subtract blend mode.
type subtract struct{}

// String implemenets fmt.Stringer interface.
func (d subtract) String() string {
	return "Subtract"
}

// Draw implements image.Drawer interface.
func (d subtract) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d subtract) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d subtract) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSubtractNRGBAToNRGBAProtectAlpha
	} else {
		f = drawSubtractNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d subtract) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSubtractRGBAToNRGBAProtectAlpha
	} else {
		f = drawSubtractRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d subtract) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSubtractNRGBAToRGBAProtectAlpha
	} else {
		f = drawSubtractNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d subtract) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSubtractRGBAToRGBAProtectAlpha
	} else {
		f = drawSubtractRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSubtractNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d subtract) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// divide implements the divide blend mode.
type divide struct{}

// String implemenets fmt.Stringer interface.
func (d divide) String() string {
	return "Divide"
}

// Draw implements image.Drawer interface.
func (d divide) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d divide) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d divide) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDivideNRGBAToNRGBAProtectAlpha
	} else {
		f = drawDivideNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d divide) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDivideRGBAToNRGBAProtectAlpha
	} else {
		f = drawDivideRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d divide) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDivideNRGBAToRGBAProtectAlpha
	} else {
		f = drawDivideNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d divide) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawDivideRGBAToRGBAProtectAlpha
	} else {
		f = drawDivideRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawDivideNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
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

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

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

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d divide) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

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

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// hue implements the hue blend mode.
type hue struct{}

// String implemenets fmt.Stringer interface.
func (d hue) String() string {
	return "Hue"
}

// Draw implements image.Drawer interface.
func (d hue) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d hue) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d hue) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHueNRGBAToNRGBAProtectAlpha
	} else {
		f = drawHueNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hue) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHueRGBAToNRGBAProtectAlpha
	} else {
		f = drawHueRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hue) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHueNRGBAToRGBAProtectAlpha
	} else {
		f = drawHueNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d hue) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawHueRGBAToRGBAProtectAlpha
	} else {
		f = drawHueRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawHueNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d hue) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r, g, b = setSat(sr, sg, sb, sat(dr, dg, db))
				r, g, b = setLum16(r, g, b, lum16(dr, dg, db))

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// saturation implements the saturation blend mode.
type saturation struct{}

// String implemenets fmt.Stringer interface.
func (d saturation) String() string {
	return "Saturation"
}

// Draw implements image.Drawer interface.
func (d saturation) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d saturation) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d saturation) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSaturationNRGBAToNRGBAProtectAlpha
	} else {
		f = drawSaturationNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d saturation) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSaturationRGBAToNRGBAProtectAlpha
	} else {
		f = drawSaturationRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d saturation) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSaturationNRGBAToRGBAProtectAlpha
	} else {
		f = drawSaturationNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d saturation) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawSaturationRGBAToRGBAProtectAlpha
	} else {
		f = drawSaturationRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawSaturationNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
			r, g, b = setLum8(r, g, b, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d saturation) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r, g, b = setSat(dr, dg, db, sat(sr, sg, sb))
				r, g, b = setLum16(r, g, b, lum16(dr, dg, db))

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// color implements the color blend mode.
type color struct{}

// String implemenets fmt.Stringer interface.
func (d color) String() string {
	return "Color"
}

// Draw implements image.Drawer interface.
func (d color) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d color) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d color) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorNRGBAToNRGBAProtectAlpha
	} else {
		f = drawColorNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d color) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorRGBAToNRGBAProtectAlpha
	} else {
		f = drawColorRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d color) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorNRGBAToRGBAProtectAlpha
	} else {
		f = drawColorNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d color) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawColorRGBAToRGBAProtectAlpha
	} else {
		f = drawColorRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawColorNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(sr, sg, sb, lum8(dr, dg, db))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d color) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r, g, b = setLum16(sr, sg, sb, lum16(dr, dg, db))

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}

// luminosity implements the luminosity blend mode.
type luminosity struct{}

// String implemenets fmt.Stringer interface.
func (d luminosity) String() string {
	return "Luminosity"
}

// Draw implements image.Drawer interface.
func (d luminosity) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d luminosity) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d luminosity) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLuminosityNRGBAToNRGBAProtectAlpha
	} else {
		f = drawLuminosityNRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d luminosity) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLuminosityRGBAToNRGBAProtectAlpha
	} else {
		f = drawLuminosityRGBAToNRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d luminosity) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLuminosityNRGBAToRGBAProtectAlpha
	} else {
		f = drawLuminosityNRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

func (d luminosity) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	var f drawfunc
	if protectAlpha {
		f = drawLuminosityRGBAToRGBAProtectAlpha
	} else {
		f = drawLuminosityRGBAToRGBA
	}
	f.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)

}

var drawLuminosityNRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToNRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityNRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToRGBA drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			if sa == 0 {
				continue
			}

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}

			var r, g, b uint32

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityNRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToNRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) / da)
			dpix[i+1] = uint8((g*a1 + dg*a3) / da)
			dpix[i+0] = uint8((r*a1 + dr*a3) / da)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityNRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToRGBAProtectAlpha drawfunc = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

	alpha *= 32897
	dPos, sPos := 0, 0
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i := xMin; i != xMax; i += xDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			if da == 0 || sa == 0 {
				continue
			}

			a1 := sa * alpha >> 23
			a3 := 255 - a1

			sr = sr * 0xff / sa
			sg = sg * 0xff / sa
			sb = sb * 0xff / sa

			dr = dr * 0xff / da
			dg = dg * 0xff / da
			db = db * 0xff / da

			var tmp, r, g, b uint32
			_ = tmp

			r, g, b = setLum8(dr, dg, db, lum8(sr, sg, sb))

			dpix[i+3] = uint8(da)
			dpix[i+2] = uint8((b*a1 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + dr*a3) * 32897 >> 23)

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d luminosity) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
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
				if da == 0 {
					continue
				}

				a1 := sa * ma / 0xffff
				a3 := 0xffff - a1
				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
					dr = dr * 0xffff / da
					dg = dg * 0xffff / da
					db = db * 0xffff / da
				}

				var tmp, r, g, b uint32
				_ = tmp

				r, g, b = setLum16(dr, dg, db, lum16(sr, sg, sb))

				out.R = uint16((r*a1 + dr*a3) / 0xffff)
				out.G = uint16((g*a1 + dg*a3) / 0xffff)
				out.B = uint16((b*a1 + db*a3) / 0xffff)
				out.A = uint16(da)

				dst.Set(x, y, &out)
			}
		}
	} else {
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

				tmp := sa * ma / 0xffff
				a1 := tmp * da / 0xffff
				a2 := tmp * (0xffff - da) / 0xffff
				a3 := (0xffff - tmp) * da / 0xffff
				a := a1 + a2 + a3
				if a == 0 {
					continue
				}

				if sa > 0 {
					sr = sr * 0xffff / sa
					sg = sg * 0xffff / sa
					sb = sb * 0xffff / sa
				}
				if da > 0 {
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
}
