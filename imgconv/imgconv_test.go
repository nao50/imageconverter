package imgconv

import (
	"log"
	"testing"
)

func TestNewImages(t *testing.T) {
	filename, imglist, err := NewImages("../images")

	if filename == nil {
		t.Fatal("fail error func test")
	}
	if imglist == nil {
		t.Fatal("fail error func test")
	}
	if err != nil {
		t.Fatal("fail error func test")
	}
}

func TestImgconv(t *testing.T) {
	filepath, image, err := NewImages("../images")
	if err != nil {
		log.Fatal(err)
	}

	err = Imgconv("jpg", filepath, image)
	if err != nil {
		t.Fatal("fail error func test")
	}
}
