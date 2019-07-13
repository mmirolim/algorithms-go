package hk

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
	t := 2
	for t <= n {
		if t%2 == 0 {
			game.SetBombs(t)
		}
		if t&1 == 1 {
			game.DetonateBombs(t)
		}
		t++
	}
	return game.ToString()
}
