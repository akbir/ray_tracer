package misc

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}

}
func OpenFile() *os.File {
	// open the file
	dir, err := os.Getwd()
	f, err := os.Create(dir + "/image.png")
	check(err, "Unable to open file")
	return f
}

func WriteFile(f *os.File, img *image.NRGBA) {
	err := png.Encode(f, img)
	check(err, "Error writing to file: %v\n")
}
