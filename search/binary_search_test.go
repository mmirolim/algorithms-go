package search

import (
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []struct {
		in    []int
		v, id int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 10, -1},
	}

	for i, d := range data {
		out := BinarySearch(d.in, d.v)
		if out != d.id {
			t.Errorf("case [%v] expected %v, got %v", i, d.id, out)
		}
	}
}

func TestCountKElemInSortedSet(t *testing.T) {
	data := []struct {
		in   []int
		v, c int
	}{
		{[]int{1, 2, 3, 4, 4, 4, 4, 5, 5, 6, 7, 8}, 4, 4},
		{[]int{1, 1, 1, 1, 4, 4, 4, 5, 5, 6, 7, 8}, 1, 4},
		{[]int{1, 1, 1, 1, 4, 4, 4, 5, 5, 6, 7, 8, 8}, 8, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 10, 0},
		{[]int{1, 1, 1, 1, 1, 1}, 1, 6},
	}

	for i, d := range data {
		out := CountKElemInSortedSet(d.in, d.v)
		if out != d.c {
			t.Errorf("case [%v] expected %v, got %v", i, d.c, out)
		}
		out2 := CountKElemInSortedSetLinear(d.in, d.v)
		if out != out2 {
			t.Errorf("case [%v] expected f1==f2 %v, got %v", i, out, out2)
		}
	}

}
func BenchmarkCountKElemInSortedSet(b *testing.B) {
	_, out := generateSortedSetWithKNnums(1000, 10, 30)
	for i := 0; i < b.N; i++ {
		_ = CountKElemInSortedSet(out, 10)
	}
}

func BenchmarkCountKElemInSortedSetLinear(b *testing.B) {
	_, out := generateSortedSetWithKNnums(1000, 10, 30)
	for i := 0; i < b.N; i++ {
		_ = CountKElemInSortedSetLinear(out, 10)
	}

}

func generateSortedSetWithKNnums(l, k, n int) (int, []int) {
	out := make([]int, l)
	for i := range out {
		out[i] = i
	}

	index := rand.Intn(l)
	if index+n > l {
		n = l - index
	}
	for i := index; i < index+n; i++ {
		out[i] = k
	}

	return index, out
}
