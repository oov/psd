package psd

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// ImageResource represents the image resource that is used in psd file.
type ImageResource struct {
	ID   int
	Name string
	Data []byte
}

func readImageResource(r io.Reader) (resMap map[int]ImageResource, read int, err error) {
	if Debug != nil {
		Debug.Println("start - image resources section")
	}
	// Image Resources Section
	// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_69883
	var l int
	b := make([]byte, 6)
	if l, err = io.ReadFull(r, b[:4]); err != nil {
		return nil, 0, err
	}
	read += l
	imageResourceLen := int(readUint32(b, 0))
	if imageResourceLen == 0 {
		return map[int]ImageResource{}, read, nil
	}

	resMap = map[int]ImageResource{}

	// Image Resource Blocks
	// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_46269
	var id int
	for read < imageResourceLen {
		var res ImageResource

		// Signature: '8BIM'
		if l, err = io.ReadFull(r, b[:6]); err != nil {
			return nil, read, err
		}
		read += l

		// http://fileformats.archiveteam.org/wiki/Photoshop_Image_Resources
		// This document says:
		// Each block usually begins with the ASCII characters "8BIM",
		// a signature that appears in several Photoshop formats.
		// Other block signatures that have been observed are
		// "MeSa" (apparently associated with ImageReady),
		// "PHUT" (apparently associated with PhotoDeluxe), "AgHg", and "DCSR".
		sig := string(b[:4])
		if sig != "8BIM" && sig != "MeSa" && sig != "PHUT" && sig != "AgHg" && sig != "DCSR" {
			return nil, read, errors.New("psd: invalid image resource signature")
		}

		// Unique identifier for the resource.
		// Image resource IDs contains a list of resource IDs used by Photoshop.
		// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_38034
		id = int(readUint16(b, 4))
		res.ID = id

		// Name: Pascal string, padded to make the size even
		// (a null name consists of two bytes of 0)
		if res.Name, l, err = readPascalString(r); err != nil {
			return nil, read, err
		}
		read += l
		if l, err = adjustAlign2(r, l); err != nil {
			return nil, read, err
		}
		read += l

		// The resource data, described in the sections on the individual resource types.
		// It is padded to make the size even.
		if l, err = io.ReadFull(r, b[:4]); err != nil {
			return nil, read, err
		}
		read += l
		res.Data = make([]byte, readUint32(b, 0))
		if l, err = io.ReadFull(r, res.Data); err != nil {
			return nil, read, err
		}
		read += l
		if l, err = adjustAlign2(r, l); err != nil {
			return nil, read, err
		}
		read += l
		if Debug != nil {
			Debug.Printf("  id: %04x(%d) name: %s dataLen: %d", id, id, res.Name, len(res.Data))
		}
		resMap[id] = res
	}
	if Debug != nil {
		Debug.Println("end - image resources section")
	}
	return resMap, read, nil
}

func hasAlphaID0(Res map[int]ImageResource) bool {
	// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_38034
	// 0x041D(1053) | (Photoshop 6.0) Alpha Identifiers.
	//                4 bytes of length, followed by 4 bytes each for every alpha identifier.
	//
	// This description is probably a little different from the actual structure.
	// "4 bytes of length" is included in Image Resource Blocks structure, not in the resource data block.
	// so alpha identifier appears from the head in the resource data block.
	aid, ok := Res[0x041D]
	if !ok {
		return true
	}
	for i, ln := 0, len(aid.Data); i < ln; i += 4 {
		if readUint32(aid.Data, i) == 0 {
			return true
		}
	}
	return false
}

type AlphaNames struct {
	Names []string // pascal format strings
}

const idAlphaNames = 1006

func (c *Config) ParseAlphaNames() (*AlphaNames, error) {
	res, ok := c.Res[idAlphaNames]
	if !ok {
		return nil, fmt.Errorf("no alpha names defined")
	}
	r := bytes.NewReader(res.Data)
	names := make([]string, 0)
	read := 0
	for read < len(res.Data) {
		s, n, err := readPascalString(r)
		if err != nil {
			return nil, err
		}
		names = append(names, s)
		read += n
	}
	return &AlphaNames{names}, nil
}
func (an *AlphaNames) Encode() (*ImageResource, error) {
	data := []byte{}
	for _, n := range an.Names {
		b, err := stringToPascalBytes(n)
		if err != nil {
			return nil, err
		}
		data = append(data, b...)
	}
	return &ImageResource{
		ID:   idAlphaNames,
		Data: data,
	}, nil
}

type DisplayInfo struct {
	Channels []DisplayInfoChannel
}
type DisplayInfoChannel struct {
	ColorSpace uint16 // TODO: define enum for this
	Color      [4]uint16
	Opacity    uint16
	Mode       DisplayInfoChannelMode
}

type DisplayInfoChannelMode uint8

const (
	DisplayChannelModeAlpha DisplayInfoChannelMode = iota
	DisplayChannelModeAlphaInverted
	DisplayChannelModeSpot
)

const idDisplayInfo = 1077

func (c *Config) ParseDisplayInfo() (*DisplayInfo, error) {
	res, ok := c.Res[idDisplayInfo]
	if !ok {
		return nil, fmt.Errorf("no alpha names defined")
	}
	read := 4 // start reading after version
	channels := make([]DisplayInfoChannel, 0)
	for read < len(res.Data) {
		r := bytes.NewReader(res.Data[read : read+13])
		c := &DisplayInfoChannel{}
		if err := binary.Read(r, binary.BigEndian, c); err != nil {
			return nil, err
		}
		read += 13
		channels = append(channels, *c)
	}
	return &DisplayInfo{channels}, nil
}
func (di *DisplayInfo) Encode() (*ImageResource, error) {
	data := bytes.NewBuffer([]byte{})
	// version
	if err := binaryWrite(data, uint32(1)); err != nil {
		return nil, err
	}
	// channels
	for _, c := range di.Channels {
		if err := binaryWrite(data, c); err != nil {
			return nil, err
		}
	}
	return &ImageResource{
		ID:   idDisplayInfo,
		Data: data.Bytes(),
	}, nil
}
