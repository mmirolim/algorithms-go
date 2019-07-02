package graph

var (
	// space separated vertex, edge directed keyword, weighted keyword, hascodes keyword
	// next line space separated vertices of edge and weight it exists
	graphDataWithNamesNotDirectedWithoutWeight = `
5 4 hascodes
1 Bill Clinton
2 Hillary Clinton
3 John McCain
4 George Bush
5 Saddam Hussein
1 2
1 3
2 3
3 4
5
`
	graphDataDirectedWeightedWithCodes = `
5 5 directed weighted hascodes
1 Bill Clinton
2 Hillary Clinton
3 John McCain
4 George Bush
5 Saddam Hussein
1 2 10
1 3 2
2 3 5
2 1 8
3 4 4
5
`
	graphDataSimple = `
5 4
1 2
1 3
2 3
3 4
5
`
	codes = map[int]string{
		1: "Bill Clinton",
		2: "Hillary Clinton",
		3: "John McCain",
		4: "George Bush",
		5: "Saddam Hussein",
	}
	edgesWithWeights = [][3]int{
		[3]int{1, 2, 10},
		[3]int{1, 3, 2},
		[3]int{2, 3, 5},
		[3]int{3, 4, 4},
	}
	edgesWithNoWeights = [][3]int{
		[3]int{1, 2, 0},
		[3]int{1, 3, 0},
		[3]int{2, 3, 0},
		[3]int{3, 4, 0},
	}
	justEdges = [][2]int{
		[2]int{1, 2},
		[2]int{1, 3},
		[2]int{2, 3},
		[2]int{3, 4},
	}
	vertices = []int{1, 2, 3, 4, 5}

	twoRhombus = `7 8
1 3
1 2
2 4
3 4
4 5
4 7
5 6
6 7
`
)
