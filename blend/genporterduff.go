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
	Name     name
	Code     code
	Alpha    code
	Code16   code
	Alpha16  code
	SrcRead  string
	DestRead string
}

var porterDuffModes = []mode{
	// ----------------------------------------------------------------
	// References:
	// https://www.w3.org/TR/compositing-1/#porterduffcompositingoperators
	// ----------------------------------------------------------------
	{
		Name:     "Clear",
		SrcRead:  "skip",
		DestRead: "skip",
	},
	{
		Name:     "Copy",
		DestRead: "skip",
		Alpha: `
			ret = src;
		`,
		Code: `
			ret = src;
		`,
	},
	{
		Name:    "Dest",
		SrcRead: "skip",
		Alpha: `
			ret = dest;
		`,
		Code: `
			ret = dest;
		`,
	},
	{
		Name: "SrcOver",
		Alpha: `
			tmp = maxValue - src;
			ret = src + (tmp * dest * 32768) >> 23;
		`,
		Alpha16: `
			tmp = maxValue - src;
			ret = src + (tmp * dest) / maxValue;
		`,
		Code: `
			ret = src + (tmp * dest * 32768) >> 23;
		`,
		Code16: `
			ret = src + (tmp * dest) / maxValue;
		`,
	},
	{
		Name: "DestOver",
		Alpha: `
			tmp = maxValue - dest;
			ret = dest + (tmp * src * 32768) >> 23;
		`,
		Alpha16: `
			tmp = maxValue - dest;
			ret = dest + (tmp * src) / maxValue;
		`,
		Code: `
			ret = dest + (tmp * src * 32768) >> 23;
		`,
		Code16: `
			ret = dest + (tmp * src) / maxValue;
		`,
	},
	{
		Name:     "SrcIn",
		DestRead: "alpha",
		Alpha: `
			ret = (dest * src * 32768) >> 23;
		`,
		Alpha16: `
			ret = (dest * src) / maxValue;
		`,
		Code: `
			ret = (da * src * 32768) >> 23;
		`,
		Code16: `
			ret = (da * src) / maxValue;
		`,
	},
	{
		Name:    "DestIn",
		SrcRead: "alpha",
		Alpha: `
			ret = (dest * src * 32768) >> 23;
		`,
		Alpha16: `
			ret = (dest * src) / maxValue;
		`,
		Code: `
			ret = (sa * dest * 32768) >> 23;
		`,
		Code16: `
			ret = (sa * dest) / maxValue;
		`,
	},
	{
		Name:     "SrcOut",
		DestRead: "alpha",
		Alpha: `
			tmp = maxValue - dest;
			ret = (tmp * src * 32768) >> 23;
		`,
		Alpha16: `
			tmp = maxValue - dest;
			ret = (tmp * src) / maxValue;
		`,
		Code: `
			ret = (tmp * src * 32768) >> 23;
		`,
		Code16: `
			ret = (tmp * src) / maxValue;
		`,
	},
	{
		Name:    "DestOut",
		SrcRead: "alpha",
		Alpha: `
			tmp = maxValue - src;
			ret = (tmp * dest * 32768) >> 23;
		`,
		Alpha16: `
			tmp = maxValue - src;
			ret = (tmp * dest) / maxValue;
		`,
		Code: `
			ret = (tmp * dest * 32768) >> 23;
		`,
		Code16: `
			ret = (tmp * dest) / maxValue;
		`,
	},
	{
		Name: "SrcAtop",
		Alpha: `
			ret = dest;
			tmp = maxValue - src;
		`,
		Code: `
			ret = (src * da + dest * tmp) * 32768 >> 23;
		`,
		Code16: `
			ret = (src * da + dest * tmp) / maxValue;
		`,
	},
	{
		Name: "DestAtop",
		Alpha: `
			ret = src;
			tmp = maxValue - dest;
		`,
		Code: `
			ret = (src * tmp + dest * sa) * 32768 >> 23;
		`,
		Code16: `
			ret = (src * tmp + dest * sa) / maxValue;
		`,
	},
	{
		Name: "XOR",
		Alpha: `
			ret = (src * (maxValue - dest) + dest * (maxValue - src)) * 32768 >> 23;
			`,
		Alpha16: `
			ret = (src * (maxValue - dest) + dest * (maxValue - src)) / maxValue;
			`,
		Code: `
			ret = (src * (maxValue - da) + dest * (maxValue - sa)) * 32768 >> 23;
		`,
		Code16: `
			ret = (src * (maxValue - da) + dest * (maxValue - sa)) / maxValue;
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

// porter/duff compositing modes
var (
{{range .porterDuffModes}}	{{.Name}} draw.Drawer = {{.Name.Lower}}{}
{{end}}
)

{{range .porterDuffModes}}{{template "porterDuffBase" .}}{{end}}
`

var porterDuffBase = `
// {{.Name.Lower}} implements the {{.Name.Lower}} porter-duff mode.
type {{.Name.Lower}} struct {}

// String implemenets fmt.Stringer interface.
func (d {{.Name.Lower}}) String() string {
	return {{printf "%q" .Name}}
}

// Draw implements image.Drawer interface.
func (d {{.Name.Lower}}) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	// d.drawFallback(dst, r, src, sp, nil, image.Point{}, false)
	drawMask(d, dst, r, src, sp, nil, image.Point{}, false)
}

func (d {{.Name.Lower}}) drawNRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {
{{define "draw1"}}
	alpha := uint32(0xff)
	if mask != nil {
		_, _, _, alpha = mask.C.RGBA()
		if alpha == 0 {
			return
		}
		alpha >>= 8
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
{{end}}
{{define "draw2"}}
	{{.}}(dst.Pix[d0:], src.Pix[s0:], alpha, dy, i0, i1, ddelta, sdelta, idelta)
{{end}}
{{template "draw1" "NRGBAToNRGBA"}}
{{template "draw2" printf "draw%sNRGBAToNRGBA" .Name}}
}

func (d {{.Name.Lower}}) drawRGBAToNRGBAUniform(dst *image.NRGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {
{{template "draw1" "RGBAToNRGBA"}}
{{template "draw2" printf "draw%sRGBAToNRGBA" .Name}}
}

func (d {{.Name.Lower}}) drawNRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.NRGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {
{{template "draw1" "NRGBAToRGBA"}}
{{template "draw2" printf "draw%sNRGBAToRGBA" .Name}}
}

func (d {{.Name.Lower}}) drawRGBAToRGBAUniform(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, mask *image.Uniform, protectAlpha bool) {
{{template "draw1" "RGBAToRGBA"}}
{{template "draw2" printf "draw%sRGBAToRGBA" .Name}}
}

var draw{{.Name}}NRGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {
{{define "drawMain1"}}
	var dPos, sPos int
	alpha *= 32897
	for ; y > 0; y-- {
		dpix := dest[dPos:]
{{if ne .SrcRead "skip"}}
		spix := src[sPos:]
{{end}}
		for i := xMin; i != xMax; i += xDelta {
{{if eq .SrcRead "alpha"}}
			sa := (uint32(spix[i+3]) * alpha) >> 23
{{else if ne .SrcRead "skip"}}
			sa := (uint32(spix[i+3]) * alpha) >> 23
			sb := uint32(spix[i+2])
			sg := uint32(spix[i+1])
			sr := uint32(spix[i])
{{end}}

{{if eq .DestRead "alpha"}}
			da := uint32(dpix[i+3])
{{else if ne .DestRead "skip"}}
			da := uint32(dpix[i+3])
			db := uint32(dpix[i+2])
			dg := uint32(dpix[i+1])
			dr := uint32(dpix[i])
{{end}}
{{end}}
{{define "drawMain1_SrcNRGBAToRGBA"}}
			if sa == 0 {
				sr = 0
				sg = 0
				sb = 0
			} else if sa < 255 {
				sr = (sr * sa * 32897) >> 23
				sg = (sg * sa * 32897) >> 23
				sb = (sb * sa * 32897) >> 23
			}
{{end}}
{{define "drawMain1_DestNRGBAToRGBA"}}
			if da == 0 {
				dr = 0
				dg = 0
				db = 0
			} else if da < 255 {
				dr = (dr * da * 32897) >> 23
				dg = (dg * da * 32897) >> 23
				db = (db * da * 32897) >> 23
			}
{{end}}
{{define "drawMain2"}}
			var r, g, b, a, tmp uint32
			_ = tmp
			{{if .Alpha}}
				{{.Alpha.To8.Channel "a"}}
			{{else if .Alpha16}}
				{{.Alpha16.To8.Channel "a"}}
			{{end}}
			{{if .Code}}
				{{.Code.To8.Channel "r"}}
				{{.Code.To8.Channel "g"}}
				{{.Code.To8.Channel "b"}}
			{{else if .Code16}}
				{{.Code16.To8.Channel "r"}}
				{{.Code16.To8.Channel "g"}}
				{{.Code16.To8.Channel "b"}}
			{{end}}
{{end}}
{{define "drawMain2_SetByRGBA"}}
			dpix[i+3] = uint8(a)
			dpix[i+2] = uint8(b)
			dpix[i+1] = uint8(g)
			dpix[i+0] = uint8(r)
{{end}}
{{define "drawMain2_SetByNRGBA"}}
			dpix[i+3] = uint8(a)
			if a == 255 {
				dpix[i+2] = uint8(b)
				dpix[i+1] = uint8(g)
				dpix[i+0] = uint8(r)
			} else if a == 0 {
				dpix[i+2] = 0
				dpix[i+1] = 0
				dpix[i+0] = 0
			} else {
				dpix[i+2] = uint8(b * 0xff / a)
				dpix[i+1] = uint8(g * 0xff / a)
				dpix[i+0] = uint8(r * 0xff / a)
			}
{{end}}
{{define "drawMain3"}}
		}
		dPos += dDelta
		sPos += sDelta
	}
{{end}}
{{template "drawMain1" .}}
{{if and (ne .SrcRead "skip") (ne .SrcRead "alpha")}}
	{{template "drawMain1_SrcNRGBAToRGBA" .}}
{{end}}
{{if and (ne .DestRead "skip") (ne .DestRead "alpha")}}
	{{template "drawMain1_DestNRGBAToRGBA" .}}
{{end}}
{{template "drawMain2" .}}
{{template "drawMain2_SetByNRGBA" .}}
{{template "drawMain3" .}}
}

var draw{{.Name}}RGBAToNRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {
	{{template "drawMain1" .}}
	{{if and (ne .DestRead "skip") (ne .DestRead "alpha")}}
		{{template "drawMain1_DestNRGBAToRGBA" .}}
	{{end}}
	{{template "drawMain2" .}}
	{{template "drawMain2_SetByNRGBA" .}}
	{{template "drawMain3" .}}
}

var draw{{.Name}}NRGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {
	{{template "drawMain1" .}}
	{{if and (ne .SrcRead "skip") (ne .SrcRead "alpha")}}
		{{template "drawMain1_SrcNRGBAToRGBA" .}}
	{{end}}
	{{template "drawMain2" .}}
	{{template "drawMain2_SetByRGBA" .}}
	{{template "drawMain3" .}}
}

var draw{{.Name}}RGBAToRGBA = func(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int) {
	{{template "drawMain1" .}}
	{{template "drawMain2" .}}
	{{template "drawMain2_SetByRGBA" .}}
	{{template "drawMain3" .}}
}

func (d {{.Name.Lower}}) drawFallback(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, protectAlpha bool) {
	x0, x1, dx := r.Min.X, r.Max.X, 1
	y0, y1, dy := r.Min.Y, r.Max.Y, 1
	if processBackward(dst, r, src, sp) {
		x0, x1, dx = x1-1, x0-1, -1
		y0, y1, dy = y1-1, y0-1, -1
	}

	var out stdcolor.RGBA64
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

{{if eq .SrcRead "alpha"}}
			_, _, _, sa := src.At(sx, sy).RGBA()
{{else if ne .SrcRead "skip"}}
			sr, sg, sb, sa := src.At(sx, sy).RGBA()
{{end}}
{{if eq .DestRead "alpha"}}
			_, _, _, da := dst.At(x, y).RGBA()
{{else if ne .DestRead "skip"}}
			dr, dg, db, da := dst.At(x, y).RGBA()
{{end}}

			var a, r, g, b, tmp uint32
			_ = tmp
{{if .Alpha16}}
	{{.Alpha16.To16.Channel "a"}}
{{else if .Alpha}}
	{{.Alpha.To16.Channel "a"}}
{{end}}
{{if .Code16}}
	{{.Code16.To16.Channel "r"}}
	{{.Code16.To16.Channel "g"}}
	{{.Code16.To16.Channel "b"}}
{{else if .Code}}
	{{.Code.To16.Channel "r"}}
	{{.Code.To16.Channel "g"}}
	{{.Code.To16.Channel "b"}}
{{else if .Code}}
	{{.Code.To16}}
{{end}}
			out.R = uint16(r)
			out.G = uint16(g)
			out.B = uint16(b)
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

{{range .}}
	func TestPorterDuffFallback{{.Name}}(t *testing.T) {
		testDrawFallback(t, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func TestPorterDuffNRGBAToNRGBA{{.Name}}(t *testing.T) {
		testDrawNRGBAToNRGBA(t, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func TestPorterDuffRGBAToNRGBA{{.Name}}(t *testing.T) {
		testDrawRGBAToNRGBA(t, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func TestPorterDuffNRGBAToRGBA{{.Name}}(t *testing.T) {
		testDrawNRGBAToRGBA(t, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func TestPorterDuffRGBAToRGBA{{.Name}}(t *testing.T) {
		testDrawRGBAToRGBA(t, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func BenchmarkPorterDuffFallback{{.Name}}(b *testing.B) {
		benchmarkDrawFallback(b, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func BenchmarkPorterDuffNRGBAToNRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawNRGBAToNRGBA(b, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func BenchmarkPorterDuffRGBAToNRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawRGBAToNRGBA(b, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func BenchmarkPorterDuffNRGBAToRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawNRGBAToRGBA(b, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}

	func BenchmarkPorterDuffRGBAToRGBA{{.Name}}(b *testing.B) {
		benchmarkDrawRGBAToRGBA(b, "png/a.png", "png/b.png", {{.Name.Lower}}{}, false)
	}
{{end}}
`

func main() {
	t := template.Must(template.New("").Parse(source))
	template.Must(t.New("porterDuffBase").Parse(porterDuffBase))

	b := bytes.NewBufferString("")
	if err := t.Execute(b, map[string]interface{}{
		"porterDuffModes": porterDuffModes,
	}); err != nil {
		log.Fatal(err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("porterduffs.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write(buf); err != nil {
		log.Fatal(err)
	}

	tt := template.Must(template.New("").Parse(testSource))
	b.Reset()
	if err = tt.Execute(b, porterDuffModes); err != nil {
		log.Fatal(err)
	}
	buf, err = format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	f2, err := os.Create("porterduffs_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	if _, err = f2.Write(buf); err != nil {
		log.Fatal(err)
	}
}
