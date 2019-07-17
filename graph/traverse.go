package graph

import (
	"errors"
)

func (g *Graph) BFS(
	start int,
	processEarly, processLate func(v int),
	processEdge func(v, u int),
	parents *[]int,
) {
	// init
	discovered := g.vertexFlagStorage()
	processed := g.vertexFlagStorage()
	if parents != nil {
		*parents = g.vertexStorage()
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

func (g *Graph) DFS(
	start int,
	processEarly,
	processLate func(v int) error,
	processEdge func(v, u int) error,
	parent, entryTime, exitTime *[]int,
	discovered, processed *[]bool,
) error {
	time := 0
	var queue []int
	var par, entryT, exitT []int
	var procEd, disEd []bool

	for _, sl := range []struct {
		s1 *[]bool
		s2 *[]bool
	}{{&procEd, processed}, {&disEd, discovered}} {
		if sl.s2 != nil {
			if len(*sl.s2) != g.numOfVertices+1 {
				return errors.New("provided size should be number of vertices + 1")
			}
			*sl.s1 = *sl.s2
		} else {
			*sl.s1 = g.vertexFlagStorage()
		}
	}
	for _, sl := range []struct {
		s1 *[]int
		s2 *[]int
	}{{&par, parent}, {&entryT, entryTime}, {&exitT, exitTime}} {
		if sl.s2 != nil {
			if len(*sl.s2) != g.numOfVertices+1 {
				return errors.New("provided size should be number of vertices + 1")
			}
			*sl.s1 = *sl.s2
		} else {
			*sl.s1 = g.vertexStorage()
		}
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

	disEd[start] = true
	time++
	entryT[start] = time
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
				if !disEd[next] {
					time++
					entryT[next] = time
					disEd[next] = true
					par[next] = cur
					if err := processEarly(next); err != nil {
						return err
					}
					push(&queue, next)
					if err := processEdge(cur, next); err != nil {
						return err
					}
					continue OUTER
				} else if !procEd[next] || (g.IsDirected() && cur != par[next]) {
					if err := processEdge(cur, next); err != nil {
						return err
					}
				}
			}
		}
		cur = pop(&queue)
		processLate(cur)
		time++
		exitT[cur] = time
		procEd[cur] = true
	}

	return nil
}
