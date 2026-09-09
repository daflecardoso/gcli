package ui

import "testing"

func TestParseHex(t *testing.T) {
	tests := []struct {
		name   string
		hex    string
		wantR  int
		wantG  int
		wantB  int
		wantOK bool
	}{
		{"with hash", "#A020F0", 160, 32, 240, true},
		{"without hash", "A020F0", 160, 32, 240, true},
		{"black", "#000000", 0, 0, 0, true},
		{"white", "#FFFFFF", 255, 255, 255, true},
		{"too short", "#FFF", 0, 0, 0, false},
		{"not hex", "#GGGGGG", 0, 0, 0, false},
		{"empty", "", 0, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b, ok := parseHex(tt.hex)
			if ok != tt.wantOK {
				t.Fatalf("parseHex(%q) ok = %v, want %v", tt.hex, ok, tt.wantOK)
			}
			if !tt.wantOK {
				return
			}
			if r != tt.wantR || g != tt.wantG || b != tt.wantB {
				t.Errorf("parseHex(%q) = (%d,%d,%d), want (%d,%d,%d)", tt.hex, r, g, b, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}
}

func TestHexColorFallback(t *testing.T) {
	p := HexColor("not-a-color")
	if p.prefix != "" {
		t.Errorf("HexColor with invalid hex should have empty prefix, got %q", p.prefix)
	}
}
