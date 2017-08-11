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
	drAbort                  drawResult = -1
	drDrew                   drawResult = 0
	drDrewFromCache          drawResult = 1
	drDrewFromCacheInvisible drawResult = 2
	drDefered                drawResult = 3
	drSkipped                drawResult = 4
	drInvisible              drawResult = 5
)

type cacheState bool

const (
	csCached      cacheState = false
	csCachedEmpty cacheState = true
)

type cachedMap map[image.Point]cacheState
type cache struct {
	Image  tiledImage
	Cached cachedMap
	M      sync.RWMutex
}

// Renderer is a renderer.
type Renderer struct {
	layertree *Root
	renderM   sync.Mutex

	cache           map[int]*cache
	deferedDirtyReq []*Layer
	setDirtyM       sync.Mutex
	cacheM          sync.Mutex

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
	if b == nil {
		return
	}
	buf := b.Pix
	for i := range buf {
		buf[i] = 0
	}
	r.pool.Put(buf)
}

func (r *Renderer) getCache(l *Layer, createIfNotExists bool) *cache {
	var seqID int
	if l != nil {
		seqID = l.SeqID
	} else {
		seqID = -1
	}
	r.cacheM.Lock()
	c, ok := r.cache[seqID]
	if !ok && createIfNotExists {
		c = &cache{
			Image:  tiledImage{},
			Cached: cachedMap{},
		}
		r.cache[seqID] = c
	}
	r.cacheM.Unlock()
	return c
}

func (r *Renderer) buildAreaMap(l *Layer) []image.Point {
	var ret []image.Point
	ld := r.layertree.layerImage[l.SeqID]
	if ld.Canvas != nil {
		ret = make([]image.Point, len(ld.Canvas))
		i := 0
		for pt := range ld.Canvas {
			ret[i] = pt
			i++
		}
	} else {
		rect, tileSize := l.Rect, r.layertree.tileSize
		rx0, ry0, rx1, ry1 := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
		for ty := (ry0 / tileSize) * tileSize; ty < ry1; ty += tileSize {
			for tx := (rx0 / tileSize) * tileSize; tx < rx1; tx += tileSize {
				ret = append(ret, image.Pt(tx, ty))
			}
		}
	}
	return ret
}

func (r *Renderer) setDirtyByLayerPassThrough(l *Layer, areaMap []image.Point) {
	if c := r.getCache(l, false); c != nil {
		c.M.Lock()
		for _, pt := range areaMap {
			delete(c.Cached, pt)
		}
		c.M.Unlock()
	}
	for i := range l.Children {
		l2 := &l.Children[i]
		if l2.BlendMode == psd.BlendModePassThrough {
			r.setDirtyByLayerPassThrough(l2, areaMap)
		}
	}
}

func (r *Renderer) setDirtyByLayerRecursive(l *Layer, areaMap []image.Point) {
	if c := r.getCache(l, false); c != nil {
		c.M.Lock()
		for _, pt := range areaMap {
			delete(c.Cached, pt)
		}
		c.M.Unlock()
	}
	if l != nil {
		if l.Parent != nil {
			sib := l.Parent.Children
			for i := len(sib) - 1; i >= 0; i-- {
				l2 := &sib[i]
				if l == l2 {
					break
				}
				if l2.BlendMode != psd.BlendModePassThrough {
					continue
				}
				r.setDirtyByLayerPassThrough(l2, areaMap)
			}
		}
		r.setDirtyByLayerRecursive(l.Parent, areaMap)
	}
}

func (r *Renderer) setDirtyByLayer(l *Layer) {
	areaMap := r.buildAreaMap(l)
	if l.Clipping {
		r.setDirtyByLayerRecursive(l.ClippedBy, areaMap)
		return
	}
	r.setDirtyByLayerRecursive(l, areaMap)
}

func (r *Renderer) SetDirtyByLayer(l *Layer) {
	r.setDirtyM.Lock()
	defer r.setDirtyM.Unlock()
	if r.deferedDirtyReq != nil {
		r.deferedDirtyReq = append(r.deferedDirtyReq, l)
		return
	}
	r.setDirtyByLayer(l)
}

type changeList struct {
	Map map[int]*[]image.Point
	M   sync.Mutex
}

func (cl *changeList) Add(l *Layer, pt image.Point) {
	var seqID int
	if l != nil {
		seqID = l.SeqID
	} else {
		seqID = -1
	}
	cl.M.Lock()
	if a, ok := cl.Map[seqID]; ok {
		*a = append(*a, pt)
	} else {
		a = &[]image.Point{pt}
		cl.Map[seqID] = a
	}
	cl.M.Unlock()
}

// Render renders image.
func (r *Renderer) Render(ctx context.Context) (*image.RGBA, error) {
	r.renderM.Lock()
	defer r.renderM.Unlock()

	img := image.NewRGBA(r.layertree.CanvasRect)
	rect := r.layertree.Rect
	tileSize := r.layertree.tileSize

	r.setDirtyM.Lock()
	r.deferedDirtyReq = make([]*Layer, 0)
	r.setDirtyM.Unlock()
	defer func() {
		r.setDirtyM.Lock()
		for _, l := range r.deferedDirtyReq {
			r.setDirtyByLayer(l)
		}
		r.deferedDirtyReq = nil
		r.setDirtyM.Unlock()
	}()

	c := r.getCache(nil, true)
	clst := &changeList{
		Map: map[int]*[]image.Point{},
	}

	x0, x1 := (rect.Min.X/tileSize)*tileSize, rect.Max.X
	y0, y1 := (rect.Min.Y/tileSize)*tileSize, rect.Max.Y
	ylen := (y1 - y0) / tileSize

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > ylen {
		n--
	}
	pc := &parallelContext{}
	step := (ylen / n) * tileSize

	pc.Wg.Add(n)
	for i := 1; i < n; i++ {
		go r.renderInner(pc, img, c, clst, x0, x1, y0, y0+step)
		y0 += step
	}
	go r.renderInner(pc, img, c, clst, x0, x1, y0, y1)
	if err := pc.Wait(ctx); err == ErrAborted {
		for seqID, pts := range clst.Map {
			r.cacheM.Lock()
			c := r.cache[seqID]
			r.cacheM.Unlock()
			c.M.Lock()
			for _, pt := range *pts {
				delete(c.Cached, pt)
			}
			c.M.Unlock()
		}
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return img, nil
}

func (r *Renderer) renderInner(pc *parallelContext, img *image.RGBA, c *cache, clst *changeList, x0, x1, y0, y1 int) {
	defer pc.Done()
	tileSize := r.layertree.tileSize
	for ty := y0; ty < y1; ty += tileSize {
		if pc.Aborted() {
			return
		}
		for tx := x0; tx < x1; tx += tileSize {
			pt := image.Pt(tx, ty)
			c.M.RLock()
			cs, ok := c.Cached[pt]
			c.M.RUnlock()
			if !ok {
				ldBuffer := r.getBuffer(pt)
				dr := r.renderTile(ldBuffer, c, clst, pt)
				r.updateCache(dr, pt, ldBuffer, c, clst)
				switch {
				case dr < 0:
					return
				case dr == drDrew || dr == drDrewFromCache:
					blend.Copy.Draw(img, ldBuffer.Rect, ldBuffer, pt)
				}
				r.putBuffer(ldBuffer)
			} else if cs == csCached {
				c.M.RLock()
				img2 := c.Image[pt]
				c.M.RUnlock()
				blend.Copy.Draw(img, img2.Rect, img2, pt)
			}
		}
	}
}

func (r *Renderer) updateCache(dr drawResult, pt image.Point, b *image.RGBA, c *cache, clst *changeList) drawResult {
	switch dr {
	case drAbort:
		c.M.Lock()
		delete(c.Cached, pt)
		c.M.Unlock()
	case drDrew:
		c.M.Lock()
		c.Cached[pt] = csCached
		cb, _ := c.Image.Get(r.layertree.tileSize, pt)
		blend.Copy.Draw(cb, cb.Rect, b, pt)
		c.M.Unlock()
		clst.Add(nil, pt)
	case drDrewFromCache, drDrewFromCacheInvisible:
		// do nothing
	default:
		c.M.Lock()
		c.Cached[pt] = csCachedEmpty
		delete(c.Image, pt)
		c.M.Unlock()
		clst.Add(nil, pt)
	}
	return dr
}

func (r *Renderer) renderTile(b *image.RGBA, c *cache, clst *changeList, pt image.Point) drawResult {
	c.M.RLock()
	cs, ok := c.Cached[pt]
	if ok {
		if cs == csCached {
			blend.Copy.Draw(b, b.Rect, c.Image[pt], pt)
			c.M.RUnlock()
			return drDrewFromCache
		}
		c.M.RUnlock()
		return drDrewFromCacheInvisible
	}
	c.M.RUnlock()

	var drew bool
	for _, l := range r.layertree.Children {
		switch dr := r.drawLayer(pt, b, clst, &l, l.Opacity, l.BlendMode, false); {
		case dr < 0:
			return dr
		case dr == drDrew || dr == drDrewFromCache:
			drew = true
		}
	}
	if !drew {
		return drInvisible
	}
	return drDrew
}

func (r *Renderer) drawLayer(pt image.Point, b *image.RGBA, clst *changeList, l *Layer, opacity int, blendMode psd.BlendMode, forceNoClip bool) drawResult {
	var fc *cache
	var ldBuffer *image.RGBA
	ret := func(dr drawResult) drawResult {
		if fc == nil {
			if len(l.Children) > 0 {
				fc = r.getCache(l, true)
			} else {
				return dr
			}
		}
		switch dr {
		case drAbort:
			fc.M.Lock()
			delete(fc.Cached, pt)
			fc.M.Unlock()
		case drDrew:
			fc.M.Lock()
			fc.Cached[pt] = csCached
			cb, _ := fc.Image.Get(r.layertree.tileSize, pt)
			fc.M.Unlock()
			clst.Add(l, pt)
			blend.Copy.Draw(cb, cb.Rect, ldBuffer, pt)
		case drDrewFromCache, drDrewFromCacheInvisible:
			// do nothing
		default:
			fc.M.Lock()
			fc.Cached[pt] = csCachedEmpty
			delete(fc.Image, pt)
			fc.M.Unlock()
			clst.Add(l, pt)
		}
		return dr
	}

	if !l.Visible || opacity == 0 {
		return ret(drInvisible)
	}
	if l.Clipping && !forceNoClip {
		return ret(drDefered)
	}

	ld := r.layertree.layerImage[l.SeqID]
	ldCanvas, ok := ld.Canvas[pt]
	if len(l.Children) == 0 && !ok {
		return ret(drSkipped)
	}

	if len(l.Children) > 0 {
		if fc = r.getCache(l, false); fc != nil {
			fc.M.RLock()
			if cs, ok := fc.Cached[pt]; ok {
				if cs == csCached {
					cb, _ := fc.Image.Get(r.layertree.tileSize, pt)
					if blendMode == psd.BlendModePassThrough {
						blend.Copy.Draw(b, b.Rect, cb, pt)
					} else {
						drawWithOpacity(b, b.Rect, cb, pt, opacity, blendMode)
					}
					fc.M.RUnlock()
					return ret(drDrewFromCache)
				}
				fc.M.RUnlock()
				return ret(drDrewFromCacheInvisible)
			}
			fc.M.RUnlock()
		}

		ldBuffer = r.getBuffer(pt)
		defer r.putBuffer(ldBuffer)

		var a int
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(ldBuffer, ldBuffer.Rect, b, pt)
			a = opacity * 32897
		} else {
			a = 255 * 32897
		}
		drew := false
		for _, l2 := range l.Children {
			switch dr := r.drawLayer(pt, ldBuffer, clst, &l2, l2.Opacity*a>>23, l2.BlendMode, false); {
			case dr < 0:
				return ret(dr)
			case dr == drDrew || dr == drDrewFromCache:
				drew = true
			}
		}
		if !drew {
			return ret(drInvisible)
		}
	} else if ld.Canvas != nil {
		ldBuffer = r.getBuffer(pt)
		blend.Copy.Draw(ldBuffer, ldCanvas.Rect, ldCanvas, pt)
	}

	if l.MaskEnabled {
		if ldMask, ok := ld.Mask[pt]; ok {
			if l.MaskDefaultColor == 255 {
				blend.DestOut.Draw(ldBuffer, ldBuffer.Rect, ldMask, pt)
			} else {
				blend.DestIn.Draw(ldBuffer, ldBuffer.Rect, ldMask, pt)
			}
		} else {
			// ???
		}
		if blendMode == psd.BlendModePassThrough {
			blend.DestOver.Draw(ldBuffer, ldBuffer.Rect, b, pt)
		}
	}

	if len(l.Clip) == 0 {
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(b, b.Rect, ldBuffer, pt)
		} else {
			drawWithOpacity(b, ldBuffer.Rect, ldBuffer, pt, opacity, blendMode)
		}
		return ret(drDrew)
	}

	ldClipBuffer := r.getBuffer(pt)
	defer r.putBuffer(ldClipBuffer)

	if l.BlendClippedElements {
		blend.Copy.Draw(ldClipBuffer, ldBuffer.Rect, ldBuffer, pt)
		removeAlpha(ldClipBuffer)
		for _, cl := range l.Clip {
			dr := r.drawLayer(pt, ldClipBuffer, clst, cl, cl.Opacity, cl.BlendMode, true)
			if dr < 0 {
				return ret(dr)
			}
		}
		blend.DestIn.Draw(ldClipBuffer, ldBuffer.Rect, ldBuffer, pt)
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(b, ldClipBuffer.Rect, ldClipBuffer, pt)
		} else {
			drawWithOpacity(b, ldClipBuffer.Rect, ldClipBuffer, pt, opacity, blendMode)
		}
		ldClipBuffer, ldBuffer = ldBuffer, ldClipBuffer
		return ret(drDrew)
	}

	// this is minor code path.
	// it is only used when "Blend Clipped Layers as Group" is unchecked in Photoshop's Layer Style dialog.
	// TODO: implement
	drawWithOpacity(b, ldBuffer.Rect, ldBuffer, pt, opacity, blendMode)
	for _, cl := range l.Clip {
		if dr := r.drawLayer(pt, b, clst, cl, cl.Opacity, cl.BlendMode, false); dr < 0 {
			return ret(dr)
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
	return ret(drDrew)
}
