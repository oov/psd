package psd

import (
	"io"
	"unicode/utf16"
)

func itoa(x int) string {
	var b [32]byte
	var minus bool
	if x < 0 {
		minus = true
		x = -x
	}
	i := len(b) - 1
	for x > 9 {
		b[i] = byte(x%10 + '0')
		x /= 10
		i--
	}
	b[i] = byte(x + '0')
	if minus {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

func readUint16(b []byte, offset int) uint16 {
	return uint16(b[offset])<<8 | uint16(b[offset+1])
}

func writeUint16(b []byte, v uint16, offset int) {
	b[offset] = uint8(v >> 8)
	b[offset+1] = uint8(v)
}

func readUint32(b []byte, offset int) uint32 {
	return uint32(b[offset])<<24 | uint32(b[offset+1])<<16 | uint32(b[offset+2])<<8 | uint32(b[offset+3])
}

func writeUint32(b []byte, v uint32, offset int) {
	b[offset] = uint8(v >> 24)
	b[offset+1] = uint8(v >> 16)
	b[offset+2] = uint8(v >> 8)
	b[offset+3] = uint8(v)
}

func readUnicodeString(b []byte) string {
	ln := readUint32(b, 0)
	if ln == 0 {
		return ""
	}
	buf := make([]uint16, ln)
	for i := range buf {
		buf[i] = readUint16(b, 4+i<<1)
	}
	return string(utf16.Decode(buf))
}

func adjustAlign2(r io.Reader, l int) (read int, err error) {
	if l&1 != 0 {
		var b [1]byte
		return r.Read(b[:])
	}
	return 0, nil
}

func adjustAlign4(r io.Reader, l int) (read int, err error) {
	if gap := l & 3; gap > 0 {
		var b [4]byte
		return r.Read(b[:4-gap])
	}
	return 0, nil
}

func readPascalString(r io.Reader) (str string, read int, err error) {
	b := make([]byte, 1)
	if _, err := io.ReadFull(r, b); err != nil {
		return "", 0, err
	}
	if b[0] == 0 {
		return "", 1, nil
	}
	buf := make([]byte, b[0])
	if _, err := io.ReadFull(r, buf); err != nil {
		return "", 1, err
	}
	return string(buf), len(buf) + 1, nil
}
