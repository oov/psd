// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestBlendFallbackNormal(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestBlendFallbackNormalProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestBlendNRGBAToNRGBANormal(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestBlendNRGBAToNRGBANormalProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestBlendRGBAToNRGBANormal(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestBlendRGBAToNRGBANormalProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestBlendNRGBAToRGBANormal(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestBlendNRGBAToRGBANormalProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestBlendRGBAToRGBANormal(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, false)
}

func TestBlendRGBAToRGBANormalProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkBlendFallbackNormal(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkBlendFallbackNormalProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkBlendNRGBAToNRGBANormal(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkBlendNRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkBlendRGBAToNRGBANormal(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkBlendRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkBlendNRGBAToRGBANormal(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkBlendNRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func BenchmarkBlendRGBAToRGBANormal(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, false)
}

func BenchmarkBlendRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{}, true)
}

func TestBlendFallbackDarken(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestBlendFallbackDarkenProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestBlendNRGBAToNRGBADarken(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestBlendNRGBAToNRGBADarkenProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestBlendRGBAToNRGBADarken(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestBlendRGBAToNRGBADarkenProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestBlendNRGBAToRGBADarken(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestBlendNRGBAToRGBADarkenProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestBlendRGBAToRGBADarken(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, false)
}

func TestBlendRGBAToRGBADarkenProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkBlendFallbackDarken(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkBlendFallbackDarkenProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkBlendNRGBAToNRGBADarken(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkBlendNRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkBlendRGBAToNRGBADarken(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkBlendRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkBlendNRGBAToRGBADarken(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkBlendNRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func BenchmarkBlendRGBAToRGBADarken(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, false)
}

func BenchmarkBlendRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{}, true)
}

func TestBlendFallbackMultiply(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestBlendFallbackMultiplyProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestBlendNRGBAToNRGBAMultiply(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestBlendNRGBAToNRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestBlendRGBAToNRGBAMultiply(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestBlendRGBAToNRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestBlendNRGBAToRGBAMultiply(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestBlendNRGBAToRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestBlendRGBAToRGBAMultiply(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, false)
}

func TestBlendRGBAToRGBAMultiplyProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkBlendFallbackMultiply(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkBlendFallbackMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkBlendNRGBAToNRGBAMultiply(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkBlendNRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkBlendRGBAToNRGBAMultiply(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkBlendRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkBlendNRGBAToRGBAMultiply(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkBlendNRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func BenchmarkBlendRGBAToRGBAMultiply(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, false)
}

func BenchmarkBlendRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{}, true)
}

func TestBlendFallbackColorBurn(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestBlendFallbackColorBurnProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestBlendNRGBAToNRGBAColorBurn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestBlendNRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestBlendRGBAToNRGBAColorBurn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestBlendRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestBlendNRGBAToRGBAColorBurn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestBlendNRGBAToRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestBlendRGBAToRGBAColorBurn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func TestBlendRGBAToRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkBlendFallbackColorBurn(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkBlendFallbackColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkBlendNRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkBlendNRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkBlendRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkBlendRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkBlendNRGBAToRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkBlendNRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func BenchmarkBlendRGBAToRGBAColorBurn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, false)
}

func BenchmarkBlendRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{}, true)
}

func TestBlendFallbackLinearBurn(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestBlendFallbackLinearBurnProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestBlendNRGBAToNRGBALinearBurn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestBlendNRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestBlendRGBAToNRGBALinearBurn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestBlendRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestBlendNRGBAToRGBALinearBurn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestBlendNRGBAToRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestBlendRGBAToRGBALinearBurn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func TestBlendRGBAToRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkBlendFallbackLinearBurn(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkBlendFallbackLinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkBlendNRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkBlendNRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkBlendRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkBlendRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkBlendNRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkBlendNRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func BenchmarkBlendRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, false)
}

func BenchmarkBlendRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{}, true)
}

func TestBlendFallbackDarkerColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestBlendFallbackDarkerColorProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestBlendNRGBAToNRGBADarkerColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestBlendNRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestBlendRGBAToNRGBADarkerColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestBlendRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestBlendNRGBAToRGBADarkerColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestBlendNRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestBlendRGBAToRGBADarkerColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func TestBlendRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkBlendFallbackDarkerColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkBlendFallbackDarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkBlendNRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkBlendNRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkBlendRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkBlendRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkBlendNRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkBlendNRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func BenchmarkBlendRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, false)
}

func BenchmarkBlendRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{}, true)
}

func TestBlendFallbackLighten(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestBlendFallbackLightenProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestBlendNRGBAToNRGBALighten(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestBlendNRGBAToNRGBALightenProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestBlendRGBAToNRGBALighten(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestBlendRGBAToNRGBALightenProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestBlendNRGBAToRGBALighten(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestBlendNRGBAToRGBALightenProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestBlendRGBAToRGBALighten(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, false)
}

func TestBlendRGBAToRGBALightenProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkBlendFallbackLighten(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkBlendFallbackLightenProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkBlendNRGBAToNRGBALighten(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkBlendNRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkBlendRGBAToNRGBALighten(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkBlendRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkBlendNRGBAToRGBALighten(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkBlendNRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func BenchmarkBlendRGBAToRGBALighten(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, false)
}

func BenchmarkBlendRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{}, true)
}

func TestBlendFallbackScreen(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestBlendFallbackScreenProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestBlendNRGBAToNRGBAScreen(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestBlendNRGBAToNRGBAScreenProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestBlendRGBAToNRGBAScreen(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestBlendRGBAToNRGBAScreenProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestBlendNRGBAToRGBAScreen(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestBlendNRGBAToRGBAScreenProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestBlendRGBAToRGBAScreen(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, false)
}

func TestBlendRGBAToRGBAScreenProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkBlendFallbackScreen(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkBlendFallbackScreenProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkBlendNRGBAToNRGBAScreen(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkBlendNRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkBlendRGBAToNRGBAScreen(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkBlendRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkBlendNRGBAToRGBAScreen(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkBlendNRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func BenchmarkBlendRGBAToRGBAScreen(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, false)
}

func BenchmarkBlendRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{}, true)
}

func TestBlendFallbackColorDodge(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestBlendFallbackColorDodgeProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestBlendNRGBAToNRGBAColorDodge(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestBlendNRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestBlendRGBAToNRGBAColorDodge(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestBlendRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestBlendNRGBAToRGBAColorDodge(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestBlendNRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestBlendRGBAToRGBAColorDodge(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func TestBlendRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkBlendFallbackColorDodge(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkBlendFallbackColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkBlendNRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkBlendNRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkBlendRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkBlendRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkBlendNRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkBlendNRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func BenchmarkBlendRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, false)
}

func BenchmarkBlendRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{}, true)
}

func TestBlendFallbackLinearDodge(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestBlendFallbackLinearDodgeProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestBlendNRGBAToNRGBALinearDodge(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestBlendNRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestBlendRGBAToNRGBALinearDodge(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestBlendRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestBlendNRGBAToRGBALinearDodge(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestBlendNRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestBlendRGBAToRGBALinearDodge(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func TestBlendRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkBlendFallbackLinearDodge(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkBlendFallbackLinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkBlendNRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkBlendNRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkBlendRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkBlendRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkBlendNRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkBlendNRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func BenchmarkBlendRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, false)
}

func BenchmarkBlendRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{}, true)
}

func TestBlendFallbackLighterColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestBlendFallbackLighterColorProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestBlendNRGBAToNRGBALighterColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestBlendNRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestBlendRGBAToNRGBALighterColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestBlendRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestBlendNRGBAToRGBALighterColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestBlendNRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestBlendRGBAToRGBALighterColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func TestBlendRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkBlendFallbackLighterColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkBlendFallbackLighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkBlendNRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkBlendNRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkBlendRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkBlendRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkBlendNRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkBlendNRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func BenchmarkBlendRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, false)
}

func BenchmarkBlendRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{}, true)
}

func TestBlendFallbackAdd(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestBlendFallbackAddProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestBlendNRGBAToNRGBAAdd(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestBlendNRGBAToNRGBAAddProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestBlendRGBAToNRGBAAdd(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestBlendRGBAToNRGBAAddProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestBlendNRGBAToRGBAAdd(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestBlendNRGBAToRGBAAddProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func TestBlendRGBAToRGBAAdd(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, false)
}

func TestBlendRGBAToRGBAAddProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkBlendFallbackAdd(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkBlendFallbackAddProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkBlendNRGBAToNRGBAAdd(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkBlendNRGBAToNRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkBlendRGBAToNRGBAAdd(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkBlendRGBAToNRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkBlendNRGBAToRGBAAdd(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkBlendNRGBAToRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func BenchmarkBlendRGBAToRGBAAdd(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, false)
}

func BenchmarkBlendRGBAToRGBAAddProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{}, true)
}

func TestBlendFallbackOverlay(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestBlendFallbackOverlayProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestBlendNRGBAToNRGBAOverlay(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestBlendNRGBAToNRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestBlendRGBAToNRGBAOverlay(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestBlendRGBAToNRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestBlendNRGBAToRGBAOverlay(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestBlendNRGBAToRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestBlendRGBAToRGBAOverlay(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, false)
}

func TestBlendRGBAToRGBAOverlayProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkBlendFallbackOverlay(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkBlendFallbackOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkBlendNRGBAToNRGBAOverlay(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkBlendNRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkBlendRGBAToNRGBAOverlay(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkBlendRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkBlendNRGBAToRGBAOverlay(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkBlendNRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func BenchmarkBlendRGBAToRGBAOverlay(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, false)
}

func BenchmarkBlendRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{}, true)
}

func TestBlendFallbackSoftLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestBlendFallbackSoftLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestBlendNRGBAToNRGBASoftLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestBlendNRGBAToNRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestBlendRGBAToNRGBASoftLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestBlendRGBAToNRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestBlendNRGBAToRGBASoftLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestBlendNRGBAToRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestBlendRGBAToRGBASoftLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, false)
}

func TestBlendRGBAToRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkBlendFallbackSoftLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkBlendFallbackSoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkBlendNRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkBlendNRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkBlendRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkBlendRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkBlendNRGBAToRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkBlendNRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func BenchmarkBlendRGBAToRGBASoftLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, false)
}

func BenchmarkBlendRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{}, true)
}

func TestBlendFallbackHardLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestBlendFallbackHardLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestBlendNRGBAToNRGBAHardLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestBlendNRGBAToNRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestBlendRGBAToNRGBAHardLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestBlendRGBAToNRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestBlendNRGBAToRGBAHardLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestBlendNRGBAToRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestBlendRGBAToRGBAHardLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func TestBlendRGBAToRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkBlendFallbackHardLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkBlendFallbackHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkBlendNRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkBlendNRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkBlendRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkBlendRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkBlendNRGBAToRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkBlendNRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func BenchmarkBlendRGBAToRGBAHardLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, false)
}

func BenchmarkBlendRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{}, true)
}

func TestBlendFallbackLinearLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestBlendFallbackLinearLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestBlendNRGBAToNRGBALinearLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestBlendNRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestBlendRGBAToNRGBALinearLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestBlendRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestBlendNRGBAToRGBALinearLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestBlendNRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestBlendRGBAToRGBALinearLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func TestBlendRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkBlendFallbackLinearLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkBlendFallbackLinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkBlendNRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkBlendNRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkBlendRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkBlendRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkBlendNRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkBlendNRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func BenchmarkBlendRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, false)
}

func BenchmarkBlendRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{}, true)
}

func TestBlendFallbackVividLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestBlendFallbackVividLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestBlendNRGBAToNRGBAVividLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestBlendNRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestBlendRGBAToNRGBAVividLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestBlendRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestBlendNRGBAToRGBAVividLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestBlendNRGBAToRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestBlendRGBAToRGBAVividLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func TestBlendRGBAToRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkBlendFallbackVividLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkBlendFallbackVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkBlendNRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkBlendNRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkBlendRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkBlendRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkBlendNRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkBlendNRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func BenchmarkBlendRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, false)
}

func BenchmarkBlendRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{}, true)
}

func TestBlendFallbackPinLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestBlendFallbackPinLightProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestBlendNRGBAToNRGBAPinLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestBlendNRGBAToNRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestBlendRGBAToNRGBAPinLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestBlendRGBAToNRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestBlendNRGBAToRGBAPinLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestBlendNRGBAToRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestBlendRGBAToRGBAPinLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func TestBlendRGBAToRGBAPinLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkBlendFallbackPinLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkBlendFallbackPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkBlendNRGBAToNRGBAPinLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkBlendNRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkBlendRGBAToNRGBAPinLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkBlendRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkBlendNRGBAToRGBAPinLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkBlendNRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func BenchmarkBlendRGBAToRGBAPinLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, false)
}

func BenchmarkBlendRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{}, true)
}

func TestBlendFallbackHardMix(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestBlendFallbackHardMixProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestBlendNRGBAToNRGBAHardMix(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestBlendNRGBAToNRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestBlendRGBAToNRGBAHardMix(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestBlendRGBAToNRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestBlendNRGBAToRGBAHardMix(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestBlendNRGBAToRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestBlendRGBAToRGBAHardMix(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func TestBlendRGBAToRGBAHardMixProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkBlendFallbackHardMix(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkBlendFallbackHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkBlendNRGBAToNRGBAHardMix(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkBlendNRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkBlendRGBAToNRGBAHardMix(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkBlendRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkBlendNRGBAToRGBAHardMix(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkBlendNRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func BenchmarkBlendRGBAToRGBAHardMix(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, false)
}

func BenchmarkBlendRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{}, true)
}

func TestBlendFallbackDifference(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestBlendFallbackDifferenceProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestBlendNRGBAToNRGBADifference(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestBlendNRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestBlendRGBAToNRGBADifference(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestBlendRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestBlendNRGBAToRGBADifference(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestBlendNRGBAToRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestBlendRGBAToRGBADifference(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, false)
}

func TestBlendRGBAToRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkBlendFallbackDifference(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkBlendFallbackDifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkBlendNRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkBlendNRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkBlendRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkBlendRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkBlendNRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkBlendNRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func BenchmarkBlendRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, false)
}

func BenchmarkBlendRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{}, true)
}

func TestBlendFallbackExclusion(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestBlendFallbackExclusionProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestBlendNRGBAToNRGBAExclusion(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestBlendNRGBAToNRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestBlendRGBAToNRGBAExclusion(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestBlendRGBAToNRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestBlendNRGBAToRGBAExclusion(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestBlendNRGBAToRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestBlendRGBAToRGBAExclusion(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func TestBlendRGBAToRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkBlendFallbackExclusion(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkBlendFallbackExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkBlendNRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkBlendNRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkBlendRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkBlendRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkBlendNRGBAToRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkBlendNRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func BenchmarkBlendRGBAToRGBAExclusion(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, false)
}

func BenchmarkBlendRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{}, true)
}

func TestBlendFallbackSubtract(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestBlendFallbackSubtractProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestBlendNRGBAToNRGBASubtract(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestBlendNRGBAToNRGBASubtractProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestBlendRGBAToNRGBASubtract(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestBlendRGBAToNRGBASubtractProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestBlendNRGBAToRGBASubtract(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestBlendNRGBAToRGBASubtractProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestBlendRGBAToRGBASubtract(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, false)
}

func TestBlendRGBAToRGBASubtractProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkBlendFallbackSubtract(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkBlendFallbackSubtractProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkBlendNRGBAToNRGBASubtract(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkBlendNRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkBlendRGBAToNRGBASubtract(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkBlendRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkBlendNRGBAToRGBASubtract(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkBlendNRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func BenchmarkBlendRGBAToRGBASubtract(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, false)
}

func BenchmarkBlendRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{}, true)
}

func TestBlendFallbackDivide(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestBlendFallbackDivideProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestBlendNRGBAToNRGBADivide(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestBlendNRGBAToNRGBADivideProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestBlendRGBAToNRGBADivide(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestBlendRGBAToNRGBADivideProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestBlendNRGBAToRGBADivide(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestBlendNRGBAToRGBADivideProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestBlendRGBAToRGBADivide(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, false)
}

func TestBlendRGBAToRGBADivideProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkBlendFallbackDivide(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkBlendFallbackDivideProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkBlendNRGBAToNRGBADivide(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkBlendNRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkBlendRGBAToNRGBADivide(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkBlendRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkBlendNRGBAToRGBADivide(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkBlendNRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func BenchmarkBlendRGBAToRGBADivide(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, false)
}

func BenchmarkBlendRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{}, true)
}

func TestBlendFallbackHue(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestBlendFallbackHueProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestBlendNRGBAToNRGBAHue(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestBlendNRGBAToNRGBAHueProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestBlendRGBAToNRGBAHue(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestBlendRGBAToNRGBAHueProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestBlendNRGBAToRGBAHue(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestBlendNRGBAToRGBAHueProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestBlendRGBAToRGBAHue(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, false)
}

func TestBlendRGBAToRGBAHueProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkBlendFallbackHue(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkBlendFallbackHueProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkBlendNRGBAToNRGBAHue(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkBlendNRGBAToNRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkBlendRGBAToNRGBAHue(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkBlendRGBAToNRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkBlendNRGBAToRGBAHue(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkBlendNRGBAToRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func BenchmarkBlendRGBAToRGBAHue(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, false)
}

func BenchmarkBlendRGBAToRGBAHueProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{}, true)
}

func TestBlendFallbackSaturation(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestBlendFallbackSaturationProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestBlendNRGBAToNRGBASaturation(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestBlendNRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestBlendRGBAToNRGBASaturation(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestBlendRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestBlendNRGBAToRGBASaturation(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestBlendNRGBAToRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestBlendRGBAToRGBASaturation(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, false)
}

func TestBlendRGBAToRGBASaturationProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkBlendFallbackSaturation(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkBlendFallbackSaturationProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkBlendNRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkBlendNRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkBlendRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkBlendRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkBlendNRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkBlendNRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func BenchmarkBlendRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, false)
}

func BenchmarkBlendRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{}, true)
}

func TestBlendFallbackColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestBlendFallbackColorProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestBlendNRGBAToNRGBAColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestBlendNRGBAToNRGBAColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestBlendRGBAToNRGBAColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestBlendRGBAToNRGBAColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestBlendNRGBAToRGBAColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestBlendNRGBAToRGBAColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func TestBlendRGBAToRGBAColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, false)
}

func TestBlendRGBAToRGBAColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkBlendFallbackColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkBlendFallbackColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkBlendNRGBAToNRGBAColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkBlendNRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkBlendRGBAToNRGBAColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkBlendRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkBlendNRGBAToRGBAColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkBlendNRGBAToRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func BenchmarkBlendRGBAToRGBAColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, false)
}

func BenchmarkBlendRGBAToRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{}, true)
}

func TestBlendFallbackLuminosity(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestBlendFallbackLuminosityProtectAlpha(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestBlendNRGBAToNRGBALuminosity(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestBlendNRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestBlendRGBAToNRGBALuminosity(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestBlendRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestBlendNRGBAToRGBALuminosity(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestBlendNRGBAToRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func TestBlendRGBAToRGBALuminosity(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func TestBlendRGBAToRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkBlendFallbackLuminosity(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkBlendFallbackLuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkBlendNRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkBlendNRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkBlendRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkBlendRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkBlendNRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkBlendNRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}

func BenchmarkBlendRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, false)
}

func BenchmarkBlendRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{}, true)
}
