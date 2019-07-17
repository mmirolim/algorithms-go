package graph

import (
	"errors"
	"testing"

	"github.com/mmirolim/algos/checks"
)

func TestGraphComponents(t *testing.T) {
	graph1 := `7 6
1 2
1 3
2 3
3 4
5 6
5 7
`

	graph2 := `7 5
1 2
1 3
2 3
3 4
5 6
7
`

	data := []struct {
		text     string
		expected [][]int
	}{
		{graph1, [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7}}},
		{graph2, [][]int{[]int{1, 2, 3, 4}, []int{5, 6}, []int{7}}},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.text)
		if !checks.AssertEqErr(t, nil, err, caseStr(i)) {
			continue
		}
		comps := g.Components()
		checks.AssertEq(t, d.expected, comps, caseStr(i))
	}
}

func TestGraphBipartite(t *testing.T) {
	graphBipartite := `10 9
1 2
3 4
3 6
7 10
7 6
10 5
1 8
9 8
9 6
`

	graphBipartiteWith2Comp := `6 5
1 3
1 5
3 1
2 4
4 6
`
	graphNotBipartite := `6 7
1 2
1 4
3 2
3 4
4 5
4 6
1 6
`

	data := []struct {
		graph     string
		bipartite bool
		groups    [2][]int
	}{
		{graphBipartite, true, [2][]int{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}}},
		{graphBipartiteWith2Comp, true, [2][]int{[]int{1, 2, 6}, []int{3, 4, 5}}},
		{graphNotBipartite, false, [2][]int{}},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.graph)
		if !checks.AssertEqErr(t, nil, err, caseStr(i)) {
			continue
		}
		groups, ok := g.Bipartite()
		checks.AssertEq(t, d.bipartite, ok, caseStr(i))
		checks.AssertEq(t, d.groups, groups, caseStr(i))
	}
}

func TestFindGraphCycle(t *testing.T) {
	data := []struct {
		graph string
		start int
		path  []int
		err   error
	}{
		{graphDataDirectedWeightedWithCodes, 1, nil, errors.New("path not found")},
		// vertex with no edges
		{graphDataDirectedWeightedWithCodes, 5, nil, nil},
		{twoRhombus, 1, []int{2, 4, 3, 1}, nil},
		// // going from highest degree vertex
		{twoRhombus, 4, []int{3, 1, 2, 4}, nil},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.graph)
		if err != nil {
			t.Errorf("case [%d] NewGraphFrom unexpected err %+v", i, err)
			continue
		}

		path, err := g.FindCycle(d.start)
		checks.AssertEqErr(t, d.err, err, caseStr(i))
		checks.AssertEq(t, d.path, path, caseStr(i))

	}
}

func TestFindArticulations(t *testing.T) {
	data := []struct {
		graph                  string
		componentArticulations [][]int
		err                    error
	}{
		{graphDataDirectedWeightedWithCodes, nil, errors.New("error graph is directed")},
		{graphDataWithNamesNotDirectedWithoutWeight, [][]int{[]int{3}}, nil},
		{twoRhombus, [][]int{[]int{4}}, nil},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.graph)
		if err != nil {
			t.Errorf("case [%d] NewGraphFrom unexpected err %+v", i, err)
			continue
		}

		out, err := g.FindArticulations()
		checks.AssertEqErr(t, d.err, err, caseStr(i))
		checks.AssertEq(t, d.componentArticulations, out, caseStr(i))

	}
}

func TestTopologicalSort(t *testing.T) {
	twoRhombusDirected := `7 8 directed
1 3
1 2
2 4
3 4
4 5
4 7
5 6
6 7
`
	graphDirectedWith2Component := `5 4 directed
	1 2
	1 3
	2 3
	3 4
	5
	`
	graphDirectedWithCycle := `4 5 directed
	1 2
	1 3
	2 3
	3 4
	4 1
	`
	data := []struct {
		graph  string
		sorted []int
		err    error
	}{
		{twoRhombus, nil, errors.New("error graph is not directed")},
		{graphDirectedWithCycle, nil, errors.New("error directed cycle found, not a DAG")},
		{graphDirectedWith2Component, []int{4, 3, 2, 1, 5}, nil},
		{twoRhombusDirected, []int{7, 6, 5, 4, 3, 2, 1}, nil},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.graph)
		if err != nil {
			t.Errorf("case [%d] NewGraphFrom unexpected err %+v", i, err)
			continue
		}

		out, err := g.TopologicalSort()
		checks.AssertEqErr(t, d.err, err, caseStr(i))
		checks.AssertEq(t, d.sorted, out, caseStr(i))

	}
}
