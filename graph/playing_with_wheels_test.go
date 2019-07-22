package graph

import (
	"testing"

	"github.com/mmirolim/algos/checks"
)

/*
9.6.2 Playing With Wheels
PC/UVa IDs: 110902/10067, Popularity: C, Success rate: average Level: 2
*/
func TestPlayingWithWheels(t *testing.T) {
	data := []struct {
		start, target [4]int
		forbidden     [][4]int
		res           int
	}{
		{[4]int{8, 0, 5, 6}, [4]int{6, 5, 0, 8},
			[][4]int{[4]int{8, 0, 5, 7}, [4]int{8, 0, 4, 7}, [4]int{5, 5, 0, 8},
				[4]int{7, 5, 0, 8}, [4]int{6, 4, 0, 8}}, 14},
		{[4]int{0, 0, 0, 0}, [4]int{5, 3, 1, 7}, [][4]int{
			[4]int{0, 0, 0, 1}, [4]int{0, 0, 0, 9}, [4]int{0, 0, 1, 0}, [4]int{0, 0, 9, 0},
			[4]int{0, 1, 0, 0}, [4]int{0, 9, 0, 0}, [4]int{1, 0, 0, 0}, [4]int{9, 0, 0, 0},
		}, -1},
	}
	for i, d := range data {
		out := PlayingWithWheels(d.start, d.target, d.forbidden)
		checks.AssertEq(t, d.res, out, caseStr(i))
	}
}

func PlayingWithWheels(start, target [4]int, forbidden [][4]int) int {
	var q []int
	// use bfs
	discovered := make([]bool, 10000)
	tonum := func(cfg [4]int) int {
		return cfg[0]*1000 + cfg[1]*100 + cfg[2]*10 + cfg[3]
	}
	for i := range forbidden {
		discovered[tonum(forbidden[i])] = true
	}
	getAdjacentCfg := func(cur int) []int {
		var adjacent []int
		for _, p := range []int{1, 10, 100, 1000} {
			d := (cur % (10 * p)) / p
			dp1 := (d + 1) % 10
			dm1 := d - 1
			if dm1 < 0 {
				dm1 = 9
			}
			next := cur - d*p
			adjacent = append(adjacent, next+dp1*p)
			adjacent = append(adjacent, next+dm1*p)
		}
		return adjacent
	}
	startnum := tonum(start)
	enqueue(&q, startnum)
	targetnum := tonum(target)
	discovered[startnum] = true
	parent := make([]int, 10000)
	for len(q) > 0 {
		cur := dequeue(&q)
		if cur == targetnum {
			p := parent[cur]
			count := 1
			for startnum != p {
				p = parent[p]
				count++
			}
			return count
		}
		for _, v := range getAdjacentCfg(cur) {
			if !discovered[v] {
				discovered[v] = true
				parent[v] = cur
				push(&q, v)
			}
		}
	}
	return -1
}
