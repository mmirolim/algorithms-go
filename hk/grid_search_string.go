package hk

// https://www.hackerrank.com/challenges/the-grid-search/problem
func TheGridSearch(g, p []string) (found bool) {
	for r := 0; r <= len(g)-len(p); r++ {
		indices := subStringSearch(g[r], p[0])
		for i := range indices {
			if subGridTest(r, indices[i], g, p) {
				return true
			}
		}
	}
	return
}

// find all matches
func subStringSearch(str, sub string) (indices []int) {
	// rabin-karp algo
	h := 1
	const MAXCHAR = 255
	const PRIME = 619
	// precompute alpha^(M-1) mod prime
	for i := 0; i < len(sub)-1; i++ {
		h = (h * MAXCHAR) % PRIME
	}
	subHash, subStrHash := 0, 0
	for i := range sub {
		subHash = (MAXCHAR*subHash + int(sub[i])) % PRIME
		subStrHash = (MAXCHAR*subStrHash + int(str[i])) % PRIME
	}
	for i := 0; i <= len(str)-len(sub); i++ {
		if subHash == subStrHash {
			// verify
			found := true
			for j := range sub {
				if sub[j] != str[i+j] {
					found = false
					break
				}
			}
			if found {
				indices = append(indices, i)
			}
		}
		if i < len(str)-len(sub) {
			subStrHash = (MAXCHAR*(subStrHash-h*int(str[i])) + int(str[i+len(sub)])) % PRIME
			if subStrHash < 0 {
				subStrHash += PRIME
			}
		}
	}
	return
}

// search from r, c in grid for pattern
func subGridTest(r, c int, g, p []string) bool {
	maxRow, maxCol := r+len(p), c+len(p[0])
	for i := r; i < maxRow; i++ {
		for j := c; j < maxCol; j++ {
			if g[i][j] != p[i-r][j-c] {
				return false
			}
		}
	}
	return true
}
