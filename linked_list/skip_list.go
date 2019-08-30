package linked_list

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type SkipList struct {
	lines []*sknode
	p     float64
	lpath *sknodeStack
}

type sknode struct {
	k     int
	v     interface{}
	next  *sknode
	nextl *sknode
}

// NewSkipList create new skip list
func NewSkipList(p float64) *SkipList {
	if p < 0 || p >= 1.0 {
		panic("invalid p argument should be [0, 1.0)")
	}
	lines := make([]*sknode, 1, 10)
	lines[0] = &sknode{k: -math.MaxInt32, v: nil, next: nil, nextl: nil}
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
	return sl.count(sl.lines[0])
}

// Insert new element to list with k key and v data
func (sl *SkipList) Insert(k int, v interface{}) {
	nnode := &sknode{k: k, v: v, next: nil, nextl: nil}
	node := sl.findNode(k, sl.getFastestLineNode())
	// insert to base line
	nnode.next = node.next
	node.next = nnode
	sl.promote(nnode)
}

func (sl *SkipList) findNode(k int, start *sknode) *sknode {
	// find where to change lines
	for ; start.next != nil && start.next.k < k; start = start.next {
	}
	if start.nextl == nil {
		// base line
		return start
	}
	sl.lpath.push(start)
	return sl.findNode(k, start.nextl)
}

func (sl *SkipList) promote(basenode *sknode) {
	lvl := 0
	node := basenode
	for {
		if rand.Float64() > sl.p {
			node = &sknode{k: basenode.k, v: basenode, next: nil, nextl: node}
			prevLn := sl.lpath.pop()
			if prevLn == nil {
				// add new express line
				newstartnode := &sknode{
					k: -math.MaxInt32, v: nil,
					next: nil, nextl: sl.lines[lvl]}
				sl.lines = append(sl.lines, newstartnode)
				newstartnode.next = node
				lvl++
				continue
			}
			node.next = prevLn.next
			prevLn.next = node
			continue
		}
		// reset path
		sl.lpath.reset()
		return
	}
}

func (sl *SkipList) getFastestLineNode() *sknode {
	if len(sl.lines) == 0 {
		return sl.lines[0]
	}
	return sl.lines[len(sl.lines)-1]
}

func (sl *SkipList) ToString() string {
	var str strings.Builder
	basecount := 0
	start := sl.lines[0]
	basespace := 5
	str.WriteString("head>")
	str.WriteString(fmt.Sprintf("%p>", start))
	for ; start != nil; start = start.next {
		str.WriteString(strconv.Itoa(start.k))
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
		str.WriteString(fmt.Sprintf("%p>", start))
		for ; start != nil; start = start.next {
			str.WriteString(strconv.Itoa(start.k))
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
