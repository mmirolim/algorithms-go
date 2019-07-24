package graph

import (
	"testing"

	"github.com/mmirolim/algos/checks"
	tree "github.com/mmirolim/algos/trees"
)

/*
9.6.3 The Tourist Guide
PC/UVa IDs: 110903/10099, Popularity: B, Success rate: average Level: 3
*/
func TestTouristGuide(t *testing.T) {
	cityMap1 := `7 10 weighted
1 2 30
1 3 15
1 4 10
2 4 25
2 5 60
3 4 40
3 6 20
4 7 35
5 7 20
6 7 30
`
	cityMap2 := `3 2 weighted
1 2 15
2 3 10`
	data := []struct {
		start, end    int
		cityMap       string
		numOfTourists int
		numOfTrips    int
		solverName    string
		solver        func(start, end, numOfTourist int, cityMap string) int
	}{
		{start: 1, end: 7, cityMap: cityMap1, numOfTourists: 99, numOfTrips: 5,
			solverName: "TheTouristGuideSolveByFindingAllPaths", solver: TheTouristGuideSolveByFindingAllPaths},
		{start: 1, end: 7, cityMap: cityMap1, numOfTourists: 99, numOfTrips: 5,
			solverName: "TheTouristGuideSolveByFindingSpanningTree", solver: TheTouristGuideSolveByFindingSpanningTree},
		{start: 1, end: 3, cityMap: cityMap2, numOfTourists: 30, numOfTrips: 4,
			solverName: "TheTouristGuideSolveByFindingAllPaths", solver: TheTouristGuideSolveByFindingAllPaths},
		{start: 1, end: 3, cityMap: cityMap2, numOfTourists: 30, numOfTrips: 4,
			solverName: "TheTouristGuideSolveByFindingSpanningTree", solver: TheTouristGuideSolveByFindingSpanningTree},
	}
	for i, d := range data {
		out := d.solver(d.start, d.end, d.numOfTourists, d.cityMap)
		checks.AssertEq(t, d.numOfTrips, out, caseStr(i))
	}
}

func TheTouristGuideSolveByFindingAllPaths(start, end, numOfTourist int, cityMap string) int {
	var paths [][]int
	cityGraph, err := NewGraphFrom(cityMap)
	if err != nil {
		panic(err)
	}
	discovered := make([]bool, cityGraph.NumOfVertices()+1)
	discovered[start] = true
	resetDiscovered := func() {
		for i := range discovered {
			discovered[i] = false
		}
		discovered[start] = true
	}
	var stack []int
	storePath := func() {
		var path []int
		path = append(path, start)
		for i := range stack {
			path = append(path, stack[i])
		}
		paths = append(paths, path)
	}

	var dfs func(v int)
	dfs = func(start int) {
		push(&stack, start)
		discovered[start] = true
	OUTER:
		for len(stack) > 0 {
			cur := peek(&stack)
			if cur == end {
				storePath()
				// pop end from stack
				pop(&stack)
				continue
			}
			for _, v := range cityGraph.edges[cur].ids() {
				if !discovered[v] {
					discovered[v] = true
					push(&stack, v)
					continue OUTER
				}
			}
			pop(&stack)
			discovered[end] = false
		}
	}

	for _, v := range cityGraph.edges[start].ids() {
		dfs(v)
		resetDiscovered()
	}

	edges := cityGraph.EdgesWithWeights()
	maxMin := 0
	minTrips := -1
	findWeight := func(v1, v2 int) int {
		for _, e := range edges {
			if (e[0] == v1 && e[1] == v2) || (e[0] == v2 && e[1] == v1) {
				return e[2]
			}
		}
		panic("edge not found")
	}
	for _, p := range paths {
		min := 1 << 30
		for i := 0; i < len(p)-1; i++ {
			w := findWeight(p[i], p[i+1])
			if min > w {
				min = w
			}
		}
		if maxMin < min {
			maxMin = min
			// minus because Guide will go with each group
			if numOfTourist%(maxMin-1) == 0 {
				minTrips = numOfTourist / (maxMin - 1)
			} else {
				// add one more trip if some tourist left
				minTrips = 1 + numOfTourist/(maxMin-1)
			}
		}
	}
	return minTrips
}

func TheTouristGuideSolveByFindingSpanningTree(start, end, numOfTourist int, cityMap string) int {
	cityGraph, err := NewGraphFrom(cityMap)
	if err != nil {
		panic(err)
	}
	priorityQ, err := tree.NewMaxHeap(-1)
	if err != nil {
		panic(err)
	}
	alreadyInTree := make([]bool, cityGraph.NumOfVertices()+1)
	// val is max weight to edge
	Tree := make([]int, cityGraph.NumOfVertices()+1)
	// find spanning tree of graph with max weight
	edges := map[int][][3]int{}
	for _, e := range cityGraph.EdgesWithWeights() {
		edges[e[0]] = append(edges[e[0]], e)
	}
	for _, e := range edges[start] {
		priorityQ.Insert(e[2], e)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	Tree[start] = numOfTourist
	alreadyInTree[start] = true
	for !priorityQ.IsEmpty() {
		_, data, err := priorityQ.ExtractTop()
		if err != nil {
			panic(err)
		}
		edge := data.([3]int)
		if !alreadyInTree[edge[1]] {
			alreadyInTree[edge[1]] = true
			Tree[edge[1]] = min(Tree[edge[0]], edge[2])
			for _, e := range edges[edge[1]] {
				priorityQ.Insert(e[2], e)
			}
		}
	}
	// minus because Guide will go with each group
	numOfTrips := numOfTourist / (Tree[end] - 1)
	if numOfTourist%(Tree[end]-1) != 0 {
		numOfTrips++
	}

	return numOfTrips
}
