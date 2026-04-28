package utils

import (
	"fmt"
	"strings"
)

// support upto 34 different colors.
var PaintColors = []int{
	196, 202, 208, 214, 220, 226,
	118, 82, 46, 47, 48, 49,
	51, 45, 39, 33, 27, 21,
	57, 93, 129, 165, 201, 200,
	160, 124, 88, 52, 94, 130,
	136, 142, 148, 154,
}

// using ptr nc we have to change the value of nc in the caller function, so we can keep track of the next color to use across multiple lines.
func PaintLine(line string, colorMap map[rune]int, nextColor *int) string {
	var sb strings.Builder
	for _, ch := range line {
		if ch == ' ' { //skip spaces
			sb.WriteRune(ch)
			continue
		}
		val, ok := colorMap[ch]
		if !ok {
			val = *nextColor % len(paintColors)
			colorMap[ch] = val
			*nextColor++
		}
		fmt.Fprintf(&sb, "\033[38;5;%dm%c\033[0m", paintColors[val], ch)
	}
	return sb.String()
}
