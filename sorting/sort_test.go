package sorting

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := []struct {
		in, out []int
	}{
		{[]int{2, 3, 10, -1, 23, 6, 7}, []int{-1, 2, 3, 6, 7, 10, 23}},
		{[]int{0, 0, 12, 3, -1}, []int{-1, 0, 0, 3, 12}},
		{[]int{10, 80, 30, 90, 40, 50, 70}, []int{10, 30, 40, 50, 70, 80, 90}},
		{[]int{4, 3, 5, 2, 1, 3, 2, 3}, []int{1, 2, 2, 3, 3, 3, 4, 5}},
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
		{[]int{10, 80, 30, 90, 40, 50, 70}, []int{10, 30, 40, 50, 70, 80, 90}},
		{[]int{4, 3, 5, 2, 1, 3, 2, 3}, []int{1, 2, 2, 3, 3, 3, 4, 5}},
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
func TestQuickSortConcurrent(t *testing.T) {
	data := []struct {
		in, out []int
	}{
		{[]int{2, 3, 10, -1, 23, 6, 7}, []int{-1, 2, 3, 6, 7, 10, 23}},
		{[]int{0, 0, 12, 3, -1}, []int{-1, 0, 0, 3, 12}},
		{[]int{10, 80, 30, 90, 40, 50, 70}, []int{10, 30, 40, 50, 70, 80, 90}},
		{[]int{4, 3, 5, 2, 1, 3, 2, 3}, []int{1, 2, 2, 3, 3, 3, 4, 5}},
	}

	for i, d := range data {
		out := QuickSortConcurrent(d.in)
		for j := range d.out {
			if out[j] != d.out[j] {
				t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
				t.FailNow()
			}
		}
	}
}

func TestFindMaxWallIntersectAimPoint(t *testing.T) {
	data := []struct {
		points []Point
		q, aim Point
	}{
		{[]Point{Point{1, 2.0, 1.0}, Point{2, 3.0, 2.0}, Point{3, 4.0, 1.0}, Point{4, 4.0, 3.0}, Point{5, 2.0, 3.0}}, Point{0, 1.0, 1.0}, Point{1, 2.0, 1.0}},
		{[]Point{Point{1, 2.0, 1.0}, Point{2, 4.0, 1.0}, Point{3, 4.0, 3.0}, Point{4, 3.0, 2.0}, Point{5, 2.0, 3.0}}, Point{0, 1.0, 1.0}, Point{4, 3.0, 2.0}},
		{[]Point{Point{1, 2.0, 1.0}, Point{2, 4.0, 1.0}, Point{3, 4.0, 3.0}, Point{4, 3.0, 3.0}, Point{5, 3.0, 2.0}, Point{6, 2.0, 2.0}}, Point{0, 1.0, 1.0}, Point{3, 4.0, 3.0}},
	}
	for i, d := range data {
		aimPoint := FindMaxWallIntersectAimPoint(d.points, d.q)
		if aimPoint.id != d.aim.id {
			t.Errorf("case [%v] expected %+v, got %+v", i, d.aim, aimPoint)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := rand.Perm(1e7)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		QuickSort(&arr)
	}
}

func BenchmarkQuickSortConcurrent(b *testing.B) {
	arr := rand.Perm(1e7)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = QuickSortConcurrent(arr)
	}
}
