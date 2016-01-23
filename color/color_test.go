package color

import "testing"

func BenchmarkNCMYKAToRGBA(b *testing.B) {
	c := NCMYKA{C: 0x80, M: 0x80, Y: 0x80, K: 0x80, A: 0x80}
	for i := 0; i < b.N; i++ {
		c.RGBA()
	}
}
