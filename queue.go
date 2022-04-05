package algorithm

import "errors"

type Queue[T any] struct {
	head int
	tail int
	data []T
}

func isPowerOfTwo(x int) bool {
	if x == 0 {
		return false
	}

	return x&(x-1) == 0
}

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

func (q *Queue[T]) Size() int {
	return q.head - q.tail
}

func (q *Queue[T]) IsFull() bool {
	return q.head == (q.tail+1)&(len(q.data)-1)
}

func (q *Queue[T]) Enqueue(d T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.tail] = d

	q.tail = q.tail + 1&(len(q.data)-1)

	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Size() == 0 {
		var zero T
		return zero, errors.New("is Empty")
	}

	result := q.data[q.head]

	q.head = q.head + 1&(len(q.data)-1)

	return result, nil
}
