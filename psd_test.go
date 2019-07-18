package psd

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
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
	{
		Name: "RGB Depth 32bit",
		PSD:  "rgb32bit.psd",
	},
	{
		Name: "CMYK Depth 8bit Min Max",
		PSD:  "cmyk8bit_minmax.psd",
	},
	{
		Name: "CMYK Depth 8bit with Background layer",
		PSD:  "cmyk8bit.psd",
	},
	{
		Name: "CMYK Depth 8bit without Background layer(partially transparent)",
		PSD:  "cmyk8bit_transparent.psd",
	},
	{
		Name: "CMYK Depth 16bit with Background layer",
		PSD:  "cmyk16bit.psd",
	},
	{
		Name: "Clipping Mask",
		PSD:  "clipping.psd",
	},
	{
		Name: "Layer Mask & Vector Mask",
		PSD:  "mask.psd",
	},
	{
		Name: "Normal Format(maximize compatibility on)",
		PSD:  "psd_compat.psd",
	},
	{
		Name: "Normal Format(maximize compatibility off)",
		PSD:  "psd_nocompat.psd",
	},
	{
		Name: "Big Document Format(maximize compatibility on)",
		PSD:  "psb_compat.psb",
	},
	{
		Name: "Big Document Format(maximize compatibility off)",
		PSD:  "psb_nocompat.psb",
	},
}

func verifyChannel(t *testing.T, name string, filename string, ch image.Image) {
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("%s: cannot open %s %v", name, filename, err)
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		t.Errorf("%s: cannot decode %s %v", name, filename, err)
		return
	}
	w, h := ch.Bounds().Dx(), ch.Bounds().Dy()
	if img.Bounds().Dx() != w {
		t.Errorf("%s: width: want %d got %d", name, img.Bounds().Dx(), w)
		return
	}
	if img.Bounds().Dy() != h {
		t.Errorf("%s: height: want %d got %d", name, img.Bounds().Dy(), h)
		return
	}
	ofsX, ofsY := ch.Bounds().Min.X, ch.Bounds().Min.Y
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c1, _, _, _ := ch.At(ofsX+x, ofsY+y).RGBA()
			c2, _, _, _ := img.At(x, y).RGBA()
			if c1 != c2 {
				t.Errorf("%s: (%d, %d): want %d got %d", name, x, y, c2, c1)
			}
		}
	}
}

func verifyImage(t *testing.T, name string, filename string, l image.Image) {
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("%s: cannot open %s %v", name, filename, err)
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		t.Errorf("%s: cannot decode %s %v", name, filename, err)
		return
	}
	w, h := l.Bounds().Dx(), l.Bounds().Dy()
	if img.Bounds().Dx() != w {
		t.Errorf("%s width: want %d got %d", name, img.Bounds().Dx(), w)
		return
	}
	if img.Bounds().Dy() != h {
		t.Errorf("%s height: want %d got %d", name, img.Bounds().Dy(), h)
		return
	}
	ofsX, ofsY := l.Bounds().Min.X, l.Bounds().Min.Y
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// RGBA returns the alpha-premultiplied color values,
			// but img stored the non-alpha-premultiplied color values.
			// So it allowed the error.
			r1, g1, b1, a1 := l.At(ofsX+x, ofsY+y).RGBA()
			r2, g2, b2, a2 := img.At(x, y).RGBA()
			if abs(r1-r2) > 1 || abs(g1-g2) > 1 || abs(b1-b2) > 1 || a1 != a2 {
				t.Errorf(
					"%s (%d, %d): want 0x%04x,0x%04x,0x%04x,0x%04x got 0x%04x,0x%04x,0x%04x,0x%04x",
					name, x, y,
					r2, g2, b2, a2,
					r1, g1, b1, a1,
				)
			}
		}
	}
}

func processLayer(t *testing.T, f string, l *Layer) error {
	if len(l.Layer) > 0 {
		for i, ll := range l.Layer {
			if err := processLayer(t, fmt.Sprintf("%s_%d", f, i), &ll); err != nil {
				return err
			}
		}
	}
	if !l.HasImage() {
		return nil
	}

	// write layer image per channel
	for id, ch := range l.Channel {
		if err := func() error {
			o, err := os.Create(fmt.Sprintf("output/%s_Ch%d.png", f, id))
			if err != nil {
				return err
			}
			defer o.Close()
			return png.Encode(o, ch.Picker)
		}(); err != nil {
			return err
		}
	}

	for id, ch := range l.Channel {
		verifyChannel(
			t,
			fmt.Sprintf("%s Ch:%d", f, id),
			fmt.Sprintf("png/%s_Ch%d.png", f, id),
			ch.Picker,
		)
	}

	// write layer image
	o, err := os.Create(fmt.Sprintf("output/%s.png", f))
	if err != nil {
		return err
	}
	defer o.Close()
	err = png.Encode(o, l.Picker)
	if err != nil {
		return err
	}

	verifyImage(t, f, fmt.Sprintf("png/%s.png", f), l.Picker)
	return nil
}

func testOne(tImg testImage, t *testing.T) {
	t.Logf("%s begin\n", tImg.Name)
	filepath := "testdata/" + tImg.PSD
	f, err := os.Open(filepath)
	if err != nil {
		t.Errorf("%s: cannot open %q\n%v", tImg.Name, filepath, err)
		return
	}
	defer f.Close()

	psdImg, _, err := Decode(f, nil)
	if err != nil {
		t.Errorf("%s: error occurred in psd.Decode\n%v", tImg.Name, err)
		return
	}
	if psdImg == nil {
		t.Errorf("%s: got nil want image.Image", tImg.Name)
		return
	}

	fnBase := tImg.PSD[:len(tImg.PSD)-4]

	// write merged image per channel
	for id, ch := range psdImg.Channel {
		func() {
			filename := fmt.Sprintf("output/%s_!merged_Ch%d.png", fnBase, id)
			var o *os.File
			if o, err = os.Create(filename); err != nil {
				t.Errorf("%s: cannot create file %q\n%v", tImg.Name, filename, err)
			}
			defer o.Close()
			if err = png.Encode(o, ch.Picker); err != nil {
				t.Errorf("%s: cannot encode to %q\n%v", tImg.Name, filename, err)
			}
		}()
	}
	for id, ch := range psdImg.Channel {
		verifyChannel(
			t,
			fmt.Sprintf("%s !merged Ch:%d", tImg.Name, id),
			fmt.Sprintf("png/%s_!merged_Ch%d.png", fnBase, id),
			ch.Picker,
		)
	}

	for i, layer := range psdImg.Layer {
		if err = processLayer(t, fmt.Sprintf("%s_%d", fnBase, i), &layer); err != nil {
			t.Errorf("%s: cannot create file\n%v", tImg.Name, err)
		}
	}

	// write merged image
	func() {
		filename := fmt.Sprintf("output/%s_!merged.png", fnBase)
		o, err := os.Create(filename)
		if err != nil {
			t.Errorf("%s: cannot create file %q\n%v", tImg.Name, filename, err)
		}
		defer o.Close()
		if err = png.Encode(o, psdImg.Picker); err != nil {
			t.Errorf("%s: cannot encode to %q\n%v", tImg.Name, filename, err)
		}
	}()
	verifyImage(
		t,
		tImg.Name,
		fmt.Sprintf("png/%s_!merged.png", fnBase),
		psdImg.Picker,
	)
}

func abs(a uint32) uint32 {
	if a < 0 {
		return -a
	}
	return a
}

type testLogger struct {
	t *testing.T
}

func (l *testLogger) Printf(format string, v ...interface{}) {
	if _, file, line, ok := runtime.Caller(1); ok {
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
		l.t.Log(fmt.Sprintf("%s:%d:", file, line), fmt.Sprintf(format, v...))
		return
	}
	l.t.Logf(format, v...)
}

func (l *testLogger) Println(v ...interface{}) {
	if _, file, line, ok := runtime.Caller(1); ok {
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
		l.t.Log(fmt.Sprintf("%s:%d:", file, line), strings.TrimRight(fmt.Sprintln(v...), "\r\n"))
		return
	}
	l.t.Log(v...)
}

func TestOneShot(t *testing.T) {
	Debug = &testLogger{t}
	tImg := testImages[20]
	tImg.PSD = "test.psb"
	testOne(tImg, t)
}

func TestAll(t *testing.T) {
	for _, tImg := range testImages {
		testOne(tImg, t)
	}
}

func TestDetectFolderBlendMode(t *testing.T) {
	for _, filename := range []string{
		"mod.psd",
		"mod2.psd",
	} {
		file, err := os.Open("testdata/" + filename)
		if err != nil {
			t.Errorf("cannot open %s: %v", filename, err)
			return
		}
		defer file.Close()
		psd, _, err := Decode(file, nil)
		if err != nil {
			t.Errorf("cannot open %s as psd: %v", filename, err)
			return
		}
		{
			want := "フォルダー3"
			got := psd.Layer[2].Layer[0].UnicodeName
			if want != got {
				t.Error(filename, "want", want, "but got", got)
				return
			}
		}
		{
			want := BlendModeNormal
			got := psd.Layer[2].Layer[0].SectionDividerSetting.BlendMode
			if want != got {
				t.Error(filename, "want", want, "but got", got)
				return
			}
		}
	}
}

func Benchmark2MB(b *testing.B) {
	b.StopTimer()
	f, err := os.Open("testdata/benchmark.psd")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		b.Fatal(err)
	}
	r := bytes.NewReader(buf)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		Decode(r, nil)
		b.StopTimer()
		r.Reset(buf)
	}
}

func TestDecodeFail(t *testing.T) {
	brokenImages := []struct {
		b []byte
		n string
	}{
		{
			b: []byte{'8', 'B', 'P', 'S', 0, 1},
			n: "psd",
		},
		{
			b: []byte{'8', 'B', 'P', 'S', 0, 2},
			n: "psb",
		},
	}
	for _, bi := range brokenImages {
		img, name, err := image.Decode(bytes.NewReader(bi.b))
		if err != io.ErrUnexpectedEOF {
			t.Error("unexpected error:", err)
		}
		if name != bi.n {
			t.Error("unexpected format name: want", bi.n, "got", name)
		}
		if img != nil {
			t.Error("img is not nil")
		}
	}
}

func TestSketchBookPSDFail(t *testing.T) {
	// https://github.com/oov/psd/issues/3
	Debug = &testLogger{t}
	testOne(testImage{
		Name: "PSD(Autodesk SketchBook)",
		PSD:  "sketchbook.psd",
	}, t)
}
