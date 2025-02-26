package doublylinkedlist

import (
	"fmt"
	"iter"
	"strings"
)

type Node[T any] struct {
	Val        T
	prev, next *Node[T]
}

// LinkedList represents a doubly linked list.
type LinkedList[T any] struct {
	headstub, tailstub Node[T]
	length             int
}

// String returns a human-readable representation of the linked list.
func (dll LinkedList[T]) String() string {
	nodes := make([]string, dll.length)
	cur := &dll.headstub
	for i := range dll.length {
		cur = cur.next
		nodes[i] = fmt.Sprintf("%v", cur.Val)
	}
	return strings.Join(nodes, "<->")
}

// Size returns the number of nodes in the linked list.
func (dll LinkedList[T]) Size() int {
	return dll.length
}

// Head returns the head node, or nil if it does not exist.
func (dll LinkedList[T]) Head() *Node[T] {
	if dll.length == 0 {
		return nil
	}
	return dll.headstub.next
}

// Tail returns the last node, or nil if it does not exist.
func (dll LinkedList[T]) Tail() *Node[T] {
	if dll.length == 0 {
		return nil
	}
	return dll.tailstub.prev
}

// Get returns the node at position `pos` in the doubly linked list. If `pos`
// is negative or greater than the length of the linked list at time of
// insertion, an error will be returned.
func (dll *LinkedList[T]) Get(pos int) (*Node[T], error) {
	if pos < 0 || pos >= dll.length {
		return nil, fmt.Errorf("invalid pos %d for length %d", pos, dll.length)
	}
	if pos > dll.length/2 {
		cur := &dll.tailstub
		for i := dll.length - 1; i >= pos; i-- {
			cur = cur.prev
		}
		return cur, nil
	} else {
		cur := &dll.headstub
		for i := 0; i <= pos; i++ {
			cur = cur.next
		}
		return cur, nil
	}
}

func (dll *LinkedList[T]) insertNode(pos int, node *Node[T]) (*Node[T], error) {
	if pos < 0 || pos > dll.length {
		return nil, fmt.Errorf("invalid pos %d for length %d", pos, dll.length)
	}
	if pos > dll.length/2 {
		cur := &dll.tailstub
		for i := dll.length - 1; i >= pos; i-- {
			cur = cur.prev
		}
		cur.prev, cur.prev.next, node.prev, node.next = node, node, cur.prev, cur
	} else {
		cur := &dll.headstub
		for range pos {
			cur = cur.next
		}
		cur.next, cur.next.prev, node.next, node.prev = node, node, cur.next, cur
	}
	dll.length++
	return node, nil
}

// Insert inserts a node with value `val` into the doubly linked list so that
// it's index will be `pos` after insertion. Returns a pointer to the inserted
// node. If `pos` is negative or greater than the length of the linked list at
// time of insertion, an error will be returned.
func (dll *LinkedList[T]) Insert(pos int, val T) (*Node[T], error) {
	node := &Node[T]{Val: val}
	return dll.insertNode(pos, node)
}

// Remove removes the node at position `pos` and returns a pointer to the
// removed node. If `pos` is negative or greater than the length of the linked
// list at time of insertion, an error will be returned.
func (dll *LinkedList[T]) Remove(pos int) (*Node[T], error) {
	if pos < 0 || pos >= dll.length {
		return nil, fmt.Errorf("invalid pos %d for length %d", pos, dll.length)
	}
	var removed *Node[T]
	if pos > dll.length/2 {
		cur := &dll.tailstub
		for i := dll.length - 1; i >= pos; i-- {
			cur = cur.prev
		}
		removed = cur
	} else {
		cur := &dll.headstub
		for i := 0; i <= pos; i++ {
			cur = cur.next
		}
		removed = cur
	}
	removed.prev.next, removed.next.prev = removed.next, removed.prev
	dll.length--
	return removed, nil
}

// MoveToHead moves the node at position `pos` to the head of the linked list,
// and returns the node.
func (dll *LinkedList[T]) MoveToHead(pos int) *Node[T] {
	node, _ := dll.Remove(pos)
	_, _ = dll.insertNode(0, node)
	return node
}

// MoveToTail moves the node at position `pos` to the tail of the linked list,
// and returns the node.
func (dll *LinkedList[T]) MoveToTail(pos int) *Node[T] {
	node, _ := dll.Remove(pos)
	_, _ = dll.insertNode(dll.length, node)
	return node
}

// PopHead removes the node at the head of the linked list,
// and returns the removed node. If there are no nodes
// to remove, nil is returned.
func (dll *LinkedList[T]) PopHead() *Node[T] {
	node, _ := dll.Remove(0)
	return node
}

// PopTail removes the node at the head of the linked list,
// and returns the removed node. If there are no nodes
// to remove, nil is returned.
func (dll *LinkedList[T]) PopTail() *Node[T] {
	node, _ := dll.Remove(max(0, dll.length-1))
	return node
}

// Prepend adds a node to the head of the linked list,
// and returns the node.
func (dll *LinkedList[T]) Prepend(val T) *Node[T] {
	node, _ := dll.Insert(0, val)
	return node
}

// Append adds a node to the tail of the linked list,
// and returns the node.
func (dll *LinkedList[T]) Append(val T) *Node[T] {
	node, _ := dll.Insert(dll.length, val)
	return node
}

// Reverse reverses the linked list's nodes.
func (dll *LinkedList[T]) Reverse() *LinkedList[T] {
	cur := dll.Head()
	for range dll.length {
		cur.prev, cur.next, cur = cur.next, cur.prev, cur.next
	}
	dll.headstub.next, dll.tailstub.prev = dll.tailstub.prev, dll.headstub.next
	return dll
}

// Swap swaps the linked list nodes at positions i and j.
func (dll *LinkedList[T]) Swap(i, j int) error {
	if dll.length < 2 {
		return fmt.Errorf("not enough elements to swap")
	}
	if i < 0 || i >= dll.length || j < 0 || j >= dll.length {
		return fmt.Errorf("invalid pos %d, %d for length %d", i, j, dll.length)
	}
	if i == j {
		return nil
	}
	i, j = min(i, j), max(i, j)
	var in, jn *Node[T]
	if j <= dll.length/2 {
		cur := dll.Head()
		for p := 0; p <= j; p++ {
			if p == i {
				in = cur
			}
			if p == j {
				jn = cur
			}
			cur = cur.next
		}
	} else if i >= dll.length/2 {
		cur := dll.Tail()
		for p := dll.length - 1; p >= i; p-- {
			if p == i {
				in = cur
			}
			if p == j {
				jn = cur
			}
			cur = cur.prev
		}
	} else {
		in, _ = dll.Get(i)
		jn, _ = dll.Get(j)
	}
	if in.next == jn {
		in.prev, in.next, jn.prev, jn.next = jn, jn.next, in.prev, in
		jn.prev.next, in.next.prev = jn, in
	} else {
		in.prev, in.next, jn.prev, jn.next = jn.prev, jn.next, in.prev, in.next
		jn.prev.next, jn.next.prev, in.prev.next, in.next.prev = jn, jn, in, in
	}
	return nil
}

// Copy makes a copy of the linked list's nodes and returns a
// pointer to it.
func (dll LinkedList[T]) Copy() *LinkedList[T] {
	cpy := New[T]()
	cur := &cpy.headstub
	for val := range dll.IterVals() {
		cur.next = &Node[T]{Val: val}
		cur = cur.next
		cpy.length++
	}
	cur.next = &cpy.tailstub
	return cpy
}

// IterNodes returns an iterator over the nodes in the linked list,
// starting from the head.
func (dll LinkedList[T]) iterNodes() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		cur := dll.Head()
		for range dll.length {
			if !yield(cur) {
				return
			}
			cur = cur.next
		}
	}
}

// iterNodesReversed returns an iterator over the nodes in the linked list,
// starting from the tail.
func (dll LinkedList[T]) iterNodesReversed() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		cur := dll.Tail()
		for range dll.length {
			if !yield(cur) {
				return
			}
			cur = cur.prev
		}
	}
}

// IterVals returns an iterator over the values in the linked list,
// starting from the head.
func (dll LinkedList[T]) IterVals() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := range dll.iterNodes() {
			if !yield(node.Val) {
				return
			}
		}
	}
}

// IterValsReversed returns an iterator over the values in the linked list,
// starting from the head.
func (dll LinkedList[T]) IterValsReversed() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := range dll.iterNodesReversed() {
			if !yield(node.Val) {
				return
			}
		}
	}
}

// AppendList takes another linked list, adds its nodes to the end
// of the receiver, and then returns the modified receiver.
// The other linked list is not modified.
func (this *LinkedList[T]) AppendList(other *LinkedList[T]) *LinkedList[T] {
	cur := this.Tail()
	if cur == nil {
		cur = &this.headstub
	}
	for val := range other.IterVals() {
		node := &Node[T]{Val: val}
		cur.next, node.prev = node, cur
		cur = cur.next
		this.length++
	}
	cur.next = &this.tailstub
	this.tailstub.prev = cur
	return this
}

// PrependList takes another linked list, adds its nodes to the end
// of the receiver, and then returns the modified receiver.
func (this *LinkedList[T]) PrependList(other *LinkedList[T]) *LinkedList[T] {
	cur := this.Head()
	if cur == nil {
		cur = &this.tailstub
	}
	for val := range other.IterValsReversed() {
		node := &Node[T]{Val: val}
		cur.prev, node.next = node, cur
		cur = cur.prev
		this.length++
	}
	cur.prev = &this.headstub
	this.headstub.next = cur
	return this
}

// New initializes a doubly linked list and returns a pointer to it.
func New[T any]() *LinkedList[T] {
	dll := LinkedList[T]{}
	dll.headstub.next = &dll.tailstub
	dll.tailstub.prev = &dll.headstub
	return &dll
}
