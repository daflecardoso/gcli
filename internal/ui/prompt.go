package ui

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

var commitTypes = []string{"feat", "fix", "docs", "style", "refactor", "test", "chore"}

// customScopeLabel is the always-first entry in the scope select, letting
// the user type an arbitrary scope instead of picking a configured one.
const customScopeLabel = "Custom"

// scopeSelectTemplate is SelectQuestionTemplate with the custom-scope
// option rendered in a different color so it stands out from the
// configured scopes.
var scopeSelectTemplate = `
{{- define "option"}}
    {{- if eq .SelectedIndex .CurrentIndex }}{{color .Config.Icons.SelectFocus.Format }}{{ .Config.Icons.SelectFocus.Text }} {{else}}{{color "default"}}  {{end}}
    {{- if eq .CurrentOpt.Value "` + customScopeLabel + `"}}{{color "magenta"}}{{end}}
    {{- .CurrentOpt.Value}}{{ if ne ($.GetDescription .CurrentOpt) "" }} - {{color "cyan"}}{{ $.GetDescription .CurrentOpt }}{{end}}
    {{- color "reset"}}
{{end}}
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{.Answer}}{{color "reset"}}{{"\n"}}
{{- else}}
  {{- "  "}}{{- color "cyan"}}[Use arrows to move, type to filter{{- if and .Help (not .ShowHelp)}}, {{ .Config.HelpInput }} for more help{{end}}]{{color "reset"}}
  {{- "\n"}}
  {{- range $ix, $option := .PageEntries}}
    {{- template "option" $.IterateOption $ix $option}}
  {{- end}}
{{- end}}`

var slugSeparators = regexp.MustCompile(`\s+`)

// slugify normalizes freely typed text into a commit-scope-friendly slug:
// trimmed, lowercased, with internal whitespace collapsed to hyphens.
func slugify(s string) string {
	s = strings.TrimSpace(s)
	s = slugSeparators.ReplaceAllString(s, "-")
	return strings.ToLower(s)
}

const tutorial = `
feat: (new feature for the user, not a new feature for build script)
fix: (bug fix for the user, not a fix to a build script)
docs: (changes to the documentation)
style: (formatting, missing semi colons, etc; no production code change)
refactor: (refactoring production code, eg. renaming a variable)
test: (adding missing tests, refactoring tests; no production code change)
chore: (updating grunt tasks etc; no production code change)
`

func PrintTutorial() {
	fmt.Print(tutorial)
}

func AskCommitType() (string, error) {
	var answer string
	prompt := &survey.Select{
		Message: "What is commit type?",
		Options: commitTypes,
	}
	err := survey.AskOne(prompt, &answer)
	return answer, err
}

// AskScope prompts the user to pick a scope. "Custom" is always the first
// option, rendered in a different color, and lets the user type a scope
// that isn't in the configured list.
func AskScope(scopes []string) (string, error) {
	options := append([]string{customScopeLabel}, scopes...)

	prevTemplate := survey.SelectQuestionTemplate
	survey.SelectQuestionTemplate = scopeSelectTemplate
	defer func() { survey.SelectQuestionTemplate = prevTemplate }()

	var answer string
	prompt := &survey.Select{
		Message: "What is scope?",
		Options: options,
	}
	if err := survey.AskOne(prompt, &answer); err != nil {
		return "", err
	}

	if answer == customScopeLabel {
		return askCustomScope()
	}
	return answer, nil
}

// askCustomScope prompts for a freely typed scope and slugifies it so it
// fits cleanly inside a commit's `type(scope):` parentheses.
func askCustomScope() (string, error) {
	var answer string
	prompt := &survey.Input{
		Message: "Type the scope:",
	}
	if err := survey.AskOne(prompt, &answer, survey.WithValidator(survey.Required)); err != nil {
		return "", err
	}
	return slugify(answer), nil
}

func AskCommitMessage() (string, error) {
	var answer string
	prompt := &survey.Input{
		Message: "Now type what did you do:",
	}
	err := survey.AskOne(prompt, &answer, survey.WithValidator(survey.Required))
	return answer, err
}

func AskBreakingChange() (string, error) {
	var answer string
	prompt := &survey.Input{
		Message: "Breaking change:",
	}
	err := survey.AskOne(prompt, &answer)
	return answer, err
}

func AskConfirm(message string, defaultValue bool) (bool, error) {
	answer := defaultValue
	prompt := &survey.Confirm{
		Message: message,
		Default: defaultValue,
	}
	err := survey.AskOne(prompt, &answer)
	return answer, err
}
