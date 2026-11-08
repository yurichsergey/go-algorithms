package store

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestLoadSaveRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "users.json")
	want := []string{"alice", "bob", "charlie"}

	if err := Save(path, want); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if !slices.Equal(got, want) {
		t.Errorf("Load() = %v, want %v", got, want)
	}
}

func TestLoadNonExistentFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nonexistent.json")

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if len(got) != 0 {
		t.Errorf("Load() = %v, want empty slice", got)
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "invalid.json")
	if err := os.WriteFile(path, []byte("invalid json"), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if _, err := Load(path); err == nil {
		t.Fatalf("Load() expected error for invalid JSON, got nil")
	}
}
