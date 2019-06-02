package linked_list

import "testing"

func TestFindLLMiddle(t *testing.T) {
	data := []struct {
		LL            *List
		MiddleNodeVal int
	}{
		{&List{head: &Node{1, nil}}, 1},
		{&List{head: &Node{1, &Node{2, &Node{3, nil}}}}, 2},
		{&List{head: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}}, 3},
	}
	for i, d := range data {
		out := FindLLMiddle(d.LL)
		if out.val != d.MiddleNodeVal {
			t.Errorf("case [%v] expected %v, got %v", i, d.MiddleNodeVal, out)
		}
	}
}

func TestListAddOneWithStack(t *testing.T) {
	var input = []struct {
		in  *List
		out string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		}, out: "head->5->6->4->nil"},
		{in: &List{head: &Node{val: 5, next: nil}}, out: "head->6->nil"},
		{in: &List{head: &Node{val: 5, next: &Node{val: 9, next: nil}}}, out: "head->6->0->nil"},
		{in: &List{head: &Node{val: 9, next: &Node{val: 9, next: nil}}}, out: "head->1->0->0->nil"},
	}
	for id, data := range input {
		data.in.AddOneWithStack()
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}

func TestListAddOneRecur(t *testing.T) {
	var input = []struct {
		in  *List
		out string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		}, out: "head->5->6->4->nil"},
		{in: &List{head: &Node{val: 5, next: nil}}, out: "head->6->nil"},
		{in: &List{head: &Node{val: 5, next: &Node{val: 9, next: nil}}}, out: "head->6->0->nil"},
		{in: &List{head: &Node{val: 9, next: &Node{val: 9, next: nil}}}, out: "head->1->0->0->nil"},
	}
	for id, data := range input {
		data.in.AddOneRecur()
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}

func TestListAddOne(t *testing.T) {
	var input = []struct {
		in  *List
		out string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		}, out: "head->5->6->4->nil"},
		{in: &List{head: &Node{val: 5, next: nil}}, out: "head->6->nil"},
		{in: &List{head: &Node{val: 5, next: &Node{val: 9, next: nil}}}, out: "head->6->0->nil"},
		{in: &List{head: &Node{val: 9, next: &Node{val: 9, next: nil}}}, out: "head->1->0->0->nil"},
	}
	for id, data := range input {
		data.in.AddOne()
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}

func TestListReverse(t *testing.T) {
	var input = []struct {
		in  *List
		out string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		}, out: "head->3->6->5->nil"},
		{in: &List{head: &Node{val: 5, next: nil}}, out: "head->5->nil"},
		{in: NewList(), out: "head->nil"},
	}
	for id, data := range input {
		data.in.Reverse()
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}

func TestListReverseRecur(t *testing.T) {
	var input = []struct {
		in  *List
		out string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		}, out: "head->3->6->5->nil"},
		{in: &List{head: &Node{val: 5, next: nil}}, out: "head->5->nil"},
		{in: NewList(), out: "head->nil"},
	}
	for id, data := range input {
		data.in.ReverseRecur()
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}

func TestListToString(t *testing.T) {
	var input = []struct {
		in  *List
		out string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		}, out: "head->5->6->3->nil"},
		{in: &List{head: &Node{val: 5, next: nil}}, out: "head->5->nil"},
		{in: NewList(), out: "head->nil"},
	}
	for id, data := range input {
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}

func TestListAppend(t *testing.T) {
	var input = []struct {
		in    *List
		nodes []*Node
		out   string
	}{
		{in: &List{
			head: &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}},
			len:  3,
		},
			nodes: []*Node{&Node{100, nil}},
			out:   "head->5->6->3->100->nil"},
		{in: &List{head: &Node{val: 5, next: nil}},
			nodes: []*Node{&Node{100, nil}, &Node{101, nil}},
			out:   "head->5->100->101->nil"},

		{in: NewList(), nodes: []*Node{&Node{100, nil}}, out: "head->100->nil"},
	}
	for id, data := range input {
		for _, n := range data.nodes {
			data.in.Append(n)
		}
		if data.in.ToString() != data.out {
			t.Errorf("case [%v] expected %v, got %v", id, data.out, data.in.ToString())
		}
	}
}
