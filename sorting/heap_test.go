package sorting

import "testing"

func TestHeapSort(t *testing.T) {
	data := []struct {
		heap func(*[]int)
		in   []int
		out  []int
	}{
		{HeapSortMin, []int{2, 1, 5, 8, -1}, []int{-1, 1, 2, 5, 8}},
		{HeapSortMax, []int{2, 1, 5, 8, -1}, []int{8, 5, 2, 1, -1}},
	}

	for i, d := range data {
		d.heap(&d.in)
		for j := range d.out {
			if d.in[j] != d.out[j] {
				t.Errorf("case [%v] expected %v, got %v", i, d.out, d.in)
			}
		}
	}
}
