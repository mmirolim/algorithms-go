package hk

import (
	"strconv"
)

// https://www.hackerrank.com/challenges/almost-sorted/problem
func AlmostSorted(arr []int32) string {
	var sets [][]int
	found := false
	almostSortedTest := func(l, r int, arr []int32) bool {
		if arr[l] >= arr[r-1] {
			if r+1 < len(arr) {
				if arr[l] > arr[r+1] {
					return false
				}
			}
			if arr[r] <= arr[l+1] {
				if l-1 >= 0 {
					if arr[r] < arr[l-1] {
						return false
					}
				}
			}
			return true
		}
		return false
	}

	// ascending order check
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			if !found {
				found = true
				sets = append(sets, []int{i - 1})
			}
		} else {
			if found {
				sets[len(sets)-1] = append(sets[len(sets)-1], i-1)
				found = false
			}
		}
	}

	l, r, res := 0, 0, ""
	if len(sets) == 1 {
		if len(sets[0]) == 1 && len(arr) == 2 {
			l = sets[0][0]
			r = sets[0][0] + 1
			res = "swap"
		} else if len(sets[0]) == 1 {
			l = sets[0][0]
			r = len(arr) - 1
			res = "reverse"
		} else {
			l = sets[0][0]
			r = sets[0][1]
			if r-l == 1 {
				res = "swap"
			} else {
				res = "reverse"
			}
		}

	} else if len(sets) == 2 {
		// test swap
		l = sets[0][0]
		r = sets[1][0] + 1
		res = "swap"
	} else {
		return ""
	}

	if almostSortedTest(l, r, arr) {
		return res + " " + strconv.Itoa(l+1) + " " + strconv.Itoa(r+1)
	}
	return ""
}
