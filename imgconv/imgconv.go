// Package imgconv provides some simple image convert function.
// These are sample functions to practice golang.
package imgconv

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func NewImages(srcDir string) ([]string, []image.Image, error) {
	var filename []string
	var img image.Image
	var imglist []image.Image

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		filename = append(filename, strings.TrimSuffix(path, filepath.Ext(path)))
		img, err = getImg(path)
		if err != nil {
			return err
		}
		imglist = append(imglist, img)

		return nil
	})

	return filename, imglist, err
}

func getImg(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func Imgconv(outType string, filename []string, img []image.Image) error {
	if _, err := os.Stat("out"); err != nil {
		if err := os.Mkdir("out", 0755); err != nil {
			fmt.Println(err)
		}
	}

	for _, filename := range filename {
		pos := strings.LastIndex(filename, "/")
		out, err := os.Create("out/" + filename[pos+1:] + "." + outType)
		if err != nil {
			return err
		}
		switch outType {
		case "jpeg", "jpg":
			for _, img := range img {
				err = jpeg.Encode(out, img, nil)
				if err != nil {
					return err
				}
			}
		case "png":
			for _, img := range img {
				err = png.Encode(out, img)
				if err != nil {
					return err
				}
			}
		default:
			return errors.New("sorry. not support this outType extend")
		}

	}
	return nil
}
