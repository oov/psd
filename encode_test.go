package psd_test

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"testing"

	"github.com/oov/psd"
	"github.com/stretchr/testify/assert"
)

func init() {
	psd.Debug = log.New(os.Stdout, "psd: ", log.Lshortfile)
}

func getOrig(t *testing.T) *psd.PSD {
	fr, err := os.Open("testdata/cmyk-spot.psd")
	if err != nil {
		t.Fatal(err)
	}
	doc, _, err := psd.Decode(fr, nil)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func TestEncodeDecode(t *testing.T) {
	doc := buildNew(t, psd.CompressionMethodRaw)

	// re-read and compare image data again
	docReread := writeRead(t, doc)
	assert.Equal(t, doc.Config.Channels, docReread.Config.Channels)
}

func TestDecodeChannelImages(t *testing.T) {
	doc := getOrig(t)
	imgs, err := doc.GetChannelImages()
	if err != nil {
		t.Fatal(err)
	}
	writeImage(t, imgs, "orig") // write imges to compare
}

func TestEncodeChannelImages(t *testing.T) {
	// add images to new doc
	doc := buildNew(t, psd.CompressionMethodRLE)
	imgs := make([]*image.Gray, doc.Config.Channels)

	var err error
	for i, fnm := range []string{
		"testdata/cmyk-spot-channel-5.png",
		"testdata/cmyk-spot-channel-6.png",
		"testdata/cmyk-spot-channel-7.png",
	} {
		imgs[i+4], err = readGrayImage(fnm)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := doc.AddImageChannelData(imgs); err != nil {
		t.Fatal(err)
	}

	// compare image data
	docOrig := getOrig(t)
	assert.Len(t, doc.Data, doc.Config.Channels*doc.Config.Rect.Dx()*doc.Config.Rect.Dy())
	assert.EqualValues(t, docOrig.Data, doc.Data)
	for i := 4; i < doc.Config.Channels; i++ {
		assert.EqualValues(t, docOrig.Channel[i].Data, doc.Channel[i].Data)
	}

	// write
	fw, err := os.Create("output/cmyk-spot.psd")
	if err != nil {
		t.Fatal(err)
	}
	defer fw.Close()
	if err := psd.Encode(doc, fw); err != nil {
		t.Fatal(err)
	}

	// read again and write those images to compare
	docNew := writeRead(t, doc)
	imgsEnc, err := docNew.GetChannelImages()
	if err != nil {
		t.Fatal(err)
	}
	writeImage(t, imgsEnc, "encoded")
}

func buildNew(t *testing.T, cmp psd.CompressionMethod) *psd.PSD {
	doc := &psd.PSD{
		Config: psd.Config{
			Version:           1,
			Rect:              image.Rect(0, 0, 640, 637),
			Channels:          7, // CMYK plus three spot channels
			Depth:             8,
			ColorMode:         psd.ColorModeCMYK,
			CompressionMethod: cmp,
		},
	}
	imgResources := make(map[int]psd.ImageResource)
	irAlphaNames, err := testAlphaNames.Encode()
	if err != nil {
		t.Fatal(err)
	}
	irDisplayInfo, err := testDisplayInfo.Encode()
	if err != nil {
		t.Fatal(err)
	}
	irResInfo, err := testResolutionInfo.Encode()
	if err != nil {
		t.Fatal(err)
	}
	imgResources[irAlphaNames.ID] = *irAlphaNames
	imgResources[irDisplayInfo.ID] = *irDisplayInfo
	imgResources[irResInfo.ID] = *irResInfo
	doc.Config.Res = imgResources

	// dummy image data
	imgs := make([]*image.Gray, doc.Config.Channels)
	for i := range imgs {
		imgs[i] = image.NewGray(doc.Config.Rect)
	}
	if err := doc.AddImageChannelData(imgs); err != nil {
		t.Fatal(err)
	}
	return doc
}

func writeRead(t *testing.T, doc *psd.PSD) *psd.PSD {
	buf := bytes.NewBuffer([]byte{})
	if err := psd.Encode(doc, buf); err != nil {
		t.Fatal(err)
	}

	log.Printf("The file is %d bytes long", len(buf.Bytes()))

	out, _, err := psd.Decode(buf, nil)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func writeImage(t *testing.T, imgs []*image.Gray, prefix string) {
	for c, img := range imgs {
		w, err := os.Create(fmt.Sprintf("output/%s-c%d.png", prefix, c))
		if err != nil {
			t.Fatal(err)
		}
		defer w.Close()
		if err := png.Encode(w, img); err != nil {
			t.Fatal(err)
		}
	}
}

func readGrayImage(fnm string) (*image.Gray, error) {
	r, err := os.Open(fnm)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return psd.ImgToGray(img), nil
}
