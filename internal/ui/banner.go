package ui

import (
	"fmt"
	"strings"
)

// Clear clears the terminal screen using an ANSI escape sequence. On
// Windows this relies on enableVirtualTerminal having run first.
func Clear() {
	fmt.Print("\x1b[H\x1b[2J")
}

// Banner prints name inside a colored box, replacing the figlet-style
// ASCII art banner from the Node.js version with a simpler, dependency-free
// equivalent.
func Banner(name, hex string) {
	c := HexColor(hex)
	top, middle, bottom := boxLines(name)

	c.Println(top)
	c.Println(middle)
	c.Println(bottom)
}

// boxLines builds the three lines of the banner box around name. It is
// split out from Banner so the layout can be unit tested without capturing
// stdout.
func boxLines(name string) (top, middle, bottom string) {
	width := len(name) + 4
	top = "┌" + strings.Repeat("─", width) + "┐"
	middle = fmt.Sprintf("│  %s  │", strings.ToUpper(name))
	bottom = "└" + strings.Repeat("─", width) + "┘"
	return top, middle, bottom
}
