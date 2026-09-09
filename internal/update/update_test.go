package update

import (
	"strings"
	"testing"
)

func TestCommandForWindows(t *testing.T) {
	cmd := commandFor("windows")

	if len(cmd.Args) < 2 || cmd.Args[0] != "powershell" {
		t.Errorf("commandFor(\"windows\") args = %v, want it to start with powershell", cmd.Args)
	}
	joined := cmd.Args[len(cmd.Args)-1]
	if !strings.Contains(joined, "install.ps1") {
		t.Errorf("commandFor(\"windows\") script = %q, want it to reference install.ps1", joined)
	}
}

func TestCommandForUnix(t *testing.T) {
	for _, goos := range []string{"linux", "darwin"} {
		cmd := commandFor(goos)

		if len(cmd.Args) < 2 || cmd.Args[0] != "sh" {
			t.Errorf("commandFor(%q) args = %v, want it to start with sh", goos, cmd.Args)
		}
		joined := cmd.Args[len(cmd.Args)-1]
		if !strings.Contains(joined, "install.sh") {
			t.Errorf("commandFor(%q) script = %q, want it to reference install.sh", goos, joined)
		}
	}
}
