//go:build windows

package ui

import (
	"os"
	"syscall"
	"unsafe"
)

const enableVirtualTerminalProcessing = 0x0004

// enableVirtualTerminal turns on ANSI escape sequence processing for the
// classic Windows console, so colors and cursor codes render instead of
// showing up as raw escape codes. Windows Terminal already supports this,
// but cmd.exe and older PowerShell hosts need it turned on explicitly.
func enableVirtualTerminal() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getMode := kernel32.NewProc("GetConsoleMode")
	setMode := kernel32.NewProc("SetConsoleMode")

	handle := syscall.Handle(os.Stdout.Fd())

	var mode uint32
	ret, _, _ := getMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	if ret == 0 {
		return
	}

	setMode.Call(uintptr(handle), uintptr(mode|enableVirtualTerminalProcessing))
}
