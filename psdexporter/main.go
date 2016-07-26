package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/oov/psd"
)

func processLayer(filename string, layerName string, l *psd.Layer) error {
	if len(l.Layer) > 0 {
		for i, ll := range l.Layer {
			if err := processLayer(
				fmt.Sprintf("%s_%03d", filename, i),
				layerName+"/"+ll.Name, &ll); err != nil {
				return err
			}
		}
	}
	if !l.HasImage() {
		return nil
	}
	fmt.Printf("%s -> %s.png\n", layerName, filename)
	out, err := os.Create(fmt.Sprintf("%s.png", filename))
	if err != nil {
		return err
	}
	defer out.Close()
	return png.Encode(out, l.Picker)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s psdfile", os.Args[0])
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := psd.Decode(file, &psd.DecodeOptions{SkipMergedImage: true})
	if err != nil {
		panic(err)
	}
	for i, layer := range img.Layer {
		if err = processLayer(fmt.Sprintf("%03d", i), layer.Name, &layer); err != nil {
			panic(err)
		}
	}
}
