package cmd

import (
	"07_cli_task_manager/db"
	"bytes"
	"io"
	"os"
	"testing"
)

func setupTestDB(t *testing.T) func() {
	t.Helper()
	f, err := os.CreateTemp("", "tasks-test-*.db")
	if err != nil {
		t.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		return nil
	}
	db.DBPath = f.Name()
	return func() {
		if err := os.Remove(f.Name()); err != nil {
			return
		}
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	err := w.Close()
	if err != nil {
		return ""
	}
	os.Stdout = old
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		return ""
	}
	return buf.String()
}

func execute(args ...string) string {
	return captureOutput(func() {
		rootCmd.SetArgs(args)
		if err := rootCmd.Execute(); err != nil {
			return
		}
	})
}

func TestAddCommand(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	out := execute("add", "clean dishes")
	if out == "" {
		t.Error("expected output from add command, got none")
	}
}

func TestListCommand(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	execute("add", "clean dishes")
	execute("add", "review proposal")

	out := execute("list")
	if out == "" {
		t.Error("expected output from list command, got none")
	}
}

func TestListCommandEmpty(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	out := execute("list")
	if out == "" {
		t.Error("expected output from list command, got none")
	}
}

func TestDoCommand(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	execute("add", "clean dishes")
	out := execute("do", "1")
	if out == "" {
		t.Error("expected output from do command, got none")
	}
}
