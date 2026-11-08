package store

import (
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