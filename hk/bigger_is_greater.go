package hk

import (
	"sort"
)

// https://www.hackerrank.com/challenges/bigger-is-greater/problem
func BiggerIsGreater(in []string) (out []string) {
	strs := make([][]byte, len(in))
	for i := range in {
		strs[i] = []byte(in[i])
		from, to := minSwap(strs[i])
		if to == -1 {
			strs[i] = []byte("no answer")
		} else {
			strs[i][from], strs[i][to] = strs[i][to], strs[i][from]
			sortFrom(to, strs[i])
		}
	}

	for i := range strs {
		out = append(out, string(strs[i]))
	}
	return
}

func sortFrom(i int, str []byte) {
	str = str[i+1:]
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
}

func minSwap(str []byte) (int, int) {
	from := len(str) - 1
	to := -1
	for j := from; j >= 0; j-- {
		for i := j; i > to; i-- {
			// use signed integers
			d := int(str[j]) - int(str[i])
			if d > 0 && i > to {
				// check if there's smaller diff swap in range [i, j)
				from = j
				to = i
				break
			}
		}
	}
	return from, to
}
