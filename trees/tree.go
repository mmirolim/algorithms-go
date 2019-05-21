package tree

import (
	"strconv"
	"strings"
)

type Tree struct {
	root *Node
}

type Node struct {
	val         int
	left, right *Node
}

func (t *Tree) Reverse() {
	if t == nil || t.root == nil {
		return
	}

	reverse := func(node *Node) {
		node.left, node.right = node.right, node.left
	}
	qu := newQueue()
	qu.push(t.root)
	for node := qu.pop(); node != nil; node = qu.pop() {
		if node.left != nil {
			qu.push(node.left)
		}
		if node.right != nil {
			qu.push(node.right)
		}
		reverse(node)
	}
}

type queue struct {
	store []*Node
}

func newQueue() *queue {
	return &queue{store: []*Node{}}
}

func (q *queue) push(node *Node) {
	q.store = append(q.store, node)
}

func (q *queue) pop() *Node {
	if len(q.store) == 0 {
		return nil
	}
	node := q.store[0]
	q.store = q.store[1:]
	return node
}

func (t *Tree) String() string {
	if t == nil || t.root == nil {
		return "empty"
	}
	queue := newQueue()
	var str strings.Builder
	queue.push(t.root)
	for node := queue.pop(); node != nil; node = queue.pop() {
		if node.left != nil {
			queue.push(node.left)
		}
		if node.right != nil {
			queue.push(node.right)
		}
		str.WriteString(strconv.Itoa(node.val))
	}
	return str.String()
}
