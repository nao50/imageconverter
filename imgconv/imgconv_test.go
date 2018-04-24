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
	"testing"
)

func TestImgConv(t *testing.T) {
	i := &MockImagefile{}

	ImgConv(i, "test", "jpg")

}

type MockImagefile struct {
	image         image.Image
	imagefilepath string
}

func (i *MockImagefile) GetImage(srcdir string) ([]Imagefile, error) {
	imagefilelist := []MockImagefile{}

	createDummyJPG()
	createDummyPNG()

	err := filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		img, err := getImg(path)
		if err != nil {
			return err
		}

		i.image = img
		i.imagefilepath = path

		imagefilelist = append(imagefilelist, *i)
		return nil
	})
	return imagefilelist, err
}

func (i *MockImagefile) ConvertImage(outputfiletype string, imagefile []Imagefile) error {
	if _, err := os.Stat("out"); err != nil {
		if err := os.Mkdir("out", 0755); err != nil {
			fmt.Println(err)
		}
	}

	for _, imagefile := range imagefile {
		f_pos := strings.LastIndex(imagefile.imagefilepath, "/")
		p_pos := strings.LastIndex(imagefile.imagefilepath, ".")
		out, err := os.Create("out/" + imagefile.imagefilepath[f_pos+1:p_pos] + "." + outputfiletype)
		if err != nil {
			return err
		}
		switch outputfiletype {
		case "jpeg", "jpg":
			err = jpeg.Encode(out, imagefile.image, nil)
			if err != nil {
				return err
			}
		case "png":
			err = png.Encode(out, imagefile.image)
			if err != nil {
				return err
			}
		default:
			return errors.New("sorry. not support this outputfiletype extend")
		}
	}
	return nil

}

// func TestNewImages(t *testing.T) {
// 	filename, imglist, err := NewImages("../images")

// 	if filename == nil {
// 		t.Fatal("fail error func test")
// 	}
// 	if imglist == nil {
// 		t.Fatal("fail error func test")
// 	}
// 	if err != nil {
// 		t.Fatal("fail error func test")
// 	}
// }

// func TestImgconv(t *testing.T) {
// 	t.Helper()
// 	dir, _ := os.Getwd()
// 	println(dir)

// 	filepath, image, err := NewImages("../images")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	outType := []struct {
// 		outtype string
// 	}{
// 		{outtype: "png"},
// 		{outtype: "jpg"},
// 	}

// 	for _, c := range outType {

// 		err = Imgconv(c.outtype, filepath, image)
// 		if err != nil {
// 			t.Fatal("fail error func test")
// 		}
// 		// _, err := os.Stat(filepath[0] + c.outtype)
// 		// if err != nil {
// 		// 	t.Fatal("fail error func test")
// 		// }
// 	}
// }

/*
package main

import (
	"fmt"
)

func main() {
	s := &StudentImpl{}
	// s := &StudentMock{}
	name, age := Show(s)
	fmt.Println("name: ", name)
	fmt.Println("age: ", age)
}

func Show(s Student) (string, int) {
	name := s.Name()
	age := s.Age()
	return name, age
}

type Student interface {
	Name() string
	Age() int
}

// db access
type StudentImpl struct{}

func (s *StudentImpl) Name() string {
	name := "Taro"
	//fmt.Println(name)
	return name
}

func (s *StudentImpl) Age() int {
	age := 15
	//fmt.Println(age)
	return age
}
*/

/*
package main

import (
	"testing"
)

// mock
type StudentMock struct{}

func (s *StudentMock) Name() string {
	name := "Mock Taro"
	return name
}

func (s *StudentMock) Age() int {
	age := 100
	return age
}

func Testshow(t *testing.T) {
	s := &StudentMock{}
	name, age := Show(s)
	if name != "Mock Taro" {
		t.Fatalf("failed test %#v", name)
	}
	if age != 100 {
		t.Fatal("failed test")
	}
}
*/

func createDummyJPG() (image.Image, error) {
	if _, err := os.Stat("test"); err != nil {
		if err := os.Mkdir("test", 0755); err != nil {
			fmt.Println(err)
		}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, 64, 64))
	file, _ := os.Create("test/" + "test.jpg")
	defer file.Close()
	if err := jpeg.Encode(file, rgba, &jpeg.Options{100}); err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func createDummyPNG() (image.Image, error) {
	if _, err := os.Stat("test"); err != nil {
		if err := os.Mkdir("test", 0755); err != nil {
			fmt.Println(err)
		}
	}
	rgba := image.NewRGBA(image.Rect(0, 0, 64, 64))
	file, _ := os.Create("test/" + "test.png")
	defer file.Close()
	if err := png.Encode(file, rgba); err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
