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

type tiledImage map[image.Point]draw.Image

func (t tiledImage) Get(tileSize int, pt image.Point) (draw.Image, bool) {
	r, ok := t[pt]
	if !ok {
		r = image.NewRGBA(image.Rect(pt.X, pt.Y, pt.X+tileSize, pt.Y+tileSize))
		t[pt] = r
	}
	return r, ok
}

func (t tiledImage) Rect() image.Rectangle {
	var r image.Rectangle
	for _, img := range t {
		r = img.Bounds().Union(r)
	}
	return r
}

func (t tiledImage) Render(ctx context.Context, img draw.Image) error {
	tileSize := t.tileSize()
	if tileSize == 0 {
		return nil
	}

	rect := img.Bounds()
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

func (t tiledImage) renderInner(pc *parallelContext, img draw.Image, tileSize, x0, x1, y0, y1 int) {
	defer pc.Done()
	for ty := y0; ty < y1; ty += tileSize {
		for tx := x0; tx < x1; tx += tileSize {
			if b, ok := t[image.Pt(tx, ty)]; ok {
				rect := b.Bounds()
				blend.Copy.Draw(img, rect, b, rect.Min)
			} else {
				blend.Clear.Draw(img, image.Rect(tx, ty, tx+tileSize, ty+tileSize), image.Transparent, image.Point{})
			}
		}
	}
}

type gammaTable struct {
	T8  [256]uint16
	T16 [65536]uint8
}

func makeGammaTable(g float64) *gammaTable {
	var t [256]uint16
	for i := range t {
		t[i] = uint16(math.Pow(float64(i)/255, g) * 65535)
	}

	g = 1.0 / g
	var rt [65536]uint8
	for i := range rt {
		rt[i] = uint8(math.Pow(float64(i)/65535, g) * 255)
	}
	return &gammaTable{t, rt}
}

func createImage(rect image.Rectangle, r []byte, g []byte, b []byte, a []byte, deltaX int) *image.NRGBA {
	if deltaX == 4 {
		return &image.NRGBA{
			Pix:    r,
			Stride: rect.Dx() * 4,
			Rect:   rect,
		}
	}
	w, h := rect.Dx(), rect.Dy()
	pix := make([]byte, w*4*h)
	var s, d int
	if a != nil {
		for d < len(pix) {
			if a[s] > 0 {
				pix[d+3] = a[s]
				pix[d+2] = b[s]
				pix[d+1] = g[s]
				pix[d+0] = r[s]
			}
			d += 4
			s += deltaX
		}
	} else {
		for d < len(pix) {
			pix[d+3] = 0xff
			pix[d+2] = b[s]
			pix[d+1] = g[s]
			pix[d+0] = r[s]
			d += 4
			s += deltaX
		}
	}
	return &image.NRGBA{
		Pix:    pix,
		Stride: w * 4,
		Rect:   rect,
	}
}

func createImageGamma(rect image.Rectangle, r []byte, g []byte, b []byte, a []byte, deltaX int, gt [256]uint16) *image.NRGBA64 {
	w, h := rect.Dx(), rect.Dy()
	pix := make([]byte, w*8*h)
	var s, d int
	if a != nil {
		for d < len(pix) {
			if a[s] > 0 {
				a8, r16, g16, b16 := a[s], gt[r[s]], gt[g[s]], gt[b[s]]
				pix[d+7] = a8
				pix[d+6] = a8
				pix[d+5] = uint8(b16)
				pix[d+4] = uint8(b16 >> 8)
				pix[d+3] = uint8(g16)
				pix[d+2] = uint8(g16 >> 8)
				pix[d+1] = uint8(r16)
				pix[d+0] = uint8(r16 >> 8)
			}
			d += 8
			s += deltaX
		}
	} else {
		for d < len(pix) {
			r16, g16, b16 := gt[r[s]], gt[g[s]], gt[b[s]]
			pix[d+7] = 0xff
			pix[d+6] = 0xff
			pix[d+5] = uint8(b16)
			pix[d+4] = uint8(b16 >> 8)
			pix[d+3] = uint8(g16)
			pix[d+2] = uint8(g16 >> 8)
			pix[d+1] = uint8(r16)
			pix[d+0] = uint8(r16 >> 8)
			d += 8
			s += deltaX
		}
	}
	return &image.NRGBA64{
		Pix:    pix,
		Stride: w * 8,
		Rect:   rect,
	}
}

func restoreGamma(img *image.NRGBA64, gt [65536]uint8) {
	pix := img.Pix
	var s int
	for s < len(pix) {
		pix[s+4] = gt[(uint16(pix[s+4])<<8)|uint16(pix[s+5])]
		pix[s+2] = gt[(uint16(pix[s+2])<<8)|uint16(pix[s+3])]
		pix[s+0] = gt[(uint16(pix[s+0])<<8)|uint16(pix[s+1])]
		s += 8
	}
}

func (t tiledImage) tileSize() int {
	for _, m := range t {
		return m.Bounds().Dx()
	}
	return 0
}

func (t tiledImage) Transform(ctx context.Context, m f64.Aff3, gt *gammaTable) (tiledImage, error) {
	rect := t.Rect()
	if rect.Empty() {
		return tiledImage{}, nil
	}
	tileSize := t.tileSize()
	if tileSize == 0 {
		return tiledImage{}, nil
	}

	tmp := image.NewNRGBA(rect)
	if err := t.Render(ctx, tmp); err != nil {
		return nil, err
	}
	return newScaledTiledImage(ctx, tileSize, rect, tmp.Pix[0:], tmp.Pix[1:], tmp.Pix[2:], tmp.Pix[3:], 4, m, gt)
}

func newScaledTiledImage(ctx context.Context, tileSize int, rect image.Rectangle, r, g, b, a []byte, deltaX int, m f64.Aff3, gt *gammaTable) (tiledImage, error) {
	if m[0] == 1 && m[1] == 0 && m[2] == 0 && m[3] == 0 && m[4] == 1 && m[5] == 0 {
		return newTiledImage(ctx, tileSize, rect, r, g, b, a, deltaX)
	}
	if gt == nil {
		tmp := createImage(rect, r, g, b, a, deltaX)
		trRect := transformRect(rect, m)
		tmp2 := image.NewNRGBA(trRect)
		draw.BiLinear.Transform(tmp2, m, tmp, rect, draw.Src, nil)
		return newTiledImage(ctx, tileSize, trRect, tmp2.Pix[0:], tmp2.Pix[1:], tmp2.Pix[2:], tmp2.Pix[3:], 4)
	}
	tmp := createImageGamma(rect, r, g, b, a, deltaX, gt.T8)
	trRect := transformRect(rect, m)
	tmp2 := image.NewNRGBA64(trRect)
	draw.BiLinear.Transform(tmp2, m, tmp, rect, draw.Src, nil)
	restoreGamma(tmp2, gt.T16)
	return newTiledImage(ctx, tileSize, trRect, tmp2.Pix[0:], tmp2.Pix[2:], tmp2.Pix[4:], tmp2.Pix[6:], 8)
}

func newTiledImage(ctx context.Context, tileSize int, rect image.Rectangle, r, g, b, a []byte, deltaX int) (tiledImage, error) {
	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}

	t := tiledImage{}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := (ylen / n) * tileSize
	for i := 1; i < n; i++ {
		go newTiledImageInner(pc, t, rect, tileSize, x0, x1, y0, y0+step, r, g, b, a, deltaX)
		y0 += step
	}
	go newTiledImageInner(pc, t, rect, tileSize, x0, x1, y0, y1, r, g, b, a, deltaX)
	if err := pc.Wait(ctx); err != nil {
		return nil, err
	}
	return t, nil
}

func newTiledImageInner(pc *parallelContext, t tiledImage, rect image.Rectangle, tileSize, x0, x1, y0, y1 int, r, g, b, a []byte, deltaX int) {
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
				t[image.Pt(tx, ty)] = &image.RGBA{
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

func (t tiledMask) Rect() image.Rectangle {
	var r image.Rectangle
	for _, img := range t {
		r = img.Rect.Union(r)
	}
	return r
}

func (t tiledMask) tileSize() int {
	for _, m := range t {
		return m.Rect.Dx()
	}
	return 0
}

func (t tiledMask) Render(ctx context.Context, img draw.Image) error {
	tileSize := t.tileSize()
	if tileSize == 0 {
		return nil
	}

	rect := img.Bounds()
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

func (t tiledMask) renderInner(pc *parallelContext, img draw.Image, tileSize, x0, x1, y0, y1 int) {
	defer pc.Done()
	for ty := y0; ty < y1; ty += tileSize {
		for tx := x0; tx < x1; tx += tileSize {
			if b, ok := t[image.Pt(tx, ty)]; ok {
				blend.Copy.Draw(img, b.Rect, b, b.Rect.Min)
			} else {
				blend.Clear.Draw(img, image.Rect(tx, ty, tx+tileSize, ty+tileSize), image.Transparent, image.Point{})
			}
		}
	}
}

func createMask(rect image.Rectangle, a []byte) *image.Alpha {
	return &image.Alpha{
		Pix:    a,
		Stride: rect.Dx(),
		Rect:   rect,
	}
}

func (t tiledMask) Transform(ctx context.Context, m f64.Aff3) (tiledMask, error) {
	rect := t.Rect()
	if rect.Empty() {
		return tiledMask{}, nil
	}
	tileSize := t.tileSize()
	if tileSize == 0 {
		return tiledMask{}, nil
	}
	tmp := image.NewAlpha(rect)
	if err := t.Render(ctx, tmp); err != nil {
		return nil, err
	}
	return newScaledTiledMask(ctx, tileSize, rect, tmp.Pix, 0, m)
}

func newScaledTiledMask(ctx context.Context, tileSize int, rect image.Rectangle, a []byte, defaultColor int, m f64.Aff3) (tiledMask, error) {
	if m[0] == 1 && m[1] == 0 && m[2] == 0 && m[3] == 0 && m[4] == 1 && m[5] == 0 {
		return newTiledMask(ctx, tileSize, rect, a, defaultColor)
	}
	tmp := createMask(rect, a)
	trRect := transformRect(rect, m)
	tmp2 := image.NewAlpha(trRect)
	// TODO: currently, it seems fallback path is used in image.Alpha.
	draw.BiLinear.Transform(tmp2, m, tmp, rect, draw.Src, nil)
	return newTiledMask(ctx, tileSize, trRect, tmp2.Pix, defaultColor)
}

func newTiledMask(ctx context.Context, tileSize int, rect image.Rectangle, a []byte, defaultColor int) (tiledMask, error) {
	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	t := tiledMask{}
	pc := &parallelContext{}
	pc.Wg.Add(n)
	step := (ylen / n) * tileSize
	for i := 1; i < n; i++ {
		go newTiledMaskInner(pc, t, rect, tileSize, x0, x1, y0, y0+step, a, defaultColor)
		y0 += step
	}
	go newTiledMaskInner(pc, t, rect, tileSize, x0, x1, y0, y1, a, defaultColor)
	if err := pc.Wait(ctx); err != nil {
		return nil, err
	}
	return t, nil
}

func newTiledMaskInner(pc *parallelContext, t tiledMask, rect image.Rectangle, tileSize, x0, x1, y0, y1 int, a []byte, defaultColor int) {
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
				t[image.Pt(tx, ty)] = &image.Alpha{
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
