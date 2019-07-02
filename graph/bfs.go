package graph

func (g *Graph) BFS(
	start int,
	processEarly, processLate func(v int),
	processEdge func(v, u int),
	parents *[]int,
) {
	// init
	discovered := make([]bool, g.numOfVertices+1)
	processed := make([]bool, g.numOfVertices+1)
	if parents != nil {
		*parents = make([]int, g.numOfVertices+1)
	}
	if processEarly == nil {
		processEarly = func(v int) { /* noop */ }
	}
	if processLate == nil {
		processLate = func(v int) { /* noop */ }
	}
	if processEdge == nil {
		processEdge = func(v, u int) { /* noop */ }
	}

	discovered[start] = true
	// queue
	var q []int
	enqueue(&q, start)
	// current and next vertices
	var cur, next int
	var edge *node
	for len(q) > 0 {
		cur = dequeue(&q)
		processEarly(cur)
		processed[cur] = true
		// if no edges
		if g.edges[cur] == nil {
			processLate(cur)
			continue
		}

		// visit all edges
		for edge = g.edges[cur].head; edge != nil; edge = edge.next {
			next = edge.id
			if !processed[next] || g.isDirected {
				processEdge(cur, next)
			}
			if !discovered[next] {
				enqueue(&q, next)
				discovered[next] = true
				if parents != nil {
					(*parents)[next] = cur
				}

			}
		}
		processLate(cur)
	}
}
