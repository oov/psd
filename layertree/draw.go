package layertree

import (
	"image"
	"image/color"

	"github.com/oov/psd"
	"github.com/oov/psd/blend"
)

var blendMode = map[psd.BlendMode]blend.Drawer{
	psd.BlendModeNormal:       blend.Normal,
	psd.BlendModeDarken:       blend.Darken,
	psd.BlendModeMultiply:     blend.Multiply,
	psd.BlendModeColorBurn:    blend.ColorBurn,
	psd.BlendModeLinearBurn:   blend.LinearBurn,
	psd.BlendModeDarkerColor:  blend.DarkerColor,
	psd.BlendModeLighten:      blend.Lighten,
	psd.BlendModeScreen:       blend.Screen,
	psd.BlendModeColorDodge:   blend.ColorDodge,
	psd.BlendModeLinearDodge:  blend.Add,
	psd.BlendModeLighterColor: blend.LighterColor,
	psd.BlendModeOverlay:      blend.Overlay,
	psd.BlendModeSoftLight:    blend.SoftLight,
	psd.BlendModeHardLight:    blend.HardLight,
	psd.BlendModeLinearLight:  blend.LinearLight,
	psd.BlendModeVividLight:   blend.VividLight,
	psd.BlendModePinLight:     blend.PinLight,
	psd.BlendModeHardMix:      blend.HardMix,
	psd.BlendModeDifference:   blend.Difference,
	psd.BlendModeExclusion:    blend.Exclusion,
	psd.BlendModeSubtract:     blend.Subtract,
	psd.BlendModeDivide:       blend.Divide,
	psd.BlendModeHue:          blend.Hue,
	psd.BlendModeSaturation:   blend.Saturation,
	psd.BlendModeColor:        blend.Color,
	psd.BlendModeLuminosity:   blend.Luminosity,
}

func (r *Renderer) draw(b *image.RGBA, src *image.RGBA, opacity int, bm psd.BlendMode) {
	blendMode[bm].DrawMask(
		b,
		src.Rect,
		src,
		src.Rect.Min,
		image.NewUniform(color.Alpha{uint8(opacity)}),
		image.Point{},
	)
}

func (r *Renderer) writeAlpha(b *image.RGBA, alpha byte) {
	d := b.Pix
	ln := len(d)
	_ = d[ln-1]
	if alpha == 255 {
		for i := 0; i < ln; i += 4 {
			if a := uint32(d[i+3]); a > 0 {
				d[i+0] = uint8(uint32(d[i+0]) * 0xff / a)
				d[i+1] = uint8(uint32(d[i+1]) * 0xff / a)
				d[i+2] = uint8(uint32(d[i+2]) * 0xff / a)
			}
			d[i+3] = 255
		}
	} else {
		for i := 3; i < ln; i += 4 {
			d[i] = alpha
		}
	}
}

func (r *Renderer) applyAlpha(b *image.RGBA, src *image.RGBA) {
	d, s := b.Pix, src.Pix
	ln := len(d)
	_, _ = d[ln-1], s[ln-1]
	for i := 0; i < ln; i += 4 {
		a := uint32(s[i+3])
		d[i+0] = uint8(uint32(d[i+0]) * a * 32768 >> 23)
		d[i+1] = uint8(uint32(d[i+1]) * a * 32768 >> 23)
		d[i+2] = uint8(uint32(d[i+2]) * a * 32768 >> 23)
		d[i+3] = s[i+3]
	}
}
