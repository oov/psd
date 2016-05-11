// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestDrawFallbackNormal(t *testing.T)       { testDrawFallback(normal{}, t, false) }
func TestDrawFallbackDarken(t *testing.T)       { testDrawFallback(darken{}, t, false) }
func TestDrawFallbackMultiply(t *testing.T)     { testDrawFallback(multiply{}, t, false) }
func TestDrawFallbackColorBurn(t *testing.T)    { testDrawFallback(colorburn{}, t, false) }
func TestDrawFallbackLinearBurn(t *testing.T)   { testDrawFallback(linearburn{}, t, false) }
func TestDrawFallbackDarkerColor(t *testing.T)  { testDrawFallback(darkercolor{}, t, false) }
func TestDrawFallbackLighten(t *testing.T)      { testDrawFallback(lighten{}, t, false) }
func TestDrawFallbackScreen(t *testing.T)       { testDrawFallback(screen{}, t, false) }
func TestDrawFallbackColorDodge(t *testing.T)   { testDrawFallback(colordodge{}, t, false) }
func TestDrawFallbackLinearDodge(t *testing.T)  { testDrawFallback(lineardodge{}, t, false) }
func TestDrawFallbackLighterColor(t *testing.T) { testDrawFallback(lightercolor{}, t, false) }
func TestDrawFallbackAdd(t *testing.T)          { testDrawFallback(add{}, t, false) }
func TestDrawFallbackOverlay(t *testing.T)      { testDrawFallback(overlay{}, t, false) }
func TestDrawFallbackSoftLight(t *testing.T)    { testDrawFallback(softlight{}, t, false) }
func TestDrawFallbackHardLight(t *testing.T)    { testDrawFallback(hardlight{}, t, false) }
func TestDrawFallbackLinearLight(t *testing.T)  { testDrawFallback(linearlight{}, t, false) }
func TestDrawFallbackVividLight(t *testing.T)   { testDrawFallback(vividlight{}, t, false) }
func TestDrawFallbackPinLight(t *testing.T)     { testDrawFallback(pinlight{}, t, false) }
func TestDrawFallbackHardMix(t *testing.T)      { testDrawFallback(hardmix{}, t, false) }
func TestDrawFallbackDifference(t *testing.T)   { testDrawFallback(difference{}, t, false) }
func TestDrawFallbackExclusion(t *testing.T)    { testDrawFallback(exclusion{}, t, false) }
func TestDrawFallbackSubtract(t *testing.T)     { testDrawFallback(subtract{}, t, false) }
func TestDrawFallbackDivide(t *testing.T)       { testDrawFallback(divide{}, t, false) }
func TestDrawFallbackHue(t *testing.T)          { testDrawFallback(hue{}, t, false) }
func TestDrawFallbackSaturation(t *testing.T)   { testDrawFallback(saturation{}, t, false) }
func TestDrawFallbackColor(t *testing.T)        { testDrawFallback(color{}, t, false) }
func TestDrawFallbackLuminosity(t *testing.T)   { testDrawFallback(luminosity{}, t, false) }

func TestDrawFallbackNormalProtectAlpha(t *testing.T)       { testDrawFallback(normal{}, t, true) }
func TestDrawFallbackDarkenProtectAlpha(t *testing.T)       { testDrawFallback(darken{}, t, true) }
func TestDrawFallbackMultiplyProtectAlpha(t *testing.T)     { testDrawFallback(multiply{}, t, true) }
func TestDrawFallbackColorBurnProtectAlpha(t *testing.T)    { testDrawFallback(colorburn{}, t, true) }
func TestDrawFallbackLinearBurnProtectAlpha(t *testing.T)   { testDrawFallback(linearburn{}, t, true) }
func TestDrawFallbackDarkerColorProtectAlpha(t *testing.T)  { testDrawFallback(darkercolor{}, t, true) }
func TestDrawFallbackLightenProtectAlpha(t *testing.T)      { testDrawFallback(lighten{}, t, true) }
func TestDrawFallbackScreenProtectAlpha(t *testing.T)       { testDrawFallback(screen{}, t, true) }
func TestDrawFallbackColorDodgeProtectAlpha(t *testing.T)   { testDrawFallback(colordodge{}, t, true) }
func TestDrawFallbackLinearDodgeProtectAlpha(t *testing.T)  { testDrawFallback(lineardodge{}, t, true) }
func TestDrawFallbackLighterColorProtectAlpha(t *testing.T) { testDrawFallback(lightercolor{}, t, true) }
func TestDrawFallbackAddProtectAlpha(t *testing.T)          { testDrawFallback(add{}, t, true) }
func TestDrawFallbackOverlayProtectAlpha(t *testing.T)      { testDrawFallback(overlay{}, t, true) }
func TestDrawFallbackSoftLightProtectAlpha(t *testing.T)    { testDrawFallback(softlight{}, t, true) }
func TestDrawFallbackHardLightProtectAlpha(t *testing.T)    { testDrawFallback(hardlight{}, t, true) }
func TestDrawFallbackLinearLightProtectAlpha(t *testing.T)  { testDrawFallback(linearlight{}, t, true) }
func TestDrawFallbackVividLightProtectAlpha(t *testing.T)   { testDrawFallback(vividlight{}, t, true) }
func TestDrawFallbackPinLightProtectAlpha(t *testing.T)     { testDrawFallback(pinlight{}, t, true) }
func TestDrawFallbackHardMixProtectAlpha(t *testing.T)      { testDrawFallback(hardmix{}, t, true) }
func TestDrawFallbackDifferenceProtectAlpha(t *testing.T)   { testDrawFallback(difference{}, t, true) }
func TestDrawFallbackExclusionProtectAlpha(t *testing.T)    { testDrawFallback(exclusion{}, t, true) }
func TestDrawFallbackSubtractProtectAlpha(t *testing.T)     { testDrawFallback(subtract{}, t, true) }
func TestDrawFallbackDivideProtectAlpha(t *testing.T)       { testDrawFallback(divide{}, t, true) }
func TestDrawFallbackHueProtectAlpha(t *testing.T)          { testDrawFallback(hue{}, t, true) }
func TestDrawFallbackSaturationProtectAlpha(t *testing.T)   { testDrawFallback(saturation{}, t, true) }
func TestDrawFallbackColorProtectAlpha(t *testing.T)        { testDrawFallback(color{}, t, true) }
func TestDrawFallbackLuminosityProtectAlpha(t *testing.T)   { testDrawFallback(luminosity{}, t, true) }

func TestDrawNRGBAToNRGBANormal(t *testing.T)       { testDrawNRGBAToNRGBA(normal{}, t, false) }
func TestDrawNRGBAToNRGBADarken(t *testing.T)       { testDrawNRGBAToNRGBA(darken{}, t, false) }
func TestDrawNRGBAToNRGBAMultiply(t *testing.T)     { testDrawNRGBAToNRGBA(multiply{}, t, false) }
func TestDrawNRGBAToNRGBAColorBurn(t *testing.T)    { testDrawNRGBAToNRGBA(colorburn{}, t, false) }
func TestDrawNRGBAToNRGBALinearBurn(t *testing.T)   { testDrawNRGBAToNRGBA(linearburn{}, t, false) }
func TestDrawNRGBAToNRGBADarkerColor(t *testing.T)  { testDrawNRGBAToNRGBA(darkercolor{}, t, false) }
func TestDrawNRGBAToNRGBALighten(t *testing.T)      { testDrawNRGBAToNRGBA(lighten{}, t, false) }
func TestDrawNRGBAToNRGBAScreen(t *testing.T)       { testDrawNRGBAToNRGBA(screen{}, t, false) }
func TestDrawNRGBAToNRGBAColorDodge(t *testing.T)   { testDrawNRGBAToNRGBA(colordodge{}, t, false) }
func TestDrawNRGBAToNRGBALinearDodge(t *testing.T)  { testDrawNRGBAToNRGBA(lineardodge{}, t, false) }
func TestDrawNRGBAToNRGBALighterColor(t *testing.T) { testDrawNRGBAToNRGBA(lightercolor{}, t, false) }
func TestDrawNRGBAToNRGBAAdd(t *testing.T)          { testDrawNRGBAToNRGBA(add{}, t, false) }
func TestDrawNRGBAToNRGBAOverlay(t *testing.T)      { testDrawNRGBAToNRGBA(overlay{}, t, false) }
func TestDrawNRGBAToNRGBASoftLight(t *testing.T)    { testDrawNRGBAToNRGBA(softlight{}, t, false) }
func TestDrawNRGBAToNRGBAHardLight(t *testing.T)    { testDrawNRGBAToNRGBA(hardlight{}, t, false) }
func TestDrawNRGBAToNRGBALinearLight(t *testing.T)  { testDrawNRGBAToNRGBA(linearlight{}, t, false) }
func TestDrawNRGBAToNRGBAVividLight(t *testing.T)   { testDrawNRGBAToNRGBA(vividlight{}, t, false) }
func TestDrawNRGBAToNRGBAPinLight(t *testing.T)     { testDrawNRGBAToNRGBA(pinlight{}, t, false) }
func TestDrawNRGBAToNRGBAHardMix(t *testing.T)      { testDrawNRGBAToNRGBA(hardmix{}, t, false) }
func TestDrawNRGBAToNRGBADifference(t *testing.T)   { testDrawNRGBAToNRGBA(difference{}, t, false) }
func TestDrawNRGBAToNRGBAExclusion(t *testing.T)    { testDrawNRGBAToNRGBA(exclusion{}, t, false) }
func TestDrawNRGBAToNRGBASubtract(t *testing.T)     { testDrawNRGBAToNRGBA(subtract{}, t, false) }
func TestDrawNRGBAToNRGBADivide(t *testing.T)       { testDrawNRGBAToNRGBA(divide{}, t, false) }
func TestDrawNRGBAToNRGBAHue(t *testing.T)          { testDrawNRGBAToNRGBA(hue{}, t, false) }
func TestDrawNRGBAToNRGBASaturation(t *testing.T)   { testDrawNRGBAToNRGBA(saturation{}, t, false) }
func TestDrawNRGBAToNRGBAColor(t *testing.T)        { testDrawNRGBAToNRGBA(color{}, t, false) }
func TestDrawNRGBAToNRGBALuminosity(t *testing.T)   { testDrawNRGBAToNRGBA(luminosity{}, t, false) }

func TestDrawNRGBAToNRGBANormalProtectAlpha(t *testing.T)   { testDrawNRGBAToNRGBA(normal{}, t, true) }
func TestDrawNRGBAToNRGBADarkenProtectAlpha(t *testing.T)   { testDrawNRGBAToNRGBA(darken{}, t, true) }
func TestDrawNRGBAToNRGBAMultiplyProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(multiply{}, t, true) }
func TestDrawNRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(colorburn{}, t, true)
}
func TestDrawNRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(linearburn{}, t, true)
}
func TestDrawNRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(darkercolor{}, t, true)
}
func TestDrawNRGBAToNRGBALightenProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(lighten{}, t, true) }
func TestDrawNRGBAToNRGBAScreenProtectAlpha(t *testing.T)  { testDrawNRGBAToNRGBA(screen{}, t, true) }
func TestDrawNRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(colordodge{}, t, true)
}
func TestDrawNRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(lineardodge{}, t, true)
}
func TestDrawNRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(lightercolor{}, t, true)
}
func TestDrawNRGBAToNRGBAAddProtectAlpha(t *testing.T)     { testDrawNRGBAToNRGBA(add{}, t, true) }
func TestDrawNRGBAToNRGBAOverlayProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(overlay{}, t, true) }
func TestDrawNRGBAToNRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(softlight{}, t, true)
}
func TestDrawNRGBAToNRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(hardlight{}, t, true)
}
func TestDrawNRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(linearlight{}, t, true)
}
func TestDrawNRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(vividlight{}, t, true)
}
func TestDrawNRGBAToNRGBAPinLightProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(pinlight{}, t, true) }
func TestDrawNRGBAToNRGBAHardMixProtectAlpha(t *testing.T)  { testDrawNRGBAToNRGBA(hardmix{}, t, true) }
func TestDrawNRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(difference{}, t, true)
}
func TestDrawNRGBAToNRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(exclusion{}, t, true)
}
func TestDrawNRGBAToNRGBASubtractProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(subtract{}, t, true) }
func TestDrawNRGBAToNRGBADivideProtectAlpha(t *testing.T)   { testDrawNRGBAToNRGBA(divide{}, t, true) }
func TestDrawNRGBAToNRGBAHueProtectAlpha(t *testing.T)      { testDrawNRGBAToNRGBA(hue{}, t, true) }
func TestDrawNRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(saturation{}, t, true)
}
func TestDrawNRGBAToNRGBAColorProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(color{}, t, true) }
func TestDrawNRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(luminosity{}, t, true)
}

func TestDrawRGBAToNRGBANormal(t *testing.T)       { testDrawRGBAToNRGBA(normal{}, t, false) }
func TestDrawRGBAToNRGBADarken(t *testing.T)       { testDrawRGBAToNRGBA(darken{}, t, false) }
func TestDrawRGBAToNRGBAMultiply(t *testing.T)     { testDrawRGBAToNRGBA(multiply{}, t, false) }
func TestDrawRGBAToNRGBAColorBurn(t *testing.T)    { testDrawRGBAToNRGBA(colorburn{}, t, false) }
func TestDrawRGBAToNRGBALinearBurn(t *testing.T)   { testDrawRGBAToNRGBA(linearburn{}, t, false) }
func TestDrawRGBAToNRGBADarkerColor(t *testing.T)  { testDrawRGBAToNRGBA(darkercolor{}, t, false) }
func TestDrawRGBAToNRGBALighten(t *testing.T)      { testDrawRGBAToNRGBA(lighten{}, t, false) }
func TestDrawRGBAToNRGBAScreen(t *testing.T)       { testDrawRGBAToNRGBA(screen{}, t, false) }
func TestDrawRGBAToNRGBAColorDodge(t *testing.T)   { testDrawRGBAToNRGBA(colordodge{}, t, false) }
func TestDrawRGBAToNRGBALinearDodge(t *testing.T)  { testDrawRGBAToNRGBA(lineardodge{}, t, false) }
func TestDrawRGBAToNRGBALighterColor(t *testing.T) { testDrawRGBAToNRGBA(lightercolor{}, t, false) }
func TestDrawRGBAToNRGBAAdd(t *testing.T)          { testDrawRGBAToNRGBA(add{}, t, false) }
func TestDrawRGBAToNRGBAOverlay(t *testing.T)      { testDrawRGBAToNRGBA(overlay{}, t, false) }
func TestDrawRGBAToNRGBASoftLight(t *testing.T)    { testDrawRGBAToNRGBA(softlight{}, t, false) }
func TestDrawRGBAToNRGBAHardLight(t *testing.T)    { testDrawRGBAToNRGBA(hardlight{}, t, false) }
func TestDrawRGBAToNRGBALinearLight(t *testing.T)  { testDrawRGBAToNRGBA(linearlight{}, t, false) }
func TestDrawRGBAToNRGBAVividLight(t *testing.T)   { testDrawRGBAToNRGBA(vividlight{}, t, false) }
func TestDrawRGBAToNRGBAPinLight(t *testing.T)     { testDrawRGBAToNRGBA(pinlight{}, t, false) }
func TestDrawRGBAToNRGBAHardMix(t *testing.T)      { testDrawRGBAToNRGBA(hardmix{}, t, false) }
func TestDrawRGBAToNRGBADifference(t *testing.T)   { testDrawRGBAToNRGBA(difference{}, t, false) }
func TestDrawRGBAToNRGBAExclusion(t *testing.T)    { testDrawRGBAToNRGBA(exclusion{}, t, false) }
func TestDrawRGBAToNRGBASubtract(t *testing.T)     { testDrawRGBAToNRGBA(subtract{}, t, false) }
func TestDrawRGBAToNRGBADivide(t *testing.T)       { testDrawRGBAToNRGBA(divide{}, t, false) }
func TestDrawRGBAToNRGBAHue(t *testing.T)          { testDrawRGBAToNRGBA(hue{}, t, false) }
func TestDrawRGBAToNRGBASaturation(t *testing.T)   { testDrawRGBAToNRGBA(saturation{}, t, false) }
func TestDrawRGBAToNRGBAColor(t *testing.T)        { testDrawRGBAToNRGBA(color{}, t, false) }
func TestDrawRGBAToNRGBALuminosity(t *testing.T)   { testDrawRGBAToNRGBA(luminosity{}, t, false) }

func TestDrawRGBAToNRGBANormalProtectAlpha(t *testing.T)    { testDrawRGBAToNRGBA(normal{}, t, true) }
func TestDrawRGBAToNRGBADarkenProtectAlpha(t *testing.T)    { testDrawRGBAToNRGBA(darken{}, t, true) }
func TestDrawRGBAToNRGBAMultiplyProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(multiply{}, t, true) }
func TestDrawRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(colorburn{}, t, true) }
func TestDrawRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(linearburn{}, t, true)
}
func TestDrawRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(darkercolor{}, t, true)
}
func TestDrawRGBAToNRGBALightenProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(lighten{}, t, true) }
func TestDrawRGBAToNRGBAScreenProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(screen{}, t, true) }
func TestDrawRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(colordodge{}, t, true)
}
func TestDrawRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(lineardodge{}, t, true)
}
func TestDrawRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(lightercolor{}, t, true)
}
func TestDrawRGBAToNRGBAAddProtectAlpha(t *testing.T)       { testDrawRGBAToNRGBA(add{}, t, true) }
func TestDrawRGBAToNRGBAOverlayProtectAlpha(t *testing.T)   { testDrawRGBAToNRGBA(overlay{}, t, true) }
func TestDrawRGBAToNRGBASoftLightProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(softlight{}, t, true) }
func TestDrawRGBAToNRGBAHardLightProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(hardlight{}, t, true) }
func TestDrawRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(linearlight{}, t, true)
}
func TestDrawRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(vividlight{}, t, true)
}
func TestDrawRGBAToNRGBAPinLightProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(pinlight{}, t, true) }
func TestDrawRGBAToNRGBAHardMixProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(hardmix{}, t, true) }
func TestDrawRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(difference{}, t, true)
}
func TestDrawRGBAToNRGBAExclusionProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(exclusion{}, t, true) }
func TestDrawRGBAToNRGBASubtractProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(subtract{}, t, true) }
func TestDrawRGBAToNRGBADivideProtectAlpha(t *testing.T)    { testDrawRGBAToNRGBA(divide{}, t, true) }
func TestDrawRGBAToNRGBAHueProtectAlpha(t *testing.T)       { testDrawRGBAToNRGBA(hue{}, t, true) }
func TestDrawRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(saturation{}, t, true)
}
func TestDrawRGBAToNRGBAColorProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(color{}, t, true) }
func TestDrawRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(luminosity{}, t, true)
}

func TestDrawNRGBAToRGBANormal(t *testing.T)       { testDrawNRGBAToRGBA(normal{}, t, false) }
func TestDrawNRGBAToRGBADarken(t *testing.T)       { testDrawNRGBAToRGBA(darken{}, t, false) }
func TestDrawNRGBAToRGBAMultiply(t *testing.T)     { testDrawNRGBAToRGBA(multiply{}, t, false) }
func TestDrawNRGBAToRGBAColorBurn(t *testing.T)    { testDrawNRGBAToRGBA(colorburn{}, t, false) }
func TestDrawNRGBAToRGBALinearBurn(t *testing.T)   { testDrawNRGBAToRGBA(linearburn{}, t, false) }
func TestDrawNRGBAToRGBADarkerColor(t *testing.T)  { testDrawNRGBAToRGBA(darkercolor{}, t, false) }
func TestDrawNRGBAToRGBALighten(t *testing.T)      { testDrawNRGBAToRGBA(lighten{}, t, false) }
func TestDrawNRGBAToRGBAScreen(t *testing.T)       { testDrawNRGBAToRGBA(screen{}, t, false) }
func TestDrawNRGBAToRGBAColorDodge(t *testing.T)   { testDrawNRGBAToRGBA(colordodge{}, t, false) }
func TestDrawNRGBAToRGBALinearDodge(t *testing.T)  { testDrawNRGBAToRGBA(lineardodge{}, t, false) }
func TestDrawNRGBAToRGBALighterColor(t *testing.T) { testDrawNRGBAToRGBA(lightercolor{}, t, false) }
func TestDrawNRGBAToRGBAAdd(t *testing.T)          { testDrawNRGBAToRGBA(add{}, t, false) }
func TestDrawNRGBAToRGBAOverlay(t *testing.T)      { testDrawNRGBAToRGBA(overlay{}, t, false) }
func TestDrawNRGBAToRGBASoftLight(t *testing.T)    { testDrawNRGBAToRGBA(softlight{}, t, false) }
func TestDrawNRGBAToRGBAHardLight(t *testing.T)    { testDrawNRGBAToRGBA(hardlight{}, t, false) }
func TestDrawNRGBAToRGBALinearLight(t *testing.T)  { testDrawNRGBAToRGBA(linearlight{}, t, false) }
func TestDrawNRGBAToRGBAVividLight(t *testing.T)   { testDrawNRGBAToRGBA(vividlight{}, t, false) }
func TestDrawNRGBAToRGBAPinLight(t *testing.T)     { testDrawNRGBAToRGBA(pinlight{}, t, false) }
func TestDrawNRGBAToRGBAHardMix(t *testing.T)      { testDrawNRGBAToRGBA(hardmix{}, t, false) }
func TestDrawNRGBAToRGBADifference(t *testing.T)   { testDrawNRGBAToRGBA(difference{}, t, false) }
func TestDrawNRGBAToRGBAExclusion(t *testing.T)    { testDrawNRGBAToRGBA(exclusion{}, t, false) }
func TestDrawNRGBAToRGBASubtract(t *testing.T)     { testDrawNRGBAToRGBA(subtract{}, t, false) }
func TestDrawNRGBAToRGBADivide(t *testing.T)       { testDrawNRGBAToRGBA(divide{}, t, false) }
func TestDrawNRGBAToRGBAHue(t *testing.T)          { testDrawNRGBAToRGBA(hue{}, t, false) }
func TestDrawNRGBAToRGBASaturation(t *testing.T)   { testDrawNRGBAToRGBA(saturation{}, t, false) }
func TestDrawNRGBAToRGBAColor(t *testing.T)        { testDrawNRGBAToRGBA(color{}, t, false) }
func TestDrawNRGBAToRGBALuminosity(t *testing.T)   { testDrawNRGBAToRGBA(luminosity{}, t, false) }

func TestDrawNRGBAToRGBANormalProtectAlpha(t *testing.T)    { testDrawNRGBAToRGBA(normal{}, t, true) }
func TestDrawNRGBAToRGBADarkenProtectAlpha(t *testing.T)    { testDrawNRGBAToRGBA(darken{}, t, true) }
func TestDrawNRGBAToRGBAMultiplyProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(multiply{}, t, true) }
func TestDrawNRGBAToRGBAColorBurnProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(colorburn{}, t, true) }
func TestDrawNRGBAToRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(linearburn{}, t, true)
}
func TestDrawNRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(darkercolor{}, t, true)
}
func TestDrawNRGBAToRGBALightenProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(lighten{}, t, true) }
func TestDrawNRGBAToRGBAScreenProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(screen{}, t, true) }
func TestDrawNRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(colordodge{}, t, true)
}
func TestDrawNRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(lineardodge{}, t, true)
}
func TestDrawNRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(lightercolor{}, t, true)
}
func TestDrawNRGBAToRGBAAddProtectAlpha(t *testing.T)       { testDrawNRGBAToRGBA(add{}, t, true) }
func TestDrawNRGBAToRGBAOverlayProtectAlpha(t *testing.T)   { testDrawNRGBAToRGBA(overlay{}, t, true) }
func TestDrawNRGBAToRGBASoftLightProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(softlight{}, t, true) }
func TestDrawNRGBAToRGBAHardLightProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(hardlight{}, t, true) }
func TestDrawNRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(linearlight{}, t, true)
}
func TestDrawNRGBAToRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(vividlight{}, t, true)
}
func TestDrawNRGBAToRGBAPinLightProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(pinlight{}, t, true) }
func TestDrawNRGBAToRGBAHardMixProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(hardmix{}, t, true) }
func TestDrawNRGBAToRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(difference{}, t, true)
}
func TestDrawNRGBAToRGBAExclusionProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(exclusion{}, t, true) }
func TestDrawNRGBAToRGBASubtractProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(subtract{}, t, true) }
func TestDrawNRGBAToRGBADivideProtectAlpha(t *testing.T)    { testDrawNRGBAToRGBA(divide{}, t, true) }
func TestDrawNRGBAToRGBAHueProtectAlpha(t *testing.T)       { testDrawNRGBAToRGBA(hue{}, t, true) }
func TestDrawNRGBAToRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(saturation{}, t, true)
}
func TestDrawNRGBAToRGBAColorProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(color{}, t, true) }
func TestDrawNRGBAToRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(luminosity{}, t, true)
}

func TestDrawRGBAToRGBANormal(t *testing.T)       { testDrawRGBAToRGBA(normal{}, t, false) }
func TestDrawRGBAToRGBADarken(t *testing.T)       { testDrawRGBAToRGBA(darken{}, t, false) }
func TestDrawRGBAToRGBAMultiply(t *testing.T)     { testDrawRGBAToRGBA(multiply{}, t, false) }
func TestDrawRGBAToRGBAColorBurn(t *testing.T)    { testDrawRGBAToRGBA(colorburn{}, t, false) }
func TestDrawRGBAToRGBALinearBurn(t *testing.T)   { testDrawRGBAToRGBA(linearburn{}, t, false) }
func TestDrawRGBAToRGBADarkerColor(t *testing.T)  { testDrawRGBAToRGBA(darkercolor{}, t, false) }
func TestDrawRGBAToRGBALighten(t *testing.T)      { testDrawRGBAToRGBA(lighten{}, t, false) }
func TestDrawRGBAToRGBAScreen(t *testing.T)       { testDrawRGBAToRGBA(screen{}, t, false) }
func TestDrawRGBAToRGBAColorDodge(t *testing.T)   { testDrawRGBAToRGBA(colordodge{}, t, false) }
func TestDrawRGBAToRGBALinearDodge(t *testing.T)  { testDrawRGBAToRGBA(lineardodge{}, t, false) }
func TestDrawRGBAToRGBALighterColor(t *testing.T) { testDrawRGBAToRGBA(lightercolor{}, t, false) }
func TestDrawRGBAToRGBAAdd(t *testing.T)          { testDrawRGBAToRGBA(add{}, t, false) }
func TestDrawRGBAToRGBAOverlay(t *testing.T)      { testDrawRGBAToRGBA(overlay{}, t, false) }
func TestDrawRGBAToRGBASoftLight(t *testing.T)    { testDrawRGBAToRGBA(softlight{}, t, false) }
func TestDrawRGBAToRGBAHardLight(t *testing.T)    { testDrawRGBAToRGBA(hardlight{}, t, false) }
func TestDrawRGBAToRGBALinearLight(t *testing.T)  { testDrawRGBAToRGBA(linearlight{}, t, false) }
func TestDrawRGBAToRGBAVividLight(t *testing.T)   { testDrawRGBAToRGBA(vividlight{}, t, false) }
func TestDrawRGBAToRGBAPinLight(t *testing.T)     { testDrawRGBAToRGBA(pinlight{}, t, false) }
func TestDrawRGBAToRGBAHardMix(t *testing.T)      { testDrawRGBAToRGBA(hardmix{}, t, false) }
func TestDrawRGBAToRGBADifference(t *testing.T)   { testDrawRGBAToRGBA(difference{}, t, false) }
func TestDrawRGBAToRGBAExclusion(t *testing.T)    { testDrawRGBAToRGBA(exclusion{}, t, false) }
func TestDrawRGBAToRGBASubtract(t *testing.T)     { testDrawRGBAToRGBA(subtract{}, t, false) }
func TestDrawRGBAToRGBADivide(t *testing.T)       { testDrawRGBAToRGBA(divide{}, t, false) }
func TestDrawRGBAToRGBAHue(t *testing.T)          { testDrawRGBAToRGBA(hue{}, t, false) }
func TestDrawRGBAToRGBASaturation(t *testing.T)   { testDrawRGBAToRGBA(saturation{}, t, false) }
func TestDrawRGBAToRGBAColor(t *testing.T)        { testDrawRGBAToRGBA(color{}, t, false) }
func TestDrawRGBAToRGBALuminosity(t *testing.T)   { testDrawRGBAToRGBA(luminosity{}, t, false) }

func TestDrawRGBAToRGBANormalProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(normal{}, t, true) }
func TestDrawRGBAToRGBADarkenProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(darken{}, t, true) }
func TestDrawRGBAToRGBAMultiplyProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(multiply{}, t, true) }
func TestDrawRGBAToRGBAColorBurnProtectAlpha(t *testing.T)  { testDrawRGBAToRGBA(colorburn{}, t, true) }
func TestDrawRGBAToRGBALinearBurnProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(linearburn{}, t, true) }
func TestDrawRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(darkercolor{}, t, true)
}
func TestDrawRGBAToRGBALightenProtectAlpha(t *testing.T)    { testDrawRGBAToRGBA(lighten{}, t, true) }
func TestDrawRGBAToRGBAScreenProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(screen{}, t, true) }
func TestDrawRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(colordodge{}, t, true) }
func TestDrawRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(lineardodge{}, t, true)
}
func TestDrawRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(lightercolor{}, t, true)
}
func TestDrawRGBAToRGBAAddProtectAlpha(t *testing.T)       { testDrawRGBAToRGBA(add{}, t, true) }
func TestDrawRGBAToRGBAOverlayProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(overlay{}, t, true) }
func TestDrawRGBAToRGBASoftLightProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(softlight{}, t, true) }
func TestDrawRGBAToRGBAHardLightProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(hardlight{}, t, true) }
func TestDrawRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(linearlight{}, t, true)
}
func TestDrawRGBAToRGBAVividLightProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(vividlight{}, t, true) }
func TestDrawRGBAToRGBAPinLightProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(pinlight{}, t, true) }
func TestDrawRGBAToRGBAHardMixProtectAlpha(t *testing.T)    { testDrawRGBAToRGBA(hardmix{}, t, true) }
func TestDrawRGBAToRGBADifferenceProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(difference{}, t, true) }
func TestDrawRGBAToRGBAExclusionProtectAlpha(t *testing.T)  { testDrawRGBAToRGBA(exclusion{}, t, true) }
func TestDrawRGBAToRGBASubtractProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(subtract{}, t, true) }
func TestDrawRGBAToRGBADivideProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(divide{}, t, true) }
func TestDrawRGBAToRGBAHueProtectAlpha(t *testing.T)        { testDrawRGBAToRGBA(hue{}, t, true) }
func TestDrawRGBAToRGBASaturationProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(saturation{}, t, true) }
func TestDrawRGBAToRGBAColorProtectAlpha(t *testing.T)      { testDrawRGBAToRGBA(color{}, t, true) }
func TestDrawRGBAToRGBALuminosityProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(luminosity{}, t, true) }

func BenchmarkDrawFallbackNormal(b *testing.B)       { benchmarkDrawFallback(normal{}, b, false) }
func BenchmarkDrawFallbackDarken(b *testing.B)       { benchmarkDrawFallback(darken{}, b, false) }
func BenchmarkDrawFallbackMultiply(b *testing.B)     { benchmarkDrawFallback(multiply{}, b, false) }
func BenchmarkDrawFallbackColorBurn(b *testing.B)    { benchmarkDrawFallback(colorburn{}, b, false) }
func BenchmarkDrawFallbackLinearBurn(b *testing.B)   { benchmarkDrawFallback(linearburn{}, b, false) }
func BenchmarkDrawFallbackDarkerColor(b *testing.B)  { benchmarkDrawFallback(darkercolor{}, b, false) }
func BenchmarkDrawFallbackLighten(b *testing.B)      { benchmarkDrawFallback(lighten{}, b, false) }
func BenchmarkDrawFallbackScreen(b *testing.B)       { benchmarkDrawFallback(screen{}, b, false) }
func BenchmarkDrawFallbackColorDodge(b *testing.B)   { benchmarkDrawFallback(colordodge{}, b, false) }
func BenchmarkDrawFallbackLinearDodge(b *testing.B)  { benchmarkDrawFallback(lineardodge{}, b, false) }
func BenchmarkDrawFallbackLighterColor(b *testing.B) { benchmarkDrawFallback(lightercolor{}, b, false) }
func BenchmarkDrawFallbackAdd(b *testing.B)          { benchmarkDrawFallback(add{}, b, false) }
func BenchmarkDrawFallbackOverlay(b *testing.B)      { benchmarkDrawFallback(overlay{}, b, false) }
func BenchmarkDrawFallbackSoftLight(b *testing.B)    { benchmarkDrawFallback(softlight{}, b, false) }
func BenchmarkDrawFallbackHardLight(b *testing.B)    { benchmarkDrawFallback(hardlight{}, b, false) }
func BenchmarkDrawFallbackLinearLight(b *testing.B)  { benchmarkDrawFallback(linearlight{}, b, false) }
func BenchmarkDrawFallbackVividLight(b *testing.B)   { benchmarkDrawFallback(vividlight{}, b, false) }
func BenchmarkDrawFallbackPinLight(b *testing.B)     { benchmarkDrawFallback(pinlight{}, b, false) }
func BenchmarkDrawFallbackHardMix(b *testing.B)      { benchmarkDrawFallback(hardmix{}, b, false) }
func BenchmarkDrawFallbackDifference(b *testing.B)   { benchmarkDrawFallback(difference{}, b, false) }
func BenchmarkDrawFallbackExclusion(b *testing.B)    { benchmarkDrawFallback(exclusion{}, b, false) }
func BenchmarkDrawFallbackSubtract(b *testing.B)     { benchmarkDrawFallback(subtract{}, b, false) }
func BenchmarkDrawFallbackDivide(b *testing.B)       { benchmarkDrawFallback(divide{}, b, false) }
func BenchmarkDrawFallbackHue(b *testing.B)          { benchmarkDrawFallback(hue{}, b, false) }
func BenchmarkDrawFallbackSaturation(b *testing.B)   { benchmarkDrawFallback(saturation{}, b, false) }
func BenchmarkDrawFallbackColor(b *testing.B)        { benchmarkDrawFallback(color{}, b, false) }
func BenchmarkDrawFallbackLuminosity(b *testing.B)   { benchmarkDrawFallback(luminosity{}, b, false) }

func BenchmarkDrawFallbackNormalProtectAlpha(b *testing.B) { benchmarkDrawFallback(normal{}, b, true) }
func BenchmarkDrawFallbackDarkenProtectAlpha(b *testing.B) { benchmarkDrawFallback(darken{}, b, true) }
func BenchmarkDrawFallbackMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(multiply{}, b, true)
}
func BenchmarkDrawFallbackColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(colorburn{}, b, true)
}
func BenchmarkDrawFallbackLinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(linearburn{}, b, true)
}
func BenchmarkDrawFallbackDarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(darkercolor{}, b, true)
}
func BenchmarkDrawFallbackLightenProtectAlpha(b *testing.B) { benchmarkDrawFallback(lighten{}, b, true) }
func BenchmarkDrawFallbackScreenProtectAlpha(b *testing.B)  { benchmarkDrawFallback(screen{}, b, true) }
func BenchmarkDrawFallbackColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(colordodge{}, b, true)
}
func BenchmarkDrawFallbackLinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(lineardodge{}, b, true)
}
func BenchmarkDrawFallbackLighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(lightercolor{}, b, true)
}
func BenchmarkDrawFallbackAddProtectAlpha(b *testing.B)     { benchmarkDrawFallback(add{}, b, true) }
func BenchmarkDrawFallbackOverlayProtectAlpha(b *testing.B) { benchmarkDrawFallback(overlay{}, b, true) }
func BenchmarkDrawFallbackSoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(softlight{}, b, true)
}
func BenchmarkDrawFallbackHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(hardlight{}, b, true)
}
func BenchmarkDrawFallbackLinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(linearlight{}, b, true)
}
func BenchmarkDrawFallbackVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(vividlight{}, b, true)
}
func BenchmarkDrawFallbackPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(pinlight{}, b, true)
}
func BenchmarkDrawFallbackHardMixProtectAlpha(b *testing.B) { benchmarkDrawFallback(hardmix{}, b, true) }
func BenchmarkDrawFallbackDifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(difference{}, b, true)
}
func BenchmarkDrawFallbackExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(exclusion{}, b, true)
}
func BenchmarkDrawFallbackSubtractProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(subtract{}, b, true)
}
func BenchmarkDrawFallbackDivideProtectAlpha(b *testing.B) { benchmarkDrawFallback(divide{}, b, true) }
func BenchmarkDrawFallbackHueProtectAlpha(b *testing.B)    { benchmarkDrawFallback(hue{}, b, true) }
func BenchmarkDrawFallbackSaturationProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(saturation{}, b, true)
}
func BenchmarkDrawFallbackColorProtectAlpha(b *testing.B) { benchmarkDrawFallback(color{}, b, true) }
func BenchmarkDrawFallbackLuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(luminosity{}, b, true)
}

func BenchmarkDrawNRGBAToNRGBANormal(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(normal{}, b, false) }
func BenchmarkDrawNRGBAToNRGBADarken(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(darken{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAMultiply(b *testing.B) { benchmarkDrawNRGBAToNRGBA(multiply{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(colorburn{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(linearburn{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(darkercolor{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALighten(b *testing.B) { benchmarkDrawNRGBAToNRGBA(lighten{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAScreen(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(screen{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(colordodge{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(lineardodge{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(lightercolor{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAAdd(b *testing.B)     { benchmarkDrawNRGBAToNRGBA(add{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAOverlay(b *testing.B) { benchmarkDrawNRGBAToNRGBA(overlay{}, b, false) }
func BenchmarkDrawNRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(softlight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(hardlight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(linearlight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(vividlight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAPinLight(b *testing.B) { benchmarkDrawNRGBAToNRGBA(pinlight{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAHardMix(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(hardmix{}, b, false) }
func BenchmarkDrawNRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(difference{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(exclusion{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBASubtract(b *testing.B) { benchmarkDrawNRGBAToNRGBA(subtract{}, b, false) }
func BenchmarkDrawNRGBAToNRGBADivide(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(divide{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAHue(b *testing.B)      { benchmarkDrawNRGBAToNRGBA(hue{}, b, false) }
func BenchmarkDrawNRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(saturation{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAColor(b *testing.B) { benchmarkDrawNRGBAToNRGBA(color{}, b, false) }
func BenchmarkDrawNRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(luminosity{}, b, false)
}

func BenchmarkDrawNRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(normal{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(darken{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(multiply{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(colorburn{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(linearburn{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(darkercolor{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(lighten{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(screen{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(colordodge{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(lineardodge{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(lightercolor{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToNRGBA(add{}, b, true) }
func BenchmarkDrawNRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(overlay{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(softlight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(hardlight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(linearlight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(vividlight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(pinlight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(hardmix{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(difference{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(exclusion{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(subtract{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(divide{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToNRGBA(hue{}, b, true) }
func BenchmarkDrawNRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(saturation{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(color{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(luminosity{}, b, true)
}

func BenchmarkDrawRGBAToNRGBANormal(b *testing.B)    { benchmarkDrawRGBAToNRGBA(normal{}, b, false) }
func BenchmarkDrawRGBAToNRGBADarken(b *testing.B)    { benchmarkDrawRGBAToNRGBA(darken{}, b, false) }
func BenchmarkDrawRGBAToNRGBAMultiply(b *testing.B)  { benchmarkDrawRGBAToNRGBA(multiply{}, b, false) }
func BenchmarkDrawRGBAToNRGBAColorBurn(b *testing.B) { benchmarkDrawRGBAToNRGBA(colorburn{}, b, false) }
func BenchmarkDrawRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(linearburn{}, b, false)
}
func BenchmarkDrawRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(darkercolor{}, b, false)
}
func BenchmarkDrawRGBAToNRGBALighten(b *testing.B) { benchmarkDrawRGBAToNRGBA(lighten{}, b, false) }
func BenchmarkDrawRGBAToNRGBAScreen(b *testing.B)  { benchmarkDrawRGBAToNRGBA(screen{}, b, false) }
func BenchmarkDrawRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(colordodge{}, b, false)
}
func BenchmarkDrawRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(lineardodge{}, b, false)
}
func BenchmarkDrawRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(lightercolor{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAAdd(b *testing.B)       { benchmarkDrawRGBAToNRGBA(add{}, b, false) }
func BenchmarkDrawRGBAToNRGBAOverlay(b *testing.B)   { benchmarkDrawRGBAToNRGBA(overlay{}, b, false) }
func BenchmarkDrawRGBAToNRGBASoftLight(b *testing.B) { benchmarkDrawRGBAToNRGBA(softlight{}, b, false) }
func BenchmarkDrawRGBAToNRGBAHardLight(b *testing.B) { benchmarkDrawRGBAToNRGBA(hardlight{}, b, false) }
func BenchmarkDrawRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(linearlight{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(vividlight{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAPinLight(b *testing.B) { benchmarkDrawRGBAToNRGBA(pinlight{}, b, false) }
func BenchmarkDrawRGBAToNRGBAHardMix(b *testing.B)  { benchmarkDrawRGBAToNRGBA(hardmix{}, b, false) }
func BenchmarkDrawRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(difference{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAExclusion(b *testing.B) { benchmarkDrawRGBAToNRGBA(exclusion{}, b, false) }
func BenchmarkDrawRGBAToNRGBASubtract(b *testing.B)  { benchmarkDrawRGBAToNRGBA(subtract{}, b, false) }
func BenchmarkDrawRGBAToNRGBADivide(b *testing.B)    { benchmarkDrawRGBAToNRGBA(divide{}, b, false) }
func BenchmarkDrawRGBAToNRGBAHue(b *testing.B)       { benchmarkDrawRGBAToNRGBA(hue{}, b, false) }
func BenchmarkDrawRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(saturation{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAColor(b *testing.B) { benchmarkDrawRGBAToNRGBA(color{}, b, false) }
func BenchmarkDrawRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(luminosity{}, b, false)
}

func BenchmarkDrawRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(normal{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(darken{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(multiply{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(colorburn{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(linearburn{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(darkercolor{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(lighten{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(screen{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(colordodge{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(lineardodge{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(lightercolor{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawRGBAToNRGBA(add{}, b, true) }
func BenchmarkDrawRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(overlay{}, b, true)
}
func BenchmarkDrawRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(softlight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(hardlight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(linearlight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(vividlight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(pinlight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(hardmix{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(difference{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(exclusion{}, b, true)
}
func BenchmarkDrawRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(subtract{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(divide{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawRGBAToNRGBA(hue{}, b, true) }
func BenchmarkDrawRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(saturation{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(color{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(luminosity{}, b, true)
}

func BenchmarkDrawNRGBAToRGBANormal(b *testing.B)    { benchmarkDrawNRGBAToRGBA(normal{}, b, false) }
func BenchmarkDrawNRGBAToRGBADarken(b *testing.B)    { benchmarkDrawNRGBAToRGBA(darken{}, b, false) }
func BenchmarkDrawNRGBAToRGBAMultiply(b *testing.B)  { benchmarkDrawNRGBAToRGBA(multiply{}, b, false) }
func BenchmarkDrawNRGBAToRGBAColorBurn(b *testing.B) { benchmarkDrawNRGBAToRGBA(colorburn{}, b, false) }
func BenchmarkDrawNRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(linearburn{}, b, false)
}
func BenchmarkDrawNRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(darkercolor{}, b, false)
}
func BenchmarkDrawNRGBAToRGBALighten(b *testing.B) { benchmarkDrawNRGBAToRGBA(lighten{}, b, false) }
func BenchmarkDrawNRGBAToRGBAScreen(b *testing.B)  { benchmarkDrawNRGBAToRGBA(screen{}, b, false) }
func BenchmarkDrawNRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(colordodge{}, b, false)
}
func BenchmarkDrawNRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(lineardodge{}, b, false)
}
func BenchmarkDrawNRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(lightercolor{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAAdd(b *testing.B)       { benchmarkDrawNRGBAToRGBA(add{}, b, false) }
func BenchmarkDrawNRGBAToRGBAOverlay(b *testing.B)   { benchmarkDrawNRGBAToRGBA(overlay{}, b, false) }
func BenchmarkDrawNRGBAToRGBASoftLight(b *testing.B) { benchmarkDrawNRGBAToRGBA(softlight{}, b, false) }
func BenchmarkDrawNRGBAToRGBAHardLight(b *testing.B) { benchmarkDrawNRGBAToRGBA(hardlight{}, b, false) }
func BenchmarkDrawNRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(linearlight{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(vividlight{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAPinLight(b *testing.B) { benchmarkDrawNRGBAToRGBA(pinlight{}, b, false) }
func BenchmarkDrawNRGBAToRGBAHardMix(b *testing.B)  { benchmarkDrawNRGBAToRGBA(hardmix{}, b, false) }
func BenchmarkDrawNRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(difference{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAExclusion(b *testing.B) { benchmarkDrawNRGBAToRGBA(exclusion{}, b, false) }
func BenchmarkDrawNRGBAToRGBASubtract(b *testing.B)  { benchmarkDrawNRGBAToRGBA(subtract{}, b, false) }
func BenchmarkDrawNRGBAToRGBADivide(b *testing.B)    { benchmarkDrawNRGBAToRGBA(divide{}, b, false) }
func BenchmarkDrawNRGBAToRGBAHue(b *testing.B)       { benchmarkDrawNRGBAToRGBA(hue{}, b, false) }
func BenchmarkDrawNRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(saturation{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAColor(b *testing.B) { benchmarkDrawNRGBAToRGBA(color{}, b, false) }
func BenchmarkDrawNRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(luminosity{}, b, false)
}

func BenchmarkDrawNRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(normal{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(darken{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(multiply{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(colorburn{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(linearburn{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(darkercolor{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(lighten{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(screen{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(colordodge{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(lineardodge{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(lightercolor{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToRGBA(add{}, b, true) }
func BenchmarkDrawNRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(overlay{}, b, true)
}
func BenchmarkDrawNRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(softlight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(hardlight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(linearlight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(vividlight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(pinlight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(hardmix{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(difference{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(exclusion{}, b, true)
}
func BenchmarkDrawNRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(subtract{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(divide{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToRGBA(hue{}, b, true) }
func BenchmarkDrawNRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(saturation{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(color{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(luminosity{}, b, true)
}

func BenchmarkDrawRGBAToRGBANormal(b *testing.B)     { benchmarkDrawRGBAToRGBA(normal{}, b, false) }
func BenchmarkDrawRGBAToRGBADarken(b *testing.B)     { benchmarkDrawRGBAToRGBA(darken{}, b, false) }
func BenchmarkDrawRGBAToRGBAMultiply(b *testing.B)   { benchmarkDrawRGBAToRGBA(multiply{}, b, false) }
func BenchmarkDrawRGBAToRGBAColorBurn(b *testing.B)  { benchmarkDrawRGBAToRGBA(colorburn{}, b, false) }
func BenchmarkDrawRGBAToRGBALinearBurn(b *testing.B) { benchmarkDrawRGBAToRGBA(linearburn{}, b, false) }
func BenchmarkDrawRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(darkercolor{}, b, false)
}
func BenchmarkDrawRGBAToRGBALighten(b *testing.B)    { benchmarkDrawRGBAToRGBA(lighten{}, b, false) }
func BenchmarkDrawRGBAToRGBAScreen(b *testing.B)     { benchmarkDrawRGBAToRGBA(screen{}, b, false) }
func BenchmarkDrawRGBAToRGBAColorDodge(b *testing.B) { benchmarkDrawRGBAToRGBA(colordodge{}, b, false) }
func BenchmarkDrawRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(lineardodge{}, b, false)
}
func BenchmarkDrawRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(lightercolor{}, b, false)
}
func BenchmarkDrawRGBAToRGBAAdd(b *testing.B)       { benchmarkDrawRGBAToRGBA(add{}, b, false) }
func BenchmarkDrawRGBAToRGBAOverlay(b *testing.B)   { benchmarkDrawRGBAToRGBA(overlay{}, b, false) }
func BenchmarkDrawRGBAToRGBASoftLight(b *testing.B) { benchmarkDrawRGBAToRGBA(softlight{}, b, false) }
func BenchmarkDrawRGBAToRGBAHardLight(b *testing.B) { benchmarkDrawRGBAToRGBA(hardlight{}, b, false) }
func BenchmarkDrawRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(linearlight{}, b, false)
}
func BenchmarkDrawRGBAToRGBAVividLight(b *testing.B) { benchmarkDrawRGBAToRGBA(vividlight{}, b, false) }
func BenchmarkDrawRGBAToRGBAPinLight(b *testing.B)   { benchmarkDrawRGBAToRGBA(pinlight{}, b, false) }
func BenchmarkDrawRGBAToRGBAHardMix(b *testing.B)    { benchmarkDrawRGBAToRGBA(hardmix{}, b, false) }
func BenchmarkDrawRGBAToRGBADifference(b *testing.B) { benchmarkDrawRGBAToRGBA(difference{}, b, false) }
func BenchmarkDrawRGBAToRGBAExclusion(b *testing.B)  { benchmarkDrawRGBAToRGBA(exclusion{}, b, false) }
func BenchmarkDrawRGBAToRGBASubtract(b *testing.B)   { benchmarkDrawRGBAToRGBA(subtract{}, b, false) }
func BenchmarkDrawRGBAToRGBADivide(b *testing.B)     { benchmarkDrawRGBAToRGBA(divide{}, b, false) }
func BenchmarkDrawRGBAToRGBAHue(b *testing.B)        { benchmarkDrawRGBAToRGBA(hue{}, b, false) }
func BenchmarkDrawRGBAToRGBASaturation(b *testing.B) { benchmarkDrawRGBAToRGBA(saturation{}, b, false) }
func BenchmarkDrawRGBAToRGBAColor(b *testing.B)      { benchmarkDrawRGBAToRGBA(color{}, b, false) }
func BenchmarkDrawRGBAToRGBALuminosity(b *testing.B) { benchmarkDrawRGBAToRGBA(luminosity{}, b, false) }

func BenchmarkDrawRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(normal{}, b, true)
}
func BenchmarkDrawRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(darken{}, b, true)
}
func BenchmarkDrawRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(multiply{}, b, true)
}
func BenchmarkDrawRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(colorburn{}, b, true)
}
func BenchmarkDrawRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(linearburn{}, b, true)
}
func BenchmarkDrawRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(darkercolor{}, b, true)
}
func BenchmarkDrawRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(lighten{}, b, true)
}
func BenchmarkDrawRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(screen{}, b, true)
}
func BenchmarkDrawRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(colordodge{}, b, true)
}
func BenchmarkDrawRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(lineardodge{}, b, true)
}
func BenchmarkDrawRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(lightercolor{}, b, true)
}
func BenchmarkDrawRGBAToRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawRGBAToRGBA(add{}, b, true) }
func BenchmarkDrawRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(overlay{}, b, true)
}
func BenchmarkDrawRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(softlight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(hardlight{}, b, true)
}
func BenchmarkDrawRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(linearlight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(vividlight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(pinlight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(hardmix{}, b, true)
}
func BenchmarkDrawRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(difference{}, b, true)
}
func BenchmarkDrawRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(exclusion{}, b, true)
}
func BenchmarkDrawRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(subtract{}, b, true)
}
func BenchmarkDrawRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(divide{}, b, true)
}
func BenchmarkDrawRGBAToRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawRGBAToRGBA(hue{}, b, true) }
func BenchmarkDrawRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(saturation{}, b, true)
}
func BenchmarkDrawRGBAToRGBAColorProtectAlpha(b *testing.B) { benchmarkDrawRGBAToRGBA(color{}, b, true) }
func BenchmarkDrawRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(luminosity{}, b, true)
}
