package search

import (
	"fmt"
	"math/rand"
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
func TestFindKthMinElementLinear(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	k := 5
	res := 5
	out := FindKthMinElementLinear(in, k)
	if out != res {
		panic(fmt.Sprintf("expected %v, got %v", res, out))
	}
}

func TestFindKthElement(t *testing.T) {
	data := []struct {
		in  []int
		k   int
		res int
		cmp func(a, b int) bool
	}{
		{[]int{3, 6, 100, 9, 10, 12, 7, -1, 10}, 2, 3, func(a, b int) bool { return a < b }},
		{[]int{3, 6, 100, 9, 10, 12, 7, -1, 10}, 2, 12, func(a, b int) bool { return a > b }},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 5, 14, func(a, b int) bool { return a > b }},
	}
	for i, d := range data {
		kth := FindKthElement(d.in, d.k, d.cmp)
		if kth != d.res {
			t.Errorf("case [%v] expected %v, got %v", i, d.res, kth)
		}
	}
}

func TestFind2thElementTA(t *testing.T) {
	data := []struct {
		in  []int
		res int
		cmp func(a, b int) bool
	}{
		{[]int{3, 6, 100, 9, 10, 12, 7, -1, 10}, 3, func(a, b int) bool { return a < b }},
		{[]int{3, 6, 100, 9, 10, 12, 7, -1, 10}, 12, func(a, b int) bool { return a > b }},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}, 17, func(a, b int) bool { return a > b }},
	}
	for i, d := range data {
		kth := Find2thElementTA(d.in, d.cmp)
		if kth != d.res {
			t.Errorf("case [%v] expected %v, got %v", i, d.res, kth)
		}
	}
}

func BenchmarkFind2thElementTA(b *testing.B) {
	in := rand.Perm(1000)
	cmp := func(a, b int) bool {
		return a < b
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Find2thElementTA(in, cmp)
	}
}

func BenchmarkFindKthElement(b *testing.B) {
	in := rand.Perm(1000)
	cmp := func(a, b int) bool {
		return a < b
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindKthElement(in, 5, cmp)
	}
}

func BenchmarkFindKthElementHeapifyAll(b *testing.B) {
	in := rand.Perm(1000)
	cmp := func(a, b int) bool {
		return a < b
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindKthElementHeapifyAll(in, 5, cmp)
	}
}

func BenchmarkFindKthElementLinear(b *testing.B) {
	in := rand.Perm(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindKthMinElementLinear(in, 5)
	}
}
