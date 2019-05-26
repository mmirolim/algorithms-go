package strings

func LongestSubSeqDP(s1, s2 string) string {
	out := []byte{}
	dp := longestSubSeqLenDP(s1, s2)
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			out = append(out, s1[i])
			i++
			j++
		} else if dp[i+1][j] >= dp[i][j+1] {
			i++
		} else {
			j++
		}
	}

	return string(out)
}

func LongestSubSeqLenDP(s1, s2 string) int {
	dp := longestSubSeqLenDP(s1, s2)
	return dp[0][0]
}

func LongestSubSeqLenDPReducedSpace(s1, s2 string) int {
	rowPrev := make([]int, len(s2)+1)
	row := make([]int, len(s2)+1)

	for i := len(s1); i >= 0; i-- {
		for j := len(s2); j >= 0; j-- {
			if len(s1) == i || len(s2) == j {
				row[j] = 0
			} else if s1[i] == s2[j] {
				row[j] = 1 + rowPrev[j+1]
			} else {
				l1 := rowPrev[j]
				l2 := row[j+1]
				if l1 < l2 {
					l1 = l2
				}
				row[j] = l1
			}
		}
		copy(rowPrev[:], row)
	}

	return row[0]
}

func longestSubSeqLenDP(s1, s2 string) [][]int {
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
	}

	for i := len(s1); i >= 0; i-- {
		for j := len(s2); j >= 0; j-- {
			if len(s1) == i || len(s2) == j {
				memo[i][j] = 0
			} else if s1[i] == s2[j] {
				memo[i][j] = 1 + memo[i+1][j+1]
			} else {
				l1 := memo[i+1][j]
				l2 := memo[i][j+1]
				if l1 < l2 {
					l1 = l2
				}
				memo[i][j] = l1
			}
		}
	}
	return memo
}
func LongestSubSeqLenWithMemoization(s1, s2 string) int {
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var lcs func(id1, id2 int) int
	lcs = func(id1, id2 int) int {
		if memo[id1][id2] == -1 {
			if len(s1) == id1 || len(s2) == id2 {
				memo[id1][id2] = 0
			} else if s1[id1] == s2[id2] {
				memo[id1][id2] = 1 + lcs(id1+1, id2+1)
			} else {
				l1 := lcs(id1+1, id2)
				l2 := lcs(id1, id2+1)
				if l1 < l2 {
					l1 = l2
				}
				memo[id1][id2] = l1
			}
		}
		return memo[id1][id2]
	}

	return lcs(0, 0)

}
func LongestSubSeqLen(s1, s2 string, s1idx, s2idx int) int {
	if len(s1) == s1idx || len(s2) == s2idx {
		return 0
	} else if s1[s1idx] == s2[s2idx] {
		return 1 + LongestSubSeqLen(s1, s2, s1idx+1, s2idx+1)
	}
	l1 := LongestSubSeqLen(s1, s2, s1idx, s2idx+1)
	l2 := LongestSubSeqLen(s1, s2, s1idx+1, s2idx)
	if l1 > l2 {
		return l1
	}
	return l2

}

func SubSeq(pat, s string) bool {
	i := 0
	j := 0
	for i < len(pat) {
		if j == len(s) {
			return false
		}
		if pat[i] != s[j] {
			j++
		} else {
			i++
			j++
		}
	}

	return true
}

// first my BF brute force approach
func LongestSubSeqBF(s1, s2 string) string {
	check := func(ch byte, s string) int {
		id := -1
		for i := 0; i < len(s); i++ {
			if s[i] == ch {
				id = i
				break
			}
		}
		return id
	}
	generateAllSubSeq := func(s string) map[string]bool {
		out := map[string]bool{}
		var subSeq func(sub string, id int)
		subSeq = func(sub string, id int) {
			if len(sub) == 0 {
				return
			}
			seq := append([]byte(sub[:id]), sub[id+1:]...)
			out[string(seq)] = true
			for i := 0; i < len(seq); i++ {
				subSeq(string(seq), i)
			}
		}
		out[s] = true
		for i := 0; i < len(s); i++ {
			subSeq(s, i)
		}

		return out
	}
	sets := generateAllSubSeq(s2)
	maxSub := ""
	maxLen := 0
	for set := range sets {
		count := 0
		cursor := 0
		for j := 0; j < len(set); j++ {
			id := check(set[j], s1[cursor:])
			if id > -1 {
				count++
				cursor = cursor + id + 1
			} else {
				break
			}
		}

		if count > maxLen {
			maxLen = count
			maxSub = set
		}
	}

	return maxSub
}
