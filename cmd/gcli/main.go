package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/daflecardoso/gcli/internal/config"
	"github.com/daflecardoso/gcli/internal/gitutil"
	"github.com/daflecardoso/gcli/internal/ui"
	"github.com/daflecardoso/gcli/internal/update"
)

func main() {
	if len(os.Args) > 1 && os.Args[len(os.Args)-1] == "--update" {
		runUpdate()
		return
	}

	ui.Clear()

	cfg, err := config.Load()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("\nPlease create a %s file\n", config.FileName)
			return
		}
		fmt.Printf("\nFailed to read %s: %v\n", config.FileName, err)
		os.Exit(1)
	}

	ui.Banner(cfg.Name, cfg.Color)

	if cfg.ShowTutorial {
		ui.PrintTutorial()
	}

	if err := run(cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runUpdate() {
	fmt.Println("Updating gcli...")
	if err := update.Run(); err != nil {
		fmt.Printf("\n🔴 Update failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("\n🎉 Successfully updated")
}

func run(cfg *config.Config) error {
	commitType, err := ui.AskCommitType()
	if err != nil {
		return err
	}

	scope, err := ui.AskScope(cfg.Scopes)
	if err != nil {
		return err
	}

	message, err := ui.AskCommitMessage()
	if err != nil {
		return err
	}

	breaking, err := ui.AskBreakingChange()
	if err != nil {
		return err
	}

	commitMessage := fmt.Sprintf("%s(%s): %s", commitType, scope, message)
	if breaking != "" {
		commitMessage += fmt.Sprintf("\n\nBREAKING CHANGE: %s", breaking)
	}

	fmt.Printf("\n\x1b[33m%s\x1b[0m\n\n", commitMessage)

	looksGood, err := ui.AskConfirm("All right? ☝🏼", true)
	if err != nil {
		return err
	}
	if !looksGood {
		return nil
	}

	fmt.Println("✅ git add .")
	if err := gitutil.AddAll(); err != nil {
		return fmt.Errorf("git add failed: %w", err)
	}

	fmt.Printf("✅ git commit -m %q\n", commitMessage)
	if err := gitutil.Commit(commitMessage); err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}

	fmt.Println("✅ git push")
	if err := gitutil.Push(); err != nil {
		fmt.Printf("🔴 git push\n\n%v\n", err)
	}

	return nil
}
