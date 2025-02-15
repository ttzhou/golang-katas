package heap_test

import (
	"golang-katas/ds/heap"
	"golang-katas/internal/assert"
	"slices"
	"testing"
)

func TestHeap(t *testing.T) {

	assert := assert.New(t)

	t.Run("Heapify", func(t *testing.T) {
		minHeap := heap.New[int](heap.MIN)
		minHeap.Values = []int{8, 5, 1}
		minHeap.Heapify()
		assert.That(minHeap.Values).Equals([]int{1, 5, 8})

		maxHeap := heap.New[int](heap.MAX)
		maxHeap.Values = []int{1, 5, 8}
		maxHeap.Heapify()
		assert.That(maxHeap.Values).Equals([]int{8, 5, 1})
	})

	t.Run("Min heap", func(t *testing.T) {
		minHeap := heap.New[int](heap.MIN)
		assert.That(minHeap.Empty()).IsTrue()
		assert.That(minHeap.Size()).Equals(0)
		assert.That(minHeap.String()).Equals("[]")

		_, err := minHeap.Pop()
		assert.That(err).IsNotNil()
		assert.That(minHeap.Size()).Equals(0)

		minHeap.Push(7)
		minHeap.Push(3)
		minHeap.Push(15)
		minHeap.Push(10)
		minHeap.Push(5)
		minHeap.Push(5)

		assert.That(minHeap.String()).Equals("[3 5 5 10 7 15]")
		assert.That(minHeap.Size()).Equals(6)

		val, err := minHeap.Pop()

		assert.That(err).IsNil()
		assert.That(val).Equals(3)
		assert.That(minHeap.String()).Equals("[5 5 10 7 15]")

		val, _ = minHeap.Pop()
		assert.That(val).Equals(5)
		assert.That(minHeap.String()).Equals("[5 10 7 15]")

		assert.That(slices.Min(minHeap.Values)).Equals(minHeap.Values[0])
	})

	t.Run("Max heap", func(t *testing.T) {
		maxHeap := heap.New[int](heap.MAX)
		assert.That(maxHeap.Size()).Equals(0)

		_, err := maxHeap.Pop()
		assert.That(err).IsNotNil()
		assert.That(maxHeap.Size()).Equals(0)

		maxHeap.Push(7)
		maxHeap.Push(3)
		maxHeap.Push(15)
		assert.That(maxHeap.String()).Equals("[15 3 7]")

		_, _ = maxHeap.Pop()

		assert.That(maxHeap.String()).Equals("[7 3]")

		maxHeap.Push(10)
		maxHeap.Push(15)
		maxHeap.Push(5)
		maxHeap.Push(6)
		maxHeap.Push(6)

		assert.That(maxHeap.String()).Equals("[15 10 7 3 5 6 6]")
		assert.That(maxHeap.Size()).Equals(7)

		val, err := maxHeap.Pop()

		assert.That(err).IsNil()
		assert.That(val).Equals(15)
		assert.That(maxHeap.String()).Equals("[10 7 6 5 6 3]")
		assert.That(maxHeap.Size()).Equals(6)

		val, _ = maxHeap.Pop()
		assert.That(val).Equals(10)
		assert.That(maxHeap.String()).Equals("[7 6 5 6 3]")
		assert.That(maxHeap.Size()).Equals(5)

		assert.That(slices.Max(maxHeap.Values)).Equals(maxHeap.Values[0])
	})

	t.Run("Heapsort (max)", func(t *testing.T) {
		maxHeap := heap.New[int](heap.MAX)
		unsortedVals := []int{8, 19, -2, -5, 0, 0, 2, 7, 7, -2}
		for _, val := range unsortedVals {
			maxHeap.Push(val)
		}
		descSortedVals := make([]int, len(maxHeap.Values))
		for i := range len(descSortedVals) {
			val, _ := maxHeap.Pop()
			descSortedVals[i] = val
		}
		assert.That(descSortedVals).Equals([]int{19, 8, 7, 7, 2, 0, 0, -2, -2, -5})
	})

	t.Run("Heapsort (min)", func(t *testing.T) {
		minHeap := heap.New[int](heap.MIN)
		unsortedVals := []int{8, 19, -2, -5, 0, 0, 2, 7, 7, -2}
		for _, val := range unsortedVals {
			minHeap.Push(val)
		}
		ascSortedVals := make([]int, len(minHeap.Values))
		for i := range len(ascSortedVals) {
			val, _ := minHeap.Pop()
			ascSortedVals[i] = val
		}
		assert.That(ascSortedVals).Equals([]int{-5, -2, -2, 0, 0, 2, 7, 7, 8, 19})
	})
}
