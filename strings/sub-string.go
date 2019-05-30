package strings

func RabinKarpSubStringSearch(pat, str string, prime int) []int {
	var out []int
	if len(pat) > len(str) {
		return []int{-1}
	}
	h := 1
	const MAXCHAR = 256 // ASCII ext

	// compute alpha^(M-1) mod prime
	for i := 0; i < len(pat)-1; i++ {
		h = (h * MAXCHAR) % prime
	}
	patHash, subStrHash := 0, 0

	for i := range pat {
		patHash = (MAXCHAR*patHash + int(pat[i])) % prime
		subStrHash = (MAXCHAR*subStrHash + int(str[i])) % prime
	}

	for i := 0; i <= len(str)-len(pat); i++ {
		if patHash == subStrHash {
			// verify
			found := true
			for j := range pat {
				if pat[j] != str[i+j] {
					found = false
					break
				}
			}
			if found {
				out = append(out, i)
			}
		}
		if i < len(str)-len(pat) {
			subStrHash =
				(MAXCHAR*(subStrHash-h*int(str[i])) + int(str[i+len(pat)])) % prime
			if subStrHash < 0 {
				subStrHash += prime
			}
		}
	}

	if len(out) == 0 {
		return []int{-1}
	}
	return out
}

func searchMinWindow(pat, str string) (index, minlen int) {
	strDic := map[byte]int{}
	patDic := map[byte]int{}
	for i := 0; i < len(pat); i++ {
		patDic[pat[i]]++
	}

	start := 0
	minlen = 1 << 30
	startIndex := -1
	count := 0
	// TODO make it work with bytes then with unicode
	for i := 0; i < len(str); i++ {
		ch := str[i]
		strDic[ch]++
		v1, ok := patDic[ch]
		if ok && v1 >= strDic[ch] {
			count++
		}

		if count == len(pat) {
			for {
				v1, ok := patDic[str[start]]
				if !ok || v1 < strDic[str[start]] {
					strDic[str[start]]--
					start++
				} else {
					break
				}
			}

			winsize := i - start + 1
			if minlen > winsize {
				minlen = winsize
				startIndex = start
			}
		}

	}
	return startIndex, minlen
}
