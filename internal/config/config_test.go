package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func chdir(t *testing.T, dir string) {
	t.Helper()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(orig); err != nil {
			t.Fatal(err)
		}
	})
}

func TestLoad(t *testing.T) {
	dir := t.TempDir()
	chdir(t, dir)

	content := `{
		"name": "gcli",
		"color": "#A020F0",
		"showTutorial": true,
		"scopes": ["core", "utils"]
	}`
	if err := os.WriteFile(filepath.Join(dir, FileName), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.Name != "gcli" {
		t.Errorf("Name = %q, want %q", cfg.Name, "gcli")
	}
	if cfg.Color != "#A020F0" {
		t.Errorf("Color = %q, want %q", cfg.Color, "#A020F0")
	}
	if !cfg.ShowTutorial {
		t.Error("ShowTutorial = false, want true")
	}
	if want := []string{"core", "utils"}; len(cfg.Scopes) != len(want) || cfg.Scopes[0] != want[0] || cfg.Scopes[1] != want[1] {
		t.Errorf("Scopes = %v, want %v", cfg.Scopes, want)
	}
}

func TestLoadMissingFile(t *testing.T) {
	chdir(t, t.TempDir())

	_, err := Load()
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Load() error = %v, want os.ErrNotExist", err)
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	chdir(t, dir)

	if err := os.WriteFile(filepath.Join(dir, FileName), []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want a JSON parse error")
	}
}
