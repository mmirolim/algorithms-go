package hk

import (
	"math"
	"strings"
)

// https://www.hackerrank.com/challenges/encryption/problem
func Encryption(in string) string {
	in = strings.Replace(in, " ", "", -1)
	r, c := getRC(len(in))
	if r*c < len(in) {
		r++
	}

	var out []string
	str := make([]byte, r)
	for i := 0; i < c; i++ {
		str = str[:0]
		for j := 0; j < r; j++ {
			id := j*c + i
			if id > len(in)-1 {
				break
			}
			str = append(str, in[id])
		}
		out = append(out, string(str))
	}

	return strings.Join(out, " ")
}

func getRC(n int) (int, int) {
	root := math.Sqrt(float64(n))
	return int(math.Floor(root)), int(math.Ceil(root))
}
