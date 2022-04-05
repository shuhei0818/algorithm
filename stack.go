package algorithm

import "errors"

type Stack[T any] struct {
	top   int
	array []T
}

func NewStack[T any](size int) Stack[T] {
	s := Stack[T]{
		top:   0,
		array: make([]T, size),
	}
	return s
}

func (s *Stack[T]) Push(data T) error {
	if s.top > len(s.array) {
		return errors.New("stack overflow happend")
	}

	s.array[s.top] = data
	s.top++

	return nil
}

func (s *Stack[T]) Pop() (T, error) {

	if s.top <= 0 {
		var zero T
		return zero, errors.New("no data")
	}

	s.top--

	return s.array[s.top], nil
}
