package main

import "flag"

var (
	encodeMode bool
	multiMode  bool
	paintMode  bool
	helpMode   bool
)

// flags registeration. No need to run bc init() will be called automatically before main() is executed.
func init() {
	flag.BoolVar(&encodeMode, "encode", false, "Encode plain text into art-decoder notation")
	flag.BoolVar(&encodeMode, "e", false, "")

	flag.BoolVar(&multiMode, "multi", false, "Read multiple lines from stdin")
	flag.BoolVar(&multiMode, "m", false, "")

	flag.BoolVar(&paintMode, "paint", false, "Colorize output: each unique character gets its own ANSI color")
	flag.BoolVar(&paintMode, "p", false, "")

	flag.BoolVar(&helpMode, "help", false, "Show help")
	flag.BoolVar(&helpMode, "h", false, "")
}
