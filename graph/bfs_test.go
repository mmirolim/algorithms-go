package graph

import (
	"testing"
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

		if len(visitOrder) != len(d.expectedOrder) {
			t.Errorf("case [%d] expected visit order %v, got %v", i, d.expectedOrder, visitOrder)
			continue
		}

		for j, id := range d.expectedOrder {
			if id != visitOrder[j] {
				t.Errorf("case [%d] expected visit order %v, got %v", i, d.expectedOrder, visitOrder)
				break
			}
		}

		if len(parents) != len(d.expectedParents) {
			t.Errorf("case [%d] expected parents %v, got %v", i, d.expectedParents, parents)
			continue
		}

		for j, id := range d.expectedParents {
			if id != parents[j] {
				t.Errorf("case [%d] expected parents %v, got %v", i, d.expectedParents, parents)
				break
			}
		}

	}
}
