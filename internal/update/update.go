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
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		script := fmt.Sprintf("iwr -useb %s/install.ps1 | iex", rawBaseURL)
		cmd = exec.Command("powershell", "-NoProfile", "-Command", script)
	} else {
		script := fmt.Sprintf("curl -fsSL %s/install.sh | sh", rawBaseURL)
		cmd = exec.Command("sh", "-c", script)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
