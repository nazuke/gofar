package main

import (
	"flag"
	"fmt"

	"github.com/nazuke/gofar/src/gofar"
)

func main() {

	var width *int = flag.Int("width", 640, "The width of the area")
	var height *int = flag.Int("height", 480, "The height of the rectangle")
	var noNewline *bool = flag.Bool("nonewline", false, "Omit the newline in the output")

	flag.Parse()

	var ar string = gofar.FindAspectRatio(*width, *height)

	if *noNewline {
		fmt.Printf("%s", ar)
	} else {
		fmt.Printf("%s\n", ar)
	}

}
