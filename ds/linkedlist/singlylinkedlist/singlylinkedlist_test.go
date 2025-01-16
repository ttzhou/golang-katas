package singlylinkedlist_test

import (
	"golang-katas/ds/linkedlist/singlylinkedlist"
	"golang-katas/internal/assert"
	"slices"
	"strings"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {

	t.Run("Size()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.Size()).Equals(0)

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(0, 0)
		assert.That(t, sll.Size()).Equals(6)
	})

	t.Run("Get()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		node, err := sll.Get(0)
		assert.That(t, node).IsNil()
		assert.That(t, err).IsNotNil()

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)

		node, err = sll.Get(0)
		assert.That(t, err).IsNil()
		assert.That(t, node).IsNotNil()
		assert.That(t, node.Val).Equals(1)

		node, err = sll.Get(-3)
		assert.That(t, node).IsNil()
		assert.That(t, err).IsNotNil()

		node, err = sll.Get(6)
		assert.That(t, node).IsNil()
		assert.That(t, err).IsNotNil()
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
		assert.That(t, node.Val).Equals(0)
		assert.That(t, err).IsNil()
		assert.That(t, sll.String()).Equals("0->1->3->3->5->8->7")

		node, err = sll.Insert(-1, 5)
		assert.That(t, err).IsNotNil()
		assert.That(t, node).IsNil()

		node, err = sll.Insert(11, -1)
		assert.That(t, err).IsNotNil()
		assert.That(t, node).IsNil()
	})

	t.Run("Remove()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()

		node, err := sll.Remove(-1)
		assert.That(t, node).IsNil()
		assert.That(t, err).IsNotNil()

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)

		node, err = sll.Remove(11)
		assert.That(t, node).IsNil()
		assert.That(t, err).IsNotNil()
		assert.That(t, sll.String()).Equals("1->3->5")

		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)
		assert.That(t, sll.String()).Equals("1->3->3->5->8->7")

		node, err = sll.Remove(4)
		assert.That(t, err).IsNil()
		assert.That(t, node).IsNotNil()
		assert.That(t, node.Val).Equals(8)
		assert.That(t, sll.String()).Equals("1->3->3->5->7")
	})

	t.Run("Head()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.Head()).IsNil()

		_, _ = sll.Insert(0, 5)
		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)
		assert.That(t, sll.Head()).IsNotNil()
		assert.That(t, sll.Head().Val).Equals(1)
	})

	t.Run("Tail()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.Tail()).IsNil()

		_, _ = sll.Insert(0, 5)
		assert.That(t, sll.Tail()).IsNotNil()
		assert.That(t, sll.Tail().Val).Equals(5)

		_, _ = sll.Insert(0, 1)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(1, 3)
		_, _ = sll.Insert(4, 7)
		_, _ = sll.Insert(4, 8)
		assert.That(t, sll.Tail()).IsNotNil()
		assert.That(t, sll.Tail().Val).Equals(7)
	})

	t.Run("PopHead()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.PopHead()).IsNil()
		sll.Prepend(5)
		sll.Prepend(5)
		sll.Prepend(1)
		assert.That(t, sll.PopHead().Val).Equals(1)
		assert.That(t, sll.Head().Val).Equals(5)
		assert.That(t, sll.String()).Equals("5->5")
	})

	t.Run("PopTail()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.PopTail()).IsNil()
		sll.Append(5)
		sll.Append(5)
		sll.Append(1)
		assert.That(t, sll.PopTail().Val).Equals(1)
		assert.That(t, sll.Tail().Val).Equals(5)
		assert.That(t, sll.String()).Equals("5->5")
	})

	t.Run("MoveToHead()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		sll.Prepend(5)
		sll.Prepend(5)
		sll.Prepend(1)
		assert.That(t, sll.String()).Equals("1->5->5")
		sll.MoveToHead(2)
		assert.That(t, sll.String()).Equals("5->1->5")
	})

	t.Run("MoveToTail()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		sll.Append(5)
		sll.Append(5)
		sll.Append(1)
		assert.That(t, sll.String()).Equals("5->5->1")
		sll.MoveToTail(0)
		assert.That(t, sll.String()).Equals("5->1->5")
	})

	t.Run("Prepend()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.Head()).IsNil()
		assert.That(t, sll.String()).Equals("")

		sll.Prepend(5)
		assert.That(t, sll.String()).Equals("5")
		sll.Prepend(5)
		assert.That(t, sll.String()).Equals("5->5")
		sll.Prepend(1)
		assert.That(t, sll.String()).Equals("1->5->5")
		assert.That(t, sll.Head().Val).Equals(1)
	})

	t.Run("Append()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.Tail()).IsNil()
		assert.That(t, sll.String()).Equals("")

		sll.Append(5)
		assert.That(t, sll.String()).Equals("5")
		sll.Append(5)
		assert.That(t, sll.String()).Equals("5->5")
		sll.Append(1)
		assert.That(t, sll.String()).Equals("5->5->1")
		assert.That(t, sll.Tail().Val).Equals(1)
	})

	t.Run("Reverse()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()
		assert.That(t, sll.Head()).IsNil()
		assert.That(t, sll.Size()).Equals(0)
		sll.Append(5)
		sll.Reverse()
		assert.That(t, sll.String()).Equals("5")
		sll.Append(5)
		sll.Append(1)
		assert.That(t, sll.String()).Equals("5->5->1")
		sll.Reverse()
		assert.That(t, sll.String()).Equals("1->5->5")
	})

	t.Run("Swap()", func(t *testing.T) {
		sll := singlylinkedlist.New[int]()

		err := sll.Swap(0, 4)
		assert.That(t, err).IsNotNil()

		sll.Append(5)
		sll.Append(5)

		err = sll.Swap(0, 7)
		assert.That(t, err).IsNotNil()

		err = sll.Swap(-2, 11)
		assert.That(t, err).IsNotNil()

		sll.Append(1)
		sll.Append(1)
		sll.Append(2)
		assert.That(t, sll.String()).Equals("5->5->1->1->2")
		assert.That(t, sll.Size()).Equals(5)

		_ = sll.Swap(1, 4)
		assert.That(t, sll.String()).Equals("5->2->1->1->5")
		assert.That(t, sll.Size()).Equals(5)

		_ = sll.Swap(4, 1)
		assert.That(t, sll.String()).Equals("5->5->1->1->2")

		_ = sll.Swap(1, 2)
		assert.That(t, sll.String()).Equals("5->1->5->1->2")
		assert.That(t, sll.Size()).Equals(5)

		_ = sll.Swap(3, 4)
		assert.That(t, sll.String()).Equals("5->1->5->2->1")

		err = sll.Swap(4, 4)
		assert.That(t, err).IsNil()
		assert.That(t, sll.String()).Equals("5->1->5->2->1")
	})

	t.Run("Copy()", func(t *testing.T) {
		sll := singlylinkedlist.New[string]()
		sll.Append("1")
		sll.Append("1")
		sll.Append("2")

		cpy := sll.Copy()
		assert.That(t, sll.String()).Equals(cpy.String())
	})

	t.Run("IterVals()", func(t *testing.T) {
		sll := singlylinkedlist.New[string]()
		sll.Append("1")
		sll.Append("1")
		sll.Append("2")

		sllSlice := slices.Collect(sll.IterVals())
		assert.That(t, strings.Join(sllSlice, "->")).Equals("1->1->2")
	})

	t.Run("AppendList()", func(t *testing.T) {
		sll1 := singlylinkedlist.New[int]()
		sll2 := singlylinkedlist.New[int]()

		sll1.AppendList(sll2)
		assert.That(t, sll1.String()).Equals("")
		assert.That(t, sll1.Size()).Equals(0)

		sll1 = singlylinkedlist.New[int]()
		sll1.Append(5)
		sll1.Append(5)
		sll1.Append(1)
		assert.That(t, sll1.String()).Equals("5->5->1")
		assert.That(t, sll1.Size()).Equals(3)

		sll1.AppendList(sll2)
		assert.That(t, sll1.String()).Equals("5->5->1")
		assert.That(t, sll1.Size()).Equals(3)

		sll2 = singlylinkedlist.New[int]()
		sll2.Append(1)
		sll2.Append(1)
		sll2.Append(5)
		assert.That(t, sll2.String()).Equals("1->1->5")

		sll1.AppendList(sll2)
		assert.That(t, sll1.Size()).Equals(6)
		assert.That(t, sll1.String()).Equals("5->5->1->1->1->5")
	})

	t.Run("PrependList()", func(t *testing.T) {
		sll1 := singlylinkedlist.New[int]()
		sll2 := singlylinkedlist.New[int]()

		sll1.PrependList(sll2)
		assert.That(t, sll1.String()).Equals("")
		assert.That(t, sll1.Size()).Equals(0)

		sll1 = singlylinkedlist.New[int]()
		sll1.Prepend(5)
		sll1.Prepend(5)
		sll1.Prepend(1)
		assert.That(t, sll1.String()).Equals("1->5->5")
		assert.That(t, sll1.Size()).Equals(3)

		sll1.PrependList(sll2)
		assert.That(t, sll1.String()).Equals("1->5->5")
		assert.That(t, sll1.Size()).Equals(3)

		sll2 = singlylinkedlist.New[int]()
		sll2.Append(1)
		sll2.Append(1)
		sll2.Append(5)
		assert.That(t, sll2.String()).Equals("1->1->5")

		sll1.PrependList(sll2)
		assert.That(t, sll1.Size()).Equals(6)
		assert.That(t, sll1.String()).Equals("1->1->5->1->5->5")
	})
}
