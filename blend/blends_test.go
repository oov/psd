// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestDrawFallbackNormal(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestDrawFallbackNormalProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestDrawNRGBAToNRGBANormal(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestDrawNRGBAToNRGBANormalProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestDrawRGBAToNRGBANormal(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestDrawRGBAToNRGBANormalProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestDrawNRGBAToRGBANormal(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestDrawNRGBAToRGBANormalProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestDrawRGBAToRGBANormal(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestDrawRGBAToRGBANormalProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkDrawFallbackNormal(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkDrawFallbackNormalProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkDrawNRGBAToNRGBANormal(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkDrawNRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkDrawRGBAToNRGBANormal(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkDrawRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkDrawNRGBAToRGBANormal(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkDrawNRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkDrawRGBAToRGBANormal(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkDrawRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestDrawFallbackDarken(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestDrawFallbackDarkenProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestDrawNRGBAToNRGBADarken(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestDrawNRGBAToNRGBADarkenProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestDrawRGBAToNRGBADarken(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestDrawRGBAToNRGBADarkenProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestDrawNRGBAToRGBADarken(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestDrawNRGBAToRGBADarkenProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestDrawRGBAToRGBADarken(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestDrawRGBAToRGBADarkenProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkDrawFallbackDarken(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkDrawFallbackDarkenProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkDrawNRGBAToNRGBADarken(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkDrawNRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkDrawRGBAToNRGBADarken(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkDrawRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkDrawNRGBAToRGBADarken(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkDrawNRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkDrawRGBAToRGBADarken(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkDrawRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestDrawFallbackMultiply(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestDrawFallbackMultiplyProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestDrawNRGBAToNRGBAMultiply(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestDrawNRGBAToNRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestDrawRGBAToNRGBAMultiply(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestDrawRGBAToNRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestDrawNRGBAToRGBAMultiply(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestDrawNRGBAToRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestDrawRGBAToRGBAMultiply(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestDrawRGBAToRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkDrawFallbackMultiply(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkDrawFallbackMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkDrawNRGBAToNRGBAMultiply(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkDrawNRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkDrawRGBAToNRGBAMultiply(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkDrawRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkDrawNRGBAToRGBAMultiply(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkDrawNRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkDrawRGBAToRGBAMultiply(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkDrawRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestDrawFallbackColorBurn(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestDrawFallbackColorBurnProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestDrawNRGBAToNRGBAColorBurn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestDrawNRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestDrawRGBAToNRGBAColorBurn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestDrawRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestDrawNRGBAToRGBAColorBurn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestDrawNRGBAToRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestDrawRGBAToRGBAColorBurn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestDrawRGBAToRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkDrawFallbackColorBurn(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkDrawFallbackColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkDrawNRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkDrawNRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkDrawRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkDrawRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkDrawNRGBAToRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkDrawNRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkDrawRGBAToRGBAColorBurn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkDrawRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestDrawFallbackLinearBurn(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestDrawFallbackLinearBurnProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestDrawNRGBAToNRGBALinearBurn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestDrawNRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestDrawRGBAToNRGBALinearBurn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestDrawRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestDrawNRGBAToRGBALinearBurn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestDrawNRGBAToRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestDrawRGBAToRGBALinearBurn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestDrawRGBAToRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkDrawFallbackLinearBurn(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkDrawFallbackLinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkDrawNRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkDrawNRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkDrawRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkDrawRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkDrawNRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkDrawNRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkDrawRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkDrawRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestDrawFallbackDarkerColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestDrawFallbackDarkerColorProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestDrawNRGBAToNRGBADarkerColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestDrawNRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestDrawRGBAToNRGBADarkerColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestDrawRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestDrawNRGBAToRGBADarkerColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestDrawNRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestDrawRGBAToRGBADarkerColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestDrawRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkDrawFallbackDarkerColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkDrawFallbackDarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkDrawNRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkDrawNRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkDrawRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkDrawRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkDrawNRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkDrawNRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkDrawRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkDrawRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestDrawFallbackLighten(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestDrawFallbackLightenProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestDrawNRGBAToNRGBALighten(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestDrawNRGBAToNRGBALightenProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestDrawRGBAToNRGBALighten(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestDrawRGBAToNRGBALightenProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestDrawNRGBAToRGBALighten(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestDrawNRGBAToRGBALightenProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestDrawRGBAToRGBALighten(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestDrawRGBAToRGBALightenProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkDrawFallbackLighten(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkDrawFallbackLightenProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkDrawNRGBAToNRGBALighten(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkDrawNRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkDrawRGBAToNRGBALighten(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkDrawRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkDrawNRGBAToRGBALighten(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkDrawNRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkDrawRGBAToRGBALighten(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkDrawRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestDrawFallbackScreen(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestDrawFallbackScreenProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestDrawNRGBAToNRGBAScreen(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestDrawNRGBAToNRGBAScreenProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestDrawRGBAToNRGBAScreen(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestDrawRGBAToNRGBAScreenProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestDrawNRGBAToRGBAScreen(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestDrawNRGBAToRGBAScreenProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestDrawRGBAToRGBAScreen(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestDrawRGBAToRGBAScreenProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkDrawFallbackScreen(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkDrawFallbackScreenProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkDrawNRGBAToNRGBAScreen(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkDrawNRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkDrawRGBAToNRGBAScreen(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkDrawRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkDrawNRGBAToRGBAScreen(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkDrawNRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkDrawRGBAToRGBAScreen(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkDrawRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestDrawFallbackColorDodge(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestDrawFallbackColorDodgeProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestDrawNRGBAToNRGBAColorDodge(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestDrawNRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestDrawRGBAToNRGBAColorDodge(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestDrawRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestDrawNRGBAToRGBAColorDodge(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestDrawNRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestDrawRGBAToRGBAColorDodge(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestDrawRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkDrawFallbackColorDodge(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkDrawFallbackColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkDrawNRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkDrawNRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkDrawRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkDrawRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkDrawNRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkDrawNRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkDrawRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkDrawRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestDrawFallbackLinearDodge(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestDrawFallbackLinearDodgeProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestDrawNRGBAToNRGBALinearDodge(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestDrawNRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestDrawRGBAToNRGBALinearDodge(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestDrawRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestDrawNRGBAToRGBALinearDodge(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestDrawNRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestDrawRGBAToRGBALinearDodge(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestDrawRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkDrawFallbackLinearDodge(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkDrawFallbackLinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkDrawNRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkDrawNRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkDrawRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkDrawRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkDrawNRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkDrawNRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkDrawRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkDrawRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestDrawFallbackLighterColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestDrawFallbackLighterColorProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestDrawNRGBAToNRGBALighterColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestDrawNRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestDrawRGBAToNRGBALighterColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestDrawRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestDrawNRGBAToRGBALighterColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestDrawNRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestDrawRGBAToRGBALighterColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestDrawRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkDrawFallbackLighterColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkDrawFallbackLighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkDrawNRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkDrawNRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkDrawRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkDrawRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkDrawNRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkDrawNRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkDrawRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkDrawRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestDrawFallbackAdd(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestDrawFallbackAddProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestDrawNRGBAToNRGBAAdd(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestDrawNRGBAToNRGBAAddProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestDrawRGBAToNRGBAAdd(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestDrawRGBAToNRGBAAddProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestDrawNRGBAToRGBAAdd(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestDrawNRGBAToRGBAAddProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestDrawRGBAToRGBAAdd(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestDrawRGBAToRGBAAddProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkDrawFallbackAdd(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkDrawFallbackAddProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkDrawNRGBAToNRGBAAdd(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkDrawNRGBAToNRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkDrawRGBAToNRGBAAdd(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkDrawRGBAToNRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkDrawNRGBAToRGBAAdd(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkDrawNRGBAToRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkDrawRGBAToRGBAAdd(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkDrawRGBAToRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func TestDrawFallbackOverlay(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestDrawFallbackOverlayProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestDrawNRGBAToNRGBAOverlay(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestDrawNRGBAToNRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestDrawRGBAToNRGBAOverlay(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestDrawRGBAToNRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestDrawNRGBAToRGBAOverlay(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestDrawNRGBAToRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestDrawRGBAToRGBAOverlay(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestDrawRGBAToRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkDrawFallbackOverlay(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkDrawFallbackOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkDrawNRGBAToNRGBAOverlay(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkDrawNRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkDrawRGBAToNRGBAOverlay(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkDrawRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkDrawNRGBAToRGBAOverlay(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkDrawNRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkDrawRGBAToRGBAOverlay(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkDrawRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestDrawFallbackSoftLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestDrawFallbackSoftLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestDrawNRGBAToNRGBASoftLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestDrawNRGBAToNRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestDrawRGBAToNRGBASoftLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestDrawRGBAToNRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestDrawNRGBAToRGBASoftLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestDrawNRGBAToRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestDrawRGBAToRGBASoftLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestDrawRGBAToRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkDrawFallbackSoftLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkDrawFallbackSoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkDrawNRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkDrawNRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkDrawRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkDrawRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkDrawNRGBAToRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkDrawNRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkDrawRGBAToRGBASoftLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkDrawRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestDrawFallbackHardLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestDrawFallbackHardLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestDrawNRGBAToNRGBAHardLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestDrawNRGBAToNRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestDrawRGBAToNRGBAHardLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestDrawRGBAToNRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestDrawNRGBAToRGBAHardLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestDrawNRGBAToRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestDrawRGBAToRGBAHardLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestDrawRGBAToRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkDrawFallbackHardLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkDrawFallbackHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkDrawNRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkDrawNRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkDrawRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkDrawRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkDrawNRGBAToRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkDrawNRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkDrawRGBAToRGBAHardLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkDrawRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestDrawFallbackLinearLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestDrawFallbackLinearLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestDrawNRGBAToNRGBALinearLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestDrawNRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestDrawRGBAToNRGBALinearLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestDrawRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestDrawNRGBAToRGBALinearLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestDrawNRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestDrawRGBAToRGBALinearLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestDrawRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkDrawFallbackLinearLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkDrawFallbackLinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkDrawNRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkDrawNRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkDrawRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkDrawRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkDrawNRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkDrawNRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkDrawRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkDrawRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestDrawFallbackVividLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestDrawFallbackVividLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestDrawNRGBAToNRGBAVividLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestDrawNRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestDrawRGBAToNRGBAVividLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestDrawRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestDrawNRGBAToRGBAVividLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestDrawNRGBAToRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestDrawRGBAToRGBAVividLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestDrawRGBAToRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkDrawFallbackVividLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkDrawFallbackVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkDrawNRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkDrawNRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkDrawRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkDrawRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkDrawNRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkDrawNRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkDrawRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkDrawRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestDrawFallbackPinLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestDrawFallbackPinLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestDrawNRGBAToNRGBAPinLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestDrawNRGBAToNRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestDrawRGBAToNRGBAPinLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestDrawRGBAToNRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestDrawNRGBAToRGBAPinLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestDrawNRGBAToRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestDrawRGBAToRGBAPinLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestDrawRGBAToRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkDrawFallbackPinLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkDrawFallbackPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkDrawNRGBAToNRGBAPinLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkDrawNRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkDrawRGBAToNRGBAPinLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkDrawRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkDrawNRGBAToRGBAPinLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkDrawNRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkDrawRGBAToRGBAPinLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkDrawRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestDrawFallbackHardMix(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestDrawFallbackHardMixProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestDrawNRGBAToNRGBAHardMix(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestDrawNRGBAToNRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestDrawRGBAToNRGBAHardMix(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestDrawRGBAToNRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestDrawNRGBAToRGBAHardMix(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestDrawNRGBAToRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestDrawRGBAToRGBAHardMix(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestDrawRGBAToRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkDrawFallbackHardMix(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkDrawFallbackHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkDrawNRGBAToNRGBAHardMix(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkDrawNRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkDrawRGBAToNRGBAHardMix(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkDrawRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkDrawNRGBAToRGBAHardMix(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkDrawNRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkDrawRGBAToRGBAHardMix(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkDrawRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestDrawFallbackDifference(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestDrawFallbackDifferenceProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestDrawNRGBAToNRGBADifference(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestDrawNRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestDrawRGBAToNRGBADifference(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestDrawRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestDrawNRGBAToRGBADifference(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestDrawNRGBAToRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestDrawRGBAToRGBADifference(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestDrawRGBAToRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkDrawFallbackDifference(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkDrawFallbackDifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkDrawNRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkDrawNRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkDrawRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkDrawRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkDrawNRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkDrawNRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkDrawRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkDrawRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestDrawFallbackExclusion(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestDrawFallbackExclusionProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestDrawNRGBAToNRGBAExclusion(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestDrawNRGBAToNRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestDrawRGBAToNRGBAExclusion(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestDrawRGBAToNRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestDrawNRGBAToRGBAExclusion(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestDrawNRGBAToRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestDrawRGBAToRGBAExclusion(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestDrawRGBAToRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkDrawFallbackExclusion(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkDrawFallbackExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkDrawNRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkDrawNRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkDrawRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkDrawRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkDrawNRGBAToRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkDrawNRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkDrawRGBAToRGBAExclusion(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkDrawRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestDrawFallbackSubtract(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestDrawFallbackSubtractProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestDrawNRGBAToNRGBASubtract(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestDrawNRGBAToNRGBASubtractProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestDrawRGBAToNRGBASubtract(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestDrawRGBAToNRGBASubtractProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestDrawNRGBAToRGBASubtract(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestDrawNRGBAToRGBASubtractProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestDrawRGBAToRGBASubtract(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestDrawRGBAToRGBASubtractProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkDrawFallbackSubtract(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkDrawFallbackSubtractProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkDrawNRGBAToNRGBASubtract(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkDrawNRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkDrawRGBAToNRGBASubtract(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkDrawRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkDrawNRGBAToRGBASubtract(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkDrawNRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkDrawRGBAToRGBASubtract(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkDrawRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestDrawFallbackDivide(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestDrawFallbackDivideProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestDrawNRGBAToNRGBADivide(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestDrawNRGBAToNRGBADivideProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestDrawRGBAToNRGBADivide(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestDrawRGBAToNRGBADivideProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestDrawNRGBAToRGBADivide(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestDrawNRGBAToRGBADivideProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestDrawRGBAToRGBADivide(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestDrawRGBAToRGBADivideProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkDrawFallbackDivide(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkDrawFallbackDivideProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkDrawNRGBAToNRGBADivide(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkDrawNRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkDrawRGBAToNRGBADivide(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkDrawRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkDrawNRGBAToRGBADivide(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkDrawNRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkDrawRGBAToRGBADivide(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkDrawRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestDrawFallbackHue(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestDrawFallbackHueProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestDrawNRGBAToNRGBAHue(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestDrawNRGBAToNRGBAHueProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestDrawRGBAToNRGBAHue(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestDrawRGBAToNRGBAHueProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestDrawNRGBAToRGBAHue(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestDrawNRGBAToRGBAHueProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestDrawRGBAToRGBAHue(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestDrawRGBAToRGBAHueProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkDrawFallbackHue(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkDrawFallbackHueProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkDrawNRGBAToNRGBAHue(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkDrawNRGBAToNRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkDrawRGBAToNRGBAHue(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkDrawRGBAToNRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkDrawNRGBAToRGBAHue(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkDrawNRGBAToRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkDrawRGBAToRGBAHue(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkDrawRGBAToRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestDrawFallbackSaturation(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestDrawFallbackSaturationProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestDrawNRGBAToNRGBASaturation(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestDrawNRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestDrawRGBAToNRGBASaturation(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestDrawRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestDrawNRGBAToRGBASaturation(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestDrawNRGBAToRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestDrawRGBAToRGBASaturation(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestDrawRGBAToRGBASaturationProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkDrawFallbackSaturation(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkDrawFallbackSaturationProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkDrawNRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkDrawNRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkDrawRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkDrawRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkDrawNRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkDrawNRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkDrawRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkDrawRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestDrawFallbackColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestDrawFallbackColorProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestDrawNRGBAToNRGBAColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestDrawNRGBAToNRGBAColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestDrawRGBAToNRGBAColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestDrawRGBAToNRGBAColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestDrawNRGBAToRGBAColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestDrawNRGBAToRGBAColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestDrawRGBAToRGBAColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestDrawRGBAToRGBAColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkDrawFallbackColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkDrawFallbackColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkDrawNRGBAToNRGBAColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkDrawNRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkDrawRGBAToNRGBAColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkDrawRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkDrawNRGBAToRGBAColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkDrawNRGBAToRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkDrawRGBAToRGBAColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkDrawRGBAToRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func TestDrawFallbackLuminosity(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestDrawFallbackLuminosityProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestDrawNRGBAToNRGBALuminosity(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestDrawNRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestDrawRGBAToNRGBALuminosity(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestDrawRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestDrawNRGBAToRGBALuminosity(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestDrawNRGBAToRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestDrawRGBAToRGBALuminosity(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestDrawRGBAToRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkDrawFallbackLuminosity(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkDrawFallbackLuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkDrawNRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkDrawNRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkDrawRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkDrawRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkDrawNRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkDrawNRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkDrawRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkDrawRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}
