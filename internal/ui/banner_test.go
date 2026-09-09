package ui

import (
	"strings"
	"testing"
)

func TestBoxLines(t *testing.T) {
	top, middle, bottom := boxLines("gcli")

	if !strings.HasPrefix(top, "┌") || !strings.HasSuffix(top, "┐") {
		t.Errorf("top = %q, want it wrapped in ┌ ┐", top)
	}
	if !strings.HasPrefix(bottom, "└") || !strings.HasSuffix(bottom, "┘") {
		t.Errorf("bottom = %q, want it wrapped in └ ┘", bottom)
	}
	if !strings.Contains(middle, "GCLI") {
		t.Errorf("middle = %q, want it to contain the uppercased name", middle)
	}
	if len(top) != len(bottom) {
		t.Errorf("top and bottom have different widths: %d vs %d", len(top), len(bottom))
	}
}
