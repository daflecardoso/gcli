package ui

import "testing"

func TestSlugify(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "single word",
			in:   "core",
			want: "core",
		},
		{
			name: "spaces become hyphens",
			in:   "custom scope typed",
			want: "custom-scope-typed",
		},
		{
			name: "trims surrounding whitespace",
			in:   "  api  ",
			want: "api",
		},
		{
			name: "lowercases",
			in:   "API Gateway",
			want: "api-gateway",
		},
		{
			name: "collapses repeated whitespace",
			in:   "billing   service",
			want: "billing-service",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slugify(tt.in)
			if got != tt.want {
				t.Errorf("slugify(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
