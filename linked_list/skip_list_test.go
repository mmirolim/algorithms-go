package linked_list

import (
	"math/rand"
	"testing"
)

func TestSkipListInsert(t *testing.T) {
	sl := NewSkipList(100)
	nodesnum := 10
	for _, v := range rand.Perm(nodesnum) {
		sl.Insert(v, v)
	}
	if sl.Count() != uint64(nodesnum) {
		t.Errorf("expected number of nodes %d, got %d\n", nodesnum, sl.Count())
		t.FailNow()
	}
	// test invariants
	// all lines are sorted and double linked
	for l := 0; l < len(sl.lines); l++ {
		for start := sl.lines[l].next; start != nil && start.next != nil; start = start.next {
			// sorted
			if start.k > start.next.k {
				t.Errorf("line %d expected ordered list got current %d next %d\n", l, start.k, start.next.k)
				break
			}
			// correctly linked
			if start.next.prev != start {
				t.Errorf("line %d expected pointer to %p got %p in node %+v\n", l, start.next.prev, start, start)

				break
			}
		}
	}
}

func TestSkipListGetByIndex(t *testing.T) {
	sl := NewSkipList(100)
	nodesnum := 10
	for _, v := range rand.Perm(nodesnum) {
		sl.Insert(v, v)
	}
	for l := 1; l <= nodesnum; l++ {
		ok, v := sl.GetByIndex(uint64(l))
		var inf interface{} = l - 1
		if ok && v != inf {
			t.Errorf("expected value %d, got %d\n", l, v)
		}
	}
	ok, v := sl.GetByIndex(100)
	if ok || v != nil {
		t.Errorf("expected not found got ok %v and value %v", ok, v)
	}
}

func TestSkipListFind(t *testing.T) {
	sl := NewSkipList(100)
	nodesnum := 15
	for _, v := range rand.Perm(nodesnum) {
		sl.Insert(v, v)
	}
	testcases := []struct {
		k     int
		found bool
		v     interface{}
	}{
		{-1, false, nil}, {0, true, 0}, {15, false, nil}, {7, true, 7},
	}
	for i, c := range testcases {
		ok, v := sl.Find(c.k)
		if ok != c.found {
			t.Errorf("case [%d] expected %v got %v", i, c.found, ok)
			break
		}
		if v != c.v {
			t.Errorf("case [%d] expected value %v got %v", i, c.v, v)
			break
		}
	}
}

func TestSkipListDelete(t *testing.T) {
	sl := NewSkipList(100)
	nodesnum := 15
	for _, v := range rand.Perm(nodesnum) {
		sl.Insert(v, v)
	}
	deleteNum := 5
	for i := deleteNum; i > 0; i-- {
		sl.Delete(nodesnum - i - 1)
	}
	shouldLeft := nodesnum - deleteNum
	if sl.Count() != uint64(shouldLeft) {
		t.Errorf("expected number of nodes left %d, got %d\n", shouldLeft, sl.Count())
	}
}

func BenchmarkInsertInOrderSkipListInsert(b *testing.B) {
	sk := NewSkipList(100)
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
