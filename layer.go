package psd

import (
	"errors"
	"image"
	"image/color"
	"io"
)

// Layer represents layer.
type Layer struct {
	SeqID int

	Name        string
	UnicodeName string
	MBCSName    string

	Rect image.Rectangle
	// Channel key is the channel ID.
	// 	0 = red, 1 = green, etc.
	// 	-1 = transparency mask
	// 	-2 = user supplied layer mask
	// 	-3 = real user supplied layer mask
	//         (when both a user mask and a vector mask are present)
	Channel               map[int]Channel
	BlendMode             BlendMode
	Opacity               uint8
	Clipping              bool
	BlendClippedElements  bool
	Flags                 uint8
	Mask                  Mask
	Picker                Picker
	AdditionalLayerInfo   map[AdditionalInfoKey][]byte
	SectionDividerSetting struct {
		Type      int
		BlendMode BlendMode
		SubType   int
	}

	Layer []Layer
}

// ExperimentalExportImage creates image.Image.
func (l *Layer) ExperimentalExportImage() (image.Image, error) {
	w, h := l.Rect.Dx(), l.Rect.Dy()
	if l.Picker.ColorModel() == color.NRGBAModel {
		r, g, b := l.Channel[0], l.Channel[1], l.Channel[2]
		img := image.NewNRGBA(l.Rect)
		if a, ok := l.Channel[-1]; ok {
			rp, gp, bp, ap := r.Data, g.Data, b.Data, a.Data
			var ofsd, ofss, x, y, sx, dx int
			for y = 0; y < h; y++ {
				ofss = y * w
				ofsd = ofss << 2
				for x = 0; x < w; x++ {
					sx, dx = ofss+x, ofsd+x<<2
					img.Pix[dx+0] = rp[sx]
					img.Pix[dx+1] = gp[sx]
					img.Pix[dx+2] = bp[sx]
					img.Pix[dx+3] = ap[sx]
				}
			}
		} else {
			rp, gp, bp := r.Data, g.Data, b.Data
			var ofsd, ofss, x, y, sx, dx int
			for y = 0; y < h; y++ {
				ofss = y * w
				ofsd = ofss << 2
				for x = 0; x < w; x++ {
					sx, dx = ofss+x, ofsd+x<<2
					img.Pix[dx+0] = rp[sx]
					img.Pix[dx+1] = gp[sx]
					img.Pix[dx+2] = bp[sx]
					img.Pix[dx+3] = 0xff
				}
			}
		}
		return img, nil
	}
	return nil, errors.New("psd: unsupported color mode")
}

// TransparencyProtected returns whether the layer's transparency being protected.
func (l *Layer) TransparencyProtected() bool {
	return l.Flags&1 != 0
}

// Visible returns whether the layer is visible.
func (l *Layer) Visible() bool {
	return l.Flags&2 == 0
}

// HasImage returns whether the layer has an image data.
func (l *Layer) HasImage() bool {
	return l.SectionDividerSetting.Type == 0
}

// Folder returns whether the layer is folder(group layer).
func (l *Layer) Folder() bool {
	return l.SectionDividerSetting.Type == 1 || l.SectionDividerSetting.Type == 2
}

// FolderIsOpen returns whether the folder is opened when layer is folder.
func (l *Layer) FolderIsOpen() bool {
	return l.SectionDividerSetting.Type == 1
}

// ColorModel returns the Image's color model.
func (l *Layer) ColorModel() color.Model {
	return l.Picker.ColorModel()
}

// Bounds returns the domain for which At can return non-zero color.
func (l *Layer) Bounds() image.Rectangle {
	return l.Picker.Bounds()
}

// At returns the color of the pixel at (x, y).
func (l *Layer) At(x, y int) color.Color {
	return l.Picker.At(x, y)
}

func (l *Layer) String() string {
	return l.Name
}

// Mask represents layer mask and vector mask.
type Mask struct {
	Rect         image.Rectangle
	DefaultColor int
	Flags        int

	RealRect            image.Rectangle
	RealBackgroundColor int
	RealFlags           int

	UserMaskDensity   int
	UserMaskFeather   float64
	VectorMaskDensity int
	VectorMaskFeather float64
}

// Enabled returns whether the mask is enabled.
func (m *Mask) Enabled() bool {
	return m.Flags&2 == 0
}

// RealEnabled returns whether the user real mask is enabled.
func (m *Mask) RealEnabled() bool {
	return m.RealFlags&2 == 0
}

// Channel represents a channel of the color.
type Channel struct {
	// Data is uncompressed channel image data.
	Data   []byte
	Picker Picker
}

// ColorModel returns the color model.
func (c *Channel) ColorModel() color.Model {
	return c.Picker.ColorModel()
}

// Bounds returns the domain for which At can return non-zero color.
func (c *Channel) Bounds() image.Rectangle {
	return c.Picker.Bounds()
}

// At returns the color of the pixel at (x, y).
func (c *Channel) At(x, y int) color.Color {
	return c.Picker.At(x, y)
}

func readLayerAndMaskInfo(r io.Reader, psd *PSD, o *DecodeOptions) (read int, err error) {
	if Debug != nil {
		Debug.Println("start - layer and mask information section")
	}
	// Layer and Mask Information Section
	// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_75067
	b := make([]byte, 4)
	var l int
	if l, err = io.ReadFull(r, b); err != nil {
		return 0, err
	}
	read += l
	layerAndMaskInfoLen := int(readUint32(b, 0))
	if layerAndMaskInfoLen == 0 {
		return read, nil
	}
	if Debug != nil {
		Debug.Println("  layerAndMaskInfoLen:", layerAndMaskInfoLen)
		reportReaderPosition("  file offset: 0x%08x", r)
	}

	var layer []Layer
	if read < layerAndMaskInfoLen+4 {
		if layer, l, err = readLayerInfo(r, psd.Config.ColorMode, psd.Config.Depth, o); err != nil {
			return read, err
		}
		read += l
	}

	if read < layerAndMaskInfoLen+4 {
		// Global layer mask info
		// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_17115
		if l, err = io.ReadFull(r, b); err != nil {
			return read, err
		}
		read += l
		if globalLayerMaskInfoLen := int(readUint32(b, 0)); globalLayerMaskInfoLen > 0 {
			if Debug != nil {
				Debug.Println("  globalLayerMaskInfoLen:", globalLayerMaskInfoLen)
				reportReaderPosition("    file offset: 0x%08x", r)
			}
			// TODO(oov): implement
			if l, err = io.ReadFull(r, make([]byte, globalLayerMaskInfoLen)); err != nil {
				return read, err
			}
			read += l
		}
	}

	if read < layerAndMaskInfoLen+4 {
		var layer2 []Layer
		if psd.AdditinalLayerInfo, layer2, l, err = readAdditionalLayerInfo(r, layerAndMaskInfoLen+4-read, psd.Config.ColorMode, psd.Config.Depth, o); err != nil {
			return read, err
		}
		read += l
		if layer != nil && layer2 != nil {
			return read, errors.New("psd: unexpected layer structure")
		}
		if layer2 != nil {
			layer = layer2
		}
	}

	if psd.Layer, err = buildTree(layer); err != nil {
		return read, err
	}

	if read != layerAndMaskInfoLen+4 {
		return read, errors.New("psd: layer and mask info read size mismatched. expected " + itoa(layerAndMaskInfoLen+4) + " actual " + itoa(read))
	}
	if Debug != nil {
		Debug.Println("end - layer and mask information section")
	}
	return read, nil
}

func buildTree(layer []Layer) ([]Layer, error) {
	// idx type
	// 0: 3
	// 1:  3
	// 2:   0
	// 3:   0
	// 4:  1
	// 5:  0
	// 6:  0
	// 7: 1
	stack := make([][]Layer, 0, 8)
	n := []Layer{}
	for i, l := range layer {
		switch l.SectionDividerSetting.Type {
		case 1, 2:
			l.Layer = n
			if len(stack) == 0 {
				return nil, errors.New("psd: layer tree structure is broken(#" + itoa(i) + ")")
			}
			n = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n = append(n, l)
		case 3:
			stack = append(stack, n)
			n = []Layer{}
		default:
			n = append(n, l)
		}
	}
	return n, nil
}

func readSectionDividerSetting(l *Layer) (typ int, blendMode BlendMode, subType int, err error) {
	b, ok := l.AdditionalLayerInfo[AdditionalInfoKeySectionDividerSetting]
	if ok {
		typ = int(readUint32(b, 0))
		if len(b) < 12 {
			return typ, BlendModeNormal, 0, nil
		}
		if string(b[4:8]) != "8BIM" {
			return 0, "", 0, errors.New("psd: unexpected signature in section divider setting")
		}
		blendMode = BlendMode(b[8:12])
		if len(b) < 16 {
			return typ, blendMode, 0, nil
		}
		subType = int(readUint32(b, 12))
		return typ, blendMode, subType, nil
	}
	b, ok = l.AdditionalLayerInfo[AdditionalInfoKeySectionDividerSetting2]
	if ok {
		typ = int(readUint32(b, 0))
		if len(b) < 12 {
			return typ, BlendModeNormal, 0, nil
		}
		if string(b[4:8]) != "8BIM" {
			return 0, "", 0, errors.New("psd: unexpected signature in section divider setting 2")
		}
		blendMode = BlendMode(b[8:12])
		if len(b) < 16 {
			return typ, blendMode, 0, nil
		}
		subType = int(readUint32(b, 12))
		return typ, blendMode, subType, nil
	}
	return 0, BlendModeNormal, 0, nil
}

func readLayerInfo(r io.Reader, colorMode ColorMode, depth int, o *DecodeOptions) (layer []Layer, read int, err error) {
	if Debug != nil {
		Debug.Println("start - layer info section")
	}
	b := make([]byte, 16)
	var l int
	if l, err = io.ReadFull(r, b[:4]); err != nil {
		return nil, 0, err
	}
	read += l
	layerInfoLen := int(readUint32(b, 0))
	if Debug != nil {
		Debug.Println("  layerInfoLen:", layerInfoLen)
	}
	if layerInfoLen == 0 {
		if Debug != nil {
			Debug.Println("end - layer info section")
		}
		return nil, read, nil
	}

	if l, err = io.ReadFull(r, b[:2]); err != nil {
		return nil, read, err
	}
	read += l
	numLayers := readUint16(b, 0)
	if numLayers&0x8000 != 0 {
		numLayers = ^numLayers + 1
		// how can I use this info...?
		// hasAlphaForMergedResult = true
	}
	if Debug != nil {
		Debug.Println("  numLayers:", numLayers)
	}

	layerSlice := make([]Layer, numLayers)
	chLen := make([][][2]int, numLayers)
	for i := range layerSlice {
		if Debug != nil {
			Debug.Printf("start - layer #%d structure", i)
		}
		layer := &layerSlice[i]
		layer.SeqID = i

		if l, err = io.ReadFull(r, b[:16]); err != nil {
			return nil, read, err
		}
		read += l
		layer.Rect = image.Rect(
			int(readUint32(b, 4)), int(readUint32(b, 0)),
			int(readUint32(b, 12)), int(readUint32(b, 8)),
		)

		if l, err = io.ReadFull(r, b[:2]); err != nil {
			return nil, read, err
		}
		read += l
		numChannels := int(readUint16(b, 0))
		layer.Channel = map[int]Channel{}
		chLen[i] = make([][2]int, numChannels)
		for j := 0; j < numChannels; j++ {
			if l, err = io.ReadFull(r, b[:6]); err != nil {
				return nil, read, err
			}
			read += l
			id := int(int16(readUint16(b, 0)))
			layer.Channel[id] = Channel{
				Picker: findGrayPicker(depth, false),
			}
			layer.Channel[id].Picker.SetSource(layer.Rect, layer.Channel[id].Data)
			chLen[i][j] = [2]int{id, int(readUint32(b, 2))}
		}

		if l, err = io.ReadFull(r, b[:12]); err != nil {
			return nil, read, err
		}
		read += l
		if string(b[:4]) != "8BIM" {
			return nil, read, errors.New("psd: unexpected the blend mode signature")
		}
		layer.BlendMode = BlendMode(b[4:8])
		layer.Opacity = b[8]
		layer.Clipping = b[9] != 0
		layer.Flags = b[10]
		// b[11] - Filler(zero)
		if l, err = readLayerExtraData(r, layer, colorMode, depth, o); err != nil {
			return nil, read, err
		}
		read += l

		if data, ok := layer.AdditionalLayerInfo[AdditionalInfoKeyUnicodeLayerName]; ok && data != nil {
			layer.UnicodeName = readUnicodeString(data)
			layer.Name = layer.UnicodeName
			delete(layer.AdditionalLayerInfo, AdditionalInfoKeyUnicodeLayerName)
		} else {
			layer.Name = layer.MBCSName
		}

		// this is called "Blend Clipped Layers as Group" in Photoshop's Layer Style dialog.
		if data, ok := layer.AdditionalLayerInfo[AdditionalInfoKeyBlendClippingElements]; ok && data != nil {
			layer.BlendClippedElements = data[0] != 0
			delete(layer.AdditionalLayerInfo, AdditionalInfoKeyBlendClippingElements)
		} else {
			layer.BlendClippedElements = true
		}

		if layer.SectionDividerSetting.Type,
			layer.SectionDividerSetting.BlendMode,
			layer.SectionDividerSetting.SubType,
			err = readSectionDividerSetting(layer); err != nil {
			return nil, read, err
		}
		delete(layer.AdditionalLayerInfo, AdditionalInfoKeySectionDividerSetting)

		// convert mask position from relative to absolute.
		if layer.Mask.Flags&1 != 0 {
			layer.Mask.Rect = layer.Mask.Rect.Sub(layer.Rect.Min)
			layer.Mask.Flags &= ^1
		}
		if layer.Mask.RealFlags&1 != 0 {
			layer.Mask.RealRect = layer.Mask.RealRect.Sub(layer.Rect.Min)
			layer.Mask.RealFlags &= ^1
		}

		if Debug != nil {
			Debug.Printf("layer #%d", layer.SeqID)
			Debug.Println("  Name:", layer.Name)
			Debug.Println("  Rect:", layer.Rect)
			Debug.Println("  Channels:", len(layer.Channel))
			Debug.Println("  BlendMode:", layer.BlendMode)
			Debug.Println("  Opacity:", layer.Opacity)
			Debug.Println("  Clipping:", layer.Clipping)
			Debug.Println("  Flags:", layer.Flags)
			Debug.Println("    TransparecyProtected:", layer.TransparencyProtected())
			Debug.Println("    Visible:", layer.Visible())
			Debug.Println("    Photoshop 5.0 and later:", layer.Flags&8 != 0)
			// this information is useful for the group layer,
			// but sadly it isn't written correctly by some other softwares.
			Debug.Println("    Pixel data irrelevant to appearance of document:", layer.Flags&8 != 0 && layer.Flags&16 != 0)
			Debug.Println("  SectionDividerSetting:")
			Debug.Println("    Type:", layer.SectionDividerSetting.Type)
			Debug.Println("    BlendMode:", layer.SectionDividerSetting.BlendMode)
			Debug.Println("    SubType:", layer.SectionDividerSetting.SubType)
			Debug.Println("  LayerMask:")
			Debug.Println("    Rect:", layer.Mask.Rect)
			Debug.Println("    DefaultColor:", layer.Mask.DefaultColor)
			Debug.Println("    Flags:", layer.Mask.Flags)
			Debug.Println("      Position relative to layer:", layer.Mask.Flags&1 != 0)
			Debug.Println("      Layer mask enabled:", layer.Mask.Enabled())
			Debug.Println("      Invert layer mask when blending (Obsolete):", layer.Mask.Flags&4 != 0)
			Debug.Println("      The user mask actually came from rendering other data:", layer.Mask.Flags&8 != 0)
			Debug.Printf("end - layer #%d structure", i)
		}
	}
	if !o.SkipLayerImage {
		for i, layer := range layerSlice {
			for _, j := range chLen[i] {
				ch := layer.Channel[j[0]]
				var readCh int
				if l, err = io.ReadFull(r, b[:2]); err != nil {
					return nil, read, err
				}
				readCh += l
				read += l

				cmpMethod := CompressionMethod(readUint16(b, 0))
				if Debug != nil {
					Debug.Printf("  layer #%d %q channel #%d image data", i, layer.Name, j[0])
					Debug.Println("    length:", j[1], "compression method:", cmpMethod)
					reportReaderPosition("    file offset: 0x%08x", r)
				}

				var rect image.Rectangle
				switch j[0] {
				case -3:
					rect = layer.Mask.RealRect
				case -2:
					rect = layer.Mask.Rect
				default:
					rect = layer.Rect
				}
				ch.Data = make([]byte, (rect.Dx()*depth+7)>>3*rect.Dy())
				ch.Picker.SetSource(rect, ch.Data)

				if l, err = cmpMethod.Decode(ch.Data, r, int64(j[1]-2), rect, depth, 1); err != nil {
					return nil, read, err
				}
				readCh += l
				read += l

				if readCh != j[1] {
					return nil, read, errors.New("psd: layer: " + itoa(i) + " channel: " + itoa(j[0]) + " read size mismatched. expected " + itoa(j[1]) + " actual " + itoa(readCh))
				}

				layer.Channel[j[0]] = ch
			}

			chs := make([][]byte, colorMode.Channels(), 8)
			for j := range chs {
				chs[j] = layer.Channel[j].Data
			}
			if ch, ok := layer.Channel[-1]; ok {
				chs = append(chs, ch.Data)
			}
			layer.Picker = findPicker(depth, colorMode, colorMode.Channels() < len(chs))
			layer.Picker.SetSource(layer.Rect, chs...)
			layerSlice[i] = layer
			if o.LayerImageLoaded != nil {
				o.LayerImageLoaded(&layerSlice[i], i, len(layerSlice))
			}
		}
		if l, err = adjustAlign4(r, read+4-layerInfoLen); err != nil {
			return nil, read, err
		}
		read += l
	} else {
		type discarder interface {
			Discard(n int) (discarded int, err error)
		}
		skip := layerInfoLen + 4 - read
		switch rr := r.(type) {
		case discarder:
			if l, err = rr.Discard(skip); err != nil {
				return nil, read, err
			}
			read += l
		case io.Seeker:
			if _, err = rr.Seek(int64(skip), 1); err != nil {
				return nil, read, err
			}
			read += skip
		default:
			if l, err = io.ReadFull(r, make([]byte, skip)); err != nil {
				return nil, read, err
			}
			read += l
		}
	}

	if layerInfoLen+4 != read {
		return nil, read, errors.New("psd: layer info read size mismatched. expected " + itoa(layerInfoLen+4) + " actual " + itoa(read))
	}
	if Debug != nil {
		Debug.Println("end - layer info section")
	}
	return layerSlice, read, nil
}

func readLayerExtraData(r io.Reader, layer *Layer, colorMode ColorMode, depth int, o *DecodeOptions) (read int, err error) {
	// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_22582
	// Layer mask / adjustment layer data
	if Debug != nil {
		Debug.Println("start - layer extra data section")
	}
	b := make([]byte, 16)
	var l int
	if l, err = io.ReadFull(r, b[:4]); err != nil {
		return 0, err
	}
	read += l
	extraDataLen := int(readUint32(b, 0))
	if Debug != nil {
		Debug.Println("  extraDataLen:", extraDataLen)
	}
	if extraDataLen == 0 {
		return read, err
	}

	if l, err = io.ReadFull(r, b[:4]); err != nil {
		return read, err
	}
	read += l
	if maskLen := int(readUint32(b, 0)); maskLen > 0 {
		var readMask int
		// Layer mask / adjustment layer data
		// Can be 40 bytes, 24 bytes, or 4 bytes if no layer mask.
		// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_22582
		if l, err = io.ReadFull(r, b); err != nil {
			return read, err
		}
		read += l
		readMask += l
		layer.Mask.Rect = image.Rect(
			int(readUint32(b, 4)), int(readUint32(b, 0)),
			int(readUint32(b, 12)), int(readUint32(b, 8)),
		)

		if l, err = io.ReadFull(r, b[:2]); err != nil {
			return read, err
		}
		read += l
		readMask += l
		layer.Mask.DefaultColor = int(b[0])

		// bit 0 = position relative to layer
		// bit 1 = layer mask disabled
		// bit 2 = invert layer mask when blending (Obsolete)
		// bit 3 = indicates that the user mask actually came from rendering other data
		// bit 4 = indicates that the user and/or vector masks have parameters applied to them
		layer.Mask.Flags = int(b[1])

		layer.Mask.UserMaskDensity = 255
		layer.Mask.VectorMaskDensity = 255

		// Mask Parameters. Only present if bit 4 of Flags set above.
		if layer.Mask.Flags&16 != 0 {
			if l, err = io.ReadFull(r, b[:1]); err != nil {
				return read, err
			}
			read += l
			readMask += l

			// Mask Parameters bit flags present as follows:
			// bit 0 = user mask density, 1 byte
			// bit 1 = user mask feather, 8 byte, double
			// bit 2 = vector mask density, 1 byte
			// bit 3 = vector mask feather, 8 bytes, double
			maskParam := int(b[0])
			if maskParam&1 != 0 {
				if l, err = io.ReadFull(r, b[:1]); err != nil {
					return read, err
				}
				read += l
				readMask += l
				layer.Mask.UserMaskDensity = int(b[0])
			}
			if maskParam&2 != 0 {
				if l, err = io.ReadFull(r, b[:8]); err != nil {
					return read, err
				}
				read += l
				readMask += l
				layer.Mask.UserMaskFeather = readFloat64(b, 0)
			}
			if maskParam&4 != 0 {
				if l, err = io.ReadFull(r, b[:1]); err != nil {
					return read, err
				}
				read += l
				readMask += l
				layer.Mask.VectorMaskDensity = int(b[0])
			}
			if maskParam&8 != 0 {
				if l, err = io.ReadFull(r, b[:8]); err != nil {
					return read, err
				}
				read += l
				readMask += l
				layer.Mask.VectorMaskFeather = readFloat64(b, 0)
			}
		}

		if maskLen == 20 {
			// Padding. Only present if size = 20.
			if l, err = io.ReadFull(r, b[:2]); err != nil {
				return read, err
			}
			read += l
			readMask += l
		} else {
			if l, err = io.ReadFull(r, b[:2]); err != nil {
				return read, err
			}
			read += l
			readMask += l
			layer.Mask.RealFlags = int(b[0])
			layer.Mask.RealBackgroundColor = int(b[1])

			if l, err = io.ReadFull(r, b); err != nil {
				return read, err
			}
			read += l
			readMask += l
			layer.Mask.RealRect = image.Rect(
				int(readUint32(b, 4)), int(readUint32(b, 0)),
				int(readUint32(b, 12)), int(readUint32(b, 8)),
			)
		}
		if maskLen != readMask {
			return read, errors.New("psd: layer mask / adjustment layer data read size mismatched. expected " + itoa(maskLen) + " actual " + itoa(readMask))
		}
	}

	// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_21332
	// Layer blending ranges data
	if l, err = io.ReadFull(r, b[:4]); err != nil {
		return read, err
	}
	read += l
	if blendingRangesLen := int(readUint32(b, 0)); blendingRangesLen > 0 {
		if Debug != nil {
			Debug.Println("  layer blending ranges data skipped:", blendingRangesLen)
		}
		// TODO(oov): implement
		if l, err = io.ReadFull(r, make([]byte, blendingRangesLen)); err != nil {
			return read, err
		}
		read += l
	}

	// Layer name: Pascal string, padded to a multiple of 4 bytes.
	if layer.MBCSName, l, err = readPascalString(r); err != nil {
		return read, err
	}
	read += l
	if l, err = adjustAlign4(r, l); err != nil {
		return read, err
	}
	read += l

	if read < extraDataLen+4 {
		var layers []Layer
		if layer.AdditionalLayerInfo, layers, l, err = readAdditionalLayerInfo(r, extraDataLen+4-read, colorMode, depth, o); err != nil {
			return read, err
		}
		read += l
		if len(layers) > 0 {
			return read, errors.New("psd: unexpected layer structure")
		}
	} else {
		layer.AdditionalLayerInfo = map[AdditionalInfoKey][]byte{}
	}

	if extraDataLen+4 != read {
		return read, errors.New("psd: layer extra info read size mismatched. expected " + itoa(extraDataLen+4) + " actual " + itoa(read))
	}
	if Debug != nil {
		Debug.Println("end - layer extra data section")
	}
	return read, nil
}

func readAdditionalLayerInfo(r io.Reader, infoLen int, colorMode ColorMode, depth int, o *DecodeOptions) (infos map[AdditionalInfoKey][]byte, layers []Layer, read int, err error) {
	if Debug != nil {
		Debug.Println("start - additional layer info section")
		Debug.Println("  infoLen:", infoLen)
	}
	infos = map[AdditionalInfoKey][]byte{}
	b := make([]byte, 8)
	var l int
	for read < infoLen {
		// Sometimes I encountered padded section and unpadded section,
		// so find the signature to avoid failure.
		b[0] = 0
		for read < infoLen && b[0] != '8' {
			if l, err = io.ReadFull(r, b[:1]); err != nil {
				return nil, nil, read, err
			}
			read += l
		}
		if b[0] != '8' {
			break
		}
		if l, err = io.ReadFull(r, b[1:]); err != nil {
			return nil, nil, read, err
		}
		read += l

		if sig := string(b[:4]); sig != "8BIM" && sig != "8B64" {
			return nil, nil, read, errors.New("psd: unexpected the signature found in the additional layer information block")
		}

		key := AdditionalInfoKey(b[4:])
		switch key {
		case AdditionalInfoKeyLayerInfo,
			AdditionalInfoKeyLayerInfo16,
			AdditionalInfoKeyLayerInfo32:
			if Debug != nil {
				Debug.Println("  key:", key)
			}
			var layrs []Layer
			if layrs, l, err = readLayerInfo(r, colorMode, depth, o); err != nil {
				return nil, nil, read, err
			}
			read += l
			layers = append(layers, layrs...)
		default:
			if l, err = io.ReadFull(r, b[:4]); err != nil {
				return nil, nil, read, err
			}
			read += l
			var data []byte
			if ln := int(readUint32(b, 0)); ln > 0 {
				data = make([]byte, ln)
				if l, err = io.ReadFull(r, data); err != nil {
					return nil, nil, read, err
				}
				read += l
			}
			infos[key] = data
			if Debug != nil {
				Debug.Println("  key:", key, "dataLen:", len(data))
			}
		}
	}
	if infoLen != read {
		return nil, nil, read, errors.New("psd: additional layer info read size mismatched. expected " + itoa(infoLen) + " actual " + itoa(read))
	}
	if Debug != nil {
		Debug.Println("end - additional layer info section")
	}
	return infos, layers, read, nil
}
