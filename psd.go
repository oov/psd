package psd

import (
	"errors"
	"image"
	"image/color"
	"io"
)

// logger is subset of log.Logger.
type logger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Debug is useful for debugging.
//
// You can use by performing the following steps.
// 	psd.Debug = log.New(os.Stdout, "psd: ", log.Lshortfile)
var Debug logger

// AdditionalInfoKey represents the key of the additional layer information.
type AdditionalInfoKey string

// These keys are defined in this document.
//
// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_71546
const (
	AdditionalInfoKeyLayerInfo             = AdditionalInfoKey("Layr")
	AdditionalInfoKeyLayerInfo16           = AdditionalInfoKey("Lr16")
	AdditionalInfoKeyLayerInfo32           = AdditionalInfoKey("Lr32")
	AdditionalInfoKeyUnicodeLayerName      = AdditionalInfoKey("luni")
	AdditionalInfoKeyBlendClippingElements = AdditionalInfoKey("clbl")
	AdditionalInfoKeySectionDividerSetting = AdditionalInfoKey("lsct")
)

// Config represents Photoshop image file configuration.
type Config struct {
	Rect      image.Rectangle
	Channels  int
	Depth     int // 1 or 8 or 16 or 32
	ColorMode ColorMode
	Picker    Picker
	Res       map[int]ImageResource
}

// PSD represents Photoshop image file.
type PSD struct {
	Config             Config
	Channel            map[int]Channel
	Layer              []Layer
	AdditinalLayerInfo map[AdditionalInfoKey][]byte
	// Data is uncompressed merged image data.
	Data []byte
}

// ColorModel returns the Image's color model.
func (p *PSD) ColorModel() color.Model {
	return p.Config.Picker.ColorModel()
}

// Bounds returns the domain for which At can return non-zero color.
func (p *PSD) Bounds() image.Rectangle {
	return p.Config.Picker.Bounds()
}

// At returns the color of the pixel at (x, y).
func (p *PSD) At(x, y int) color.Color {
	return p.Config.Picker.At(x, y)
}

// DecodeConfig returns the color model and dimensions of a image without decoding the entire image.
func DecodeConfig(r io.Reader) (cfg Config, read int, err error) {
	const (
		fileHeaderLen    = 4 + 2 + 6 + 2 + 4 + 4 + 2 + 2
		colorModeDataLen = 4
	)
	b := make([]byte, fileHeaderLen+colorModeDataLen)
	var l int
	if l, err = io.ReadFull(r, b); err != nil {
		return Config{}, 0, err
	}
	read += l

	if string(b[:4]) != "8BPS" { // Signature
		return Config{}, read, errors.New("psd: invalid format")
	}
	if v := readUint16(b, 4); v != 1 { // Version
		return Config{}, read, errors.New("psd: unexpected file version: " + itoa(int(v)))
	}
	// b[6:12] Reserved

	cfg.Channels = int(readUint16(b, 12))
	if cfg.Channels < 1 || cfg.Channels > 56 {
		return Config{}, read, errors.New("psd: unexpected the number of channels")
	}

	width, height := int(readUint32(b, 18)), int(readUint32(b, 14))
	if height < 1 || height > 30000 {
		return Config{}, read, errors.New("psd: unexpected the height of the image")
	}
	if width < 1 || width > 30000 {
		return Config{}, read, errors.New("psd: unexpected the width of the image")
	}
	cfg.Rect = image.Rect(0, 0, width, height)

	cfg.Depth = int(readUint16(b, 22))
	if cfg.Depth != 1 && cfg.Depth != 8 && cfg.Depth != 16 && cfg.Depth != 32 {
		return Config{}, read, errors.New("psd: unexpected color depth")
	}

	cfg.ColorMode = ColorMode(readUint16(b, 24))

	var colorModeData []byte
	if ln := readUint32(b, 26); ln > 0 {
		colorModeData = make([]byte, ln)
		if l, err = io.ReadFull(r, colorModeData); err != nil {
			return Config{}, read, err
		}
		read += l
	}
	if Debug != nil {
		Debug.Println("start - header")
		Debug.Printf("  width: %d height: %d", width, height)
		Debug.Printf("  channels: %d depth: %d colorMode %d", cfg.Channels, cfg.Depth, cfg.ColorMode)
		Debug.Printf("  colorModeDataLen: %d", len(colorModeData))
		Debug.Println("end - header")
	}

	if cfg.Res, l, err = readImageResource(r); err != nil {
		return Config{}, read, err
	}
	read += l

	switch cfg.ColorMode {
	case ColorModeBitmap:
		cfg.Picker = &pickerGray1{}
	case ColorModeGrayscale:
		cfg.Picker = findGrayPicker(cfg.Depth, cfg.Channels > cfg.ColorMode.Channels() && hasAlphaID0(cfg.Res))
	case ColorModeIndexed:
		// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_38034
		// 0x0417(1046) | (Photoshop 6.0) Transparency Index.
		//                2 bytes for the index of transparent color, if any.
		transparentColorIndex := -1
		if r, ok := cfg.Res[0x0417]; ok {
			transparentColorIndex = int(readUint16(r.Data, 0))
		}
		pal := make(color.Palette, len(colorModeData)/3)
		for i := range pal {
			c := color.NRGBA{
				colorModeData[i],
				colorModeData[i+256],
				colorModeData[i+512],
				0xFF,
			}
			if i == transparentColorIndex {
				c.A = 0
			}
			pal[i] = c
		}
		cfg.Picker = &pickerPalette{Palette: pal}
	case ColorModeRGB:
		cfg.Picker = findNRGBAPicker(cfg.Depth, cfg.Channels > cfg.ColorMode.Channels() && hasAlphaID0(cfg.Res))
	case ColorModeCMYK:
		cfg.Picker = findNCMYKAPicker(cfg.Depth, cfg.Channels > cfg.ColorMode.Channels() && hasAlphaID0(cfg.Res))
	case ColorModeMultichannel, ColorModeDuotone, ColorModeLab:
		return Config{}, read, errors.New("psd: color mode " + itoa(int(cfg.ColorMode)) + " is not implemented")
	default:
		return Config{}, read, errors.New("psd: unexpected color mode: " + itoa(int(cfg.ColorMode)))
	}
	if cfg.Picker == nil {
		return Config{}, read, errors.New("psd: unsupported combination of color mode(" + itoa(int(cfg.ColorMode)) + ") and depth(" + itoa(cfg.Depth) + ")")
	}
	return cfg, read, nil
}

// DecodeOptions are the decoding options.
type DecodeOptions struct {
	SkipLayerImage  bool
	SkipMergedImage bool
}

// Decode reads a image from r and returns it.
// Default parameters are used if a nil *DecodeOptions is passed.
func Decode(r io.Reader, o *DecodeOptions) (psd *PSD, read int, err error) {
	if o == nil {
		o = &DecodeOptions{}
	}
	psd = &PSD{}
	var l int
	if psd.Config, l, err = DecodeConfig(r); err != nil {
		return nil, 0, err
	}
	read += l

	b := make([]byte, 4)

	if l, err = readLayerAndMaskInfo(r, psd, o); err != nil {
		return nil, read, err
	}
	read += l

	if !o.SkipMergedImage {
		// allocate image buffer.
		// `+7` and `>>3` is the padding for the 1bit depth color mode.
		plane := (psd.Config.Rect.Dx()*psd.Config.Depth + 7) >> 3 * psd.Config.Rect.Dy()
		psd.Data = make([]byte, plane*psd.Config.Channels)
		chs := make([][]byte, psd.Config.Channels)
		psd.Channel = make(map[int]Channel)
		for i := 0; i < psd.Config.Channels; i++ {
			psd.Channel[i] = Channel{
				Data:   psd.Data[plane*i : plane*(i+1) : plane*(i+1)],
				Picker: findGrayPicker(psd.Config.Depth, false),
			}
			psd.Channel[i].Picker.SetSource(psd.Config.Rect, psd.Channel[i].Data)
			chs[i] = psd.Channel[i].Data
		}
		psd.Config.Picker.SetSource(psd.Config.Rect, chs...)

		if l, err = io.ReadFull(r, b[:2]); err != nil {
			return nil, read, err
		}
		read += l
		cmpMethod := CompressionMethod(readUint16(b, 0))
		l, err = cmpMethod.Decode(
			psd.Data,
			r,
			0,
			psd.Config.Rect,
			psd.Config.Depth,
			psd.Config.Channels,
		)
		if err != nil {
			return nil, read, err
		}
		read += l
	}
	return psd, read, nil
}

func init() {
	image.RegisterFormat(
		"psd",
		"8BPS\x00\x01",
		func(r io.Reader) (image.Image, error) {
			psd, _, err := Decode(r, &DecodeOptions{SkipLayerImage: true})
			return psd, err
		},
		func(r io.Reader) (image.Config, error) {
			cfg, _, err := DecodeConfig(r)
			if err != nil {
				return image.Config{}, err
			}
			return image.Config{
				Width:      cfg.Rect.Dx(),
				Height:     cfg.Rect.Dy(),
				ColorModel: cfg.Picker.ColorModel(),
			}, nil
		},
	)
}
