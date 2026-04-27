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
	flag.BoolVar(&encodeMode, "encode", false, "enable encode mode")
	flag.BoolVar(&encodeMode, "e", false, "alias for encode")

	flag.BoolVar(&multiMode, "multi", false, "enable multi mode")
	flag.BoolVar(&multiMode, "m", false, "alias for multi")

	flag.BoolVar(&paintMode, "paint", false, "enable paint mode")
	flag.BoolVar(&paintMode, "p", false, "alias for paint")

	flag.BoolVar(&helpMode, "help", false, "show help")
	flag.BoolVar(&helpMode, "h", false, "alias for help")
}
