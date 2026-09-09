package update

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const rawBaseURL = "https://raw.githubusercontent.com/daflecardoso/gcli/main"

// Run re-downloads and re-runs the install script, which always fetches
// the latest release binary for the current OS/arch and replaces gcli
// in place.
func Run() error {
	cmd := commandFor(runtime.GOOS)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// commandFor builds the shell invocation that re-runs the installer for
// goos, split out from Run so it can be unit tested without touching the
// network or the current process's file descriptors.
func commandFor(goos string) *exec.Cmd {
	if goos == "windows" {
		script := fmt.Sprintf("iwr -useb %s/install.ps1 | iex", rawBaseURL)
		return exec.Command("powershell", "-NoProfile", "-Command", script)
	}

	script := fmt.Sprintf("curl -fsSL %s/install.sh | sh", rawBaseURL)
	return exec.Command("sh", "-c", script)
}
