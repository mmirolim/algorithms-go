package math

import (
	"fmt"
	"math/rand"
	"testing"
)

type Matrix struct {
	// [c][r]float64
	mi  interface{}
	m   [][]float32
	mc  []float32
	mci interface{}
}

var m1e4 = NewMat(1, 1, 1, true)
var dim = int(5e2)
var randPerms = rand.Perm(dim)

func NewMat(row, col int, bitsize int, float bool) *Matrix {
	d := make([][]float32, dim)
	for i := range d {
		d[i] = make([]float32, dim)
		for j := range d[i] {
			d[i][j] = float32(randPerms[j])
		}
	}

	d2 := make([][]float32, dim)
	d3 := make([]float32, dim*dim)
	d4 := make([]float32, dim*dim)
	for i := range d {
		d2[i] = make([]float32, dim)
		for j := range d2[i] {
			d2[i][j] = float32(randPerms[j])
			d3[i*dim+j] = d2[i][j]
			d4[i*dim+j] = d2[i][j]
		}
	}

	return &Matrix{mi: d, m: d2, mc: d3, mci: d4}
}

func (m1 *Matrix) Mul(m2 *Matrix) *Matrix {
	m3 := make([][]float32, dim)
	for i := range m3 {
		m3[i] = make([]float32, dim)
	}
	var sum float32
	for i := range m3 {
		for j := range m3[i] {
			for x := 0; x < dim; x++ {
				sum += m1.m[x][j] + m2.m[i][x]
			}
			m3[i][j] = sum

		}
	}
	mn := new(Matrix)
	mn.m = m3
	return mn
}
func (m1 *Matrix) MulWithoutSum(m2 *Matrix) *Matrix {
	m3 := make([][]float32, dim)
	for i := range m3 {
		m3[i] = make([]float32, dim)
	}

	for i := range m3 {
		for j := range m3[i] {
			for x := 0; x < dim; x++ {
				m3[i][j] += m1.m[x][j] + m2.m[i][x]
			}

		}
	}
	mn := new(Matrix)
	mn.m = m3
	return mn
}

func (m1 *Matrix) MulCont(m2 *Matrix) *Matrix {
	m3 := make([]float32, dim*dim)
	var sum float32
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			for x := 0; x < dim; x++ {
				sum += m1.mc[x*dim+j] + m2.mc[i*dim+x]
			}
			m3[i*dim+j] = sum
		}
	}
	mn := new(Matrix)
	mn.mc = m3
	return mn
}

func (m1 *Matrix) MulContWithoutSum(m2 *Matrix) *Matrix {
	m3 := make([]float32, dim*dim)
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			for x := 0; x < dim; x++ {
				m3[i*dim+j] += m1.mc[x*dim+j] + m2.mc[i*dim+x]
			}

		}
	}
	mn := new(Matrix)
	mn.mc = m3
	return mn
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.Mul(m1e4)
	}
}

func BenchmarkMulWithoutSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.MulWithoutSum(m1e4)
	}
}

func BenchmarkMulCont(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.MulCont(m1e4)
	}
}

func BenchmarkMulContWithoutSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.MulContWithoutSum(m1e4)
	}
}

func (m *Matrix) IterateSumInterface() int {
	sum := 0
	d := m.mi.([][]float32)
	for i := range d {
		for j := range d[i] {
			sum += int(d[i][j])
		}
	}
	return sum
}

func (m *Matrix) IterateSumf32CR() int {
	sum := 0
	d := m.m
	for i := range d {
		for j := range d {
			sum += int(d[i][j])
		}
	}
	return sum
}
func (m *Matrix) IterateSumf32RC() int {
	sum := 0
	d := m.m
	for i := range d {
		for j := range d {
			sum += int(d[j][i])
		}
	}
	return sum
}

func (m *Matrix) IterateSumf32Continuous() int {
	sum := 0
	d := m.mc
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			sum += int(d[r*dim+c])
		}
	}
	return sum
}

func (m *Matrix) IterateSumf32ContinuousInterface() int {
	sum := 0
	d := m.mci.([]float32)
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			sum += int(d[r*dim+c])
		}
	}
	return sum
}

func TestSum(t *testing.T) {
	m := NewMat(1, 1, 1, true)
	sInterface := m.IterateSumInterface()
	sf32CR := m.IterateSumf32CR()
	sf32RC := m.IterateSumf32RC()
	sf32Cont := m.IterateSumf32Continuous()
	sf32ContInf := m.IterateSumf32ContinuousInterface()
	fmt.Printf("%v sInterface\n%v sf32CR\n%v sf32RC\n%v sf32Cont\n%v sf32ContInf\n", sInterface, sf32CR, sf32RC, sf32Cont, sf32ContInf) // output for debug

}

func BenchmarkIterateSumInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.IterateSumInterface()
	}
}

func BenchmarkIterateSumf32CR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.IterateSumf32CR()
	}
}

func BenchmarkIterateSumf32RC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.IterateSumf32RC()
	}
}

func BenchmarkIterateSumf32Continuous(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.IterateSumf32Continuous()
	}
}

func BenchmarkIterateSumf32ContinuousInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = m1e4.IterateSumf32ContinuousInterface()
	}
}

/*
dim = 5e2
goos: linux
goarch: amd64
pkg: github.com/mmirolim/algos/math
BenchmarkMul-4                                	       3	 385947903 ns/op
BenchmarkMulWithoutSum-4                      	       2	 531958756 ns/op
BenchmarkMulCont-4                            	      10	 207222959 ns/op
BenchmarkMulContWithoutSum-4                  	       5	 389149370 ns/op
BenchmarkIterateSumInterface-4                	   10000	    150211 ns/op
BenchmarkIterateSumf32CR-4                    	   10000	    225037 ns/op
BenchmarkIterateSumf32RC-4                    	    3000	    423330 ns/op
BenchmarkIterateSumf32Continuous-4            	    5000	    256719 ns/op
BenchmarkIterateSumf32ContinuousInterface-4   	    5000	    225407 ns/op
*/

func TestMatDim(t *testing.T) {
	// 	data := []struct{

	// 	}
	// 	{1*2} *quantity {2*1} ===> value {1*1}
	// 	{2*3} times {3*5} ==> {2*5}
	// 	{3*2} times {2*4} ==> {3*4}
	// 	{1*2} times {2*1} ==> {1*1}
}
