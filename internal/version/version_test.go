package version

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	origVersion, origCommit, origDate := Version, Commit, Date
	Version, Commit, Date = "1.2.3", "abc123", "2026-09-09"
	defer func() { Version, Commit, Date = origVersion, origCommit, origDate }()

	got := String()

	for _, want := range []string{"1.2.3", "abc123", "2026-09-09"} {
		if !strings.Contains(got, want) {
			t.Errorf("String() = %q, missing %q", got, want)
		}
	}
}
