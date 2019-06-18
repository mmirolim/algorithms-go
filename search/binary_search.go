package search

func BinarySearch(in []int, v int) int {
	var bs func(int, int) int
	bs = func(l, h int) int {
		if l > h {
			return -1
		}

		middle := (l + h) / 2
		if in[middle] == v {
			return middle
		}
		if in[middle] < v {
			return bs(middle+1, h)
		}
		return bs(l, middle-1)
	}

	return bs(0, len(in)-1)
}

// ADM problem, count number of V elements in sorted array
func CountKElemInSortedSet(in []int, v int) int {
	var bs func(lb bool, l, h int) int
	bs = func(lb bool, l, h int) int {
		if l > h {
			return -1
		}
		middle := (l + h) / 2
		if in[middle] == v {
			if middle == 0 || middle == len(in)-1 {
				return middle
			}
		}

		if lb {
			if in[middle-1] < v && in[middle] == v {
				return middle
			}
		} else {
			if in[middle-1] == v && in[middle] > v {
				return middle - 1
			}
		}
		// inside the k set
		if in[middle] == v {
			if lb {
				return bs(lb, l, middle-1)
			}
			return bs(lb, middle+1, h)
		}

		if in[middle] < v {
			return bs(lb, middle+1, h)
		}
		return bs(lb, l, middle-1)
	}

	start := bs(true, 0, len(in)-1)
	if start == -1 {
		return 0
	}
	end := bs(false, 0, len(in)-1)

	return end - start + 1
}

func CountKElemInSortedSetLinear(in []int, v int) int {
	count := 0
	for i := range in {
		if in[i] == v {
			count++
		}
	}
	return count
}

// ADM sorting 4.10
// in slice should be sorted
func CheckSumOfKEqToT(in []int, k, T int) bool {
	if k == 1 {
		return BinarySearch(in, T) > -1
	} else if k == 2 {
		for i := range in {
			if BinarySearch(in, T-in[i]) > -1 {
				return true
			}
		}
		return false
	}
	// k >= 3
	for i := range in {
		if CheckSumOfKEqToT(append(in[0:i], in[i+1:]...), k-1, T-in[i]) {
			return true
		}
	}
	return false
}
