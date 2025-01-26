package stack_test

import (
	"golang-katas/ds/stack"
	"golang-katas/internal/assert"
	"testing"
)

func TestStack(t *testing.T) {

	assert := assert.New(t)

	t.Run("Size()", func(t *testing.T) {
		s := stack.New[int]()
		assert.That(s.Size()).Equals(0)
		assert.That(s.IsEmpty()).Equals(true)
		assert.That(s.String()).Equals("")

		s.Push(5)
		assert.That(s.Size()).Equals(1)

		s.Pop()
		assert.That(s.Size()).Equals(0)
	})

	t.Run("Peek()", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(5)
		s.Push(5)
		s.Push(1)

		top, ok := s.Peek()
		assert.That(ok).Equals(true)
		assert.That(top).Equals(1)
	})

	t.Run("Push()", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(5)
		s.Push(5)
		s.Push(1)

		assert.That(s.String()).Equals("1->5->5")
		assert.That(s.Size()).Equals(3)
	})

	t.Run("Pop()", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(5)
		s.Push(5)
		s.Push(1)

		assert.That(s.String()).Equals("1->5->5")

		val, ok := s.Pop()
		assert.That(ok).Equals(true)
		assert.That(val).Equals(1)
		assert.That(s.String()).Equals("5->5")
	})

	t.Run("Flush()", func(t *testing.T) {
		s := stack.New[int]()
		s.Push(5)
		s.Push(5)
		s.Push(1)

		assert.That(s.String()).Equals("1->5->5")
		assert.That(s.Size()).Equals(3)

		s.Flush()
		assert.That(s.IsEmpty()).Equals(true)
		assert.That(s.String()).Equals("")
		assert.That(s.Size()).Equals(0)
	})
}
