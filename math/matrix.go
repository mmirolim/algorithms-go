package math

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// TODO inverse, tranpose matrix float64 operations
type Matrix struct {
	m        []float64
	row, col int
	mlen     int
}

// [r][c]float64
func NewMat(data [][]float64) *Matrix {
	mat := new(Matrix)
	mat.row = len(data)
	if mat.row == 0 {
		return nil
	} else if len(data[0]) == 0 {
		return nil
	}
	col := len(data[0])
	for i := 0; i < mat.row; i++ {
		if col != len(data[i]) {
			return nil
		}
	}
	mat.col = col
	mat.mlen = mat.row * mat.col
	mat.m = make([]float64, mat.mlen)
	for r := 0; r < mat.row; r++ {
		for c := 0; c < mat.col; c++ {
			mat.m[r*mat.col+c] = data[r][c]
		}
	}
	return mat
}

func ZeroMat(row, col int) *Matrix {
	if row == 0 || col == 0 {
		return nil
	}
	mat := new(Matrix)
	mat.row, mat.col = row, col
	mat.mlen = row * col
	mat.m = make([]float64, mat.mlen)
	return mat
}

func IdentityMat(dim int) *Matrix {
	if dim == 0 {
		return nil
	}
	mat := new(Matrix)
	mat.row, mat.col = dim, dim
	mat.mlen = dim * dim
	mat.m = make([]float64, mat.mlen)
	for i := 0; i < dim; i++ {
		mat.m[i*dim+i] = 1.0
	}
	return mat
}

// panics if not square matrix
func (m *Matrix) Cofactor(r, c int) *Matrix {
	if m == nil || m.row != m.col {
		panic("defined only for square matrix")
	}
	dim := m.col
	out := ZeroMat(dim-1, dim-1)
	i := 0
	for row := 0; row < dim; row++ {
		if row == r {
			continue
		}
		j := 0
		for col := 0; col < dim; col++ {
			if col == c {
				continue
			}
			out.m[i*out.col+j] = m.m[row*dim+col]
			j++
		}
		i++
	}

	return out
}

// panics if not square matrix
func (m *Matrix) Det() float64 {
	if m == nil || m.row != m.col {
		panic("defined only for square matrix")
	}
	dim := m.col
	// base cases
	if dim == 1 {
		return m.m[0]
	} else if dim == 2 {
		return m.m[0]*m.m[3] - m.m[1]*m.m[2]
	}
	sign := func(i, j int) float64 {
		if (i+j)&1 == 1 {
			return -1.0
		}
		return 1.0
	}
	res := 0.0
	for col := 0; col < dim; col++ {
		res += sign(0, col) * m.m[col] * m.Cofactor(0, col).Det()
	}

	return res
}

func (m *Matrix) Inverse() (*Matrix, error) {
	if m.row != m.col {
		return nil, errors.New("not a square matrix")
	}

	detM := m.Det()
	if detM == 0 {
		return nil, errors.New("inverse matrix not found")
	}

	sign := func(i, j int) float64 {
		if (i+j)&1 == 1 {
			return -1.0
		}
		return 1.0
	}
	dim := m.row
	out := ZeroMat(dim, dim)
	// find cofactor matrix
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			out.m[r*dim+c] = sign(r, c) * m.Cofactor(r, c).Det()
		}
	}
	// find Adjoint matrix by transposing it
	adjM := out.Transpose()
	return adjM.ScalarMul(1 / detM), nil
}

func (m *Matrix) Dim() (row, col int) {
	return m.row, m.col
}

// prec - precision
func (m *Matrix) ToString(prec int) string {
	var str strings.Builder
	for r := 0; r < m.row; r++ {
		str.WriteString("| ")
		for c := 0; c < m.col; c++ {
			str.WriteString(strconv.FormatFloat(m.m[r*m.col+c], 'f', prec, 64))
			str.WriteByte(' ')
		}
		str.WriteByte('|')
		str.WriteByte('\n')
	}

	return str.String()
}

func (m1 *Matrix) IsDimEq(m2 *Matrix) bool {
	if m2 == nil {
		return false
	}
	r1, c1 := m1.Dim()
	r2, c2 := m2.Dim()
	if r1 != r2 || c1 != c2 {
		return false
	}
	return true
}

func (m1 *Matrix) IsEq(m2 *Matrix) bool {
	if !m1.IsDimEq(m2) {
		return false
	}
	for i := 0; i < m1.mlen; i++ {
		if math.Abs(m1.m[i]-m2.m[i]) > 1.0E-5 {
			return false
		}
	}

	return true
}

func (m1 *Matrix) Mul(m2 *Matrix) *Matrix {
	// check inner dimension
	// {a*d} X {d*b} ===> value {a*b}
	if m2 == nil || m1.col != m2.row {
		return nil
	}
	dim := m1.col
	out := ZeroMat(m1.row, m2.col)
	for i := 0; i < out.row; i++ {
		for j := 0; j < out.col; j++ {
			sum := 0.0
			for x := 0; x < dim; x++ {
				sum += m1.m[i*m1.col+x] * m2.m[x*m2.col+j]
			}
			out.m[i*out.col+j] = sum
		}
	}

	return out
}

func (m1 *Matrix) Add(m2 *Matrix) *Matrix {
	// check inner dimension
	// {a*d} X {d*b} ===> value {a*b}
	if !m1.IsDimEq(m2) {
		return nil
	}
	row, col := m1.row, m2.col
	out := ZeroMat(row, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			out.m[i*col+j] = m1.m[i*col+j] + m2.m[i*col+j]
		}
	}

	return out
}

func (m1 *Matrix) ScalarMul(s float64) *Matrix {
	if m1 == nil {
		return nil
	}
	out := ZeroMat(m1.row, m1.col)
	for i := 0; i < m1.mlen; i++ {
		out.m[i] = s * m1.m[i]
	}

	return out
}

func (m1 *Matrix) Transpose() *Matrix {
	if m1 == nil {
		return nil
	}
	out := ZeroMat(m1.col, m1.row)
	for i := 0; i < m1.row; i++ {
		for j := 0; j < m1.col; j++ {
			out.m[j*out.col+i] = m1.m[i*m1.col+j]
		}
	}

	return out
}
