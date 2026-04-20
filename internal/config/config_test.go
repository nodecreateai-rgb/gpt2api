package config

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestLoadFallsBackToExampleConfig(t *testing.T) {
	resetConfigStateForTest()

	root := t.TempDir()
	configsDir := filepath.Join(root, "configs")
	examplePath := filepath.Join(configsDir, "config.example.yaml")
	writeTestFile(t, examplePath, `
app:
  name: gpt2api
`)

	cfg, err := Load(filepath.Join(configsDir, "config.yaml"))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg == nil {
		t.Fatalf("Load() returned nil config")
	}
	if cfg.App.Name != "gpt2api" {
		t.Fatalf("cfg.App.Name = %q, want %q", cfg.App.Name, "gpt2api")
	}
}

func TestLoadMissingCustomPathFails(t *testing.T) {
	resetConfigStateForTest()

	_, err := Load(filepath.Join(t.TempDir(), "missing.yaml"))
	if err == nil {
		t.Fatalf("Load() error = nil, want non-nil")
	}
}

func resetConfigStateForTest() {
	global = nil
	once = sync.Once{}
}

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
}
