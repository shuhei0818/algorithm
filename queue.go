package algorithm

import "errors"

type Queue[T any] struct {
	head int
	tail int
	data []T
}

const size = 4
const mask = size - 1

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		head: 0,
		tail: 0,
		data: make([]T, size),
	}
}

func (q *Queue[T]) Size() int {
	return q.head - q.tail
}

func (q *Queue[T]) IsFull() bool {
	return q.head == (q.tail+1)&mask
}

func (q *Queue[T]) Enqueue(d T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.tail] = d

	q.tail = q.tail + 1&mask

	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Size() == 0 {
		var zero T
		return zero, errors.New("is Empty")
	}

	result := q.data[q.head]

	q.head = q.head + 1&mask

	return result, nil
}
