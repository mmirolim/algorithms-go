package graph

import (
	"errors"
	"fmt"
)

var ErrFinished = errors.New("finished")

func (g *Graph) FindCycle(start int) ([]int, error) {
	processEarly := func(v int) error { return nil }
	processLate := func(v int) error { return nil }
	parent := make([]int, g.NumOfVertices()+1)
	var path []int
	processEdge := func(c, n int) error {
		var err error
		if parent[n] != c {
			path, err = FindPath(n, c, parent)
			if err != nil {
				return err
			}
			return ErrFinished
		}
		return nil
	}

	err := g.DFS(start, processEarly, processLate, processEdge,
		&parent, nil, nil, nil, nil)
	if err == ErrFinished {
		return path, nil
	}
	return nil, err
}

// TODO handle directed graphs
func (g *Graph) Components() (components [][]int) {
	discovered := make([]bool, g.numOfVertices+1)
	var component []int
	addToComponent := func(v int) {
		discovered[v] = true
		component = append(component, v)
	}
	for i := 1; i <= g.numOfVertices; i++ {
		if !discovered[i] {
			g.BFS(i, addToComponent, nil, nil, nil)
			components = append(components, make([]int, len(component)))
			copy(components[len(components)-1][0:], component)
			component = component[:0]
		}
	}
	return
}

func (g *Graph) Bipartite() ([2][]int, bool) {
	var out [2][]int
	bipartite := true
	uncolored, white, black := 0, 1, 2
	discovered := make([]bool, g.numOfVertices+1)
	color := make([]int, g.numOfVertices+1)
	complement := func(c int) int {
		if c == white {
			return black
		} else if c == black {
			return white
		}
		return uncolored
	}
	processEarly := func(v int) {
		discovered[v] = true
	}
	processEdge := func(v1, v2 int) {
		if color[v1] == color[v2] {
			// same color
			bipartite = false
		}
		color[v2] = complement(color[v1])
	}
	for i := 1; i <= g.numOfVertices; i++ {
		if !discovered[i] {
			color[i] = white
			g.BFS(i, processEarly, nil, processEdge, nil)
			if !bipartite {
				break
			}
		}
	}
	if bipartite {
		for i := 1; i <= g.numOfVertices; i++ {
			if color[i] == uncolored {
				panic("uncolored found")
			}
			out[color[i]-1] = append(out[color[i]-1], i)
		}
	}
	return out, bipartite
}

func FindPath(start, end int, parent []int) ([]int, error) {
	var path []int
	if end > len(parent)-1 {
		return nil, errors.New("no end vertices in path")
	}
	count := 0
	for start != end {
		path = append(path, end)
		end = parent[end]
		count++
		if count > len(parent) {
			return nil, errors.New("path not found")
		}
	}
	path = append(path, start)
	return path, nil
}

type edgeType int

var (
	Unclassifiededge edgeType = 0
	TreeEdge         edgeType = 1
	BackEdge         edgeType = 2
	ForwardEdge      edgeType = 3
	CrossEdge        edgeType = 4
)

func edgeClassification(
	v1, v2 int,
	discovered []bool,
	processed []bool,
	parent []int,
	entryTime []int) edgeType {
	if parent[v2] == v1 {
		return TreeEdge
	}
	if discovered[v2] && !processed[v2] {
		return BackEdge
	}
	if processed[v2] && entryTime[v2] > entryTime[v1] {
		return ForwardEdge
	}

	if processed[v2] && entryTime[v2] < entryTime[v1] {
		return CrossEdge
	}
	fmt.Println("Warning: unclassified edge", v1, v2) // output for debug
	return Unclassifiededge
}

func (g *Graph) FindArticulations() ([][]int, error) {
	if g.isDirected {
		return nil, errors.New("error graph is directed")
	}
	var articulations []int
	reachableAncestor := g.vertexStorage()
	treeOutDegree := g.vertexStorage()
	discovered := g.vertexFlagStorage()
	processed := g.vertexFlagStorage()
	entryTime := g.vertexStorage()
	parent := g.vertexStorage()

	processEarly := func(v int) error {
		reachableAncestor[v] = v
		return nil
	}
	processEdge := func(v1, v2 int) error {
		edgeClass := edgeClassification(v1, v2, discovered, processed, parent, entryTime)
		if edgeClass == TreeEdge {
			treeOutDegree[v1]++
		}
		if edgeClass == BackEdge && parent[v1] != v2 {
			if entryTime[v2] < entryTime[reachableAncestor[v1]] {
				reachableAncestor[v1] = v2
			}
		}
		if edgeClass == Unclassifiededge {
			return errors.New("error unclassified edge")
		}
		return nil
	}
	processLate := func(v int) error {
		// is it root vertex
		isRoot := false
		// earliest entry time reachable vertex
		var earliestReachableTime int
		var earliestReachableParentTime int
		if parent[v] < 1 {
			// possible root
			if treeOutDegree[v] > 1 {
				// Root articulation vertex
				articulations = append(articulations, v)
			}
			return nil
		}
		isRoot = parent[parent[v]] < 1

		if reachableAncestor[v] == parent[v] && !isRoot {
			// Parent articulation vertex
			articulations = append(articulations, parent[v])
		}
		if reachableAncestor[v] == v {
			// Bridge articulation vertex
			articulations = append(articulations, parent[v])
			if treeOutDegree[v] > 0 {
				// vertex is a list ?
				// Bridge articulation vertex
				articulations = append(articulations, v)
			}
		}
		earliestReachableTime = entryTime[reachableAncestor[v]]
		earliestReachableParentTime = entryTime[reachableAncestor[parent[v]]]
		if earliestReachableTime < earliestReachableParentTime {
			reachableAncestor[parent[v]] = reachableAncestor[v]
		}
		return nil
	}
	var componentArticulations [][]int
	for _, v := range g.Vertices() {
		if !discovered[v] {
			g.DFS(v, processEarly, processLate, processEdge,
				&parent, &entryTime, nil,
				&discovered, &processed)

			if len(articulations) > 0 {
				componentArticulations = append(componentArticulations, articulations)
				articulations = make([]int, 0, len(articulations))
			}

		}
	}
	return componentArticulations, nil
}

func (g *Graph) TopologicalSort() ([]int, error) {
	if !g.isDirected {
		return nil, errors.New("error graph is not directed")
	}
	var sorted []int
	processLate := func(v int) error {
		sorted = append(sorted, v)
		return nil
	}
	discovered, processed := g.vertexFlagStorage(), g.vertexFlagStorage()
	parent, entryTime := g.vertexStorage(), g.vertexStorage()
	processEdge := func(v1, v2 int) error {
		if BackEdge == edgeClassification(v1, v2, discovered, processed, parent, entryTime) {
			return errors.New("error directed cycle found, not a DAG")

		}
		return nil
	}
	for _, v := range g.Vertices() {
		if !discovered[v] {
			err := g.DFS(v, nil, processLate, processEdge,
				&parent, &entryTime, nil, &discovered, &processed)
			if err != nil {
				return nil, err
			}
			for i := range entryTime {
				entryTime[i] = 0
				parent[i] = 0
			}
		}
	}

	return sorted, nil
}

func (g *Graph) StrongComponents() []int {
	if !g.isDirected {
		return nil
	}
	// strong component number for each vertex
	scc := g.vertexStorage()
	// oldest vertex surely in component of v
	low := g.vertexStorage()
	for i := range low {
		low[i] = i
	}
	discovered, processed := g.vertexFlagStorage(), g.vertexFlagStorage()
	parent, entryTime := g.vertexStorage(), g.vertexStorage()
	countComponents := 0
	var stack []int
	processEarly := func(v int) error {
		push(&stack, v)
		return nil
	}
	popComponent := func(v int) {
		countComponents++
		scc[v] = countComponents
		for {
			t := pop(&stack)
			if t == v {
				break
			}
			scc[t] = countComponents
		}
	}

	processlate := func(v int) error {
		if low[v] == v {
			// edge parent[v], v cuts off scc
			popComponent(v)
		}
		if entryTime[low[v]] < entryTime[low[parent[v]]] {
			low[parent[v]] = low[v]
		}
		return nil
	}

	processEdge := func(v1, v2 int) error {
		class := edgeClassification(v1, v2, discovered, processed, parent, entryTime)

		if class == BackEdge {
			if entryTime[v2] < entryTime[low[v1]] {
				low[v1] = v2
			}
		}
		if class == CrossEdge {
			if scc[v2] == 0 { // component not yet assigned
				if entryTime[v2] < entryTime[low[v1]] {
					low[v1] = v2
				}
			}
		}
		return nil
	}
	for _, v := range g.Vertices() {
		if !discovered[v] {
			g.DFS(v,
				processEarly, processlate, processEdge,
				&parent, &entryTime, nil, &discovered, &processed)
		}
	}
	return scc
}
