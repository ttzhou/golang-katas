package singlylinkedlist_test

import (
	"golang-katas/ds/linkedlist/singlylinkedlist"
	"golang-katas/internal/assert"
	"slices"
	"strings"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {

	assert := assert.New(t)

	t.Run("Size()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.Size()).Equals(0)

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(0, 0)
		assert.That(sll.Size()).Equals(6)
	})

	t.Run("Get()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		node, err := sll.Get(0)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)

		node, err = sll.Get(0)
		assert.That(err).IsNil()
		assert.That(node).IsNotNil()
		assert.That(node.Val).Equals(1)

		node, err = sll.Get(-3)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		node, err = sll.Get(6)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()
	})

	t.Run("Insert()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)

		node, err := sll.Insert(0, 0)
		assert.That(node.Val).Equals(0)
		assert.That(err).IsNil()
		assert.That(sll.String()).Equals("0->1->3->3->5->8->7")

		node, err = sll.Insert(-1, 5)
		assert.That(err).IsNotNil()
		assert.That(node).IsNil()

		node, err = sll.Insert(11, -1)
		assert.That(err).IsNotNil()
		assert.That(node).IsNil()
	})

	t.Run("Remove()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()

		node, err := sll.Remove(-1)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)

		node, err = sll.Remove(11)
		assert.That(node).IsNil()
		assert.That(err).IsNotNil()
		assert.That(sll.String()).Equals("1->3->5")

		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)
		assert.That(sll.String()).Equals("1->3->3->5->8->7")

		node, err = sll.Remove(4)
		assert.That(err).IsNil()
		assert.That(node).IsNotNil()
		assert.That(node.Val).Equals(8)
		assert.That(sll.String()).Equals("1->3->3->5->7")
	})

	t.Run("Head()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.Head()).IsNil()

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)
		assert.That(sll.Head()).IsNotNil()
		assert.That(sll.Head().Val).Equals(1)
	})

	t.Run("Tail()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.Tail()).IsNil()

		_, _ = sll.Insert(0, 5)
		assert.That(sll.Tail()).IsNotNil()
		assert.That(sll.Tail().Val).Equals(5)

		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)
		assert.That(sll.Tail()).IsNotNil()
		assert.That(sll.Tail().Val).Equals(7)
	})

	t.Run("PopHead()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.PopHead()).IsNil()
		sll.Prepend(5)
		sll.Prepend(5)
		sll.Prepend(1)
		assert.That(sll.PopHead().Val).Equals(1)
		assert.That(sll.Head().Val).Equals(5)
		assert.That(sll.String()).Equals("5->5")
	})

	t.Run("PopTail()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.PopTail()).IsNil()
		sll.Append(5)
		sll.Append(5)
		sll.Append(1)
		assert.That(sll.PopTail().Val).Equals(1)
		assert.That(sll.Tail().Val).Equals(5)
		assert.That(sll.String()).Equals("5->5")
	})

	t.Run("MoveToHead()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		sll.Prepend(5)
		sll.Prepend(5)
		sll.Prepend(1)
		assert.That(sll.String()).Equals("1->5->5")
		sll.MoveToHead(2)
		assert.That(sll.String()).Equals("5->1->5")
	})

	t.Run("MoveToTail()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		sll.Append(5)
		sll.Append(5)
		sll.Append(1)
		assert.That(sll.String()).Equals("5->5->1")
		sll.MoveToTail(0)
		assert.That(sll.String()).Equals("5->1->5")
	})

	t.Run("Prepend()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.Head()).IsNil()
		assert.That(sll.String()).Equals("")

		sll.Prepend(5)
		assert.That(sll.String()).Equals("5")
		sll.Prepend(5)
		assert.That(sll.String()).Equals("5->5")
		sll.Prepend(1)
		assert.That(sll.String()).Equals("1->5->5")
		assert.That(sll.Head().Val).Equals(1)
	})

	t.Run("Append()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.Tail()).IsNil()
		assert.That(sll.String()).Equals("")

		sll.Append(5)
		assert.That(sll.String()).Equals("5")
		sll.Append(5)
		assert.That(sll.String()).Equals("5->5")
		sll.Append(1)
		assert.That(sll.String()).Equals("5->5->1")
		assert.That(sll.Tail().Val).Equals(1)
	})

	t.Run("Reverse()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(sll.Head()).IsNil()
		assert.That(sll.Size()).Equals(0)
		sll.Append(5)
		sll.Reverse()
		assert.That(sll.String()).Equals("5")
		sll.Append(5)
		sll.Append(1)
		assert.That(sll.String()).Equals("5->5->1")
		sll.Reverse()
		assert.That(sll.String()).Equals("1->5->5")
	})

	t.Run("Swap()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()

		err := sll.Swap(0, 4)
		assert.That(err).IsNotNil()

		sll.Append(5)
		sll.Append(5)

		err = sll.Swap(0, 7)
		assert.That(err).IsNotNil()

		err = sll.Swap(-2, 11)
		assert.That(err).IsNotNil()

		sll.Append(1)
		sll.Append(1)
		sll.Append(2)
		assert.That(sll.String()).Equals("5->5->1->1->2")
		assert.That(sll.Size()).Equals(5)

		_ = sll.Swap(1, 4)
		assert.That(sll.String()).Equals("5->2->1->1->5")
		assert.That(sll.Size()).Equals(5)

		_ = sll.Swap(4, 1)
		assert.That(sll.String()).Equals("5->5->1->1->2")

		_ = sll.Swap(1, 2)
		assert.That(sll.String()).Equals("5->1->5->1->2")
		assert.That(sll.Size()).Equals(5)

		_ = sll.Swap(3, 4)
		assert.That(sll.String()).Equals("5->1->5->2->1")

		err = sll.Swap(4, 4)
		assert.That(err).IsNil()
		assert.That(sll.String()).Equals("5->1->5->2->1")
	})

	t.Run("Copy()", func(t *testing.T) {
		sll := singlylinkedlist.New[string]()
		sll.Append("1")
		sll.Append("1")
		sll.Append("2")

		cpy := sll.Copy()
		assert.That(sll.String()).Equals(cpy.String())
	})

	t.Run("IterVals()", func(t *testing.T) {
		sll := singlylinkedlist.New[string]()
		sll.Append("1")
		sll.Append("1")
		sll.Append("2")

		for range sll.IterVals() {
			break
		}

		sllSlice := slices.Collect(sll.IterVals())
		assert.That(strings.Join(sllSlice, "->")).Equals("1->1->2")
	})

	t.Run("AppendList()", func(t *testing.T) {
		sll1 := singlylinkedlist.New[int]()
		sll2 := singlylinkedlist.New[int]()

		sll1.AppendList(sll2)
		assert.That(sll1.String()).Equals("")
		assert.That(sll1.Size()).Equals(0)

		sll1 = singlylinkedlist.New[int]()
		sll1.Append(5)
		sll1.Append(5)
		sll1.Append(1)
		assert.That(sll1.String()).Equals("5->5->1")
		assert.That(sll1.Size()).Equals(3)

		sll1.AppendList(sll2)
		assert.That(sll1.String()).Equals("5->5->1")
		assert.That(sll1.Size()).Equals(3)

		sll2 = singlylinkedlist.New[int]()
		sll2.Append(1)
		sll2.Append(1)
		sll2.Append(5)
		assert.That(sll2.String()).Equals("1->1->5")

		sll1.AppendList(sll2)
		assert.That(sll1.Size()).Equals(6)
		assert.That(sll1.String()).Equals("5->5->1->1->1->5")
	})

	t.Run("PrependList()", func(t *testing.T) {
		sll1 := singlylinkedlist.New[int]()
		sll2 := singlylinkedlist.New[int]()

		sll1.PrependList(sll2)
		assert.That(sll1.String()).Equals("")
		assert.That(sll1.Size()).Equals(0)

		sll1 = singlylinkedlist.New[int]()
		sll1.Prepend(5)
		sll1.Prepend(5)
		sll1.Prepend(1)
		assert.That(sll1.String()).Equals("1->5->5")
		assert.That(sll1.Size()).Equals(3)

		sll1.PrependList(sll2)
		assert.That(sll1.String()).Equals("1->5->5")
		assert.That(sll1.Size()).Equals(3)

		sll2 = singlylinkedlist.New[int]()
		sll2.Append(1)
		sll2.Append(1)
		sll2.Append(5)
		assert.That(sll2.String()).Equals("1->1->5")

		sll1.PrependList(sll2)
		assert.That(sll1.Size()).Equals(6)
		assert.That(sll1.String()).Equals("1->1->5->1->5->5")
	})
}
