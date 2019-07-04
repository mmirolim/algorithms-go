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

func TestIdentityMat(t *testing.T) {
	data := []struct {
		dim         int
		expectedMat *Matrix
	}{
		{dim: 0, expectedMat: nil},
		{dim: 1, expectedMat: &Matrix{row: 1, col: 1, mlen: 1, m: []float64{1}}},
		{dim: 2, expectedMat: &Matrix{row: 2, col: 2, mlen: 2, m: []float64{1, 0, 0, 1}}},
		{dim: 3, expectedMat: &Matrix{row: 3, col: 3, mlen: 9, m: []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}}}}
	for i, d := range data {
		m := IdentityMat(d.dim)
		if m == nil {
			if d.expectedMat == nil {
				continue
			}
			t.Errorf("case [%v] expected %#v, got %#v", i, d.expectedMat, m)
			continue
		}

		if m.row != d.dim || m.col != d.dim {
			t.Errorf("case [%v] expected dim (row, col) (%d, %d), got (%d, %d)", i, d.dim, d.dim, m.row, m.col)
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
func TestMatScalarMul(t *testing.T) {
	data := []struct {
		s      float64
		m, sxm *Matrix
	}{
		{
			10,
			NewMat([][]float64{[]float64{2, 3, 4}, []float64{1, 0, 0}}),
			NewMat([][]float64{[]float64{20, 30, 40}, []float64{10, 0, 0}}),
		},
		{
			2,
			NewMat([][]float64{[]float64{0, 1000}, []float64{1, 100}, []float64{0, 10}}),
			NewMat([][]float64{[]float64{0, 2000}, []float64{2, 200}, []float64{0, 20}})},
	}

	for i, d := range data {
		sxm := d.m.ScalarMul(d.s)
		if !sxm.IsEq(d.sxm) {
			t.Errorf("case [%v] expected mat %#v, got %#v", i, d.sxm, sxm)
		}
	}
}

func TestMatTranspose(t *testing.T) {
	data := []struct {
		m, tm *Matrix
	}{
		{
			NewMat([][]float64{[]float64{1, 2, 3}, []float64{0, -6, 7}}),
			NewMat([][]float64{[]float64{1, 0}, []float64{2, -6}, []float64{3, 7}}),
		},

		{
			NewMat([][]float64{[]float64{2, 3, 4}, []float64{1, 0, 0}}),
			NewMat([][]float64{[]float64{2, 1}, []float64{3, 0}, []float64{4, 0}}),
		},
	}

	for i, d := range data {
		tm := d.m.Transpose()
		if !tm.IsEq(d.tm) {
			t.Errorf("case [%v] expected mat %#v, got %#v", i, d.tm, tm)
		}
	}
}

func TestMatAdd(t *testing.T) {
	data := []struct {
		m1, m2  *Matrix
		m1Addm2 *Matrix
	}{
		{
			m1:      NewMat([][]float64{[]float64{2, 3}, []float64{1, 0}}),
			m2:      NewMat([][]float64{[]float64{0, 1000}, []float64{1, 100}}),
			m1Addm2: NewMat([][]float64{[]float64{2, 1003}, []float64{2, 100}}),
		},
		{
			m1:      NewMat([][]float64{[]float64{1, 2}, []float64{3, 4}}),
			m2:      NewMat([][]float64{[]float64{0, 1}, []float64{0, 0}}),
			m1Addm2: NewMat([][]float64{[]float64{1, 3}, []float64{3, 4}}),
		},
	}

	for i, d := range data {
		m1Addm2 := d.m1.Add(d.m2)
		if !m1Addm2.IsEq(d.m1Addm2) {
			t.Errorf("case [%v] expected mat %#v, got %#v", i, d.m1Addm2, m1Addm2)
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
