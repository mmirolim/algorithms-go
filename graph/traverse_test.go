package graph

import (
	"errors"
	"testing"

	"github.com/mmirolim/algos/checks"
)

func TestBFS(t *testing.T) {
	var visitOrder, parents []int
	processEarly := func(v int) {
		visitOrder = append(visitOrder, v)
	}

	data := []struct {
		graph           string
		start           int
		expectedOrder   []int
		returnParents   bool
		expectedParents []int
	}{
		{graphDataDirectedWeightedWithCodes, 1, []int{1, 2, 3, 4}, false, nil},
		{graphDataDirectedWeightedWithCodes, 1, []int{1, 2, 3, 4}, true, []int{0, 0, 1, 1, 3, 0}},
		// vertex with no edges
		{graphDataDirectedWeightedWithCodes, 5, []int{5}, true, []int{0, 0, 0, 0, 0, 0}},

		{twoRhombus, 1, []int{1, 3, 2, 4, 5, 7, 6}, true, []int{0, 0, 1, 1, 3, 4, 5, 4}},
		// going from highest degree vertex
		{twoRhombus, 4, []int{4, 2, 3, 5, 7, 1, 6}, true, []int{0, 2, 4, 4, 0, 4, 5, 4}},
	}
	for i, d := range data {
		visitOrder, parents = visitOrder[:0], parents[:0]
		g, err := NewGraphFrom(d.graph)
		if err != nil {
			t.Errorf("case [%d] NewGraphFrom unexpected err %+v", i, err)
			continue
		}
		// Test method
		if d.returnParents {
			g.BFS(d.start, processEarly, nil, nil, &parents)
		} else {
			g.BFS(d.start, processEarly, nil, nil, nil)
		}

		checks.AssertEq(t, d.expectedOrder, visitOrder, caseStr(i))
		checks.AssertEq(t, d.expectedParents, parents, caseStr(i))

	}
}

func TestDFS(t *testing.T) {
	var visitOrder, parents []int
	processEarly := func(v int) error {
		visitOrder = append(visitOrder, v)
		return nil
	}
	processLate := func(v int) error {
		// noop
		return nil
	}
	countEdge := 0
	processEdge := func(c, n int) error {
		countEdge++
		return nil
	}

	data := []struct {
		graph         string
		start         int
		expectedOrder []int
		edgesNumber   int
	}{
		{graphDataDirectedWeightedWithCodes, 1, []int{1, 2, 3, 4}, 4},
		// vertex with no edges
		{graphDataDirectedWeightedWithCodes, 5, []int{5}, 0},
		{twoRhombus, 1, []int{1, 3, 4, 2, 5, 6, 7}, 8},
		// // going from highest degree vertex
		{twoRhombus, 4, []int{4, 2, 1, 3, 5, 6, 7}, 8},
	}
	for i, d := range data {
		countEdge = 0
		visitOrder, parents = visitOrder[:0], parents[:0]
		g, err := NewGraphFrom(d.graph)
		if err != nil {
			t.Errorf("case [%d] NewGraphFrom unexpected err %+v", i, err)
			continue
		}
		g.DFS(d.start, processEarly, processLate, processEdge, nil)

		checks.AssertEq(t, d.edgesNumber, countEdge, caseStr(i))
		checks.AssertEq(t, d.expectedOrder, visitOrder, caseStr(i))

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
