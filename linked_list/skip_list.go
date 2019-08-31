package linked_list

import (
	"fmt"
	"math/rand"
	"strings"
)

const promoteChance = 0.5

// SkipList container
type SkipList struct {
	size  uint64
	lines []*sknode
	p     float64
	lpath *sknodeStack
	rng   *rand.Rand
}

type sknode struct {
	k            int64
	v            interface{}
	prev, next   *sknode
	dist         uint64
	nextl, prevl *sknode
}

// reset pointer for GC
func (n *sknode) reset() {
	n.next = nil
	n.nextl = nil
	n.v = nil
}

// NewSkipList create new skip list
// seed for random generator
func NewSkipList(seed int64) *SkipList {
	lines := make([]*sknode, 1, 10)
	lines[0] = &sknode{}
	return &SkipList{
		0, lines, promoteChance,
		&sknodeStack{}, rand.New(rand.NewSource(seed))}
}

func (sl *SkipList) count(node *sknode) uint64 {
	var count uint64
	for start := node; start != nil; start = start.next {
		count++
	}
	return count
}

// Count returns size of skiplist
func (sl *SkipList) Count() uint64 {
	return sl.size
}

// Insert new element to list with k key and v data
func (sl *SkipList) Insert(key int, v interface{}) {
	k := int64(key)
	sl.lpath.reset()
	nnode := &sknode{k: k, v: v, next: nil, nextl: nil}
	node, pos := sl.findNode(k)
	// insert to base line
	if node != sl.lines[0] && node.k == k { // skip head
		// update
		node.v = v
		return
	}
	sl.size++
	nnode.dist = 1
	connectNodes(&node, &nnode)
	sl.promote(nnode, pos+1)
	// skip column, dist already set by promote()
	for ; nnode.prevl != nil; nnode = nnode.prevl {
	}
	// update distance for all right top nnodes
	for nnode = nnode.next; nnode != nil; nnode = nnode.next {
		for clmnode := nnode.prevl; clmnode != nil; clmnode = clmnode.prevl {
			clmnode.dist++
			nnode = clmnode
		}

	}
}

// Find returns true and nodes's value by key or false and nil
func (sl *SkipList) Find(key int) (bool, interface{}) {
	k := int64(key)
	node, _ := sl.findNode(k)
	if node.k == k {
		return true, node.v
	}
	return false, nil
}

// Delete node by key
func (sl *SkipList) Delete(key int) {
	k := int64(key)
	sl.lpath.reset()
	node, _ := sl.findNode(k)
	if node.k != k {
		return // not found
	}
	sl.size--
	// detach deleted node from list
	detachNode(node, true)
	prevl, _ := sl.lpath.pop()
	if prevl.k == k { // delete column of deleted node
		for ; prevl != nil && prevl.k == k; prevl, _ = sl.lpath.pop() {
			detachNode(prevl, false)
			node.reset()
		}
	}
	node.reset()
}

// returns node and its position
func (sl *SkipList) findNode(k int64) (*sknode, uint64) {
	var pos uint64
	start := sl.lines[len(sl.lines)-1]
	if start.next != nil {
		pos += start.next.dist - 1 // pos starts from 0
	}
	for l := len(sl.lines) - 1; l >= 0; l-- {
		// find where to change lines
		for ; start.next != nil && start.next.k <= k; start = start.next {
			pos += start.next.dist
		}
		if start.nextl == nil {
			return start, pos
		}
		sl.lpath.push(start, pos)
		start = start.nextl // change line
	}
	return nil, pos
}

// GetByIndex return value at pos, starts from 1
// or false and nil if out of range
func (sl *SkipList) GetByIndex(pos uint64) (bool, interface{}) {
	if pos == 0 || pos >= sl.size {
		return false, nil
	}
	start := sl.lines[len(sl.lines)-1]
	for l := len(sl.lines) - 1; l >= 0; l-- {
		// find where to change lines
		for ; start.next != nil && pos >= start.next.dist && pos-start.next.dist >= 0; start = start.next {
			pos -= start.next.dist

		}
		if start.nextl == nil {
			return true, start.v
		}
		sl.lpath.push(start, pos)
		start = start.nextl // change line
	}
	return false, nil
}

func (sl *SkipList) promote(basenode *sknode, pos uint64) {
	node := basenode
	for {
		if sl.rng.Float64() < sl.p {
			nnode := &sknode{}
			nnode.k = basenode.k
			nnode.nextl = node
			node.prevl = nnode
			prevLn, prevpos := sl.lpath.pop()
			if prevLn == nil {
				// add new express line
				prevLn = sl.addNewExpressLine()
				nnode.dist = pos
			} else {
				// update distances between nodes
				nnode.dist = pos - prevpos
				if prevLn.next != nil {
					prevLn.next.dist = prevLn.next.dist - nnode.dist + 1

				}
			}
			connectNodes(&prevLn, &nnode)
			node = nnode
			continue
		}
		// reset path
		sl.lpath.reset()
		return
	}
}

// double link <->n1<->n2<->
func connectNodes(node, next **sknode) {
	prev := *node
	nxt := *next
	nxt.prev = prev
	nxt.next = prev.next
	if prev.next != nil {
		prev.next.prev = nxt
	}
	prev.next = nxt
}

// double link prev<->node<->next => prev<->next
func detachNode(node *sknode, baselvl bool) {
	node.prev.next = node.next
	if node.next != nil {
		if !baselvl {
			node.next.dist += node.next.dist + node.dist
		}
		node.next.prev = node.prev
	}
}

func (sl *SkipList) addNewExpressLine() *sknode {
	nsnode := sl.newStartNode()
	sl.lines = append(sl.lines, nsnode)
	return nsnode
}

func (sl *SkipList) newStartNode() *sknode {
	return &sknode{
		k: 0, v: nil,
		next: nil, nextl: sl.lines[len(sl.lines)-1]}
}

// ToString returns skiplist serialized to string
func (sl *SkipList) ToString() string {
	var str strings.Builder
	basecount := sl.size
	var start *sknode
	var basespace uint64 = 5

	for i := 0; i < len(sl.lines); i++ {
		start = sl.lines[i]
		count := sl.count(start)
		space := basecount * basespace / count
		str.WriteString("head>")
		str.WriteString(">")
		for ; start != nil; start = start.next {
			str.WriteString(fmt.Sprintf("{%d, d: %d}", start.k, start.dist))
			str.WriteString(strings.Repeat("-", int(space)))
			str.WriteByte('>')
			count++
		}
		str.WriteString("nil\n")
	}

	return str.String()
}

type sknodeStack struct {
	s []stackElm
}

type stackElm struct {
	n   *sknode
	pos uint64
}

func (ns *sknodeStack) reset() {
	ns.s = ns.s[:0]
}

func (ns *sknodeStack) push(node *sknode, pos uint64) {
	ns.s = append(ns.s, stackElm{node, pos})
}

func (ns *sknodeStack) peek() (*sknode, uint64) {
	elm := ns.s[len(ns.s)-1]
	return elm.n, elm.pos
}

func (ns *sknodeStack) pop() (*sknode, uint64) {
	l := len(ns.s)
	if l == 0 {
		return nil, 0
	}
	elm := ns.s[l-1]
	ns.s = ns.s[:l-1]
	return elm.n, elm.pos
}
