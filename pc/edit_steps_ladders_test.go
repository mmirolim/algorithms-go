package pc

import (
	"sort"
	"strconv"
	"testing"

	"github.com/mmirolim/algos/checks"
)

/*
9.6.5
Edit Step Ladders
PC/UVa IDs: 110905/10029, Popularity: B, Success rate: low Level: 3
*/
func TestEditStepLadders(t *testing.T) {
	dic1 := []string{"cat", "dig", "dog", "fig", "fin", "fine", "fog", "log", "wine"}
	dic2 := []string{"age", "ale", "ate", "bag", "big", "cat", "date", "dig", "dog", "fig", "fin", "fine", "fog", "line", "log", "swine", "wine"}
	data := []struct {
		dic []string
		res int
	}{
		{dic1, 5}, {dic2, 8},
	}
	for i, d := range data {
		out := EditStepLadders(d.dic)
		checks.AssertEq(t, d.res, out, caseStr(i))
	}
}

func EditStepLadders(sortedWords []string) int {
	// sorted list of words
	sort.Strings(sortedWords)
	// get all edges which connects to reachable words
	findEdges := func(order int, w string) []string {
		var out []string
		// in lexicographic order
		for i := order; i < len(sortedWords); i++ {
			countDiff := 0
			// possible edge
			if len(sortedWords[i]) == len(w) {
				for j := range w {
					if w[j] != sortedWords[i][j] {
						countDiff++
						if countDiff > 1 {
							break
						}
					}
				}
				if countDiff == 1 {
					out = append(out, sortedWords[i])
				}

			} else if abs(len(sortedWords[i]), len(w)) == 1 {
				for j := range w {
					if j < len(sortedWords[i]) {
						if w[j] != sortedWords[i][j] {
							countDiff++
							if countDiff > 0 {
								break
							}
						}
					} else {
						break
					}
				}
				if countDiff == 0 {
					out = append(out, sortedWords[i])
				}
			}
		}
		return out
	}
	mapEdges := map[string][]string{}
	for i, w := range sortedWords {
		mapEdges[w] = findEdges(i, w)
	}
	maxLadder := 0
	var stack []string
	longestDistanceFrom := map[string]int{}
	// go DFS and store distance of explored paths
	dfs := func(w string) {
		pushStr(&stack, w)
	OUTER:
		for len(stack) > 0 {
			cur := peekStr(&stack)
			maxDist := 0
			for _, next := range mapEdges[cur] {
				if longestDistanceFrom[next] == 0 {
					pushStr(&stack, next)
					continue OUTER
				} else {
					if longestDistanceFrom[next] > maxDist {
						maxDist = longestDistanceFrom[next]
					}
				}
			}
			longestDistanceFrom[cur] = maxDist + 1
			if maxLadder < longestDistanceFrom[cur] {
				maxLadder = longestDistanceFrom[cur]
			}
			popStr(&stack)
		}
	}
	// explore all components
	for i := range sortedWords {
		if longestDistanceFrom[sortedWords[i]] == 0 {
			dfs(sortedWords[i])
		}
	}
	return maxLadder
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}
func pushStr(stack *[]string, str string) {
	*stack = append(*stack, str)
}
func peekStr(stack *[]string) string {
	return (*stack)[len(*stack)-1]
}

func popStr(stack *[]string) string {
	s := *stack
	str := s[len(s)-1]
	s = s[:len(s)-1]
	*stack = s
	return str
}

func caseStr(i int) string {
	return "case [" + strconv.Itoa(i) + "]"
}
