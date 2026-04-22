package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// decode expands encoded art notation into plain text.
func decode(input string) (string, error) {
	var result strings.Builder
	i := 0
	for i < len(input) {
		if input[i] == '[' {
			end := strings.Index(input[i:], "]")
			if end == -1 {
				return "", fmt.Errorf("unbalanced brackets")
			}
			end += i

			inner := input[i+1 : end]

			spaceIdx := strings.Index(inner, " ")
			if spaceIdx == -1 {
				return "", fmt.Errorf("missing space separator")
			}

			countStr := inner[:spaceIdx]
			pattern := inner[spaceIdx+1:]

			count, err := strconv.Atoi(countStr)
			if err != nil || count < 0 {
				return "", fmt.Errorf("invalid number: %q", countStr)
			}

			if pattern == "" {
				return "", fmt.Errorf("empty pattern")
			}

			for k := 0; k < count; k++ {
				result.WriteString(pattern)
			}

			i = end + 1
		} else if input[i] == ']' {
			return "", fmt.Errorf("unbalanced closing bracket")
		} else {
			result.WriteByte(input[i])
			i++
		}
	}
	return result.String(), nil
}

// encode converts plain text into art-decoder notation.
func encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	i := 0
	for i < len(input) {
		remaining := input[i:]
		bestSaving := 0
		bestCount := 1
		bestUnit := string(input[i])

		for unitLen := 1; unitLen <= len(remaining)/2; unitLen++ {
			unit := remaining[:unitLen]
			count := 1
			for j := unitLen; j+unitLen <= len(remaining); j += unitLen {
				if remaining[j:j+unitLen] == unit {
					count++
				} else {
					break
				}
			}
			if count > 1 {
				expanded := unitLen * count
				encoded := len(fmt.Sprintf("[%d %s]", count, unit))
				saving := expanded - encoded
				if saving > bestSaving {
					bestSaving = saving
					bestCount = count
					bestUnit = unit
				}
			}
		}

		if bestSaving > 0 {
			result.WriteString(fmt.Sprintf("[%d %s]", bestCount, bestUnit))
			i += len(bestUnit) * bestCount
		} else {
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String()
}

func decodeMultiLine(lines []string) ([]string, error) {
	result := make([]string, len(lines))
	for idx, line := range lines {
		decoded, err := decode(line)
		if err != nil {
			return nil, err
		}
		result[idx] = decoded
	}
	return result, nil
}

func encodeMultiLine(lines []string) []string {
	result := make([]string, len(lines))
	for idx, line := range lines {
		result[idx] = encode(line)
	}
	return result
}

func printUsage() {
	fmt.Fprintln(os.Stderr, `Usage:
  art-decoder [flags] "<encoded_string>"
  art-decoder --multi [flags]   (reads lines from stdin)

Flags:
  --encode, -e    Encode plain text into art-decoder notation
  --multi,  -m    Read multiple lines from stdin
  --help,   -h    Show this help

Examples:
  art-decoder "[5 #][5 -_]-[5 #]"
  art-decoder --encode "#####-_-_-_-_-_-#####"
  art-decoder --multi < file.encoded
  art-decoder --encode --multi < file.art`)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

	encodeMode := false
	multiMode := false
	var positional []string

	for _, arg := range args {
		switch arg {
		case "--encode", "-e":
			encodeMode = true
		case "--multi", "-m":
			multiMode = true
		case "--help", "-h":
			printUsage()
			os.Exit(0)
		default:
			if strings.HasPrefix(arg, "-") {
				fmt.Fprintf(os.Stderr, "Unknown flag: %s\n", arg)
				os.Exit(1)
			}
			positional = append(positional, arg)
		}
	}

	if multiMode {
		var lines []string
		scanner := bufio.NewScanner(os.Stdin)
		// Increase scanner buffer for long lines (large art files)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if encodeMode {
			for _, line := range encodeMultiLine(lines) {
				fmt.Println(line)
			}
		} else {
			decoded, err := decodeMultiLine(lines)
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
			for _, line := range decoded {
				fmt.Println(line)
			}
		}
		return
	}

	if len(positional) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input string provided")
		printUsage()
		os.Exit(1)
	}

	input := positional[0]

	if encodeMode {
		fmt.Println(encode(input))
	} else {
		decoded, err := decode(input)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		fmt.Println(decoded)
	}
}
