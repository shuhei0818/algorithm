package list

import "errors"

type Stack[T any] struct {
	top   int
	array []T
}

var _ List[string] = (*Stack[string])(nil)

func NewStack[T any](size int) Stack[T] {
	s := Stack[T]{
		top:   0,
		array: make([]T, size),
	}
	return s
}

// Add is Push.
func (s *Stack[T]) Add(data T) error {
	if s.top > len(s.array) {
		return errors.New("stack overflow happend")
	}

	s.array[s.top] = data
	s.top++

	return nil
}

// Get is Pop.
func (s *Stack[T]) Get() (T, error) {

	if s.top <= 0 {
		var zero T
		return zero, errors.New("no data")
	}

	s.top--

	return s.array[s.top], nil
}
