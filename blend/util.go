//go:generate go run gen.go

package blend

import (
	"image"
	"image/draw"
	"math"
)

type Drawer interface {
	draw.Drawer
	DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool)
}

type drawer interface {
	Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)
	DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool)
	drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool)
}

func drawMask(d drawer, dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	clip(dst, &r, src, &sp, mask, &mp)
	if r.Empty() {
		return
	}

	switch dst0 := dst.(type) {
	case *image.RGBA:
		if mask == nil {
			switch src0 := src.(type) {
			case *image.RGBA:
				d.drawRGBAToRGBAUniform(dst0, r, src0, sp, nil, protectAlpha)
				return
			case *image.NRGBA:
				d.drawNRGBAToRGBAUniform(dst0, r, src0, sp, nil, protectAlpha)
				return
			}
		} else {
			switch mask0 := mask.(type) {
			case *image.Uniform:
				switch src0 := src.(type) {
				case *image.RGBA:
					d.drawRGBAToRGBAUniform(dst0, r, src0, sp, mask0, protectAlpha)
					return
				case *image.NRGBA:
					d.drawNRGBAToRGBAUniform(dst0, r, src0, sp, mask0, protectAlpha)
					return
				}
			}
		}
	case *image.NRGBA:
		if mask == nil {
			switch src0 := src.(type) {
			case *image.RGBA:
				d.drawRGBAToNRGBAUniform(dst0, r, src0, sp, nil, protectAlpha)
				return
			case *image.NRGBA:
				d.drawNRGBAToNRGBAUniform(dst0, r, src0, sp, nil, protectAlpha)
				return
			}
		} else {
			switch mask0 := mask.(type) {
			case *image.Uniform:
				switch src0 := src.(type) {
				case *image.RGBA:
					d.drawRGBAToNRGBAUniform(dst0, r, src0, sp, mask0, protectAlpha)
					return
				case *image.NRGBA:
					d.drawNRGBAToNRGBAUniform(dst0, r, src0, sp, mask0, protectAlpha)
					return
				}
			}
		}
	}
	d.drawFallback(dst, r, src, sp, mask, mp, protectAlpha)
}

func sqrt(a uint32, b float64) uint32 {
	return uint32(math.Sqrt(float64(a)/b) * b)
}

func max(a uint32, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func maxs(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func maxf(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a uint32, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func mins(a int32, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func minf(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func sat(r uint32, g uint32, b uint32) uint32 {
	return max(r, max(g, b)) - min(r, min(g, b))
}

func setSat(r uint32, g uint32, b uint32, sat uint32) (uint32, uint32, uint32) {
	if r <= g {
		if g <= b {
			return setSatMinMidMax(r, g, b, sat)
		} else if r <= b {
			r, b, g = setSatMinMidMax(r, b, g, sat)
			return r, g, b
		}
		b, r, g = setSatMinMidMax(b, r, g, sat)
		return r, g, b
	} else if r <= b {
		g, r, b = setSatMinMidMax(g, r, b, sat)
		return r, g, b
	} else if g <= b {
		g, b, r = setSatMinMidMax(g, b, r, sat)
		return r, g, b
	}
	b, g, r = setSatMinMidMax(b, g, r, sat)
	return r, g, b
}

func setSatMinMidMax(min uint32, mid uint32, max uint32, sat uint32) (uint32, uint32, uint32) {
	if max > min {
		return 0, (mid - min) * sat / (max - min), sat
	}
	return 0, 0, 0
}

func lum8(r uint32, g uint32, b uint32) int32 {
	return int32((r*77 + g*151 + b*28) >> 8)
}

func lum8s(r int32, g int32, b int32) int32 {
	return (r*77 + g*151 + b*28) >> 8
}

func lum16(r uint32, g uint32, b uint32) float64 {
	return (float64(r)*19661 + float64(g)*38666 + float64(b)*7209) / 0x10000
}

func lum16f(r float64, g float64, b float64) float64 {
	return (r*19661 + g*38666 + b*7209) / 0x10000
}

func setLum8(r uint32, g uint32, b uint32, lm int32) (uint32, uint32, uint32) {
	lm -= lum8(r, g, b)
	return clipColor8(int32(r)+lm, int32(g)+lm, int32(b)+lm)
}

func setLum16(r uint32, g uint32, b uint32, lm float64) (uint32, uint32, uint32) {
	lm -= lum16(r, g, b)
	return clipColor16(float64(r)+lm, float64(g)+lm, float64(b)+lm)
}

func clipColor8(r int32, g int32, b int32) (uint32, uint32, uint32) {
	lm := lum8s(r, g, b)
	min := mins(r, mins(g, b))
	max := maxs(r, maxs(g, b))
	if min < 0 {
		tmp := lm - min
		r = lm + (r-lm)*lm/tmp
		g = lm + (g-lm)*lm/tmp
		b = lm + (b-lm)*lm/tmp
	}
	if max > 0xff {
		tmp, tmp2 := 0xff-lm, max-lm
		r = lm + (r-lm)*tmp/tmp2
		g = lm + (g-lm)*tmp/tmp2
		b = lm + (b-lm)*tmp/tmp2
	}
	return clamp8s(r), clamp8s(g), clamp8s(b)
}

func clipColor16(r float64, g float64, b float64) (uint32, uint32, uint32) {
	lm := lum16f(r, g, b)
	mn := minf(r, minf(g, b))
	mx := maxf(r, maxf(g, b))
	if mn < 0 {
		tmp := lm - mn
		r = lm + (r-lm)*lm/tmp
		g = lm + (g-lm)*lm/tmp
		b = lm + (b-lm)*lm/tmp
	}
	if mx > 0xffff {
		tmp, tmp2 := 0xffff-lm, mx-lm
		r = lm + (r-lm)*tmp/tmp2
		g = lm + (g-lm)*tmp/tmp2
		b = lm + (b-lm)*tmp/tmp2
	}
	return clamp16f(r), clamp16f(g), clamp16f(b)
}

func clip0(a uint32) uint32 {
	if a&0x80000000 == 0 {
		return a
	}
	return 0x0000
}

func clip8(a uint32) uint32 {
	if a < 0xff {
		return a
	}
	return 0xff
}

func clip16(a uint32) uint32 {
	if a < 0xffff {
		return a
	}
	return 0xffff
}

func clamp8s(a int32) uint32 {
	if a < 0 {
		return 0
	}
	if a > 0xff {
		return 0xff
	}
	return uint32(a)
}

func clamp16f(a float64) uint32 {
	if a < 0 {
		return 0
	}
	if a > 0xffff {
		return 0xffff
	}
	return uint32(a)
}

func processBackward(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) bool {
	return image.Image(dst) == src &&
		r.Overlaps(r.Add(sp.Sub(r.Min))) &&
		(sp.Y < r.Min.Y || (sp.Y == r.Min.Y && sp.X < r.Min.X))
}

// clip clips r against each image's bounds (after translating into the
// destination image's coordinate space) and shifts the points sp and mp by
// the same amount as the change in r.Min.
func clip(dst draw.Image, r *image.Rectangle, src image.Image, sp *image.Point, mask image.Image, mp *image.Point) {
	orig := r.Min
	*r = r.Intersect(dst.Bounds())
	*r = r.Intersect(src.Bounds().Add(orig.Sub(*sp)))
	if mask != nil {
		*r = r.Intersect(mask.Bounds().Add(orig.Sub(*mp)))
	}
	dx := r.Min.X - orig.X
	dy := r.Min.Y - orig.Y
	if dx == 0 && dy == 0 {
		return
	}
	sp.X += dx
	sp.Y += dy
	if mp != nil {
		mp.X += dx
		mp.Y += dy
	}
}
