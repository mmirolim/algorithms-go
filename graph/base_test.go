package graph

import (
	"testing"

	"github.com/mmirolim/algos/checks"
)

var (
	// space separated vertex, edge directed keyword, weighted keyword, hascodes keyword
	// next line space separated vertices of edge and weight it exists
	graphDataWithNamesNotDirectedWithoutWeight = `
5 4 hascodes
1 Bill Clinton
2 Hillary Clinton
3 John McCain
4 George Bush
5 Saddam Hussein
1 2
1 3
2 3
3 4
5
`
	graphDataDirectedWeightedWithCodes = `
5 4 directed weighted hascodes
1 Bill Clinton
2 Hillary Clinton
3 John McCain
4 George Bush
5 Saddam Hussein
1 2 10
1 3 2
2 3 5
3 4 4
5
`
	graphDataSimple = `
5 4
1 2
1 3
2 3
3 4
5
`
	codes = map[int]string{
		1: "Bill Clinton",
		2: "Hillary Clinton",
		3: "John McCain",
		4: "George Bush",
		5: "Saddam Hussein",
	}
	edgesWithWeights = [][3]int{
		[3]int{1, 2, 10},
		[3]int{1, 3, 2},
		[3]int{2, 3, 5},
		[3]int{3, 4, 4},
	}
	edgesWithNoWeights = [][3]int{
		[3]int{1, 2, 0},
		[3]int{1, 3, 0},
		[3]int{2, 3, 0},
		[3]int{3, 4, 0},
	}
	justEdges = [][2]int{
		[2]int{1, 2},
		[2]int{1, 3},
		[2]int{2, 3},
		[2]int{3, 4},
	}
	vertices = []int{1, 2, 3, 4, 5}
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
		if checks.NeqErrs(err, d.expectedErr) {
			t.Errorf("case [%v] expected err %v, got err %v",
				i, d.expectedErr, err)
			continue
		}
		vertices := g.Vertices()
		if len(vertices) != d.numOfVertices {
			t.Errorf("case [%v] expected vertices %v, got %v",
				i, d.numOfVertices, len(vertices))
		}
		for j, ver := range d.vertices {
			if ver != vertices[j] {
				t.Errorf("case [%v] expected vertices %v, got %v",
					i, d.vertices, vertices)
				break
			}
		}

		edgesFromGraph := g.Edges()
		if len(edgesFromGraph) != d.numOfEdges {
			t.Errorf("case [%v] expected edges %v, got %v", i, d.numOfEdges, len(edgesFromGraph))
		}
		for j, edge := range d.edges {
			if edge[0] != edgesFromGraph[j][0] || edge[1] != edgesFromGraph[j][1] {
				t.Errorf("case [%v] expected edge %v, got %v",
					i, d.edges, edgesFromGraph)
				break
			}
		}
		if g.IsDirected() != d.isDirected {
			t.Errorf("case [%v] expected directed %v, got %v",
				i, d.isDirected, g.IsDirected())
		}
		if g.IsWeighted() != d.isWeighted {
			t.Errorf("case [%v] expected weighted %v, got %v",
				i, d.isWeighted, g.IsWeighted())
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
	data := []struct {
		wrongData   string
		expectedErr error
	}{
		{wrongNumOfCodes, ErrWrongNumOfEdges},
		{wrongNumOfVertices, ErrWrongNumOfVertices},
		{weightMissing, ErrWeightMissing},
		{wrongNumOfEdges, ErrWrongNumOfEdges},
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
