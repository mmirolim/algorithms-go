package math

import (
	"strconv"
	"testing"

	"github.com/mmirolim/algos/checks"
)

func caseStr(i int) string {
	return "case [" + strconv.Itoa(i) + "]"
}
func TestBigAdd(t *testing.T) {
	data := []struct {
		s1, s2   []uint8
		expected []uint8
	}{
		{
			[]uint8{3, 2, 1},
			[]uint8{7, 0, 1, 5},
			[]uint8{0, 3, 2, 5},
		},
		{
			[]uint8{0, 0, 9, 6, 2, 3, 2, 1},
			[]uint8{7, 0, 1, 5},
			[]uint8{7, 0, 0, 2, 3, 3, 2, 1},
		},
	}
	for i, d := range data {
		add(&d.s1, &d.s2)
		checks.AssertEq(t, d.expected, d.s1, caseStr(i))
	}
}

func TestScalarMul(t *testing.T) {
	data := []struct {
		s        []uint8
		c        uint8
		expected []uint8
	}{
		{[]uint8{3, 2, 1}, 3, []uint8{9, 6, 3}},
		{[]uint8{7, 0, 1, 5}, 9, []uint8{3, 6, 9, 5, 4}},
	}
	for i, d := range data {
		scalarMul(d.c, &d.s)
		if !checks.AssertEq(t, d.expected, d.s, caseStr(i)) {
			break
		}
	}

}

func TestBigMul(t *testing.T) {
	data := []struct {
		n, m     int
		expected string
	}{
		{123029, 124, "15255596"},
		{981, 982, "963342"},
	}
	for i, d := range data {
		bn := NewBig(d.n)
		bn.Mul(d.m)
		checks.AssertEq(t, d.expected, bn.String(), caseStr(i))
	}

}

func TestBigString(t *testing.T) {
	bn := NewBig(123029)
	out := bn.String()
	checks.AssertEqStr(t, "123029", out)
}

func TestBigFactorial(t *testing.T) {
	// solution to
	// https://www.hackerrank.com/challenges/extra-long-factorials/problem
	/*
		Factorials of  can't be stored even in a  long long variable.
		   Big integers must be used for such calculations
	*/
	expected := "15511210043330985984000000"
	n := 25
	bn := NewBig(1)
	for i := 1; i <= n; i++ {
		bn.Mul(i)
	}
	checks.AssertEq(t, expected, bn.String())
}
