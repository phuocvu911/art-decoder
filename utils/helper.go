package utils

import (
	"fmt"
	"os"
)

func PrintUsage() {
	fmt.Fprintln(os.Stderr, `Usage:
  ./art-decoder [flags] "<encoded_string>"
  ./art-decoder --multi [flags] < <encoded_file>

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
