package graph

import (
	"strconv"
	"testing"

	"github.com/mmirolim/algos/checks"
)

func TestNewGraphFrom(t *testing.T) {
	data := []struct {
		text                      string
		expectedErr               error
		numOfVertices, numOfEdges int
		isDirected, isWeighted    bool
		codes                     map[int]string
		vertices                  []int
		edges                     [][2]int
		edgesWithWeights          [][3]int
	}{
		{graphDataSimple, nil, 5, 4, false, false,
			nil, vertices, justEdges, edgesWithNoWeights},
		{graphDataDirectedWeightedWithCodes, nil, 5, 4, true, true,
			codes, vertices, justEdges, edgesWithWeights},
		{graphDataWithNamesNotDirectedWithoutWeight, nil, 5, 4, false, false,
			codes, vertices, justEdges, edgesWithNoWeights},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.text)
		if !checks.AssertEqErr(t, d.expectedErr, err, caseStr(i)) {
			continue
		}
		checks.AssertEq(t, d.isDirected, g.IsDirected(), caseStr(i))
		checks.AssertEq(t, d.isWeighted, g.IsWeighted(), caseStr(i))

		vertices := g.Vertices()
		checks.AssertEq(t, d.vertices, vertices, caseStr(i)+" vertices")

		edgesFromGraph := g.Edges()
		checks.AssertEq(t, d.numOfEdges, len(edgesFromGraph), caseStr(i)+" edges")

		for j, edge := range d.edges {
			if edge[0] != edgesFromGraph[j][0] || edge[1] != edgesFromGraph[j][1] {
				t.Errorf("case [%v] expected edge %v, got %v",
					i, d.edges, edgesFromGraph)
				break
			}
		}

		edges := g.EdgesWithWeights()
		for j, edge := range d.edgesWithWeights {
			if edge[0] != edges[j][0] ||
				edge[1] != edges[j][1] ||
				edge[2] != edges[j][2] {
				t.Errorf("case [%v] [%v] expected edge %v, got %v",
					i, j, d.edgesWithWeights, edges)
				break
			}
		}
	}
}
func TestNewGraphFromConsistencyCheck(t *testing.T) {
	wrongNumOfCodes := `5 4 directed weighted hascodes
2  Hillary Clinton
3  John McCain
4  George Bush
5  Saddam Hussein
1 2 10
1 3 2
2 3 5
3 4 4
4
5
`
	wrongNumOfVertices := `5 4
1 2
1 3
2 3
3 4`
	weightMissing := `5 4 directed weighted hascodes
1  Bill Clinton
2  Hillary Clinton
3  John McCain
4  George Bush
5  Saddam Hussein
1 2 10
1 3 2
2 3
3 4 4
4
5`
	wrongNumOfEdges := `5 4
1 2
2 3
3 4
5`
	wrongNumOfVerticesAndEdges := `5 4
1 2
1 3
2 3
3 4
5 6
5 7
`

	data := []struct {
		wrongData   string
		expectedErr error
	}{
		{wrongNumOfCodes, ErrWrongNumOfEdges},
		{wrongNumOfVertices, ErrWrongNumOfVertices},
		{weightMissing, ErrWeightMissing},
		{wrongNumOfEdges, ErrWrongNumOfEdges},
		{wrongNumOfVerticesAndEdges, ErrWrongNumOfVertices},
	}
	for i, d := range data {
		_, err := NewGraphFrom(d.wrongData)
		if checks.NeqErrs(err, d.expectedErr) {
			t.Errorf("case [%v] expected err %v, got %v", i, d.expectedErr, err)
		}
	}
}

func TestGraphToString(t *testing.T) {
	expectedSimple := `5 4
1 2
1 3
2 3
3 4
5
`
	expectedDirected := `5 4 directed weighted hascodes
1  Bill Clinton
2  Hillary Clinton
3  John McCain
4  George Bush
5  Saddam Hussein
1 2 10
1 3 2
2 3 5
3 4 4
4
5
`
	expectedHasCodes := `5 4 hascodes
1  Bill Clinton
2  Hillary Clinton
3  John McCain
4  George Bush
5  Saddam Hussein
1 2
1 3
2 3
3 4
5
`
	data := []struct {
		text           string
		expectedString string
	}{
		{graphDataSimple, expectedSimple},
		{graphDataDirectedWeightedWithCodes, expectedDirected},
		{graphDataWithNamesNotDirectedWithoutWeight, expectedHasCodes},
	}
	for i, d := range data {
		g, err := NewGraphFrom(d.text)
		if err != nil {
			t.Errorf("case [%v] unexpected err %+v", i, err)
			t.FailNow()
		}
		out := g.ToString()
		if out != d.expectedString {
			t.Errorf("case [%v] expected %#v, got %#v", i, d.expectedString, out)
		}
	}

}

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

func caseStr(i int) string {
	return "case [" + strconv.Itoa(i) + "]"
}
