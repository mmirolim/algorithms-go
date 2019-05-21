package strings

import (
	"sort"
	"testing"
)

func TestCountNumOfBinaryStringsNWithNoConsec1s(t *testing.T) {
	data := []struct {
		n, out int
	}{
		{2, 3},
		{3, 5},
		{4, 8},
	}

	for i, d := range data {
		res := CountNumOfBinaryStringsNWithNoConsec1s(d.n)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestCountNumOfBinaryStringsNWithAtLeastKAdjacent1Bits(t *testing.T) {
	data := []struct {
		n, k, out int
	}{
		{3, 3, 1},
		{4, 3, 3},
	}

	for i, d := range data {
		res := CountNumOfBinaryStringsNWithAtLeastKAdjacent1Bits(d.n, d.k)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestNumOfBinaryStringsNWithAtLeastKAdjacent1Bits(t *testing.T) {
	data := []struct {
		n, k int
		out  []string
	}{
		{3, 3, []string{"111"}},
		{4, 3, []string{"1110", "0111", "1111"}},
	}

	for i, d := range data {
		res := NumOfBinaryStringsNWithAtLeastKAdjacent1Bits(d.n, d.k)
		sort.Strings(res)
		sort.Strings(d.out)
		if len(res) != len(d.out) {
			t.Errorf("case [%v] expected %v, got %v", i, len(d.out), len(res))
		}
		for j := range res {
			if res[j] != d.out[j] {
				t.Errorf("case [%v] expected %v, got %v", i, d.out[j], res[j])
			}
		}
	}
}

func TestNumOfBinaryStringsWithSameSumOfFirstAndSecondHalf(t *testing.T) {
	data := []struct {
		n   int
		out []string
	}{
		{2, []string{"0101", "1111", "1001", "0110", "0000", "1010"}},
		{3, []string{"011011", "001001", "011101", "010001", "101011", "111111",
			"110011", "101101", "100001", "110101", "001010", "011110",
			"010010", "001100", "000000", "010100", "101110", "100010",
			"110110", "100100",
		}},
	}

	for i, d := range data {
		res := NumOfBinaryStringsWithSameSumOfFirstAndSecondHalf(d.n)
		sort.Strings(res)
		sort.Strings(d.out)
		if len(res) != len(d.out) {
			t.Errorf("case [%v] expected %v, got %v", i, len(d.out), len(res))
		}
		for j := range res {
			if res[j] != d.out[j] {
				t.Errorf("case [%v] expected %v, got %v", i, d.out[j], res[j])
			}
		}
	}
}

func TestNumOfBinaryStringsNWithWithKAdjacentSetBits(t *testing.T) {
	data := []struct {
		n, k, out int
	}{
		{4, 2, 2},
		{5, 2, 6},
	}

	for i, d := range data {
		res := NumOfBinaryStringsNWithWithKAdjacentSetBits(d.n, d.k)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}
