package number

import (
	"strings"
)

func LightMoreLight(n int) bool {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count%2 != 0
}

func StanVsOllie(n int) bool {
	x := 1
	stanWins := false
	for n > x {
		x *= 9
		if x >= n {
			stanWins = true
			break
		}
		x *= 2
	}
	return stanWins
}
func CountCarryInt(a, b int) int {
	digitSum := func(n int) int {
		sum := 0
		for ; n > 0; n /= 10 {
			sum += n % 10
		}
		return sum
	}

	return (digitSum(a) + digitSum(b) - digitSum(a+b)) / 9
}

func CountCarryIter(num1, num2 string) int {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	count := 0
	carry := 0
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		n2 := int(num2[i] - '0')
		if n1+n2+carry >= 10 {
			carry = 1
			count++
		} else {
			carry = 0
		}
	}

	if len(num2) > len(num1) {
		n2 := int(num2[len(num2)-1-len(num1)] - '0')
		if n2+carry >= 10 {
			count++
		}
	}
	return count
}

func CountCarry(num1, num2 string) int {
	var countCarry func(id1, id2, carry int) int
	countCarry = func(id1, id2, carry int) int {
		if id1 < 0 && id2 < 0 {
			return 0
		}

		n1, n2 := 0, 0
		if id1 >= 0 {
			n1 = int(num1[id1] - '0')
		}
		if id2 >= 0 {
			n2 = int(num2[id2] - '0')
		}
		if n1+n2+carry >= 10 {
			return 1 + countCarry(id1-1, id2-1, 1)
		}

		return countCarry(id1-1, id2-1, 0)
	}

	return countCarry(len(num1)-1, len(num2)-1, 0)
}

func FindDropsKMarbleBreaksInNFloorBuilding(n, k int) int {
	var findMaxFloor func(d, k int) int
	findMaxFloor = func(d, k int) int {
		if d == 0 || k == 0 {
			return 0
		} else if d == 1 {
			return 1
		} else if k == 1 {
			return d
		}

		return findMaxFloor(d-1, k-1) + findMaxFloor(d-1, k) + 1
	}
	d := k
	for {
		if findMaxFloor(d, k) >= n {
			break
		}
		d++
	}

	return d
}

func StairCaseProblemDPSolution(totalSteps int, steps []int) int {
	dp := make([]int, totalSteps+1)

	var countWays func(totalSteps int) int

	countWays = func(totalSteps int) int {
		if totalSteps == 0 {
			return 1
		} else if totalSteps < 0 {
			return 0
		}
		if dp[totalSteps] > 0 {
			return dp[totalSteps]
		}

		for _, s := range steps {
			dp[totalSteps] += countWays(totalSteps - s)
		}
		return dp[totalSteps]
	}

	return countWays(totalSteps)
}

func StairCaseProblemOrderDoestNotMatter(totalSteps int) int {
	// possible steps [1, 2]
	if totalSteps%2 == 0 {
		return 1 + totalSteps/2
	} else {
		return 1 + (totalSteps-1)/2
	}

}

func StairCaseProblem(totalSteps int, steps []int) int {
	if totalSteps == 0 {
		return 1
	} else if totalSteps < 0 {
		return 0
	}
	res := 0
	for _, s := range steps {
		res += StairCaseProblem(totalSteps-s, steps)
	}

	return res
}

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

/*
You have an unordered array X of n integers.
Find the array M containing n elements where Mi
is the product of all integers in X except for Xi.
You may not use division. You can use extra memory.
(Hint: There are solutions faster than O(n2).)
*/
func ADM_3_28(X []int) []int {
	M := make([]int, len(X))
	RP := make([]int, len(X))
	LP := make([]int, len(X))
	rp := 1
	lp := 1
	for i := range X {
		rp = X[i] * rp
		RP[i] = rp
		lp = X[len(X)-i-1] * lp
		LP[len(X)-i-1] = lp
	}

	M[0] = LP[0]
	M[len(X)-1] = RP[len(X)-2]

	for i := 1; i < len(M)-1; i++ {
		M[i] = RP[i-1] * LP[i+1]
	}

	return M
}
