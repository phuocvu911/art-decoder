package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"

	u "art-decoder/utils"
)

func main() {
	flag.Parse()

	if helpMode {
		u.PrintUsage()
		os.Exit(0)
	}
	args := flag.Args()

	colorMap := make(map[rune]int)
	nextColor := rand.IntN(len(u.PaintColors))

	if multiMode {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				os.Exit(1)
			}

			if encodeMode {
				fmt.Println(u.Encode(line))
			} else {
				decoded, err := u.Decode(line)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error")
					//fmt.Fprintln(os.Stderr, err)
					continue //log only that error line and continue the rest
				}
				if paintMode {
					fmt.Println(u.PaintLine(decoded, colorMap, &nextColor))
				} else {
					fmt.Println(decoded)
				}
			}
		}
	} else { // single mode, accept only 1 arg.
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "Error: expected 1 argument, got", len(args))
			u.PrintUsage()
			os.Exit(1)
		}
		if encodeMode {
			fmt.Println(u.Encode(args[0]))
		} else {
			decoded, err := u.Decode(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error")
				//fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if paintMode {
				fmt.Println(u.PaintLine(decoded, colorMap, &nextColor))
			} else {
				fmt.Println(decoded)
			}
		}
	}
}
