package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	data := []struct {
		in, out []int
	}{
		{[]int{2, 3, 10, -1, 23, 6, 7}, []int{-1, 2, 3, 6, 7, 10, 23}},
		{[]int{0, 0, 12, 3, -1}, []int{-1, 0, 0, 3, 12}},
	}

	for i, d := range data {
		out := MergeSort(d.in)
		for j := range d.out {
			if out[j] != d.out[j] {
				t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
				t.FailNow()
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	data := []struct {
		in, out []int
	}{
		{[]int{2, 3, 10, -1, 23, 6, 7}, []int{-1, 2, 3, 6, 7, 10, 23}},
		{[]int{0, 0, 12, 3, -1}, []int{-1, 0, 0, 3, 12}},
	}

	for i, d := range data {
		QuickSort(&d.in)
		for j := range d.out {
			if d.in[j] != d.out[j] {
				t.Errorf("case [%v] expected %v, got %v", i, d.out, d.in)
				t.FailNow()
			}
		}
	}
}
