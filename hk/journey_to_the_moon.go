package hk

// https://www.hackerrank.com/challenges/journey-to-the-moon
func JourneyToTheMoon(n int32, astronaut [][]int32) int {
	mapOfEdges := map[int32][]int32{}
	for i := range astronaut {
		edges := mapOfEdges[astronaut[i][0]]
		edges = append(edges, astronaut[i][1])
		mapOfEdges[astronaut[i][0]] = edges
		// undirected graph
		edges = mapOfEdges[astronaut[i][1]]
		edges = append(edges, astronaut[i][0])
		mapOfEdges[astronaut[i][1]] = edges
	}

	discovered := make([]bool, n)
	var queue []int32
	enqueue := func(v int32) {
		queue = append(queue, v)
	}
	dequeue := func() int32 {
		v := queue[0]
		queue = queue[1:]
		return v
	}
	var components [][]int32
	var component []int32
	bfs := func(start int32) {
		enqueue(start)
		discovered[start] = true
		for len(queue) > 0 {
			cur := dequeue()
			component = append(component, cur)
			for _, e := range mapOfEdges[cur] {
				if !discovered[e] {
					enqueue(e)
					discovered[e] = true
				}
			}
		}
	}
	countOnes := 0
	for i := 0; i < int(n); i++ {
		if !discovered[i] {
			// if has edges
			bfs(int32(i))
			if len(component) == 1 {
				countOnes++
			} else {
				comp := make([]int32, len(component))
				copy(comp[0:], component)
				components = append(components, comp)
			}
			component = component[:0]
		}
	}
	countPairs := 0
	lenOfComponents := 0
	for i := range components {
		lenOfComponents += len(components[i])
		for j := i + 1; j < len(components); j++ {
			countPairs += len(components[i]) * len(components[j])
		}
	}
	countPairs += lenOfComponents * countOnes
	combinationOfOnes := func(n int) int {
		// number of variations n!/(r!(n-r!))
		return n * (n - 1) / 2
	}
	countPairs += combinationOfOnes(countOnes)
	return countPairs
}
