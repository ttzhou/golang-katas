package singlylinkedlist

import (
	"fmt"
	"iter"
	"strings"
)

type Node[T any] struct {
	Val  T
	next *Node[T]
}

// LinkedList represents a singly linked list.
type LinkedList[T any] struct {
	stub   Node[T]
	length int
}

// String returns a human-readable representation of the linked list.
func (sll LinkedList[T]) String() string {
	nodes := make([]string, sll.length)
	cur := &sll.stub
	for i := range sll.length {
		cur = cur.next
		nodes[i] = fmt.Sprintf("%v", cur.Val)
	}
	return strings.Join(nodes, "->")
}

// Size returns the number of nodes in the linked list.
func (sll LinkedList[T]) Size() int {
	return sll.length
}

// Get returns the node at position `pos` in the singly linked list. If `pos`
// is negative or greater than the length of the linked list at time of
// insertion, an error will be returned.
func (sll *LinkedList[T]) Get(pos int) (*Node[T], error) {
	if pos < 0 || pos >= sll.length {
		return nil, fmt.Errorf("invalid pos %d for length %d", pos, sll.length)
	}
	cur := &sll.stub
	for ; pos >= 0; pos-- {
		cur = cur.next
	}
	return cur, nil
}

func (sll *LinkedList[T]) insertNode(pos int, node *Node[T]) (*Node[T], error) {
	if pos < 0 || pos > sll.length {
		return nil, fmt.Errorf("invalid pos %d for length %d", pos, sll.length)
	}
	cur := &sll.stub
	for ; pos > 0; pos-- {
		cur = cur.next
	}
	node.next, cur.next = cur.next, node
	sll.length++
	return node, nil
}

// Insert inserts a node with value val into the singly linked list so that
// it's index will be `pos` after insertion. Returns a pointer to the inserted
// node. If `pos` is negative or greater than the length of the linked list at
// time of insertion, an error will be returned.
func (sll *LinkedList[T]) Insert(pos int, val T) (*Node[T], error) {
	node := &Node[T]{Val: val}
	return sll.insertNode(pos, node)
}

// Remove removes a node from the singly linked list at
// Returns a pointer to the removed node.
// If index is not within bounds (negative or beyond end of list),
// an error will be returned.
func (sll *LinkedList[T]) Remove(pos int) (*Node[T], error) {
	if pos < 0 || pos >= sll.length {
		return nil, fmt.Errorf("invalid pos %d for length %d", pos, sll.length)
	}
	prev, cur := &sll.stub, sll.stub.next
	for ; pos > 0; pos-- {
		prev, cur = cur, cur.next
	}
	prev.next = cur.next
	sll.length--
	return cur, nil
}

// Prepend adds a node to the head of the linked list,
// and returns the node.
func (sll *LinkedList[T]) Prepend(val T) *Node[T] {
	node, _ := sll.Insert(0, val)
	return node
}

// Append adds a node to the tail of the linked list,
// and returns the node.
func (sll *LinkedList[T]) Append(val T) *Node[T] {
	node, _ := sll.Insert(sll.length, val)
	return node
}

// PopHead removes the node at the head of the linked list,
// and returns the removed node. If there are no nodes
// to remove, nil is returned.
func (sll *LinkedList[T]) PopHead() *Node[T] {
	node, _ := sll.Remove(0)
	return node
}

// PopTail removes the node at the head of the linked list,
// and returns the removed node. If there are no nodes
// to remove, nil is returned.
func (sll *LinkedList[T]) PopTail() *Node[T] {
	node, _ := sll.Remove(max(0, sll.length-1))
	return node
}

// MoveToHead moves the node at position `pos` to the head of the linked list,
// and returns the node.
func (sll *LinkedList[T]) MoveToHead(pos int) *Node[T] {
	node, _ := sll.Remove(pos)
	_, _ = sll.insertNode(0, node)
	return node
}

// MoveToTail moves the node at position `pos` to the tail of the linked list,
// and returns the node.
func (sll *LinkedList[T]) MoveToTail(pos int) *Node[T] {
	node, _ := sll.Remove(pos)
	_, _ = sll.insertNode(sll.length, node)
	return node
}

// Head returns the head node, or nil if it does not exist.
func (sll LinkedList[T]) Head() *Node[T] {
	return sll.stub.next
}

// Tail returns the last node, or nil if it does not exist.
func (sll LinkedList[T]) Tail() *Node[T] {
	node, _ := sll.Get(max(sll.length-1, 0))
	return node
}

// Reverse reverses the linked list in place, and returns the head
// of the linked list.
func (sll *LinkedList[T]) Reverse() *Node[T] {
	var prev *Node[T] = nil
	cur := sll.Head()
	for cur != nil {
		prev, cur, cur.next = cur, cur.next, prev
	}
	sll.stub.next = prev
	return sll.Head()
}

// Swap swaps the linked list nodes at positions i and j.
func (sll *LinkedList[T]) Swap(i, j int) error {
	if sll.length < 2 {
		return fmt.Errorf("not enough elements to swap")
	}
	if i < 0 || i >= sll.length || j < 0 || j >= sll.length {
		return fmt.Errorf("invalid pos %d, %d for length %d", i, j, sll.length)
	}
	if i == j {
		return nil
	}
	i, j = min(i, j), max(i, j)
	ip, jp, in, jn := &sll.stub, &sll.stub, &sll.stub, &sll.stub
	cur := sll.Head()
	for p := 0; p <= j; p++ {
		if p == i-1 {
			ip = cur
		}
		if p == j-1 {
			jp = cur
		}
		if p == i {
			in = cur
		}
		if p == j {
			jn = cur
		}
		cur = cur.next
	}
	if in.next == jn {
		ip.next, jn.next, in.next = jn, in, jn.next
	} else {
		ip.next, in.next, jp.next, jn.next = jn, jn.next, in, in.next
	}
	return nil
}

// IterNodes returns an iterator over the nodes in the linked list.
func (sll LinkedList[T]) iterNodes() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		cur := sll.Head()
		for range sll.length {
			if !yield(cur) {
				return
			}
			cur = cur.next
		}
	}
}

// IterVals returns an iterator over the values in the linked list.
func (sll LinkedList[T]) IterVals() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := range sll.iterNodes() {
			if !yield(node.Val) {
				return
			}
		}
	}
}

// Copy makes a copy of the linked list's nodes and returns a
// pointer to it.
func (this LinkedList[T]) Copy() *LinkedList[T] {
	cpy := New[T]()
	cur := &cpy.stub
	for val := range this.IterVals() {
		cur.next = &Node[T]{Val: val}
		cur = cur.next
		cpy.length++
	}
	return cpy
}

// AppendList takes another linked list, adds its nodes to the end
// of the receiver, and then returns the modified receiver.
// The other linked list is not modified.
func (this *LinkedList[T]) AppendList(other *LinkedList[T]) *LinkedList[T] {
	cur := &this.stub
	for cur.next != nil {
		cur = cur.next
	}
	for val := range other.IterVals() {
		cur.next = &Node[T]{Val: val}
		cur = cur.next
		this.length++
	}
	return this
}

// PrependList takes another linked list, adds its nodes to the end
// of the receiver, and then returns the modified receiver.
func (this *LinkedList[T]) PrependList(other *LinkedList[T]) *LinkedList[T] {
	cpy := other.Copy()
	cur := &cpy.stub
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = this.Head()
	this.stub = cpy.stub
	this.length += cpy.length
	return this
}

// New initializes a singly linked list and returns a pointer to it.
func New[T any]() *LinkedList[T] {
	sll := LinkedList[T]{}
	return &sll
}
