package number

import (
	"math"
	"testing"
)

func TestLightMoreLight(t *testing.T) {
	data := []struct {
		n      int
		lastOn bool
	}{
		{3, false}, {6241, true}, {8191, false},
	}
	for i, d := range data {
		out := LightMoreLight(d.n)
		if out != d.lastOn {
			t.Errorf("case [%v] expected %v, got %v", i, d.lastOn, out)
		}
	}

}

func TestStanVsOllie(t *testing.T) {
	data := []struct {
		n        int
		stanWins bool
	}{
		{162, true}, {17, false}, {34012226, true},
		{324, false},
	}
	for i, d := range data {
		out := StanVsOllie(d.n)
		if out != d.stanWins {
			t.Errorf("case [%v] expected %v, got %v", i, d.stanWins, out)
		}
	}

}

func TestCountCarryInt(t *testing.T) {
	data := []struct {
		a, b int
		out  int
	}{
		{123, 456, 0}, {555, 555, 3}, {123, 594, 1},
		{9823, 290, 3}, {99, 1, 2},
	}
	for i, d := range data {
		out := CountCarryInt(d.a, d.b)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}

}
func TestCountCarry(t *testing.T) {
	data := []struct {
		a, b string
		out  int
	}{
		{"123", "456", 0}, {"555", "555", 3}, {"123", "594", 1},
		{"9823", "290", 3}, {"99", "1", 2},
	}
	for i, d := range data {
		out := CountCarry(d.a, d.b)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}

}
func TestCountCarryIter(t *testing.T) {
	data := []struct {
		a, b string
		out  int
	}{
		{"123", "456", 0}, {"555", "555", 3}, {"123", "594", 1},
		{"9823", "290", 3}, {"99", "1", 2},
	}
	for i, d := range data {
		out := CountCarryIter(d.a, d.b)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}

}

func TestFindDropsKMarbleBreaksInNFloorBuilding(t *testing.T) {
	data := []struct {
		N, K int
		out  int
	}{
		{100, 1, 100}, {100, 2, 14}, {100, 3, 9}, {56, 4, 6},
	}
	for i, d := range data {
		out := FindDropsKMarbleBreaksInNFloorBuilding(d.N, d.K)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}
}

func TestStairCaseProblemDPSolution(t *testing.T) {
	data := []struct {
		tsteps int
		steps  []int
		out    int
	}{{2, []int{1, 2}, 2}, {4, []int{1, 2}, 5}}

	for i, d := range data {
		out := StairCaseProblemDPSolution(d.tsteps, d.steps)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}
}

func TestStairCaseProblem(t *testing.T) {
	data := []struct {
		tsteps int
		steps  []int
		out    int
	}{{2, []int{1, 2}, 2}, {4, []int{1, 2}, 5}}

	for i, d := range data {
		out := StairCaseProblem(d.tsteps, d.steps)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}
}

func TestStairCaseProblemOrderDoesNotMatter(t *testing.T) {
	data := []struct {
		tsteps int
		out    int
	}{{2, 2}, {4, 3}, {5, 3}, {6, 4}}

	for i, d := range data {
		out := StairCaseProblemOrderDoestNotMatter(d.tsteps)
		if out != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, out)
		}
	}
}

func TestDelannoyNumber(t *testing.T) {
	data := []struct {
		n, m, out int
	}{
		{3, 3, 63}, {4, 5, 681},
	}

	for i, d := range data {
		res := DelannoyNumber(d.n, d.m)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestDelannoyNumberDP(t *testing.T) {
	data := []struct {
		n, m, out int
	}{
		{3, 3, 63}, {4, 5, 681},
	}

	for i, d := range data {
		res := DelannoyNumberDP(d.n, d.m)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestConvertNumberToWords(t *testing.T) {
	data := []struct {
		num, out string
	}{
		{"9", "nine"},
		{"11", "eleven"},
		{"119", "one hundred and nineteen"},
		{"100", "one hundred"},
		{"578", "five hundred and seventy eight"},
		{"9923", "nine thousand nine hundred and twenty three"},
		{"19008", "nineteen thousand and eight"},
		{"120411", "one hundred twenty thousand four hundred and eleven"},
		{"916120410", "nine hundred sixteen million one hundred twenty thousand four hundred and ten"},
	}

	for i, d := range data {
		res := ConvertNumberToWords(d.num)
		if res != d.out {
			t.Errorf("case [%v] num %v expected >%v<, got >%v<", i, d.num, d.out, res)
		}
	}
}

func TestADM_3_28(t *testing.T) {
	data := []struct {
		X, M []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{720, 360, 240, 180, 144, 120}},
	}

	for i, d := range data {
		res := ADM_3_28(d.X)
		if len(d.M) != len(res) {
			t.Errorf("case [%v] expected %v, got %v", i, len(d.M), res)
		}
		for k := 0; k < len(d.M); k++ {
			if res[k] != d.M[k] {
				t.Errorf("case [%v] index %v expected %v, got %v", i, k, d.M[k], res[k])
			}
		}
	}
}

func TestIsJollyJumpers(t *testing.T) {
	data := []struct {
		seq []int
		out bool
	}{
		{[]int{1}, true},
		{[]int{5, 1, 4, 2, -1, 6}, false},
		{[]int{11, 7, 4, 2, 1, 6}, true},
		{[]int{1, 4, 2, 3}, true},
	}

	for i, d := range data {
		res := IsJollyJumpers(d.seq)
		if res != d.out {
			t.Errorf("case [%v] expected %+v, got %+v", i, d.out, res)
		}
	}
}

func TestAbs(t *testing.T) {
	for i, d := range []int{-10, 23, -332, 12} {
		if int(math.Abs(float64(d))) != abs(d) {
			t.Errorf("case [%v] expected %+v, got %+v", i, math.Abs(float64(d)), abs(d))
		}
	}
}

func TestFindTimeWhenAlexeyWillWakeUp(t *testing.T) {
	data := []struct {
		N, X, K int
		alarms  []int
		out     int
	}{
		{6, 5, 10, []int{1, 2, 3, 4, 5, 6}, 10},
		{5, 7, 12, []int{5, 22, 17, 13, 8}, 27},
	}
	for i, d := range data {
		res := FindTimeWhenAlexeyWillWakeUp(d.N, d.X, d.K, d.alarms)
		if res != d.out {
			t.Errorf("case [%v] expected %+v, got %+v", i, d.out, res)
		}
	}

}
