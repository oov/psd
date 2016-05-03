package blend

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"
	"testing"
)

type drawer interface {
	Drawer
	drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform)
	drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform)
	drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform)
	drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform)
	drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point)
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func loadImages(path1 string, path2 string) (image.Image, image.Image, error) {
	img1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	img2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}
	return img1, img2, nil
}

func saveImage(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func verify(a image.Image, b image.Image) (float64, error) {
	if a.Bounds().String() != a.Bounds().String() {
		return 0, fmt.Errorf("image sizes are not same\nA: %v / B: %v", a.Bounds(), b.Bounds())
	}

	bounds := a.Bounds()
	er := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ar, ag, ab, aa := a.At(x, y).RGBA()
			br, bg, bb, ba := b.At(x, y).RGBA()
			if math.Abs(float64(ar)-float64(br)) >= 0x0200 ||
				math.Abs(float64(ag)-float64(bg)) >= 0x0200 ||
				math.Abs(float64(ab)-float64(bb)) >= 0x0200 ||
				math.Abs(float64(aa)-float64(ba)) >= 0x0200 {
				er++
			}
		}
	}
	return float64(er) / float64(bounds.Dx()*bounds.Dy()), nil
}

func testDrawFallback(d drawer, t *testing.T) {
	name := "DrawFallback"
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		t.Fatalf("%v", err)
	}
	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	d.drawFallback(img, img2.Bounds(), img2, image.Pt(0, 0), nil, image.Pt(0, 0))

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	errorRate, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("ErrorRate: %3.2f%%", errorRate*100)
}

func TestDrawFallbackNormal(t *testing.T)       { testDrawFallback(Normal{}, t) }
func TestDrawFallbackDarken(t *testing.T)       { testDrawFallback(Darken{}, t) }
func TestDrawFallbackMultiply(t *testing.T)     { testDrawFallback(Multiply{}, t) }
func TestDrawFallbackColorBurn(t *testing.T)    { testDrawFallback(ColorBurn{}, t) }
func TestDrawFallbackLinearBurn(t *testing.T)   { testDrawFallback(LinearBurn{}, t) }
func TestDrawFallbackDarkerColor(t *testing.T)  { testDrawFallback(DarkerColor{}, t) }
func TestDrawFallbackLighten(t *testing.T)      { testDrawFallback(Lighten{}, t) }
func TestDrawFallbackScreen(t *testing.T)       { testDrawFallback(Screen{}, t) }
func TestDrawFallbackColorDodge(t *testing.T)   { testDrawFallback(ColorDodge{}, t) }
func TestDrawFallbackLinearDodge(t *testing.T)  { testDrawFallback(LinearDodge{}, t) }
func TestDrawFallbackLighterColor(t *testing.T) { testDrawFallback(LighterColor{}, t) }
func TestDrawFallbackOverlay(t *testing.T)      { testDrawFallback(Overlay{}, t) }
func TestDrawFallbackSoftLight(t *testing.T)    { testDrawFallback(SoftLight{}, t) }
func TestDrawFallbackHardLight(t *testing.T)    { testDrawFallback(HardLight{}, t) }
func TestDrawFallbackVividLight(t *testing.T)   { testDrawFallback(VividLight{}, t) }
func TestDrawFallbackLinearLight(t *testing.T)  { testDrawFallback(LinearLight{}, t) }
func TestDrawFallbackPinLight(t *testing.T)     { testDrawFallback(PinLight{}, t) }
func TestDrawFallbackHardMix(t *testing.T)      { testDrawFallback(HardMix{}, t) }
func TestDrawFallbackDifference(t *testing.T)   { testDrawFallback(Difference{}, t) }
func TestDrawFallbackExclusion(t *testing.T)    { testDrawFallback(Exclusion{}, t) }
func TestDrawFallbackSubtract(t *testing.T)     { testDrawFallback(Subtract{}, t) }
func TestDrawFallbackDivide(t *testing.T)       { testDrawFallback(Divide{}, t) }
func TestDrawFallbackHue(t *testing.T)          { testDrawFallback(Hue{}, t) }
func TestDrawFallbackSaturation(t *testing.T)   { testDrawFallback(Saturation{}, t) }
func TestDrawFallbackColor(t *testing.T)        { testDrawFallback(Color{}, t) }
func TestDrawFallbackLuminosity(t *testing.T)   { testDrawFallback(Luminosity{}, t) }

func testDrawNRGBAToNRGBA(d drawer, t *testing.T) {
	name := "DrawNRGBAToNRGBA"
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		t.Fatalf("%v", err)
	}
	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	d.drawNRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	errorRate, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("ErrorRate: %3.2f%%", errorRate*100)
}

func TestDrawNRGBAToNRGBANormal(t *testing.T)       { testDrawNRGBAToNRGBA(Normal{}, t) }
func TestDrawNRGBAToNRGBADarken(t *testing.T)       { testDrawNRGBAToNRGBA(Darken{}, t) }
func TestDrawNRGBAToNRGBAMultiply(t *testing.T)     { testDrawNRGBAToNRGBA(Multiply{}, t) }
func TestDrawNRGBAToNRGBAColorBurn(t *testing.T)    { testDrawNRGBAToNRGBA(ColorBurn{}, t) }
func TestDrawNRGBAToNRGBALinearBurn(t *testing.T)   { testDrawNRGBAToNRGBA(LinearBurn{}, t) }
func TestDrawNRGBAToNRGBADarkerColor(t *testing.T)  { testDrawNRGBAToNRGBA(DarkerColor{}, t) }
func TestDrawNRGBAToNRGBALighten(t *testing.T)      { testDrawNRGBAToNRGBA(Lighten{}, t) }
func TestDrawNRGBAToNRGBAScreen(t *testing.T)       { testDrawNRGBAToNRGBA(Screen{}, t) }
func TestDrawNRGBAToNRGBAColorDodge(t *testing.T)   { testDrawNRGBAToNRGBA(ColorDodge{}, t) }
func TestDrawNRGBAToNRGBALinearDodge(t *testing.T)  { testDrawNRGBAToNRGBA(LinearDodge{}, t) }
func TestDrawNRGBAToNRGBALighterColor(t *testing.T) { testDrawNRGBAToNRGBA(LighterColor{}, t) }
func TestDrawNRGBAToNRGBAOverlay(t *testing.T)      { testDrawNRGBAToNRGBA(Overlay{}, t) }
func TestDrawNRGBAToNRGBASoftLight(t *testing.T)    { testDrawNRGBAToNRGBA(SoftLight{}, t) }
func TestDrawNRGBAToNRGBAHardLight(t *testing.T)    { testDrawNRGBAToNRGBA(HardLight{}, t) }
func TestDrawNRGBAToNRGBAVividLight(t *testing.T)   { testDrawNRGBAToNRGBA(VividLight{}, t) }
func TestDrawNRGBAToNRGBALinearLight(t *testing.T)  { testDrawNRGBAToNRGBA(LinearLight{}, t) }
func TestDrawNRGBAToNRGBAPinLight(t *testing.T)     { testDrawNRGBAToNRGBA(PinLight{}, t) }
func TestDrawNRGBAToNRGBAHardMix(t *testing.T)      { testDrawNRGBAToNRGBA(HardMix{}, t) }
func TestDrawNRGBAToNRGBADifference(t *testing.T)   { testDrawNRGBAToNRGBA(Difference{}, t) }
func TestDrawNRGBAToNRGBAExclusion(t *testing.T)    { testDrawNRGBAToNRGBA(Exclusion{}, t) }
func TestDrawNRGBAToNRGBASubtract(t *testing.T)     { testDrawNRGBAToNRGBA(Subtract{}, t) }
func TestDrawNRGBAToNRGBADivide(t *testing.T)       { testDrawNRGBAToNRGBA(Divide{}, t) }
func TestDrawNRGBAToNRGBAHue(t *testing.T)          { testDrawNRGBAToNRGBA(Hue{}, t) }
func TestDrawNRGBAToNRGBASaturation(t *testing.T)   { testDrawNRGBAToNRGBA(Saturation{}, t) }
func TestDrawNRGBAToNRGBAColor(t *testing.T)        { testDrawNRGBAToNRGBA(Color{}, t) }
func TestDrawNRGBAToNRGBALuminosity(t *testing.T)   { testDrawNRGBAToNRGBA(Luminosity{}, t) }

func testDrawRGBAToNRGBA(d drawer, t *testing.T) {
	name := "DrawRGBAToNRGBA"
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		t.Fatalf("%v", err)
	}
	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	d.drawRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	errorRate, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("ErrorRate: %3.2f%%", errorRate*100)
}

func TestDrawRGBAToNRGBANormal(t *testing.T)       { testDrawRGBAToNRGBA(Normal{}, t) }
func TestDrawRGBAToNRGBADarken(t *testing.T)       { testDrawRGBAToNRGBA(Darken{}, t) }
func TestDrawRGBAToNRGBAMultiply(t *testing.T)     { testDrawRGBAToNRGBA(Multiply{}, t) }
func TestDrawRGBAToNRGBAColorBurn(t *testing.T)    { testDrawRGBAToNRGBA(ColorBurn{}, t) }
func TestDrawRGBAToNRGBALinearBurn(t *testing.T)   { testDrawRGBAToNRGBA(LinearBurn{}, t) }
func TestDrawRGBAToNRGBADarkerColor(t *testing.T)  { testDrawRGBAToNRGBA(DarkerColor{}, t) }
func TestDrawRGBAToNRGBALighten(t *testing.T)      { testDrawRGBAToNRGBA(Lighten{}, t) }
func TestDrawRGBAToNRGBAScreen(t *testing.T)       { testDrawRGBAToNRGBA(Screen{}, t) }
func TestDrawRGBAToNRGBAColorDodge(t *testing.T)   { testDrawRGBAToNRGBA(ColorDodge{}, t) }
func TestDrawRGBAToNRGBALinearDodge(t *testing.T)  { testDrawRGBAToNRGBA(LinearDodge{}, t) }
func TestDrawRGBAToNRGBALighterColor(t *testing.T) { testDrawRGBAToNRGBA(LighterColor{}, t) }
func TestDrawRGBAToNRGBAOverlay(t *testing.T)      { testDrawRGBAToNRGBA(Overlay{}, t) }
func TestDrawRGBAToNRGBASoftLight(t *testing.T)    { testDrawRGBAToNRGBA(SoftLight{}, t) }
func TestDrawRGBAToNRGBAHardLight(t *testing.T)    { testDrawRGBAToNRGBA(HardLight{}, t) }
func TestDrawRGBAToNRGBAVividLight(t *testing.T)   { testDrawRGBAToNRGBA(VividLight{}, t) }
func TestDrawRGBAToNRGBALinearLight(t *testing.T)  { testDrawRGBAToNRGBA(LinearLight{}, t) }
func TestDrawRGBAToNRGBAPinLight(t *testing.T)     { testDrawRGBAToNRGBA(PinLight{}, t) }
func TestDrawRGBAToNRGBAHardMix(t *testing.T)      { testDrawRGBAToNRGBA(HardMix{}, t) }
func TestDrawRGBAToNRGBADifference(t *testing.T)   { testDrawRGBAToNRGBA(Difference{}, t) }
func TestDrawRGBAToNRGBAExclusion(t *testing.T)    { testDrawRGBAToNRGBA(Exclusion{}, t) }
func TestDrawRGBAToNRGBASubtract(t *testing.T)     { testDrawRGBAToNRGBA(Subtract{}, t) }
func TestDrawRGBAToNRGBADivide(t *testing.T)       { testDrawRGBAToNRGBA(Divide{}, t) }
func TestDrawRGBAToNRGBAHue(t *testing.T)          { testDrawRGBAToNRGBA(Hue{}, t) }
func TestDrawRGBAToNRGBASaturation(t *testing.T)   { testDrawRGBAToNRGBA(Saturation{}, t) }
func TestDrawRGBAToNRGBAColor(t *testing.T)        { testDrawRGBAToNRGBA(Color{}, t) }
func TestDrawRGBAToNRGBALuminosity(t *testing.T)   { testDrawRGBAToNRGBA(Luminosity{}, t) }

func testDrawNRGBAToRGBA(d drawer, t *testing.T) {
	name := "DrawNRGBAToRGBA"
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		t.Fatalf("%v", err)
	}
	img := image.NewRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	d.drawNRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	errorRate, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("ErrorRate: %3.2f%%", errorRate*100)
}

func TestDrawNRGBAToRGBANormal(t *testing.T)       { testDrawNRGBAToRGBA(Normal{}, t) }
func TestDrawNRGBAToRGBADarken(t *testing.T)       { testDrawNRGBAToRGBA(Darken{}, t) }
func TestDrawNRGBAToRGBAMultiply(t *testing.T)     { testDrawNRGBAToRGBA(Multiply{}, t) }
func TestDrawNRGBAToRGBAColorBurn(t *testing.T)    { testDrawNRGBAToRGBA(ColorBurn{}, t) }
func TestDrawNRGBAToRGBALinearBurn(t *testing.T)   { testDrawNRGBAToRGBA(LinearBurn{}, t) }
func TestDrawNRGBAToRGBADarkerColor(t *testing.T)  { testDrawNRGBAToRGBA(DarkerColor{}, t) }
func TestDrawNRGBAToRGBALighten(t *testing.T)      { testDrawNRGBAToRGBA(Lighten{}, t) }
func TestDrawNRGBAToRGBAScreen(t *testing.T)       { testDrawNRGBAToRGBA(Screen{}, t) }
func TestDrawNRGBAToRGBAColorDodge(t *testing.T)   { testDrawNRGBAToRGBA(ColorDodge{}, t) }
func TestDrawNRGBAToRGBALinearDodge(t *testing.T)  { testDrawNRGBAToRGBA(LinearDodge{}, t) }
func TestDrawNRGBAToRGBALighterColor(t *testing.T) { testDrawNRGBAToRGBA(LighterColor{}, t) }
func TestDrawNRGBAToRGBAOverlay(t *testing.T)      { testDrawNRGBAToRGBA(Overlay{}, t) }
func TestDrawNRGBAToRGBASoftLight(t *testing.T)    { testDrawNRGBAToRGBA(SoftLight{}, t) }
func TestDrawNRGBAToRGBAHardLight(t *testing.T)    { testDrawNRGBAToRGBA(HardLight{}, t) }
func TestDrawNRGBAToRGBAVividLight(t *testing.T)   { testDrawNRGBAToRGBA(VividLight{}, t) }
func TestDrawNRGBAToRGBALinearLight(t *testing.T)  { testDrawNRGBAToRGBA(LinearLight{}, t) }
func TestDrawNRGBAToRGBAPinLight(t *testing.T)     { testDrawNRGBAToRGBA(PinLight{}, t) }
func TestDrawNRGBAToRGBAHardMix(t *testing.T)      { testDrawNRGBAToRGBA(HardMix{}, t) }
func TestDrawNRGBAToRGBADifference(t *testing.T)   { testDrawNRGBAToRGBA(Difference{}, t) }
func TestDrawNRGBAToRGBAExclusion(t *testing.T)    { testDrawNRGBAToRGBA(Exclusion{}, t) }
func TestDrawNRGBAToRGBASubtract(t *testing.T)     { testDrawNRGBAToRGBA(Subtract{}, t) }
func TestDrawNRGBAToRGBADivide(t *testing.T)       { testDrawNRGBAToRGBA(Divide{}, t) }
func TestDrawNRGBAToRGBAHue(t *testing.T)          { testDrawNRGBAToRGBA(Hue{}, t) }
func TestDrawNRGBAToRGBASaturation(t *testing.T)   { testDrawNRGBAToRGBA(Saturation{}, t) }
func TestDrawNRGBAToRGBAColor(t *testing.T)        { testDrawNRGBAToRGBA(Color{}, t) }
func TestDrawNRGBAToRGBALuminosity(t *testing.T)   { testDrawNRGBAToRGBA(Luminosity{}, t) }

func testDrawRGBAToRGBA(d drawer, t *testing.T) {
	name := "DrawRGBAToRGBA"
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		t.Fatalf("%v", err)
	}
	img := image.NewRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	d.drawRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	errorRate, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("ErrorRate: %3.2f%%", errorRate*100)
}

func TestDrawRGBAToRGBANormal(t *testing.T)       { testDrawRGBAToRGBA(Normal{}, t) }
func TestDrawRGBAToRGBADarken(t *testing.T)       { testDrawRGBAToRGBA(Darken{}, t) }
func TestDrawRGBAToRGBAMultiply(t *testing.T)     { testDrawRGBAToRGBA(Multiply{}, t) }
func TestDrawRGBAToRGBAColorBurn(t *testing.T)    { testDrawRGBAToRGBA(ColorBurn{}, t) }
func TestDrawRGBAToRGBALinearBurn(t *testing.T)   { testDrawRGBAToRGBA(LinearBurn{}, t) }
func TestDrawRGBAToRGBADarkerColor(t *testing.T)  { testDrawRGBAToRGBA(DarkerColor{}, t) }
func TestDrawRGBAToRGBALighten(t *testing.T)      { testDrawRGBAToRGBA(Lighten{}, t) }
func TestDrawRGBAToRGBAScreen(t *testing.T)       { testDrawRGBAToRGBA(Screen{}, t) }
func TestDrawRGBAToRGBAColorDodge(t *testing.T)   { testDrawRGBAToRGBA(ColorDodge{}, t) }
func TestDrawRGBAToRGBALinearDodge(t *testing.T)  { testDrawRGBAToRGBA(LinearDodge{}, t) }
func TestDrawRGBAToRGBALighterColor(t *testing.T) { testDrawRGBAToRGBA(LighterColor{}, t) }
func TestDrawRGBAToRGBAOverlay(t *testing.T)      { testDrawRGBAToRGBA(Overlay{}, t) }
func TestDrawRGBAToRGBASoftLight(t *testing.T)    { testDrawRGBAToRGBA(SoftLight{}, t) }
func TestDrawRGBAToRGBAHardLight(t *testing.T)    { testDrawRGBAToRGBA(HardLight{}, t) }
func TestDrawRGBAToRGBAVividLight(t *testing.T)   { testDrawRGBAToRGBA(VividLight{}, t) }
func TestDrawRGBAToRGBALinearLight(t *testing.T)  { testDrawRGBAToRGBA(LinearLight{}, t) }
func TestDrawRGBAToRGBAPinLight(t *testing.T)     { testDrawRGBAToRGBA(PinLight{}, t) }
func TestDrawRGBAToRGBAHardMix(t *testing.T)      { testDrawRGBAToRGBA(HardMix{}, t) }
func TestDrawRGBAToRGBADifference(t *testing.T)   { testDrawRGBAToRGBA(Difference{}, t) }
func TestDrawRGBAToRGBAExclusion(t *testing.T)    { testDrawRGBAToRGBA(Exclusion{}, t) }
func TestDrawRGBAToRGBASubtract(t *testing.T)     { testDrawRGBAToRGBA(Subtract{}, t) }
func TestDrawRGBAToRGBADivide(t *testing.T)       { testDrawRGBAToRGBA(Divide{}, t) }
func TestDrawRGBAToRGBAHue(t *testing.T)          { testDrawRGBAToRGBA(Hue{}, t) }
func TestDrawRGBAToRGBASaturation(t *testing.T)   { testDrawRGBAToRGBA(Saturation{}, t) }
func TestDrawRGBAToRGBAColor(t *testing.T)        { testDrawRGBAToRGBA(Color{}, t) }
func TestDrawRGBAToRGBALuminosity(t *testing.T)   { testDrawRGBAToRGBA(Luminosity{}, t) }

func benchmarkDrawFallback(d drawer, b *testing.B) {
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		b.Fatalf("%v", err)
	}
	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawFallback(img, img2.Bounds(), img2, image.Pt(0, 0), nil, image.Pt(0, 0))
	}
}

func BenchmarkDrawFallbackNormal(b *testing.B)       { benchmarkDrawFallback(Normal{}, b) }
func BenchmarkDrawFallbackDarken(b *testing.B)       { benchmarkDrawFallback(Darken{}, b) }
func BenchmarkDrawFallbackMultiply(b *testing.B)     { benchmarkDrawFallback(Multiply{}, b) }
func BenchmarkDrawFallbackColorBurn(b *testing.B)    { benchmarkDrawFallback(ColorBurn{}, b) }
func BenchmarkDrawFallbackLinearBurn(b *testing.B)   { benchmarkDrawFallback(LinearBurn{}, b) }
func BenchmarkDrawFallbackDarkerColor(b *testing.B)  { benchmarkDrawFallback(DarkerColor{}, b) }
func BenchmarkDrawFallbackLighten(b *testing.B)      { benchmarkDrawFallback(Lighten{}, b) }
func BenchmarkDrawFallbackScreen(b *testing.B)       { benchmarkDrawFallback(Screen{}, b) }
func BenchmarkDrawFallbackColorDodge(b *testing.B)   { benchmarkDrawFallback(ColorDodge{}, b) }
func BenchmarkDrawFallbackLinearDodge(b *testing.B)  { benchmarkDrawFallback(LinearDodge{}, b) }
func BenchmarkDrawFallbackLighterColor(b *testing.B) { benchmarkDrawFallback(LighterColor{}, b) }
func BenchmarkDrawFallbackOverlay(b *testing.B)      { benchmarkDrawFallback(Overlay{}, b) }
func BenchmarkDrawFallbackSoftLight(b *testing.B)    { benchmarkDrawFallback(SoftLight{}, b) }
func BenchmarkDrawFallbackHardLight(b *testing.B)    { benchmarkDrawFallback(HardLight{}, b) }
func BenchmarkDrawFallbackVividLight(b *testing.B)   { benchmarkDrawFallback(VividLight{}, b) }
func BenchmarkDrawFallbackLinearLight(b *testing.B)  { benchmarkDrawFallback(LinearLight{}, b) }
func BenchmarkDrawFallbackPinLight(b *testing.B)     { benchmarkDrawFallback(PinLight{}, b) }
func BenchmarkDrawFallbackHardMix(b *testing.B)      { benchmarkDrawFallback(HardMix{}, b) }
func BenchmarkDrawFallbackDifference(b *testing.B)   { benchmarkDrawFallback(Difference{}, b) }
func BenchmarkDrawFallbackExclusion(b *testing.B)    { benchmarkDrawFallback(Exclusion{}, b) }
func BenchmarkDrawFallbackSubtract(b *testing.B)     { benchmarkDrawFallback(Subtract{}, b) }
func BenchmarkDrawFallbackDivide(b *testing.B)       { benchmarkDrawFallback(Divide{}, b) }
func BenchmarkDrawFallbackHue(b *testing.B)          { benchmarkDrawFallback(Hue{}, b) }
func BenchmarkDrawFallbackSaturation(b *testing.B)   { benchmarkDrawFallback(Saturation{}, b) }
func BenchmarkDrawFallbackColor(b *testing.B)        { benchmarkDrawFallback(Color{}, b) }
func BenchmarkDrawFallbackLuminosity(b *testing.B)   { benchmarkDrawFallback(Luminosity{}, b) }

func benchmarkDrawNRGBAToNRGBA(d drawer, b *testing.B) {
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		b.Fatalf("%v", err)
	}
	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawNRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func BenchmarkDrawNRGBAToNRGBANormal(b *testing.B)       { benchmarkDrawNRGBAToNRGBA(Normal{}, b) }
func BenchmarkDrawNRGBAToNRGBADarken(b *testing.B)       { benchmarkDrawNRGBAToNRGBA(Darken{}, b) }
func BenchmarkDrawNRGBAToNRGBAMultiply(b *testing.B)     { benchmarkDrawNRGBAToNRGBA(Multiply{}, b) }
func BenchmarkDrawNRGBAToNRGBAColorBurn(b *testing.B)    { benchmarkDrawNRGBAToNRGBA(ColorBurn{}, b) }
func BenchmarkDrawNRGBAToNRGBALinearBurn(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(LinearBurn{}, b) }
func BenchmarkDrawNRGBAToNRGBADarkerColor(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(DarkerColor{}, b) }
func BenchmarkDrawNRGBAToNRGBALighten(b *testing.B)      { benchmarkDrawNRGBAToNRGBA(Lighten{}, b) }
func BenchmarkDrawNRGBAToNRGBAScreen(b *testing.B)       { benchmarkDrawNRGBAToNRGBA(Screen{}, b) }
func BenchmarkDrawNRGBAToNRGBAColorDodge(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(ColorDodge{}, b) }
func BenchmarkDrawNRGBAToNRGBALinearDodge(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(LinearDodge{}, b) }
func BenchmarkDrawNRGBAToNRGBALighterColor(b *testing.B) { benchmarkDrawNRGBAToNRGBA(LighterColor{}, b) }
func BenchmarkDrawNRGBAToNRGBAOverlay(b *testing.B)      { benchmarkDrawNRGBAToNRGBA(Overlay{}, b) }
func BenchmarkDrawNRGBAToNRGBASoftLight(b *testing.B)    { benchmarkDrawNRGBAToNRGBA(SoftLight{}, b) }
func BenchmarkDrawNRGBAToNRGBAHardLight(b *testing.B)    { benchmarkDrawNRGBAToNRGBA(HardLight{}, b) }
func BenchmarkDrawNRGBAToNRGBAVividLight(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(VividLight{}, b) }
func BenchmarkDrawNRGBAToNRGBALinearLight(b *testing.B)  { benchmarkDrawNRGBAToNRGBA(LinearLight{}, b) }
func BenchmarkDrawNRGBAToNRGBAPinLight(b *testing.B)     { benchmarkDrawNRGBAToNRGBA(PinLight{}, b) }
func BenchmarkDrawNRGBAToNRGBAHardMix(b *testing.B)      { benchmarkDrawNRGBAToNRGBA(HardMix{}, b) }
func BenchmarkDrawNRGBAToNRGBADifference(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(Difference{}, b) }
func BenchmarkDrawNRGBAToNRGBAExclusion(b *testing.B)    { benchmarkDrawNRGBAToNRGBA(Exclusion{}, b) }
func BenchmarkDrawNRGBAToNRGBASubtract(b *testing.B)     { benchmarkDrawNRGBAToNRGBA(Subtract{}, b) }
func BenchmarkDrawNRGBAToNRGBADivide(b *testing.B)       { benchmarkDrawNRGBAToNRGBA(Divide{}, b) }
func BenchmarkDrawNRGBAToNRGBAHue(b *testing.B)          { benchmarkDrawNRGBAToNRGBA(Hue{}, b) }
func BenchmarkDrawNRGBAToNRGBASaturation(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(Saturation{}, b) }
func BenchmarkDrawNRGBAToNRGBAColor(b *testing.B)        { benchmarkDrawNRGBAToNRGBA(Color{}, b) }
func BenchmarkDrawNRGBAToNRGBALuminosity(b *testing.B)   { benchmarkDrawNRGBAToNRGBA(Luminosity{}, b) }

func benchmarkDrawRGBAToNRGBA(d drawer, b *testing.B) {
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		b.Fatalf("%v", err)
	}
	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func BenchmarkDrawRGBAToNRGBANormal(b *testing.B)       { benchmarkDrawRGBAToNRGBA(Normal{}, b) }
func BenchmarkDrawRGBAToNRGBADarken(b *testing.B)       { benchmarkDrawRGBAToNRGBA(Darken{}, b) }
func BenchmarkDrawRGBAToNRGBAMultiply(b *testing.B)     { benchmarkDrawRGBAToNRGBA(Multiply{}, b) }
func BenchmarkDrawRGBAToNRGBAColorBurn(b *testing.B)    { benchmarkDrawRGBAToNRGBA(ColorBurn{}, b) }
func BenchmarkDrawRGBAToNRGBALinearBurn(b *testing.B)   { benchmarkDrawRGBAToNRGBA(LinearBurn{}, b) }
func BenchmarkDrawRGBAToNRGBADarkerColor(b *testing.B)  { benchmarkDrawRGBAToNRGBA(DarkerColor{}, b) }
func BenchmarkDrawRGBAToNRGBALighten(b *testing.B)      { benchmarkDrawRGBAToNRGBA(Lighten{}, b) }
func BenchmarkDrawRGBAToNRGBAScreen(b *testing.B)       { benchmarkDrawRGBAToNRGBA(Screen{}, b) }
func BenchmarkDrawRGBAToNRGBAColorDodge(b *testing.B)   { benchmarkDrawRGBAToNRGBA(ColorDodge{}, b) }
func BenchmarkDrawRGBAToNRGBALinearDodge(b *testing.B)  { benchmarkDrawRGBAToNRGBA(LinearDodge{}, b) }
func BenchmarkDrawRGBAToNRGBALighterColor(b *testing.B) { benchmarkDrawRGBAToNRGBA(LighterColor{}, b) }
func BenchmarkDrawRGBAToNRGBAOverlay(b *testing.B)      { benchmarkDrawRGBAToNRGBA(Overlay{}, b) }
func BenchmarkDrawRGBAToNRGBASoftLight(b *testing.B)    { benchmarkDrawRGBAToNRGBA(SoftLight{}, b) }
func BenchmarkDrawRGBAToNRGBAHardLight(b *testing.B)    { benchmarkDrawRGBAToNRGBA(HardLight{}, b) }
func BenchmarkDrawRGBAToNRGBAVividLight(b *testing.B)   { benchmarkDrawRGBAToNRGBA(VividLight{}, b) }
func BenchmarkDrawRGBAToNRGBALinearLight(b *testing.B)  { benchmarkDrawRGBAToNRGBA(LinearLight{}, b) }
func BenchmarkDrawRGBAToNRGBAPinLight(b *testing.B)     { benchmarkDrawRGBAToNRGBA(PinLight{}, b) }
func BenchmarkDrawRGBAToNRGBAHardMix(b *testing.B)      { benchmarkDrawRGBAToNRGBA(HardMix{}, b) }
func BenchmarkDrawRGBAToNRGBADifference(b *testing.B)   { benchmarkDrawRGBAToNRGBA(Difference{}, b) }
func BenchmarkDrawRGBAToNRGBAExclusion(b *testing.B)    { benchmarkDrawRGBAToNRGBA(Exclusion{}, b) }
func BenchmarkDrawRGBAToNRGBASubtract(b *testing.B)     { benchmarkDrawRGBAToNRGBA(Subtract{}, b) }
func BenchmarkDrawRGBAToNRGBADivide(b *testing.B)       { benchmarkDrawRGBAToNRGBA(Divide{}, b) }
func BenchmarkDrawRGBAToNRGBAHue(b *testing.B)          { benchmarkDrawRGBAToNRGBA(Hue{}, b) }
func BenchmarkDrawRGBAToNRGBASaturation(b *testing.B)   { benchmarkDrawRGBAToNRGBA(Saturation{}, b) }
func BenchmarkDrawRGBAToNRGBAColor(b *testing.B)        { benchmarkDrawRGBAToNRGBA(Color{}, b) }
func BenchmarkDrawRGBAToNRGBALuminosity(b *testing.B)   { benchmarkDrawRGBAToNRGBA(Luminosity{}, b) }

func benchmarkDrawNRGBAToRGBA(d drawer, b *testing.B) {
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		b.Fatalf("%v", err)
	}
	img := image.NewRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawNRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func BenchmarkDrawNRGBAToRGBANormal(b *testing.B)       { benchmarkDrawNRGBAToRGBA(Normal{}, b) }
func BenchmarkDrawNRGBAToRGBADarken(b *testing.B)       { benchmarkDrawNRGBAToRGBA(Darken{}, b) }
func BenchmarkDrawNRGBAToRGBAMultiply(b *testing.B)     { benchmarkDrawNRGBAToRGBA(Multiply{}, b) }
func BenchmarkDrawNRGBAToRGBAColorBurn(b *testing.B)    { benchmarkDrawNRGBAToRGBA(ColorBurn{}, b) }
func BenchmarkDrawNRGBAToRGBALinearBurn(b *testing.B)   { benchmarkDrawNRGBAToRGBA(LinearBurn{}, b) }
func BenchmarkDrawNRGBAToRGBADarkerColor(b *testing.B)  { benchmarkDrawNRGBAToRGBA(DarkerColor{}, b) }
func BenchmarkDrawNRGBAToRGBALighten(b *testing.B)      { benchmarkDrawNRGBAToRGBA(Lighten{}, b) }
func BenchmarkDrawNRGBAToRGBAScreen(b *testing.B)       { benchmarkDrawNRGBAToRGBA(Screen{}, b) }
func BenchmarkDrawNRGBAToRGBAColorDodge(b *testing.B)   { benchmarkDrawNRGBAToRGBA(ColorDodge{}, b) }
func BenchmarkDrawNRGBAToRGBALinearDodge(b *testing.B)  { benchmarkDrawNRGBAToRGBA(LinearDodge{}, b) }
func BenchmarkDrawNRGBAToRGBALighterColor(b *testing.B) { benchmarkDrawNRGBAToRGBA(LighterColor{}, b) }
func BenchmarkDrawNRGBAToRGBAOverlay(b *testing.B)      { benchmarkDrawNRGBAToRGBA(Overlay{}, b) }
func BenchmarkDrawNRGBAToRGBASoftLight(b *testing.B)    { benchmarkDrawNRGBAToRGBA(SoftLight{}, b) }
func BenchmarkDrawNRGBAToRGBAHardLight(b *testing.B)    { benchmarkDrawNRGBAToRGBA(HardLight{}, b) }
func BenchmarkDrawNRGBAToRGBAVividLight(b *testing.B)   { benchmarkDrawNRGBAToRGBA(VividLight{}, b) }
func BenchmarkDrawNRGBAToRGBALinearLight(b *testing.B)  { benchmarkDrawNRGBAToRGBA(LinearLight{}, b) }
func BenchmarkDrawNRGBAToRGBAPinLight(b *testing.B)     { benchmarkDrawNRGBAToRGBA(PinLight{}, b) }
func BenchmarkDrawNRGBAToRGBAHardMix(b *testing.B)      { benchmarkDrawNRGBAToRGBA(HardMix{}, b) }
func BenchmarkDrawNRGBAToRGBADifference(b *testing.B)   { benchmarkDrawNRGBAToRGBA(Difference{}, b) }
func BenchmarkDrawNRGBAToRGBAExclusion(b *testing.B)    { benchmarkDrawNRGBAToRGBA(Exclusion{}, b) }
func BenchmarkDrawNRGBAToRGBASubtract(b *testing.B)     { benchmarkDrawNRGBAToRGBA(Subtract{}, b) }
func BenchmarkDrawNRGBAToRGBADivide(b *testing.B)       { benchmarkDrawNRGBAToRGBA(Divide{}, b) }
func BenchmarkDrawNRGBAToRGBAHue(b *testing.B)          { benchmarkDrawNRGBAToRGBA(Hue{}, b) }
func BenchmarkDrawNRGBAToRGBASaturation(b *testing.B)   { benchmarkDrawNRGBAToRGBA(Saturation{}, b) }
func BenchmarkDrawNRGBAToRGBAColor(b *testing.B)        { benchmarkDrawNRGBAToRGBA(Color{}, b) }
func BenchmarkDrawNRGBAToRGBALuminosity(b *testing.B)   { benchmarkDrawNRGBAToRGBA(Luminosity{}, b) }

func benchmarkDrawRGBAToRGBA(d drawer, b *testing.B) {
	bg, fg, err := loadImages("png/bg.png", "png/fg.png")
	if err != nil {
		b.Fatalf("%v", err)
	}
	img := image.NewRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(fg.Bounds())
	draw.Draw(img2, bg.Bounds(), fg, image.Pt(0, 0), draw.Over)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func BenchmarkDrawRGBAToRGBANormal(b *testing.B)       { benchmarkDrawRGBAToRGBA(Normal{}, b) }
func BenchmarkDrawRGBAToRGBADarken(b *testing.B)       { benchmarkDrawRGBAToRGBA(Darken{}, b) }
func BenchmarkDrawRGBAToRGBAMultiply(b *testing.B)     { benchmarkDrawRGBAToRGBA(Multiply{}, b) }
func BenchmarkDrawRGBAToRGBAColorBurn(b *testing.B)    { benchmarkDrawRGBAToRGBA(ColorBurn{}, b) }
func BenchmarkDrawRGBAToRGBALinearBurn(b *testing.B)   { benchmarkDrawRGBAToRGBA(LinearBurn{}, b) }
func BenchmarkDrawRGBAToRGBADarkerColor(b *testing.B)  { benchmarkDrawRGBAToRGBA(DarkerColor{}, b) }
func BenchmarkDrawRGBAToRGBALighten(b *testing.B)      { benchmarkDrawRGBAToRGBA(Lighten{}, b) }
func BenchmarkDrawRGBAToRGBAScreen(b *testing.B)       { benchmarkDrawRGBAToRGBA(Screen{}, b) }
func BenchmarkDrawRGBAToRGBAColorDodge(b *testing.B)   { benchmarkDrawRGBAToRGBA(ColorDodge{}, b) }
func BenchmarkDrawRGBAToRGBALinearDodge(b *testing.B)  { benchmarkDrawRGBAToRGBA(LinearDodge{}, b) }
func BenchmarkDrawRGBAToRGBALighterColor(b *testing.B) { benchmarkDrawRGBAToRGBA(LighterColor{}, b) }
func BenchmarkDrawRGBAToRGBAOverlay(b *testing.B)      { benchmarkDrawRGBAToRGBA(Overlay{}, b) }
func BenchmarkDrawRGBAToRGBASoftLight(b *testing.B)    { benchmarkDrawRGBAToRGBA(SoftLight{}, b) }
func BenchmarkDrawRGBAToRGBAHardLight(b *testing.B)    { benchmarkDrawRGBAToRGBA(HardLight{}, b) }
func BenchmarkDrawRGBAToRGBAVividLight(b *testing.B)   { benchmarkDrawRGBAToRGBA(VividLight{}, b) }
func BenchmarkDrawRGBAToRGBALinearLight(b *testing.B)  { benchmarkDrawRGBAToRGBA(LinearLight{}, b) }
func BenchmarkDrawRGBAToRGBAPinLight(b *testing.B)     { benchmarkDrawRGBAToRGBA(PinLight{}, b) }
func BenchmarkDrawRGBAToRGBAHardMix(b *testing.B)      { benchmarkDrawRGBAToRGBA(HardMix{}, b) }
func BenchmarkDrawRGBAToRGBADifference(b *testing.B)   { benchmarkDrawRGBAToRGBA(Difference{}, b) }
func BenchmarkDrawRGBAToRGBAExclusion(b *testing.B)    { benchmarkDrawRGBAToRGBA(Exclusion{}, b) }
func BenchmarkDrawRGBAToRGBASubtract(b *testing.B)     { benchmarkDrawRGBAToRGBA(Subtract{}, b) }
func BenchmarkDrawRGBAToRGBADivide(b *testing.B)       { benchmarkDrawRGBAToRGBA(Divide{}, b) }
func BenchmarkDrawRGBAToRGBAHue(b *testing.B)          { benchmarkDrawRGBAToRGBA(Hue{}, b) }
func BenchmarkDrawRGBAToRGBASaturation(b *testing.B)   { benchmarkDrawRGBAToRGBA(Saturation{}, b) }
func BenchmarkDrawRGBAToRGBAColor(b *testing.B)        { benchmarkDrawRGBAToRGBA(Color{}, b) }
func BenchmarkDrawRGBAToRGBALuminosity(b *testing.B)   { benchmarkDrawRGBAToRGBA(Luminosity{}, b) }
