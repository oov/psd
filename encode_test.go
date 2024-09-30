package psd_test

import (
	"image"
	"image/draw"
	"os"
	"testing"

	"github.com/oov/psd"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	doc, an, di := buildNew(t)
	docOrig, alphaNames, displayInfo := getOrig(t)

	// config matches original
	assert.Equal(t, docOrig.Config.Channels, doc.Config.Channels)
	assert.Equal(t, docOrig.Config.ColorMode, doc.Config.ColorMode)
	assert.Equal(t, docOrig.Config.Depth, doc.Config.Depth)
	assert.Equal(t, docOrig.Config.Rect, doc.Config.Rect)
	assert.Equal(t, docOrig.Config.CompressionMethod, doc.Config.CompressionMethod)
	assert.EqualValues(t, an, alphaNames)
	assert.EqualValues(t, di, displayInfo)

	// read images
	imgs := make([]*image.Gray, 6)
	for c := range imgs {
		imgs[c] = imgToGray(docOrig.Channel[c].Picker)
	}
	// add images to new doc
	if err := doc.AddImageChannelData(imgs); err != nil {
		t.Fatal(err)
	}

	// compare image data
	assert.Len(t, doc.Data, doc.Config.Channels*doc.Config.Rect.Dx()*doc.Config.Rect.Dy())
	assert.EqualValues(t, docOrig.Data, doc.Data)
	assert.EqualValues(t, docOrig.Channel[3].Data, doc.Channel[3].Data)
	assert.EqualValues(t, docOrig.Channel[4].Data, doc.Channel[4].Data)
	assert.EqualValues(t, docOrig.Channel[5].Data, doc.Channel[5].Data)

	// write
	fw, err := os.Create("testdata/cmyk-adam-encode.psd")
	if err != nil {
		t.Fatal(err)
	}
	defer fw.Close()
	if err := psd.Encode(doc, fw); err != nil {
		t.Fatal(err)
	}

	// re-read and compare image data again
	docReread := readNew(t)
	assert.EqualValues(t, docOrig.Data, docReread.Data)
}

func getOrig(t *testing.T) (*psd.PSD, *psd.AlphaNames, *psd.DisplayInfo) {
	fr, err := os.Open("testdata/cmyk-adam.psd")
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

func buildNew(t *testing.T) (*psd.PSD, *psd.AlphaNames, *psd.DisplayInfo) {
	doc := &psd.PSD{
		Config: psd.Config{
			Version:           1,
			Rect:              image.Rect(0, 0, 2400, 3554),
			Channels:          6, // CMYK plus two spot channels
			Depth:             8,
			ColorMode:         psd.ColorModeCMYK,
			CompressionMethod: psd.CompressionMethodRLE,
		},
	}
	an := &psd.AlphaNames{[]string{"coral", "light teal"}}

	di := &psd.DisplayInfo{
		Channels: []psd.DisplayInfoChannel{
			{
				ColorSpace: 9009, // what is this??
				Color:      [4]uint16{25650, 25954, 25956, 0},
				Opacity:    0,
				Mode:       psd.DisplayChannelModeSpot,
			},
			{
				ColorSpace: 9009,
				Color:      [4]uint16{13670, 25400, 13157, 0},
				Opacity:    0,
				Mode:       psd.DisplayChannelModeSpot,
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

	return doc, an, di
}

func readNew(t *testing.T) *psd.PSD {
	fr, err := os.Open("testdata/cmyk-adam-encode.psd")
	if err != nil {
		t.Fatal(err)
	}
	doc, _, err := psd.Decode(fr, nil)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func imgToGray(img image.Image) *image.Gray {
	out := image.NewGray(img.Bounds())
	draw.Draw(out, out.Rect, img, image.Point{}, draw.Src)
	return out
}
