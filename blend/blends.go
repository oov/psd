// DO NOT EDIT.
// Generate with: go generate

package blend

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oov/psd"
)

var Mode = map[psd.BlendMode]Drawer{
	psd.BlendModePassThrough:  Normal{},
	psd.BlendModeNormal:       Normal{},
	psd.BlendModeDarken:       Darken{},
	psd.BlendModeMultiply:     Multiply{},
	psd.BlendModeColorBurn:    ColorBurn{},
	psd.BlendModeLinearBurn:   LinearBurn{},
	psd.BlendModeDarkerColor:  DarkerColor{},
	psd.BlendModeLighten:      Lighten{},
	psd.BlendModeScreen:       Screen{},
	psd.BlendModeColorDodge:   ColorDodge{},
	psd.BlendModeLinearDodge:  LinearDodge{},
	psd.BlendModeLighterColor: LighterColor{},
	psd.BlendModeOverlay:      Overlay{},
	psd.BlendModeSoftLight:    SoftLight{},
	psd.BlendModeHardLight:    HardLight{},
	psd.BlendModeLinearLight:  LinearLight{},
	psd.BlendModeVividLight:   VividLight{},
	psd.BlendModePinLight:     PinLight{},
	psd.BlendModeHardMix:      HardMix{},
	psd.BlendModeDifference:   Difference{},
	psd.BlendModeExclusion:    Exclusion{},
	psd.BlendModeSubtract:     Subtract{},
	psd.BlendModeDivide:       Divide{},
	psd.BlendModeHue:          Hue{},
	psd.BlendModeSaturation:   Saturation{},
	psd.BlendModeColor:        Color{},
	psd.BlendModeLuminosity:   Luminosity{},
}

// Normal implements the Normal blend mode.
type Normal struct{}

// String implemenets fmt.Stringer interface.
func (d Normal) String() string {
	return "Normal"
}

// BlendMode returns psd.BlendMode.
func (d Normal) BlendMode() psd.BlendMode {
	return psd.BlendModeNormal
}

// Draw implements image.Drawer interface.
func (d Normal) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Normal) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Normal) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawNormalNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawNormalNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Normal) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawNormalRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawNormalRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Normal) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawNormalNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawNormalNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Normal) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawNormalRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawNormalRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawNormalNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawNormalRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Normal) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Darken implements the Darken blend mode.
type Darken struct{}

// String implemenets fmt.Stringer interface.
func (d Darken) String() string {
	return "Darken"
}

// BlendMode returns psd.BlendMode.
func (d Darken) BlendMode() psd.BlendMode {
	return psd.BlendModeDarken
}

// Draw implements image.Drawer interface.
func (d Darken) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Darken) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Darken) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkenNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkenNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Darken) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkenRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkenRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Darken) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkenNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkenNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Darken) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkenRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkenRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawDarkenNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkenRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Darken) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Multiply implements the Multiply blend mode.
type Multiply struct{}

// String implemenets fmt.Stringer interface.
func (d Multiply) String() string {
	return "Multiply"
}

// BlendMode returns psd.BlendMode.
func (d Multiply) BlendMode() psd.BlendMode {
	return psd.BlendModeMultiply
}

// Draw implements image.Drawer interface.
func (d Multiply) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Multiply) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Multiply) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawMultiplyNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawMultiplyNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Multiply) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawMultiplyRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawMultiplyRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Multiply) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawMultiplyNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawMultiplyNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Multiply) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawMultiplyRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawMultiplyRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawMultiplyNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawMultiplyRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Multiply) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// ColorBurn implements the ColorBurn blend mode.
type ColorBurn struct{}

// String implemenets fmt.Stringer interface.
func (d ColorBurn) String() string {
	return "ColorBurn"
}

// BlendMode returns psd.BlendMode.
func (d ColorBurn) BlendMode() psd.BlendMode {
	return psd.BlendModeColorBurn
}

// Draw implements image.Drawer interface.
func (d ColorBurn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d ColorBurn) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d ColorBurn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorBurnNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorBurnNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d ColorBurn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorBurnRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorBurnRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d ColorBurn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorBurnNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorBurnNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d ColorBurn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorBurnRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorBurnRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawColorBurnNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorBurnRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d ColorBurn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// LinearBurn implements the LinearBurn blend mode.
type LinearBurn struct{}

// String implemenets fmt.Stringer interface.
func (d LinearBurn) String() string {
	return "LinearBurn"
}

// BlendMode returns psd.BlendMode.
func (d LinearBurn) BlendMode() psd.BlendMode {
	return psd.BlendModeLinearBurn
}

// Draw implements image.Drawer interface.
func (d LinearBurn) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d LinearBurn) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d LinearBurn) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearBurnNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearBurnNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearBurn) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearBurnRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearBurnRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearBurn) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearBurnNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearBurnNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearBurn) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearBurnRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearBurnRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawLinearBurnNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearBurnRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d LinearBurn) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// DarkerColor implements the DarkerColor blend mode.
type DarkerColor struct{}

// String implemenets fmt.Stringer interface.
func (d DarkerColor) String() string {
	return "DarkerColor"
}

// BlendMode returns psd.BlendMode.
func (d DarkerColor) BlendMode() psd.BlendMode {
	return psd.BlendModeDarkerColor
}

// Draw implements image.Drawer interface.
func (d DarkerColor) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d DarkerColor) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d DarkerColor) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkerColorNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkerColorNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d DarkerColor) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkerColorRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkerColorRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d DarkerColor) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkerColorNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkerColorNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d DarkerColor) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDarkerColorRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDarkerColorRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawDarkerColorNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDarkerColorRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d DarkerColor) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Lighten implements the Lighten blend mode.
type Lighten struct{}

// String implemenets fmt.Stringer interface.
func (d Lighten) String() string {
	return "Lighten"
}

// BlendMode returns psd.BlendMode.
func (d Lighten) BlendMode() psd.BlendMode {
	return psd.BlendModeLighten
}

// Draw implements image.Drawer interface.
func (d Lighten) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Lighten) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Lighten) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLightenNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLightenNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Lighten) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLightenRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLightenRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Lighten) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLightenNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLightenNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Lighten) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLightenRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLightenRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawLightenNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLightenRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Lighten) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Screen implements the Screen blend mode.
type Screen struct{}

// String implemenets fmt.Stringer interface.
func (d Screen) String() string {
	return "Screen"
}

// BlendMode returns psd.BlendMode.
func (d Screen) BlendMode() psd.BlendMode {
	return psd.BlendModeScreen
}

// Draw implements image.Drawer interface.
func (d Screen) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Screen) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Screen) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawScreenNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawScreenNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Screen) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawScreenRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawScreenRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Screen) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawScreenNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawScreenNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Screen) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawScreenRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawScreenRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawScreenNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawScreenRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Screen) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// ColorDodge implements the ColorDodge blend mode.
type ColorDodge struct{}

// String implemenets fmt.Stringer interface.
func (d ColorDodge) String() string {
	return "ColorDodge"
}

// BlendMode returns psd.BlendMode.
func (d ColorDodge) BlendMode() psd.BlendMode {
	return psd.BlendModeColorDodge
}

// Draw implements image.Drawer interface.
func (d ColorDodge) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d ColorDodge) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d ColorDodge) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorDodgeNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorDodgeNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d ColorDodge) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorDodgeRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorDodgeRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d ColorDodge) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorDodgeNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorDodgeNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d ColorDodge) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorDodgeRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorDodgeRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawColorDodgeNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorDodgeRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d ColorDodge) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// LinearDodge implements the LinearDodge blend mode.
type LinearDodge struct{}

// String implemenets fmt.Stringer interface.
func (d LinearDodge) String() string {
	return "LinearDodge"
}

// BlendMode returns psd.BlendMode.
func (d LinearDodge) BlendMode() psd.BlendMode {
	return psd.BlendModeLinearDodge
}

// Draw implements image.Drawer interface.
func (d LinearDodge) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d LinearDodge) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d LinearDodge) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearDodgeNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearDodgeNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearDodge) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearDodgeRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearDodgeRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearDodge) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearDodgeNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearDodgeNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearDodge) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearDodgeRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearDodgeRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawLinearDodgeNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearDodgeRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d LinearDodge) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// LighterColor implements the LighterColor blend mode.
type LighterColor struct{}

// String implemenets fmt.Stringer interface.
func (d LighterColor) String() string {
	return "LighterColor"
}

// BlendMode returns psd.BlendMode.
func (d LighterColor) BlendMode() psd.BlendMode {
	return psd.BlendModeLighterColor
}

// Draw implements image.Drawer interface.
func (d LighterColor) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d LighterColor) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d LighterColor) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLighterColorNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLighterColorNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LighterColor) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLighterColorRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLighterColorRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LighterColor) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLighterColorNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLighterColorNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LighterColor) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLighterColorRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLighterColorRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawLighterColorNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLighterColorRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d LighterColor) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Add implements the Add blend mode.
type Add struct{}

// String implemenets fmt.Stringer interface.
func (d Add) String() string {
	return "Add"
}

// Draw implements image.Drawer interface.
func (d Add) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Add) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Add) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawAddNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawAddNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Add) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawAddRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawAddRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Add) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawAddNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawAddNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Add) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawAddRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawAddRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawAddNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

var drawAddRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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

func (d Add) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

// Overlay implements the Overlay blend mode.
type Overlay struct{}

// String implemenets fmt.Stringer interface.
func (d Overlay) String() string {
	return "Overlay"
}

// BlendMode returns psd.BlendMode.
func (d Overlay) BlendMode() psd.BlendMode {
	return psd.BlendModeOverlay
}

// Draw implements image.Drawer interface.
func (d Overlay) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Overlay) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Overlay) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawOverlayNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawOverlayNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Overlay) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawOverlayRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawOverlayRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Overlay) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawOverlayNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawOverlayNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Overlay) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawOverlayRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawOverlayRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawOverlayNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawOverlayRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Overlay) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// SoftLight implements the SoftLight blend mode.
type SoftLight struct{}

// String implemenets fmt.Stringer interface.
func (d SoftLight) String() string {
	return "SoftLight"
}

// BlendMode returns psd.BlendMode.
func (d SoftLight) BlendMode() psd.BlendMode {
	return psd.BlendModeSoftLight
}

// Draw implements image.Drawer interface.
func (d SoftLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d SoftLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d SoftLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSoftLightNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSoftLightNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d SoftLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSoftLightRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSoftLightRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d SoftLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSoftLightNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSoftLightNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d SoftLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSoftLightRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSoftLightRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawSoftLightNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSoftLightRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d SoftLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// HardLight implements the HardLight blend mode.
type HardLight struct{}

// String implemenets fmt.Stringer interface.
func (d HardLight) String() string {
	return "HardLight"
}

// BlendMode returns psd.BlendMode.
func (d HardLight) BlendMode() psd.BlendMode {
	return psd.BlendModeHardLight
}

// Draw implements image.Drawer interface.
func (d HardLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d HardLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d HardLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardLightNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardLightNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d HardLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardLightRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardLightRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d HardLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardLightNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardLightNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d HardLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardLightRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardLightRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawHardLightNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardLightRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d HardLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// LinearLight implements the LinearLight blend mode.
type LinearLight struct{}

// String implemenets fmt.Stringer interface.
func (d LinearLight) String() string {
	return "LinearLight"
}

// BlendMode returns psd.BlendMode.
func (d LinearLight) BlendMode() psd.BlendMode {
	return psd.BlendModeLinearLight
}

// Draw implements image.Drawer interface.
func (d LinearLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d LinearLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d LinearLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearLightNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearLightNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearLightRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearLightRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearLightNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearLightNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d LinearLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLinearLightRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLinearLightRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawLinearLightNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLinearLightRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d LinearLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// VividLight implements the VividLight blend mode.
type VividLight struct{}

// String implemenets fmt.Stringer interface.
func (d VividLight) String() string {
	return "VividLight"
}

// BlendMode returns psd.BlendMode.
func (d VividLight) BlendMode() psd.BlendMode {
	return psd.BlendModeVividLight
}

// Draw implements image.Drawer interface.
func (d VividLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d VividLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d VividLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawVividLightNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawVividLightNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d VividLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawVividLightRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawVividLightRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d VividLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawVividLightNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawVividLightNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d VividLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawVividLightRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawVividLightRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawVividLightNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawVividLightRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d VividLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// PinLight implements the PinLight blend mode.
type PinLight struct{}

// String implemenets fmt.Stringer interface.
func (d PinLight) String() string {
	return "PinLight"
}

// BlendMode returns psd.BlendMode.
func (d PinLight) BlendMode() psd.BlendMode {
	return psd.BlendModePinLight
}

// Draw implements image.Drawer interface.
func (d PinLight) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d PinLight) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d PinLight) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawPinLightNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawPinLightNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d PinLight) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawPinLightRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawPinLightRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d PinLight) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawPinLightNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawPinLightNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d PinLight) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawPinLightRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawPinLightRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawPinLightNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawPinLightRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d PinLight) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// HardMix implements the HardMix blend mode.
type HardMix struct{}

// String implemenets fmt.Stringer interface.
func (d HardMix) String() string {
	return "HardMix"
}

// BlendMode returns psd.BlendMode.
func (d HardMix) BlendMode() psd.BlendMode {
	return psd.BlendModeHardMix
}

// Draw implements image.Drawer interface.
func (d HardMix) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d HardMix) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d HardMix) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardMixNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardMixNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d HardMix) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardMixRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardMixRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d HardMix) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardMixNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardMixNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d HardMix) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHardMixRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHardMixRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawHardMixNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHardMixRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d HardMix) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Difference implements the Difference blend mode.
type Difference struct{}

// String implemenets fmt.Stringer interface.
func (d Difference) String() string {
	return "Difference"
}

// BlendMode returns psd.BlendMode.
func (d Difference) BlendMode() psd.BlendMode {
	return psd.BlendModeDifference
}

// Draw implements image.Drawer interface.
func (d Difference) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Difference) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Difference) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDifferenceNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDifferenceNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Difference) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDifferenceRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDifferenceRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Difference) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDifferenceNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDifferenceNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Difference) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDifferenceRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDifferenceRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawDifferenceNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDifferenceRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Difference) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Exclusion implements the Exclusion blend mode.
type Exclusion struct{}

// String implemenets fmt.Stringer interface.
func (d Exclusion) String() string {
	return "Exclusion"
}

// BlendMode returns psd.BlendMode.
func (d Exclusion) BlendMode() psd.BlendMode {
	return psd.BlendModeExclusion
}

// Draw implements image.Drawer interface.
func (d Exclusion) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Exclusion) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Exclusion) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawExclusionNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawExclusionNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Exclusion) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawExclusionRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawExclusionRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Exclusion) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawExclusionNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawExclusionNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Exclusion) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawExclusionRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawExclusionRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawExclusionNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawExclusionRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Exclusion) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Subtract implements the Subtract blend mode.
type Subtract struct{}

// String implemenets fmt.Stringer interface.
func (d Subtract) String() string {
	return "Subtract"
}

// BlendMode returns psd.BlendMode.
func (d Subtract) BlendMode() psd.BlendMode {
	return psd.BlendModeSubtract
}

// Draw implements image.Drawer interface.
func (d Subtract) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Subtract) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Subtract) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSubtractNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSubtractNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Subtract) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSubtractRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSubtractRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Subtract) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSubtractNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSubtractNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Subtract) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSubtractRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSubtractRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawSubtractNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSubtractRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Subtract) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Divide implements the Divide blend mode.
type Divide struct{}

// String implemenets fmt.Stringer interface.
func (d Divide) String() string {
	return "Divide"
}

// BlendMode returns psd.BlendMode.
func (d Divide) BlendMode() psd.BlendMode {
	return psd.BlendModeDivide
}

// Draw implements image.Drawer interface.
func (d Divide) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Divide) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Divide) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDivideNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDivideNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Divide) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDivideRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDivideRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Divide) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDivideNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDivideNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Divide) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawDivideRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawDivideRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawDivideNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawDivideRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Divide) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Hue implements the Hue blend mode.
type Hue struct{}

// String implemenets fmt.Stringer interface.
func (d Hue) String() string {
	return "Hue"
}

// BlendMode returns psd.BlendMode.
func (d Hue) BlendMode() psd.BlendMode {
	return psd.BlendModeHue
}

// Draw implements image.Drawer interface.
func (d Hue) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Hue) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Hue) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHueNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHueNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Hue) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHueRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHueRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Hue) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHueNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHueNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Hue) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawHueRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawHueRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawHueNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawHueRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Hue) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Saturation implements the Saturation blend mode.
type Saturation struct{}

// String implemenets fmt.Stringer interface.
func (d Saturation) String() string {
	return "Saturation"
}

// BlendMode returns psd.BlendMode.
func (d Saturation) BlendMode() psd.BlendMode {
	return psd.BlendModeSaturation
}

// Draw implements image.Drawer interface.
func (d Saturation) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Saturation) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Saturation) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSaturationNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSaturationNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Saturation) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSaturationRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSaturationRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Saturation) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSaturationNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSaturationNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Saturation) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawSaturationRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawSaturationRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawSaturationNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawSaturationRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Saturation) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Color implements the Color blend mode.
type Color struct{}

// String implemenets fmt.Stringer interface.
func (d Color) String() string {
	return "Color"
}

// BlendMode returns psd.BlendMode.
func (d Color) BlendMode() psd.BlendMode {
	return psd.BlendModeColor
}

// Draw implements image.Drawer interface.
func (d Color) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Color) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Color) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Color) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Color) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Color) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawColorRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawColorRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawColorNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawColorRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Color) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}

// Luminosity implements the Luminosity blend mode.
type Luminosity struct{}

// String implemenets fmt.Stringer interface.
func (d Luminosity) String() string {
	return "Luminosity"
}

// BlendMode returns psd.BlendMode.
func (d Luminosity) BlendMode() psd.BlendMode {
	return psd.BlendModeLuminosity
}

// Draw implements image.Drawer interface.
func (d Luminosity) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d Luminosity) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	drawMask(d, dst, r, src, sp, mask, mp, protectAlpha)
}

func (d Luminosity) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLuminosityNRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLuminosityNRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Luminosity) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLuminosityRGBAToNRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLuminosityRGBAToNRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Luminosity) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLuminosityNRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLuminosityNRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

func (d Luminosity) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {

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

	if protectAlpha {
		drawLuminosityRGBAToRGBAProtectAlpha(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	} else {
		drawLuminosityRGBAToRGBA(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
	}

}

var drawLuminosityNRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityNRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityNRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToNRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) / da))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) / da))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) / da))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityNRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

var drawLuminosityRGBAToRGBAProtectAlpha = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {

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
			dpix[i+2] = uint8(clip8((b*a1 + db*a3) * 32897 >> 23))
			dpix[i+1] = uint8(clip8((g*a1 + dg*a3) * 32897 >> 23))
			dpix[i+0] = uint8(clip8((r*a1 + dr*a3) * 32897 >> 23))

		}
		dPos += dDelta
		sPos += sDelta
	}

}

func (d Luminosity) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	if protectAlpha {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + db*a3) / 0xffff))
				out.A = uint16(da)
				dst.Set(x, y, &out)
			}
		}
	} else {
		var out color.RGBA64
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

				out.R = uint16(clip16((r*a1 + sr*a2 + dr*a3) / 0xffff))
				out.G = uint16(clip16((g*a1 + sg*a2 + dg*a3) / 0xffff))
				out.B = uint16(clip16((b*a1 + sb*a2 + db*a3) / 0xffff))
				out.A = uint16(a)

				dst.Set(x, y, &out)
			}
		}
	}
}
