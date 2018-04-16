package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"reflect"
	"strings"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func Init() (image.Image, string, int, int, string) {
	inputfilename := flag.String("file", "test.png", "Use -file <filesource>")
	outputformat := flag.String("out", "jpg", "Use -out <output file format>")
	flag.Parse()

	f, err := os.Open(*inputfilename)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	f, err = os.Open(*inputfilename)
	if err != nil {
		log.Fatal(err)
	}

	conf, _, err := image.DecodeConfig(f)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	return img, *inputfilename, conf.Width, conf.Height, *outputformat
}

// func ConvertFomat(img image.Image, outputformat, inputformat, inputfilename string, w, h int) error {
func ConvertFomat(img image.Image, inputfilename, outputformat string, w, h int) error {

	file := strings.LastIndex(inputfilename, "/")
	pos := strings.LastIndex(inputfilename[file+1:], ".")

	a := inputfilename[file+1:]

	fmt.Println(inputfilename[file+1:])
	fmt.Println(a[:pos])

	// pos := strings.LastIndex(inputfilename, ".")

	if _, err := os.Stat("out"); err != nil {
		if err := os.Mkdir("out", 0755); err != nil {
			fmt.Println(err)
		}
	}

	df, err := os.Create("out/" + a[:pos] + "." + outputformat)
	if err != nil {
		return fmt.Errorf("can't write images")
	}
	defer df.Close()

	switch strings.ToLower(outputformat) {
	case "png":
		err = png.Encode(df, img)
		fmt.Println("png called")
	case "jpeg", "jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{jpeg.DefaultQuality})
		fmt.Println("jpg called")
	case "ascii":
		table := []byte(ASCIISTR)
		buf := new(bytes.Buffer)

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				g := color.GrayModel.Convert(img.At(j, i))
				y := reflect.ValueOf(g).FieldByName("Y").Uint()
				pos := int(y * 16 / 255)
				_ = buf.WriteByte(table[pos])
			}
			_ = buf.WriteByte('\n')
		}
		fmt.Print(string(buf.Bytes()))
	}

	return nil
}

func main() {
	img, inputfilename, width, height, outputformat := Init()
	fmt.Printf("inputfilename: %v\n", inputfilename)
	fmt.Printf("outputformat: %v\n", outputformat)

	// p := Convert2Ascii(img, width, height)
	// fmt.Print(string(p))

	ConvertFomat(img, inputfilename, outputformat, width, height)
}

// imgconv -out png ./
// outディレクトリに変換後のimg
