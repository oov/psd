package blend

import (
	"image"
	"image/draw"
	"runtime"
	"sync"
)

type Drawer interface {
	draw.Drawer
	DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool)
}

type drawer interface {
	drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool)
}

type alphaDrawer interface {
	drawAlphaToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform, protectAlpha bool)
	drawAlphaToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.Alpha, sp image.Point, mask *image.Uniform, protectAlpha bool)
}

type drawFunc func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int)

func (f drawFunc) Parallel(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {
	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > y {
		n--
	}
	if n == 1 {
		f(dest, src, alpha, y, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)
		return
	}
	var wg sync.WaitGroup
	wg.Add(n)
	step := y / n
	for i := 1; i < n; i++ {
		go func(d []byte, s []byte) {
			defer wg.Done()
			f(d, s, alpha, step, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)
		}(dest, src)
		dest = dest[dyDelta*step:]
		src = src[syDelta*step:]
		y -= step
	}
	go func() {
		defer wg.Done()
		f(dest, src, alpha, y, sx0, sx1, sxDelta, syDelta, dx0, dx1, dxDelta, dyDelta)
	}()
	wg.Wait()
}

type drawFallbackFunc func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int)

func (f drawFallbackFunc) Parallel(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var y int
	if dy > 0 {
		y = endY - pY
	} else {
		y = pY - endY
	}
	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > y {
		n--
	}
	if n == 1 {
		f(dst, pX, pY, src, spX, spY, mask, mpX, mpY, endX, endY, dx, dy)
		return
	}
	var wg sync.WaitGroup
	wg.Add(n)
	step := (y / n) * dy
	endY2 := pY + step
	for i := 1; i < n; i++ {
		go func(pY int, spY int, mpY int, endY2 int) {
			defer wg.Done()
			f(dst, pX, pY, src, spX, spY, mask, mpX, mpY, endX, endY2, dx, dy)
		}(pY, spY, mpY, endY2)
		pY += step
		spY += step
		mpY += step
		endY2 += step
	}
	go func() {
		defer wg.Done()
		f(dst, pX, pY, src, spX, spY, mask, mpX, mpY, endX, endY, dx, dy)
	}()
	wg.Wait()
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
	if ad, ok := d.(alphaDrawer); ok {
		switch dst0 := dst.(type) {
		case *image.RGBA:
			if mask == nil {
				switch src0 := src.(type) {
				case *image.Alpha:
					ad.drawAlphaToRGBAUniform(dst0, r, src0, sp, nil, protectAlpha)
					return
				}
			} else {
				switch mask0 := mask.(type) {
				case *image.Uniform:
					switch src0 := src.(type) {
					case *image.Alpha:
						ad.drawAlphaToRGBAUniform(dst0, r, src0, sp, mask0, protectAlpha)
						return
					}
				}
			}
		case *image.NRGBA:
			if mask == nil {
				switch src0 := src.(type) {
				case *image.Alpha:
					ad.drawAlphaToNRGBAUniform(dst0, r, src0, sp, nil, protectAlpha)
					return
				}
			} else {
				switch mask0 := mask.(type) {
				case *image.Uniform:
					switch src0 := src.(type) {
					case *image.Alpha:
						ad.drawAlphaToNRGBAUniform(dst0, r, src0, sp, mask0, protectAlpha)
						return
					}
				}
			}
		}
	}
	d.drawFallback(dst, r, src, sp, mask, mp, protectAlpha)
}
