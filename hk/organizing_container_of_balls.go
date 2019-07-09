package hk

// https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem
func OrganizingContainersOfBalls(cons [][]int32) bool {
	sumOfRow := sumOfFirstRow(cons)
	for i := 0; i < len(cons); i++ {
		if sumOfRow == sumOfFirstCol(i, cons) {
			return true
		}
	}

	return false
}

func sumOfFirstRow(cons [][]int32) (sum int32) {
	for i := 0; i < len(cons); i++ {
		sum += cons[0][i]
	}
	return
}

func sumOfFirstCol(c int, cons [][]int32) (sum int32) {
	for i := 0; i < len(cons); i++ {
		sum += cons[i][c]
	}
	return
}
