// +build ignore

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
)

type name string

func (n name) Lower() name {
	if len(n) == 0 {
		return ""
	}
	return name(strings.ToLower(string(n[:1])) + string(n[1:]))
}

type code string

func (c code) Channel(ch string) code {
	return code(strings.NewReplacer(
		"src", "s"+ch,
		"dest", "d"+ch,
		"ret", ch,
	).Replace(string(c)))
}

func (c code) To16() code {
	return code(strings.NewReplacer(
		"maxValue", "0xffff",
		"halfValue", "0x8000",
		"quarterValue", "0x4000",
		"lum(", "lum16(",
		"setLum(", "setLum16(",
		"clip(", "clip16(",
	).Replace(string(c)))
}

func (c code) To8() code {
	return code(strings.NewReplacer(
		"maxValue", "0xff",
		"halfValue", "0x80",
		"quarterValue", "0x40",
		"lum(", "lum8(",
		"setLum(", "setLum8(",
		"clip(", "clip8(",
	).Replace(string(c)))
}

type mode struct {
	Name             name
	OverMax          bool
	Code             code
	CodePerChannel   code
	Code16           code
	CodePerChannel16 code
}

var blendModes = []mode{
	// ----------------------------------------------------------------
	// References:
	// https://www.w3.org/TR/compositing-1/#blending
	// http://dunnbypaul.net/blends/
	// https://mouaif.wordpress.com/2009/01/05/photoshop-math-with-glsl-shaders/
	// ----------------------------------------------------------------
	{
		Name: "Normal",
		CodePerChannel: `
			ret = src;
		`,
	},
	// ----------------------------------------------------------------
	{
		Name: "Darken",
		CodePerChannel: `
			if (src < dest) {
				ret = src;
			} else {
				ret = dest;
			}
		`,
	},
	{
		Name: "Multiply",
		CodePerChannel: `
			ret = src * dest * 32897 >> 23;
		`,
		CodePerChannel16: `
			ret = src * dest / maxValue;
		`,
	},
	{
		Name: "ColorBurn",
		CodePerChannel: `
			if (dest == maxValue) {
				ret = maxValue;
			} else if (src == 0) {
				ret = 0;
			} else {
				ret = maxValue - clip((maxValue - dest) * maxValue / src);
			}
		`,
	},
	{
		Name: "LinearBurn",
		CodePerChannel: `
			tmp = dest + src;
			if (tmp > maxValue) {
				ret = tmp - maxValue;
			} else {
				ret = 0;
			}
		`,
	},
	{
		Name: "DarkerColor",
		Code: `
			if (lum(sr, sg, sb) < lum(dr, dg, db)) {
				r = sr;
				g = sg;
				b = sb;
			} else {
				r = dr;
				g = dg;
				b = db;
			}
		`,
	},
	// ----------------------------------------------------------------
	{
		Name: "Lighten",
		CodePerChannel: `
			if (src > dest) {
				ret = src;
			} else {
				ret = dest;
			}
		`,
	},
	{
		Name: "Screen",
		CodePerChannel: `
			ret = src + dest - (src * dest * 32897 >> 23);
		`,
		CodePerChannel16: `
			ret = src + dest - (src * dest / maxValue);
		`,
	},
	{
		Name: "ColorDodge",
		CodePerChannel: `
			if (dest == 0) {
				ret = 0;
			} else if (src == maxValue) {
				ret = maxValue;
			} else {
				ret = clip(dest * maxValue / (maxValue - src));
			}
		`,
	},
	{
		Name: "LinearDodge",
		CodePerChannel: `
			ret = clip(src + dest);
		`,
	},
	{
		Name: "LighterColor",
		Code: `
			if (lum(sr, sg, sb) > lum(dr, dg, db)) {
				r = sr;
				g = sg;
				b = sb;
			} else {
				r = dr;
				g = dg;
				b = db;
			}
		`,
	},
	{
		Name:    "Add",
		OverMax: true,
		CodePerChannel: `
			ret = src + dest;
		`,
	},
	// ----------------------------------------------------------------
	{
		Name: "Overlay",
		CodePerChannel: `
			if (dest < halfValue) {
				ret = src * dest * 32897 >> 22;
			} else {
				ret = maxValue - ((maxValue - ((dest - halfValue) << 1)) * (maxValue - src) * 32897 >> 23);
			}
		`,
		CodePerChannel16: `
			if (dest < halfValue) {
				ret = src * dest / halfValue;
			} else {
				ret = maxValue - ((maxValue - ((dest - halfValue) << 1)) * (maxValue - src) / maxValue);
			}
		`,
	},
	{
		Name: "SoftLight",
		CodePerChannel: `
			if (src < halfValue) {
				ret = dest - (((maxValue - (src << 1)) * dest * 32897 >> 23) * (maxValue - dest) * 32897 >> 23);
			} else {
				if (dest < quarterValue) {
					tmp = uint32((((int32(dest)<<4 - maxValue*12) * 32897 >> 23) * int32(dest) + maxValue*4) * int32(dest) * 32897 >> 23);
				} else {
					tmp = sqrt(dest, maxValue);
				}
				ret = dest + (((src << 1) - maxValue) * (tmp - dest) * 32897 >> 23);
			}
		`,
		CodePerChannel16: `
			if (src < halfValue) {
				ret = dest - (((maxValue - (src << 1)) * dest / maxValue) * (maxValue - dest) / maxValue);
			} else {
				if (dest < quarterValue) {
					tmp = uint32((((int32(dest)<<4 - maxValue*12) / maxValue) * int32(dest) + maxValue*4) * int32(dest) / maxValue);
				} else {
					tmp = sqrt(dest, maxValue);
				}
				ret = dest + (((src << 1) - maxValue) * (tmp - dest) / maxValue);
			}
		`,
	},
	{
		Name: "HardLight",
		CodePerChannel: `
			if (src < halfValue) {
				ret = dest * src * 32897 >> 22;
			} else {
				tmp = (src << 1) - maxValue;
				ret = dest + tmp - (dest * tmp * 32897 >> 23);
			}
		`,
		CodePerChannel16: `
			if (src < halfValue) {
				ret = dest * src / halfValue;
			} else {
				tmp = (src << 1) - maxValue;
				ret = dest + tmp - (dest * tmp / maxValue);
			}
		`,
	},
	{
		Name: "LinearLight",
		CodePerChannel: `
			if (src < halfValue) {
				ret = clip0(dest + (src << 1) - maxValue);
			} else {
				ret = clip(dest + ((src - halfValue) << 1));
			}
		`,
	},
	{
		Name: "VividLight",
		CodePerChannel: `
			if (src < halfValue) {
				tmp = src << 1;
					if (tmp == 0) {
					ret = 0;
				} else {
					ret = maxValue - clip((maxValue - dest) * maxValue / tmp);
				}
			} else {
				tmp = (src << 1) - maxValue;
					if (tmp == maxValue) {
					ret = maxValue;
				} else {
					ret = clip((dest * maxValue) / (maxValue - tmp));
				}
			}
		`,
	},
	{
		Name: "PinLight",
		CodePerChannel: `
			if (src < halfValue) {
				tmp = src << 1;
				if (tmp < dest) {
				  ret = tmp;
				} else {
				  ret = dest;
				}
			} else {
				tmp = (src - halfValue) << 1;
				if (tmp > dest) {
				  ret = tmp;
				} else {
				  ret = dest;
				}
			}
		`,
	},
	{
		Name: "HardMix",
		CodePerChannel: `
			if (src < halfValue) {
				tmp = src << 1;
				if (dest == maxValue) {
					tmp = maxValue;
				} else if (tmp == 0) {
					tmp = 0;
				} else {
					tmp = maxValue - clip((maxValue - dest) * maxValue / tmp);
				}
			} else {
				tmp = ((src - halfValue) << 1);
				if (dest == 0) {
					tmp = 0;
				} else if (tmp == maxValue) {
					tmp = maxValue;
				} else {
					tmp = clip(dest * maxValue / (maxValue - tmp));
				}
			}
			if (tmp <= halfValue) {
				ret = 0;
			} else {
				ret = maxValue;
			}
		`,
	},
	// ----------------------------------------------------------------
	{
		Name: "Difference",
		CodePerChannel: `
			if (dest < src) {
				ret = src - dest;
			} else {
				ret = dest - src;
			}
		`,
	},
	{
		Name: "Exclusion",
		CodePerChannel: `
			ret = dest + src - (dest * src * 32897 >> 22);
		`,
		CodePerChannel16: `
			ret = dest + src - (dest * src / halfValue);
		`,
	},
	{
		Name: "Subtract",
		CodePerChannel: `
			if (dest < src) {
				ret = 0;
			} else {
				ret = dest - src;
			}
		`,
	},
	{
		Name: "Divide",
		CodePerChannel: `
			if (dest == 0) {
				ret = 0;
			} else if (src == 0) {
				ret = maxValue;
			} else {
				ret = clip(dest * maxValue / src);
			}
		`,
	},
	// ----------------------------------------------------------------
	{
		Name: "Hue",
		Code: `
			r, g, b = setSat(sr, sg, sb, sat(dr, dg, db));
			r, g, b = setLum(r, g, b, lum(dr, dg, db));
		`,
	},
	{
		Name: "Saturation",
		Code: `
			r, g, b = setSat(dr, dg, db, sat(sr, sg, sb));
			r, g, b = setLum(r, g, b, lum(dr, dg, db));
		`,
	},
	{
		Name: "Color",
		Code: `
			r, g, b = setLum(sr, sg, sb, lum(dr, dg, db));
		`,
	},
	{
		Name: "Luminosity",
		Code: `
			r, g, b = setLum(dr, dg, db, lum(sr, sg, sb));
		`,
	},
}

var source = `// DO NOT EDIT.
// Generate with: go generate

package blend

import (
	"image"
	stdcolor "image/color"
	"image/draw"
)

// blend modes
var (
{{range .blendModes}}	{{.Name}} Drawer = {{.Name.Lower}}{}
{{end}}
)

{{range .blendModes}}{{template "blendBase" .}}{{end}}
`

var blendBase = `
// {{.Name.Lower}} implements the {{.Name.Lower}} blend mode.
type {{.Name.Lower}} struct {}

// String implemenets fmt.Stringer interface.
func (d {{.Name.Lower}}) String() string {
	return {{printf "%q" .Name}}
}

// Draw implements image.Drawer interface.
func (d {{.Name.Lower}}) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d {{.Name.Lower}}) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d {{.Name.Lower}}) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {
{{define "draw"}}
	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		alpha >>= 8
		if alpha == 0 {
			return
		}
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		dyDelta, syDelta int
		x0, x1, xDelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		dyDelta = dst.Stride
		syDelta = src.Stride
		x0, x1, xDelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		dyDelta = -dst.Stride
		syDelta = -src.Stride
		x0, x1, xDelta = (dx-1)<<2, -4, -4
	}
	{{.}}.Parallel(dst.Pix[d0:], src.Pix[s0:], alpha, dy, x0, x1, xDelta, syDelta, x0, x1, xDelta, dyDelta)
{{end}}
{{template "draw" printf "draw%sNRGBAToNRGBA" .Name}}
}

func (d {{.Name.Lower}}) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {
{{template "draw" printf "draw%sRGBAToNRGBA" .Name}}
}

func (d {{.Name.Lower}}) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {
{{template "draw" printf "draw%sNRGBAToRGBA" .Name}}
}

func (d {{.Name.Lower}}) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {
{{template "draw" printf "draw%sRGBAToRGBA" .Name}}
}

var draw{{.Name}}NRGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {
{{define "drawMain1"}}
	dPos, sPos := 0, 0
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
		spix := src[sPos:]
		for i, j := sx0, dx0; i != sx1; i, j = i+sxDelta, j+dxDelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
			tmp := (sa * alpha >> 23) * 32897
			if tmp == 0 {
				continue
			}

			da := uint32(dpix[j+3])
			db := uint32(dpix[j+2])
			dg := uint32(dpix[j+1])
			dr := uint32(dpix[j])

			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}
{{end}}
{{define "drawMain1_SrcRGBAToNRGBA"}}
			if sa < 0xff {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}
{{end}}
{{define "drawMain1_DestRGBAToNRGBA"}}
			if 0x00 < da && da < 0xff {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}
{{end}}
{{define "drawMain2"}}
			var r, g, b uint32
			{{if .CodePerChannel}}
				{{.CodePerChannel.To8.Channel "r"}}
				{{.CodePerChannel.To8.Channel "g"}}
				{{.CodePerChannel.To8.Channel "b"}}
			{{else if .Code}}
				{{.Code.To8}}
			{{else if .CodePerChannel16}}
				{{.CodePerChannel16.To8.Channel "r"}}
				{{.CodePerChannel16.To8.Channel "g"}}
				{{.CodePerChannel16.To8.Channel "b"}}
			{{else if .Code16}}
				{{.Code16.To8}}
			{{end}}
{{end}}
{{define "drawMain2_SetNRGBA"}}
{{if .OverMax}}
			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a))
			dpix[j+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a))
			dpix[j+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a))
{{else}}
			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)
{{end}}
{{end}}
{{define "drawMain2_SetRGBA"}}
{{if .OverMax}}
			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8(clip8((b*a1 + sb*a2 + db*a3) / a) * a * 32897 >> 23)
			dpix[j+1] = uint8(clip8((g*a1 + sg*a2 + dg*a3) / a) * a * 32897 >> 23)
			dpix[j+0] = uint8(clip8((r*a1 + sr*a2 + dr*a3) / a) * a * 32897 >> 23)
{{else}}
			dpix[j+3] = uint8(a)
			dpix[j+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[j+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[j+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)
{{end}}
{{end}}
{{define "drawMain3"}}
		}
		dPos += dyDelta
		sPos += syDelta
	}
{{end}}
{{template "drawMain1" .}}
{{template "drawMain2" .}}
{{template "drawMain2_SetNRGBA" .}}
{{template "drawMain3" .}}
}

var draw{{.Name}}RGBAToNRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {
	{{template "drawMain1" .}}
	{{template "drawMain1_SrcRGBAToNRGBA" .}}
	{{template "drawMain2" .}}
	{{template "drawMain2_SetNRGBA" .}}
	{{template "drawMain3" .}}
}

var draw{{.Name}}NRGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {
	{{template "drawMain1" .}}
	{{template "drawMain1_DestRGBAToNRGBA" .}}
	{{template "drawMain2" .}}
	{{template "drawMain2_SetRGBA" .}}
	{{template "drawMain3" .}}
}

var draw{{.Name}}RGBAToRGBA drawFunc = func(dest []byte, src []byte, alpha uint32, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int) {
	{{template "drawMain1" .}}
	{{template "drawMain1_SrcRGBAToNRGBA" .}}
	{{template "drawMain1_DestRGBAToNRGBA" .}}
	{{template "drawMain2" .}}
	{{template "drawMain2_SetRGBA" .}}
	{{template "drawMain3" .}}
}

var draw{{.Name}}Fallback drawFallbackFunc = func(dst draw.Image, pX int, pY int, src image.Image, spX int, spY int, mask image.Image, mpX int, mpY int, endX int, endY int, dx int, dy int) {
	var out stdcolor.RGBA64
	for y, sy, my := pY, spY, mpY; y != endY; y, sy, my = y+dy, sy+dy, my+dy {
		for x, sx, mx := pX, spX, mpX; x != endX; x, sx, mx = x+dx, sx+dx, mx+dx {
			ma := uint32(0xffff)
			if mask != nil {
				_, _, _, ma = mask.At(mx, my).RGBA()
			}
			if ma == 0 {
				continue
			}

			sr, sg, sb, sa := src.At(sx, sy).RGBA()
			dr, dg, db, da := dst.At(x, y).RGBA()

			tmp := sa * ma / 0xffff
			a1 := tmp * da / 0xffff
			a2 := tmp * (0xffff - da) / 0xffff
			a3 := (0xffff - tmp) * da / 0xffff
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}

			if sa == 0x0000 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 0xffff {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da == 0x0000 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 0xffff {
				dr = dr * 0xffff / da
				dg = dg * 0xffff / da
				db = db * 0xffff / da
			}

		 var r, g, b uint32
{{if .CodePerChannel16}}
{{.CodePerChannel16.To16.Channel "r"}}
{{.CodePerChannel16.To16.Channel "g"}}
{{.CodePerChannel16.To16.Channel "b"}}
{{else if .Code16}}
{{.Code16.To16}}
{{else if .CodePerChannel}}
{{.CodePerChannel.To16.Channel "r"}}
{{.CodePerChannel.To16.Channel "g"}}
{{.CodePerChannel.To16.Channel "b"}}
{{else if .Code}}
{{.Code.To16}}
{{end}}
{{if .OverMax}}
			out.R = uint16(clip6416(uint64(sr*a2 + dr*a3)/0xffff + uint64(r)*uint64(a1) / uint64(a)) * a / 0xffff)
			out.G = uint16(clip6416(uint64(sg*a2 + dg*a3)/0xffff + uint64(g)*uint64(a1) / uint64(a)) * a / 0xffff)
			out.B = uint16(clip6416(uint64(sb*a2 + db*a3)/0xffff + uint64(b)*uint64(a1) / uint64(a)) * a / 0xffff)
			out.A = uint16(a)
{{else}}
			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)
{{end}}
			dst.Set(x, y, &out)
		}
	}
}

func (d {{.Name.Lower}}) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	pX, pY, endX, endY, delta := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		pX, pY, endX, endY, delta = r.Max.X-1, r.Max.Y-1, r.Min.X-1, r.Min.Y-1, -1
	}
	ofsX, ofsY := pX - r.Min.X, pY - r.Min.Y
	draw{{.Name}}Fallback.Parallel(
		dst, pX, pY,
		src, sp.X + ofsX, sp.Y + ofsY,
		mask, mp.X + ofsX, mp.Y + ofsY,
		endX, endY, delta, delta,
	)
}
`

var testSource = `// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

{{range .}}
	func TestBlendFallback{{.Name}}(t *testing.T) {
		testDrawFallback(t, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func TestBlendNRGBAToNRGBA{{.Name}}(t *testing.T) {
		testDrawNRGBAToNRGBA(t, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func TestBlendRGBAToNRGBA{{.Name}}(t *testing.T) {
		testDrawRGBAToNRGBA(t, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func TestBlendNRGBAToRGBA{{.Name}}(t *testing.T) {
		testDrawNRGBAToRGBA(t, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func TestBlendRGBAToRGBA{{.Name}}(t *testing.T) {
		testDrawRGBAToRGBA(t, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func BenchmarkBlendFallback{{.Name}}(b *testing.B) {
		benchmarkDrawFallback(b, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func BenchmarkBlendNRGBAToNRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawNRGBAToNRGBA(b, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func BenchmarkBlendRGBAToNRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawRGBAToNRGBA(b, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func BenchmarkBlendNRGBAToRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawNRGBAToRGBA(b, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}

	func BenchmarkBlendRGBAToRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawRGBAToRGBA(b, "png/bg.png", "png/fg.png", {{.Name.Lower}}{})
	}
{{end}}
`

func main() {
	t := template.Must(template.New("").Parse(source))
	template.Must(t.New("blendBase").Parse(blendBase))

	b := bytes.NewBufferString("")
	if err := t.Execute(b, map[string]interface{}{
		"blendModes": blendModes,
	}); err != nil {
		log.Fatal(err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("blends.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write(buf); err != nil {
		log.Fatal(err)
	}

	tt := template.Must(template.New("").Parse(testSource))
	b.Reset()
	if err = tt.Execute(b, blendModes); err != nil {
		log.Fatal(err)
	}
	buf, err = format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	f2, err := os.Create("blends_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	if _, err = f2.Write(buf); err != nil {
		log.Fatal(err)
	}
}
