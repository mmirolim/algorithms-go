package tree

import (
	"testing"
)

var treeMirrorTestData = []struct {
	t1, t2 *Tree
	out    bool
}{
	{t1: &Tree{root: nil}, t2: &Tree{root: nil}, out: true},
	{t1: &Tree{root: &Node{val: 1,
		left: &Node{val: 2,
			left:  nil,
			right: nil},
		right: &Node{val: 3,
			left:  nil,
			right: nil}}},
		t2: &Tree{root: &Node{val: 1,
			right: &Node{val: 2,
				left:  nil,
				right: nil},
			left: &Node{val: 3,
				left: &Node{val: 5,
					left:  nil,
					right: nil},
				right: nil}}}, out: false},
	{t1: &Tree{root: &Node{val: 1,
		left: &Node{val: 2,
			left:  nil,
			right: nil},
		right: &Node{val: 3,
			left: &Node{val: 4,
				left:  &Node{val: 6, left: nil, right: nil},
				right: &Node{val: 5, left: nil, right: nil},
			},
			right: &Node{val: 7,
				left:  &Node{val: 8, left: nil, right: nil},
				right: nil}}}},
		t2: &Tree{root: &Node{val: 1,
			right: &Node{val: 2,
				left:  nil,
				right: nil},
			left: &Node{val: 3,
				right: &Node{val: 4,
					right: &Node{val: 6, left: nil, right: nil},
					left:  &Node{val: 5, left: nil, right: nil},
				},
				left: &Node{val: 7,
					right: &Node{val: 8, left: nil, right: nil},
					left:  nil}}}},
		out: true},
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
		{in: &Tree{root: &Node{val: 1,
			left: &Node{val: 2,
				left:  nil,
				right: nil},
			right: &Node{val: 3,
				left: &Node{val: 4,
					left:  &Node{val: 6, left: nil, right: nil},
					right: &Node{val: 5, left: nil, right: nil},
				},
				right: &Node{val: 7,
					left:  &Node{val: 8, left: nil, right: nil},
					right: nil}}}}, out: "13274856"},
	}

	for i, d := range data {
		d.in.Reverse()
		if d.in.String() != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, d.in.String())
		}
	}
}

func TestTreeString(t *testing.T) {
	data := []struct {
		in  *Tree
		out string
	}{
		{in: nil, out: "empty"},
		{in: &Tree{root: &Node{val: 1,
			left: &Node{val: 2,
				left:  nil,
				right: nil},
			right: &Node{val: 3,
				left: &Node{val: 4,
					left:  &Node{val: 6, left: nil, right: nil},
					right: &Node{val: 5, left: nil, right: nil},
				},
				right: &Node{val: 7,
					left:  &Node{val: 8, left: nil, right: nil},
					right: nil}}}}, out: "12347658"},
	}

	for i, d := range data {
		if d.in.String() != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, d.in.String())
		}
	}
}
