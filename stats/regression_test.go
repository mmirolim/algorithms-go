package stats

import (
	"testing"

	"github.com/mmirolim/algos/checks"
	"github.com/mmirolim/algos/math"
)

func TestLeastSquareRegressionLine(t *testing.T) {
	// https://www.hackerrank.com/challenges/s10-least-square-regression-line/tutorial
	// problem from example
	X := []float64{1, 2, 3, 4, 5}
	Y := []float64{2, 1, 4, 3, 5}
	expectedA, expectedB := 0.6, 0.8
	a, b := LeastSquareRegressionLine(X, Y)
	if !checks.Eqf64(a, expectedA) || !checks.Eqf64(b, expectedB) {
		t.Errorf("case 1 expected (a, b) (%g, %g), got (%g, %g)", expectedA, expectedB, a, b)
	}
}

func TestMultipleLinearRegression(t *testing.T) {
	// Tutorial https://www.hackerrank.com/challenges/s10-multiple-linear-regression/tutorial
	// https://www.hackerrank.com/challenges/s10-multiple-linear-regression/problem
	X := [][]float64{
		[]float64{0.18, 0.89},
		[]float64{1.0, 0.26},
		[]float64{0.92, 0.11},
		[]float64{0.07, 0.37},
		[]float64{0.85, 0.16},
		[]float64{0.99, 0.41},
		[]float64{0.87, 0.47},
	}
	Y := []float64{
		109.85,
		155.72,
		137.66,
		76.17,
		139.75,
		162.6,
		151.77,
	}
	B, err := MultipleLinearRegression(X, Y)
	if err != nil {
		t.Errorf("unexpected Inverse error %v", err)
		t.FailNow()
	}

	X2 := [][]float64{
		[]float64{1, 0.49, 0.18},
		[]float64{1, 0.57, 0.83},
		[]float64{1, 0.56, 0.64},
		[]float64{1, 0.76, 0.18},
	}

	// Find Y2 for X2
	Y2 := math.NewMat(X2).Mul(B)
	expectedY2 := math.NewMat([][]float64{
		[]float64{105.22}, []float64{142.68}, []float64{132.94}, []float64{129.71}},
	)
	prec := 1
	y2str := Y2.ToString(prec)
	expectedY2str := expectedY2.ToString(prec)
	if y2str != expectedY2str {
		t.Errorf("expected mat\n%s\ngot\n%s\n", expectedY2str, y2str)
	}
}
