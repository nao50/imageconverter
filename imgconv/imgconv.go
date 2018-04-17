// Package imgconv provides some simple image convert function.
// These are sample functions to practice golang.
package imgconv

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var ASCIISTR = "MND80Z$7I?+=~:,.."

func Imgconv() error {
	if len(os.Args) < 2 {
		fmt.Errorf("usage")
	}

	outputformat := flag.String("outputformat", "png", "Use -outputformat <outputformat>")
	flag.Parse()

	imagedirectory := flag.Args()

	if _, err := os.Stat("out"); err != nil {
		if err := os.Mkdir("out", 0755); err != nil {
			fmt.Println(err)
		}
	}

	err := filepath.Walk(imagedirectory[0], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			f, err := os.Open(imagedirectory[0] + "/" + info.Name())
			if err != nil {
				panic(err)
			}

			img, _, err := image.Decode(f)
			if err != nil {
				fmt.Println(err)
			}

			f, err = os.Open(imagedirectory[0] + "/" + info.Name())
			if err != nil {
				panic(err)
			}

			conf, _, err := image.DecodeConfig(f)
			if err != nil {
				fmt.Println(err)
			}

			pos := strings.LastIndex(info.Name(), ".")
			dest, err := os.Create("out/" + info.Name()[:pos] + "." + *outputformat)
			if err != nil {
				fmt.Println("error")
			}

			switch strings.ToLower(*outputformat) {
			case "png":
				err = png.Encode(dest, img)
			case "jpg", "jpeg":
				err = jpeg.Encode(dest, img, &jpeg.Options{jpeg.DefaultQuality})
			case "ascii":
				table := []byte(ASCIISTR)
				buf := new(bytes.Buffer)

				for i := 0; i < conf.Height; i++ {
					for j := 0; j < conf.Width; j++ {
						g := color.GrayModel.Convert(img.At(j, i))
						y := reflect.ValueOf(g).FieldByName("Y").Uint()
						pos := int(y * 16 / 255)
						err = buf.WriteByte(table[pos])
						if err != nil {
							panic(err)
						}
					}
					_ = buf.WriteByte('\n')
				}
				fmt.Print(string(buf.Bytes()))
				_, err = dest.WriteString(string(buf.Bytes()))
				if err != nil {
					panic(err)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error on filepath.Walk : ", err)
	}

	return nil
}
