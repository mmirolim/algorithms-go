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

func (t *Tree) IsMirrorRecur(t2 *Tree) bool {
	if t.root == nil && t2.root == nil {
		return true
	}
	var isMirror func(n1, n2 *Node) bool
	isMirror = func(n1, n2 *Node) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		return n1.val == n2.val &&
			isMirror(n1.left, n2.right) && isMirror(n1.right, n2.left)
	}

	return isMirror(t.root, t2.root)
}
func (t *Tree) IsMirror(t2 *Tree) bool {
	if t.root == nil && t2.root == nil {
		return true
	}

	tq1 := newQueue()
	tq2 := newQueue()

	t1node := t.root
	t2node := t2.root

	for ; t1node != nil && t2node != nil; t1node, t2node = tq1.pop(), tq2.pop() {
		if t1node.val != t2node.val {
			return false
		}
		nodes := []*Node{t1node.left, t2node.right, t1node.right, t2node.left}
		for i := 0; i < len(nodes)-1; i += 2 {
			if nodes[i] != nil && nodes[i+1] != nil {
				tq1.push(nodes[i])
				tq2.push(nodes[i+1])
			} else if nodes[i] != nil || nodes[i+1] != nil {
				return false
			}

		}

	}

	return true
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
