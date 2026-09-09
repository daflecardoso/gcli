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
	width := len(name) + 4
	top := "┌" + strings.Repeat("─", width) + "┐"
	bottom := "└" + strings.Repeat("─", width) + "┘"

	c.Println(top)
	c.Printf("│  %s  │\n", strings.ToUpper(name))
	c.Println(bottom)
}
