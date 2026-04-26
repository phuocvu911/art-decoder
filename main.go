package main

import (
	u "art-decoder/utils"
	"fmt"
	"os"
	"strings"
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
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	encodeMode := false
	multiMode := false
	paintMode := false
	var input string

	for _, arg := range args {
		switch arg {
		case "--encode", "-e":
			encodeMode = true
		case "--multi", "-m":
			multiMode = true
		case "--paint", "-p":
			paintMode = true
		case "--help", "-h":
			printUsage()
			os.Exit(0)
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Fprintf(os.Stderr, "Unknown flag: %s\n", arg)
				os.Exit(1)
			}
			input = arg
		}
	}

	if encodeMode {
		fmt.Println(u.Encode(input))
	} else {
		decoded, err := u.Decode(input)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		fmt.Println(decoded)
	}
	if paintMode {
		colorMap := make(map[rune]int)
		nextColor := 0
		fmt.Println(boiler.PaintLine(decoded, colorMap, &nextColor))
	} else {
		fmt.Println(decoded)
	}
}
