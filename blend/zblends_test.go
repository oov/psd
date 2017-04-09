// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestBlendFallbackNormal(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", normal{})
}

func TestBlendNRGBAToNRGBANormal(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{})
}

func TestBlendRGBAToNRGBANormal(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", normal{})
}

func TestBlendNRGBAToRGBANormal(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{})
}

func TestBlendRGBAToRGBANormal(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", normal{})
}

func BenchmarkBlendFallbackNormal(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", normal{})
}

func BenchmarkBlendNRGBAToNRGBANormal(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{})
}

func BenchmarkBlendRGBAToNRGBANormal(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", normal{})
}

func BenchmarkBlendNRGBAToRGBANormal(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{})
}

func BenchmarkBlendRGBAToRGBANormal(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", normal{})
}

func TestBlendFallbackDarken(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darken{})
}

func TestBlendNRGBAToNRGBADarken(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{})
}

func TestBlendRGBAToNRGBADarken(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darken{})
}

func TestBlendNRGBAToRGBADarken(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{})
}

func TestBlendRGBAToRGBADarken(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darken{})
}

func BenchmarkBlendFallbackDarken(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darken{})
}

func BenchmarkBlendNRGBAToNRGBADarken(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{})
}

func BenchmarkBlendRGBAToNRGBADarken(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darken{})
}

func BenchmarkBlendNRGBAToRGBADarken(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{})
}

func BenchmarkBlendRGBAToRGBADarken(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darken{})
}

func TestBlendFallbackMultiply(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", multiply{})
}

func TestBlendNRGBAToNRGBAMultiply(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{})
}

func TestBlendRGBAToNRGBAMultiply(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", multiply{})
}

func TestBlendNRGBAToRGBAMultiply(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{})
}

func TestBlendRGBAToRGBAMultiply(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", multiply{})
}

func BenchmarkBlendFallbackMultiply(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", multiply{})
}

func BenchmarkBlendNRGBAToNRGBAMultiply(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{})
}

func BenchmarkBlendRGBAToNRGBAMultiply(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", multiply{})
}

func BenchmarkBlendNRGBAToRGBAMultiply(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{})
}

func BenchmarkBlendRGBAToRGBAMultiply(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", multiply{})
}

func TestBlendFallbackColorBurn(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorBurn{})
}

func TestBlendNRGBAToNRGBAColorBurn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{})
}

func TestBlendRGBAToNRGBAColorBurn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorBurn{})
}

func TestBlendNRGBAToRGBAColorBurn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{})
}

func TestBlendRGBAToRGBAColorBurn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorBurn{})
}

func BenchmarkBlendFallbackColorBurn(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorBurn{})
}

func BenchmarkBlendNRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{})
}

func BenchmarkBlendRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorBurn{})
}

func BenchmarkBlendNRGBAToRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{})
}

func BenchmarkBlendRGBAToRGBAColorBurn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorBurn{})
}

func TestBlendFallbackLinearBurn(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearBurn{})
}

func TestBlendNRGBAToNRGBALinearBurn(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{})
}

func TestBlendRGBAToNRGBALinearBurn(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearBurn{})
}

func TestBlendNRGBAToRGBALinearBurn(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{})
}

func TestBlendRGBAToRGBALinearBurn(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearBurn{})
}

func BenchmarkBlendFallbackLinearBurn(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearBurn{})
}

func BenchmarkBlendNRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{})
}

func BenchmarkBlendRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearBurn{})
}

func BenchmarkBlendNRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{})
}

func BenchmarkBlendRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearBurn{})
}

func TestBlendFallbackDarkerColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", darkerColor{})
}

func TestBlendNRGBAToNRGBADarkerColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{})
}

func TestBlendRGBAToNRGBADarkerColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", darkerColor{})
}

func TestBlendNRGBAToRGBADarkerColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{})
}

func TestBlendRGBAToRGBADarkerColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", darkerColor{})
}

func BenchmarkBlendFallbackDarkerColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", darkerColor{})
}

func BenchmarkBlendNRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{})
}

func BenchmarkBlendRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", darkerColor{})
}

func BenchmarkBlendNRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{})
}

func BenchmarkBlendRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", darkerColor{})
}

func TestBlendFallbackLighten(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighten{})
}

func TestBlendNRGBAToNRGBALighten(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{})
}

func TestBlendRGBAToNRGBALighten(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighten{})
}

func TestBlendNRGBAToRGBALighten(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{})
}

func TestBlendRGBAToRGBALighten(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighten{})
}

func BenchmarkBlendFallbackLighten(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighten{})
}

func BenchmarkBlendNRGBAToNRGBALighten(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{})
}

func BenchmarkBlendRGBAToNRGBALighten(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighten{})
}

func BenchmarkBlendNRGBAToRGBALighten(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{})
}

func BenchmarkBlendRGBAToRGBALighten(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighten{})
}

func TestBlendFallbackScreen(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", screen{})
}

func TestBlendNRGBAToNRGBAScreen(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{})
}

func TestBlendRGBAToNRGBAScreen(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", screen{})
}

func TestBlendNRGBAToRGBAScreen(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{})
}

func TestBlendRGBAToRGBAScreen(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", screen{})
}

func BenchmarkBlendFallbackScreen(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", screen{})
}

func BenchmarkBlendNRGBAToNRGBAScreen(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{})
}

func BenchmarkBlendRGBAToNRGBAScreen(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", screen{})
}

func BenchmarkBlendNRGBAToRGBAScreen(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{})
}

func BenchmarkBlendRGBAToRGBAScreen(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", screen{})
}

func TestBlendFallbackColorDodge(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", colorDodge{})
}

func TestBlendNRGBAToNRGBAColorDodge(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{})
}

func TestBlendRGBAToNRGBAColorDodge(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", colorDodge{})
}

func TestBlendNRGBAToRGBAColorDodge(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{})
}

func TestBlendRGBAToRGBAColorDodge(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", colorDodge{})
}

func BenchmarkBlendFallbackColorDodge(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", colorDodge{})
}

func BenchmarkBlendNRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{})
}

func BenchmarkBlendRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", colorDodge{})
}

func BenchmarkBlendNRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{})
}

func BenchmarkBlendRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", colorDodge{})
}

func TestBlendFallbackLinearDodge(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearDodge{})
}

func TestBlendNRGBAToNRGBALinearDodge(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{})
}

func TestBlendRGBAToNRGBALinearDodge(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearDodge{})
}

func TestBlendNRGBAToRGBALinearDodge(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{})
}

func TestBlendRGBAToRGBALinearDodge(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearDodge{})
}

func BenchmarkBlendFallbackLinearDodge(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearDodge{})
}

func BenchmarkBlendNRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{})
}

func BenchmarkBlendRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearDodge{})
}

func BenchmarkBlendNRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{})
}

func BenchmarkBlendRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearDodge{})
}

func TestBlendFallbackLighterColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", lighterColor{})
}

func TestBlendNRGBAToNRGBALighterColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{})
}

func TestBlendRGBAToNRGBALighterColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", lighterColor{})
}

func TestBlendNRGBAToRGBALighterColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{})
}

func TestBlendRGBAToRGBALighterColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", lighterColor{})
}

func BenchmarkBlendFallbackLighterColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", lighterColor{})
}

func BenchmarkBlendNRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{})
}

func BenchmarkBlendRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", lighterColor{})
}

func BenchmarkBlendNRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{})
}

func BenchmarkBlendRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", lighterColor{})
}

func TestBlendFallbackAdd(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", add{})
}

func TestBlendNRGBAToNRGBAAdd(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{})
}

func TestBlendRGBAToNRGBAAdd(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", add{})
}

func TestBlendNRGBAToRGBAAdd(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{})
}

func TestBlendRGBAToRGBAAdd(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", add{})
}

func BenchmarkBlendFallbackAdd(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", add{})
}

func BenchmarkBlendNRGBAToNRGBAAdd(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{})
}

func BenchmarkBlendRGBAToNRGBAAdd(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", add{})
}

func BenchmarkBlendNRGBAToRGBAAdd(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{})
}

func BenchmarkBlendRGBAToRGBAAdd(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", add{})
}

func TestBlendFallbackOverlay(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", overlay{})
}

func TestBlendNRGBAToNRGBAOverlay(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{})
}

func TestBlendRGBAToNRGBAOverlay(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", overlay{})
}

func TestBlendNRGBAToRGBAOverlay(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{})
}

func TestBlendRGBAToRGBAOverlay(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", overlay{})
}

func BenchmarkBlendFallbackOverlay(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", overlay{})
}

func BenchmarkBlendNRGBAToNRGBAOverlay(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{})
}

func BenchmarkBlendRGBAToNRGBAOverlay(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", overlay{})
}

func BenchmarkBlendNRGBAToRGBAOverlay(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{})
}

func BenchmarkBlendRGBAToRGBAOverlay(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", overlay{})
}

func TestBlendFallbackSoftLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", softLight{})
}

func TestBlendNRGBAToNRGBASoftLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{})
}

func TestBlendRGBAToNRGBASoftLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", softLight{})
}

func TestBlendNRGBAToRGBASoftLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{})
}

func TestBlendRGBAToRGBASoftLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", softLight{})
}

func BenchmarkBlendFallbackSoftLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", softLight{})
}

func BenchmarkBlendNRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{})
}

func BenchmarkBlendRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", softLight{})
}

func BenchmarkBlendNRGBAToRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{})
}

func BenchmarkBlendRGBAToRGBASoftLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", softLight{})
}

func TestBlendFallbackHardLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardLight{})
}

func TestBlendNRGBAToNRGBAHardLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{})
}

func TestBlendRGBAToNRGBAHardLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardLight{})
}

func TestBlendNRGBAToRGBAHardLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{})
}

func TestBlendRGBAToRGBAHardLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardLight{})
}

func BenchmarkBlendFallbackHardLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardLight{})
}

func BenchmarkBlendNRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{})
}

func BenchmarkBlendRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardLight{})
}

func BenchmarkBlendNRGBAToRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{})
}

func BenchmarkBlendRGBAToRGBAHardLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardLight{})
}

func TestBlendFallbackLinearLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", linearLight{})
}

func TestBlendNRGBAToNRGBALinearLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{})
}

func TestBlendRGBAToNRGBALinearLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", linearLight{})
}

func TestBlendNRGBAToRGBALinearLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{})
}

func TestBlendRGBAToRGBALinearLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", linearLight{})
}

func BenchmarkBlendFallbackLinearLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", linearLight{})
}

func BenchmarkBlendNRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{})
}

func BenchmarkBlendRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", linearLight{})
}

func BenchmarkBlendNRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{})
}

func BenchmarkBlendRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", linearLight{})
}

func TestBlendFallbackVividLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", vividLight{})
}

func TestBlendNRGBAToNRGBAVividLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{})
}

func TestBlendRGBAToNRGBAVividLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", vividLight{})
}

func TestBlendNRGBAToRGBAVividLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{})
}

func TestBlendRGBAToRGBAVividLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", vividLight{})
}

func BenchmarkBlendFallbackVividLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", vividLight{})
}

func BenchmarkBlendNRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{})
}

func BenchmarkBlendRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", vividLight{})
}

func BenchmarkBlendNRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{})
}

func BenchmarkBlendRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", vividLight{})
}

func TestBlendFallbackPinLight(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", pinLight{})
}

func TestBlendNRGBAToNRGBAPinLight(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{})
}

func TestBlendRGBAToNRGBAPinLight(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", pinLight{})
}

func TestBlendNRGBAToRGBAPinLight(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{})
}

func TestBlendRGBAToRGBAPinLight(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", pinLight{})
}

func BenchmarkBlendFallbackPinLight(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", pinLight{})
}

func BenchmarkBlendNRGBAToNRGBAPinLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{})
}

func BenchmarkBlendRGBAToNRGBAPinLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", pinLight{})
}

func BenchmarkBlendNRGBAToRGBAPinLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{})
}

func BenchmarkBlendRGBAToRGBAPinLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", pinLight{})
}

func TestBlendFallbackHardMix(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hardMix{})
}

func TestBlendNRGBAToNRGBAHardMix(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{})
}

func TestBlendRGBAToNRGBAHardMix(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hardMix{})
}

func TestBlendNRGBAToRGBAHardMix(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{})
}

func TestBlendRGBAToRGBAHardMix(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hardMix{})
}

func BenchmarkBlendFallbackHardMix(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hardMix{})
}

func BenchmarkBlendNRGBAToNRGBAHardMix(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{})
}

func BenchmarkBlendRGBAToNRGBAHardMix(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hardMix{})
}

func BenchmarkBlendNRGBAToRGBAHardMix(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{})
}

func BenchmarkBlendRGBAToRGBAHardMix(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hardMix{})
}

func TestBlendFallbackDifference(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", difference{})
}

func TestBlendNRGBAToNRGBADifference(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{})
}

func TestBlendRGBAToNRGBADifference(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", difference{})
}

func TestBlendNRGBAToRGBADifference(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{})
}

func TestBlendRGBAToRGBADifference(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", difference{})
}

func BenchmarkBlendFallbackDifference(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", difference{})
}

func BenchmarkBlendNRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{})
}

func BenchmarkBlendRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", difference{})
}

func BenchmarkBlendNRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{})
}

func BenchmarkBlendRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", difference{})
}

func TestBlendFallbackExclusion(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", exclusion{})
}

func TestBlendNRGBAToNRGBAExclusion(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{})
}

func TestBlendRGBAToNRGBAExclusion(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", exclusion{})
}

func TestBlendNRGBAToRGBAExclusion(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{})
}

func TestBlendRGBAToRGBAExclusion(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", exclusion{})
}

func BenchmarkBlendFallbackExclusion(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", exclusion{})
}

func BenchmarkBlendNRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{})
}

func BenchmarkBlendRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", exclusion{})
}

func BenchmarkBlendNRGBAToRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{})
}

func BenchmarkBlendRGBAToRGBAExclusion(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", exclusion{})
}

func TestBlendFallbackSubtract(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", subtract{})
}

func TestBlendNRGBAToNRGBASubtract(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{})
}

func TestBlendRGBAToNRGBASubtract(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", subtract{})
}

func TestBlendNRGBAToRGBASubtract(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{})
}

func TestBlendRGBAToRGBASubtract(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", subtract{})
}

func BenchmarkBlendFallbackSubtract(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", subtract{})
}

func BenchmarkBlendNRGBAToNRGBASubtract(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{})
}

func BenchmarkBlendRGBAToNRGBASubtract(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", subtract{})
}

func BenchmarkBlendNRGBAToRGBASubtract(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{})
}

func BenchmarkBlendRGBAToRGBASubtract(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", subtract{})
}

func TestBlendFallbackDivide(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", divide{})
}

func TestBlendNRGBAToNRGBADivide(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{})
}

func TestBlendRGBAToNRGBADivide(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", divide{})
}

func TestBlendNRGBAToRGBADivide(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{})
}

func TestBlendRGBAToRGBADivide(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", divide{})
}

func BenchmarkBlendFallbackDivide(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", divide{})
}

func BenchmarkBlendNRGBAToNRGBADivide(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{})
}

func BenchmarkBlendRGBAToNRGBADivide(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", divide{})
}

func BenchmarkBlendNRGBAToRGBADivide(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{})
}

func BenchmarkBlendRGBAToRGBADivide(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", divide{})
}

func TestBlendFallbackHue(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", hue{})
}

func TestBlendNRGBAToNRGBAHue(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{})
}

func TestBlendRGBAToNRGBAHue(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", hue{})
}

func TestBlendNRGBAToRGBAHue(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{})
}

func TestBlendRGBAToRGBAHue(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", hue{})
}

func BenchmarkBlendFallbackHue(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", hue{})
}

func BenchmarkBlendNRGBAToNRGBAHue(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{})
}

func BenchmarkBlendRGBAToNRGBAHue(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", hue{})
}

func BenchmarkBlendNRGBAToRGBAHue(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{})
}

func BenchmarkBlendRGBAToRGBAHue(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", hue{})
}

func TestBlendFallbackSaturation(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", saturation{})
}

func TestBlendNRGBAToNRGBASaturation(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{})
}

func TestBlendRGBAToNRGBASaturation(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", saturation{})
}

func TestBlendNRGBAToRGBASaturation(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{})
}

func TestBlendRGBAToRGBASaturation(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", saturation{})
}

func BenchmarkBlendFallbackSaturation(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", saturation{})
}

func BenchmarkBlendNRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{})
}

func BenchmarkBlendRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", saturation{})
}

func BenchmarkBlendNRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{})
}

func BenchmarkBlendRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", saturation{})
}

func TestBlendFallbackColor(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", color{})
}

func TestBlendNRGBAToNRGBAColor(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{})
}

func TestBlendRGBAToNRGBAColor(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", color{})
}

func TestBlendNRGBAToRGBAColor(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{})
}

func TestBlendRGBAToRGBAColor(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", color{})
}

func BenchmarkBlendFallbackColor(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", color{})
}

func BenchmarkBlendNRGBAToNRGBAColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{})
}

func BenchmarkBlendRGBAToNRGBAColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", color{})
}

func BenchmarkBlendNRGBAToRGBAColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{})
}

func BenchmarkBlendRGBAToRGBAColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", color{})
}

func TestBlendFallbackLuminosity(t *testing.T) {
	testDrawFallback(t, "png/bg.png", "png/fg.png", luminosity{})
}

func TestBlendNRGBAToNRGBALuminosity(t *testing.T) {
	testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{})
}

func TestBlendRGBAToNRGBALuminosity(t *testing.T) {
	testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", luminosity{})
}

func TestBlendNRGBAToRGBALuminosity(t *testing.T) {
	testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{})
}

func TestBlendRGBAToRGBALuminosity(t *testing.T) {
	testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", luminosity{})
}

func BenchmarkBlendFallbackLuminosity(b *testing.B) {
	benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", luminosity{})
}

func BenchmarkBlendNRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{})
}

func BenchmarkBlendRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", luminosity{})
}

func BenchmarkBlendNRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{})
}

func BenchmarkBlendRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", luminosity{})
}
