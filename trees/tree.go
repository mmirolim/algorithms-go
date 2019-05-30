package tree

import (
	"strconv"
	"strings"
)

type Tree struct {
	val                 int
	parent, left, right *Tree
}

func Delete(t **Tree, val int) {
	if (*t).val == val {
		// case leaf node
		if (*t).left == nil && (*t).right == nil {
			*t = nil
			return
		} else if (*t).right == nil {
			// case one left node
			(*t).left.parent = (*t).parent
			if (*t).parent.left != nil && (*t).parent.left.val == (*t).val {
				(*t).parent.left = (*t).left
			} else {
				(*t).parent.right = (*t).left
			}
		} else if (*t).left == nil {
			// case one right node
			(*t).right.parent = (*t).parent
			if (*t).parent.left != nil && (*t).parent.left.val == (*t).val {
				(*t).parent.left = (*t).right
			} else {
				(*t).parent.right = (*t).right

			}
		} else {
			// case has subtree
			min := (*t).right
			for ; min.left != nil; min = min.left {
			}
			if min.right != nil {
				// case one right node
				min.right.parent = min.parent
				min.parent.left = min.right
			}
			(*t).val = min.val
			min.parent.left = nil
		}
	} else if (*t).val > val {
		Delete(&(*t).left, val)
	} else {
		Delete(&(*t).right, val)
	}
}

func Insert(t **Tree, val int) {
	var insert func(t **Tree, p *Tree)
	insert = func(t **Tree, p *Tree) {
		if *t == nil {
			*t = &Tree{
				val:    val,
				parent: p,
				left:   nil,
				right:  nil}
		} else if (*t).val > val {
			insert(&(*t).left, *t)
		} else {
			insert(&(*t).right, *t)
		}
	}

	insert(t, nil)
}

func (t *Tree) FindMin() int {
	min := t
	for ; min.left != nil; min = min.left {
	}

	return min.val
}

func (t *Tree) FindMax() int {
	max := t
	for ; max.right != nil; max = max.right {
	}

	return max.val
}

func (t *Tree) Search(val int) bool {
	if t == nil {
		return false
	}
	if t.val == val {
		return true
	}

	if t.val > val {
		return t.right.Search(val)
	}
	return t.left.Search(val)
}

func (t *Tree) IsMirrorRecur(t2 *Tree) bool {
	if t == nil && t2 == nil {
		return true
	}
	var isMirror func(n1, n2 *Tree) bool
	isMirror = func(n1, n2 *Tree) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		return n1.val == n2.val &&
			isMirror(n1.left, n2.right) && isMirror(n1.right, n2.left)
	}

	return isMirror(t, t2)
}

func (t *Tree) IsMirror(t2 *Tree) bool {
	if t == nil && t2 == nil {
		return true
	}

	tq1 := newQueue()
	tq2 := newQueue()

	t1node := t
	t2node := t2

	for ; t1node != nil && t2node != nil; t1node, t2node = tq1.pop(), tq2.pop() {
		if t1node.val != t2node.val {
			return false
		}
		nodes := []*Tree{t1node.left, t2node.right, t1node.right, t2node.left}
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
	if t == nil {
		return
	}

	reverse := func(node *Tree) {
		node.left, node.right = node.right, node.left
	}
	qu := newQueue()
	qu.push(t)
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
	store []*Tree
}

func newQueue() *queue {
	return &queue{store: []*Tree{}}
}

func (q *queue) push(node *Tree) {
	q.store = append(q.store, node)
}

func (q *queue) pop() *Tree {
	if len(q.store) == 0 {
		return nil
	}
	node := q.store[0]
	q.store = q.store[1:]
	return node
}

func (t *Tree) ToString() string {
	if t == nil {
		return "empty"
	}
	queue := newQueue()
	var str strings.Builder
	queue.push(t)
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
