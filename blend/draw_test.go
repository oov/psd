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

func loadTestImages() (image.Image, image.Image, error) {
	return loadImages("png/bg.png", "png/fg.png")
}

func loadNRGBAToNRGBAImages() (*image.NRGBA, *image.NRGBA, error) {
	bg, fg, err := loadTestImages()
	if err != nil {
		return nil, nil, err
	}

	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, fg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadRGBAToNRGBAImages() (*image.NRGBA, *image.RGBA, error) {
	bg, fg, err := loadTestImages()
	if err != nil {
		return nil, nil, err
	}

	img := image.NewNRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(fg.Bounds())
	draw.Draw(img2, fg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadNRGBAToRGBAImages() (*image.RGBA, *image.NRGBA, error) {
	bg, fg, err := loadTestImages()
	if err != nil {
		return nil, nil, err
	}

	img := image.NewRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(fg.Bounds())
	draw.Draw(img2, fg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadRGBAToRGBAImages() (*image.RGBA, *image.RGBA, error) {
	bg, fg, err := loadTestImages()
	if err != nil {
		return nil, nil, err
	}

	img := image.NewRGBA(bg.Bounds())
	draw.Draw(img, bg.Bounds(), bg, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(fg.Bounds())
	draw.Draw(img2, fg.Bounds(), fg, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func testDrawFallback(d drawer, t *testing.T, protectAlpha bool) {
	name := "DrawFallback"
	img, img2, err := loadNRGBAToNRGBAImages()
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawFallback(img, img2.Bounds(), img2, image.Pt(0, 0), nil, image.Pt(0, 0), protectAlpha)

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
	if errorRate > 0.262 {
		t.Errorf("too many erros: %3.2f%%", errorRate*100)
	}
}

func testDrawNRGBAToNRGBA(d drawer, t *testing.T, protectAlpha bool) {
	name := "DrawNRGBAToNRGBA"
	img, img2, err := loadNRGBAToNRGBAImages()
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawNRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)

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
	if errorRate > 0.262 {
		t.Errorf("too many erros: %3.2f%%", errorRate*100)
	}
}

func testDrawRGBAToNRGBA(d drawer, t *testing.T, protectAlpha bool) {
	name := "DrawRGBAToNRGBA"
	img, img2, err := loadRGBAToNRGBAImages()
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)

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
	if errorRate > 0.262 {
		t.Errorf("too many erros: %3.2f%%", errorRate*100)
	}
}

func testDrawNRGBAToRGBA(d drawer, t *testing.T, protectAlpha bool) {
	name := "DrawNRGBAToRGBA"
	img, img2, err := loadNRGBAToRGBAImages()
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawNRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)

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
	if errorRate > 0.262 {
		t.Errorf("too many erros: %3.2f%%", errorRate*100)
	}
}

func testDrawRGBAToRGBA(d drawer, t *testing.T, protectAlpha bool) {
	name := "DrawRGBAToRGBA"
	img, img2, err := loadRGBAToRGBAImages()
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)

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
	if errorRate > 0.262 {
		t.Errorf("too many erros: %3.2f%%", errorRate*100)
	}
}

func benchmarkDrawFallback(d drawer, b *testing.B, protectAlpha bool) {
	img, img2, err := loadNRGBAToNRGBAImages()
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawFallback(img, img2.Bounds(), img2, image.Pt(0, 0), nil, image.Pt(0, 0), protectAlpha)
	}
}

func benchmarkDrawNRGBAToNRGBA(d drawer, b *testing.B, protectAlpha bool) {
	img, img2, err := loadNRGBAToNRGBAImages()
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawNRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)
	}
}

func benchmarkDrawRGBAToNRGBA(d drawer, b *testing.B, protectAlpha bool) {
	img, img2, err := loadRGBAToNRGBAImages()
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)
	}
}

func benchmarkDrawNRGBAToRGBA(d drawer, b *testing.B, protectAlpha bool) {
	img, img2, err := loadNRGBAToRGBAImages()
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawNRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)
	}
}

func benchmarkDrawRGBAToRGBA(d drawer, b *testing.B, protectAlpha bool) {
	img, img2, err := loadRGBAToRGBAImages()
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil, protectAlpha)
	}
}
