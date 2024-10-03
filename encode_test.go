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

func TestEncodeDecode(t *testing.T) {
	doc, _, _ := buildNew(t, psd.CompressionMethodRaw)

	// re-read and compare image data again
	docReread := writeRead(t, doc)
	assert.Equal(t, doc.Config.Channels, docReread.Config.Channels)
}

func TestDecodeImageResources(t *testing.T) {
	_, alphaNames, displayInfo := getOrig(t)
	_, an, di := buildNew(t, psd.CompressionMethodRLE)

	// image resources decoding matches
	assert.EqualValues(t, an, alphaNames)
	assert.EqualValues(t, di, displayInfo)
}

func TestEncodeAlphaNames(t *testing.T) {
	doc, _, _ := getOrig(t)
	an := &psd.AlphaNames{[]string{"blue", "fluorescent pink", "yellow"}}
	ir, err := an.Encode()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, doc.Config.Res[ir.ID].Data, ir.Data)
}

func TestDecodeChannelImages(t *testing.T) {
	doc, _, _ := getOrig(t)
	imgs, err := doc.GetChannelImages()
	if err != nil {
		t.Fatal(err)
	}
	writeImage(t, imgs, "orig") // write imges to compare
}

func TestEncodeChannelImages(t *testing.T) {
	// add images to new doc
	doc, _, _ := buildNew(t, psd.CompressionMethodRLE)
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
	docOrig, _, _ := getOrig(t)
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

func getOrig(t *testing.T) (*psd.PSD, *psd.AlphaNames, *psd.DisplayInfo) {
	fr, err := os.Open("testdata/cmyk-spot.psd")
	if err != nil {
		t.Fatal(err)
	}
	docOrig, _, err := psd.Decode(fr, nil)
	if err != nil {
		t.Fatal(err)
	}
	alphaNames, err := docOrig.Config.ParseAlphaNames()
	if err != nil {
		t.Fatal(err)
	}
	displayInfo, err := docOrig.Config.ParseDisplayInfo()
	if err != nil {
		t.Fatal(err)
	}
	return docOrig, alphaNames, displayInfo
}

func buildNew(t *testing.T, cmp psd.CompressionMethod) (*psd.PSD, *psd.AlphaNames, *psd.DisplayInfo) {
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
	an := &psd.AlphaNames{[]string{"blue", "fluorescent pink", "yellow"}}

	di := &psd.DisplayInfo{
		Channels: []psd.DisplayInfoChannel{
			{
				Color: [4]uint16{0, 30840, 49087, 0},
				Mode:  psd.DisplayChannelModeSpot,
			},
			{
				Color: [4]uint16{65535, 18504, 45232, 0},
				Mode:  psd.DisplayChannelModeSpot,
			},
			{
				Color: [4]uint16{65535, 59624, 0, 0},
				Mode:  psd.DisplayChannelModeSpot,
			},
		},
	}

	imgResources := make(map[int]psd.ImageResource)
	irAlphaNames, err := an.Encode()
	if err != nil {
		t.Fatal(err)
	}
	irDisplayInfo, err := di.Encode()
	if err != nil {
		t.Fatal(err)
	}
	imgResources[irAlphaNames.ID] = *irAlphaNames
	imgResources[irDisplayInfo.ID] = *irDisplayInfo
	doc.Config.Res = imgResources

	// dummy image data
	imgs := make([]*image.Gray, doc.Config.Channels)
	for i := range imgs {
		imgs[i] = image.NewGray(doc.Config.Rect)
	}
	if err := doc.AddImageChannelData(imgs); err != nil {
		t.Fatal(err)
	}

	return doc, an, di
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
