package hk

// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
// testcase 6, 8, 9 failing (n 200000) with runtime error
// TODO fix
func ClimbingTheLeaderboard(board, scores []int) (ranks []int) {
	createRankings := func(board []int) []int {
		var out []int
		out = append(out, board[0])
		prev := out[0]
		for i := 1; i < len(board); i++ {
			if board[i] != prev {
				out = append(out, board[i])
			}
			prev = board[i]
		}
		return out
	}

	scoresToRanks := createRankings(board)

	// recursive solution does not work with big score list
	binarySearch := func(l, r, score int, scores []int) (rank int) {
		for l <= r {
			m := l + (r-l)/2
			if scores[m] == score {
				return m
			} else if scores[m] < score && score < scores[m-1] {
				return m
			} else if scores[m] > score && score >= scores[m+1] {
				return m + 1
			} else if scores[m] > score {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		return
	}
	id := len(scoresToRanks) - 1
	for _, s := range scores {
		rank := -1
		if s < scoresToRanks[len(scoresToRanks)-1] {
			rank = len(scoresToRanks) + 1
		} else if s > scoresToRanks[0] {
			rank = 1
		} else {
			id = binarySearch(0, id, s, scoresToRanks)
			rank = id + 1 // zero baseed index
			if s < scoresToRanks[id] {
				rank++
			}
		}
		ranks = append(ranks, rank)
	}
	return
}
