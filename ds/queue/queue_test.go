package queue_test

import (
	"golang-katas/ds/queue"
	"golang-katas/internal/assert"
	"testing"
)

func TestQueue(t *testing.T) {

	assert := assert.New(t)

	t.Run("Size()", func(t *testing.T) {
		q := queue.New[int]()
		assert.That(q.Size()).Equals(0)
		assert.That(q.IsEmpty()).Equals(true)
		assert.That(q.String()).Equals("")

		q.Push(5)
		assert.That(q.Size()).Equals(1)

		q.Pop()
		assert.That(q.Size()).Equals(0)
	})

	t.Run("Push()", func(t *testing.T) {
		q := queue.New[int]()
		q.Push(5)
		q.Push(5)
		q.Push(1)

		assert.That(q.String()).Equals("5<->5<->1")
		assert.That(q.Size()).Equals(3)
	})

	t.Run("Pop()", func(t *testing.T) {
		q := queue.New[int]()
		q.Push(5)
		q.Push(5)
		q.Push(1)

		assert.That(q.String()).Equals("5<->5<->1")

		val, ok := q.Pop()
		assert.That(ok).Equals(true)
		assert.That(val).Equals(5)
		assert.That(q.String()).Equals("5<->1")
	})

	t.Run("Flush()", func(t *testing.T) {
		q := queue.New[int]()
		q.Push(5)
		q.Push(5)
		q.Push(1)

		assert.That(q.String()).Equals("5<->5<->1")
		assert.That(q.Size()).Equals(3)

		q.Flush()
		assert.That(q.IsEmpty()).Equals(true)
		assert.That(q.String()).Equals("")
		assert.That(q.Size()).Equals(0)
	})
}
