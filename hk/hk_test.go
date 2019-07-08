package hk

import (
	"strconv"
	"testing"

	"github.com/mmirolim/algos/checks"
)

func caseStr(i int) string {
	return "case [" + strconv.Itoa(i) + "]"
}

func TestFormingMagicSquare(t *testing.T) {
	data := []struct {
		in   [][]int
		out  [][]int
		cost int
	}{
		{
			in: [][]int{
				[]int{5, 3, 4},
				[]int{1, 5, 8},
				[]int{6, 4, 2},
			},
			out: [][]int{
				[]int{8, 3, 4},
				[]int{1, 5, 9},
				[]int{6, 7, 2},
			},
			cost: 7,
		},
		{
			in: [][]int{
				[]int{4, 5, 8},
				[]int{2, 4, 1},
				[]int{1, 9, 7},
			},
			out:  [][]int{},
			cost: 14,
		},
		{
			in: [][]int{
				[]int{2, 5, 4},
				[]int{4, 6, 9},
				[]int{4, 5, 2},
			},
			out:  [][]int{},
			cost: 16,
		},
	}
	for i, d := range data {
		msg := caseStr(i)
		cost, _ := FormingMagicSquare(d.in)
		checks.AssertEq(t, d.cost, cost, msg)
	}
}

func TestPermutateNums(t *testing.T) {
	data := []struct {
		n   int
		out [][]int
	}{
		{2, [][]int{[]int{1, 2}, []int{2, 1}}},
		{3, [][]int{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 2, 1},
			[]int{3, 1, 2},
		},
		},
	}

	for i, d := range data {
		out := Permutate(d.n)
		checks.AssertEq(t, d.out, out, caseStr(i))
	}
}

func TestPermutateNumsHeapsAlgorithm(t *testing.T) {
	data := []struct {
		n   int
		out [][]int
	}{
		{2, [][]int{[]int{1, 2}, []int{2, 1}}},
		{3, [][]int{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			[]int{3, 1, 2},
			[]int{1, 3, 2},
			[]int{2, 3, 1},
			[]int{3, 2, 1},
		},
		},
	}

	for i, d := range data {
		out := PermutateHeapsAlgorithm(d.n)
		checks.AssertEq(t, d.out, out, caseStr(i))
	}
}

func TestQueensAttack2(t *testing.T) {
	data := []struct {
		size int
		qpos [2]int
		obs  [][2]int
		out  int
	}{
		{size: 1, qpos: [2]int{1, 1}, obs: nil, out: 0},
		{size: 4, qpos: [2]int{4, 4}, obs: nil, out: 9},
		{size: 5, qpos: [2]int{4, 3}, obs: [][2]int{[2]int{5, 5}, [2]int{4, 2}, [2]int{2, 3}}, out: 10},
		{size: 100000, qpos: [2]int{4187, 5068}, obs: nil, out: 308369},
		{size: 8, qpos: [2]int{1, 1}, obs: nil, out: 21},
		{size: 100, qpos: [2]int{48, 81}, obs: queensAttack2Case6, out: 40},
	}
	for i, d := range data {
		out := QueensAttack2(d.size, d.qpos, d.obs)
		checks.AssertEq(t, d.out, out, caseStr(i))

	}
}

func TestEncryption(t *testing.T) {
	data := []struct {
		in, out string
	}{
		{"haveaniceday", "hae and via ecy"},
		{"feedthedog", "fto ehg ee dd"},
		{"chillout", "clu hlt io"},
	}
	for i, d := range data {
		out := Encryption(d.in)
		checks.AssertEq(t, d.out, out, caseStr(i))
	}
}

var (
	queensAttack2Case6 = [][2]int{
		[2]int{54, 87},
		[2]int{64, 97},
		[2]int{42, 75},
		[2]int{32, 65},
		[2]int{42, 87},
		[2]int{32, 97},
		[2]int{54, 75},
		[2]int{64, 65},
		[2]int{48, 87},
		[2]int{48, 75},
		[2]int{54, 81},
		[2]int{42, 81},
		[2]int{45, 17},
		[2]int{14, 24},
		[2]int{35, 15},
		[2]int{95, 64},
		[2]int{63, 87},
		[2]int{25, 72},
		[2]int{71, 38},
		[2]int{96, 97},
		[2]int{16, 30},
		[2]int{60, 34},
		[2]int{31, 67},
		[2]int{26, 82},
		[2]int{20, 93},
		[2]int{81, 38},
		[2]int{51, 94},
		[2]int{75, 41},
		[2]int{79, 84},
		[2]int{79, 65},
		[2]int{76, 80},
		[2]int{52, 87},
		[2]int{81, 54},
		[2]int{89, 52},
		[2]int{20, 31},
		[2]int{10, 41},
		[2]int{32, 73},
		[2]int{83, 98},
		[2]int{87, 61},
		[2]int{82, 52},
		[2]int{80, 64},
		[2]int{82, 46},
		[2]int{49, 21},
		[2]int{73, 86},
		[2]int{37, 70},
		[2]int{43, 12},
		[2]int{94, 28},
		[2]int{10, 93},
		[2]int{52, 25},
		[2]int{50, 61},
		[2]int{52, 68},
		[2]int{52, 23},
		[2]int{60, 91},
		[2]int{79, 17},
		[2]int{93, 82},
		[2]int{12, 18},
		[2]int{75, 64},
		[2]int{69, 69},
		[2]int{94, 74},
		[2]int{61, 61},
		[2]int{46, 57},
		[2]int{67, 45},
		[2]int{96, 64},
		[2]int{83, 89},
		[2]int{58, 87},
		[2]int{76, 53},
		[2]int{79, 21},
		[2]int{94, 70},
		[2]int{16, 10},
		[2]int{50, 82},
		[2]int{92, 20},
		[2]int{40, 51},
		[2]int{49, 28},
		[2]int{51, 82},
		[2]int{35, 16},
		[2]int{15, 86},
		[2]int{78, 89},
		[2]int{41, 98},
		[2]int{70, 46},
		[2]int{79, 79},
		[2]int{24, 40},
		[2]int{91, 13},
		[2]int{59, 73},
		[2]int{35, 32},
		[2]int{40, 31},
		[2]int{14, 31},
		[2]int{71, 35},
		[2]int{96, 18},
		[2]int{27, 39},
		[2]int{28, 38},
		[2]int{41, 36},
		[2]int{31, 63},
		[2]int{52, 48},
		[2]int{81, 25},
		[2]int{49, 90},
		[2]int{32, 65},
		[2]int{25, 45},
		[2]int{63, 94},
		[2]int{89, 50},
		[2]int{43, 41},
	}
)
