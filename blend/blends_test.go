// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

func TestDrawFallbackNormal(t *testing.T)       { testDrawFallback(Normal{}, t, false) }
func TestDrawFallbackDarken(t *testing.T)       { testDrawFallback(Darken{}, t, false) }
func TestDrawFallbackMultiply(t *testing.T)     { testDrawFallback(Multiply{}, t, false) }
func TestDrawFallbackColorBurn(t *testing.T)    { testDrawFallback(ColorBurn{}, t, false) }
func TestDrawFallbackLinearBurn(t *testing.T)   { testDrawFallback(LinearBurn{}, t, false) }
func TestDrawFallbackDarkerColor(t *testing.T)  { testDrawFallback(DarkerColor{}, t, false) }
func TestDrawFallbackLighten(t *testing.T)      { testDrawFallback(Lighten{}, t, false) }
func TestDrawFallbackScreen(t *testing.T)       { testDrawFallback(Screen{}, t, false) }
func TestDrawFallbackColorDodge(t *testing.T)   { testDrawFallback(ColorDodge{}, t, false) }
func TestDrawFallbackLinearDodge(t *testing.T)  { testDrawFallback(LinearDodge{}, t, false) }
func TestDrawFallbackLighterColor(t *testing.T) { testDrawFallback(LighterColor{}, t, false) }
func TestDrawFallbackAdd(t *testing.T)          { testDrawFallback(Add{}, t, false) }
func TestDrawFallbackOverlay(t *testing.T)      { testDrawFallback(Overlay{}, t, false) }
func TestDrawFallbackSoftLight(t *testing.T)    { testDrawFallback(SoftLight{}, t, false) }
func TestDrawFallbackHardLight(t *testing.T)    { testDrawFallback(HardLight{}, t, false) }
func TestDrawFallbackLinearLight(t *testing.T)  { testDrawFallback(LinearLight{}, t, false) }
func TestDrawFallbackVividLight(t *testing.T)   { testDrawFallback(VividLight{}, t, false) }
func TestDrawFallbackPinLight(t *testing.T)     { testDrawFallback(PinLight{}, t, false) }
func TestDrawFallbackHardMix(t *testing.T)      { testDrawFallback(HardMix{}, t, false) }
func TestDrawFallbackDifference(t *testing.T)   { testDrawFallback(Difference{}, t, false) }
func TestDrawFallbackExclusion(t *testing.T)    { testDrawFallback(Exclusion{}, t, false) }
func TestDrawFallbackSubtract(t *testing.T)     { testDrawFallback(Subtract{}, t, false) }
func TestDrawFallbackDivide(t *testing.T)       { testDrawFallback(Divide{}, t, false) }
func TestDrawFallbackHue(t *testing.T)          { testDrawFallback(Hue{}, t, false) }
func TestDrawFallbackSaturation(t *testing.T)   { testDrawFallback(Saturation{}, t, false) }
func TestDrawFallbackColor(t *testing.T)        { testDrawFallback(Color{}, t, false) }
func TestDrawFallbackLuminosity(t *testing.T)   { testDrawFallback(Luminosity{}, t, false) }

func TestDrawFallbackNormalProtectAlpha(t *testing.T)       { testDrawFallback(Normal{}, t, true) }
func TestDrawFallbackDarkenProtectAlpha(t *testing.T)       { testDrawFallback(Darken{}, t, true) }
func TestDrawFallbackMultiplyProtectAlpha(t *testing.T)     { testDrawFallback(Multiply{}, t, true) }
func TestDrawFallbackColorBurnProtectAlpha(t *testing.T)    { testDrawFallback(ColorBurn{}, t, true) }
func TestDrawFallbackLinearBurnProtectAlpha(t *testing.T)   { testDrawFallback(LinearBurn{}, t, true) }
func TestDrawFallbackDarkerColorProtectAlpha(t *testing.T)  { testDrawFallback(DarkerColor{}, t, true) }
func TestDrawFallbackLightenProtectAlpha(t *testing.T)      { testDrawFallback(Lighten{}, t, true) }
func TestDrawFallbackScreenProtectAlpha(t *testing.T)       { testDrawFallback(Screen{}, t, true) }
func TestDrawFallbackColorDodgeProtectAlpha(t *testing.T)   { testDrawFallback(ColorDodge{}, t, true) }
func TestDrawFallbackLinearDodgeProtectAlpha(t *testing.T)  { testDrawFallback(LinearDodge{}, t, true) }
func TestDrawFallbackLighterColorProtectAlpha(t *testing.T) { testDrawFallback(LighterColor{}, t, true) }
func TestDrawFallbackAddProtectAlpha(t *testing.T)          { testDrawFallback(Add{}, t, true) }
func TestDrawFallbackOverlayProtectAlpha(t *testing.T)      { testDrawFallback(Overlay{}, t, true) }
func TestDrawFallbackSoftLightProtectAlpha(t *testing.T)    { testDrawFallback(SoftLight{}, t, true) }
func TestDrawFallbackHardLightProtectAlpha(t *testing.T)    { testDrawFallback(HardLight{}, t, true) }
func TestDrawFallbackLinearLightProtectAlpha(t *testing.T)  { testDrawFallback(LinearLight{}, t, true) }
func TestDrawFallbackVividLightProtectAlpha(t *testing.T)   { testDrawFallback(VividLight{}, t, true) }
func TestDrawFallbackPinLightProtectAlpha(t *testing.T)     { testDrawFallback(PinLight{}, t, true) }
func TestDrawFallbackHardMixProtectAlpha(t *testing.T)      { testDrawFallback(HardMix{}, t, true) }
func TestDrawFallbackDifferenceProtectAlpha(t *testing.T)   { testDrawFallback(Difference{}, t, true) }
func TestDrawFallbackExclusionProtectAlpha(t *testing.T)    { testDrawFallback(Exclusion{}, t, true) }
func TestDrawFallbackSubtractProtectAlpha(t *testing.T)     { testDrawFallback(Subtract{}, t, true) }
func TestDrawFallbackDivideProtectAlpha(t *testing.T)       { testDrawFallback(Divide{}, t, true) }
func TestDrawFallbackHueProtectAlpha(t *testing.T)          { testDrawFallback(Hue{}, t, true) }
func TestDrawFallbackSaturationProtectAlpha(t *testing.T)   { testDrawFallback(Saturation{}, t, true) }
func TestDrawFallbackColorProtectAlpha(t *testing.T)        { testDrawFallback(Color{}, t, true) }
func TestDrawFallbackLuminosityProtectAlpha(t *testing.T)   { testDrawFallback(Luminosity{}, t, true) }

func TestDrawNRGBAToNRGBANormal(t *testing.T)       { testDrawNRGBAToNRGBA(Normal{}, t, false) }
func TestDrawNRGBAToNRGBADarken(t *testing.T)       { testDrawNRGBAToNRGBA(Darken{}, t, false) }
func TestDrawNRGBAToNRGBAMultiply(t *testing.T)     { testDrawNRGBAToNRGBA(Multiply{}, t, false) }
func TestDrawNRGBAToNRGBAColorBurn(t *testing.T)    { testDrawNRGBAToNRGBA(ColorBurn{}, t, false) }
func TestDrawNRGBAToNRGBALinearBurn(t *testing.T)   { testDrawNRGBAToNRGBA(LinearBurn{}, t, false) }
func TestDrawNRGBAToNRGBADarkerColor(t *testing.T)  { testDrawNRGBAToNRGBA(DarkerColor{}, t, false) }
func TestDrawNRGBAToNRGBALighten(t *testing.T)      { testDrawNRGBAToNRGBA(Lighten{}, t, false) }
func TestDrawNRGBAToNRGBAScreen(t *testing.T)       { testDrawNRGBAToNRGBA(Screen{}, t, false) }
func TestDrawNRGBAToNRGBAColorDodge(t *testing.T)   { testDrawNRGBAToNRGBA(ColorDodge{}, t, false) }
func TestDrawNRGBAToNRGBALinearDodge(t *testing.T)  { testDrawNRGBAToNRGBA(LinearDodge{}, t, false) }
func TestDrawNRGBAToNRGBALighterColor(t *testing.T) { testDrawNRGBAToNRGBA(LighterColor{}, t, false) }
func TestDrawNRGBAToNRGBAAdd(t *testing.T)          { testDrawNRGBAToNRGBA(Add{}, t, false) }
func TestDrawNRGBAToNRGBAOverlay(t *testing.T)      { testDrawNRGBAToNRGBA(Overlay{}, t, false) }
func TestDrawNRGBAToNRGBASoftLight(t *testing.T)    { testDrawNRGBAToNRGBA(SoftLight{}, t, false) }
func TestDrawNRGBAToNRGBAHardLight(t *testing.T)    { testDrawNRGBAToNRGBA(HardLight{}, t, false) }
func TestDrawNRGBAToNRGBALinearLight(t *testing.T)  { testDrawNRGBAToNRGBA(LinearLight{}, t, false) }
func TestDrawNRGBAToNRGBAVividLight(t *testing.T)   { testDrawNRGBAToNRGBA(VividLight{}, t, false) }
func TestDrawNRGBAToNRGBAPinLight(t *testing.T)     { testDrawNRGBAToNRGBA(PinLight{}, t, false) }
func TestDrawNRGBAToNRGBAHardMix(t *testing.T)      { testDrawNRGBAToNRGBA(HardMix{}, t, false) }
func TestDrawNRGBAToNRGBADifference(t *testing.T)   { testDrawNRGBAToNRGBA(Difference{}, t, false) }
func TestDrawNRGBAToNRGBAExclusion(t *testing.T)    { testDrawNRGBAToNRGBA(Exclusion{}, t, false) }
func TestDrawNRGBAToNRGBASubtract(t *testing.T)     { testDrawNRGBAToNRGBA(Subtract{}, t, false) }
func TestDrawNRGBAToNRGBADivide(t *testing.T)       { testDrawNRGBAToNRGBA(Divide{}, t, false) }
func TestDrawNRGBAToNRGBAHue(t *testing.T)          { testDrawNRGBAToNRGBA(Hue{}, t, false) }
func TestDrawNRGBAToNRGBASaturation(t *testing.T)   { testDrawNRGBAToNRGBA(Saturation{}, t, false) }
func TestDrawNRGBAToNRGBAColor(t *testing.T)        { testDrawNRGBAToNRGBA(Color{}, t, false) }
func TestDrawNRGBAToNRGBALuminosity(t *testing.T)   { testDrawNRGBAToNRGBA(Luminosity{}, t, false) }

func TestDrawNRGBAToNRGBANormalProtectAlpha(t *testing.T)   { testDrawNRGBAToNRGBA(Normal{}, t, true) }
func TestDrawNRGBAToNRGBADarkenProtectAlpha(t *testing.T)   { testDrawNRGBAToNRGBA(Darken{}, t, true) }
func TestDrawNRGBAToNRGBAMultiplyProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(Multiply{}, t, true) }
func TestDrawNRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(ColorBurn{}, t, true)
}
func TestDrawNRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(LinearBurn{}, t, true)
}
func TestDrawNRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(DarkerColor{}, t, true)
}
func TestDrawNRGBAToNRGBALightenProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(Lighten{}, t, true) }
func TestDrawNRGBAToNRGBAScreenProtectAlpha(t *testing.T)  { testDrawNRGBAToNRGBA(Screen{}, t, true) }
func TestDrawNRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(ColorDodge{}, t, true)
}
func TestDrawNRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(LinearDodge{}, t, true)
}
func TestDrawNRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(LighterColor{}, t, true)
}
func TestDrawNRGBAToNRGBAAddProtectAlpha(t *testing.T)     { testDrawNRGBAToNRGBA(Add{}, t, true) }
func TestDrawNRGBAToNRGBAOverlayProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(Overlay{}, t, true) }
func TestDrawNRGBAToNRGBASoftLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(SoftLight{}, t, true)
}
func TestDrawNRGBAToNRGBAHardLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(HardLight{}, t, true)
}
func TestDrawNRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(LinearLight{}, t, true)
}
func TestDrawNRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(VividLight{}, t, true)
}
func TestDrawNRGBAToNRGBAPinLightProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(PinLight{}, t, true) }
func TestDrawNRGBAToNRGBAHardMixProtectAlpha(t *testing.T)  { testDrawNRGBAToNRGBA(HardMix{}, t, true) }
func TestDrawNRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(Difference{}, t, true)
}
func TestDrawNRGBAToNRGBAExclusionProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(Exclusion{}, t, true)
}
func TestDrawNRGBAToNRGBASubtractProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(Subtract{}, t, true) }
func TestDrawNRGBAToNRGBADivideProtectAlpha(t *testing.T)   { testDrawNRGBAToNRGBA(Divide{}, t, true) }
func TestDrawNRGBAToNRGBAHueProtectAlpha(t *testing.T)      { testDrawNRGBAToNRGBA(Hue{}, t, true) }
func TestDrawNRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(Saturation{}, t, true)
}
func TestDrawNRGBAToNRGBAColorProtectAlpha(t *testing.T) { testDrawNRGBAToNRGBA(Color{}, t, true) }
func TestDrawNRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToNRGBA(Luminosity{}, t, true)
}

func TestDrawRGBAToNRGBANormal(t *testing.T)       { testDrawRGBAToNRGBA(Normal{}, t, false) }
func TestDrawRGBAToNRGBADarken(t *testing.T)       { testDrawRGBAToNRGBA(Darken{}, t, false) }
func TestDrawRGBAToNRGBAMultiply(t *testing.T)     { testDrawRGBAToNRGBA(Multiply{}, t, false) }
func TestDrawRGBAToNRGBAColorBurn(t *testing.T)    { testDrawRGBAToNRGBA(ColorBurn{}, t, false) }
func TestDrawRGBAToNRGBALinearBurn(t *testing.T)   { testDrawRGBAToNRGBA(LinearBurn{}, t, false) }
func TestDrawRGBAToNRGBADarkerColor(t *testing.T)  { testDrawRGBAToNRGBA(DarkerColor{}, t, false) }
func TestDrawRGBAToNRGBALighten(t *testing.T)      { testDrawRGBAToNRGBA(Lighten{}, t, false) }
func TestDrawRGBAToNRGBAScreen(t *testing.T)       { testDrawRGBAToNRGBA(Screen{}, t, false) }
func TestDrawRGBAToNRGBAColorDodge(t *testing.T)   { testDrawRGBAToNRGBA(ColorDodge{}, t, false) }
func TestDrawRGBAToNRGBALinearDodge(t *testing.T)  { testDrawRGBAToNRGBA(LinearDodge{}, t, false) }
func TestDrawRGBAToNRGBALighterColor(t *testing.T) { testDrawRGBAToNRGBA(LighterColor{}, t, false) }
func TestDrawRGBAToNRGBAAdd(t *testing.T)          { testDrawRGBAToNRGBA(Add{}, t, false) }
func TestDrawRGBAToNRGBAOverlay(t *testing.T)      { testDrawRGBAToNRGBA(Overlay{}, t, false) }
func TestDrawRGBAToNRGBASoftLight(t *testing.T)    { testDrawRGBAToNRGBA(SoftLight{}, t, false) }
func TestDrawRGBAToNRGBAHardLight(t *testing.T)    { testDrawRGBAToNRGBA(HardLight{}, t, false) }
func TestDrawRGBAToNRGBALinearLight(t *testing.T)  { testDrawRGBAToNRGBA(LinearLight{}, t, false) }
func TestDrawRGBAToNRGBAVividLight(t *testing.T)   { testDrawRGBAToNRGBA(VividLight{}, t, false) }
func TestDrawRGBAToNRGBAPinLight(t *testing.T)     { testDrawRGBAToNRGBA(PinLight{}, t, false) }
func TestDrawRGBAToNRGBAHardMix(t *testing.T)      { testDrawRGBAToNRGBA(HardMix{}, t, false) }
func TestDrawRGBAToNRGBADifference(t *testing.T)   { testDrawRGBAToNRGBA(Difference{}, t, false) }
func TestDrawRGBAToNRGBAExclusion(t *testing.T)    { testDrawRGBAToNRGBA(Exclusion{}, t, false) }
func TestDrawRGBAToNRGBASubtract(t *testing.T)     { testDrawRGBAToNRGBA(Subtract{}, t, false) }
func TestDrawRGBAToNRGBADivide(t *testing.T)       { testDrawRGBAToNRGBA(Divide{}, t, false) }
func TestDrawRGBAToNRGBAHue(t *testing.T)          { testDrawRGBAToNRGBA(Hue{}, t, false) }
func TestDrawRGBAToNRGBASaturation(t *testing.T)   { testDrawRGBAToNRGBA(Saturation{}, t, false) }
func TestDrawRGBAToNRGBAColor(t *testing.T)        { testDrawRGBAToNRGBA(Color{}, t, false) }
func TestDrawRGBAToNRGBALuminosity(t *testing.T)   { testDrawRGBAToNRGBA(Luminosity{}, t, false) }

func TestDrawRGBAToNRGBANormalProtectAlpha(t *testing.T)    { testDrawRGBAToNRGBA(Normal{}, t, true) }
func TestDrawRGBAToNRGBADarkenProtectAlpha(t *testing.T)    { testDrawRGBAToNRGBA(Darken{}, t, true) }
func TestDrawRGBAToNRGBAMultiplyProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(Multiply{}, t, true) }
func TestDrawRGBAToNRGBAColorBurnProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(ColorBurn{}, t, true) }
func TestDrawRGBAToNRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(LinearBurn{}, t, true)
}
func TestDrawRGBAToNRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(DarkerColor{}, t, true)
}
func TestDrawRGBAToNRGBALightenProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(Lighten{}, t, true) }
func TestDrawRGBAToNRGBAScreenProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(Screen{}, t, true) }
func TestDrawRGBAToNRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(ColorDodge{}, t, true)
}
func TestDrawRGBAToNRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(LinearDodge{}, t, true)
}
func TestDrawRGBAToNRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(LighterColor{}, t, true)
}
func TestDrawRGBAToNRGBAAddProtectAlpha(t *testing.T)       { testDrawRGBAToNRGBA(Add{}, t, true) }
func TestDrawRGBAToNRGBAOverlayProtectAlpha(t *testing.T)   { testDrawRGBAToNRGBA(Overlay{}, t, true) }
func TestDrawRGBAToNRGBASoftLightProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(SoftLight{}, t, true) }
func TestDrawRGBAToNRGBAHardLightProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(HardLight{}, t, true) }
func TestDrawRGBAToNRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(LinearLight{}, t, true)
}
func TestDrawRGBAToNRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(VividLight{}, t, true)
}
func TestDrawRGBAToNRGBAPinLightProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(PinLight{}, t, true) }
func TestDrawRGBAToNRGBAHardMixProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(HardMix{}, t, true) }
func TestDrawRGBAToNRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(Difference{}, t, true)
}
func TestDrawRGBAToNRGBAExclusionProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(Exclusion{}, t, true) }
func TestDrawRGBAToNRGBASubtractProtectAlpha(t *testing.T)  { testDrawRGBAToNRGBA(Subtract{}, t, true) }
func TestDrawRGBAToNRGBADivideProtectAlpha(t *testing.T)    { testDrawRGBAToNRGBA(Divide{}, t, true) }
func TestDrawRGBAToNRGBAHueProtectAlpha(t *testing.T)       { testDrawRGBAToNRGBA(Hue{}, t, true) }
func TestDrawRGBAToNRGBASaturationProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(Saturation{}, t, true)
}
func TestDrawRGBAToNRGBAColorProtectAlpha(t *testing.T) { testDrawRGBAToNRGBA(Color{}, t, true) }
func TestDrawRGBAToNRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawRGBAToNRGBA(Luminosity{}, t, true)
}

func TestDrawNRGBAToRGBANormal(t *testing.T)       { testDrawNRGBAToRGBA(Normal{}, t, false) }
func TestDrawNRGBAToRGBADarken(t *testing.T)       { testDrawNRGBAToRGBA(Darken{}, t, false) }
func TestDrawNRGBAToRGBAMultiply(t *testing.T)     { testDrawNRGBAToRGBA(Multiply{}, t, false) }
func TestDrawNRGBAToRGBAColorBurn(t *testing.T)    { testDrawNRGBAToRGBA(ColorBurn{}, t, false) }
func TestDrawNRGBAToRGBALinearBurn(t *testing.T)   { testDrawNRGBAToRGBA(LinearBurn{}, t, false) }
func TestDrawNRGBAToRGBADarkerColor(t *testing.T)  { testDrawNRGBAToRGBA(DarkerColor{}, t, false) }
func TestDrawNRGBAToRGBALighten(t *testing.T)      { testDrawNRGBAToRGBA(Lighten{}, t, false) }
func TestDrawNRGBAToRGBAScreen(t *testing.T)       { testDrawNRGBAToRGBA(Screen{}, t, false) }
func TestDrawNRGBAToRGBAColorDodge(t *testing.T)   { testDrawNRGBAToRGBA(ColorDodge{}, t, false) }
func TestDrawNRGBAToRGBALinearDodge(t *testing.T)  { testDrawNRGBAToRGBA(LinearDodge{}, t, false) }
func TestDrawNRGBAToRGBALighterColor(t *testing.T) { testDrawNRGBAToRGBA(LighterColor{}, t, false) }
func TestDrawNRGBAToRGBAAdd(t *testing.T)          { testDrawNRGBAToRGBA(Add{}, t, false) }
func TestDrawNRGBAToRGBAOverlay(t *testing.T)      { testDrawNRGBAToRGBA(Overlay{}, t, false) }
func TestDrawNRGBAToRGBASoftLight(t *testing.T)    { testDrawNRGBAToRGBA(SoftLight{}, t, false) }
func TestDrawNRGBAToRGBAHardLight(t *testing.T)    { testDrawNRGBAToRGBA(HardLight{}, t, false) }
func TestDrawNRGBAToRGBALinearLight(t *testing.T)  { testDrawNRGBAToRGBA(LinearLight{}, t, false) }
func TestDrawNRGBAToRGBAVividLight(t *testing.T)   { testDrawNRGBAToRGBA(VividLight{}, t, false) }
func TestDrawNRGBAToRGBAPinLight(t *testing.T)     { testDrawNRGBAToRGBA(PinLight{}, t, false) }
func TestDrawNRGBAToRGBAHardMix(t *testing.T)      { testDrawNRGBAToRGBA(HardMix{}, t, false) }
func TestDrawNRGBAToRGBADifference(t *testing.T)   { testDrawNRGBAToRGBA(Difference{}, t, false) }
func TestDrawNRGBAToRGBAExclusion(t *testing.T)    { testDrawNRGBAToRGBA(Exclusion{}, t, false) }
func TestDrawNRGBAToRGBASubtract(t *testing.T)     { testDrawNRGBAToRGBA(Subtract{}, t, false) }
func TestDrawNRGBAToRGBADivide(t *testing.T)       { testDrawNRGBAToRGBA(Divide{}, t, false) }
func TestDrawNRGBAToRGBAHue(t *testing.T)          { testDrawNRGBAToRGBA(Hue{}, t, false) }
func TestDrawNRGBAToRGBASaturation(t *testing.T)   { testDrawNRGBAToRGBA(Saturation{}, t, false) }
func TestDrawNRGBAToRGBAColor(t *testing.T)        { testDrawNRGBAToRGBA(Color{}, t, false) }
func TestDrawNRGBAToRGBALuminosity(t *testing.T)   { testDrawNRGBAToRGBA(Luminosity{}, t, false) }

func TestDrawNRGBAToRGBANormalProtectAlpha(t *testing.T)    { testDrawNRGBAToRGBA(Normal{}, t, true) }
func TestDrawNRGBAToRGBADarkenProtectAlpha(t *testing.T)    { testDrawNRGBAToRGBA(Darken{}, t, true) }
func TestDrawNRGBAToRGBAMultiplyProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(Multiply{}, t, true) }
func TestDrawNRGBAToRGBAColorBurnProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(ColorBurn{}, t, true) }
func TestDrawNRGBAToRGBALinearBurnProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(LinearBurn{}, t, true)
}
func TestDrawNRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(DarkerColor{}, t, true)
}
func TestDrawNRGBAToRGBALightenProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(Lighten{}, t, true) }
func TestDrawNRGBAToRGBAScreenProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(Screen{}, t, true) }
func TestDrawNRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(ColorDodge{}, t, true)
}
func TestDrawNRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(LinearDodge{}, t, true)
}
func TestDrawNRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(LighterColor{}, t, true)
}
func TestDrawNRGBAToRGBAAddProtectAlpha(t *testing.T)       { testDrawNRGBAToRGBA(Add{}, t, true) }
func TestDrawNRGBAToRGBAOverlayProtectAlpha(t *testing.T)   { testDrawNRGBAToRGBA(Overlay{}, t, true) }
func TestDrawNRGBAToRGBASoftLightProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(SoftLight{}, t, true) }
func TestDrawNRGBAToRGBAHardLightProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(HardLight{}, t, true) }
func TestDrawNRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(LinearLight{}, t, true)
}
func TestDrawNRGBAToRGBAVividLightProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(VividLight{}, t, true)
}
func TestDrawNRGBAToRGBAPinLightProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(PinLight{}, t, true) }
func TestDrawNRGBAToRGBAHardMixProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(HardMix{}, t, true) }
func TestDrawNRGBAToRGBADifferenceProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(Difference{}, t, true)
}
func TestDrawNRGBAToRGBAExclusionProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(Exclusion{}, t, true) }
func TestDrawNRGBAToRGBASubtractProtectAlpha(t *testing.T)  { testDrawNRGBAToRGBA(Subtract{}, t, true) }
func TestDrawNRGBAToRGBADivideProtectAlpha(t *testing.T)    { testDrawNRGBAToRGBA(Divide{}, t, true) }
func TestDrawNRGBAToRGBAHueProtectAlpha(t *testing.T)       { testDrawNRGBAToRGBA(Hue{}, t, true) }
func TestDrawNRGBAToRGBASaturationProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(Saturation{}, t, true)
}
func TestDrawNRGBAToRGBAColorProtectAlpha(t *testing.T) { testDrawNRGBAToRGBA(Color{}, t, true) }
func TestDrawNRGBAToRGBALuminosityProtectAlpha(t *testing.T) {
	testDrawNRGBAToRGBA(Luminosity{}, t, true)
}

func TestDrawRGBAToRGBANormal(t *testing.T)       { testDrawRGBAToRGBA(Normal{}, t, false) }
func TestDrawRGBAToRGBADarken(t *testing.T)       { testDrawRGBAToRGBA(Darken{}, t, false) }
func TestDrawRGBAToRGBAMultiply(t *testing.T)     { testDrawRGBAToRGBA(Multiply{}, t, false) }
func TestDrawRGBAToRGBAColorBurn(t *testing.T)    { testDrawRGBAToRGBA(ColorBurn{}, t, false) }
func TestDrawRGBAToRGBALinearBurn(t *testing.T)   { testDrawRGBAToRGBA(LinearBurn{}, t, false) }
func TestDrawRGBAToRGBADarkerColor(t *testing.T)  { testDrawRGBAToRGBA(DarkerColor{}, t, false) }
func TestDrawRGBAToRGBALighten(t *testing.T)      { testDrawRGBAToRGBA(Lighten{}, t, false) }
func TestDrawRGBAToRGBAScreen(t *testing.T)       { testDrawRGBAToRGBA(Screen{}, t, false) }
func TestDrawRGBAToRGBAColorDodge(t *testing.T)   { testDrawRGBAToRGBA(ColorDodge{}, t, false) }
func TestDrawRGBAToRGBALinearDodge(t *testing.T)  { testDrawRGBAToRGBA(LinearDodge{}, t, false) }
func TestDrawRGBAToRGBALighterColor(t *testing.T) { testDrawRGBAToRGBA(LighterColor{}, t, false) }
func TestDrawRGBAToRGBAAdd(t *testing.T)          { testDrawRGBAToRGBA(Add{}, t, false) }
func TestDrawRGBAToRGBAOverlay(t *testing.T)      { testDrawRGBAToRGBA(Overlay{}, t, false) }
func TestDrawRGBAToRGBASoftLight(t *testing.T)    { testDrawRGBAToRGBA(SoftLight{}, t, false) }
func TestDrawRGBAToRGBAHardLight(t *testing.T)    { testDrawRGBAToRGBA(HardLight{}, t, false) }
func TestDrawRGBAToRGBALinearLight(t *testing.T)  { testDrawRGBAToRGBA(LinearLight{}, t, false) }
func TestDrawRGBAToRGBAVividLight(t *testing.T)   { testDrawRGBAToRGBA(VividLight{}, t, false) }
func TestDrawRGBAToRGBAPinLight(t *testing.T)     { testDrawRGBAToRGBA(PinLight{}, t, false) }
func TestDrawRGBAToRGBAHardMix(t *testing.T)      { testDrawRGBAToRGBA(HardMix{}, t, false) }
func TestDrawRGBAToRGBADifference(t *testing.T)   { testDrawRGBAToRGBA(Difference{}, t, false) }
func TestDrawRGBAToRGBAExclusion(t *testing.T)    { testDrawRGBAToRGBA(Exclusion{}, t, false) }
func TestDrawRGBAToRGBASubtract(t *testing.T)     { testDrawRGBAToRGBA(Subtract{}, t, false) }
func TestDrawRGBAToRGBADivide(t *testing.T)       { testDrawRGBAToRGBA(Divide{}, t, false) }
func TestDrawRGBAToRGBAHue(t *testing.T)          { testDrawRGBAToRGBA(Hue{}, t, false) }
func TestDrawRGBAToRGBASaturation(t *testing.T)   { testDrawRGBAToRGBA(Saturation{}, t, false) }
func TestDrawRGBAToRGBAColor(t *testing.T)        { testDrawRGBAToRGBA(Color{}, t, false) }
func TestDrawRGBAToRGBALuminosity(t *testing.T)   { testDrawRGBAToRGBA(Luminosity{}, t, false) }

func TestDrawRGBAToRGBANormalProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(Normal{}, t, true) }
func TestDrawRGBAToRGBADarkenProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(Darken{}, t, true) }
func TestDrawRGBAToRGBAMultiplyProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(Multiply{}, t, true) }
func TestDrawRGBAToRGBAColorBurnProtectAlpha(t *testing.T)  { testDrawRGBAToRGBA(ColorBurn{}, t, true) }
func TestDrawRGBAToRGBALinearBurnProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(LinearBurn{}, t, true) }
func TestDrawRGBAToRGBADarkerColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(DarkerColor{}, t, true)
}
func TestDrawRGBAToRGBALightenProtectAlpha(t *testing.T)    { testDrawRGBAToRGBA(Lighten{}, t, true) }
func TestDrawRGBAToRGBAScreenProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(Screen{}, t, true) }
func TestDrawRGBAToRGBAColorDodgeProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(ColorDodge{}, t, true) }
func TestDrawRGBAToRGBALinearDodgeProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(LinearDodge{}, t, true)
}
func TestDrawRGBAToRGBALighterColorProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(LighterColor{}, t, true)
}
func TestDrawRGBAToRGBAAddProtectAlpha(t *testing.T)       { testDrawRGBAToRGBA(Add{}, t, true) }
func TestDrawRGBAToRGBAOverlayProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(Overlay{}, t, true) }
func TestDrawRGBAToRGBASoftLightProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(SoftLight{}, t, true) }
func TestDrawRGBAToRGBAHardLightProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(HardLight{}, t, true) }
func TestDrawRGBAToRGBALinearLightProtectAlpha(t *testing.T) {
	testDrawRGBAToRGBA(LinearLight{}, t, true)
}
func TestDrawRGBAToRGBAVividLightProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(VividLight{}, t, true) }
func TestDrawRGBAToRGBAPinLightProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(PinLight{}, t, true) }
func TestDrawRGBAToRGBAHardMixProtectAlpha(t *testing.T)    { testDrawRGBAToRGBA(HardMix{}, t, true) }
func TestDrawRGBAToRGBADifferenceProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(Difference{}, t, true) }
func TestDrawRGBAToRGBAExclusionProtectAlpha(t *testing.T)  { testDrawRGBAToRGBA(Exclusion{}, t, true) }
func TestDrawRGBAToRGBASubtractProtectAlpha(t *testing.T)   { testDrawRGBAToRGBA(Subtract{}, t, true) }
func TestDrawRGBAToRGBADivideProtectAlpha(t *testing.T)     { testDrawRGBAToRGBA(Divide{}, t, true) }
func TestDrawRGBAToRGBAHueProtectAlpha(t *testing.T)        { testDrawRGBAToRGBA(Hue{}, t, true) }
func TestDrawRGBAToRGBASaturationProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(Saturation{}, t, true) }
func TestDrawRGBAToRGBAColorProtectAlpha(t *testing.T)      { testDrawRGBAToRGBA(Color{}, t, true) }
func TestDrawRGBAToRGBALuminosityProtectAlpha(t *testing.T) { testDrawRGBAToRGBA(Luminosity{}, t, true) }

func BenchmarkDrawFallbackNormal(b *testing.B)       { benchmarkDrawFallback(Normal{}, b, false) }
func BenchmarkDrawFallbackDarken(b *testing.B)       { benchmarkDrawFallback(Darken{}, b, false) }
func BenchmarkDrawFallbackMultiply(b *testing.B)     { benchmarkDrawFallback(Multiply{}, b, false) }
func BenchmarkDrawFallbackColorBurn(b *testing.B)    { benchmarkDrawFallback(ColorBurn{}, b, false) }
func BenchmarkDrawFallbackLinearBurn(b *testing.B)   { benchmarkDrawFallback(LinearBurn{}, b, false) }
func BenchmarkDrawFallbackDarkerColor(b *testing.B)  { benchmarkDrawFallback(DarkerColor{}, b, false) }
func BenchmarkDrawFallbackLighten(b *testing.B)      { benchmarkDrawFallback(Lighten{}, b, false) }
func BenchmarkDrawFallbackScreen(b *testing.B)       { benchmarkDrawFallback(Screen{}, b, false) }
func BenchmarkDrawFallbackColorDodge(b *testing.B)   { benchmarkDrawFallback(ColorDodge{}, b, false) }
func BenchmarkDrawFallbackLinearDodge(b *testing.B)  { benchmarkDrawFallback(LinearDodge{}, b, false) }
func BenchmarkDrawFallbackLighterColor(b *testing.B) { benchmarkDrawFallback(LighterColor{}, b, false) }
func BenchmarkDrawFallbackAdd(b *testing.B)          { benchmarkDrawFallback(Add{}, b, false) }
func BenchmarkDrawFallbackOverlay(b *testing.B)      { benchmarkDrawFallback(Overlay{}, b, false) }
func BenchmarkDrawFallbackSoftLight(b *testing.B)    { benchmarkDrawFallback(SoftLight{}, b, false) }
func BenchmarkDrawFallbackHardLight(b *testing.B)    { benchmarkDrawFallback(HardLight{}, b, false) }
func BenchmarkDrawFallbackLinearLight(b *testing.B)  { benchmarkDrawFallback(LinearLight{}, b, false) }
func BenchmarkDrawFallbackVividLight(b *testing.B)   { benchmarkDrawFallback(VividLight{}, b, false) }
func BenchmarkDrawFallbackPinLight(b *testing.B)     { benchmarkDrawFallback(PinLight{}, b, false) }
func BenchmarkDrawFallbackHardMix(b *testing.B)      { benchmarkDrawFallback(HardMix{}, b, false) }
func BenchmarkDrawFallbackDifference(b *testing.B)   { benchmarkDrawFallback(Difference{}, b, false) }
func BenchmarkDrawFallbackExclusion(b *testing.B)    { benchmarkDrawFallback(Exclusion{}, b, false) }
func BenchmarkDrawFallbackSubtract(b *testing.B)     { benchmarkDrawFallback(Subtract{}, b, false) }
func BenchmarkDrawFallbackDivide(b *testing.B)       { benchmarkDrawFallback(Divide{}, b, false) }
func BenchmarkDrawFallbackHue(b *testing.B)          { benchmarkDrawFallback(Hue{}, b, false) }
func BenchmarkDrawFallbackSaturation(b *testing.B)   { benchmarkDrawFallback(Saturation{}, b, false) }
func BenchmarkDrawFallbackColor(b *testing.B)        { benchmarkDrawFallback(Color{}, b, false) }
func BenchmarkDrawFallbackLuminosity(b *testing.B)   { benchmarkDrawFallback(Luminosity{}, b, false) }

func BenchmarkDrawFallbackNormalProtectAlpha(b *testing.B) { benchmarkDrawFallback(Normal{}, b, true) }
func BenchmarkDrawFallbackDarkenProtectAlpha(b *testing.B) { benchmarkDrawFallback(Darken{}, b, true) }
func BenchmarkDrawFallbackMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(Multiply{}, b, true)
}
func BenchmarkDrawFallbackColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(ColorBurn{}, b, true)
}
func BenchmarkDrawFallbackLinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(LinearBurn{}, b, true)
}
func BenchmarkDrawFallbackDarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(DarkerColor{}, b, true)
}
func BenchmarkDrawFallbackLightenProtectAlpha(b *testing.B) { benchmarkDrawFallback(Lighten{}, b, true) }
func BenchmarkDrawFallbackScreenProtectAlpha(b *testing.B)  { benchmarkDrawFallback(Screen{}, b, true) }
func BenchmarkDrawFallbackColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(ColorDodge{}, b, true)
}
func BenchmarkDrawFallbackLinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(LinearDodge{}, b, true)
}
func BenchmarkDrawFallbackLighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(LighterColor{}, b, true)
}
func BenchmarkDrawFallbackAddProtectAlpha(b *testing.B)     { benchmarkDrawFallback(Add{}, b, true) }
func BenchmarkDrawFallbackOverlayProtectAlpha(b *testing.B) { benchmarkDrawFallback(Overlay{}, b, true) }
func BenchmarkDrawFallbackSoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(SoftLight{}, b, true)
}
func BenchmarkDrawFallbackHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(HardLight{}, b, true)
}
func BenchmarkDrawFallbackLinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(LinearLight{}, b, true)
}
func BenchmarkDrawFallbackVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(VividLight{}, b, true)
}
func BenchmarkDrawFallbackPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(PinLight{}, b, true)
}
func BenchmarkDrawFallbackHardMixProtectAlpha(b *testing.B) { benchmarkDrawFallback(HardMix{}, b, true) }
func BenchmarkDrawFallbackDifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(Difference{}, b, true)
}
func BenchmarkDrawFallbackExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(Exclusion{}, b, true)
}
func BenchmarkDrawFallbackSubtractProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(Subtract{}, b, true)
}
func BenchmarkDrawFallbackDivideProtectAlpha(b *testing.B) { benchmarkDrawFallback(Divide{}, b, true) }
func BenchmarkDrawFallbackHueProtectAlpha(b *testing.B)    { benchmarkDrawFallback(Hue{}, b, true) }
func BenchmarkDrawFallbackSaturationProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(Saturation{}, b, true)
}
func BenchmarkDrawFallbackColorProtectAlpha(b *testing.B) { benchmarkDrawFallback(Color{}, b, true) }
func BenchmarkDrawFallbackLuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawFallback(Luminosity{}, b, true)
}

func BenchmarkDrawNRGBAToNRGBANormal(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(Normal{}, b, false) }
func BenchmarkDrawNRGBAToNRGBADarken(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(Darken{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAMultiply(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Multiply{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAColorBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(ColorBurn{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LinearBurn{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(DarkerColor{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALighten(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Lighten{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAScreen(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(Screen{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(ColorDodge{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LinearDodge{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LighterColor{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAAdd(b *testing.B)     { benchmarkDrawNRGBAToNRGBA(Add{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAOverlay(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Overlay{}, b, false) }
func BenchmarkDrawNRGBAToNRGBASoftLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(SoftLight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAHardLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(HardLight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LinearLight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(VividLight{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAPinLight(b *testing.B) { benchmarkDrawNRGBAToNRGBA(PinLight{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAHardMix(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(HardMix{}, b, false) }
func BenchmarkDrawNRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Difference{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAExclusion(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Exclusion{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBASubtract(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Subtract{}, b, false) }
func BenchmarkDrawNRGBAToNRGBADivide(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(Divide{}, b, false) }
func BenchmarkDrawNRGBAToNRGBAHue(b *testing.B)      { benchmarkDrawNRGBAToNRGBA(Hue{}, b, false) }
func BenchmarkDrawNRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Saturation{}, b, false)
}
func BenchmarkDrawNRGBAToNRGBAColor(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Color{}, b, false) }
func BenchmarkDrawNRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Luminosity{}, b, false)
}

func BenchmarkDrawNRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Normal{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Darken{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Multiply{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(ColorBurn{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LinearBurn{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(DarkerColor{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Lighten{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Screen{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(ColorDodge{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LinearDodge{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LighterColor{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Add{}, b, true) }
func BenchmarkDrawNRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Overlay{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(SoftLight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(HardLight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(LinearLight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(VividLight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(PinLight{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(HardMix{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Difference{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Exclusion{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Subtract{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Divide{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToNRGBA(Hue{}, b, true) }
func BenchmarkDrawNRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Saturation{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Color{}, b, true)
}
func BenchmarkDrawNRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToNRGBA(Luminosity{}, b, true)
}

func BenchmarkDrawRGBAToNRGBANormal(b *testing.B)    { benchmarkDrawRGBAToNRGBA(Normal{}, b, false) }
func BenchmarkDrawRGBAToNRGBADarken(b *testing.B)    { benchmarkDrawRGBAToNRGBA(Darken{}, b, false) }
func BenchmarkDrawRGBAToNRGBAMultiply(b *testing.B)  { benchmarkDrawRGBAToNRGBA(Multiply{}, b, false) }
func BenchmarkDrawRGBAToNRGBAColorBurn(b *testing.B) { benchmarkDrawRGBAToNRGBA(ColorBurn{}, b, false) }
func BenchmarkDrawRGBAToNRGBALinearBurn(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LinearBurn{}, b, false)
}
func BenchmarkDrawRGBAToNRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(DarkerColor{}, b, false)
}
func BenchmarkDrawRGBAToNRGBALighten(b *testing.B) { benchmarkDrawRGBAToNRGBA(Lighten{}, b, false) }
func BenchmarkDrawRGBAToNRGBAScreen(b *testing.B)  { benchmarkDrawRGBAToNRGBA(Screen{}, b, false) }
func BenchmarkDrawRGBAToNRGBAColorDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(ColorDodge{}, b, false)
}
func BenchmarkDrawRGBAToNRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LinearDodge{}, b, false)
}
func BenchmarkDrawRGBAToNRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LighterColor{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAAdd(b *testing.B)       { benchmarkDrawRGBAToNRGBA(Add{}, b, false) }
func BenchmarkDrawRGBAToNRGBAOverlay(b *testing.B)   { benchmarkDrawRGBAToNRGBA(Overlay{}, b, false) }
func BenchmarkDrawRGBAToNRGBASoftLight(b *testing.B) { benchmarkDrawRGBAToNRGBA(SoftLight{}, b, false) }
func BenchmarkDrawRGBAToNRGBAHardLight(b *testing.B) { benchmarkDrawRGBAToNRGBA(HardLight{}, b, false) }
func BenchmarkDrawRGBAToNRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LinearLight{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAVividLight(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(VividLight{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAPinLight(b *testing.B) { benchmarkDrawRGBAToNRGBA(PinLight{}, b, false) }
func BenchmarkDrawRGBAToNRGBAHardMix(b *testing.B)  { benchmarkDrawRGBAToNRGBA(HardMix{}, b, false) }
func BenchmarkDrawRGBAToNRGBADifference(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Difference{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAExclusion(b *testing.B) { benchmarkDrawRGBAToNRGBA(Exclusion{}, b, false) }
func BenchmarkDrawRGBAToNRGBASubtract(b *testing.B)  { benchmarkDrawRGBAToNRGBA(Subtract{}, b, false) }
func BenchmarkDrawRGBAToNRGBADivide(b *testing.B)    { benchmarkDrawRGBAToNRGBA(Divide{}, b, false) }
func BenchmarkDrawRGBAToNRGBAHue(b *testing.B)       { benchmarkDrawRGBAToNRGBA(Hue{}, b, false) }
func BenchmarkDrawRGBAToNRGBASaturation(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Saturation{}, b, false)
}
func BenchmarkDrawRGBAToNRGBAColor(b *testing.B) { benchmarkDrawRGBAToNRGBA(Color{}, b, false) }
func BenchmarkDrawRGBAToNRGBALuminosity(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Luminosity{}, b, false)
}

func BenchmarkDrawRGBAToNRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Normal{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Darken{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Multiply{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(ColorBurn{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LinearBurn{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(DarkerColor{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Lighten{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Screen{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(ColorDodge{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LinearDodge{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LighterColor{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawRGBAToNRGBA(Add{}, b, true) }
func BenchmarkDrawRGBAToNRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Overlay{}, b, true)
}
func BenchmarkDrawRGBAToNRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(SoftLight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(HardLight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(LinearLight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(VividLight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(PinLight{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(HardMix{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Difference{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Exclusion{}, b, true)
}
func BenchmarkDrawRGBAToNRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Subtract{}, b, true)
}
func BenchmarkDrawRGBAToNRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Divide{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawRGBAToNRGBA(Hue{}, b, true) }
func BenchmarkDrawRGBAToNRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Saturation{}, b, true)
}
func BenchmarkDrawRGBAToNRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Color{}, b, true)
}
func BenchmarkDrawRGBAToNRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToNRGBA(Luminosity{}, b, true)
}

func BenchmarkDrawNRGBAToRGBANormal(b *testing.B)    { benchmarkDrawNRGBAToRGBA(Normal{}, b, false) }
func BenchmarkDrawNRGBAToRGBADarken(b *testing.B)    { benchmarkDrawNRGBAToRGBA(Darken{}, b, false) }
func BenchmarkDrawNRGBAToRGBAMultiply(b *testing.B)  { benchmarkDrawNRGBAToRGBA(Multiply{}, b, false) }
func BenchmarkDrawNRGBAToRGBAColorBurn(b *testing.B) { benchmarkDrawNRGBAToRGBA(ColorBurn{}, b, false) }
func BenchmarkDrawNRGBAToRGBALinearBurn(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LinearBurn{}, b, false)
}
func BenchmarkDrawNRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(DarkerColor{}, b, false)
}
func BenchmarkDrawNRGBAToRGBALighten(b *testing.B) { benchmarkDrawNRGBAToRGBA(Lighten{}, b, false) }
func BenchmarkDrawNRGBAToRGBAScreen(b *testing.B)  { benchmarkDrawNRGBAToRGBA(Screen{}, b, false) }
func BenchmarkDrawNRGBAToRGBAColorDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(ColorDodge{}, b, false)
}
func BenchmarkDrawNRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LinearDodge{}, b, false)
}
func BenchmarkDrawNRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LighterColor{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAAdd(b *testing.B)       { benchmarkDrawNRGBAToRGBA(Add{}, b, false) }
func BenchmarkDrawNRGBAToRGBAOverlay(b *testing.B)   { benchmarkDrawNRGBAToRGBA(Overlay{}, b, false) }
func BenchmarkDrawNRGBAToRGBASoftLight(b *testing.B) { benchmarkDrawNRGBAToRGBA(SoftLight{}, b, false) }
func BenchmarkDrawNRGBAToRGBAHardLight(b *testing.B) { benchmarkDrawNRGBAToRGBA(HardLight{}, b, false) }
func BenchmarkDrawNRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LinearLight{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAVividLight(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(VividLight{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAPinLight(b *testing.B) { benchmarkDrawNRGBAToRGBA(PinLight{}, b, false) }
func BenchmarkDrawNRGBAToRGBAHardMix(b *testing.B)  { benchmarkDrawNRGBAToRGBA(HardMix{}, b, false) }
func BenchmarkDrawNRGBAToRGBADifference(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Difference{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAExclusion(b *testing.B) { benchmarkDrawNRGBAToRGBA(Exclusion{}, b, false) }
func BenchmarkDrawNRGBAToRGBASubtract(b *testing.B)  { benchmarkDrawNRGBAToRGBA(Subtract{}, b, false) }
func BenchmarkDrawNRGBAToRGBADivide(b *testing.B)    { benchmarkDrawNRGBAToRGBA(Divide{}, b, false) }
func BenchmarkDrawNRGBAToRGBAHue(b *testing.B)       { benchmarkDrawNRGBAToRGBA(Hue{}, b, false) }
func BenchmarkDrawNRGBAToRGBASaturation(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Saturation{}, b, false)
}
func BenchmarkDrawNRGBAToRGBAColor(b *testing.B) { benchmarkDrawNRGBAToRGBA(Color{}, b, false) }
func BenchmarkDrawNRGBAToRGBALuminosity(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Luminosity{}, b, false)
}

func BenchmarkDrawNRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Normal{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Darken{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Multiply{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(ColorBurn{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LinearBurn{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(DarkerColor{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Lighten{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Screen{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(ColorDodge{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LinearDodge{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LighterColor{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToRGBA(Add{}, b, true) }
func BenchmarkDrawNRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Overlay{}, b, true)
}
func BenchmarkDrawNRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(SoftLight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(HardLight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(LinearLight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(VividLight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(PinLight{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(HardMix{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Difference{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Exclusion{}, b, true)
}
func BenchmarkDrawNRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Subtract{}, b, true)
}
func BenchmarkDrawNRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Divide{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawNRGBAToRGBA(Hue{}, b, true) }
func BenchmarkDrawNRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Saturation{}, b, true)
}
func BenchmarkDrawNRGBAToRGBAColorProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Color{}, b, true)
}
func BenchmarkDrawNRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawNRGBAToRGBA(Luminosity{}, b, true)
}

func BenchmarkDrawRGBAToRGBANormal(b *testing.B)     { benchmarkDrawRGBAToRGBA(Normal{}, b, false) }
func BenchmarkDrawRGBAToRGBADarken(b *testing.B)     { benchmarkDrawRGBAToRGBA(Darken{}, b, false) }
func BenchmarkDrawRGBAToRGBAMultiply(b *testing.B)   { benchmarkDrawRGBAToRGBA(Multiply{}, b, false) }
func BenchmarkDrawRGBAToRGBAColorBurn(b *testing.B)  { benchmarkDrawRGBAToRGBA(ColorBurn{}, b, false) }
func BenchmarkDrawRGBAToRGBALinearBurn(b *testing.B) { benchmarkDrawRGBAToRGBA(LinearBurn{}, b, false) }
func BenchmarkDrawRGBAToRGBADarkerColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(DarkerColor{}, b, false)
}
func BenchmarkDrawRGBAToRGBALighten(b *testing.B)    { benchmarkDrawRGBAToRGBA(Lighten{}, b, false) }
func BenchmarkDrawRGBAToRGBAScreen(b *testing.B)     { benchmarkDrawRGBAToRGBA(Screen{}, b, false) }
func BenchmarkDrawRGBAToRGBAColorDodge(b *testing.B) { benchmarkDrawRGBAToRGBA(ColorDodge{}, b, false) }
func BenchmarkDrawRGBAToRGBALinearDodge(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LinearDodge{}, b, false)
}
func BenchmarkDrawRGBAToRGBALighterColor(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LighterColor{}, b, false)
}
func BenchmarkDrawRGBAToRGBAAdd(b *testing.B)       { benchmarkDrawRGBAToRGBA(Add{}, b, false) }
func BenchmarkDrawRGBAToRGBAOverlay(b *testing.B)   { benchmarkDrawRGBAToRGBA(Overlay{}, b, false) }
func BenchmarkDrawRGBAToRGBASoftLight(b *testing.B) { benchmarkDrawRGBAToRGBA(SoftLight{}, b, false) }
func BenchmarkDrawRGBAToRGBAHardLight(b *testing.B) { benchmarkDrawRGBAToRGBA(HardLight{}, b, false) }
func BenchmarkDrawRGBAToRGBALinearLight(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LinearLight{}, b, false)
}
func BenchmarkDrawRGBAToRGBAVividLight(b *testing.B) { benchmarkDrawRGBAToRGBA(VividLight{}, b, false) }
func BenchmarkDrawRGBAToRGBAPinLight(b *testing.B)   { benchmarkDrawRGBAToRGBA(PinLight{}, b, false) }
func BenchmarkDrawRGBAToRGBAHardMix(b *testing.B)    { benchmarkDrawRGBAToRGBA(HardMix{}, b, false) }
func BenchmarkDrawRGBAToRGBADifference(b *testing.B) { benchmarkDrawRGBAToRGBA(Difference{}, b, false) }
func BenchmarkDrawRGBAToRGBAExclusion(b *testing.B)  { benchmarkDrawRGBAToRGBA(Exclusion{}, b, false) }
func BenchmarkDrawRGBAToRGBASubtract(b *testing.B)   { benchmarkDrawRGBAToRGBA(Subtract{}, b, false) }
func BenchmarkDrawRGBAToRGBADivide(b *testing.B)     { benchmarkDrawRGBAToRGBA(Divide{}, b, false) }
func BenchmarkDrawRGBAToRGBAHue(b *testing.B)        { benchmarkDrawRGBAToRGBA(Hue{}, b, false) }
func BenchmarkDrawRGBAToRGBASaturation(b *testing.B) { benchmarkDrawRGBAToRGBA(Saturation{}, b, false) }
func BenchmarkDrawRGBAToRGBAColor(b *testing.B)      { benchmarkDrawRGBAToRGBA(Color{}, b, false) }
func BenchmarkDrawRGBAToRGBALuminosity(b *testing.B) { benchmarkDrawRGBAToRGBA(Luminosity{}, b, false) }

func BenchmarkDrawRGBAToRGBANormalProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Normal{}, b, true)
}
func BenchmarkDrawRGBAToRGBADarkenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Darken{}, b, true)
}
func BenchmarkDrawRGBAToRGBAMultiplyProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Multiply{}, b, true)
}
func BenchmarkDrawRGBAToRGBAColorBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(ColorBurn{}, b, true)
}
func BenchmarkDrawRGBAToRGBALinearBurnProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LinearBurn{}, b, true)
}
func BenchmarkDrawRGBAToRGBADarkerColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(DarkerColor{}, b, true)
}
func BenchmarkDrawRGBAToRGBALightenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Lighten{}, b, true)
}
func BenchmarkDrawRGBAToRGBAScreenProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Screen{}, b, true)
}
func BenchmarkDrawRGBAToRGBAColorDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(ColorDodge{}, b, true)
}
func BenchmarkDrawRGBAToRGBALinearDodgeProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LinearDodge{}, b, true)
}
func BenchmarkDrawRGBAToRGBALighterColorProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LighterColor{}, b, true)
}
func BenchmarkDrawRGBAToRGBAAddProtectAlpha(b *testing.B) { benchmarkDrawRGBAToRGBA(Add{}, b, true) }
func BenchmarkDrawRGBAToRGBAOverlayProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Overlay{}, b, true)
}
func BenchmarkDrawRGBAToRGBASoftLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(SoftLight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAHardLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(HardLight{}, b, true)
}
func BenchmarkDrawRGBAToRGBALinearLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(LinearLight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAVividLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(VividLight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAPinLightProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(PinLight{}, b, true)
}
func BenchmarkDrawRGBAToRGBAHardMixProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(HardMix{}, b, true)
}
func BenchmarkDrawRGBAToRGBADifferenceProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Difference{}, b, true)
}
func BenchmarkDrawRGBAToRGBAExclusionProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Exclusion{}, b, true)
}
func BenchmarkDrawRGBAToRGBASubtractProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Subtract{}, b, true)
}
func BenchmarkDrawRGBAToRGBADivideProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Divide{}, b, true)
}
func BenchmarkDrawRGBAToRGBAHueProtectAlpha(b *testing.B) { benchmarkDrawRGBAToRGBA(Hue{}, b, true) }
func BenchmarkDrawRGBAToRGBASaturationProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Saturation{}, b, true)
}
func BenchmarkDrawRGBAToRGBAColorProtectAlpha(b *testing.B) { benchmarkDrawRGBAToRGBA(Color{}, b, true) }
func BenchmarkDrawRGBAToRGBALuminosityProtectAlpha(b *testing.B) {
	benchmarkDrawRGBAToRGBA(Luminosity{}, b, true)
}
