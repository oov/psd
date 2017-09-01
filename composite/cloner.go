package composite

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

func (cl cloner) Clone(t *Tree) *Tree {
	nt := *t
	nt.Renderer = &Renderer{
		tree:  &nt,
		cache: map[int]*cache{},
	}
	nt.Renderer.pool.New = nt.Renderer.allocate
	cl.clone(&t.Root, &nt.Root)
	cl.registerClippingGroup(&t.Root, &nt.Root)
	return &nt
}
