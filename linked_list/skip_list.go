package linked_list

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type SkipList struct {
	lines []*sknode
	p     float64
	lpath *sknodeStack
}

type sknode struct {
	k          int64
	v          interface{}
	prev, next *sknode
	nextl      *sknode
}

// reset pointer for GC
func (n *sknode) reset() {
	n.next = nil
	n.nextl = nil
	n.v = nil
}

// NewSkipList create new skip list
// TODO seed rand and separate generator
// TODO smallest k is -math.MaxInt32
func NewSkipList(p float64) *SkipList {
	if p < 0 || p >= 1.0 {
		panic("invalid p argument should be [0, 1.0)")
	}
	lines := make([]*sknode, 1, 10)
	lines[0] = &sknode{k: -math.MaxInt64, v: nil, prev: nil, next: nil, nextl: nil}
	return &SkipList{lines, p, &sknodeStack{}}
}
func (sl *SkipList) count(node *sknode) int {
	count := 0
	for start := node; start != nil; start = start.next {
		count++
	}
	return count
}

func (sl *SkipList) Count() int {
	count := sl.count(sl.lines[0])
	if sl.lines[0].v != nil {
		return count
	}
	return count - 1
}

// Insert new element to list with k key and v data
func (sl *SkipList) Insert(key int, v interface{}) {
	k := int64(key)
	sl.lpath.reset()
	nnode := &sknode{k: k, v: v, next: nil, nextl: nil}
	node := sl.findNode(k, sl.getFastestLineNode())
	// insert to base line
	if node.k == k {
		// update
		node.v = v
		return
	}

	connectNodes(&node, &nnode)
	sl.promote(nnode)
}

func (sl *SkipList) Find(key int) (bool, interface{}) {
	k := int64(key)
	node := sl.findNode(k, sl.getFastestLineNode())
	if node.k == k {
		return true, node.v
	}
	return false, nil
}

func (sl *SkipList) Delete(key int) {
	k := int64(key)
	sl.lpath.reset()
	node := sl.findNode(k, sl.getFastestLineNode())
	if node.k != k {
		return // not found
	}
	// special case, don't delete just clear val
	if node.k == -math.MaxInt64 {
		node.v = nil
		return
	}

	// detach deleted node from list
	detachNode(node)
	prevl := sl.lpath.pop()
	if prevl.k == k { // delete column of deleted node
		for ; prevl != nil && prevl.k == k; prevl = sl.lpath.pop() {
			detachNode(prevl)
			node.reset()
		}
	}
	node.reset()
}

func (sl *SkipList) findNode(k int64, start *sknode) *sknode {
	// find where to change lines
	for ; start.next != nil && start.next.k <= k; start = start.next {
	}
	if start.nextl == nil {
		// base line
		return start
	}
	sl.lpath.push(start)
	return sl.findNode(k, start.nextl)
}

func (sl *SkipList) promote(basenode *sknode) {
	node := basenode
	for {
		// TODO use separate generator
		if rand.Float64() > sl.p {
			node = &sknode{k: basenode.k, v: nil, next: nil, nextl: node}
			prevLn := sl.lpath.pop()
			if prevLn == nil {
				// add new express line
				prevLn = sl.addNewExpressLine()
			}
			connectNodes(&prevLn, &node)
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
func detachNode(node *sknode) {
	node.prev.next = node.next
	if node.next != nil {
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
		k: -math.MaxInt64, v: nil,
		next: nil, nextl: sl.lines[len(sl.lines)-1]}
}

func (sl *SkipList) getFastestLineNode() *sknode {
	if len(sl.lines) == 0 {
		return sl.lines[0]
	}
	return sl.lines[len(sl.lines)-1]
}

// TODO add debug version
func (sl *SkipList) ToString() string {
	var str strings.Builder
	basecount := 0
	start := sl.lines[0]
	basespace := 5
	str.WriteString("head>")
	str.WriteString(">")
	for ; start != nil; start = start.next {
		str.WriteString(fmt.Sprintf("{%d, %v}", start.k, start.v))
		str.WriteString(strings.Repeat("-", basespace))
		str.WriteByte('>')
		basecount++
	}
	str.WriteString("nil\n")
	for i := 1; i < len(sl.lines); i++ {
		start = sl.lines[i]
		count := sl.count(start)
		space := basecount * basespace / count
		str.WriteString("head>")
		str.WriteString(">")
		for ; start != nil; start = start.next {
			str.WriteString(fmt.Sprintf("{%d, %v}", start.k, start.v))
			str.WriteString(strings.Repeat("-", space))
			str.WriteByte('>')
			count++
		}
		str.WriteString("nil\n")
	}

	return str.String()
}

type sknodeStack struct {
	s []*sknode
}

func (ns *sknodeStack) reset() {
	ns.s = ns.s[:0]
}

func (ns *sknodeStack) push(node *sknode) {
	ns.s = append(ns.s, node)
}

func (ns *sknodeStack) pop() *sknode {
	l := len(ns.s)
	if l == 0 {
		return nil
	}
	node := ns.s[l-1]
	ns.s = ns.s[:l-1]
	return node
}
