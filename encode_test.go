package psd_test

import (
	"image"
	"os"
	"testing"

	"github.com/oov/psd"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	doc := &psd.PSD{
		Config: psd.Config{
			Version:   1,
			Rect:      image.Rect(0, 0, 2400, 3554),
			Channels:  6, // CMYK plus two spot channels
			Depth:     8,
			ColorMode: psd.ColorModeCMYK,
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
	var imgData []byte
	// TODO
	doc.Data = imgData

	fw, err := os.Create("testdata/cmyk-adam-encode.psd")
	if err != nil {
		t.Fatal(err)
	}
	defer fw.Close()

	if err := psd.Encode(doc, fw); err != nil {
		t.Fatal(err)
	}

	// compare with original
	fr, err := os.Open("testdata/cmyk-adam.psd")
	if err != nil {
		t.Fatal(err)
	}
	docOrig, _, err := psd.Decode(fr, &psd.DecodeOptions{SkipMergedImage: true})
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

	assert.Equal(t, docOrig.Config.Channels, doc.Config.Channels)
	assert.Equal(t, docOrig.Config.ColorMode, doc.Config.ColorMode)
	assert.Equal(t, docOrig.Config.Depth, doc.Config.Depth)
	assert.Equal(t, docOrig.Config.Rect, doc.Config.Rect)
	assert.Equal(t, *an, *alphaNames)
	assert.Equal(t, *di, *displayInfo)

}
