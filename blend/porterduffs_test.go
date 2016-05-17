// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestPorterDuffFallbackClear(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", clear{}, false)
}

func TestPorterDuffNRGBAToNRGBAClear(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", clear{}, false)
}

func TestPorterDuffRGBAToNRGBAClear(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", clear{}, false)
}

func TestPorterDuffNRGBAToRGBAClear(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", clear{}, false)
}

func TestPorterDuffRGBAToRGBAClear(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", clear{}, false)
}

func BenchmarkPorterDuffFallbackClear(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", clear{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBAClear(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", clear{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBAClear(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", clear{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBAClear(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", clear{}, false)
}

func BenchmarkPorterDuffRGBAToRGBAClear(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", clear{}, false)
}

func TestPorterDuffFallbackCopy(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", copy{}, false)
}

func TestPorterDuffNRGBAToNRGBACopy(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", copy{}, false)
}

func TestPorterDuffRGBAToNRGBACopy(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", copy{}, false)
}

func TestPorterDuffNRGBAToRGBACopy(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", copy{}, false)
}

func TestPorterDuffRGBAToRGBACopy(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", copy{}, false)
}

func BenchmarkPorterDuffFallbackCopy(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", copy{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBACopy(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", copy{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBACopy(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", copy{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBACopy(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", copy{}, false)
}

func BenchmarkPorterDuffRGBAToRGBACopy(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", copy{}, false)
}

func TestPorterDuffFallbackDest(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", dest{}, false)
}

func TestPorterDuffNRGBAToNRGBADest(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", dest{}, false)
}

func TestPorterDuffRGBAToNRGBADest(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", dest{}, false)
}

func TestPorterDuffNRGBAToRGBADest(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", dest{}, false)
}

func TestPorterDuffRGBAToRGBADest(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", dest{}, false)
}

func BenchmarkPorterDuffFallbackDest(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", dest{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBADest(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", dest{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBADest(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", dest{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBADest(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", dest{}, false)
}

func BenchmarkPorterDuffRGBAToRGBADest(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", dest{}, false)
}

func TestPorterDuffFallbackSrcOver(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcOver{}, false)
}

func TestPorterDuffNRGBAToNRGBASrcOver(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOver{}, false)
}

func TestPorterDuffRGBAToNRGBASrcOver(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOver{}, false)
}

func TestPorterDuffNRGBAToRGBASrcOver(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcOver{}, false)
}

func TestPorterDuffRGBAToRGBASrcOver(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcOver{}, false)
}

func BenchmarkPorterDuffFallbackSrcOver(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcOver{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBASrcOver(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOver{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBASrcOver(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOver{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBASrcOver(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcOver{}, false)
}

func BenchmarkPorterDuffRGBAToRGBASrcOver(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcOver{}, false)
}

func TestPorterDuffFallbackDestOver(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destOver{}, false)
}

func TestPorterDuffNRGBAToNRGBADestOver(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destOver{}, false)
}

func TestPorterDuffRGBAToNRGBADestOver(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destOver{}, false)
}

func TestPorterDuffNRGBAToRGBADestOver(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destOver{}, false)
}

func TestPorterDuffRGBAToRGBADestOver(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destOver{}, false)
}

func BenchmarkPorterDuffFallbackDestOver(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destOver{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBADestOver(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destOver{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBADestOver(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destOver{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBADestOver(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destOver{}, false)
}

func BenchmarkPorterDuffRGBAToRGBADestOver(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destOver{}, false)
}

func TestPorterDuffFallbackSrcIn(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcIn{}, false)
}

func TestPorterDuffNRGBAToNRGBASrcIn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcIn{}, false)
}

func TestPorterDuffRGBAToNRGBASrcIn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcIn{}, false)
}

func TestPorterDuffNRGBAToRGBASrcIn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcIn{}, false)
}

func TestPorterDuffRGBAToRGBASrcIn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcIn{}, false)
}

func BenchmarkPorterDuffFallbackSrcIn(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcIn{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBASrcIn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcIn{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBASrcIn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcIn{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBASrcIn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcIn{}, false)
}

func BenchmarkPorterDuffRGBAToRGBASrcIn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcIn{}, false)
}

func TestPorterDuffFallbackDestIn(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destIn{}, false)
}

func TestPorterDuffNRGBAToNRGBADestIn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destIn{}, false)
}

func TestPorterDuffRGBAToNRGBADestIn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destIn{}, false)
}

func TestPorterDuffNRGBAToRGBADestIn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destIn{}, false)
}

func TestPorterDuffRGBAToRGBADestIn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destIn{}, false)
}

func BenchmarkPorterDuffFallbackDestIn(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destIn{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBADestIn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destIn{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBADestIn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destIn{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBADestIn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destIn{}, false)
}

func BenchmarkPorterDuffRGBAToRGBADestIn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destIn{}, false)
}

func TestPorterDuffFallbackSrcOut(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcOut{}, false)
}

func TestPorterDuffNRGBAToNRGBASrcOut(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOut{}, false)
}

func TestPorterDuffRGBAToNRGBASrcOut(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcOut{}, false)
}

func TestPorterDuffNRGBAToRGBASrcOut(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcOut{}, false)
}

func TestPorterDuffRGBAToRGBASrcOut(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcOut{}, false)
}

func BenchmarkPorterDuffFallbackSrcOut(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcOut{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBASrcOut(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOut{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBASrcOut(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcOut{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBASrcOut(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcOut{}, false)
}

func BenchmarkPorterDuffRGBAToRGBASrcOut(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcOut{}, false)
}

func TestPorterDuffFallbackDestOut(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destOut{}, false)
}

func TestPorterDuffNRGBAToNRGBADestOut(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destOut{}, false)
}

func TestPorterDuffRGBAToNRGBADestOut(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destOut{}, false)
}

func TestPorterDuffNRGBAToRGBADestOut(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destOut{}, false)
}

func TestPorterDuffRGBAToRGBADestOut(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destOut{}, false)
}

func BenchmarkPorterDuffFallbackDestOut(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destOut{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBADestOut(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destOut{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBADestOut(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destOut{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBADestOut(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destOut{}, false)
}

func BenchmarkPorterDuffRGBAToRGBADestOut(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destOut{}, false)
}

func TestPorterDuffFallbackSrcAtop(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", srcAtop{}, false)
}

func TestPorterDuffNRGBAToNRGBASrcAtop(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", srcAtop{}, false)
}

func TestPorterDuffRGBAToNRGBASrcAtop(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", srcAtop{}, false)
}

func TestPorterDuffNRGBAToRGBASrcAtop(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", srcAtop{}, false)
}

func TestPorterDuffRGBAToRGBASrcAtop(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", srcAtop{}, false)
}

func BenchmarkPorterDuffFallbackSrcAtop(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", srcAtop{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBASrcAtop(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", srcAtop{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBASrcAtop(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", srcAtop{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBASrcAtop(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", srcAtop{}, false)
}

func BenchmarkPorterDuffRGBAToRGBASrcAtop(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", srcAtop{}, false)
}

func TestPorterDuffFallbackDestAtop(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", destAtop{}, false)
}

func TestPorterDuffNRGBAToNRGBADestAtop(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", destAtop{}, false)
}

func TestPorterDuffRGBAToNRGBADestAtop(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", destAtop{}, false)
}

func TestPorterDuffNRGBAToRGBADestAtop(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", destAtop{}, false)
}

func TestPorterDuffRGBAToRGBADestAtop(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", destAtop{}, false)
}

func BenchmarkPorterDuffFallbackDestAtop(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", destAtop{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBADestAtop(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", destAtop{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBADestAtop(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", destAtop{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBADestAtop(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", destAtop{}, false)
}

func BenchmarkPorterDuffRGBAToRGBADestAtop(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", destAtop{}, false)
}

func TestPorterDuffFallbackXOR(t *testing.T) {
	testDrawFallback(t, "png/a.png", "png/b.png", xOR{}, false)
}

func TestPorterDuffNRGBAToNRGBAXOR(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", xOR{}, false)
}

func TestPorterDuffRGBAToNRGBAXOR(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", xOR{}, false)
}

func TestPorterDuffNRGBAToRGBAXOR(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", xOR{}, false)
}

func TestPorterDuffRGBAToRGBAXOR(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", xOR{}, false)
}

func BenchmarkPorterDuffFallbackXOR(b *testing.B) {
	benchmarkDrawFallback(b, "png/a.png", "png/b.png", xOR{}, false)
}

func BenchmarkPorterDuffNRGBAToNRGBAXOR(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", xOR{}, false)
}

func BenchmarkPorterDuffRGBAToNRGBAXOR(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", xOR{}, false)
}

func BenchmarkPorterDuffNRGBAToRGBAXOR(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", xOR{}, false)
}

func BenchmarkPorterDuffRGBAToRGBAXOR(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", xOR{}, false)
}
