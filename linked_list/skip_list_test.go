package linked_list

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSkipListInsert(t *testing.T) {
	sl := NewSkipList(0.5)
	nodesnum := 15
	for _, v := range rand.Perm(nodesnum) {
		sl.Insert(v, v)
	}
	if sl.Count() != nodesnum {
		t.Errorf("expected number of nodes %d, got %d\n", nodesnum, sl.Count())
		t.FailNow()
	}
	for start := sl.lines[0]; start.next != nil; start = start.next {
		if start.k > start.next.k {
			t.Errorf("expected ordered list got current %d next %d\n", start.k, start.next.k)
			break
		}
	}
	fmt.Println(sl.ToString())
}

func BenchmarkInsertInOrderSkipListInsert(b *testing.B) {
	sk := NewSkipList(0.5)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sk.Insert(i, i)
	}
}

func BenchmarkInsertInOrderList(b *testing.B) {
	list := NewList()
	list.head = &Node{-1, nil}
	b.ReportAllocs()
	set := false
	for i := 0; i < b.N; i++ {
		// insert in order
		current := list.head
		set = false
		for current.next != nil {
			if current.next.val > i {
				node := &Node{i, current.next}
				current.next = node
				set = true
				break
			}
			current = current.next
		}
		if !set {
			current.next = &Node{i, nil}
		}
	}
}
