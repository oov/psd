// +build ignore

package main

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"
)

var source = `// DO NOT EDIT.
// Generate with: go generate

package blend

var divTable = [256]uint32{
	{{range $i, $v := .N}}{{printf "0x%08x, // %d\n" $v $i}}{{end}}}
`

func main() {
	t := template.Must(template.New("").Parse(source))

	n := make([]int64, 256)
	for i := int64(1); i < 256; i++ {
		n[i] = 0xffffffff / i
	}

	b := bytes.NewBufferString("")
	if err := t.Execute(b, map[string]interface{}{
		"N": n,
	}); err != nil {
		log.Fatal(err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("zdivtable.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write(buf); err != nil {
		log.Fatal(err)
	}
}
