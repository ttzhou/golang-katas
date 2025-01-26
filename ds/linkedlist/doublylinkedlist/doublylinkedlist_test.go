package doublylinkedlist_test

import (
	"golang-katas/ds/linkedlist/doublylinkedlist"
	"golang-katas/internal/assert"
	"slices"
	"strings"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {

	assert := assert.New(t)

	t.Run("Size()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Size()).Equals(0)

		_, _ = dll.Insert(0, 5)
		_, _ = dll.Insert(0, 1)
		_, _ = dll.Insert(1, 3)
		_, _ = dll.Insert(1, 3)
		_, _ = dll.Insert(4, 7)
		_, _ = dll.Insert(0, 0)
		assert.That(dll.Size()).Equals(6)
	})

	t.Run("Head()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Head()).IsNil()

		_, _ = dll.Insert(0, 5)
		assert.That(dll.Head().Val).Equals(5)
		_, _ = dll.Insert(0, 1)
		assert.That(dll.Head().Val).Equals(1)
	})

	t.Run("Tail()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Tail()).IsNil()

		_, _ = dll.Insert(0, 5)
		assert.That(dll.Tail().Val).Equals(5)
		_, _ = dll.Insert(1, 1)
		assert.That(dll.Tail().Val).Equals(1)
	})

	t.Run("Get()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()

		node, err := dll.Get(0)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		node, err = dll.Get(-1)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		_, _ = dll.Insert(0, 1)
		_, _ = dll.Insert(0, 2)
		node, err = dll.Get(0)
		assert.That(err).IsNil()
		assert.That(node.Val).Equals(2)

		node, err = dll.Get(1)
		assert.That(err).IsNil()
		assert.That(node.Val).Equals(1)

		_, _ = dll.Insert(2, 3)
		node, err = dll.Get(2)
		assert.That(err).IsNil()
		assert.That(node.Val).Equals(3)
	})

	t.Run("Insert()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Tail()).IsNil()
		_, _ = dll.Insert(0, 5)
		_, _ = dll.Insert(0, 1)
		_, _ = dll.Insert(2, 3)
		assert.That(dll.String()).Equals("1<->5<->3")
		assert.That(dll.Head().Val).Equals(1)
		assert.That(dll.Tail().Val).Equals(3)

		_, _ = dll.Insert(0, 1)
		_, _ = dll.Insert(1, 3)
		_, _ = dll.Insert(1, 3)
		_, _ = dll.Insert(4, 7)
		_, _ = dll.Insert(4, 8)
		node, err := dll.Insert(0, 0)
		assert.That(err).IsNil()
		assert.That(node.Val).Equals(0)
		assert.That(dll.String()).Equals("0<->1<->3<->3<->1<->8<->7<->5<->3")

		node, err = dll.Insert(-1, 5)
		assert.That(err).IsNotNil()
		assert.That(node).IsNil()

		node, err = dll.Insert(11, -1)
		assert.That(err).IsNotNil()
		assert.That(node).IsNil()
	})

	t.Run("Remove()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		node, err := dll.Remove(0)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		_, _ = dll.Insert(0, 5)
		_, _ = dll.Insert(0, 1)
		_, _ = dll.Insert(2, 3)
		_, _ = dll.Insert(0, 1)
		assert.That(dll.String()).Equals("1<->1<->5<->3")

		node, err = dll.Remove(3)
		assert.That(err).IsNil()
		assert.That(node.Val).Equals(3)
		assert.That(dll.String()).Equals("1<->1<->5")

		node, err = dll.Remove(0)
		assert.That(err).IsNil()
		assert.That(node.Val).Equals(1)
		assert.That(dll.String()).Equals("1<->5")

		node, err = dll.Remove(-1)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		node, err = dll.Remove(11)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()
	})

	t.Run("PopHead()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.PopHead()).IsNil()
		dll.Prepend(5)
		dll.Prepend(5)
		dll.Prepend(1)
		assert.That(dll.PopHead().Val).Equals(1)
		assert.That(dll.Head().Val).Equals(5)
		assert.That(dll.String()).Equals("5<->5")
	})

	t.Run("PopTail()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.PopTail()).IsNil()
		dll.Append(5)
		dll.Append(5)
		dll.Append(1)
		assert.That(dll.PopTail().Val).Equals(1)
		assert.That(dll.Tail().Val).Equals(5)
		assert.That(dll.String()).Equals("5<->5")
	})

	t.Run("MoveToHead()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		dll.Prepend(5)
		dll.Prepend(5)
		dll.Prepend(1)
		assert.That(dll.String()).Equals("1<->5<->5")
		dll.MoveToHead(2)
		assert.That(dll.String()).Equals("5<->1<->5")
	})

	t.Run("MoveToTail()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		dll.Append(5)
		dll.Append(5)
		dll.Append(1)
		assert.That(dll.String()).Equals("5<->5<->1")
		dll.MoveToTail(0)
		assert.That(dll.String()).Equals("5<->1<->5")
	})

	t.Run("Prepend()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Head()).IsNil()
		assert.That(dll.String()).Equals("")

		dll.Prepend(5)
		assert.That(dll.String()).Equals("5")
		dll.Prepend(5)
		assert.That(dll.String()).Equals("5<->5")
		dll.Prepend(1)
		assert.That(dll.String()).Equals("1<->5<->5")
		assert.That(dll.Head().Val).Equals(1)
	})

	t.Run("Append()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Tail()).IsNil()
		assert.That(dll.String()).Equals("")

		dll.Append(5)
		assert.That(dll.String()).Equals("5")
		dll.Append(5)
		assert.That(dll.String()).Equals("5<->5")
		dll.Append(1)
		assert.That(dll.String()).Equals("5<->5<->1")
		assert.That(dll.Tail().Val).Equals(1)
	})

	t.Run("Reverse()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()
		assert.That(dll.Head()).IsNil()
		assert.That(dll.Size()).Equals(0)
		dll.Append(5)
		dll.Reverse()
		assert.That(dll.String()).Equals("5")
		dll.Append(5)
		dll.Append(1)
		assert.That(dll.String()).Equals("5<->5<->1")
		dll.Reverse()
		assert.That(dll.String()).Equals("1<->5<->5")
	})

	t.Run("Swap()", func(t *testing.T) {
		dll := doublylinkedlist.New[int]()

		err := dll.Swap(0, 4)
		assert.That(err).IsNotNil()

		dll.Append(5)
		dll.Append(5)

		err = dll.Swap(0, 7)
		assert.That(err).IsNotNil()

		err = dll.Swap(-2, 11)
		assert.That(err).IsNotNil()

		dll.Append(1)
		dll.Append(1)
		dll.Append(2)
		assert.That(dll.String()).Equals("5<->5<->1<->1<->2")
		assert.That(dll.Size()).Equals(5)

		_ = dll.Swap(1, 4)
		assert.That(dll.String()).Equals("5<->2<->1<->1<->5")
		assert.That(dll.Size()).Equals(5)

		_ = dll.Swap(4, 1)
		assert.That(dll.String()).Equals("5<->5<->1<->1<->2")

		_ = dll.Swap(1, 2)
		assert.That(dll.String()).Equals("5<->1<->5<->1<->2")
		assert.That(dll.Size()).Equals(5)

		_ = dll.Swap(3, 4)
		assert.That(dll.String()).Equals("5<->1<->5<->2<->1")

		err = dll.Swap(4, 4)
		assert.That(err).IsNil()
		assert.That(dll.String()).Equals("5<->1<->5<->2<->1")
	})

	t.Run("Copy()", func(t *testing.T) {
		dll := doublylinkedlist.New[string]()
		dll.Append("1")
		dll.Append("1")
		dll.Append("2")

		cpy := dll.Copy()
		assert.That(dll.String()).Equals(cpy.String())
	})

	t.Run("IterVals()", func(t *testing.T) {
		dll := doublylinkedlist.New[string]()
		dll.Append("1")
		dll.Append("1")
		dll.Append("2")

		for range dll.IterVals() {
			break
		}

		dllSlice := slices.Collect(dll.IterVals())
		assert.That(strings.Join(dllSlice, "<->")).Equals("1<->1<->2")
	})

	t.Run("IterValsReversed()", func(t *testing.T) {
		dll := doublylinkedlist.New[string]()
		dll.Append("1")
		dll.Append("1")
		dll.Append("2")

		for range dll.IterValsReversed() {
			break
		}

		dllSlice := slices.Collect(dll.IterValsReversed())
		assert.That(strings.Join(dllSlice, "<->")).Equals("2<->1<->1")
	})

	t.Run("AppendList()", func(t *testing.T) {
		dll1 := doublylinkedlist.New[int]()
		dll2 := doublylinkedlist.New[int]()

		dll1.AppendList(dll2)
		assert.That(dll1.String()).Equals("")
		assert.That(dll1.Size()).Equals(0)

		dll1 = doublylinkedlist.New[int]()
		dll1.Append(5)
		dll1.Append(5)
		dll1.Append(1)
		assert.That(dll1.String()).Equals("5<->5<->1")
		assert.That(dll1.Size()).Equals(3)

		dll1.AppendList(dll2)
		assert.That(dll1.String()).Equals("5<->5<->1")
		assert.That(dll1.Size()).Equals(3)

		dll2 = doublylinkedlist.New[int]()
		dll2.Append(1)
		dll2.Append(1)
		dll2.Append(5)
		assert.That(dll2.String()).Equals("1<->1<->5")

		dll1.AppendList(dll2)
		assert.That(dll1.Size()).Equals(6)
		assert.That(dll1.String()).Equals("5<->5<->1<->1<->1<->5")
	})

	t.Run("PrependList()", func(t *testing.T) {
		dll1 := doublylinkedlist.New[int]()
		dll2 := doublylinkedlist.New[int]()

		dll1.PrependList(dll2)
		assert.That(dll1.String()).Equals("")
		assert.That(dll1.Size()).Equals(0)

		dll1 = doublylinkedlist.New[int]()
		dll1.Prepend(5)
		dll1.Prepend(5)
		dll1.Prepend(1)
		assert.That(dll1.String()).Equals("1<->5<->5")
		assert.That(dll1.Size()).Equals(3)

		dll1.PrependList(dll2)
		assert.That(dll1.String()).Equals("1<->5<->5")
		assert.That(dll1.Size()).Equals(3)

		dll2 = doublylinkedlist.New[int]()
		dll2.Append(1)
		dll2.Append(1)
		dll2.Append(5)
		assert.That(dll2.String()).Equals("1<->1<->5")

		dll1.PrependList(dll2)
		assert.That(dll1.Size()).Equals(6)
		assert.That(dll1.String()).Equals("1<->1<->5<->1<->5<->5")
	})

}
