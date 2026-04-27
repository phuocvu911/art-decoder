package main

import (
	"bufio"
	"flag"
	"fmt"
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

	if multiMode {
		lines := make([]string, 0)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		for _, line := range lines {
			fmt.Println(u.Decode(line))
		}
	} else {
		for _, arg := range args {
			fmt.Println(u.Decode(arg))
		}
	}
}
