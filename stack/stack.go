package stack

import "errors"

// ErrEmptyStack is returned when an operation cannot be performed on an empty collection
var ErrEmptyStack = errors.New("stack is empty")

// Stack is a generic Last-In-First-Out (LIFO) data structure
type Stack[T any] struct {
	array []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		array: make([]T, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.array = append(s.array, value)
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
func (s *Stack[T]) Pop() (T, error) {
	var value T

	if len(s.array) == 0 {
		return value, ErrEmptyStack
	}

	value = s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]

	return value, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
func (s *Stack[T]) Peek() (T, error) {
	var value T
	if len(s.array) == 0 {
		return value, ErrEmptyStack
	}

	return s.array[len(s.array)-1], nil
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.array)
}

// IsEmpty returns true if the stack contains no elements
func (s *Stack[T]) IsEmpty() bool {
	if len(s.array) == 0 {
		return true
	}
	return false
}
