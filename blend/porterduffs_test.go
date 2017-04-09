// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestPorterDuffFallbackClear(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffNRGBAToNRGBAClear(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffRGBAToNRGBAClear(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffAlphaToNRGBAClear(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffNRGBAToRGBAClear(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffRGBAToRGBAClear(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffAlphaToRGBAClear(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", clear{})
}

func BenchmarkPorterDuffFallbackClear(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", clear{})
}

func BenchmarkPorterDuffNRGBAToNRGBAClear(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", clear{})
}

func BenchmarkPorterDuffRGBAToNRGBAClear(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", clear{})
}

func BenchmarkPorterDuffNRGBAToRGBAClear(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", clear{})
}

func BenchmarkPorterDuffRGBAToRGBAClear(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", clear{})
}

func TestPorterDuffFallbackCopy(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffNRGBAToNRGBACopy(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffRGBAToNRGBACopy(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffAlphaToNRGBACopy(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffNRGBAToRGBACopy(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffRGBAToRGBACopy(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffAlphaToRGBACopy(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", copy{})
}

func BenchmarkPorterDuffFallbackCopy(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", copy{})
}

func BenchmarkPorterDuffNRGBAToNRGBACopy(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", copy{})
}

func BenchmarkPorterDuffRGBAToNRGBACopy(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", copy{})
}

func BenchmarkPorterDuffNRGBAToRGBACopy(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", copy{})
}

func BenchmarkPorterDuffRGBAToRGBACopy(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", copy{})
}

func TestPorterDuffFallbackDest(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffNRGBAToNRGBADest(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffRGBAToNRGBADest(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffAlphaToNRGBADest(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffNRGBAToRGBADest(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffRGBAToRGBADest(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffAlphaToRGBADest(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", dest{})
}

func BenchmarkPorterDuffFallbackDest(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", dest{})
}

func BenchmarkPorterDuffNRGBAToNRGBADest(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", dest{})
}

func BenchmarkPorterDuffRGBAToNRGBADest(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", dest{})
}

func BenchmarkPorterDuffNRGBAToRGBADest(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", dest{})
}

func BenchmarkPorterDuffRGBAToRGBADest(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", dest{})
}

func TestPorterDuffFallbackSrcOver(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffNRGBAToNRGBASrcOver(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffRGBAToNRGBASrcOver(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffAlphaToNRGBASrcOver(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffNRGBAToRGBASrcOver(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffRGBAToRGBASrcOver(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffAlphaToRGBASrcOver(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", srcOver{})
}

func BenchmarkPorterDuffFallbackSrcOver(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcOver{})
}

func BenchmarkPorterDuffNRGBAToNRGBASrcOver(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOver{})
}

func BenchmarkPorterDuffRGBAToNRGBASrcOver(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOver{})
}

func BenchmarkPorterDuffNRGBAToRGBASrcOver(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcOver{})
}

func BenchmarkPorterDuffRGBAToRGBASrcOver(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcOver{})
}

func TestPorterDuffFallbackDestOver(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffNRGBAToNRGBADestOver(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffRGBAToNRGBADestOver(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffAlphaToNRGBADestOver(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffNRGBAToRGBADestOver(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffRGBAToRGBADestOver(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffAlphaToRGBADestOver(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", destOver{})
}

func BenchmarkPorterDuffFallbackDestOver(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destOver{})
}

func BenchmarkPorterDuffNRGBAToNRGBADestOver(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destOver{})
}

func BenchmarkPorterDuffRGBAToNRGBADestOver(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destOver{})
}

func BenchmarkPorterDuffNRGBAToRGBADestOver(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destOver{})
}

func BenchmarkPorterDuffRGBAToRGBADestOver(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destOver{})
}

func TestPorterDuffFallbackSrcIn(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffNRGBAToNRGBASrcIn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffRGBAToNRGBASrcIn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffAlphaToNRGBASrcIn(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffNRGBAToRGBASrcIn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffRGBAToRGBASrcIn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffAlphaToRGBASrcIn(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", srcIn{})
}

func BenchmarkPorterDuffFallbackSrcIn(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcIn{})
}

func BenchmarkPorterDuffNRGBAToNRGBASrcIn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcIn{})
}

func BenchmarkPorterDuffRGBAToNRGBASrcIn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcIn{})
}

func BenchmarkPorterDuffNRGBAToRGBASrcIn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcIn{})
}

func BenchmarkPorterDuffRGBAToRGBASrcIn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcIn{})
}

func TestPorterDuffFallbackDestIn(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffNRGBAToNRGBADestIn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffRGBAToNRGBADestIn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffAlphaToNRGBADestIn(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffNRGBAToRGBADestIn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffRGBAToRGBADestIn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffAlphaToRGBADestIn(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", destIn{})
}

func BenchmarkPorterDuffFallbackDestIn(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destIn{})
}

func BenchmarkPorterDuffNRGBAToNRGBADestIn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destIn{})
}

func BenchmarkPorterDuffRGBAToNRGBADestIn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destIn{})
}

func BenchmarkPorterDuffNRGBAToRGBADestIn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destIn{})
}

func BenchmarkPorterDuffRGBAToRGBADestIn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destIn{})
}

func TestPorterDuffFallbackSrcOut(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffNRGBAToNRGBASrcOut(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffRGBAToNRGBASrcOut(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffAlphaToNRGBASrcOut(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffNRGBAToRGBASrcOut(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffRGBAToRGBASrcOut(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffAlphaToRGBASrcOut(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", srcOut{})
}

func BenchmarkPorterDuffFallbackSrcOut(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcOut{})
}

func BenchmarkPorterDuffNRGBAToNRGBASrcOut(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOut{})
}

func BenchmarkPorterDuffRGBAToNRGBASrcOut(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOut{})
}

func BenchmarkPorterDuffNRGBAToRGBASrcOut(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcOut{})
}

func BenchmarkPorterDuffRGBAToRGBASrcOut(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcOut{})
}

func TestPorterDuffFallbackDestOut(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffNRGBAToNRGBADestOut(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffRGBAToNRGBADestOut(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffAlphaToNRGBADestOut(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffNRGBAToRGBADestOut(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffRGBAToRGBADestOut(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffAlphaToRGBADestOut(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", destOut{})
}

func BenchmarkPorterDuffFallbackDestOut(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destOut{})
}

func BenchmarkPorterDuffNRGBAToNRGBADestOut(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destOut{})
}

func BenchmarkPorterDuffRGBAToNRGBADestOut(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destOut{})
}

func BenchmarkPorterDuffNRGBAToRGBADestOut(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destOut{})
}

func BenchmarkPorterDuffRGBAToRGBADestOut(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destOut{})
}

func TestPorterDuffFallbackSrcAtop(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffNRGBAToNRGBASrcAtop(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffRGBAToNRGBASrcAtop(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffAlphaToNRGBASrcAtop(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffNRGBAToRGBASrcAtop(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffRGBAToRGBASrcAtop(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffAlphaToRGBASrcAtop(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", srcAtop{})
}

func BenchmarkPorterDuffFallbackSrcAtop(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcAtop{})
}

func BenchmarkPorterDuffNRGBAToNRGBASrcAtop(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcAtop{})
}

func BenchmarkPorterDuffRGBAToNRGBASrcAtop(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcAtop{})
}

func BenchmarkPorterDuffNRGBAToRGBASrcAtop(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcAtop{})
}

func BenchmarkPorterDuffRGBAToRGBASrcAtop(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcAtop{})
}

func TestPorterDuffFallbackDestAtop(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffNRGBAToNRGBADestAtop(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffRGBAToNRGBADestAtop(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffAlphaToNRGBADestAtop(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffNRGBAToRGBADestAtop(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffRGBAToRGBADestAtop(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffAlphaToRGBADestAtop(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", destAtop{})
}

func BenchmarkPorterDuffFallbackDestAtop(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destAtop{})
}

func BenchmarkPorterDuffNRGBAToNRGBADestAtop(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destAtop{})
}

func BenchmarkPorterDuffRGBAToNRGBADestAtop(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destAtop{})
}

func BenchmarkPorterDuffNRGBAToRGBADestAtop(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destAtop{})
}

func BenchmarkPorterDuffRGBAToRGBADestAtop(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destAtop{})
}

func TestPorterDuffFallbackXOR(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", xOR{})
}

func TestPorterDuffNRGBAToNRGBAXOR(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", xOR{})
}

func TestPorterDuffRGBAToNRGBAXOR(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", xOR{})
}

func TestPorterDuffAlphaToNRGBAXOR(t *testing.T) {
	testDrawAlphaToNRGBA(t, "png/a.png", "png/b.png", xOR{})
}

func TestPorterDuffNRGBAToRGBAXOR(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", xOR{})
}

func TestPorterDuffRGBAToRGBAXOR(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", xOR{})
}

func TestPorterDuffAlphaToRGBAXOR(t *testing.T) {
	testDrawAlphaToRGBA(t, "png/a.png", "png/b.png", xOR{})
}

func BenchmarkPorterDuffFallbackXOR(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", xOR{})
}

func BenchmarkPorterDuffNRGBAToNRGBAXOR(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", xOR{})
}

func BenchmarkPorterDuffRGBAToNRGBAXOR(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", xOR{})
}

func BenchmarkPorterDuffNRGBAToRGBAXOR(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", xOR{})
}

func BenchmarkPorterDuffRGBAToRGBAXOR(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", xOR{})
}
