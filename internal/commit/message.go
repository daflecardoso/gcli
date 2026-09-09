// Package commit builds Conventional Commits messages
// (https://www.conventionalcommits.org).
package commit

import "fmt"

// Build assembles a Conventional Commits message from its parts. breaking
// is optional; when empty, no BREAKING CHANGE footer is appended.
func Build(commitType, scope, message, breaking string) string {
	msg := fmt.Sprintf("%s(%s): %s", commitType, scope, message)
	if breaking != "" {
		msg += fmt.Sprintf("\n\nBREAKING CHANGE: %s", breaking)
	}
	return msg
}
