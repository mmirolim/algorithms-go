package number

import (
	"strings"
)

func DelannoyNumber(n, m int) int {
	var D func(n, m int) int
	D = func(n, m int) int {
		if n == 0 || m == 0 {
			return 1
		}

		return D(n-1, m) + D(n, m-1) + D(n-1, m-1)
	}

	return D(n, m)
}

func DelannoyNumberDP(n, m int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] + dp[i-1][j-1]
		}
	}

	return dp[n][m]
}

func ConvertNumberToWords(num string) string {
	if len(num) > 12 {
		return "no implementaion"
	}
	digits := []string{
		"zero", "one", "two",
		"three", "four", "five",
		"six", "seven", "eight", "nine",
	}
	twoDigits := []string{"ten", "eleven", "twelve",
		"thirteen", "fourteen",
		"fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
	}
	tensMultiple := []string{
		"", "", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety",
	}

	if len(num) == 1 {
		return digits[num[0]-'0']
	}
	getDigit := func(n int) int {
		return int(num[len(num)-n] - '0')
	}
	out := []string{}
	push := func(s string) {
		if s != digits[0] {
			out = append(out, s)
		}
	}
	parseDigits := func(i int) int {
		d := getDigit(i)
		if i%3 == 0 {
			if getDigit(i) > 0 {
				push(digits[d])
				push("hundred")
			}
		} else if (i+1)%3 == 0 {
			// consume next two digits
			i--
			if d >= 2 {
				push(tensMultiple[d])
				push(digits[getDigit(i)])
			} else if d == 1 {
				push(twoDigits[getDigit(i)])
			} else if d == 0 {
				push(digits[getDigit(i)])
			}
		} else {
			push(digits[d])
		}

		return i

	}
	i := len(num)
	for {
		if i == 2 {
			break
		}
		i = parseDigits(i)
		if i == 10 {
			push("billion")
		} else if i == 7 {
			push("million")
		} else if i == 4 {
			push("thousand")
		}
		i--
	}

	if len(num) > 2 && (getDigit(2) > 0 || getDigit(1) > 0) {
		push("and")
	}
	parseDigits(2)
	return strings.Join(out, " ")
}
