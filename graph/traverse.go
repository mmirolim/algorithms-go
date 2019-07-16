package graph

import "errors"

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
		// if has edges
		if g.edges[cur] != nil {
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
		}
		processLate(cur)
		processed[cur] = true
	}
}

// entry/exitTime
func (g *Graph) DFS(
	start int,
	processEarly,
	processLate func(v int) error,
	processEdge func(v, u int) error,
	parent *[]int,
) error {

	processed := make([]bool, g.numOfVertices+1)
	discovered := make([]bool, g.numOfVertices+1)
	entryTime := make([]int, g.numOfVertices+1)
	exitTime := make([]int, g.numOfVertices+1)
	time := 0
	var queue []int
	var par []int
	if parent != nil {
		if len(*parent) != g.numOfVertices+1 {
			return errors.New("parent size should be number of vertices + 1")
		}
		par = *parent
	} else {
		par = make([]int, g.numOfVertices+1)
	}
	if processEarly == nil {
		processEarly = func(v int) error { /* noop */ return nil }
	}
	if processLate == nil {
		processLate = func(v int) error { /* noop */ return nil }
	}
	if processEdge == nil {
		processEdge = func(v, u int) error { /* noop */ return nil }
	}

	discovered[start] = true
	time++
	entryTime[start] = time
	processEarly(start)
	push(&queue, start)
OUTER:
	for len(queue) > 0 {
		cur := peek(&queue)
		if g.edges[cur] != nil {
			for edge := g.edges[cur].head; edge != nil; edge = edge.next {
				next := edge.id
				if par[cur] == next {
					// skip parent
					continue
				}
				if !discovered[next] {
					time++
					entryTime[next] = time
					discovered[next] = true
					par[next] = cur
					if err := processEarly(next); err != nil {
						return err
					}
					push(&queue, next)
					if err := processEdge(cur, next); err != nil {
						return err
					}
					continue OUTER
				} else if !processed[next] || (g.IsDirected() && cur != par[next]) {
					if err := processEdge(cur, next); err != nil {
						return err
					}
				}
			}
		}
		cur = pop(&queue)
		processLate(cur)
		time++
		exitTime[cur] = time
		processed[cur] = true
	}

	return nil
}
