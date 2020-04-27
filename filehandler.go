package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func check(e error, s string){
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}

}
func openFile() *os.File {
	// open the file
	f, err := os.Create("/tmp/pic.png")
	check(err, "Unable to open file")

	// start writing the ppm file
	//_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)
	return f
}

func writeFile(f *os.File, img *image.NRGBA) {
	err := png.Encode(f, img)
	check(err, "Error writing to file: %v\n")
}