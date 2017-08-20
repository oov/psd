package layertree

import (
	"image"
	"log"
	"strings"

	"golang.org/x/text/encoding"

	"github.com/oov/psd"
)

type builder struct {
	Img                       *psd.PSD
	Rect                      image.Rectangle
	LayerNameEncodingDetector func([]byte) encoding.Encoding
	LayerNameEncoding         encoding.Encoding
}

func (b *builder) buildLayer(l *Layer, psdl *psd.Layer) {
	l.SeqID = psdl.SeqID

	if psdl.UnicodeName == "" && psdl.MBCSName != "" {
		n, err := b.getLayerNameEncoding().NewDecoder().String(psdl.MBCSName)
		if err != nil {
			l.Name = psdl.MBCSName
		} else {
			l.Name = n
		}
	} else {
		l.Name = psdl.UnicodeName
	}

	if psdl.Folder() {
		l.BlendMode = psdl.SectionDividerSetting.BlendMode
	} else {
		if psdl.BlendMode == psd.BlendModePassThrough {
			log.Printf("NOTICE: In '%s' layer, blend mode 'pass-through' which is unsupported mode in image layer has been replaced by 'normal'.", l.Name)
			psdl.BlendMode = psd.BlendModeNormal
		}
		l.BlendMode = psdl.BlendMode
	}
	l.Opacity = int(psdl.Opacity)
	l.Clipping = psdl.Clipping
	l.BlendClippedElements = psdl.BlendClippedElements
	l.Visible = psdl.Visible()
	l.Folder = psdl.Folder()
	l.FolderOpen = psdl.FolderIsOpen()

	l.MaskDefaultColor = psdl.Mask.DefaultColor
	l.MaskEnabled = psdl.Mask.Enabled() && !psdl.Mask.Rect.Empty()

	b.Rect = b.Rect.Union(psdl.Rect)
	for i := range psdl.Layer {
		l.Children = append(l.Children, Layer{Parent: l})
		b.buildLayer(&l.Children[i], &psdl.Layer[i])
	}
}

func reverse(ls []*Layer) {
	ln := len(ls) >> 1
	for i := 0; i < ln; i++ {
		j := len(ls) - i - 1
		ls[i], ls[j] = ls[j], ls[i]
	}
}

func (b *builder) registerClippingGroup(ls []Layer) {
	var clip []*Layer
	for i := len(ls) - 1; i >= 0; i-- {
		if len(ls[i].Children) > 0 {
			b.registerClippingGroup(ls[i].Children)
		}
		if ls[i].Clipping {
			clip = append(clip, &ls[i])
			continue
		}
		if len(clip) > 0 {
			for _, cl := range clip {
				cl.ClippedBy = &ls[i]
			}
			reverse(clip)
			ls[i].Clip = clip
			clip = nil
		}
	}
}

func (b *builder) gatherLayerName(l *psd.Layer) []string {
	var s []string
	for i := range l.Layer {
		s = append(s, l.Layer[i].MBCSName)
		if len(l.Layer[i].Layer) > 0 {
			s = append(s, b.gatherLayerName(&l.Layer[i])...)
		}
	}
	return s
}

func (b *builder) getLayerNameEncoding() encoding.Encoding {
	if b.LayerNameEncoding != nil {
		return b.LayerNameEncoding
	}

	var s []string
	for i := range b.Img.Layer {
		s = append(s, b.Img.Layer[i].MBCSName)
		if len(b.Img.Layer[i].Layer) > 0 {
			s = append(s, b.gatherLayerName(&b.Img.Layer[i])...)
		}
	}
	if len(s) > 0 && b.LayerNameEncodingDetector != nil {
		b.LayerNameEncoding = b.LayerNameEncodingDetector([]byte(strings.Join(s, "")))
	} else {
		b.LayerNameEncoding = encoding.Nop
	}
	return b.LayerNameEncoding
}
