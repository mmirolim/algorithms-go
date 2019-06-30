package tree

import (
	"testing"

	"github.com/mmirolim/algos/checks"
)

func TestNewHeap(t *testing.T) {
	data := []struct {
		size                 int
		storageSize          int
		err                  error
		a, b                 int
		toSwapMin, toSwapMax bool
	}{
		// unlimited heap
		{-1, 0, nil, 10, 11, false, true},
		{-2, 0, ErrHeapWrongSize, 10, 11, false, true},
		{0, 0, ErrHeapWrongSize, 10, 11, false, true},
		{10, 10, nil, 12, 9, true, false},
		{2048, 2048, nil, -10, -9, false, true},
	}
	for i, d := range data {
		h, err := NewMinHeap(d.size)
		if checks.NeqErrs(d.err, err) {
			t.Errorf("case [%v] expected err %v, got %v", i, d.err, err)
			continue
		}
		if err == nil && d.storageSize != cap(h.arr) {
			t.Errorf("case [%v] expected %v, got %v", i, d.storageSize, cap(h.arr))
		}
		if err == nil && h.toSwap(d.a, d.b) != d.toSwapMin {
			t.Errorf("case [%v] expected a %v, b %v toSwap %v, got %v", i, d.a, d.b, d.toSwapMin, h.toSwap(d.a, d.b))
		}

		h, err = NewMaxHeap(d.size)
		if checks.NeqErrs(d.err, err) {
			t.Errorf("case [%v] expected err %v, got %v", i, d.err, err)
			continue
		}
		if err == nil && d.storageSize != cap(h.arr) {
			t.Errorf("case [%v] expected %v, got %v", i, d.storageSize, cap(h.arr))
		}
		if err == nil && h.toSwap(d.a, d.b) != d.toSwapMax {
			t.Errorf("case [%v] expected a %v, b %v toSwap %v, got %v", i, d.a, d.b, d.toSwapMax, h.toSwap(d.a, d.b))
		}

	}
}

func TestHeapInsert(t *testing.T) {
	data := []struct {
		heap             func(int) (*Heap, error)
		sl               []int
		expectedFirstElm []int
	}{
		{NewMaxHeap, []int{2, 1, 5, 8, -1}, []int{2, 2, 5, 8, 8}},
		{NewMinHeap, []int{2, 1, 5, 8, -1}, []int{2, 1, 1, 1, -1}},
	}

	for i, d := range data {
		h, err := d.heap(-1)
		if err != nil {
			t.Errorf("case [%v] unexpected error %#v", i, err)
			continue
		}
		for j := range d.sl {
			err := h.Insert(d.sl[j], j)
			if err != nil {
				t.Errorf("case [%v, %v] unexpected error %#v", i, j, err)
				continue
			}
			if h.arr[0].key != d.expectedFirstElm[j] {
				t.Errorf("case [%v, %v] expected %v, got %v", i, j, d.expectedFirstElm[j], h.arr[0])
			}
		}
	}

}

func TestHeapInsertErr(t *testing.T) {
	for i, heap := range []func(int) (*Heap, error){NewMaxHeap, NewMinHeap} {
		maxSize := 5
		h, err := heap(maxSize)
		if err != nil {
			t.Errorf("case [%v] unexpected error %#v", i, err)
			continue
		}
		for i := 0; i < maxSize; i++ {
			err := h.Insert(i, i)
			if err != nil {
				t.Errorf("case [%v] unexpected error %#v", i, err)
				t.FailNow()
			}
		}
		// additinal insert
		err = h.Insert(1, 1)
		if err == nil || err.Error() != ErrHeapIsFull.Error() {
			t.Errorf("case [%v] expected error %+v, got %+v", i, ErrHeapIsFull, err)
		}

	}
}

func TestHeapExtractTop(t *testing.T) {
	input := []int{2, 1, 5, 8, 10, -1}
	expected := []int{-1, 1, 2, 5, 8, 10}
	h, _ := NewMinHeap(20)
	for _, v := range input {
		err := h.Insert(v, v)
		if err != nil {
			t.Errorf("unexpected err %+v", err)
			t.FailNow()
		}
	}

	for _, v := range expected {
		out, val, err := h.ExtractTop()
		if err != nil {
			t.Errorf("unexpected err %+v", err)
			t.FailNow()
		}
		if out != v || val.(int) != v {
			t.Errorf("expected key %v val %v, got %v %v", v, v, out, val)
		}
	}

	// try extract from empty heap
	_, _, err := h.ExtractTop()
	if err == nil || err.Error() != ErrHeapIsEmpty.Error() {
		t.Errorf("expected err %+v, got %+v", ErrHeapIsEmpty, err)
	}

}

func TestHeapToString(t *testing.T) {
	data := []struct {
		heap  func(int) (*Heap, error)
		input []int
		str   string
	}{
		{NewMinHeap, []int{10, -1, 2, 5, -2, 7, 8, 3, 11}, "-2,-1,2,3,10,7,8,5,11"},
		{NewMaxHeap, []int{10, -1, 2, 5, -2}, "10,5,-1,2,-2"},
	}

	for i, d := range data {
		h, err := d.heap(-1)
		if err != nil {
			t.Errorf("case [%v] unexpected err %+v", i, err)
			t.FailNow()
		}
		for _, v := range d.input {
			err := h.Insert(v, v)
			if err != nil {
				t.Errorf("case [%v] unexpected err %+v", i, err)
				t.FailNow()
			}
		}
		out := h.ToString()
		if out != d.str {
			t.Errorf("case [%v] expected %v, got %v", i, d.str, out)
		}
	}
}
