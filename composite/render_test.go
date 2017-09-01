package composite

import (
	"context"
	"image"
	"os"
	"testing"
)

func load() *Tree {
	f, err := os.Open("psd/filled.psd")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	tree, err := New(context.Background(), f, nil)
	if err != nil {
		panic(err)
	}
	return tree
}

var tree = load()

func BenchmarkRender(b *testing.B) {
	ctx := context.Background()
	dest := image.NewRGBA(tree.CanvasRect)
	if err := tree.Renderer.Render(ctx, dest); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := tree.Renderer.Render(ctx, dest); err != nil {
			b.Fatal(err)
		}
	}
}
