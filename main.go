package main

import (
	u "art-decoder/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		u.PrintUsage()
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
			u.PrintUsage()
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
}
