package update

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const rawBaseURL = "https://raw.githubusercontent.com/daflecardoso/gcli/main"

// Run re-downloads and re-runs the install script, which replaces gcli in
// place with the given release. An empty version installs the latest
// release.
func Run(version string) error {
	cmd := commandFor(runtime.GOOS, version)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// commandFor builds the shell invocation that re-runs the installer for
// goos and version, split out from Run so it can be unit tested without
// touching the network or the current process's file descriptors. The
// install scripts read the target version from GCLI_VERSION, defaulting to
// the latest release when it's unset.
func commandFor(goos, version string) *exec.Cmd {
	var cmd *exec.Cmd
	if goos == "windows" {
		script := fmt.Sprintf("iwr -useb %s/install.ps1 | iex", rawBaseURL)
		cmd = exec.Command("powershell", "-NoProfile", "-Command", script)
	} else {
		script := fmt.Sprintf("curl -fsSL %s/install.sh | sh", rawBaseURL)
		cmd = exec.Command("sh", "-c", script)
	}

	if version != "" {
		cmd.Env = append(os.Environ(), "GCLI_VERSION="+version)
	}
	return cmd
}
