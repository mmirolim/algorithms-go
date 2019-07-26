package hk

func PrimsSpecialSubtree(n, start int32, edges [][]int32) int32 {
	inTree := make([]bool, n+1)
	mapOfEdges := make(map[int32][][2]int32, n)
	for _, e := range edges {
		mapOfEdges[e[0]] = append(mapOfEdges[e[0]], [2]int32{e[1], e[2]})
		// undirected graph
		mapOfEdges[e[1]] = append(mapOfEdges[e[1]], [2]int32{e[0], e[2]})
	}

	var minqueue [][2]int32
	put := func(e [2]int32) {
		minqueue = append(minqueue, e)
	}
	top := func() [2]int32 {
		var min int32 = 1000001
		id := 0
		var res [2]int32
		for i, d := range minqueue {
			if min > d[1] {
				min = d[1]
				id = i
				res = d
			}
		}
		minqueue[0], minqueue[id] = minqueue[id], minqueue[0]
		minqueue = minqueue[1:]
		return res
	}
	var sum int32

	bfs := func() {
		for len(minqueue) > 0 {
			cur := top()
			if !inTree[cur[0]] {
				inTree[cur[0]] = true
				sum += cur[1]
			}
			for _, e := range mapOfEdges[cur[0]] {
				if !inTree[e[0]] {
					put(e)
				}
			}
		}
	}
	for _, e := range mapOfEdges[start] {
		put(e)
	}
	inTree[start] = true
	bfs()
	return sum
}
