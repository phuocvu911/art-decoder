package utils

import (
	"fmt"
	"html"
	"strings"
)

// PaintLineHTML colorizes a decoded line (or multi-line string) for browser display.
// It returns HTML where each non-space rune is wrapped in a span with an inline color.
//
// Note: This is a bonus helper for the web UI. Default behavior is unchanged unless
// the web server explicitly opts into using it.
func PaintLineHTML(text string) string {
	colorMap := make(map[rune]int)
	nextColor := 0

	var sb strings.Builder
	sb.Grow(len(text) * 2)

	for _, ch := range text {
		switch ch {
		case ' ':
			sb.WriteRune(ch)
			continue
		case '\n':
			sb.WriteRune('\n')
			continue
		case '\r', '\t':
			// Keep control-ish whitespace readable, but safe.
			sb.WriteString(html.EscapeString(string(ch)))
			continue
		}

		val, ok := colorMap[ch]
		if !ok {
			val = nextColor % len(PaintColors)
			colorMap[ch] = val
			nextColor++
		}

		css := xterm256ToCSS(PaintColors[val])
		sb.WriteString(`<span class="px" style="color:`)
		sb.WriteString(css)
		sb.WriteString(`">`)
		sb.WriteString(html.EscapeString(string(ch)))
		sb.WriteString(`</span>`)
	}

	return sb.String()
}

func xterm256ToCSS(code int) string {
	r, g, b := xterm256ToRGB(code)
	return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
}

// xterm256ToRGB converts an xterm 256-color palette index into an RGB triple.
// Palette:
// - 0..15: system colors
// - 16..231: 6x6x6 color cube
// - 232..255: grayscale ramp
func xterm256ToRGB(code int) (int, int, int) {
	if code < 0 {
		code = 0
	}
	if code > 255 {
		code = 255
	}

	// Commonly-used approximations for the first 16 colors.
	system := [16][3]int{
		{0, 0, 0}, {128, 0, 0}, {0, 128, 0}, {128, 128, 0},
		{0, 0, 128}, {128, 0, 128}, {0, 128, 128}, {192, 192, 192},
		{128, 128, 128}, {255, 0, 0}, {0, 255, 0}, {255, 255, 0},
		{0, 0, 255}, {255, 0, 255}, {0, 255, 255}, {255, 255, 255},
	}
	if code < 16 {
		c := system[code]
		return c[0], c[1], c[2]
	}

	if code >= 232 {
		// 24 steps from ~8 to ~238.
		gray := 8 + (code-232)*10
		return gray, gray, gray
	}

	// 6x6x6 cube.
	c := code - 16
	r := c / 36
	g := (c % 36) / 6
	b := c % 6

	steps := [6]int{0, 95, 135, 175, 215, 255}
	return steps[r], steps[g], steps[b]
}
