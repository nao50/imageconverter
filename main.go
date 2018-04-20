package main

import (
	"fmt"
	"os"

	"github.com/naoyamaguchi/imageconverter/imgconv"
)

func main() {
	if err := imgconv.Imgconv(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
