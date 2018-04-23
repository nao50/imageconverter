package main

import (
	"flag"

	"github.com/naoyamaguchi/imageconverter/imgconv"
)

var outType string
var srcDir string

func init() {
	flag.StringVar(&outType, "out", "png", "set output image type")
	flag.StringVar(&outType, "o", "png", "shorthand 'out flag")
}

func main() {
	flag.Parse()
	srcDir = flag.Arg(0)

	// filepath, image, err := imgconv.NewImages(srcDir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := imgconv.Imgconv(outType, filepath, image); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	// 	os.Exit(1)
	// }

	i := &imgconv.Imagefile{}
	// fmt.Println(i)
	imgconv.ImgConv(i, srcDir)
}
