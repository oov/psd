package psd_test

import (
	"testing"

	"github.com/oov/psd"
	"github.com/stretchr/testify/assert"
)

var testAlphaNames = &psd.AlphaNames{[]string{"blue", "fluorescent pink", "yellow"}}
var testDisplayInfo = &psd.DisplayInfo{
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
var testResolutionInfo = &psd.ResolutionInfo{
	HorizontalRes:  19660800,
	HorizontalUnit: 1,
	WidthUnit:      1,
	VerticalRes:    19660800,
	VerticalUnit:   1,
	HeightUnit:     1,
}

func TestParseAlpha(t *testing.T) {
	doc := getOrig(t)
	alphaNames, err := doc.Config.ParseAlphaNames()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, testAlphaNames, alphaNames)
}

func TestParseDisplayInfo(t *testing.T) {
	doc := getOrig(t)
	displayInfo, err := doc.Config.ParseDisplayInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, testDisplayInfo, displayInfo)
}

func TestParseResolutionInfo(t *testing.T) {
	doc := getOrig(t)
	resInfo, err := doc.Config.ParseResolutionInfo()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, testResolutionInfo, resInfo)
}

func TestEncodeAlphaNames(t *testing.T) {
	doc := getOrig(t)
	ir, err := testAlphaNames.Encode()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, doc.Config.Res[ir.ID].Data, ir.Data)
}

func TestEncodeDisplayInfo(t *testing.T) {
	doc := getOrig(t)
	ir, err := testDisplayInfo.Encode()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, doc.Config.Res[ir.ID].Data, ir.Data)
}

func TestEncodeResolutionInfoInfo(t *testing.T) {
	doc := getOrig(t)
	ir, err := testResolutionInfo.Encode()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, doc.Config.Res[ir.ID].Data, ir.Data)
}
