// Linked list data structrue
package list

// Node defines structure of the nodes of list
type node[T comparable] struct {
	Value T
	Next  *node[T]
}

// List defines Linked list strcuture.
type List[T comparable] struct {
	Head *node[T]
	len  int
}

// This function will return new linked list pointer with the given type.
// You can pass initial value as argument.
func NewList[T comparable](initial ...T) *List[T] {
	l := &List[T]{}
	for i := 0; i < len(initial); i++ {
		l.Add(initial[i])
	}
	return l
}

// This function will add value v to list. It will also return list to make chain call possible.
func (l *List[T]) Add(v T) *List[T] {
	n := &node[T]{Value: v}
	if l.Head == nil {
		l.Head = n
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n
	}
	l.len++
	return l
}

// This function will return current length of the list.
func (l *List[T]) Len() int {
	return l.len
}

// This function will remove the given index of the list.
// It won't do anything if the index is less than zero or length of the list.
func (l *List[T]) RemoveAt(index int) *List[T] {
	if index < 0 {
		return l
	}
	if index > l.Len()-1 {
		return l
	}
	l.len--
	if index == 0 {
		l.Head = l.Head.Next
		return l
	}
	current := l.Head
	for i := 0; current.Next != nil && i < index-1; i++ {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	return l
}
