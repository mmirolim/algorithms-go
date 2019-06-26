package stats

import (
	"fmt"
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
