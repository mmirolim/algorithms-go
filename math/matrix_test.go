package math

import (
	"testing"
)

func TestNewMat(t *testing.T) {
	data := []struct {
		data        [][]float64
		row, col    int
		expectedMat *Matrix
	}{
		{data: [][]float64{}, row: 0, col: 0, expectedMat: nil},
		{
			data: [][]float64{[]float64{54, 21}},
			row:  1, col: 2,
			expectedMat: &Matrix{row: 1, col: 2, mlen: 2, m: []float64{54, 21}},
		},
		{
			data: [][]float64{[]float64{1, 3, 1}, []float64{1, 0, 0}},
			row:  2, col: 3,
			expectedMat: &Matrix{row: 2, col: 3, mlen: 6, m: []float64{1, 3, 1, 1, 0, 0}},
		},
	}
	for i, d := range data {
		m := NewMat(d.data)
		// check for nil
		if m == nil {
			if d.expectedMat == nil {
				continue
			}
			t.Errorf("case [%v] expected %#v, got %#v", i, d.expectedMat, m)
			continue
		}

		if m.row != d.row || m.col != d.col {
			t.Errorf("case [%v] expected dim (row, col) (%d, %d), got (%d, %d)", i, d.row, d.col, m.row, m.col)
			break
		}

		if !m.IsEq(d.expectedMat) {
			t.Errorf("case [%v] expected mat %#v, got %#v", i, d.expectedMat, m)
		}

	}
}

func TestZeroMat(t *testing.T) {
	data := []struct {
		row, col    int
		expectedMat *Matrix
	}{
		{row: 1, col: 0, expectedMat: nil},
		{row: 1, col: 1, expectedMat: &Matrix{row: 1, col: 1, mlen: 1, m: []float64{0}}},
		{row: 1, col: 2, expectedMat: &Matrix{row: 1, col: 2, mlen: 2, m: []float64{0, 0}}},
		{row: 3, col: 2, expectedMat: &Matrix{row: 3, col: 2, mlen: 6, m: []float64{0, 0, 0, 0, 0, 0}}},
	}
	for i, d := range data {
		m := ZeroMat(d.row, d.col)
		if m == nil {
			if d.expectedMat == nil {
				continue
			}
			t.Errorf("case [%v] expected %#v, got %#v", i, d.expectedMat, m)
			continue
		}

		if m.row != d.row || m.col != d.col {
			t.Errorf("case [%v] expected dim (row, col) (%d, %d), got (%d, %d)", i, d.row, d.col, m.row, m.col)
			break
		}

		if !m.IsEq(d.expectedMat) {
			t.Errorf("case [%v] expected mat %#v, got %#v", i, d.expectedMat, m)
		}

	}
}

func TestMatMul(t *testing.T) {
	data := []struct {
		m1, m2  *Matrix
		m1Mulm2 *Matrix
	}{
		{
			m1:      NewMat([][]float64{[]float64{2, 3, 4}, []float64{1, 0, 0}}),
			m2:      NewMat([][]float64{[]float64{0, 1000}, []float64{1, 100}, []float64{0, 10}}),
			m1Mulm2: NewMat([][]float64{[]float64{3, 2340}, []float64{0, 1000}}),
		},
		{
			m1:      NewMat([][]float64{[]float64{1, 2}, []float64{3, 4}}),
			m2:      NewMat([][]float64{[]float64{0, 1}, []float64{0, 0}}),
			m1Mulm2: NewMat([][]float64{[]float64{0, 1}, []float64{0, 3}}),
		},
	}

	for i, d := range data {
		m1xm2 := d.m1.Mul(d.m2)
		if !m1xm2.IsEq(d.m1Mulm2) {
			t.Errorf("case [%v] expected mat %#v, got %#v", i, d.m1Mulm2, m1xm2)
		}
	}
}

func TestMatToString(t *testing.T) {
	data := []struct {
		m              *Matrix
		expectedString string
	}{
		{NewMat([][]float64{[]float64{2, 3, 4}, []float64{1, 0, 0}}), "| 2 3 4 |\n| 1 0 0 |\n"},
		{NewMat([][]float64{[]float64{0, 1000}, []float64{1, 100}, []float64{0, 10}}), "| 0 1000 |\n| 1 100 |\n| 0 10 |\n"},
		{NewMat([][]float64{[]float64{0, 1}, []float64{0, 3}}), "| 0 1 |\n| 0 3 |\n"},
	}

	for i, d := range data {
		res := d.m.ToString()
		if d.expectedString != res {
			t.Errorf("case [%v] expected mat \n%s\ngot\n%s", i, d.expectedString, res)
		}
	}
}
