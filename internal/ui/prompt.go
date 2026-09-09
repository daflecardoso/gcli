package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

var commitTypes = []string{"feat", "fix", "docs", "style", "refactor", "test", "chore"}

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

func AskScope(scopes []string) (string, error) {
	var answer string
	prompt := &survey.Select{
		Message: "What is scope?",
		Options: scopes,
	}
	err := survey.AskOne(prompt, &answer)
	return answer, err
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
