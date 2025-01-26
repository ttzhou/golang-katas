package queue

import (
	doubly "golang-katas/ds/linkedlist/doublylinkedlist"
)

// Queue represents a queue data structure.
type Queue[T any] struct {
	list *doubly.LinkedList[T]
}

// Size returns the number of elements in the queue.
func (q Queue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty indicates whether the queue has no elements.
func (q Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Pop removes the element from the start (head) of the queue, and returns it,
// plus a boolean indicating whether the queue has elements to pop.
// If there are no elements in the queue, the value is
// a (meaningless) zero value, and the boolean will be
// value false.
func (q *Queue[T]) Pop() (top T, ok bool) {
	if !q.IsEmpty() {
		top = q.list.PopHead().Val
		ok = true
	}
	return top, ok
}

// Push adds a value to the end (tail) of the queue.
func (q *Queue[T]) Push(val T) {
	q.list.Append(val)
}

// Flush clears all elements from the queue.
func (q *Queue[T]) Flush() {
	for !q.IsEmpty() {
		q.list.PopHead()
	}
}

// String returns a string representation of the
// queue; the leftmost element is the top of the queue.
func (q Queue[T]) String() string {
	return q.list.String()
}

// New returns a pointer to a queue of type T.
func New[T any]() *Queue[T] {
	return &Queue[T]{list: doubly.New[T]()}
}
