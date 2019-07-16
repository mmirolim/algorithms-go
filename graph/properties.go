package graph

import "errors"

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

	err := g.DFS(start, processEarly, processLate, processEdge, &parent)
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
