package hk

// https://www.hackerrank.com/challenges/3d-surface-area/problem
func D3SurfaceArea(data [][]int32) int {
	area := 2 * len(data) * len(data[0]) // top + bottom area
	for r := 0; r < len(data); r++ {
		area += rowArea(r, data)
	}

	for c := 0; c < len(data[0]); c++ {
		area += colArea(c, data)
	}

	return area
}

func absInt32(a int32) int32 {
	if a >= 0 {
		return a
	}
	return -1 * a
}

func rowArea(r int, data [][]int32) int {
	area := data[r][0]
	for i := 1; i < len(data[0]); i++ {
		area += absInt32(data[r][i] - data[r][i-1])
	}
	area += data[r][len(data[0])-1]
	return int(area)
}

func colArea(c int, data [][]int32) int {
	area := data[0][c]
	for i := 1; i < len(data); i++ {
		area += absInt32(data[i][c] - data[i-1][c])
	}
	area += data[len(data)-1][c]
	return int(area)
}
