package hk

// TODO solve
// https://www.hackerrank.com/challenges/string-transmission/problem
func StringTransmission(k int, s string) int {
	panic("not implemented yet")
	const prime = 1000000007
	res := 0
	if k == 0 {
		// TODO hanlde properly
		if len(s) == 1 {
			return 1
		}
		return res
	}
	if k == len(s) {
		res += powModPrime(2, k, prime)
	} else {
		for r := 1; r <= k; r++ {
			// n - r
			nr := len(s) - r
			if nr < r {
				rfact := mulRangeModPrime(1, nr, prime)
				res += (mulRangeModPrime(r+1, len(s), prime) / rfact) % prime
			} else {
				rfact := mulRangeModPrime(1, r, prime)
				res += (mulRangeModPrime(nr+1, len(s), prime) / rfact) % prime
			}
		}
	}
	// TODO compute periodic permutations and remove from total
	return res
}

func powModPrime(a, n, prime int) int {
	if n == 0 {
		return 1
	}
	x := powModPrime(a, n/2, prime)
	if n%2 == 0 {
		return (x * x) % prime
	}
	return (((a * x) % prime) * x) % prime
}

func mulRangeModPrime(from, to, prime int) int {
	product := 1
	for i := from; i <= to; i++ {
		product = (product % prime) * i
		if product < 0 {
			product += prime
		}
	}
	return product
}

var stIn = `
6 3
101010
2 1
11
9 4
010100000
8 8
11111001
8 1
11010111
6 2
001101
7 4
1110000
9 8
011100101
1 1
0
6 6
011001
3 2
011
6 6
010011
7 0
0010111
1 0
1
1 0
0
7 2
1000001
3 0
000
10 5
0011110100
4 4
1101
4 3
0010`
var stout = `
33
2
252
240
9
19
97
503
2
54
5
54
1
1
1
28
0
613
12
11`
