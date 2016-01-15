package psd

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

type testImage struct {
	Name string
	PSD  string
}

var testImages = []testImage{
	{
		Name: "Bitmap",
		PSD:  "bitmap.psd",
	},
	{
		Name: "Grayscale Depth 8bit with Background layer",
		PSD:  "grayscale8bit.psd",
	},
	{
		Name: "Grayscale Depth 8bit without Background layer(partially transparent)",
		PSD:  "grayscale8bit_transparent.psd",
	},
	{
		Name: "Grayscale Depth 16bit Gradient",
		PSD:  "grayscale16bit_grad.psd",
	},
	{
		Name: "Grayscale Depth 16bit with Background layer",
		PSD:  "grayscale16bit.psd",
	},
	{
		Name: "Grayscale Depth 16bit without Background layer(partially transparent)",
		PSD:  "grayscale16bit_transparent.psd",
	},
	{
		Name: "Grayscale Depth 32bit",
		PSD:  "grayscale32bit.psd",
	},
	{
		Name: "Indexed",
		PSD:  "indexed.psd",
	},
	{
		Name: "Indexed + Transparent Color",
		PSD:  "indexed_transparent.psd",
	},
	{
		Name: "RGB Depth 8bit with Background layer",
		PSD:  "rgb8bit.psd",
	},
	{
		Name: "RGB Depth 8bit without Background layer",
		PSD:  "rgb8bit_nobg.psd",
	},
	{
		Name: "RGB Depth 8bit without Background layer(partially transparent)",
		PSD:  "rgb8bit_nobg_transparent.psd",
	},
	{
		Name: "RGB Depth 8bit with Background layer including one added alpha channel",
		PSD:  "rgb8bit+1ch.psd",
	},
	{
		Name: "RGB Depth 8bit without Background layer(partially transparent) including one added alpha channel",
		PSD:  "rgb8bit+1ch_transparent.psd",
	},
	{
		Name: "RGB Depth 8bit without Background layer(partially transparent) including two added alpha channels and one added spot color channel",
		PSD:  "rgb8bit+2ch+spot_transparent.psd",
	},
	{
		Name: "RGB Depth 16bit with Background layer",
		PSD:  "rgb16bit.psd",
	},
	// {
	// 	Name: "RGB Depth 32bit",
	// 	PSD:  "rgb32bit.psd",
	// },
	{
		Name: "CMYK Depth 8bit with Background layer",
		PSD:  "cmyk8bit.psd",
	},
	{
		Name: "CMYK Depth 8bit without Background layer(partially transparent)",
		PSD:  "cmyk8bit_transparent.psd",
	},
	// {
	// 	Name: "CMYK Depth 16bit with Background layer",
	// 	PSD:  "cmyk16bit.psd",
	// },
}

func processLayer(f string, l *Layer) error {
	if len(l.Layer) > 0 {
		for i, ll := range l.Layer {
			if err := processLayer(fmt.Sprintf("%s_%d", f, i), &ll); err != nil {
				return err
			}
		}
	}
	if !l.HasImage() {
		return nil
	}

	//	// write layer image per channel
	//	for id, ch := range l.Channel {
	//		if err := func() error {
	//			filename := fmt.Sprintf("%s_Ch%d.png", f, id)
	//			o, err := os.Create(filename)
	//			if err != nil {
	//				return err
	//			}
	//			defer o.Close()
	//			return png.Encode(o, &ch)
	//		}(); err != nil {
	//			return err
	//		}
	//	}

	filename := fmt.Sprintf("%s.png", f)
	o, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer o.Close()
	img, err := l.ExperimentalExportImage()
	if err != nil {
		return png.Encode(o, l)
	}
	return png.Encode(o, img)

}

func testOne(tImg testImage, t *testing.T) {
	fmt.Printf("%s begin\n", tImg.Name)
	filepath := "testdata/" + tImg.PSD
	f, err := os.Open(filepath)
	if err != nil {
		t.Errorf("%s: cannot open %q\n%v", tImg.Name, filepath, err)
		return
	}
	defer f.Close()

	img, _, err := Decode(f)
	if err != nil {
		t.Errorf("%s: error occurred in psd.Decode\n%v", tImg.Name, err)
		return
	}
	if img == nil {
		t.Errorf("%s: got nil want image.Image", tImg.Name)
		return
	}

	fnBase := "output/" + tImg.PSD[:len(tImg.PSD)-4]
	// write merged image
	func() {
		filename := fmt.Sprintf("%s_!merged.png", fnBase)
		o, err := os.Create(filename)
		if err != nil {
			t.Errorf("%s: cannot create file %q\n%v", tImg.Name, filename, err)
		}
		defer o.Close()
		if err = png.Encode(o, img); err != nil {
			t.Errorf("%s: cannot encode to %q\n%v", tImg.Name, filename, err)
		}
	}()
	//	// write merged image per channel
	//	for id, ch := range img.Channel {
	//		func() {
	//			filename := fmt.Sprintf("%s_!merged_Ch%d.png", fnBase, id)
	//			o, err := os.Create(filename)
	//			if err != nil {
	//				t.Errorf("%s: cannot create file %q\n%v", tImg.Name, filename, err)
	//			}
	//			defer o.Close()
	//			if err = png.Encode(o, &ch); err != nil {
	//				t.Errorf("%s: cannot encode to %q\n%v", tImg.Name, filename, err)
	//			}
	//		}()
	//	}
	for i, layer := range img.Layer {
		if err = processLayer(fmt.Sprintf("%s_%d", fnBase, i), &layer); err != nil {
			t.Errorf("%s: cannot create file\n%v", tImg.Name, err)
		}
	}
}

func TestOneShot(t *testing.T) {
	tImg := testImages[3]
	testOne(tImg, t)
}

func TestAll(t *testing.T) {
	for _, tImg := range testImages {
		testOne(tImg, t)
	}
}
