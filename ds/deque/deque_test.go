package deque_test

import (
	"golang-katas/ds/deque"
	"golang-katas/internal/assert"
	"testing"
)

func TestDeque(t *testing.T) {

	assert := assert.New(t)

	t.Run("Size()", func(t *testing.T) {
		dq := deque.New[int]()
		assert.That(dq.Size()).Equals(0)
		assert.That(dq.IsEmpty()).Equals(true)
		assert.That(dq.String()).Equals("")

		dq.PushEnd(5)
		assert.That(dq.Size()).Equals(1)

		dq.PopEnd()
		assert.That(dq.Size()).Equals(0)
	})

	t.Run("PushStart()", func(t *testing.T) {
		dq := deque.New[int]()
		dq.PushStart(5)
		dq.PushStart(1)
		dq.PushStart(3)
		assert.That(dq.String()).Equals("3<->1<->5")
	})

	t.Run("PopStart()", func(t *testing.T) {
		dq := deque.New[int]()
		dq.PushStart(5)
		dq.PushStart(1)
		dq.PushStart(5)
		assert.That(dq.String()).Equals("5<->1<->5")

		dq.PopStart()
		assert.That(dq.String()).Equals("1<->5")
	})

	t.Run("PushEnd()", func(t *testing.T) {
		dq := deque.New[int]()
		dq.PushEnd(5)
		dq.PushEnd(1)
		dq.PushEnd(3)
		assert.That(dq.String()).Equals("5<->1<->3")
	})

	t.Run("PopEnd()", func(t *testing.T) {
		dq := deque.New[int]()
		dq.PushEnd(5)
		dq.PushEnd(1)
		dq.PushEnd(5)
		assert.That(dq.String()).Equals("5<->1<->5")

		dq.PopEnd()
		assert.That(dq.String()).Equals("5<->1")
	})

	t.Run("Flush()", func(t *testing.T) {
		dq := deque.New[int]()
		dq.PushEnd(5)
		dq.PushEnd(5)
		dq.PushEnd(1)

		assert.That(dq.String()).Equals("5<->5<->1")
		assert.That(dq.Size()).Equals(3)

		dq.Flush()
		assert.That(dq.IsEmpty()).Equals(true)
		assert.That(dq.String()).Equals("")
		assert.That(dq.Size()).Equals(0)
	})
}
