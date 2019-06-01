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

type Tree3_13 struct {
	id          int
	val         int
	sum         int
	left, right *Tree3_13
}

func (t *Tree3_13) ToString() string {
	var out strings.Builder
	f := func(t *Tree3_13) {
		out.WriteString("(id=")
		out.WriteString(strconv.Itoa(t.id))
		out.WriteString(",val=")
		out.WriteString(strconv.Itoa(t.val))
		out.WriteString(",sum=")
		out.WriteString(strconv.Itoa(t.sum))
		out.WriteRune(')')
	}
	t.TraverseLevelOrder(f)
	return out.String()
}

func (t *Tree3_13) TraverseInOrder(f func(t *Tree3_13)) {
	if t.left != nil {
		t.left.TraverseInOrder(f)
	}
	f(t)
	if t.right != nil {
		t.right.TraverseInOrder(f)
	}
}
func (t *Tree3_13) TraverseLevelOrder(f func(t *Tree3_13)) {
	q := []*Tree3_13{}
	push := func(t *Tree3_13) {
		q = append(q, t)
	}
	pop := func() *Tree3_13 {
		tr := q[0]
		q = q[1:]
		return tr
	}

	push(t)
	func(t *Tree3_13, f func(t *Tree3_13)) {
		for len(q) > 0 {
			node := pop()
			f(node)
			if node.left != nil {
				push(node.left)
			}
			if node.right != nil {
				push(node.right)
			}
		}
	}(t, f)

}

type ADMProblem3_13_Solution struct {
	// balanced binary tree
	t *Tree3_13
}

func (s *ADMProblem3_13_Solution) Add(i, y int) {
	var add func(t *Tree3_13, i, y int)
	add = func(t *Tree3_13, i, y int) {
		t.sum += y
		if t.id == i {
			t.val += y
			return
		}
		if t.id > i {
			add(t.left, i, y)
			return
		}
		add(t.right, i, y)
	}
	add(s.t, i, y)
}

func (s *ADMProblem3_13_Solution) PartialSum(i int) int {
	var partialSum func(t *Tree3_13, id, sum int) int

	partialSum = func(t *Tree3_13, id, sum int) int {
		leftSum := 0
		if t.left != nil {
			leftSum = t.left.sum
		}
		if t.id == id {
			return sum + t.val + leftSum
		}

		if t.id > id {
			return partialSum(t.left, i, sum)
		}
		return partialSum(t.right, i, sum+t.val+leftSum)
	}

	if s.t.id == i {
		if s.t.left != nil {
			return s.t.val + s.t.left.sum
		} else {
			return s.t.val
		}
	}
	if s.t.id > i {
		return partialSum(s.t.left, i, 0)
	}
	return partialSum(s.t.right, i, s.t.val+s.t.left.sum)
}

func ADMProblem3_13(set []int) ADMProblem3_13_Solution {
	if len(set) == 0 {
		return ADMProblem3_13_Solution{&Tree3_13{}}
	}

	sum := func(set []int) int {
		sum := 0
		for i := 0; i < len(set); i++ {
			sum += set[i]
		}
		return sum
	}
	nodes := []*Tree3_13{}
	for i := 0; i < len(set); i++ {
		t := &Tree3_13{
			id:    i + 1,
			val:   set[i],
			left:  nil,
			right: nil,
		}
		nodes = append(nodes, t)
	}

	var assign func(nodes []*Tree3_13, set []int) *Tree3_13
	assign = func(nodes []*Tree3_13, set []int) *Tree3_13 {
		if len(nodes) == 0 {
			return nil
		}
		id := 0
		if len(nodes) > 1 {
			id = len(nodes)/2 - 1
		}
		tr := nodes[id]
		tr.sum = sum(set)
		tr.left = assign(nodes[0:id], set[0:id])
		tr.right = assign(nodes[id+1:], set[id+1:])
		return tr
	}

	t := assign(nodes, set)
	return ADMProblem3_13_Solution{t}
}
