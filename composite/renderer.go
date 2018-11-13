package composite

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
	tree    *Tree
	renderM sync.Mutex

	cache           map[int]*cache
	deferedDirtyReq []*Layer
	setDirtyM       sync.Mutex
	cacheM          sync.Mutex

	pool sync.Pool
}

func (r *Renderer) allocate() interface{} {
	return make([]byte, r.tree.tileSize*r.tree.tileSize*4)
}

func (r *Renderer) getBuffer(pt image.Point) *image.NRGBA {
	return &image.NRGBA{
		Pix:    r.pool.Get().([]byte),
		Stride: r.tree.tileSize * 4,
		Rect:   image.Rect(pt.X, pt.Y, pt.X+r.tree.tileSize, pt.Y+r.tree.tileSize),
	}
}

func (r *Renderer) putBuffer(b *image.NRGBA) {
	if b == nil {
		return
	}
	buf := b.Pix
	for i := range buf {
		buf[i] = 0
	}
	r.pool.Put(buf)
}

func (r *Renderer) getCache(seqID int, createIfNotExists bool) *cache {
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

func (r *Renderer) buildAreaMapRecursive(areaMap map[image.Point]struct{}, l *Layer) {
	for i := range l.Children {
		l2 := &l.Children[i]
		if !l2.Visible || l2.Opacity == 0 || l2.Clipping {
			continue
		}
		ld := r.tree.layerImage[l2.SeqID]
		if ld.Canvas != nil {
			for pt := range ld.Canvas {
				areaMap[pt] = struct{}{}
			}
		} else if len(l2.Children) > 0 {
			r.buildAreaMapRecursive(areaMap, l2)
		}
	}
}

func (r *Renderer) buildAreaMap(l *Layer) []image.Point {
	var ret []image.Point
	ld := r.tree.layerImage[l.SeqID]
	if ld.Canvas != nil {
		ret = make([]image.Point, len(ld.Canvas))
		i := 0
		for pt := range ld.Canvas {
			ret[i] = pt
			i++
		}
	} else if len(l.Children) > 0 {
		areaMap := map[image.Point]struct{}{}
		r.buildAreaMapRecursive(areaMap, l)
		ret = make([]image.Point, 0, len(areaMap))
		for pt := range areaMap {
			ret = append(ret, pt)
		}
	}
	return ret
}

func (r *Renderer) setDirtyByLayerPassThrough(l *Layer, areaMap []image.Point) {
	if c := r.getCache(l.SeqID, false); c != nil {
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
	if c := r.getCache(l.SeqID, false); c != nil {
		c.M.Lock()
		for _, pt := range areaMap {
			delete(c.Cached, pt)
		}
		c.M.Unlock()
	}
	if l.Clipping {
		if l.ClippedBy != nil {
			r.setDirtyByLayerRecursive(l.ClippedBy, areaMap)
		}
	} else if l.Parent != nil {
		sib := l.Parent.Children
		for i := len(sib) - 1; i >= 0; i-- {
			l2 := &sib[i]
			if l2.BlendMode != psd.BlendModePassThrough {
				continue
			}
			r.setDirtyByLayerPassThrough(l2, areaMap)
		}
		r.setDirtyByLayerRecursive(l.Parent, areaMap)
	}
}

func (r *Renderer) setDirtyByLayer(l *Layer) {
	areaMap := r.buildAreaMap(l)
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
func (r *Renderer) Render(ctx context.Context, dest *image.RGBA) error {
	return r.render(ctx, dest, false)
}

// RenderDiff renders only the place changed since last time.
func (r *Renderer) RenderDiff(ctx context.Context, dest *image.RGBA) error {
	return r.render(ctx, dest, true)
}

func (r *Renderer) render(ctx context.Context, dest *image.RGBA, diffOnly bool) error {
	r.renderM.Lock()
	defer r.renderM.Unlock()

	rect := r.tree.Rect
	tileSize := r.tree.tileSize

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

	rootCache := r.getCache(SeqIDRoot, true)
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
		go r.renderInner(pc, dest, diffOnly, rootCache, clst, x0, x1, y0, y0+step)
		y0 += step
	}
	go r.renderInner(pc, dest, diffOnly, rootCache, clst, x0, x1, y0, y1)
	if err := pc.Wait(ctx); err == ErrAborted {
		// revert
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
		return err
	} else if err != nil {
		return err
	}
	return nil
}

func (r *Renderer) renderInner(pc *parallelContext, dest *image.RGBA, diffOnly bool, rootCache *cache, clst *changeList, x0, x1, y0, y1 int) {
	defer pc.Done()
	tileSize := r.tree.tileSize
	for ty := y0; ty < y1; ty += tileSize {
		if pc.Aborted() {
			return
		}
		for tx := x0; tx < x1; tx += tileSize {
			pt := image.Pt(tx, ty)
			if modified, cs := r.updateTile(pt, rootCache, clst); modified || !diffOnly {
				if cs == csCached {
					rootCache.M.RLock()
					ci := rootCache.Image[pt]
					rootCache.M.RUnlock()
					blend.Copy.Draw(dest, ci.Bounds(), ci, pt)
				} else {
					blend.Clear.Draw(dest, image.Rect(pt.X, pt.Y, pt.X+tileSize, pt.Y+tileSize), image.Transparent, image.Point{})
				}
			}
		}
	}
}

func (r *Renderer) updateTile(pt image.Point, rootCache *cache, clst *changeList) (modified bool, cs cacheState) {
	rootCache.M.RLock()
	cs, ok := rootCache.Cached[pt]
	rootCache.M.RUnlock()
	if ok {
		return false, cs
	}

	buf := r.getBuffer(pt)
	drew := false
	for _, l := range r.tree.Root.Children {
		dr := r.drawLayer(pt, buf, clst, &l, l.Opacity, l.BlendMode, false)
		drew = drew || dr == drDrew || dr == drDrewFromCache
	}
	if drew {
		r.updateCache(drDrew, nil, pt, buf, rootCache, clst)
		r.putBuffer(buf)
		return true, csCached
	}
	r.updateCache(drInvisible, nil, pt, buf, rootCache, clst)
	r.putBuffer(buf)
	return true, csCachedEmpty
}

func (r *Renderer) updateCache(dr drawResult, l *Layer, pt image.Point, b *image.NRGBA, c *cache, clst *changeList) drawResult {
	if c == nil {
		return dr
	}
	switch dr {
	case drDrew:
		c.M.Lock()
		c.Cached[pt] = csCached
		ci, _ := c.Image.Get(r.tree.tileSize, pt)
		c.M.Unlock()
		blend.Copy.Draw(ci, ci.Bounds(), b, pt)
		clst.Add(l, pt)
	case drDrewFromCache, drDrewFromCacheInvisible:
		// do nothing
	default:
		c.M.Lock()
		c.Cached[pt] = csCachedEmpty
		delete(c.Image, pt)
		c.M.Unlock()
		clst.Add(l, pt)
	}
	return dr
}

func (r *Renderer) drawLayer(pt image.Point, b *image.NRGBA, clst *changeList, l *Layer, opacity int, blendMode psd.BlendMode, forceNoClip bool) drawResult {
	var c *cache
	var buf *image.NRGBA

	if !l.Visible || opacity == 0 {
		return r.updateCache(drInvisible, l, pt, buf, c, clst)
	}
	if l.Clipping && !forceNoClip {
		return r.updateCache(drDefered, l, pt, buf, c, clst)
	}

	ld := r.tree.layerImage[l.SeqID]
	ldCanvas, ok := ld.Canvas[pt]
	if len(l.Children) == 0 && !ok {
		return r.updateCache(drSkipped, l, pt, buf, c, clst)
	}

	if len(l.Children) > 0 {
		c = r.getCache(l.SeqID, true)
		c.M.RLock()
		if cs, ok := c.Cached[pt]; ok {
			if cs == csCached {
				ci := c.Image[pt]
				c.M.RUnlock()
				if blendMode == psd.BlendModePassThrough {
					blend.Copy.Draw(b, b.Rect, ci, pt)
				} else {
					drawWithOpacity(b, b.Rect, ci, pt, opacity, blendMode)
				}
				return r.updateCache(drDrewFromCache, l, pt, buf, c, clst)
			}
			c.M.RUnlock()
			return r.updateCache(drDrewFromCacheInvisible, l, pt, buf, c, clst)
		}
		c.M.RUnlock()

		buf = r.getBuffer(pt)
		defer r.putBuffer(buf)

		var a int
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(buf, buf.Rect, b, pt)
			a = opacity * 32897
		} else {
			a = 255 * 32897
		}
		drew := false
		for _, l2 := range l.Children {
			dr := r.drawLayer(pt, buf, clst, &l2, l2.Opacity*a>>23, l2.BlendMode, false)
			drew = drew || dr == drDrew || dr == drDrewFromCache
		}
		if !drew {
			return r.updateCache(drInvisible, l, pt, buf, c, clst)
		}
	} else if ld.Canvas != nil {
		buf = r.getBuffer(pt)
		blend.Copy.Draw(buf, ldCanvas.Bounds(), ldCanvas, pt)
	}

	if l.MaskEnabled {
		if ldMask, ok := ld.Mask[pt]; ok {
			if l.MaskDefaultColor != 255 {
				blend.DestIn.Draw(buf, buf.Rect, ldMask, pt)
			} else {
				blend.DestOut.Draw(buf, buf.Rect, ldMask, pt)
			}
		} else {
			if l.MaskDefaultColor != 255 {
				blend.Clear.Draw(buf, buf.Rect, image.Transparent, image.Point{})
			}
		}
		if blendMode == psd.BlendModePassThrough {
			blend.DestOver.Draw(buf, buf.Rect, b, pt)
		}
	}

	if len(l.Clip) == 0 {
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(b, b.Rect, buf, pt)
		} else {
			drawWithOpacity(b, buf.Rect, buf, pt, opacity, blendMode)
		}
		return r.updateCache(drDrew, l, pt, buf, c, clst)
	}

	clipBuf := r.getBuffer(pt)
	defer r.putBuffer(clipBuf)

	if l.BlendClippedElements {
		blend.Copy.Draw(clipBuf, buf.Rect, buf, pt)
		removeAlpha(clipBuf)
		for _, cl := range l.Clip {
			r.drawLayer(pt, clipBuf, clst, cl, cl.Opacity, cl.BlendMode, true)
		}
		blend.DestIn.Draw(clipBuf, buf.Rect, buf, pt)
		if blendMode == psd.BlendModePassThrough {
			blend.Copy.Draw(b, clipBuf.Rect, clipBuf, pt)
		} else {
			drawWithOpacity(b, clipBuf.Rect, clipBuf, pt, opacity, blendMode)
		}
		clipBuf.Pix, buf.Pix = buf.Pix, clipBuf.Pix
		return r.updateCache(drDrew, l, pt, buf, c, clst)
	}

	// this is minor code path.
	// it is only used when "Blend Clipped Layers as Group" is unchecked in Photoshop's Layer Style dialog.
	// TODO: implement
	drawWithOpacity(b, buf.Rect, buf, pt, opacity, blendMode)
	for _, cl := range l.Clip {
		r.drawLayer(pt, b, clst, cl, cl.Opacity, cl.BlendMode, false)
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
	return r.updateCache(drDrew, l, pt, buf, c, clst)
}
