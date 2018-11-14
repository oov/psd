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
	var score float64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ar, ag, ab, aa := a.At(x, y).RGBA()
			br, bg, bb, ba := b.At(x, y).RGBA()
			switch {
			case math.Abs(float64(ar)-float64(br)) >= 0x4000 ||
				math.Abs(float64(ag)-float64(bg)) >= 0x4000 ||
				math.Abs(float64(ab)-float64(bb)) >= 0x4000 ||
				math.Abs(float64(aa)-float64(ba)) >= 0x4000:
				score += 256.0
			case math.Abs(float64(ar)-float64(br)) >= 0x2000 ||
				math.Abs(float64(ag)-float64(bg)) >= 0x2000 ||
				math.Abs(float64(ab)-float64(bb)) >= 0x2000 ||
				math.Abs(float64(aa)-float64(ba)) >= 0x2000:
				score += 128.0
			case math.Abs(float64(ar)-float64(br)) >= 0x1000 ||
				math.Abs(float64(ag)-float64(bg)) >= 0x1000 ||
				math.Abs(float64(ab)-float64(bb)) >= 0x1000 ||
				math.Abs(float64(aa)-float64(ba)) >= 0x1000:
				score += 64.0
			case math.Abs(float64(ar)-float64(br)) >= 0x0400 ||
				math.Abs(float64(ag)-float64(bg)) >= 0x0400 ||
				math.Abs(float64(ab)-float64(bb)) >= 0x0400 ||
				math.Abs(float64(aa)-float64(ba)) >= 0x0400:
				score += 16.0
			case math.Abs(float64(ar)-float64(br)) >= 0x0100 ||
				math.Abs(float64(ag)-float64(bg)) >= 0x0100 ||
				math.Abs(float64(ab)-float64(bb)) >= 0x0100 ||
				math.Abs(float64(aa)-float64(ba)) >= 0x0100:
				score += 1.0
			}
		}
	}
	return score / float64(a.Bounds().Dx()*a.Bounds().Dy()), nil
}

func loadNRGBAToNRGBAImages(path1 string, path2 string) (*image.NRGBA, *image.NRGBA, error) {
	src1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	src2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}

	img := image.NewNRGBA(src1.Bounds())
	draw.Draw(img, src1.Bounds(), src1, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(src2.Bounds())
	draw.Draw(img2, src2.Bounds(), src2, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadRGBAToNRGBAImages(path1 string, path2 string) (*image.NRGBA, *image.RGBA, error) {
	src1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	src2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}

	img := image.NewNRGBA(src1.Bounds())
	draw.Draw(img, src1.Bounds(), src1, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(src2.Bounds())
	draw.Draw(img2, src2.Bounds(), src2, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadNRGBAToRGBAImages(path1 string, path2 string) (*image.RGBA, *image.NRGBA, error) {
	src1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	src2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}

	img := image.NewRGBA(src1.Bounds())
	draw.Draw(img, src1.Bounds(), src1, image.Pt(0, 0), draw.Over)
	img2 := image.NewNRGBA(src2.Bounds())
	draw.Draw(img2, src2.Bounds(), src2, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadRGBAToRGBAImages(path1 string, path2 string) (*image.RGBA, *image.RGBA, error) {
	src1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	src2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}

	img := image.NewRGBA(src1.Bounds())
	draw.Draw(img, src1.Bounds(), src1, image.Pt(0, 0), draw.Over)
	img2 := image.NewRGBA(src2.Bounds())
	draw.Draw(img2, src2.Bounds(), src2, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadAlphaToRGBAImages(path1 string, path2 string) (*image.RGBA, *image.Alpha, error) {
	src1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	src2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}

	img := image.NewRGBA(src1.Bounds())
	draw.Draw(img, src1.Bounds(), src1, image.Pt(0, 0), draw.Over)
	img2 := image.NewAlpha(src2.Bounds())
	draw.Draw(img2, src2.Bounds(), src2, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func loadAlphaToNRGBAImages(path1 string, path2 string) (*image.NRGBA, *image.Alpha, error) {
	src1, err := loadImage(path1)
	if err != nil {
		return nil, nil, err
	}
	src2, err := loadImage(path2)
	if err != nil {
		return nil, nil, err
	}

	img := image.NewNRGBA(src1.Bounds())
	draw.Draw(img, src1.Bounds(), src1, image.Pt(0, 0), draw.Over)
	img2 := image.NewAlpha(src2.Bounds())
	draw.Draw(img2, src2.Bounds(), src2, image.Pt(0, 0), draw.Over)
	return img, img2, nil
}

func testDrawFallback(t *testing.T, path1 string, path2 string, d drawer) {
	name := "DrawFallback"

	img, img2, err := loadNRGBAToNRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
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

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func testDrawNRGBAToNRGBA(t *testing.T, path1 string, path2 string, d drawer) {
	name := "DrawNRGBAToNRGBA"

	img, img2, err := loadNRGBAToNRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
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

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func testDrawRGBAToNRGBA(t *testing.T, path1 string, path2 string, d drawer) {
	name := "DrawRGBAToNRGBA"

	img, img2, err := loadRGBAToNRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
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

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func testDrawNRGBAToRGBA(t *testing.T, path1 string, path2 string, d drawer) {
	name := "DrawNRGBAToRGBA"

	img, img2, err := loadNRGBAToRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
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

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func testDrawRGBAToRGBA(t *testing.T, path1 string, path2 string, d drawer) {
	name := "DrawRGBAToRGBA"

	img, img2, err := loadRGBAToRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
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

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func testDrawAlphaToRGBA(t *testing.T, path1 string, path2 string, d alphaDrawer) {
	name := "DrawAlphaToRGBA"

	img, img2, err := loadAlphaToRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawAlphaToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/alpha/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func testDrawAlphaToNRGBA(t *testing.T, path1 string, path2 string, d alphaDrawer) {
	name := "DrawAlphaToNRGBA"

	img, img2, err := loadAlphaToNRGBAImages(path1, path2)
	if err != nil {
		t.Fatalf("%v", err)
	}
	d.drawAlphaToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)

	os.MkdirAll("output/"+name, os.ModePerm)
	if err = saveImage(fmt.Sprintf("output/%s/%v.png", name, d), img); err != nil {
		t.Fatalf("Cannot create png: %v", err)
	}

	created, ref, err := loadImages(
		fmt.Sprintf("output/%s/%v.png", name, d),
		fmt.Sprintf("reference/alpha/%v.png", d),
	)
	if err != nil {
		t.Fatalf("%v", err)
	}

	score, err := verify(created, ref)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Logf("score: %f", score)
	if score > 1 {
		t.Errorf("too many erros: %f", score)
	}
}

func benchmarkDrawFallback(b *testing.B, path1 string, path2 string, d drawer) {
	img, img2, err := loadNRGBAToNRGBAImages(path1, path2)
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawFallback(img, img2.Bounds(), img2, image.Pt(0, 0), nil, image.Pt(0, 0))
	}
}

func benchmarkDrawNRGBAToNRGBA(b *testing.B, path1 string, path2 string, d drawer) {
	img, img2, err := loadNRGBAToNRGBAImages(path1, path2)
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawNRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func benchmarkDrawRGBAToNRGBA(b *testing.B, path1 string, path2 string, d drawer) {
	img, img2, err := loadRGBAToNRGBAImages(path1, path2)
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawRGBAToNRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func benchmarkDrawNRGBAToRGBA(b *testing.B, path1 string, path2 string, d drawer) {
	img, img2, err := loadNRGBAToRGBAImages(path1, path2)
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawNRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func benchmarkDrawRGBAToRGBA(b *testing.B, path1 string, path2 string, d drawer) {
	img, img2, err := loadRGBAToRGBAImages(path1, path2)
	if err != nil {
		b.Fatalf("%v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.drawRGBAToRGBAUniform(img, img2.Bounds(), img2, image.Pt(0, 0), nil)
	}
}

func TestSliceBoundsOutOfRange(t *testing.T) {
	x := image.NewNRGBA(image.Rect(0, 0, 48, 48))
	x2 := image.NewNRGBA(image.Rect(0, 0, 48, 27))
	Copy.Draw(x, image.Rect(0, 10, 48, 37), x2, image.Point{})
}
