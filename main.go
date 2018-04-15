package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"
)

var ASCIISTR = "MND8OZ$7I?+=~:,.."

func Init() (image.Image, int, int, string, string) {
	width := flag.Int("w", 30, "Use -w <width>")
	height := flag.Int("h", 30, "Use -w <height>")
	fpath := flag.String("input", "test.png", "Use -input <filesource>")
	outputformat := flag.String("out", "jpg", "Use -out <output file format>")
	flag.Parse()

	f, err := os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
	}

	// buf := new(bytes.Buffer)
	// io.Copy(buf, f)
	// fmt.Println(buf)

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	f, err = os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
	}

	conf, encode, err := image.DecodeConfig(f)
	if err != nil {
		panic(err)
	}

	// サイズ表示
	fmt.Printf("INPUT Image Size width: %v px, height: %v px\n", conf.Width, conf.Height)
	// フォーマット名表示
	fmt.Printf("INPUT Image Format: %v\n", encode)

	defer f.Close()
	return img, *width, *height, *outputformat, encode
}

func Convert2Ascii(img image.Image, w, h int) []byte {
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
	return buf.Bytes()
}

func ConvertFomat(img image.Image, w, h int) []byte {
	buf := new(bytes.Buffer)

	return buf.Bytes()
}

func main() {
	img, width, height, outputformat, encode := Init()
	fmt.Printf("inputformat: %v\n", encode)
	fmt.Printf("outputformat: %v\n", outputformat)

	p := Convert2Ascii(img, width, height)
	fmt.Print(string(p))
}
