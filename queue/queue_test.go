package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue(1, 2, 3)
	if q.Len() != 3 {
		t.Errorf("expected %d but got %d", 3, q.Len())
	}
	q.Enqueue(4)
	if q.Len() != 4 {
		t.Errorf("expected %d but got %d", 4, q.Len())
	}
}
func TestDequeue(t *testing.T) {
	q := NewQueue(1, 2, 3)
	d1, err := q.Dequeue()
	if err != nil {
		t.Errorf("didn't expect error but got %v\n", err)
	}
	if d1 != 1 {
		t.Errorf("expected %d but got %d", 1, d1)
	}
	if q.Len() != 2 {
		t.Errorf("expected %d for queue length but got %d", 2, q.Len())
	}
	q.Dequeue()
	q.Dequeue()
	_, err = q.Dequeue()
	if err != ErrEmpty {
		t.Errorf("expected ErrEmpty but got %v", err)
	}
}
func TestLength(t *testing.T) {
	q := NewQueue(1)
	if q.Len() != 1 {
		t.Errorf("expected %d but got %d", 1, q.Len())
	}
	q.Dequeue()
	if q.Len() != 0 {
		t.Errorf("expected %d but got %d", 0, q.Len())
	}
	q.Enqueue(2)
	if q.Len() != 1 {
		t.Errorf("expected %d but got %d", 1, q.Len())
	}
}
func TestPeek(t *testing.T) {
	q := NewQueue(1, 3)
	p1, err := q.Peek()
	if err != nil {
		t.Errorf("didn't expect %v", err)
	}
	if p1 != 1 {
		t.Errorf("expected %d but got %d", 1, p1)
	}
	q.Dequeue()
	q.Dequeue()
	_, err = q.Peek()
	if err != ErrEmpty {
		t.Errorf("expected ErrEmpty but got %v", err)
	}

}
