package main

import (
	"flag"
	"os"

	u "art-decoder/utils"
)

func main() {
	flag.Parse()

	if helpMode {
		u.PrintUsage()
		os.Exit(0)
	}
}
