package sorting

import tree "github.com/mmirolim/algos/trees"

func HeapSortMin(in *[]int) {
	if len(*in) < 2 {
		return
	}
	h, _ := tree.NewMinHeapFrom(*in, nil)
	for i := 0; i < len(*in); i++ {
		n, _, _ := h.ExtractTop()
		(*in)[i] = n
	}
}

func HeapSortMax(in *[]int) {
	if len(*in) < 2 {
		return
	}
	h, _ := tree.NewMaxHeapFrom(*in, nil)
	for i := 0; i < len(*in); i++ {
		n, _, _ := h.ExtractTop()
		(*in)[i] = n
	}
}
