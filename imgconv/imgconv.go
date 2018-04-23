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

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func ImgConv(i Imageconverter, srcdir string) {
	fmt.Println(i.GetImage(srcdir))
}

type Imageconverter interface {
	GetImage(srcdir string) ([]Imagefile, error)
	ConvertImage(outputfiletype string, imagefile Imagefile) error
}

type Imagefile struct {
	image         image.Image
	imagefilepath string
	// imagefilename string
	// imagefiletype string
}

func (i *Imagefile) GetImage(srcdir string) ([]Imagefile, error) {
	imagefile := Imagefile{}
	imagefilelist := []Imagefile{}

	err := filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// // Get imagefilename
		// filename := strings.TrimSuffix(path, filepath.Ext(path))
		// fmt.Println("filename: ", filename)
		// // Get imagefiletype
		// pos := strings.LastIndex(filename, "/")
		// filetype := filename[pos:]
		// fmt.Println("filetype: ", filetype)
		// GEt image
		img, err := getImg(path)
		if err != nil {
			return err
		}
		imagefile = Imagefile{
			image:         img,
			imagefilepath: path,
			// imagefilename: filename,
			// imagefiletype: filetype,
		}
		// fmt.Printf("imagefile: %v \n", imagefile)
		imagefilelist = append(imagefilelist, imagefile)
		return nil
	})
	// fmt.Printf("imagefile summry: %v \n", imagefilelist)
	return imagefilelist, err
}

func (i *Imagefile) ConvertImage(outputfiletype string, imagefile Imagefile) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
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
