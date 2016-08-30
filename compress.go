package psd

import (
	"compress/zlib"
	"errors"
	"image"
	"io"
	"runtime"
	"sync"
)

// CompressionMethod represents compression method that is used in psd file.
type CompressionMethod int16

// These compression methods are defined in this document.
//
// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_pgfId-1054855
const (
	CompressionMethodRaw                  = CompressionMethod(0)
	CompressionMethodRLE                  = CompressionMethod(1)
	CompressionMethodZIPWithoutPrediction = CompressionMethod(2)
	CompressionMethodZIPWithPrediction    = CompressionMethod(3)
)

// Decode decodes the compressed image data from r.
//
// You can pass 0 to sizeHint if unknown, but in this case may read more data than necessary from r.
func (cm CompressionMethod) Decode(dest []byte, r io.Reader, sizeHint int64, rect image.Rectangle, depth int, channels int, large bool) (read int, err error) {
	switch cm {
	case CompressionMethodRaw:
		return io.ReadFull(r, dest)
	case CompressionMethodRLE:
		return decodePackBits(dest, r, rect.Dx(), rect.Dy()*channels, large)
	case CompressionMethodZIPWithoutPrediction:
		return decodeZLIB(dest, r, sizeHint)
	case CompressionMethodZIPWithPrediction:
		if read, err = decodeZLIB(dest, r, sizeHint); err != nil {
			return read, err
		}
		decodeDelta(dest, rect.Dx(), depth)
		return read, err
	}
	return 0, errors.New("psd: compression method " + itoa(int(cm)) + " is not implemented")
}

func decodePackBits(dest []byte, r io.Reader, width int, lines int, large bool) (read int, err error) {
	buf := make([]byte, lines*(get4or8(large)>>1))
	var l int
	if l, err = io.ReadFull(r, buf); err != nil {
		return
	}
	read += l

	total := 0
	lens := make([]int, lines)
	offsets := make([]int, lines)
	ofs := 0
	if large {
		for i := range lens {
			l = int(readUint32(buf, ofs))
			lens[i] = l
			offsets[i] = total
			total += l
			ofs += 4
		}
	} else {
		for i := range lens {
			l = int(readUint16(buf, ofs))
			lens[i] = l
			offsets[i] = total
			total += l
			ofs += 2
		}
	}

	buf = make([]byte, total)
	if l, err = io.ReadFull(r, buf); err != nil {
		return
	}
	read += l

	n := runtime.GOMAXPROCS(0)
	for n > 1 && n<<1 > lines {
		n--
	}
	if n == 1 {
		decodePackBitsPerLine(dest, buf, lens)
		return
	}

	var wg sync.WaitGroup
	wg.Add(n)
	step := lines / n
	ofs = 0
	for i := 1; i < n; i++ {
		go func(dest []byte, buf []byte, lens []int) {
			defer wg.Done()
			decodePackBitsPerLine(dest, buf, lens)
		}(dest[ofs*width:(ofs+step)*width], buf[offsets[ofs]:offsets[ofs+step]], lens[ofs:ofs+step])
		ofs += step
	}
	go func() {
		defer wg.Done()
		decodePackBitsPerLine(dest[ofs*width:], buf[offsets[ofs]:], lens[ofs:])
	}()
	wg.Wait()
	return
}

func decodePackBitsPerLine(dest []byte, buf []byte, lens []int) {
	var l int
	for _, ln := range lens {
		for i := 0; i < ln; {
			if buf[i] <= 0x7f {
				l = int(buf[i]) + 1
				copy(dest[:l], buf[i+1:])
				dest = dest[l:]
				i += l + 1
				continue
			}
			l = int(-buf[i]) + 1
			for j, c := 0, buf[i+1]; j < l; j++ {
				dest[j] = c
			}
			dest = dest[l:]
			i += 2
		}
		buf = buf[ln:]
	}
}

// decodeZLIB decodes compressed data by zlib.
// You can pass 0 to length if unknown,
// but in that case may read more data than necessary from r.
func decodeZLIB(dest []byte, r io.Reader, length int64) (read int, err error) {
	N := length
	if N == 0 {
		N = 0x7fffffffffffffff
	}
	lr := &io.LimitedReader{R: r, N: N}
	zr, err := zlib.NewReader(lr)
	if err != nil {
		return 0, err
	}
	if _, err = io.ReadFull(zr, dest); err != nil {
		zr.Close()
		return int(N - lr.N), err
	}
	if err = zr.Close(); err != nil {
		return int(N - lr.N), err
	}
	return int(N - lr.N), nil
}

func decodeDelta(buf []byte, width int, depth int) {
	switch depth {
	case 16:
		var d uint16
		for i := 0; i < len(buf); {
			d = 0
			for j := 0; j < width; j++ {
				d += readUint16(buf, i)
				writeUint16(buf, d, i)
				i += 2
			}
		}
	case 32:
		var d uint32
		for i := 0; i < len(buf); {
			d = 0
			for j := 0; j < width; j++ {
				d += readUint32(buf, i)
				writeUint32(buf, d, i)
				i += 4
			}
		}
	}
}
