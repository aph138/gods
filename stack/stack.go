// Implementation of the Stack data structure.
// Stack is a linear data structure that follows LIFO (Last-In-First-Out) principle.
package stack

import "errors"

var ErrEmpty = errors.New("stack is empty")

type Stack[T comparable] struct {
	items []T
}

// NewStack returns a new Stack. You can pass initial value to it.
func NewStack[T comparable](i ...T) *Stack[T] {
	l := []T{}
	l = append(l, i...)
	return &Stack[T]{
		items: l,
	}
}

// Add adds an element to the stack
func (s *Stack[T]) Add(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the last element of the stack
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		return *new(T), ErrEmpty
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, nil
}

// Peek returns the last element of the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		return *new(T), ErrEmpty
	}
	return s.items[len(s.items)-1], nil
}

// Len returns the length of the stack
func (s *Stack[T]) Len() int {
	return len(s.items)
}
