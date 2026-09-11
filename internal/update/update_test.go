package update

import (
	"strings"
	"testing"
)

func TestCommandForWindows(t *testing.T) {
	cmd := commandFor("windows", "")

	if len(cmd.Args) < 2 || cmd.Args[0] != "powershell" {
		t.Errorf("commandFor(\"windows\", \"\") args = %v, want it to start with powershell", cmd.Args)
	}
	joined := cmd.Args[len(cmd.Args)-1]
	if !strings.Contains(joined, "install.ps1") {
		t.Errorf("commandFor(\"windows\", \"\") script = %q, want it to reference install.ps1", joined)
	}
}

func TestCommandForUnix(t *testing.T) {
	for _, goos := range []string{"linux", "darwin"} {
		cmd := commandFor(goos, "")

		if len(cmd.Args) < 2 || cmd.Args[0] != "sh" {
			t.Errorf("commandFor(%q, \"\") args = %v, want it to start with sh", goos, cmd.Args)
		}
		joined := cmd.Args[len(cmd.Args)-1]
		if !strings.Contains(joined, "install.sh") {
			t.Errorf("commandFor(%q, \"\") script = %q, want it to reference install.sh", goos, joined)
		}
	}
}

func TestCommandForNoVersionLeavesEnvUnset(t *testing.T) {
	cmd := commandFor("darwin", "")

	if cmd.Env != nil {
		t.Errorf("commandFor(%q, \"\") Env = %v, want nil so the child inherits the default environment", "darwin", cmd.Env)
	}
}

func TestCommandForVersionSetsEnv(t *testing.T) {
	for _, goos := range []string{"linux", "darwin", "windows"} {
		cmd := commandFor(goos, "v1.2.3")

		found := false
		for _, e := range cmd.Env {
			if e == "GCLI_VERSION=v1.2.3" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("commandFor(%q, \"v1.2.3\") Env = %v, want it to contain GCLI_VERSION=v1.2.3", goos, cmd.Env)
		}
	}
}
