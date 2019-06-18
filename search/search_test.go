package search

import (
	"testing"
)

// ADM Sorting 4.10
func TestCheckSumOfKEqToT(t *testing.T) {
	data := []struct {
		in   []int
		k, T int
		out  bool
	}{
		{[]int{5, 7, 8, 9, 10, 12, 14, 15, 16, 17, 18, 19, 20}, 1, 15, true},
		{[]int{5, 7, 9, 10, 12, 14, 15, 16, 17, 18, 19, 20}, 2, 24, true},
		{[]int{7, 8, 10, 12, 14, 17, 18, 19, 20}, 2, 23, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 3, 18, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 3, 31, true},
		{[]int{5, 7, 8, 9, 10, 12, 14, 15, 16, 17, 18, 19, 20}, 3, 33, true},
	}
	for i, d := range data {
		ok := CheckSumOfKEqToT(d.in, d.k, d.T)
		if ok != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, ok)
		}
	}
}
