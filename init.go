package main

import "flag"

var encodeMode bool
var multiMode bool
var paintMode bool
var helpMode bool

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
