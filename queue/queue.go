// Implementation of the Queue data structure.
// Queue is a linear data structure that follows FIFO (First-In-First-Out) principle.
package queue

import "errors"

var ErrEmpty = errors.New("queue is empty")

type Queue[T comparable] struct {
	items []T
}

// NewQueue creates and returns a new Queue. You can pass initial value to it.
func NewQueue[T comparable](i ...T) *Queue[T] {
	l := []T{}
	l = append(l, i...)
	return &Queue[T]{
		items: l,
	}
}

// Enqueue adds an element to the queue
func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) == 0 {
		return *new(T), ErrEmpty
	}
	dv := q.items[0]
	q.items = q.items[1:]
	return dv, nil
}

// Peek returns the front element of the queue without removing it
func (q *Queue[T]) Peek() (T, error) {
	if len(q.items) == 0 {
		return *new(T), ErrEmpty
	}
	return q.items[0], nil
}

// Len returns length of the queue
func (q *Queue[T]) Len() int {
	return len(q.items)
}
