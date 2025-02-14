package bst

import (
	"cmp"
	"fmt"
	"golang-katas/ds/queue"
	"iter"
)

// A (BST) Node has key K of type `cmp.Ordered` and value V of type `any`
type Node[K cmp.Ordered, V any] struct {
	Key         K
	Val         V
	Left, Right *Node[K, V]
	Self        **Node[K, V]
}

func (n *Node[K, V]) inOrderTraverse(yield func(*Node[K, V]) bool) bool {
	return n == nil ||
		(n.Left.inOrderTraverse(yield) &&
			yield(n) &&
			n.Right.inOrderTraverse(yield))
}

func (n *Node[K, V]) preOrderTraverse(yield func(*Node[K, V]) bool) bool {
	return n == nil ||
		(yield(n) &&
			n.Left.preOrderTraverse(yield) &&
			n.Right.preOrderTraverse(yield))
}

func (n *Node[K, V]) postOrderTraverse(yield func(*Node[K, V]) bool) bool {
	return n == nil ||
		(n.Left.postOrderTraverse(yield) &&
			n.Right.postOrderTraverse(yield) &&
			yield(n))
}

func (n *Node[K, V]) levelOrderTraverse(yield func(*Node[K, V]) bool, q *queue.Queue[*Node[K, V]]) {
	if n == nil {
		return
	}
	for !q.IsEmpty() {
		node, _ := q.Pop()
		if !yield(node) {
			return
		}
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}
}

func (n *Node[K, V]) successor() *Node[K, V] {
	if n.Right == nil {
		return nil
	}
	cur := n.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// A BinarySearchTree is a struct with (BST) Node `root`
// and a field Size which tracks the number of nodes in the BST.
type BinarySearchTree[K cmp.Ordered, V any] struct {
	root *Node[K, V]
	Size int
}

func insert[K cmp.Ordered, V any](n **Node[K, V], key K, val V) error {
	node := &Node[K, V]{Key: key, Val: val}
	for *n != nil {
		if key < (*n).Key {
			n = &(*n).Left
			node.Self = n
		} else if key > (*n).Key {
			n = &(*n).Right
			node.Self = n
		} else {
			return fmt.Errorf("key already exists in tree")
		}
	}
	*n = node
	return nil
}

func get[K cmp.Ordered, V any](n *Node[K, V], key K) *Node[K, V] {
	cur := &n
	for *cur != nil {
		if key < (*cur).Key {
			cur = &(*cur).Left
		} else if key > (*cur).Key {
			cur = &(*cur).Right
		} else {
			break
		}
	}
	return *cur
}

func contains[K cmp.Ordered, V any](n *Node[K, V], key K) bool {
	if n == nil {
		return false
	}
	if n.Key == key {
		return true
	}
	return contains(n.Left, key) || contains(n.Right, key)
}

func delete[K cmp.Ordered, V any](n **Node[K, V], key K) error {
	for *n != nil {
		if key < (*n).Key {
			n = &((*n).Left)
		} else if key > (*n).Key {
			n = &((*n).Right)
		} else {
			successor := (*n).successor()
			if successor == nil {
				self := (*n).Self
				*n = (*n).Left
				if self != nil {
					*self = *n
				}
				return nil
			}
			(*n).Key, successor.Key = successor.Key, (*n).Key
			(*n).Val, successor.Val = successor.Val, (*n).Val
			return delete(&successor, key)
		}
	}
	return fmt.Errorf("could not find node with key %v", key)
}

func height[K cmp.Ordered, V any](n *Node[K, V]) int {
	if n == nil {
		return 0
	}
	return 1 + max(height(n.Left), height(n.Right))

}

// Insert inserts a value V with key K into a BinarySearchTree.
// If the key already exists in the tree, then an error is returned.
func (bst *BinarySearchTree[K, V]) Insert(key K, val V) error {
	err := insert(&bst.root, key, val)
	if err == nil {
		bst.Size++
	}
	return err
}

// Get returns the value V of the node with key K in the BinarySearchTree.
// If the key K cannot be found, a non-nil error is returned and the
// returned value is the zero value of the type V.
func (bst BinarySearchTree[K, V]) Get(key K) (V, error) {
	node := get(bst.root, key)
	if node == nil {
		var z V
		return z, fmt.Errorf("could not find key '%v'", key)
	}
	return node.Val, nil
}

// Contains returns a boolean indicating whether there is a node
// with key K in the BinarySearchTree.
func (bst BinarySearchTree[K, V]) Contains(key K) bool {
	return contains(bst.root, key)
}

// Delete removes the node with key K in the BinarySearchTree.
// If there is no such node, an error is returned.
func (bst *BinarySearchTree[K, V]) Delete(key K) error {
	err := delete(&bst.root, key)
	if err == nil {
		bst.Size--
	}
	return err
}

// Height returns the number of levels in the BinarySearchTree.
// An empty BinarySearchTree has height 0.
func (bst BinarySearchTree[K, V]) Height() int {
	return height(bst.root)
}

type traversalOrder int

const (
	PRE_ORDER traversalOrder = iota
	IN_ORDER
	POST_ORDER
	LEVEL_ORDER
	LEVELS
)

func (bst BinarySearchTree[K, V]) traverseNodes(to traversalOrder) iter.Seq[*Node[K, V]] {
	return func(yield func(*Node[K, V]) bool) {
		switch to {
		case PRE_ORDER:
			bst.root.preOrderTraverse(yield)
		case IN_ORDER:
			bst.root.inOrderTraverse(yield)
		case POST_ORDER:
			bst.root.postOrderTraverse(yield)
		case LEVEL_ORDER:
			q := queue.New[*Node[K, V]]()
			q.Push(bst.root)
			bst.root.levelOrderTraverse(yield, q)
		}
	}
}

// Traverse returns an iterator over the nodes of the BinarySearchTree
// following the traversal order specified.
func (bst BinarySearchTree[K, V]) Traverse(to traversalOrder) iter.Seq[K] {
	return func(yield func(K) bool) {
		for node := range bst.traverseNodes(to) {
			if !yield(node.Key) {
				return
			}
		}
	}
}

// TraverseLevels returns an iterator of string slices,
// each slice being a representation of the keys in a given level
// of the BinarySearchTree.
func (bst BinarySearchTree[K, V]) TraverseLevels() iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		if bst.root == nil {
			return
		}
		q := queue.New[*Node[K, V]]()
		q.Push(bst.root)
		for !q.IsEmpty() {
			level := make([]*Node[K, V], 0, q.Size())
			levelKeys := make([]string, 0, q.Size()*2)
			var numKeys int
			for !q.IsEmpty() {
				node, _ := q.Pop()
				level = append(level, node)
				s := "_"
				if node != nil {
					numKeys++
					s = fmt.Sprintf("%v", node.Key)
				}
				levelKeys = append(levelKeys, s)
			}
			if numKeys <= 0 {
				return
			}
			if !yield(levelKeys) {
				return
			}
			for _, node := range level {
				if node != nil {
					q.Push(node.Left)
					q.Push(node.Right)
				} else {
					q.Push(nil)
					q.Push(nil)
				}
			}
		}
	}
}

// New initializes and returns an empty BinarySearchTree.
func New[K cmp.Ordered, V any]() *BinarySearchTree[K, V] {
	return &BinarySearchTree[K, V]{}
}
