package queue

type Queue[T any] struct {
	inbox  []T
	outbox []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.inbox = append(q.inbox, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.outbox) == 0 {
		for i := len(q.inbox) - 1; i >= 0; i-- {
			q.outbox = append(q.outbox, q.inbox[i])
		}
		q.inbox = nil
	}

	if len(q.outbox) == 0 {
		var zero T
		return zero
	}

	item := q.outbox[len(q.outbox)-1]
	q.outbox = q.outbox[:len(q.outbox)-1]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.inbox) == 0 && len(q.outbox) == 0
}
