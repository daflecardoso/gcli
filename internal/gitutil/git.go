package gitutil

import (
	"os"
	"os/exec"
)

// Run executes a git subcommand, streaming its output to the current
// process's stdout/stderr, and returns an error if it exits non-zero.
func Run(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func AddAll() error {
	return Run("add", ".")
}

func Commit(message string) error {
	return Run("commit", "-m", message)
}

func Push() error {
	return Run("push")
}

func Pull() error {
	return Run("pull")
}
