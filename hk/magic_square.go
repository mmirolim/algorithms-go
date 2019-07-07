package hk

import (
	"math"
)

// https://www.hackerrank.com/challenges/magic-square-forming/problem
func FormingMagicSquare(n [][]int) (cost int, res [][]int) {
	mss := generateAllMagicSquares(len(n))
	minCost := 1000
	var minSq [][]int
	for _, sq := range mss {
		cost := computeCost(n, sq)
		if cost < minCost {
			minCost = cost
			minSq = sq
		}
	}
	return minCost, minSq
}

func computeCost(n [][]int, msq [][]int) (cost int) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			cost += int(math.Abs(float64(n[i][j] - msq[i][j])))
		}
	}
	return
}

func vSum(j int, n [][]int) int {
	return n[0][j] + n[1][j] + n[2][j]
}
func hSum(i int, n [][]int) int {
	return n[i][0] + n[i][1] + n[i][2]
}
func dSum1(n [][]int) int {
	return n[0][0] + n[1][1] + n[2][2]
}
func dSum2(n [][]int) int {
	return n[0][2] + n[1][1] + n[2][0]
}

func conv(d int, n []int) [][]int {
	out := make([][]int, d)
	if len(n) != d*d {
		panic("should be square length")
	}
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			out[i] = append(out[i], n[i*d+j])
		}
	}
	return out
}

func generateAllMagicSquares(n int) [][][]int {
	isMagicSquare := func(n [][]int) bool {
		if dSum1(n) != 15 || dSum2(n) != 15 {
			return false
		}

		for i := 0; i < len(n); i++ {
			if hSum(i, n) != 15 {
				return false
			}
		}

		for j := 0; j < len(n); j++ {
			if vSum(j, n) != 15 {
				return false
			}

		}
		return true
	}

	var out [][][]int

	var recur func(start int, nums []int)
	recur = func(start int, s []int) {
		if start == len(s) {
			grid := conv(n, s)
			if isMagicSquare(grid) {
				out = append(out, grid)
			}
			return
		}
		for i := start; i < n*n; i++ {
			s[start], s[i] = s[i], s[start]
			recur(start+1, s)
			s[start], s[i] = s[i], s[start]
		}
	}

	var nums []int
	for i := 0; i < n*n; i++ {
		nums = append(nums, i+1)
	}

	recur(0, nums)

	return out
}

func Permutate(n int) [][]int {
	var out [][]int
	var recur func(start int, nums []int)
	recur = func(start int, s []int) {
		if start == len(s) {
			res := make([]int, len(s))
			copy(res, s)
			out = append(out, res)
			return
		}
		for i := start; i < n; i++ {
			s[start], s[i] = s[i], s[start]
			recur(start+1, s)
			s[start], s[i] = s[i], s[start]
		}
	}

	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}

	recur(0, nums)

	return out

}

func PermutateHeapsAlgorithm(n int) [][]int {
	var out [][]int
	var recur func(k int, s []int)
	recur = func(k int, s []int) {
		if k == 1 {
			res := make([]int, len(s))
			copy(res, s)
			out = append(out, res)
			return
		}
		// gen with kth elem unaltered
		recur(k-1, s)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				s[i], s[k-1] = s[k-1], s[i]
			} else {
				s[0], s[k-1] = s[k-1], s[0]
			}
			recur(k-1, s)
		}
	}

	var nums []int
	for i := 0; i < n; i++ {
		nums = append(nums, i+1)
	}

	recur(len(nums), nums)

	return out

}
