package utils

import (
	"fmt"
	"strings"
)

// support upto 34 different colors.
var PaintColors = []int{
	226, 46, 21, 201, 196, 51,
	154, 129, 52, 220, 47, 27,
	200, 88, 214, 48, 33, 165,
	124, 208, 49, 39, 93, 160,
	202, 118, 45, 57, 130, 82,
	136, 94, 142, 148,
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
			val = *nextColor % len(PaintColors)
			colorMap[ch] = val
			*nextColor++
		}
		fmt.Fprintf(&sb, "\033[38;5;%dm%c\033[0m", PaintColors[val], ch)
	}
	return sb.String()
}
