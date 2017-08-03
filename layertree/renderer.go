package layertree

import (
	"context"
	"image"
	"runtime"
	"sync"

	"github.com/oov/psd"
	"github.com/oov/psd/blend"
)

type drawResult int

const (
	drAbort     drawResult = -1
	drDrew      drawResult = 0
	drDefered   drawResult = 1
	drSkipped   drawResult = 2
	drInvisible drawResult = 3
)

// Renderer is a renderer.
type Renderer struct {
	layertree *Root

	rootImage  tiledImage
	rootImageM sync.Mutex
	cached     map[image.Point]struct{}
	cacheM     sync.RWMutex

	pool sync.Pool
}

func (r *Renderer) allocate() interface{} {
	return make([]byte, r.layertree.tileSize*r.layertree.tileSize*4)
}

func (r *Renderer) getBuffer(pt image.Point) *image.RGBA {
	return &image.RGBA{
		Pix:    r.pool.Get().([]byte),
		Stride: r.layertree.tileSize * 4,
		Rect:   image.Rect(pt.X, pt.Y, pt.X+r.layertree.tileSize, pt.Y+r.layertree.tileSize),
	}
}

func (r *Renderer) putBuffer(b *image.RGBA) {
	buf := b.Pix
	for i := range buf {
		buf[i] = 0
	}
	r.pool.Put(buf)
}

func (r *Renderer) SetDirty(rect image.Rectangle) {
	tileSize := r.layertree.tileSize
	r.cacheM.Lock()
	rx0, ry0, rx1, ry1 := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	for ty := (ry0 / tileSize) * tileSize; ty < ry1; ty += tileSize {
		for tx := (rx0 / tileSize) * tileSize; tx < rx1; tx += tileSize {
			delete(r.cached, image.Pt(tx, ty))
		}
	}
	r.cacheM.Unlock()
}

// Render renders image.
func (r *Renderer) Render(ctx context.Context) (*image.RGBA, error) {
	r.rootImageM.Lock()
	defer r.rootImageM.Unlock()

	r.cacheM.RLock()
	cached := map[image.Point]struct{}{}
	for pt := range r.cached {
		cached[pt] = struct{}{}
	}
	r.cacheM.RUnlock()

	img := image.NewRGBA(r.layertree.CanvasRect)
	rect := r.layertree.Rect
	tileSize := r.layertree.tileSize
	r.rootImage.Alloc(tileSize, rect)
	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	modified := map[image.Point]struct{}{}
	pc := &parallelContext{}
	step := (ylen / n) * tileSize

	pc.Wg.Add(n)
	for i := 1; i < n; i++ {
		go r.renderInner(pc, img, cached, modified, x0, x1, y0, y0+step)
		y0 += step
	}
	go r.renderInner(pc, img, cached, modified, x0, x1, y0, y1)
	if err := pc.Wait(ctx); err != nil {
		return nil, err
	}

	r.cacheM.Lock()
	for pt := range modified {
		r.cached[pt] = struct{}{}
	}
	r.cacheM.Unlock()
	return img, nil
}

func (r *Renderer) renderInner(pc *parallelContext, img *image.RGBA, cached, modified map[image.Point]struct{}, x0, x1, y0, y1 int) {
	defer pc.Done()
	tileSize := r.layertree.tileSize
	for ty := y0; ty < y1; ty += tileSize {
		if pc.Aborted() {
			return
		}
		for tx := x0; tx < x1; tx += tileSize {
			pt := image.Pt(tx, ty)
			if _, ok := cached[pt]; !ok {
				r.renderTile(r.rootImage, pt)
				pc.M.Lock()
				modified[pt] = struct{}{}
				pc.M.Unlock()
			}
			blend.Copy.Draw(img, image.Rect(tx, ty, tx+tileSize, ty+tileSize), r.rootImage[pt], pt)
		}
	}
}

func (r *Renderer) renderTile(t tiledImage, pt image.Point) drawResult {
	buffer, ok := t[pt]
	if !ok {
		return drSkipped
	}
	blend.Clear.Draw(buffer, buffer.Rect, image.Transparent, image.Point{})
	for _, l := range r.layertree.Children {
		if dr := r.drawLayer(pt, buffer, &l, l.Opacity, l.BlendMode, false); dr < 0 {
			return dr
		}
	}
	return drDrew
}

func (r *Renderer) drawLayer(pt image.Point, b *image.RGBA, l *Layer, opacity int, blendMode psd.BlendMode, forceNoClip bool) drawResult {
	if !l.Visible || opacity == 0 {
		return drInvisible
	}
	if l.Clipping && !forceNoClip {
		return drDefered
	}

	ld := r.layertree.layerImage[l.SeqID]
	ldCanvas, ok := ld.Canvas[pt]
	if len(l.Children) == 0 && !ok {
		return drSkipped
	}

	ldBuffer := r.getBuffer(pt)
	defer r.putBuffer(ldBuffer)
	if len(l.Children) > 0 {
		var a int
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(ldBuffer, ldBuffer.Rect, b, ldBuffer.Rect.Min)
			a = opacity * 32897
		} else {
			a = 255 * 32897
		}
		for _, l2 := range l.Children {
			if dr := r.drawLayer(pt, ldBuffer, &l2, l2.Opacity*a>>23, l2.BlendMode, false); dr < 0 {
				return dr
			}
		}
	} else if ld.Canvas != nil {
		blend.Copy.Draw(ldBuffer, ldCanvas.Rect, ldCanvas, ldCanvas.Rect.Min)
	}

	if l.MaskEnabled {
		if ldMask, ok := ld.Mask[pt]; ok {
			if l.MaskDefaultColor == 255 {
				blend.DestOut.Draw(ldBuffer, ldBuffer.Rect, ldMask, ldMask.Rect.Min)
			} else {
				blend.DestIn.Draw(ldBuffer, ldBuffer.Rect, ldMask, ldMask.Rect.Min)
			}
		} else {
			// ???
		}
		if blendMode == psd.BlendModePassThrough {
			blend.DestOver.Draw(ldBuffer, ldBuffer.Rect, b, ldBuffer.Rect.Min)
		}
	}

	if len(l.Clip) == 0 {
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(b, b.Rect, ldBuffer, ldBuffer.Rect.Min)
		} else {
			drawWithOpacity(b, ldBuffer.Rect, ldBuffer, ldBuffer.Rect.Min, opacity, blendMode)
		}
		return drDrew
	}

	ldClipBuffer := r.getBuffer(pt)
	defer r.putBuffer(ldClipBuffer)

	if l.BlendClippedElements {
		blend.Copy.Draw(ldClipBuffer, ldBuffer.Rect, ldBuffer, ldBuffer.Rect.Min)
		removeAlpha(ldClipBuffer)
		for _, cl := range l.Clip {
			dr := r.drawLayer(pt, ldClipBuffer, cl, cl.Opacity, cl.BlendMode, true)
			if dr < 0 {
				return dr
			}
		}
		blend.DestIn.Draw(ldClipBuffer, ldBuffer.Rect, ldBuffer, ldBuffer.Rect.Min)
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(b, ldClipBuffer.Rect, ldClipBuffer, ldClipBuffer.Rect.Min)
		} else {
			drawWithOpacity(b, ldClipBuffer.Rect, ldClipBuffer, ldClipBuffer.Rect.Min, opacity, blendMode)
		}
		return drDrew
	}

	// this is minor code path.
	// it is only used when "Blend Clipped Layers as Group" is unchecked in Photoshop's Layer Style dialog.
	// TODO: implement
	drawWithOpacity(b, ldBuffer.Rect, ldBuffer, ldBuffer.Rect.Min, opacity, blendMode)
	for _, cl := range l.Clip {
		if dr := r.drawLayer(pt, b, cl, cl.Opacity, cl.BlendMode, false); dr < 0 {
			return dr
		}
	}

	// this.draw(ctx, bb, x + n.layer.X, y + n.layer.Y, opacity, blendMode);
	// this.clear(cbbctx);
	// for (let cn of n.clip) {
	// 	if (!this.drawLayer(cbbctx, cn, -n.layer.X, -n.layer.Y, 1, 'source-over')) {
	// 		continue;
	// 	}
	// 	this.draw(cbbctx, bb, 0, 0, 1, 'destination-in');
	// 	this.draw(ctx, cbb, x + n.layer.X, y + n.layer.Y, cn.layer.Opacity / 255, cn.layer.BlendMode);
	// 	this.clear(cbbctx);
	// }
	// n.state = n.nextState;
	return drDrew
}
