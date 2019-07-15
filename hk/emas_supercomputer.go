package hk

// https://www.hackerrank.com/challenges/two-pluses/problem
func EmasSupercomputer(grid []string) int {
	const good, bad = 'G', 'B'
	var areas [][3]int
	// store pluses metadata
	pluses := make([][][][3]int, len(grid))
	for i := 0; i < len(grid); i++ {
		pluses[i] = make([][][3]int, len(grid[0]))
	}
	sizeOfPlustAt := func(r, c int) []int {
		if grid[r][c] == bad {
			return nil
		}
		var areas []int
		limit := r
		if limit > c {
			limit = c
		}
		if limit > len(grid)-r-1 {
			limit = len(grid) - r - 1
		}
		if limit > len(grid[0])-c-1 {
			limit = len(grid[0]) - c - 1
		}

		for i := 1; i <= limit; i++ {
			if grid[r-i][c] == good && grid[r+i][c] == good &&
				grid[r][c-i] == good && grid[r][c+i] == good {
				area := i*4 + 1
				areas = append(areas, area)
			} else {
				break
			}
		}

		for i := range areas {
			for j := range areas[i:] {
				pluses[r-i-1][c] = append(pluses[r-i-1][c], [3]int{areas[i:][j], r, c})
				pluses[r+i+1][c] = append(pluses[r+i+1][c], [3]int{areas[i:][j], r, c})
				pluses[r][c-i-1] = append(pluses[r][c-i-1], [3]int{areas[i:][j], r, c})
				pluses[r][c+i+1] = append(pluses[r][c+i+1], [3]int{areas[i:][j], r, c})
			}
			pluses[r][c] = append(pluses[r][c], [3]int{areas[i], r, c})
		}

		return areas
	}
	// find all plus areas
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			pluseAreas := sizeOfPlustAt(r, c)
			if len(pluseAreas) > 0 {
				for i := range pluseAreas {
					if pluseAreas[i] > 1 {
						areas = append(areas, [3]int{pluseAreas[i], r, c})
					}
				}
			}
		}
	}

	intersect := func(p1, p2 [3]int) bool {
		limit := (p1[0] - 1) / 4
		r, c := p1[1], p1[2]
		// check center
		for _, d := range pluses[r][c] {
			if d[0] == p2[0] && d[1] == p2[1] && d[2] == p2[2] {
				return true
			}
		}
		for i := 1; i <= limit; i++ {
			// check all cells of p1 plus to intersect p2 plus
			for _, d := range [][][3]int{pluses[r-i][c], pluses[r+i][c], pluses[r][c-i], pluses[r][c+i]} {
				for _, v := range d {
					if v[0] == p2[0] && v[1] == p2[1] && v[2] == p2[2] {
						return true
					}
				}

			}
		}
		return false
	}

	max := 1
	for i := 0; i < len(areas); i++ {
		for j := len(areas) - 1; j > i; j-- {
			if !intersect(areas[i], areas[j]) {
				p := areas[i][0] * areas[j][0]
				if p > max {
					max = p
				}
			}
		}
	}

	for i := range areas {
		if areas[i][0] > max {
			max = areas[i][0]
		}
	}

	return max
}
