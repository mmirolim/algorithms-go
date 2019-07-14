package main

import (
	"fmt"
	"strings"
)

// http://codeforces.com/contest/1194/problem/C
func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	strs := make([][3]string, t)
	for i := 0; i < t; i++ {
		var s, t, p string
		fmt.Scanf("%s\n", &s)
		fmt.Scanf("%s\n", &t)
		fmt.Scanf("%s\n", &p)
		strs[i] = [3]string{s, t, p}
	}
	for i := range strs {
		fmt.Println(solve(strs[i][0], strs[i][1], strs[i][2]))
	}
}

func solve(s, t, p string) string {
	res := "YES"
	if len(s) > len(t) {
		return "NO"
	}
	freqs := map[byte]int{}
	freqt := map[byte]int{}
	freqp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		freqs[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		freqt[t[i]]++
	}
	for i := 0; i < len(p); i++ {
		freqp[p[i]]++
	}
	for k, v := range freqt {
		if v > freqp[k]+freqs[k] {
			return "NO"
		}
	}
	// check exists
	for k := range freqs {
		if freqt[k] == 0 {
			return "NO"
		}
	}
	id := 0
	for i := 0; i < len(s); i++ {
		t = t[id:]
		id = strings.IndexByte(t, s[i])
		if id == -1 {
			return "NO"
		}
		id++
	}

	return res
}
