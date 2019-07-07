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
