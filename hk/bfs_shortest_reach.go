package hk

// https://www.hackerrank.com/challenges/bfsshortreach/problem
func BFSShortestReach(n int32, m int32, allEdges [][]int32, s int32) []int32 {
	var noEdge, edgeWeight int32 = -1, 6
	discovered := make([]bool, n+1)
	// undirected
	mapOfEdges := map[int32][]int32{}
	for i := range allEdges {
		edges := mapOfEdges[allEdges[i][0]]
		edges = append(edges, allEdges[i][1])
		mapOfEdges[allEdges[i][0]] = edges
		// undirected graph
		edges = mapOfEdges[allEdges[i][1]]
		edges = append(edges, allEdges[i][0])
		mapOfEdges[allEdges[i][1]] = edges
	}
	var queue []int32
	enqueue := func(v int32) {
		queue = append(queue, v)
	}
	dequeue := func() int32 {
		v := queue[0]
		queue = queue[1:]
		return v
	}
	parent := make([]int32, n+1)
	bfs := func(start int32) {
		enqueue(start)
		discovered[start] = true
		for len(queue) > 0 {
			cur := dequeue()
			for _, e := range mapOfEdges[cur] {
				if !discovered[e] {
					parent[e] = cur
					enqueue(e)
					discovered[e] = true
				}
			}
		}
	}
	pathLen := func(node int32, parent []int32) int32 {
		var l int32 = 1
		for parent[node] != s {
			node = parent[node]
			l++
		}
		return l
	}
	bfs(s)
	var shortestReach []int32
	for i := 1; i <= int(n); i++ {
		if i == int(s) {
			// skip start
			continue
		}
		if !discovered[i] {
			shortestReach = append(shortestReach, noEdge)
		} else {
			shortestReach = append(shortestReach, pathLen(int32(i), parent)*edgeWeight)
		}
	}

	return shortestReach
}
