package hk

import "math"

// https://www.hackerrank.com/challenges/subset-component/problem
// memoization and profiling used to reduce allocation and redundant work
// slow some tests failed with timeout error
// the faster way to count components by checking for disjoint sets and bit operations
// check SubsetComponentCountUsingBitCounting
func SubsetComponentCountUsingGraphBFS(arr []int64) int {
	maxNumOfSubsets := int(math.Pow(2, float64(len(arr))))
	numOfComponents := 0
	getSubset := func(cfg int) []int64 {
		var subset []int64
		bin := uint64(cfg)
		for i := 0; i < len(arr); i++ {
			if bin&(1<<uint64(i)) > 0 {
				subset = append(subset, arr[i])
			}
		}
		return subset
	}

	var queue []int8
	enqueue := func(d int8) {
		queue = append(queue, d)
	}
	dequeue := func() int8 {
		d := queue[0]
		queue = queue[1:]
		return d
	}
	// store alledges for digit
	memo := map[int64][]int8{}
	memoAllEdges := map[int][][]int8{}
	countComponents := func(cfg int, set []int64) int {
		getEdges := func(d int64) {
			_, ok := memo[d]
			if ok {
				return
			}
			memo[d] = make([]int8, 0, 64)
			cfg := uint64(d)
			var i uint64
			for ; i < 64; i++ {
				if cfg&(1<<i) > 0 {
					memo[d] = append(memo[d], int8(i))
				}
			}
		}

		edges, ok := memoAllEdges[cfg-1]
		var allEdges [][]int8
		allEdges = make([][]int8, 64)
		if !ok {
			for _, d := range set {
				getEdges(d)
				nodes := memo[d]
				for i := 0; i < len(nodes)-1; i++ {
					allEdges[nodes[i]] = append(allEdges[nodes[i]], nodes[i+1])
					// undirected graph
					allEdges[nodes[i+1]] = append(allEdges[nodes[i+1]], nodes[i])
				}
			}
			memoAllEdges[cfg] = allEdges
			edges = allEdges
		} else {
			// we got already subset
			allEdges = edges
			getEdges(set[0])
			nodes := memo[set[0]]
			for i := 0; i < len(nodes)-1; i++ {
				allEdges[nodes[i]] = append(allEdges[nodes[i]], nodes[i+1])
				// undirected graph
				allEdges[nodes[i+1]] = append(allEdges[nodes[i+1]], nodes[i])
			}
			edges = allEdges
		}
		discovered := make([]bool, 64)
		bfs := func(start int8) {
			enqueue(start)
			discovered[start] = true
			for len(queue) > 0 {
				cur := dequeue()
				for _, e := range edges[cur] {
					if !discovered[e] {
						enqueue(e)
						discovered[e] = true
					}
				}
			}
		}
		var i int8
		components := 0
		for ; i < 64; i++ {
			if len(edges[i]) == 0 {
				components++
			} else if !discovered[i] {
				bfs(i)
				components++
			}
		}
		return components
	}
	for i := 0; i < maxNumOfSubsets; i++ {
		subset := getSubset(i)
		numOfComponents += countComponents(i, subset)
	}
	return numOfComponents
}

func SubsetComponentCountUsingBitCounting(arr []int64) int {
	maxNumOfSubsets := int(math.Pow(2, float64(len(arr))))
	numOfComponents := 0
	getSubset := func(cfg int) []int64 {
		var subset []int64
		bin := uint64(cfg)
		for i := 0; i < len(arr); i++ {
			if bin&(1<<uint64(i)) > 0 {
				subset = append(subset, arr[i])
			}
		}
		return subset
	}

	countComponents := func(set []int64) int {
		// disjoint components
		var comps []int64
		for i := 0; i < len(set); i++ {
			joined := false
			for j := i + 1; j < len(set); j++ {
				if set[i]&set[j] != 0 {
					set[j] |= set[i]
					joined = true
				}
			}
			if !joined {
				comps = append(comps, set[i])
			}
		}

		components := 64
		countOnes := func(d int64) int {
			count := 0
			bin := uint64(d)
			for i := 0; i < 64; i++ {
				if bin&(1<<uint64(i)) != 0 {
					count++
				}
			}
			return count
		}

		for _, d := range comps {
			components -= countOnes(d)
		}

		return components + len(comps)
	}
	for i := 0; i < maxNumOfSubsets; i++ {
		subset := getSubset(i)
		numOfComponents += countComponents(subset)
	}
	return numOfComponents
}
