package number

import (
	"testing"
)

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
