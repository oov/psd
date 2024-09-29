package psd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"io"
)

func (psd *PSD) AddChannelData(imgs []*image.Image) error {
	// TODO: with compression
	return nil
}

func Encode(psd *PSD, w io.Writer) error {
	if err := psd.Config.encode(w); err != nil {
		return err
	}
	if len(psd.Layer) > 0 {
		return fmt.Errorf("encoding layers not yet supported")
	}

	// image data
	if err := psd.Config.CompressionMethod.encode(psd.Data, &psd.Config, w); err != nil {
		return err
	}
	return nil
}

func (c *Config) encode(w io.Writer) error {
	// header
	h := &headerWrite{
		Signature: headerSignatureBytes,
		Version:   uint16(c.Version),
		Channels:  uint16(c.Channels),
		Height:    uint32(c.Rect.Dy()),
		Width:     uint32(c.Rect.Dx()),
		Depth:     uint16(c.Depth),
		ColorMode: uint16(c.ColorMode),
	}

	if err := binaryWrite(w, h); err != nil {
		return err
	}

	// color mode data
	if err := writeDataBlock(c.ColorModeData, w); err != nil {
		return err
	}

	// image resources
	buf := &bytes.Buffer{}
	for id, imgResource := range c.Res {
		if err := imgResource.encode(id, buf); err != nil {
			return err
		}
	}
	if err := writeDataBlock(buf.Bytes(), w); err != nil {
		return err
	}
	return nil
}

func (i *ImageResource) encode(id int, w io.Writer) error {
	if _, err := w.Write(imageResourceSignatureBytes[:]); err != nil {
		return err
	}
	if err := binaryWrite(w, uint16(id)); err != nil {
		return err
	}
	b, err := stringToPascalBytes(i.Name)
	if err != nil {
		return err
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return writeDataBlockEvenLength(i.Data, w)
}

func (c CompressionMethod) encode(imgDataRaw []byte, cfg *Config, w io.Writer) error {
	if err := binaryWrite(w, uint16(c)); err != nil {
		return err
	}
	switch c {
	case CompressionMethodRaw:
		if _, err := w.Write(imgDataRaw); err != nil {
			return err
		}
	default:
		// TODO: support RLE compression
		return fmt.Errorf("econding only supported for the raw compression method")
	}
	return nil
}

func binaryWrite(w io.Writer, data any) error {
	return binary.Write(w, binary.BigEndian, data)
}
func writeDataBlock(data []byte, w io.Writer) error {
	if err := binaryWrite(w, uint32(len(data))); err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}
func writeDataBlockEvenLength(data []byte, w io.Writer) error {
	if err := writeDataBlock(data, w); err != nil {
		return err
	}
	if len(data)%2 == 1 {
		if _, err := w.Write([]byte{0}); err != nil {
			return err
		}
	}
	return nil
}

type headerWrite struct {
	Signature [4]byte
	Version   uint16
	Reserved  [6]byte
	Channels  uint16
	Height    uint32
	Width     uint32
	Depth     uint16
	ColorMode uint16
}

var headerSignatureBytes [4]byte
var imageResourceSignatureBytes [4]byte

func init() {
	copy(headerSignatureBytes[:], []byte(headerSignature))
	copy(imageResourceSignatureBytes[:], []byte(sectionSignature))
}
