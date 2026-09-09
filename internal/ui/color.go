package ui

import (
	"fmt"
	"strconv"
	"strings"
)

// Printer prints text wrapped in a 24-bit ANSI truecolor escape sequence.
type Printer struct {
	prefix string
}

func (p Printer) Println(a ...any) {
	fmt.Print(p.prefix)
	fmt.Print(a...)
	fmt.Println("\x1b[0m")
}

func (p Printer) Printf(format string, a ...any) {
	fmt.Print(p.prefix)
	fmt.Printf(format, a...)
	fmt.Print("\x1b[0m")
}

// HexColor returns a Printer for the given "#RRGGBB" string, falling back
// to the default terminal color if it can't be parsed.
func HexColor(hex string) Printer {
	r, g, b, ok := parseHex(hex)
	if !ok {
		return Printer{}
	}
	return Printer{prefix: fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)}
}

func parseHex(hex string) (r, g, b int, ok bool) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, false
	}

	values := make([]int64, 3)
	for i := range values {
		v, err := strconv.ParseInt(hex[i*2:i*2+2], 16, 32)
		if err != nil {
			return 0, 0, 0, false
		}
		values[i] = v
	}

	return int(values[0]), int(values[1]), int(values[2]), true
}

func init() {
	enableVirtualTerminal()
}
