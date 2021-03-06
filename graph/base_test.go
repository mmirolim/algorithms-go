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

func caseStr(i int) string {
	return "case [" + strconv.Itoa(i) + "]"
}
