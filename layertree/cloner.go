package layertree

type cloner map[int]*Layer

func (cl cloner) clone(src *Layer, dest *Layer) {
	*dest = *src
	cl[dest.SeqID] = dest
	dest.Children = make([]Layer, len(src.Children))
	for i := range dest.Children {
		cl.clone(&src.Children[i], &dest.Children[i])
		dest.Children[i].Parent = dest
	}
}

func (cl cloner) registerClippingGroup(src *Layer, dest *Layer) {
	dest.Clip = make([]*Layer, len(src.Clip))
	for i := range dest.Clip {
		dest.Clip[i] = cl[src.Clip[i].SeqID]
	}
	if src.ClippedBy != nil {
		dest.ClippedBy = cl[src.ClippedBy.SeqID]
	}
	for i := range dest.Children {
		cl.registerClippingGroup(&src.Children[i], &dest.Children[i])
	}

}

func (cl cloner) Clone(r *Root) *Root {
	rr := *r
	rr.Renderer.layertree = &rr
	rr.Children = make([]Layer, len(r.Children))
	for i := range rr.Children {
		cl.clone(&r.Children[i], &rr.Children[i])
	}
	for i := range rr.Children {
		cl.registerClippingGroup(&r.Children[i], &rr.Children[i])
	}
	return &rr
}
