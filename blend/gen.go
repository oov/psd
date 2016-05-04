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

var blendModes = []struct {
	Name             string
	Code             code
	CodePerChannel   code
	Code16           code
	CodePerChannel16 code
}{
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
	"image/color"
	"image/draw"

	"github.com/oov/psd"
)

var Mode = map[psd.BlendMode]Drawer{
	psd.BlendModePassThrough: Normal{},
	{{range .}}psd.BlendMode{{.Name}}: {{.Name}}{},
{{end}}
}

{{range .}}{{template "blendBase" .}}{{end}}
`

var blendBase = `
// {{.Name}} implements the {{.Name}} blend mode.
type {{.Name}} struct {}

// String implemenets fmt.Stringer interface.
func (d {{.Name}}) String() string {
	return {{printf "%q" .Name}}
}

// BlendMode returns psd.BlendMode.
func (d {{.Name}}) BlendMode() psd.BlendMode {
	return psd.BlendMode{{.Name}}
}

// Draw implements image.Drawer interface.
func (d {{.Name}}) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	drawMask(d, dst, r, src, sp, nil, image.Point{})
}

// DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r
// in dst with the result. A nil mask is treated as opaque.
func (d {{.Name}}) DrawMask(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	drawMask(d, dst, r, src, sp, mask, mp)
}

func (d {{.Name}}) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {
{{define "blendRGBAToRGBA1"}}
	ma := uint32(0xff)
	if mask != nil {
		_, _, _, ma = mask.C.RGBA()
		if ma == 0 {
			return
		}
		ma >>= 8
	}

	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx<<2, +4
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)<<2, -4, -4
	}

	for ; dy > 0; dy-- {
		dpix := dst.Pix[d0:]
		spix := src.Pix[s0:]
		for i := i0; i != i1; i += idelta {
			sa := uint32(spix[i+3])
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])

			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])

			tmp := (sa * ma * 32897 >> 23) * 32897
			a1 := (tmp * da) >> 23
			a2 := (tmp * (255 - da)) >> 23
			a3 := ((8388735 - tmp) * da) >> 23
			a := a1 + a2 + a3
			if a == 0 {
				continue
			}
{{end}}
{{define "blendRGBAToRGBA2_Src"}}
{{if .}}
			if sa > 0 {
				sr = sr * 0xff / sa
				sg = sg * 0xff / sa
				sb = sb * 0xff / sa
			}
{{end}}
{{end}}
{{define "blendRGBAToRGBA2_Dest"}}
{{if .}}
			if da > 0 {
				dr = dr * 0xff / da
				dg = dg * 0xff / da
				db = db * 0xff / da
			}
{{end}}
{{end}}
{{define "blendRGBAToRGBA3"}}
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
{{define "blendRGBAToRGBA4_Dest"}}
{{if .}}
			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) * 32897 >> 23)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) * 32897 >> 23)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) * 32897 >> 23)
{{else}}
			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8((b*a1 + sb*a2 + db*a3) / a)
			dpix[i+1] = uint8((g*a1 + sg*a2 + dg*a3) / a)
			dpix[i+0] = uint8((r*a1 + sr*a2 + dr*a3) / a)
{{end}}
		}
		d0 += ddelta
		s0 += sdelta
	}
{{end}}
{{template "blendRGBAToRGBA1" .}}
{{template "blendRGBAToRGBA2_Src" true}}
{{template "blendRGBAToRGBA2_Dest" true}}
{{template "blendRGBAToRGBA3" .}}
{{template "blendRGBAToRGBA4_Dest" true}}
}

func (d {{.Name}}) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {
{{template "blendRGBAToRGBA1" .}}
{{template "blendRGBAToRGBA2_Src" false}}
{{template "blendRGBAToRGBA2_Dest" true}}
{{template "blendRGBAToRGBA3" .}}
{{template "blendRGBAToRGBA4_Dest" true}}
}

func (d {{.Name}}) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform) {
{{template "blendRGBAToRGBA1" .}}
{{template "blendRGBAToRGBA2_Src" true}}
{{template "blendRGBAToRGBA2_Dest" false}}
{{template "blendRGBAToRGBA3" .}}
{{template "blendRGBAToRGBA4_Dest" false}}
}

func (d {{.Name}}) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform) {
{{template "blendRGBAToRGBA1" .}}
{{template "blendRGBAToRGBA2_Src" false}}
{{template "blendRGBAToRGBA2_Dest" false}}
{{template "blendRGBAToRGBA3" .}}
{{template "blendRGBAToRGBA4_Dest" false}}
}

func (d {{.Name}}) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out color.RGBA64
	sy := sp.Y + y0 - r.Min.Y
	my := mp.Y + y0 - r.Min.Y
	for y := y0; y != y1; y, sy, my = y+dy, sy+dy, my+dy {
		sx := sp.X + x0 - r.Min.X
		mx := mp.X + x0 - r.Min.X
		for x := x0; x != x1; x, sx, mx = x+dx, sx+dx, mx+dx {
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

			if sa > 0 {
				sr = sr * 0xffff / sa
				sg = sg * 0xffff / sa
				sb = sb * 0xffff / sa
			}
			if da > 0 {
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
			out.R = uint16((r*a1 + sr*a2 + dr*a3) / 0xffff)
			out.G = uint16((g*a1 + sg*a2 + dg*a3) / 0xffff)
			out.B = uint16((b*a1 + sb*a2 + db*a3) / 0xffff)
			out.A = uint16(a)
			dst.Set(x, y, &out)
		}
	}
}
`

var testSource = `// DO NOT EDIT.
// Generate with: go generate

package blend

import "testing"

{{range .}}func TestDrawFallback{{.Name}}(t *testing.T) { testDrawFallback({{.Name}}{}, t) }
{{end}}

{{range .}}func TestDrawNRGBAToNRGBA{{.Name}}(t *testing.T) { testDrawNRGBAToNRGBA({{.Name}}{}, t) }
{{end}}

{{range .}}func TestDrawRGBAToNRGBA{{.Name}}(t *testing.T) { testDrawRGBAToNRGBA({{.Name}}{}, t) }
{{end}}

{{range .}}func TestDrawNRGBAToRGBA{{.Name}}(t *testing.T) { testDrawNRGBAToRGBA({{.Name}}{}, t) }
{{end}}

{{range .}}func TestDrawRGBAToRGBA{{.Name}}(t *testing.T) { testDrawRGBAToRGBA({{.Name}}{}, t) }
{{end}}

{{range .}}func BenchmarkDrawFallback{{.Name}}(b *testing.B) { benchmarkDrawFallback({{.Name}}{}, b) }
{{end}}

{{range .}}func BenchmarkDrawNRGBAToNRGBA{{.Name}}(b *testing.B) { benchmarkDrawNRGBAToNRGBA({{.Name}}{}, b) }
{{end}}

{{range .}}func BenchmarkDrawRGBAToNRGBA{{.Name}}(b *testing.B) { benchmarkDrawRGBAToNRGBA({{.Name}}{}, b) }
{{end}}

{{range .}}func BenchmarkDrawNRGBAToRGBA{{.Name}}(b *testing.B) { benchmarkDrawNRGBAToRGBA({{.Name}}{}, b) }
{{end}}

{{range .}}func BenchmarkDrawRGBAToRGBA{{.Name}}(b *testing.B) { benchmarkDrawRGBAToRGBA({{.Name}}{}, b) }
{{end}}

`

func main() {
	t := template.Must(template.New("").Parse(source))
	template.Must(t.New("blendBase").Parse(blendBase))

	b := bytes.NewBufferString("")
	if err := t.Execute(b, blendModes); err != nil {
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
	if err := tt.Execute(b, blendModes); err != nil {
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
