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
