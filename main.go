package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Fprintln(os.Stderr, `Usage:
  ./art-decoder [flags] "<encoded_string>"
  ./art-decoder --multi [flags]   (reads lines from stdin)

Flags:
  --encode, -e    Encode plain text into art-decoder notation
  --multi,  -m    Read multiple lines from stdin
  --paint,  -p    Colorize output: each unique character gets its own ANSI color
  --help,   -h    Show this help

Examples:
  ./art-decoder "[5 #][5 -_]-[5 #]"
  ./art-decoder --encode "#####-_-_-_-_-_-#####"
  ./art-decoder --multi < file.encoded
  ./art-decoder --encode --multi < file.art
  ./art-decoder --paint "[5 #][5 -_]-[5 #]"
  ./art-decoder --paint --multi < file.encoded`)
}

func main() {
	input := os.Args[1:]

	if len(input) == 0 {
		printUsage()
		os.Exit(1)
	}

	encodeMode := false
	multiMode := false
	paintMode := false

}
