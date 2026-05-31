package queue

import "testing"

func TestFIFOOrder(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for _, expected := range []int{1, 2, 3} {
		if got := q.Dequeue(); got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	q := Queue[int]{}
	if !q.IsEmpty() {
		t.Error("expected empty queue")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("expected non-empty queue")
	}
	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("expected empty queue after dequeue")
	}
}

func TestDequeueFromEmpty(t *testing.T) {
	q := Queue[int]{}
	if got := q.Dequeue(); got != 0 {
		t.Errorf("expected zero value 0, got %d", got)
	}
}

func TestInterleaved(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)

	if got := q.Dequeue(); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}

	q.Enqueue(3)

	if got := q.Dequeue(); got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
	if got := q.Dequeue(); got != 3 {
		t.Errorf("expected 3, got %d", got)
	}
}

func TestStringType(t *testing.T) {
	q := Queue[string]{}
	q.Enqueue("a")
	q.Enqueue("b")

	if got := q.Dequeue(); got != "a" {
		t.Errorf("expected a, got %s", got)
	}
	if got := q.Dequeue(); got != "b" {
		t.Errorf("expected b, got %s", got)
	}
}
