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

func TestBiggerIsGreater(t *testing.T) {
	data := []struct {
		in, out []string
	}{
		{
			[]string{"ab", "bb", "hefg", "dhck", "dkhc"},
			[]string{"ba", "no answer", "hegf", "dhkc", "hcdk"}},
		{
			[]string{"lmno", "dcba", "dcbb", "abdc", "abcd", "fedcbabcd"},
			[]string{"lmon", "no answer", "no answer", "acbd", "abdc", "fedcbabdc"},
		},
		{
			[]string{"gwakhcpkolybihkmxyecrdhsvycjrljajlmlqgpcnmvvkjlkvdowzdfikh"},
			[]string{"gwakhcpkolybihkmxyecrdhsvycjrljajlmlqgpcnmvvkjlkvdowzdfkhi"},
		},
	}
	for i, d := range data {
		out := BiggerIsGreater(d.in)
		checks.AssertEq(t, d.out, out, caseStr(i))
	}
}

func TestClimbingTheLeaderboard(t *testing.T) {
	data := []struct {
		board  []int
		scores []int
		ranks  []int
	}{
		{
			[]int{100, 100, 50, 40, 40, 20, 10},
			[]int{5, 25, 50, 120},
			[]int{6, 4, 2, 1},
		},
		{
			[]int{100, 90, 90, 80, 75, 60},
			[]int{50, 65, 77, 90, 102},
			[]int{6, 5, 4, 2, 1},
		},
	}
	for i, d := range data {
		out := ClimbingTheLeaderboard(d.board, d.scores)
		checks.AssertEq(t, d.ranks, out, caseStr(i))
	}
}

func TestTimeInWords(t *testing.T) {
	data := []struct {
		// hour and minutes
		hour, minutes int32
		str           string
	}{
		{5, 47, "thirteen minutes to six"},
		{3, 0, "three o' clock"},
		{7, 15, "quarter past seven"},
		{5, 1, "one minute past five"},
		{11, 28, "twenty eight minutes past eleven"},
		{5, 30, "half past five"},
	}
	for i, d := range data {
		out := TimeInWords(d.hour, d.minutes)
		checks.AssertEq(t, d.str, out, caseStr(i))
	}
}

func TestOrganizingContainersOfBalls(t *testing.T) {
	data := []struct {
		cons     [][]int32
		possible bool
	}{
		{cons: [][]int32{[]int32{1, 1}, []int32{1, 1}}, possible: true},
		{cons: [][]int32{[]int32{0, 2}, []int32{1, 1}}, possible: false},
		{cons: [][]int32{[]int32{1, 3, 1}, []int32{2, 1, 2}, []int32{3, 3, 3}},
			possible: false},
		{cons: [][]int32{[]int32{0, 2, 1}, []int32{1, 1, 1}, []int32{2, 0, 0}}, possible: true},
	}

	for i, d := range data {
		out := OrganizingContainersOfBalls(d.cons)
		checks.AssertEq(t, d.possible, out, caseStr(i))
	}
}

func TestTheGridSearch(t *testing.T) {
	data := []struct {
		grid, pattern []string
		found         bool
	}{
		{grid: []string{
			"1234567890",
			"0987654321",
			"1111111111",
			"1111111111",
			"2222222222",
		}, pattern: []string{"99", "99"}, found: false},
		{grid: []string{
			"1234567890",
			"0987654321",
			"1111111111",
			"1111111111",
			"2222222222",
		}, pattern: []string{"876543", "111111", "111111"}, found: true},
		{grid: []string{
			"7283455864",
			"6731158619",
			"8988242643",
			"3830589324",
			"2229505813",
			"5633845374",
			"6473530293",
			"7053106601",
			"0834282956",
			"4607924137",
		}, pattern: []string{"9505", "3845", "3530"}, found: true},
		{grid: []string{"999999", "121211"}, pattern: []string{"99", "11"}, found: true},
		{grid: []string{
			"111111111111111",
			"111111111111111",
			"111111011111111",
			"111111111111111",
			"111111111111111",
		}, pattern: []string{"11111", "11111", "11110"}, found: true},
	}

	for i, d := range data {
		out := TheGridSearch(d.grid, d.pattern)
		checks.AssertEq(t, d.found, out, caseStr(i))
	}
}

func Test3DSurfaceArea(t *testing.T) {
	data := []struct {
		data [][]int32
		area int
	}{
		{data: [][]int32{[]int32{1}}, area: 6},
		{data: [][]int32{
			[]int32{1, 3, 4},
			[]int32{2, 2, 3},
			[]int32{1, 2, 4},
		}, area: 60},
	}

	for i, d := range data {
		out := D3SurfaceArea(d.data)
		checks.AssertEq(t, d.area, out, caseStr(i))
	}
}

func TestAbsolutePermutation(t *testing.T) {
	data := []struct {
		n, k int
		perm []int
	}{
		{4, 2, []int{3, 4, 1, 2}}, {2, 1, []int{2, 1}}, {3, 0, []int{1, 2, 3}}, {3, 2, nil},
	}

	for i, d := range data {
		out := AbsolutePermutation(d.n, d.k)
		checks.AssertEq(t, d.perm, out, caseStr(i))
	}
}

func TestTheBombermanGame(t *testing.T) {
	data := []struct {
		n         int
		grid, out []string
	}{
		{n: 3, grid: []string{
			".......",
			"...O...",
			"....O..",
			".......",
			"OO.....",
			"OO.....",
		}, out: []string{
			"OOO.OOO",
			"OO...OO",
			"OOO...O",
			"..OO.OO",
			"...OOOO",
			"...OOOO",
		}},
		{n: 181054341, grid: []string{
			"O..OO........O..O........OO.O.OO.OO...O.....OOO...OO.O..OOOOO...O.O..O..O.O..OOO..O..O..O....O...O....O...O..O..O....O.O.O.O.....O.....OOOO..O......O.O.....OOO....OO....OO....O.O...O..OO....OO..O...O",
		},
			out: []string{
				"OOOOO........OOOO........OOOOOOOOOO...O.....OOO...OOOOOOOOOOO...OOOOOOOOOOOOOOOOOOOOOOOOO....O...O....O...OOOOOOO....OOOOOOO.....O.....OOOOOOO......OOO.....OOO....OO....OO....OOO...OOOOO....OOOOO...O",
			}},
		{n: 329973043, grid: []string{
			"OOOO.O.O...OOO.O.O........O.OOO.O.....OO..O..O...OOO....O.OOO....O...O....O..O.O.O.....OOOO.O...O....OO.O...........O.O..O.O..O...OO.OOO......O........O...O....O.O..O....O.......OOOO.O..........OO.O"}, out: []string{
			".........O.........OOOOOO.........OOO..........O.....OO.......OO...O...OO..........OOO........O...OO......OOOOOOOOO.............O........OOOO...OOOOOO...O...OO........OO...OOOOO........OOOOOOOO.....",
		}},
	}

	for i, d := range data {
		out := TheBombermanGame(d.n, d.grid)
		checks.AssertEq(t, d.out, out, caseStr(i))
	}
}

// TODO solve
func _TestStringTransmission(t *testing.T) {
	data := []struct {
		s      string
		k, res int
	}{
		{"00000", 0, 0},
		{"001", 1, 3},
		{"101", 3, 6},
		{"1111", 1, 4},
		{"1101011", 4, 98},
		{"011", 2, 5},
		{"1", 0, 1},
		{"101010", 3, 33},
		{"001101", 2, 19},
		{"0011110100", 5, 613},
		{"000000011010100011000101101111101101101000000101000000110100010000111001001110010010111101100100000100010000001011100010010001101111100100101100010001000001111001011011101011011011111001111001101000110010011111010110000100100000000001001010111100010001011011011000000100111110111011111010110100100111100111110011100001111101111110000110001100000100011100011110111000011101101110101010001010010110011101010001011111101100101100001111110110011011101011011001001000110001001000110010101001000100001001101000011111101011111100100000100101110000110011100101111100110100011101110101110100011001011011011001110010100110100100000101100001100101011011001011101111001000110010111110011100010100000101100101110101101011110100100000110110100110111010100100111111100001010001110010000010100011110110111101011110111100011010101110100010001110101011110011", 834, 623251959},
	}
	for i, d := range data {
		out := StringTransmission(d.k, d.s)
		checks.AssertEq(t, d.res, out, caseStr(i))
	}
}

func TestAlmostSorted(t *testing.T) {
	data := []struct {
		arr []int32
		res string
	}{
		{[]int32{4, 2}, "swap 1 2"}, {[]int32{3, 1, 2}, ""}, {[]int32{1, 5, 4, 3, 2, 6}, "reverse 2 5"},
		{[]int32{4, 2, 3, 1}, "swap 1 4"}, {[]int32{4, 3, 2, 1}, "reverse 1 4"},
		{[]int32{4, 3, 2, 1, 5}, "reverse 1 4"}, {[]int32{1, 2, 4, 3, 5, 6}, "swap 3 4"},
		{[]int32{1, 2, 3, 5, 4, 6}, "swap 4 5"}, {[]int32{43, 65, 1, 98, 99, 101}, ""},
	}
	for i, d := range data {
		out := AlmostSorted(d.arr)
		checks.AssertEq(t, d.res, out, caseStr(i))
	}
}

func TestEmasSupercomputer(t *testing.T) {
	data := []struct {
		grid []string
		res  int
	}{
		{grid: []string{
			"GGGGGG",
			"GBBBGB",
			"GGGGGG",
			"GGBBGB",
			"GGGGGG",
		}, res: 5},
		{grid: []string{
			"BGBBGB",
			"GGGGGG",
			"BGBBGB",
			"GGGGGG",
			"BGBBGB",
			"BGBBGB",
		}, res: 25},
		{grid: []string{
			"BBBBBGGBGG",
			"GGGGGGGGGG",
			"GGGGGGGGGG",
			"BBBBBGGBGG",
			"BBBBBGGBGG",
			"GGGGGGGGGG",
			"BBBBBGGBGG",
			"GGGGGGGGGG",
			"BBBBBGGBGG",
			"GGGGGGGGGG",
		}, res: 85},
		{grid: []string{
			"GGGGGGGGGGGG",
			"GBGGBBBBBBBG",
			"GBGGBBBBBBBG",
			"GGGGGGGGGGGG",
			"GGGGGGGGGGGG",
			"GGGGGGGGGGGG",
			"GGGGGGGGGGGG",
			"GBGGBBBBBBBG",
			"GBGGBBBBBBBG",
			"GBGGBBBBBBBG",
			"GGGGGGGGGGGG",
			"GBGGBBBBBBBG",
		}, res: 81},
		{grid: []string{
			"GGGGGGGGG",
			"GBBBGGBGG",
			"GBBBGGBGG",
			"GBBBGGBGG",
			"GBBBGGBGG",
			"GBBBGGBGG",
			"GBBBGGBGG",
			"GGGGGGGGG",
		}, res: 1},
	}
	for i, d := range data {
		out := EmasSupercomputer(d.grid)
		checks.AssertEq(t, d.res, out, caseStr(i))
	}

}

func TestCountLuck(t *testing.T) {
	data := []struct {
		grid []string
		k    int
		res  bool
	}{
		{grid: []string{
			"*.M",
			".X.",
		}, k: 1, res: true},
		{grid: []string{
			".X.X......X",
			".X*.X.XXX.X",
			".XX.X.XM...",
			"......XXXX.",
		}, k: 3, res: true},
		{grid: []string{
			".X.X......X",
			".X*.X.XXX.X",
			".XX.X.XM...",
			"......XXXX.",
		}, k: 4, res: false},
	}
	for i, d := range data {
		out := CountLuck(d.k, d.grid)
		checks.AssertEq(t, d.res, out, caseStr(i))
	}
}
