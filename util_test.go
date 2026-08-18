package psd

import (
	"bytes"
	"testing"
)

func TestReadUnicodeStringTruncated(t *testing.T) {
	// A length prefix that claims more code units than the buffer holds must
	// not read (or allocate) beyond the buffer.
	if got := readUnicodeString([]byte{0x00, 0x00, 0x00, 0x01}); got != "" {
		t.Fatalf("truncated string: got %q, want empty", got)
	}
	if got := readUnicodeString([]byte{0x00, 0x00}); got != "" {
		t.Fatalf("short buffer: got %q, want empty", got)
	}
	// A well-formed value still decodes.
	want := "Hi"
	b := []byte{0x00, 0x00, 0x00, 0x02, 0x00, 'H', 0x00, 'i'}
	if got := readUnicodeString(b); got != want {
		t.Fatalf("valid string: got %q, want %q", got, want)
	}
}

// craftPSDWithLuni builds a minimal single-layer PSD whose layer carries a
// 'luni' (Unicode layer name) block declaring one code unit but supplying no
// character bytes.
func craftPSDWithLuni() []byte {
	var b bytes.Buffer
	b.WriteString("8BPS")                   // signature
	b.Write([]byte{0x00, 0x01})             // version 1
	b.Write(make([]byte, 6))                // reserved
	b.Write([]byte{0x00, 0x01})             // channels
	b.Write([]byte{0x00, 0x00, 0x00, 0x01}) // height
	b.Write([]byte{0x00, 0x00, 0x00, 0x01}) // width
	b.Write([]byte{0x00, 0x08})             // depth
	b.Write([]byte{0x00, 0x03})             // color mode: RGB
	b.Write([]byte{0x00, 0x00, 0x00, 0x00}) // color mode data length
	b.Write([]byte{0x00, 0x00, 0x00, 0x00}) // image resource length
	b.Write([]byte{0x00, 0x00, 0x00, 0xFF}) // layer and mask info length
	b.Write([]byte{0x00, 0x00, 0x00, 0xFF}) // layer info length
	b.Write([]byte{0x00, 0x01})             // layer count
	b.Write(make([]byte, 16))               // layer rectangle
	b.Write([]byte{0x00, 0x00})             // channel count
	b.WriteString("8BIM")                   // blend mode signature
	b.WriteString("norm")                   // blend mode key
	b.Write([]byte{0x00, 0x00, 0x00, 0x00}) // opacity, clipping, flags, filler
	b.Write([]byte{0x00, 0x00, 0x00, 0x1C}) // extra data length
	b.Write([]byte{0x00, 0x00, 0x00, 0x00}) // layer mask length
	b.Write([]byte{0x00, 0x00, 0x00, 0x00}) // blending ranges length
	b.Write([]byte{0x00})                   // pascal name length
	b.Write([]byte{0x00, 0x00, 0x00})       // 4-byte alignment padding
	b.WriteString("8BIM")                   // additional info signature
	b.WriteString("luni")                   // key
	b.Write([]byte{0x00, 0x00, 0x00, 0x04}) // block length
	b.Write([]byte{0x00, 0x00, 0x00, 0x01}) // Unicode length 1, no character data
	return b.Bytes()
}

func TestDecodeTruncatedUnicodeLayerName(t *testing.T) {
	// Decoding the layer must not panic while parsing the truncated
	// Unicode layer name.
	opt := &DecodeOptions{SkipLayerImage: true, SkipMergedImage: true}
	if _, _, err := Decode(bytes.NewReader(craftPSDWithLuni()), opt); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
