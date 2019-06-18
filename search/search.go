package search

import (
	"sort"

	tree "github.com/mmirolim/algos/trees"
)

// ADM 4.15 use tournament algorithm
func Find2thElementTA(in []int, cmp func(a, b int) bool) int {
	type node struct {
		v    int
		l, r *node
	}
	newNode := func(v int) *node {
		return &node{v, nil, nil}
	}
	nodes := []*node{}

	pop := func() *node {
		n := nodes[0]
		nodes = nodes[1:]
		return n
	}

	push := func(n *node) {
		nodes = append(nodes, n)
	}
	var root *node
	// can be optimized to create only roots
	// and skip leaves creation just store index value
	for i := 0; i < len(in); i += 2 {
		n1 := newNode(i)
		if i+1 < len(in) {
			n2 := newNode(i + 1)
			if cmp(in[n1.v], in[n2.v]) {
				root = newNode(n1.v)
			} else {
				root = newNode(n2.v)
			}
			root.l = n1
			root.r = n2
			push(root)
		} else {
			push(n1)
		}
	}

	for len(nodes) != 1 {
		// check if it odd number of elements
		lastId := len(nodes) - 1
		if len(nodes)&1 == 1 {
			lastId = len(nodes) - 2
		}
		// iterate in pairs
		for i := 0; i < lastId; i += 2 {
			// get pair
			n1 := pop()
			n2 := pop()
			if cmp(in[n1.v], in[n2.v]) {
				root = newNode(n1.v)
			} else {
				root = newNode(n2.v)
			}
			root.l = n1
			root.r = n2
			push(root)
		}
		if len(nodes)&1 == 1 {
			// push odd element to the end
			push(pop())
		}
	}
	// find second elm
	var walk func(*node)
	var elm *int
	walk = func(n *node) {
		if n.l == nil {
			return
		}
		if n.v == n.l.v {
			if elm == nil {
				elm = &n.r.v
			} else if cmp(in[n.r.v], *elm) {
				elm = &in[n.r.v]
			}
			walk(n.l)
		} else {
			if elm == nil {
				elm = &n.l.v
			} else if cmp(in[n.l.v], *elm) {
				elm = &in[n.l.v]
			}
			walk(n.r)
		}
	}
	walk(root)
	return *elm
}

func FindKthElement(in []int, k int, cmp func(a, b int) bool) int {
	if len(in) < k {
		return -1
	}
	var kheap *tree.Heap
	if cmp(1, 2) {
		kheap, _ = tree.NewMaxHeapFrom(in[:k], nil)
	} else {
		kheap, _ = tree.NewMinHeapFrom(in[:k], nil)
	}

	for i := k; i < len(in); i++ {
		v, _ := kheap.PeekKey()
		if cmp(in[i], v) {
			kheap.ExtractTop()
			_ = kheap.Insert(in[i], nil)
		}
	}
	kth, _ := kheap.PeekKey()
	return kth
}

func FindKthElementHeapifyAll(in []int, k int, cmp func(a, b int) bool) int {
	var kheap *tree.Heap
	if cmp(1, 2) {
		kheap, _ = tree.NewMinHeapFrom(in, nil)
	} else {
		kheap, _ = tree.NewMaxHeapFrom(in, nil)
	}
	kth := 0
	for i := 0; i < k; i++ {
		kth, _, _ = kheap.ExtractTop()
	}
	return kth
}
func FindKthMinElementLinear(in []int, k int) int {
	kths := []int{}
	for i := 0; i < k; i++ {
		kths = append(kths, in[i])
		sort.Ints(kths)
	}

	for i := k; i < len(in); i++ {
		if kths[len(kths)-1] < in[i] {
			continue
		}
		for i := range kths {
			if kths[i] > in[i] {
				// move
				copy(kths[:i+1], kths)
				kths[i] = in[i]
				break
			}
		}
	}

	return kths[len(kths)-1]
}
