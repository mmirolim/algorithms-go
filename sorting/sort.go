package sorting

import (
	"math/rand"
)

func merge(in1, in2 []int) []int {
	out := make([]int, len(in1)+len(in2))
	i, j, k := 0, 0, 0
	for ; j < len(in1) && k < len(in2); i++ {
		if in1[j] < in2[k] {
			out[i] = in1[j]
			j++
		} else {
			out[i] = in2[k]
			k++
		}
	}
	for ; j < len(in1); j++ {
		out[i] = in1[j]
		i++
	}
	for ; k < len(in2); k++ {
		out[i] = in2[k]
		i++
	}

	return out
}

func MergeSort(in []int) []int {
	if len(in) == 1 {
		return in
	}
	return merge(
		MergeSort(in[:len(in)/2]),
		MergeSort(in[len(in)/2:]))
}

func qspartition(in *[]int) int {
	sl := *in
	p := rand.Intn(len(sl))
	last := len(sl) - 1
	sl[last], sl[p] = sl[p], sl[last]
	count := 0
	for i := 0; i < len(sl); i++ {
		if sl[i] < sl[last] {
			sl[i], sl[count] = sl[count], sl[i]
			count++
		}
	}

	sl[last], sl[count] = sl[count], sl[last]
	// return border index
	return count
}

func QuickSort(in *[]int) {
	if len(*in) > 0 {
		p := qspartition(in)
		if p > 0 {
			p1 := (*in)[:p-1]
			QuickSort(&p1)
		}
		p2 := (*in)[p+1:]
		QuickSort(&p2)
	}
}
