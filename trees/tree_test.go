package tree

import "testing"

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
