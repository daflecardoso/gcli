package commit

import "testing"

func TestBuild(t *testing.T) {
	tests := []struct {
		name       string
		commitType string
		scope      string
		message    string
		breaking   string
		want       string
	}{
		{
			name:       "without breaking change",
			commitType: "feat",
			scope:      "core",
			message:    "add login flow",
			breaking:   "",
			want:       "feat(core): add login flow",
		},
		{
			name:       "with breaking change",
			commitType: "fix",
			scope:      "utils",
			message:    "correct date parsing",
			breaking:   "changes the DateParser signature",
			want:       "fix(utils): correct date parsing\n\nBREAKING CHANGE: changes the DateParser signature",
		},
		{
			name:       "empty scope still formats",
			commitType: "chore",
			scope:      "",
			message:    "bump deps",
			breaking:   "",
			want:       "chore(): bump deps",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Build(tt.commitType, tt.scope, tt.message, tt.breaking)
			if got != tt.want {
				t.Errorf("Build() = %q, want %q", got, tt.want)
			}
		})
	}
}
