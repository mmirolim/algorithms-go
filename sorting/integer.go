package sorting

import "sync"

func mapRangeParams(arr []int) (a1, da, size int, sorted bool) {
	min := arr[0]
	max := arr[0]
	sorted = true
	prev := arr[0]
	for i := range arr {
		if sorted && (prev > arr[i]) {
			sorted = false
		}
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
		prev = arr[i]
	}
	size = len(arr) - 1
	a1 = min
	da = max - min
	return
}

type syncpool struct{ sync.Pool }

func (s *syncpool) Alloc(n int) [][]int {
	if b, _ := s.Pool.Get().([][]int); cap(b) >= n {
		return b[:n]
	}
	return make([][]int, n) // pool allocation mis-sized
}
func (s *syncpool) Free(arr [][]int) {
	for i := range arr {
		arr[i] = arr[i][:0]
	}
	s.Pool.Put(arr)
}

var pool = syncpool{}

func mapRange(arr []int, out []int) []int {
	if len(arr) == 2 {
		if arr[0] <= arr[1] {
			return arr
		}
		// swap
		arr[0], arr[1] = arr[1], arr[0]
		return arr
	}
	a1, da, size, sorted := mapRangeParams(arr)
	if sorted {
		return arr
	}
	if out == nil {
		out = make([]int, len(arr))
	} else {
		out = out[:len(arr)]
	}
	t := 0
	for i := range arr {
		t = (arr[i] - a1) * size / da
		out[i] = t
	}
	sorted = true
	prev := out[0]
	for _, d := range out {
		if prev >= d {
			sorted = false
			break
		}
		prev = d
	}
	buckets := pool.Alloc(len(arr))
	for i := range out {
		buckets[out[i]] = append(buckets[out[i]], arr[i])
	}
	arr = arr[:0]
	for i := range buckets {
		if len(buckets[i]) == 1 {
			arr = append(arr, buckets[i][0])
		} else if len(buckets[i]) > 1 {
			arr = append(arr, mapRange(buckets[i], out)...)
		}
	}
	pool.Free(buckets)
	return arr
}

func SortIntegers(arr []int) []int {
	return arr
}

func isNaN(a float64) bool {
	return a != a
}
