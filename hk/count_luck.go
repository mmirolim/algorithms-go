package hk

// https://www.hackerrank.com/challenges/count-luck/problem
func CountLuck(k int, grid []string) bool {
	const Empty, Tree, M, Key = '.', 'X', 'M', '*'
	var stack [][2]int
	push := func(p [2]int) {
		stack = append(stack, p)
	}
	pop := func() [2]int {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}
	peek := func() [2]int {
		return stack[len(stack)-1]
	}
	getNeighbors := func(p [2]int) [][2]int {
		var out [][2]int
		r, c := p[0], p[1]

		if r-1 >= 0 {
			if grid[r-1][c] != Tree {
				out = append(out, [2]int{r - 1, c})
			}
		}
		if r+1 < len(grid) {
			if grid[r+1][c] != Tree {
				out = append(out, [2]int{r + 1, c})
			}
		}
		if c-1 >= 0 {
			if grid[r][c-1] != Tree {
				out = append(out, [2]int{r, c - 1})
			}
		}
		if c+1 < len(grid[0]) {
			if grid[r][c+1] != Tree {
				out = append(out, [2]int{r, c + 1})
			}
		}
		return out
	}
	findPath := func(start [2]int) {
		visited := make([][]bool, len(grid))
		for i := range grid {
			visited[i] = make([]bool, len(grid[0]))
		}
		// dfs with backtracking
		push(start)
	OUTER:
		for len(stack) > 0 {
			cur := peek()
			if !visited[cur[0]][cur[1]] {
				visited[cur[0]][cur[1]] = true
				if grid[cur[0]][cur[1]] == Key {
					break
				}
			}
			for _, p := range getNeighbors(cur) {
				if !visited[p[0]][p[1]] {
					push(p)
					continue OUTER
				}
			}
			pop()

		}
	}
	isMultiOp := func(p [2]int) int {
		count := 0
		r, c := p[0], p[1]
		if r-1 >= 0 {
			if grid[r-1][c] != Tree {
				count++
			}
		}
		if r+1 < len(grid) {
			if grid[r+1][c] != Tree {
				count++
			}
		}
		if c-1 >= 0 {
			if grid[r][c-1] != Tree {
				count++
			}
		}
		if c+1 < len(grid[0]) {
			if grid[r][c+1] != Tree {
				count++
			}
		}

		return count
	}
	getStartPoint := func(g []string) [2]int {
		for r := 0; r < len(g); r++ {
			for c := 0; c < len(g[0]); c++ {
				if g[r][c] == M {
					return [2]int{r, c}
				}
			}
		}
		return [2]int{}
	}
	findPath(getStartPoint(grid))
	count := 0
	pop()
	// start point
	if isMultiOp(stack[0]) > 1 {
		count++
	}
	stack = stack[1:]
	for i := range stack {
		if isMultiOp(stack[i]) > 2 {
			count++
		}
	}
	return count == k
}
