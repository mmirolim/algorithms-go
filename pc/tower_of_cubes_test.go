package pc

import (
	"testing"

	"github.com/mmirolim/algos/checks"
)

type cubeSide int

const (
	front cubeSide = iota
	back
	left
	right
	top
	bottom
)

type TowerOfCubesPos struct {
	w int
	// top facing side
	side cubeSide
}

/*
9.6.6 Tower of Cubes
PC/UVa IDs: 110906/10051, Popularity: C, Success rate: high Level: 3
*/
func TestTowerOfCubes(t *testing.T) {
	cubesSet1 := [][6]int{
		[6]int{1, 2, 2, 2, 1, 2}, [6]int{3, 3, 3, 3, 3, 3}, [6]int{3, 2, 1, 1, 1, 1},
	}
	cubesSet2 := [][6]int{
		[6]int{1, 5, 10, 3, 6, 5}, [6]int{2, 6, 7, 3, 6, 9}, [6]int{5, 7, 3, 2, 1, 9}, [6]int{1, 3, 3, 5, 8, 10}, [6]int{6, 6, 2, 2, 4, 4}, [6]int{1, 2, 3, 4, 5, 6}, [6]int{10, 9, 8, 7, 6, 5}, [6]int{6, 1, 2, 3, 4, 7}, [6]int{1, 2, 3, 3, 2, 1}, [6]int{3, 2, 1, 1, 2, 3}}

	data := []struct {
		cubes     [][6]int
		towerSize int
	}{
		{cubesSet1, 2}, {cubesSet2, 8},
	}
	for i, d := range data {
		tower := TowerOfCubes(d.cubes)
		checks.AssertEq(t, d.towerSize, len(tower), caseStr(i))
	}
}

func TowerOfCubes(cubes [][6]int) [][2]int {
	oppositeSide := func(side int) int {
		if side%2 == 0 {
			return side + 1
		}
		return side - 1
	}
	returnAppropriateCubeSides := func(color int, cube [6]int) []int {
		var sides []int
		for i := range cube {
			if cube[i] == color {
				sides = append(sides, i)
			}
		}
		return sides
	}
	// order of cubes is their weight
	// returns all edges from provided position and color
	findEdges := func(w, side int) [][2]int {
		var pos [][2]int
		// DAG nodes connected in ascending order of weight
		cube := cubes[w]
		for i := w + 1; i < len(cubes); i++ {
			sides := returnAppropriateCubeSides(
				cube[oppositeSide(side)],
				cubes[i])
			for j := range sides {
				pos = append(pos, [2]int{i, sides[j]})
			}
		}
		return pos
	}
	// generate all edges
	mapEdges := map[[2]int][][2]int{}
	for i := range cubes {
		for j := range cubes[i] {
			pos := [2]int{i, j}
			edges := findEdges(pos[0], pos[1])
			mapEdges[pos] = edges
		}
	}
	var maxTower [][2]int
	var stack [][2]int
	push := func(pos [2]int) {
		stack = append(stack, pos)
	}
	peek := func() [2]int {
		return stack[len(stack)-1]
	}
	pop := func() [2]int {
		pos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return pos
	}
	discovered := map[[2]int]bool{}
	resetDiscovered := func() {
		for k := range discovered {
			discovered[k] = false
		}
	}
	dfs := func(pos [2]int) {
		push(pos)
	OUTER:
		for len(stack) > 0 {
			cur := peek()
			for _, e := range mapEdges[cur] {
				if !discovered[e] {
					push(e)
					discovered[e] = true
					continue OUTER
				}
			}
			if len(stack) > len(maxTower) {
				maxTower = maxTower[:0]
				for i := range stack {
					maxTower = append(maxTower, stack[i])
				}
			}
			// backtrack
			pop()
		}
	}
	for k := range mapEdges {
		dfs(k)
		resetDiscovered()
	}
	return maxTower
}
