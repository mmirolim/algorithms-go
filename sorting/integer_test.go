package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMapRangeSort(t *testing.T) {
	data := []struct {
		in []int
	}{
		{[]int{2, 1, 5, 8, 0}},
		{[]int{0, 10, 99, 98, 2}},
		{rand.Perm(100)[:20]},
		{[]int{1000, 2, 2, 99, 99, 1, 50}},
		{[]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}},
	}

	for i, d := range data {
		out := mapRange(d.in, nil)
		fmt.Printf("%+v\n", out) // output for debug

		if !isSorted(out) {
			t.Errorf("case [%v] expected sorted got %v", i, out)
		}
	}
}

func TestPerfOfRangeSortVSStdStort(t *testing.T) {
	max := 100000000
	n := 10000000
	start := time.Now()
	randA100 := rand.Perm(max)[:n]
	randB100 := rand.Perm(max)[:n]
	fmt.Printf("Rand arr generation time %+v\n", time.Since(start)) // output for debug

	start = time.Now()
	out := mapRange(randA100, nil)
	fmt.Printf("rangeSort time\n%+v for max %d and n %d\n", time.Since(start), max, n) // output for debug
	if !isSorted(out) {
		t.Errorf("case [%s] expected sorted got %v", "mapRange", out)
	}
	start = time.Now()
	sort.Ints(randB100)
	fmt.Printf("sort.Ints time\n%+v for max %d and n %d\n", time.Since(start), max, n)
}

func isSorted(arr []int) bool {
	prev := arr[0]
	for i := range arr {
		if prev > arr[i] {
			return false
		}
		prev = arr[i]
	}
	return true
}

func BenchmarkMapRangeSortRand100(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand100 := rand.Perm(10000)[:100]
		b.StartTimer()
		_ = mapRange(rand100, nil)
		b.StopTimer()
	}
}

func BenchmarkMapRangeSortStdSortIntsRand100(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand100 := rand.Perm(10000)[:100]
		b.StartTimer()
		sort.Ints(rand100)
		b.StopTimer()
	}
}

func BenchmarkSortRangeInt64K(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		_ = mapRange(data, nil)
		b.StopTimer()
	}
}

func BenchmarkSortInt64K(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<16)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0xcccc
		}
		b.StartTimer()
		sort.Ints(data)
		b.StopTimer()
	}
}
