package db

import (
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
	DBPath = f.Name()
	return func() {
		err := os.Remove(f.Name())
		if err != nil {
			return
		}
	}
}

func TestAddTask(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	err := AddTask("clean dishes")
	if err != nil {
		t.Fatalf("AddTask failed: %v", err)
	}
}

func TestListTasks(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	err := AddTask("clean dishes")
	if err != nil {
		return
	}
	err = AddTask("review proposal")
	if err != nil {
		return
	}

	tasks, err := ListTasks()
	if err != nil {
		t.Fatalf("ListTasks failed: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}
}

func TestCompleteTask(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	err := AddTask("clean dishes")
	if err != nil {
		return
	}

	err = CompleteTask(1)
	if err != nil {
		t.Fatalf("CompleteTask failed: %v", err)
	}

	tasks, _ := ListTasks()
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks after completion, got %d", len(tasks))
	}
}
