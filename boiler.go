package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ANSI 256-color foreground escape. We pick from a vivid spread across the
// 256-color cube, skipping the muddy low-16 system colors.
var paintColors = []int{
	196, 202, 208, 214, 220, 226,
	118, 82, 46, 47, 48, 49,
	51, 45, 39, 33, 27, 21,
	57, 93, 129, 165, 201, 200,
	160, 124, 88, 52, 94, 130,
	136, 142, 148, 154,
}

// paintLine colorizes each non-space character with a stable per-character color.
// The colorMap persists across lines so the same character always gets the same color.
func paintLine(line string, colorMap map[rune]int, nextColor *int) string {
	var sb strings.Builder
	for _, ch := range line {
		if ch == ' ' {
			sb.WriteRune(ch)
			continue
		}
		idx, ok := colorMap[ch]
		if !ok {
			idx = *nextColor % len(paintColors)
			colorMap[ch] = idx
			*nextColor++
		}
		fmt.Fprintf(&sb, "\033[38;5;%dm%c\033[0m", paintColors[idx], ch)
	}
	return sb.String()
}

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
  --paint,  -p    Colorize output: each unique character gets its own ANSI color
  --help,   -h    Show this help

Examples:
  art-decoder "[5 #][5 -_]-[5 #]"
  art-decoder --encode "#####-_-_-_-_-_-#####"
  art-decoder --multi < file.encoded
  art-decoder --encode --multi < file.art
  art-decoder --paint "[5 #][5 -_]-[5 #]"
  art-decoder --paint --multi < plane.encoded`)
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
	var positional []string

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
			positional = append(positional, arg)
		}
	}

	colorMap := make(map[rune]int)
	nextColor := 0

	if multiMode {
		var lines []string
		scanner := bufio.NewScanner(os.Stdin)
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
				if paintMode {
					fmt.Println(paintLine(line, colorMap, &nextColor))
				} else {
					fmt.Println(line)
				}
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
		if paintMode {
			fmt.Println(paintLine(decoded, colorMap, &nextColor))
		} else {
			fmt.Println(decoded)
		}
	}
}
