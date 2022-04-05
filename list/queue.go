package list

import "errors"

type Queue[T any] struct {
	head int
	tail int
	data []T
}

var _ List[string] = (*Queue[string])(nil)

// NewQueue create instance.
// The argument size must be a power of two.
func NewQueue[T any](size int) (Queue[T], error) {
	if !isPowerOfTwo(size) {
		var zero = Queue[T]{}
		return zero, errors.New("size must power of two")
	}

	return Queue[T]{
		head: 0,
		tail: 0,
		data: make([]T, size),
	}, nil
}

// Size return list size.
func (q *Queue[T]) Size() int {
	return q.head - q.tail
}

// IsFull is check if the list size is full.
func (q *Queue[T]) IsFull() bool {
	return q.head == (q.tail+1)&(len(q.data)-1)
}

// Add is Enqueue.
func (q *Queue[T]) Add(d T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.tail] = d

	q.tail = q.tail + 1&(len(q.data)-1)

	return nil
}

// Get is Dequeue.
func (q *Queue[T]) Get() (T, error) {
	if q.Size() == 0 {
		var zero T
		return zero, errors.New("is Empty")
	}

	result := q.data[q.head]

	q.head = q.head + 1&(len(q.data)-1)

	return result, nil
}

func isPowerOfTwo(x int) bool {
	if x == 0 {
		return false
	}

	return x&(x-1) == 0
}
