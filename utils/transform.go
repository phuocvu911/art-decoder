package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// decode expands encoded art notation into plain text.
func Decode(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("input line is empty")
	}
	var result strings.Builder

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '[':
			end := strings.Index(input[i:], "]")
			if end == -1 {
				return "", fmt.Errorf("unbalanced brackets")
			}
			end += i

			inner := input[i+1 : end]

			countStr, pattern, ok := strings.Cut(inner, " ")
			if !ok {
				return "", fmt.Errorf("missing space separator")
			}

			count, err := strconv.Atoi(countStr)
			if err != nil || count < 0 {
				return "", fmt.Errorf("invalid number: %q", countStr)
			}

			if pattern == "" {
				return "", fmt.Errorf("empty pattern")
			}

			//dont use strings.Repeat to avoid unnecessary memory allocation.
			for range count {
				result.WriteString(pattern)
			}
			i = end
		case ']':
			return "", fmt.Errorf("unbalanced closing bracket")
		default:
			result.WriteByte(input[i])
		}
	}
	return result.String(), nil
}

// encode converts plain text into art-decoder notation.
func Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				fmt.Fprintf(&result, "[%d %c]", count, input[i-1])
			} else {
				result.WriteByte(input[i-1])
			}
			count = 1
		}
	}

	// Handle the last run
	if count > 1 {
		fmt.Fprintf(&result, "[%d %c]", count, input[len(input)-1])
	} else {
		result.WriteByte(input[len(input)-1])
	}

	return result.String()
}
