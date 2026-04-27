package utils

import (
	"fmt"
	"strings"
)

// encode converts plain text into art-decoder notation.
func Encode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	for i := 0; i < len(input); {
		remaining := input[i:]
		bestSavings := 0
		bestCount := 0
		bestPattern := ""

		for j := 1; j <= len(remaining)/2; j++ {
			pattern := remaining[:j] //pattern can have any length
			count := 1
			for strings.HasPrefix(remaining[j*count:], pattern) { //count how many times it repeats
				count++
			}
			if count > 1 {
				originalLength := count * len(pattern)
				encodedLength := len(fmt.Sprintf("[%d %s]", count, pattern))
				savings := originalLength - encodedLength
				if savings > bestSavings { //check if it worth compressing
					bestSavings = savings
					bestCount = count
					bestPattern = pattern
				}
			}
		}

		if bestSavings > 0 {
			fmt.Fprintf(&result, "[%d %s]", bestCount, bestPattern)
			i += len(bestPattern) * bestCount
		} else {
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String()
}
