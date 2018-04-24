package imgconv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestImgConv(t *testing.T) {
	t.Helper()

	prevDir, _ := filepath.Abs(".")

	i := &Imagefile{}

	outType := []struct {
		outtype string
	}{
		{outtype: "png"},
		{outtype: "jpg"},
	}

	for _, c := range outType {
		ImgConv(i, "../images", c.outtype)
	}

	_, err := os.Stat(prevDir + "/out/")
	if err != nil {
		t.Fatal("err")
	}
}
