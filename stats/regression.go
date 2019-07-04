package stats

import (
	"github.com/mmirolim/algos/math"
)

// LeastSquareRegressionLine Y = a + b*X
// returns a, b
func LeastSquareRegressionLine(X, Y []float64) (a, b float64) {
	b = PearsonCorrelationCoefficient(X, Y) * StandardDeviation(Y) / StandardDeviation(X)
	a = Mean(Y) - b*Mean(X)
	return
}

// returns B matrix in Y = X*B
// Y = [1 x1 x2 ... xm] X B
func MultipleLinearRegression(X [][]float64, Y []float64) (*math.Matrix, error) {
	x := make([][]float64, len(X))
	for i := range X {
		x[i] = make([]float64, len(X[0])+1)
		x[i][0] = 1
		copy(x[i][1:], X[i])
	}

	y := make([][]float64, len(Y))
	for i := range Y {
		y[i] = make([]float64, 1)
		y[i][0] = Y[i]
	}

	// find first a, b1, b2n
	Xmat := math.NewMat(x)
	XTmat := Xmat.Transpose()
	Xinv, err := XTmat.Mul(Xmat).Inverse()
	if err != nil {
		return nil, err
	}
	Ymat := math.NewMat(y)
	B := Xinv.Mul(XTmat).Mul(Ymat)
	return B, nil
}
