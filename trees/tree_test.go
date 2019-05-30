package tree

import (
	"testing"
)

var treeMirrorTestData = []struct {
	t1, t2 *Tree
	out    bool
}{
	{t1: nil, t2: nil, out: true},
	{t1: &Tree{parent: nil,
		left: &Tree{val: 2,
			left:  nil,
			right: nil},
		right: &Tree{val: 3,
			left:  nil,
			right: nil}},
		t2: &Tree{parent: nil,
			right: &Tree{val: 2,
				left:  nil,
				right: nil},
			left: &Tree{val: 3,
				left: &Tree{val: 5,
					left:  nil,
					right: nil},
				right: nil}}, out: false},
	{t1: &Tree{parent: nil,
		left: &Tree{val: 2,
			left:  nil,
			right: nil},
		right: &Tree{val: 3,
			left: &Tree{val: 4,
				left:  &Tree{val: 6, left: nil, right: nil},
				right: &Tree{val: 5, left: nil, right: nil},
			},
			right: &Tree{val: 7,
				left:  &Tree{val: 8, left: nil, right: nil},
				right: nil}}},
		t2: &Tree{parent: nil,
			right: &Tree{val: 2,
				left:  nil,
				right: nil},
			left: &Tree{val: 3,
				right: &Tree{val: 4,
					right: &Tree{val: 6, left: nil, right: nil},
					left:  &Tree{val: 5, left: nil, right: nil},
				},
				left: &Tree{val: 7,
					right: &Tree{val: 8, left: nil, right: nil},
					left:  nil}}},
		out: true},
}

func TestTreeDelete(t *testing.T) {
	data := []struct {
		vals []int
		del  int
		str  string
	}{
		// case del leaf
		{[]int{2, 1, 7, 4, 3, 6, 5, 8}, 3, "2174865"},
		// case del subtree with one child
		{[]int{2, 1, 7, 4, 3, 6, 5, 8}, 6, "2174835"},
		// case del subtree with both children
		{[]int{2, 1, 7, 4, 3, 6, 5, 8}, 4, "2175836"},
	}

	for i, d := range data {
		tree := &Tree{val: d.vals[0]}
		for j := 1; j < len(d.vals); j++ {
			Insert(&tree, d.vals[j])
		}
		Delete(&tree, d.del)
		out := tree.ToString()
		if out != d.str {
			t.Errorf("case [%v] expected %v, got %v", i, d.str, out)
		}

	}
}
func TestTreeInsert(t *testing.T) {
	data := []struct {
		vals []int
		str  string
	}{
		{[]int{1, 2, 3, 4}, "1234"},
		{[]int{4, 2, 3, 4}, "4243"},
	}

	for i, d := range data {
		tree := &Tree{val: d.vals[0]}
		for j := 1; j < len(d.vals); j++ {
			Insert(&tree, d.vals[j])
		}
		out := tree.ToString()
		if out != d.str {
			t.Errorf("case [%v] expected %v, got %v", i, d.str, out)
		}

	}
}

func TestTreeFindMax(t *testing.T) {
	data := []struct {
		vals []int
		max  int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{4, 2, 3, 4, 10}, 10},
	}

	for i, d := range data {
		tree := &Tree{val: d.vals[0]}
		for j := 1; j < len(d.vals); j++ {
			Insert(&tree, d.vals[j])
		}
		out := tree.FindMax()
		if out != d.max {
			t.Errorf("case [%v] expected %v, got %v", i, d.max, out)
		}

	}
}
func TestTreeFindMin(t *testing.T) {
	data := []struct {
		vals []int
		min  int
	}{
		{[]int{1, 2, 3, 4}, 1},
		{[]int{4, 2, 3, 4, 10}, 2},
	}

	for i, d := range data {
		tree := &Tree{val: d.vals[0]}
		for j := 1; j < len(d.vals); j++ {
			Insert(&tree, d.vals[j])
		}
		out := tree.FindMin()
		if out != d.min {
			t.Errorf("case [%v] expected %v, got %v", i, d.min, out)
		}

	}
}

func TestTreeIsMirror(t *testing.T) {
	for i, d := range treeMirrorTestData {
		isMirror := d.t1.IsMirror(d.t2)
		if isMirror != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, isMirror)
		}
	}
}

func TestTreeIsMirrorRecur(t *testing.T) {
	for i, d := range treeMirrorTestData {
		isMirror := d.t1.IsMirrorRecur(d.t2)
		if isMirror != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, isMirror)
		}
	}
}

func TestTreeReverse(t *testing.T) {
	data := []struct {
		in  *Tree
		out string
	}{
		{in: nil, out: "empty"},
		{in: &Tree{parent: nil,
			val:  1,
			left: &Tree{val: 2},
			right: &Tree{val: 3,
				left: &Tree{val: 4,
					left:  &Tree{val: 6},
					right: &Tree{val: 5},
				},
				right: &Tree{val: 7,
					left:  &Tree{val: 8},
					right: nil}}}, out: "13274856"},
	}

	for i, d := range data {
		d.in.Reverse()
		if d.in.ToString() != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, d.in.ToString())
		}
	}
}

func TestTreeToString(t *testing.T) {
	data := []struct {
		in  *Tree
		out string
	}{
		{in: nil, out: "empty"},
		{in: &Tree{parent: nil,
			val:  1,
			left: &Tree{val: 2},
			right: &Tree{val: 3,
				left: &Tree{val: 4,
					left:  &Tree{val: 6},
					right: &Tree{val: 5},
				},
				right: &Tree{val: 7,
					left:  &Tree{val: 8},
					right: nil}}}, out: "12347658"},
		{in: &Tree{
			val: 4,
			left: &Tree{
				val:  2,
				left: nil,
				right: &Tree{
					val:   3,
					left:  nil,
					right: &Tree{val: 4},
				},
			},
			right: nil,
		}, out: "4234"},
	}

	for i, d := range data {
		if d.in.ToString() != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, d.in.ToString())
		}
	}
}
