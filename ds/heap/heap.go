package heap

import (
	"cmp"
	"fmt"
)

type heapType int

// Enum to represent MIN or MAX heap
// MIN represents a min heap.
// MAX represents a max heap.
const (
	MIN heapType = iota
	MAX
)

// A Heap is a binary heap implementation.
// Type represents the type of heap (MIN or MAX).
// Values are the underlying values in the heap.
type Heap[T cmp.Ordered] struct {
	Type   heapType
	Values []T
}

func (h Heap[T]) getParent(k int) int {
	return max((k-1)/2, 0)
}

func (h Heap[T]) getLeftChild(k int) int {
	index := 2*k + 1
	if index >= len(h.Values) {
		return -1
	}
	return index
}

func (h Heap[T]) getRightChild(k int) int {
	index := 2*k + 2
	if index >= len(h.Values) {
		return -1
	}
	return index
}

func heapCompare[T cmp.Ordered](ht heapType, a, b T) bool {
	if ht == MIN {
		return a < b
	}
	return a > b
}

func (h *Heap[T]) heapify(cur int) {
	lci := h.getLeftChild(cur)
	rci := h.getRightChild(cur)
	child := cur

	if lci > -1 && heapCompare(h.Type, h.Values[lci], h.Values[child]) {
		child = lci
	}
	if rci > -1 && heapCompare(h.Type, h.Values[rci], h.Values[child]) {
		child = rci
	}
	if child != cur {
		h.Values[child], h.Values[cur] = h.Values[cur], h.Values[child]
		h.heapify(child)
	}
}

// Heapify modifies the underlying values slice in the Heap
// to conform to the heap property.
func (h *Heap[T]) Heapify() {
	for i := len(h.Values)/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
}

// Push adds a value val to the heap.
func (h *Heap[T]) Push(val T) {
	h.Values = append(h.Values, val)
	cur := len(h.Values) - 1
	for cur > 0 {
		parent := h.getParent(cur)
		if heapCompare(h.Type, h.Values[parent], h.Values[cur]) {
			break
		}
		h.Values[parent], h.Values[cur] = h.Values[cur], h.Values[parent]
		cur = parent
	}
}

// Pop removes the topmost (min or max) value from the heap.
func (h *Heap[T]) Pop() (T, error) {
	var val T
	if h.Empty() {
		return val, fmt.Errorf("heap is empty")
	}
	val, h.Values = h.Values[0], h.Values[1:]
	h.Heapify()
	return val, nil
}

// String returns the string representation
// of the underlying slice of values.
func (h Heap[T]) String() string {
	return fmt.Sprintf("%v", h.Values)
}

// Size returns the number of values in the heap.
func (h Heap[T]) Size() int {
	return len(h.Values)
}

// Empty returns a boolean indicating if the heap has no values.
func (h Heap[T]) Empty() bool {
	return h.Size() == 0
}

func New[T cmp.Ordered](heapType heapType) *Heap[T] {
	return &Heap[T]{Type: heapType}
}
