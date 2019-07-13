package hk

// https://www.hackerrank.com/challenges/absolute-permutation
func AbsolutePermutation(n, k int) []int {
	var out []int
	nSum := 0
	permSum := 0
	alreadyUsed := map[int]bool{}
	for i := 1; i <= n; i++ {
		nSum += i
		pos1 := i - k
		pos2 := k + i
		if i-k > 0 {
			if !alreadyUsed[pos1] {
				pos2 = pos1
			}
		}
		if pos2 > n {
			return nil
		}
		if alreadyUsed[pos2] {
			return nil
		} else {
			alreadyUsed[pos2] = true
		}
		out = append(out, pos2)
		permSum += pos2
	}
	if permSum != nSum {
		return nil
	}
	return out
}
