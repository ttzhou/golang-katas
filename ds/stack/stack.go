package stack

import (
	singly "golang-katas/ds/linkedlist/singlylinkedlist"
)

// Stack represents a stack data structure.
type Stack[T any] struct {
	list *singly.LinkedList[T]
}

// Size returns the number of elements in the stack.
func (s Stack[T]) Size() int {
	return s.list.Size()
}

// IsEmpty indicates whether the stack has no elements.
func (s Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Peek returns the value of the stack's topmost element,
// and a boolean indicating whether the stack has elements.
// If there are no elements in the stack, the value is
// a (meaningless) zero value, and the boolean will be
// value false.
func (s Stack[T]) Peek() (top T, ok bool) {
	if !s.IsEmpty() {
		top = s.list.Head().Val
		ok = true
	}
	return top, ok
}

// Pop removes the element from the top of the stack, and returns it,
// plus a boolean indicating whether the queue has elements to pop.
// That value is also removed from the stack.
// If there are no elements in the stack, the value is
// a (meaningless) zero value, and the boolean will be
// value false.
func (s *Stack[T]) Pop() (top T, ok bool) {
	if !s.IsEmpty() {
		top = s.list.PopHead().Val
		ok = true
	}
	return top, ok
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(val T) {
	s.list.Prepend(val)
}

// Flush clears all elements from the stack.
func (s *Stack[T]) Flush() {
	for !s.IsEmpty() {
		s.list.PopHead()
	}
}

// String returns a string representation of the
// stack; the leftmost element is the top of the stack.
func (s Stack[T]) String() string {
	return s.list.String()
}

// New returns a pointer to a stack of type T.
func New[T any]() *Stack[T] {
	return &Stack[T]{list: singly.New[T]()}
}
