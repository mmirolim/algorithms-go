package strings

import (
	"math"
)

func CountNumOfBinaryStringsNWithNoConsec1s(n int) int {
	var countSeq func(index, count int) int
	countSeq = func(index, count int) int {
		if count >= 2 {
			return 0
		}
		if index == 0 {
			if count < 2 {
				return 1
			}
			return 0
		}
		return countSeq(index-1, count+1) + countSeq(index-1, 0)
	}

	return countSeq(n-1, 1) + countSeq(n-1, 0)
}

func CountNumOfBinaryStringsNWithAtLeastKAdjacent1Bits(n, k int) int {
	// dp solution
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = -1
		}
	}

	// base condition
	for i := 0; i < n; i++ {
		dp[i][k] = int(1 << uint(i+1))
	}

	var countSeq func(index, count int) int
	countSeq = func(index, count int) int {
		if index < 0 {
			if count == 3 {
				return 1
			}
			return 0
		}

		if dp[index][count] > -1 {
			return dp[index][count]
		}

		dp[index][count] = countSeq(index-1, count+1)
		dp[index][count] += countSeq(index-1, 0)

		return dp[index][count]
	}
	return countSeq(n-1, 0)
}

func NumOfBinaryStringsNWithAtLeastKAdjacent1Bits(n, k int) []string {
	out := []string{}
	var genSeq func(index, count, sets int, str []byte)
	genSeq = func(index, count, sets int, str []byte) {
		if count >= k {
			sets += count
		}
		if index == n {
			if sets >= k {
				out = append(out, string(str))
			}
			return
		}

		str[index] = '1'
		genSeq(index+1, count+1, sets, str)
		str[index] = '0'
		genSeq(index+1, 0, sets, str)
	}

	str := make([]byte, n)
	str[0] = '1'
	genSeq(1, 1, 0, str)
	str[0] = '0'
	genSeq(1, 0, 0, str)

	return out
}

func NumOfBinaryStringsWithSameSumOfFirstAndSecondHalf(n int) []string {
	out := []string{}
	var genSeq func(diff, start, end int, seq []byte)
	genSeq = func(diff, start, end int, seq []byte) {
		if int(math.Abs(float64(diff))) > (end-start+1)/2 {
			return
		}
		if start > end {
			if diff == 0 {
				out = append(out, string(seq))
			}
			return
		}

		seq[start], seq[end] = '1', '1'
		genSeq(diff, start+1, end-1, seq)

		seq[start], seq[end] = '1', '0'
		genSeq(diff+1, start+1, end-1, seq)

		seq[start], seq[end] = '0', '0'
		genSeq(diff, start+1, end-1, seq)

		seq[start], seq[end] = '0', '1'
		genSeq(diff-1, start+1, end-1, seq)

	}

	diff, start, end := 0, 0, 2*n-1
	seq := make([]byte, 2*n)
	genSeq(diff, start, end, seq)

	return out
}

func NumOfBinaryStringsNWithWithKAdjacentSetBits(n, k int) int {
	var genSeq func(n, k, count, index, lastBit int) int
	genSeq = func(n, k, count, index, lastBit int) int {
		if index == n && count == k {
			return 1
		} else if index == n || count > k {
			return 0
		}

		res := 0
		if lastBit == 1 {
			res = genSeq(n, k, count+1, index+1, 1) +
				genSeq(n, k, count, index+1, 0)
		} else if lastBit == 0 {
			res = genSeq(n, k, count, index+1, 1) +
				genSeq(n, k, count, index+1, 0)
		}
		return res
	}

	return genSeq(n, k, 0, 1, 1) + genSeq(n, k, 0, 1, 0)
}
