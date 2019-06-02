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

func TestADMProblem3_13_ToString(t *testing.T) {
	data := []struct {
		in  []int
		out string
	}{
		{[]int{}, "(id=0,val=0,sum=0)"},
		{[]int{1, 2, 3}, "(id=1,val=1,sum=6)(id=2,val=2,sum=5)(id=3,val=3,sum=3)"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "(id=5,val=5,sum=55)(id=2,val=2,sum=10)(id=7,val=7,sum=40)(id=1,val=1,sum=1)(id=3,val=3,sum=7)(id=6,val=6,sum=6)(id=8,val=8,sum=27)(id=4,val=4,sum=4)(id=9,val=9,sum=19)(id=10,val=10,sum=10)"},
	}
	for i, d := range data {
		res := ADMProblem3_13(d.in)
		out := res.t.ToString()
		if out != d.out {
			t.Errorf("case [%d] expected %v, got %v", i, d.out, out)
		}

	}
}
func TestADMProblem3_13_PartialSum(t *testing.T) {
	data := []struct {
		in     []int
		i, out int
	}{
		{[]int{1, 2, 3}, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 5, 15},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 28},
	}
	for i, d := range data {
		res := ADMProblem3_13(d.in)
		out := res.PartialSum(d.i)
		if out != d.out {
			t.Errorf("case [%d] expected %v, got %v", i, d.out, out)
		}

	}
}

func TestADMProblem3_13_Add(t *testing.T) {
	data := []struct {
		in           []int
		i, y, x, out int
	}{
		{[]int{1, 2, 3}, 1, 10, 1, 11},
		{[]int{1, 2, 3, 4, 5}, 5, 7, 3, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 5, 8, 41},
	}
	for i, d := range data {
		res := ADMProblem3_13(d.in)
		res.Add(d.i, d.y)
		out := res.PartialSum(d.x)
		if out != d.out {
			t.Errorf("case [%d] expected %v, got %v", i, d.out, out)
		}

	}

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

func TestTreesEqual(t *testing.T) {
	data := []struct {
		t1, t2 *Tree
		out    bool
	}{
		{nil, nil, true},
		{t1: &Tree{
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
			right: nil},
			t2: &Tree{
				val: 4,
				left: &Tree{
					val:  2,
					left: nil,
					right: &Tree{
						val:   30,
						left:  nil,
						right: &Tree{val: 4},
					},
				},
				right: nil}, out: false},
		{t1: &Tree{
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
		}, t2: &Tree{
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
			right: nil}, out: true},
	}
	for i, d := range data {
		out1 := d.t1.EqualLevelOrder(d.t2)
		out2 := d.t1.EqualInOrder(d.t2)
		if out1 != d.out || out1 != out2 {
			t.Errorf("case [%v] expected %v, got out1 %v out2 %v", i, d.out, out1, out2)
		}
	}

}
