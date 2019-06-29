package stats

import (
	"fmt"
	"math"
	"testing"
)

func TestFindQuartiles(t *testing.T) {
	data := []struct {
		nums       []int
		q1, q2, q3 float64
	}{
		{[]int{7, 15, 36, 39, 40, 41}, 15, 37.5, 40},
		{[]int{3, 7, 8, 5, 12, 14, 21, 13, 18}, 6, 12, 16},
		{[]int{6, 7, 15, 36, 39, 40, 41, 42, 43, 47, 49}, 15, 40, 43},
	}

	for i, d := range data {
		q1, q2, q3 := FindQuartiles(d.nums)
		if q1 != d.q1 || q2 != d.q2 || q3 != d.q3 {
			t.Errorf("case [%v] expected (q1, q2, q3) (%v, %v, %v), got (%v, %v, %v)", i, d.q1, d.q2, d.q3, q1, q2, q3)
		}
	}
}

func TestMeanMedianMode(t *testing.T) {
	data := []struct {
		nums               []float64
		mean, median, mode float64
	}{
		{[]float64{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}, 43900.6, 44627.5, 4978},
	}

	for i, d := range data {
		mean := Mean(d.nums)
		median := Median(d.nums)
		mode := Mode(d.nums)
		if mean != d.mean || median != d.median || mode != d.mode {
			t.Errorf("case [%v] expected (mean, median, mode) (%v, %v, %v), got (%v, %v, %v)", i, d.mean, d.median, d.mode, mean, median, mode)
		}
	}
}

func TestWeightedMean(t *testing.T) {
	data := []struct {
		nums, weights []float64
		wmean         float64
	}{
		{[]float64{10, 40, 30, 50, 20}, []float64{1, 2, 3, 4, 5}, 32.0},
	}
	for i, d := range data {
		wmean := WeightedMean(d.nums, d.weights)
		if wmean != d.wmean {
			t.Errorf("case [%v] expected %v, got %v", i, d.wmean, wmean)
		}
	}
}

func TestStandardDeviation(t *testing.T) {
	data := []struct {
		nums []float64
		sd   float64
	}{
		{[]float64{10, 40, 30, 50, 20}, 14.1},
	}
	for i, d := range data {
		sd := StandardDeviation(d.nums)
		if fmt.Sprintf("%0.1f", sd) != fmt.Sprintf("%0.1f", d.sd) {
			t.Errorf("case [%v] expected %v, got %v", i, d.sd, sd)
		}
	}
}

func TestInterquartileRange(t *testing.T) {
	data := []struct {
		nums  []int
		freqs []int
		ir    float64
	}{
		{[]int{6, 12, 8, 10, 20, 16}, []int{5, 4, 3, 2, 1, 5}, 9.0},
	}
	for i, d := range data {
		ir := InterquartileRange(GenerateSeqFromNumAndFreq(d.nums, d.freqs))
		if fmt.Sprintf("%0.1f", ir) != fmt.Sprintf("%0.1f", d.ir) {
			t.Errorf("case [%v] expected %v, got %v", i, d.ir, ir)
		}
	}
}

func TestBinomialDistribution(t *testing.T) {
	/*
	   A fair coin is tossed 10 times. Find the following probabilities:

	   Getting 5 heads.
	   Getting at least 5 heads.
	   Getting at most 5 heads.
	*/
	heads := 5
	n := 10  // tossed
	p := 0.5 // fair coin prob
	// case 1 5 heads
	res := BinomialDistribution(heads, n, p)
	expected := 0.24609375
	if res != expected {
		t.Errorf("case [1] expected %v, got %v", expected, res)
	}
	// case 2 at least 5 heads
	res = 0
	for i := heads; i <= n; i++ {
		res += BinomialDistribution(i, n, p)
	}
	expected = 0.623046875
	if res != expected {
		t.Errorf("case [1] expected %v, got %v", expected, res)
	}
	// case 3 at most 5 heads
	res = 0
	for i := 0; i <= heads; i++ {
		res += BinomialDistribution(i, n, p)
	}
	expected = 0.623046875
	if res != expected {
		t.Errorf("case [1] expected %v, got %v", expected, res)
	}

}

// https://www.hackerrank.com/challenges/s10-binomial-distribution-1/problem
func TestBinomialDistributionAtLeast3BoysOutOf6Children(t *testing.T) {
	/*	Task
		The ratio of boys to girls for babies born in Russia is 1.09:1. If there is  child born per birth, what proportion of Russian families with exactly 6 children will have at least 3 boys?

		Write a program to compute the answer using the above parameters. Then print your result, rounded to a scale of 3 decimal places (i.e., 1.234 format).
	*/
	n, p, x := 6, 1.09/2.09, 3
	expectedResult := 0.696
	res := 0.0
	for i := x; i <= n; i++ {
		res += BinomialDistribution(i, n, p)
	}
	if fmt.Sprintf("%0.3f", res) != fmt.Sprintf("%0.3f", expectedResult) {
		t.Errorf("expected %v, got %v", expectedResult, res)
	}
}

// https://www.hackerrank.com/challenges/s10-geometric-distribution-2/problem
func TestGeometricDistributionWhatProbToFindDefectDuringFirstNInspections(t *testing.T) {
	/*
	   Task
	   The probability that a machine produces a defective product is 1/3. What is the probability that the 1th defect is found during the first 5 inspections?
	*/
	p := 1.0 / 3
	n := 5
	expectedResult := 0.868
	res := 0.0
	for i := 1; i <= n; i++ {
		res += GeometricDistribution(i, p)
	}
	if fmt.Sprintf("%0.3f", res) != fmt.Sprintf("%0.3f", expectedResult) {
		t.Errorf("expected %v, got %v", expectedResult, res)
	}
}

// https://www.hackerrank.com/challenges/s10-poisson-distribution-1/problem
func TestPoissonDistribution(t *testing.T) {
	/*
		Task
		A random variable, X, follows Poisson distribution with mean of 2.5.
		Find the probability with which the random variable X is equal to 5.
	*/

	X := 2.5
	K := 5
	expectedResult := 0.067
	res := PoissonDistribution(X, K)
	if fmt.Sprintf("%0.3f", res) != fmt.Sprintf("%0.3f", expectedResult) {
		t.Errorf("expected %v, got %v", expectedResult, res)
	}

}

// https://www.hackerrank.com/challenges/s10-normal-distribution-2/problem
func TestNormalDistributionWhatPercentageOfTheStudentsWithXGrade(t *testing.T) {
	/*
			Task
			The final grades for a Physics exam taken by a large group of students
			have a mean of 70 and a standard deviation of 10.
			If we can approximate the distribution of these grades by
			a normal distribution, what percentage of the students:
		        Scored higher than 80 (i.e., have a grade > 80)?
			Passed the test (i.e., have a grade => 60)?
			Failed the test (i.e., have a grade < 60)?
			Find and print the answer to each question on a new line,
			rounded to a scale of 2 decimal places.
	*/
	mean := 70.0
	sd := 10.0
	m80 := 80.0
	p60 := 60.0
	f60 := 60.0
	expectedQ1Res := 15.87
	expectedQ2Res := 84.13
	expectedQ3Res := 15.87
	res := 100 * (1 - CumulativeDistributionFunctionOfNormalDistribution(m80, mean, sd))
	if fmt.Sprintf("%0.2f", res) != fmt.Sprintf("%0.2f", expectedQ1Res) {
		t.Errorf("expected %v, got %v", expectedQ1Res, res)
	}
	res = 100 * (1 - CumulativeDistributionFunctionOfNormalDistribution(p60, mean, sd))
	if fmt.Sprintf("%0.2f", res) != fmt.Sprintf("%0.2f", expectedQ2Res) {
		t.Errorf("expected %v, got %v", expectedQ2Res, res)
	}
	res = 100 * CumulativeDistributionFunctionOfNormalDistribution(f60, mean, sd)
	if fmt.Sprintf("%0.2f", res) != fmt.Sprintf("%0.2f", expectedQ3Res) {
		t.Errorf("expected %v, got %v", expectedQ3Res, res)
	}

}

func TestCentralLimitTheorem(t *testing.T) {
	// https://www.hackerrank.com/challenges/s10-the-central-limit-theorem-1/problem
	// elevator can transport a max weightww
	mw := 9800.0
	n := 49.0
	u := 205.0
	sd := 15.0
	expectedResult := 0.0098
	// Question is :  what is the probability that all
	// boxes can be safely loaded into the freight elevator and transported?
	res := CumulativeDistributionFunctionOfNormalDistribution(mw, n*u, sd*math.Sqrt(n))
	if fmt.Sprintf("%0.4f", res) != fmt.Sprintf("%0.4f", expectedResult) {
		t.Errorf("expected %v, got %v", expectedResult, res)
	}

	// https://www.hackerrank.com/challenges/s10-the-central-limit-theorem-2/problem
	nt := 250.0
	ns := 100.0
	u = 2.4
	sd = 2.0
	expectedResult = 0.6915
	res = CumulativeDistributionFunctionOfNormalDistribution(nt, ns*u, sd*math.Sqrt(ns))
	if fmt.Sprintf("%0.4f", res) != fmt.Sprintf("%0.4f", expectedResult) {
		t.Errorf("expected %v, got %v", expectedResult, res)
	}
}

func TestSpearmansRankCorrelationCoefficient(t *testing.T) {
	X := []float64{0.2, 1.3, 0.2, 1.1, 1.4, 1.5}
	Y := []float64{1.9, 2.2, 3.1, 1.2, 2.2, 2.2}

	rs := 0.158114
	res := SpearmansRankCorrelationCoefficient(X, Y)
	if fmt.Sprintf("%0.3f", rs) != fmt.Sprintf("%0.3f", res) {
		t.Errorf("expected %v, got %v", rs, res)
	}
}
