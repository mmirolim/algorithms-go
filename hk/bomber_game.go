package hk

import "fmt"

// https://www.hackerrank.com/challenges/bomber-man
type BomberGame struct {
	state [][]int
}

func NewBomberGame(grid []string) *BomberGame {
	g := new(BomberGame)
	g.state = make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		g.state[i] = make([]int, len(grid[0]))
	}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 'O' {
				g.state[r][c] = 3
			}
		}
	}
	return g
}

func (g *BomberGame) ToString() []string {
	out := make([]string, len(g.state))
	buf := make([]byte, len(g.state[0]))
	for r := 0; r < len(g.state); r++ {
		for c := 0; c < len(g.state[0]); c++ {
			if g.state[r][c] == 0 {
				buf[c] = '.'
			} else {
				buf[c] = 'O'
			}
		}
		out[r] = string(buf)
	}
	return out
}

func (g *BomberGame) SetBombs(t int) {
	for r := 0; r < len(g.state); r++ {
		for c := 0; c < len(g.state[0]); c++ {
			if g.state[r][c] == 0 {
				g.state[r][c] = t + 3
			}
		}
	}
}

func (g *BomberGame) DetonateBombs(t int) {
	for r := 0; r < len(g.state); r++ {
		for c := 0; c < len(g.state[0]); c++ {
			if g.state[r][c] == t {
				g.Detonate(r, c)
			}
		}
	}
}
func (g *BomberGame) Detonate(r, c int) {
	bombTime := g.state[r][c]
	if r-1 > -1 {
		if g.state[r-1][c] != bombTime {
			g.state[r-1][c] = 0
		}
	}
	if r+1 < len(g.state) {
		if g.state[r+1][c] != bombTime {
			g.state[r+1][c] = 0
		}
	}
	if c-1 > -1 {
		if g.state[r][c-1] != bombTime {
			g.state[r][c-1] = 0
		}
	}
	if c+1 < len(g.state[0]) {
		if g.state[r][c+1] != bombTime {
			g.state[r][c+1] = 0
		}
	}

	g.state[r][c] = 0
}

func TheBombermanGame(n int, grid []string) []string {
	game := NewBomberGame(grid)
	setBombState := game.ToString()
	var detonated1, detonated2 []string
	limit := n
	if n > 5 {
		limit = 5
	}
	t := 2
	for t <= limit {
		if t%2 == 0 {
			game.SetBombs(t)
			setBombState = game.ToString()
		}
		if t == 3 {
			game.DetonateBombs(t)
			detonated1 = game.ToString()
		}
		if t == 5 {
			game.DetonateBombs(t)
			detonated2 = game.ToString()
		}
		t++
	}
	if n < 5 {
		return game.ToString()
	}
	if n%2 == 0 {
		return setBombState
	}
	// check if a0 of arithmetic progression is 3
	if n+1 == ((n+1)/4)*4 {
		return detonated1
	}

	return detonated2
}

func (g *BomberGame) printGrid(t int) {
	fmt.Printf("Time %d\n", t) // output for debug
	for _, str := range g.ToString() {
		fmt.Printf("%+v\n", str) // output for debug

	}
}
