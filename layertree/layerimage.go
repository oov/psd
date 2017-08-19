package layertree

import (
	"context"
	"image"
	"math"
	"runtime"

	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"

	"github.com/oov/psd/blend"
)

type layerImage struct {
	Canvas tiledImage
	Mask   tiledMask
}

type tiledImage map[image.Point]*image.RGBA

func (t *tiledImage) Alloc(tileSize int, rect image.Rectangle) {
	if *t == nil {
		*t = make(tiledImage)
	}

	rx0, ry0, rx1, ry1 := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	for ty := (ry0 / tileSize) * tileSize; ty < ry1; ty += tileSize {
		for tx := (rx0 / tileSize) * tileSize; tx < rx1; tx += tileSize {
			p := image.Pt(tx, ty)
			if _, ok := (*t)[p]; !ok {
				(*t)[p] = image.NewRGBA(image.Rect(tx, ty, tx+tileSize, ty+tileSize))
			}
		}
	}
}

func (t *tiledImage) Get(tileSize int, pt image.Point) (*image.RGBA, bool) {
	r, ok := (*t)[pt]
	if !ok {
		r = image.NewRGBA(image.Rect(pt.X, pt.Y, pt.X+tileSize, pt.Y+tileSize))
		(*t)[pt] = r
	}
	return r, ok
}

func (t *tiledImage) Render(ctx context.Context, tileSize int, img *image.RGBA) error {
	if *t == nil {
		*t = make(tiledImage)
	}

	rect := img.Rect
	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := (ylen / n) * tileSize
	for i := 1; i < n; i++ {
		go t.renderInner(pc, img, tileSize, x0, x1, y0, y0+step)
		y0 += step
	}
	go t.renderInner(pc, img, tileSize, x0, x1, y0, y1)
	return pc.Wait(ctx)
}

func (t *tiledImage) renderInner(pc *parallelContext, img *image.RGBA, tileSize, x0, x1, y0, y1 int) {
	defer pc.Done()
	for ty := y0; ty < y1; ty += tileSize {
		for tx := x0; tx < x1; tx += tileSize {
			if b, ok := (*t)[image.Pt(tx, ty)]; ok {
				blend.Copy.Draw(img, b.Rect, b, b.Rect.Min)
			} else {
				blend.Clear.Draw(img, image.Rect(tx, ty, tx+tileSize, ty+tileSize), image.Transparent, image.Point{})
			}
		}
	}
}

func createImage(rect image.Rectangle, r []byte, g []byte, b []byte, a []byte) *image.NRGBA {
	img := image.NewNRGBA(rect)
	w, h := rect.Dx(), rect.Dy()
	pix := img.Pix
	var y, sx0, sx1, dx int
	if a != nil {
		for y = 0; y < h; y++ {
			sx0 = y * w
			sx1 = sx0 + w
			dx = sx0 << 2
			for sx0 < sx1 {
				if a[sx0] > 0 {
					pix[dx+3] = a[sx0]
					pix[dx+2] = b[sx0]
					pix[dx+1] = g[sx0]
					pix[dx+0] = r[sx0]
				}
				dx += 4
				sx0++
			}
		}
	} else {
		for y = 0; y < h; y++ {
			sx0 = y * w
			sx1 = sx0 + w
			dx = sx0 << 2
			for sx0 < sx1 {
				pix[dx+3] = 0xff
				pix[dx+2] = b[sx0]
				pix[dx+1] = g[sx0]
				pix[dx+0] = r[sx0]
				dx += 4
				sx0++
			}
		}
	}
	return img
}

func scaleRect(rect image.Rectangle, scale float64) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			int(math.Floor(float64(rect.Min.X) * scale)),
			int(math.Floor(float64(rect.Min.Y) * scale)),
		},
		Max: image.Point{
			int(math.Ceil(float64(rect.Max.X) * scale)),
			int(math.Ceil(float64(rect.Max.Y) * scale)),
		},
	}
}

func (t *tiledImage) Rescale(ctx context.Context, tileSize int, rect image.Rectangle, scale float64) (*tiledImage, image.Rectangle, error) {
	tmp := image.NewRGBA(rect)
	if err := t.Render(ctx, tileSize, tmp); err != nil {
		return nil, image.Rectangle{}, err
	}
	scaledRect := scaleRect(rect, scale)
	tmp2 := image.NewNRGBA(scaledRect)
	draw.BiLinear.Transform(tmp2, f64.Aff3{scale, 0, 0, 0, scale, 0}, tmp, rect, draw.Src, nil)
	var t2 tiledImage
	if err := t2.Store(ctx, tileSize, scaledRect, tmp2.Pix[0:], tmp2.Pix[1:], tmp2.Pix[2:], tmp2.Pix[3:], 4); err != nil {
		return nil, image.Rectangle{}, err
	}
	return &t2, scaledRect, nil
}

func (t *tiledImage) StoreScaled(ctx context.Context, tileSize int, rect image.Rectangle, r, g, b, a []byte, scale float64) (image.Rectangle, error) {
	if *t == nil {
		*t = make(tiledImage)
	}

	tmp := createImage(rect, r, g, b, a)
	scaledRect := scaleRect(rect, scale)
	tmp2 := image.NewNRGBA(scaledRect)
	draw.BiLinear.Transform(tmp2, f64.Aff3{scale, 0, 0, 0, scale, 0}, tmp, rect, draw.Src, nil)
	if err := t.Store(ctx, tileSize, scaledRect, tmp2.Pix[0:], tmp2.Pix[1:], tmp2.Pix[2:], tmp2.Pix[3:], 4); err != nil {
		return image.Rectangle{}, err
	}
	return scaledRect, nil
}

func (t *tiledImage) Store(ctx context.Context, tileSize int, rect image.Rectangle, r, g, b, a []byte, deltaX int) error {
	if *t == nil {
		*t = make(tiledImage)
	}

	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := (ylen / n) * tileSize
	for i := 1; i < n; i++ {
		go t.storeInner(pc, rect, tileSize, x0, x1, y0, y0+step, r, g, b, a, deltaX)
		y0 += step
	}
	go t.storeInner(pc, rect, tileSize, x0, x1, y0, y1, r, g, b, a, deltaX)
	return pc.Wait(ctx)
}

func (t *tiledImage) storeInner(pc *parallelContext, rect image.Rectangle, tileSize, x0, x1, y0, y1 int, r, g, b, a []byte, deltaX int) {
	defer pc.Done()

	rx0, ry0, rx1, ry1 := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	sw := rect.Dx() * deltaX
	tw := tileSize << 2

	buf := make([]byte, tw*tileSize)
	for ty := y0; ty < y1; ty += tileSize {
		if pc.Aborted() {
			return
		}
		for tx := x0; tx < x1; tx += tileSize {
			dxMin, dxMax := 0, tileSize
			sxMin := tx - rx0
			if sxMin < 0 {
				dxMin -= sxMin
				sxMin = 0
			}
			if rx0+sxMin+(dxMax-dxMin) > rx1 {
				dxMax -= rx0 + sxMin + (dxMax - dxMin) - rx1
			}

			dyMin, dyMax := 0, tileSize
			syMin := ty - ry0
			if syMin < 0 {
				dyMin -= syMin
				syMin = 0
			}
			if ry0+syMin+(dyMax-dyMin) > ry1 {
				dyMax -= ry0 + syMin + (dyMax - dyMin) - ry1
			}
			used := false
			sxMin = sxMin * deltaX
			dyMax = dyMax * tw
			dxMin, dxMax = dxMin<<2, dxMax<<2
			if a != nil {
				var alpha uint32
				for dy, sy := dyMin*tw, syMin*sw; dy < dyMax; dy += tw {
					for dx, sx, dEnd := dy+dxMin, sy+sxMin, dy+dxMax; dx < dEnd; dx += 4 {
						alpha = uint32(a[sx]) * 32897
						if alpha == 255*32897 {
							buf[dx+3] = 255
							buf[dx+2] = b[sx]
							buf[dx+1] = g[sx]
							buf[dx+0] = r[sx]
							used = true
						} else if alpha > 0 {
							buf[dx+3] = a[sx]
							buf[dx+2] = uint8((uint32(b[sx]) * alpha) >> 23)
							buf[dx+1] = uint8((uint32(g[sx]) * alpha) >> 23)
							buf[dx+0] = uint8((uint32(r[sx]) * alpha) >> 23)
							used = true
						}
						sx += deltaX
					}
					sy += sw
				}
			} else {
				for dy, sy := dyMin*tw, syMin*sw; dy < dyMax; dy += tw {
					for dx, sx, dEnd := dy+dxMin, sy+sxMin, dy+dxMax; dx < dEnd; dx += 4 {
						buf[dx+3] = 0xff
						buf[dx+2] = b[sx]
						buf[dx+1] = g[sx]
						buf[dx+0] = r[sx]
						used = true
						sx += deltaX
					}
					sy += sw
				}
			}
			if used {
				pc.M.Lock()
				(*t)[image.Pt(tx, ty)] = &image.RGBA{
					Pix:    buf,
					Stride: tileSize * 4,
					Rect:   image.Rect(tx, ty, tx+tileSize, ty+tileSize),
				}
				pc.M.Unlock()
				buf = make([]byte, tw*tileSize)
			}
		}
	}
}

type tiledMask map[image.Point]*image.Alpha

func (t *tiledMask) Alloc(tileSize int, rect image.Rectangle) {
	if *t == nil {
		*t = make(tiledMask)
	}
	m := *t

	rx0, ry0, rx1, ry1 := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	for ty := (ry0 / tileSize) * tileSize; ty < ry1; ty += tileSize {
		for tx := (rx0 / tileSize) * tileSize; tx < rx1; tx += tileSize {
			p := image.Pt(tx, ty)
			if _, ok := m[p]; !ok {
				m[p] = image.NewAlpha(image.Rect(tx, ty, tx+tileSize, ty+tileSize))
			}
		}
	}
}

func (t *tiledMask) Render(ctx context.Context, tileSize int, img *image.Alpha) error {
	if *t == nil {
		*t = make(tiledMask)
	}

	rect := img.Rect
	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := (ylen / n) * tileSize
	for i := 1; i < n; i++ {
		go t.renderInner(pc, img, tileSize, x0, x1, y0, y0+step)
		y0 += step
	}
	go t.renderInner(pc, img, tileSize, x0, x1, y0, y1)
	return pc.Wait(ctx)
}

func (t *tiledMask) renderInner(pc *parallelContext, img *image.Alpha, tileSize, x0, x1, y0, y1 int) {
	defer pc.Done()
	for ty := y0; ty < y1; ty += tileSize {
		for tx := x0; tx < x1; tx += tileSize {
			if b, ok := (*t)[image.Pt(tx, ty)]; ok {
				blend.Copy.Draw(img, b.Rect, b, b.Rect.Min)
			} else {
				blend.Clear.Draw(img, image.Rect(tx, ty, tx+tileSize, ty+tileSize), image.Transparent, image.Point{})
			}
		}
	}
}

func createMask(rect image.Rectangle, a []byte) *image.Alpha {
	img := image.NewAlpha(rect)
	w, h := rect.Dx(), rect.Dy()
	pix := img.Pix
	var y, sx0, sx1 int
	for y = 0; y < h; y++ {
		sx0 = y * w
		sx1 = sx0 + w
		for sx0 < sx1 {
			pix[sx0] = a[sx0]
			sx0++
		}
	}
	return img
}

func (t *tiledMask) Rescale(ctx context.Context, tileSize int, rect image.Rectangle, a []byte, defaultColor int, scale float64) (*tiledMask, image.Rectangle, error) {
	tmp := image.NewAlpha(rect)
	if err := t.Render(ctx, tileSize, tmp); err != nil {
		return nil, image.Rectangle{}, err
	}
	scaledRect := scaleRect(rect, scale)
	tmp2 := image.NewAlpha(scaledRect)
	// TODO: currently, it seems fallback path is used in image.Alpha.
	draw.BiLinear.Transform(tmp2, f64.Aff3{scale, 0, 0, 0, scale, 0}, tmp, rect, draw.Src, nil)
	var t2 tiledMask
	if err := t.Store(ctx, tileSize, scaledRect, tmp2.Pix, defaultColor); err != nil {
		return nil, image.Rectangle{}, err
	}
	return &t2, scaledRect, nil
}

func (t *tiledMask) StoreScaled(ctx context.Context, tileSize int, rect image.Rectangle, a []byte, defaultColor int, scale float64) (image.Rectangle, error) {
	if *t == nil {
		*t = make(tiledMask)
	}

	tmp := createMask(rect, a)
	scaledRect := scaleRect(rect, scale)
	tmp2 := image.NewAlpha(scaledRect)
	// TODO: currently, it seems fallback path is used in image.Alpha.
	draw.BiLinear.Transform(tmp2, f64.Aff3{scale, 0, 0, 0, scale, 0}, tmp, rect, draw.Src, nil)
	if err := t.Store(ctx, tileSize, scaledRect, tmp2.Pix, defaultColor); err != nil {
		return image.Rectangle{}, err
	}
	return scaledRect, nil
}

func (t *tiledMask) Store(ctx context.Context, tileSize int, rect image.Rectangle, a []byte, defaultColor int) error {
	if *t == nil {
		*t = make(tiledMask)
	}

	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := (ylen / n) * tileSize
	for i := 1; i < n; i++ {
		go t.storeInner(pc, rect, tileSize, x0, x1, y0, y0+step, a, defaultColor)
		y0 += step
	}
	go t.storeInner(pc, rect, tileSize, x0, x1, y0, y1, a, defaultColor)
	return pc.Wait(ctx)

}

func (t *tiledMask) storeInner(pc *parallelContext, rect image.Rectangle, tileSize, x0, x1, y0, y1 int, a []byte, defaultColor int) {
	defer pc.Done()

	rx0, ry0, rx1, ry1 := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	rw := rect.Dx()

	buf := make([]byte, tileSize*tileSize)
	for ty := y0; ty < y1; ty += tileSize {
		if pc.Aborted() {
			return
		}
		for tx := x0; tx < x1; tx += tileSize {
			dxMin, dxMax := 0, tileSize
			sxMin := tx - rx0
			if sxMin < 0 {
				dxMin -= sxMin
				sxMin = 0
			}
			if rx0+sxMin+(dxMax-dxMin) > rx1 {
				dxMax -= rx0 + sxMin + (dxMax - dxMin) - rx1
			}

			dyMin, dyMax := 0, tileSize
			syMin := ty - ry0
			if syMin < 0 {
				dyMin -= syMin
				syMin = 0
			}
			if ry0+syMin+(dyMax-dyMin) > ry1 {
				dyMax -= ry0 + syMin + (dyMax - dyMin) - ry1
			}

			used := false
			dyMax = dyMax * tileSize
			if defaultColor == 0 {
				for dy, sy := dyMin*tileSize, syMin*rw; dy < dyMax; dy += tileSize {
					for dx, sx, dEnd := dy+dxMin, sy+sxMin, dy+dxMax; dx < dEnd; dx++ {
						alpha := a[sx]
						if alpha != 0 {
							buf[dx] = alpha
							used = true
						}
						sx++
					}
					sy += rw
				}
			} else {
				for dy, sy := dyMin*tileSize, syMin*rw; dy < dyMax; dy += tileSize {
					for dx, sx, dEnd := dy+dxMin, sy+sxMin, dy+dxMax; dx < dEnd; dx++ {
						alpha := a[sx]
						if alpha != 255 {
							buf[dx] = 255 - alpha
							used = true
						}
						sx++
					}
					sy += rw
				}
			}
			if used {
				pc.M.Lock()
				(*t)[image.Pt(tx, ty)] = &image.Alpha{
					Pix:    buf,
					Stride: tileSize,
					Rect:   image.Rect(tx, ty, tx+tileSize, ty+tileSize),
				}
				pc.M.Unlock()
				buf = make([]byte, tileSize*tileSize)
			}
		}
	}
}
