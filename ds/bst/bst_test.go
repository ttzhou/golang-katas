package bst_test

import (
	"fmt"
	"golang-katas/ds/bst"
	"golang-katas/internal/assert"
	"slices"
	"testing"
)

func TestBST(t *testing.T) {

	assert := assert.New(t)

	type null struct{}

	t.Run("Operations on a purely left tree", func(t *testing.T) {

		binst := bst.New[int, any]()

		err := binst.Delete(5)
		assert.That(err).IsNotNil()
		assert.That(binst.Size).Equals(0)
		assert.That(binst.Height()).Equals(0)
		assert.That(binst.Contains(5)).IsFalse()

		_ = binst.Insert(5, null{})
		_ = binst.Insert(4, null{})
		_ = binst.Insert(3, null{})
		_ = binst.Insert(2, null{})
		_ = binst.Insert(1, null{})

		assert.That(binst.Size).Equals(5)
		assert.That(binst.Height()).Equals(5)
		assert.That(binst.Contains(3)).IsTrue()

		// Testing that these iterators break without error
		for range binst.Traverse(bst.PRE_ORDER) {
			break
		}
		for range binst.Traverse(bst.IN_ORDER) {
			break
		}
		for range binst.Traverse(bst.POST_ORDER) {
			break
		}
		for range binst.Traverse(bst.LEVEL_ORDER) {
			break
		}
		for range binst.TraverseLevels() {
			break
		}

		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.IN_ORDER)))).Equals("[1 2 3 4 5]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.PRE_ORDER)))).Equals("[5 4 3 2 1]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.POST_ORDER)))).Equals("[1 2 3 4 5]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.LEVEL_ORDER)))).Equals("[5 4 3 2 1]")

		err = binst.Delete(0)
		assert.That(err).IsNotNil()
		assert.That(binst.Size).Equals(5)
		assert.That(binst.Height()).Equals(5)
		assert.That(binst.Contains(0)).IsFalse()

		err = binst.Delete(4)
		assert.That(err).IsNil()
		assert.That(binst.Size).Equals(4)
		assert.That(binst.Height()).Equals(4)
		assert.That(binst.Contains(4)).IsFalse()
		assert.That(binst.Contains(5)).IsTrue()
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.IN_ORDER)))).Equals("[1 2 3 5]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.TraverseLevels()))).Equals("[[5] [3 _] [2 _ _ _] [1 _ _ _ _ _ _ _]]")
	})

	t.Run("Operations on a purely right tree", func(t *testing.T) {

		binst := bst.New[int, any]()

		err := binst.Delete(5)
		assert.That(err).IsNotNil()
		assert.That(binst.Size).Equals(0)
		assert.That(binst.Height()).Equals(0)
		assert.That(binst.Contains(5)).IsFalse()

		_ = binst.Insert(1, null{})
		_ = binst.Insert(2, null{})
		_ = binst.Insert(3, null{})
		_ = binst.Insert(4, null{})
		_ = binst.Insert(5, null{})

		assert.That(binst.Size).Equals(5)
		assert.That(binst.Height()).Equals(5)
		assert.That(binst.Contains(3)).IsTrue()

		// Testing that these iterators break without error
		for range binst.Traverse(bst.PRE_ORDER) {
			break
		}
		for range binst.Traverse(bst.IN_ORDER) {
			break
		}
		for range binst.Traverse(bst.POST_ORDER) {
			break
		}
		for range binst.Traverse(bst.LEVEL_ORDER) {
			break
		}
		for range binst.TraverseLevels() {
			break
		}

		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.IN_ORDER)))).Equals("[1 2 3 4 5]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.PRE_ORDER)))).Equals("[1 2 3 4 5]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.POST_ORDER)))).Equals("[5 4 3 2 1]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.LEVEL_ORDER)))).Equals("[1 2 3 4 5]")

		err = binst.Delete(0)
		assert.That(err).IsNotNil()
		assert.That(binst.Size).Equals(5)
		assert.That(binst.Height()).Equals(5)
		assert.That(binst.Contains(0)).IsFalse()

		err = binst.Delete(4)
		assert.That(err).IsNil()
		assert.That(binst.Size).Equals(4)
		assert.That(binst.Height()).Equals(4)
		assert.That(binst.Contains(4)).IsFalse()
		assert.That(binst.Contains(5)).IsTrue()
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.IN_ORDER)))).Equals("[1 2 3 5]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.TraverseLevels()))).Equals("[[1] [_ 2] [_ _ _ 3] [_ _ _ _ _ _ _ 5]]")
	})

	t.Run("Operations on a two-sided tree", func(t *testing.T) {

		binst := bst.New[int, any]()
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.LEVEL_ORDER)))).Equals("[]")

		// Testing that these iterators break without error
		for range binst.Traverse(bst.PRE_ORDER) {
			break
		}
		for range binst.Traverse(bst.IN_ORDER) {
			break
		}
		for range binst.Traverse(bst.POST_ORDER) {
			break
		}
		for range binst.Traverse(bst.LEVEL_ORDER) {
			break
		}
		for range binst.TraverseLevels() {
			break
		}

		_ = binst.Insert(8, null{})
		_ = binst.Insert(4, null{})
		_ = binst.Insert(12, null{})
		_ = binst.Insert(2, null{})
		_ = binst.Insert(6, null{})
		_ = binst.Insert(10, null{})
		_ = binst.Insert(14, null{})
		_ = binst.Insert(1, null{})
		_ = binst.Insert(3, null{})
		_ = binst.Insert(5, null{})
		_ = binst.Insert(7, null{})
		_ = binst.Insert(9, null{})
		_ = binst.Insert(11, null{})
		_ = binst.Insert(13, null{})
		_ = binst.Insert(15, null{})

		assert.That(binst.Size).Equals(15)
		assert.That(binst.Height()).Equals(4)
		assert.That(binst.Contains(11)).IsTrue()

		assert.That(fmt.Sprintf("%v", slices.Collect(binst.Traverse(bst.IN_ORDER)))).Equals("[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]")
		assert.That(fmt.Sprintf("%v", slices.Collect(binst.TraverseLevels()))).Equals("[[8] [4 12] [2 6 10 14] [1 3 5 7 9 11 13 15]]")

		err := binst.Delete(16)
		assert.That(err).IsNotNil()

		err = binst.Delete(8)
		assert.That(err).IsNil()

		assert.That(fmt.Sprintf("%v", slices.Collect(binst.TraverseLevels()))).Equals("[[9] [4 12] [2 6 10 14] [1 3 5 7 _ 11 13 15]]")

		// 15 already present
		err = binst.Insert(15, null{})
		assert.That(err).IsNotNil()

		_, err = binst.Get(10)
		assert.That(err).IsNil()

		_, err = binst.Get(100)
		assert.That(err).IsNotNil()
	})
}
