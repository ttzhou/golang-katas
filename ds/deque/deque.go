package deque

import (
	doubly "golang-katas/ds/linkedlist/doublylinkedlist"
)

// Deque represents a Deque data structure.
type Deque[T any] struct {
	list *doubly.LinkedList[T]
}

// Size returns the number of elements in the deque.
func (dq Deque[T]) Size() int {
	return dq.list.Size()
}

// IsEmpty indicates whether the deque has no elements.
func (dq Deque[T]) IsEmpty() bool {
	return dq.Size() == 0
}

// PopStart removes the element from the start (head) of the Deque, and returns it,
// plus a boolean indicating whether the Deque has elements to pop.
// If there are no elements in the Deque, the value is
// a (meaningless) zero value, and the boolean will be
// value false.
func (dq *Deque[T]) PopStart() (top T, ok bool) {
	if !dq.IsEmpty() {
		top = dq.list.PopHead().Val
		ok = true
	}
	return top, ok
}

// Push adds a value to the end (tail) of the Deque.
func (dq *Deque[T]) PushEnd(val T) {
	dq.list.Append(val)
}

// PopEnd removes the element from the end (tail) of the Deque, and returns it,
// plus a boolean indicating whether the Deque has elements to pop.
// If there are no elements in the Deque, the value is
// a (meaningless) zero value, and the boolean will be
// value false.
func (dq *Deque[T]) PopEnd() (top T, ok bool) {
	if !dq.IsEmpty() {
		top = dq.list.PopTail().Val
		ok = true
	}
	return top, ok
}

// PushStart adds a value to the start (head) of the Deque.
func (dq *Deque[T]) PushStart(val T) {
	dq.list.Prepend(val)
}

// Flush clears all elements from the Deque.
func (dq *Deque[T]) Flush() {
	for !dq.IsEmpty() {
		dq.list.PopHead()
	}
}

// String returns a string representation of the
// deque; the leftmost element is the top of the deque.
func (dq Deque[T]) String() string {
	return dq.list.String()
}

// New returns a pointer to a deque of type T.
func New[T any]() *Deque[T] {
	return &Deque[T]{list: doubly.New[T]()}
}
