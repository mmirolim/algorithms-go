package strings

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
