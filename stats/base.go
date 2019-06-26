package stats

import (
	"math"
	"sort"
)

func FindQuartiles(nums []int) (q1, q2, q3 float64) {
	sort.Ints(nums)
	medID := len(nums) / 2
	low := medID
	hi := medID
	if len(nums)&1 == 1 {
		// exclude median from hi part, if len(nums) is odd
		hi++
	}

	qs := make([]float64, 3)
	for i, arr := range [][]int{nums[:low], nums, nums[hi:]} {
		med := len(arr) / 2
		if len(arr)&1 == 1 {
			qs[i] = float64(arr[med])
		} else {
			qs[i] = float64(arr[med-1]+arr[med]) / 2
		}
	}
	q1, q2, q3 = qs[0], qs[1], qs[2]
	return
}

func Mean(nums []float64) float64 {
	mean := 0.0
	for i := range nums {
		mean += nums[i]
	}

	return mean / float64(len(nums))
}

func Median(nums []float64) float64 {
	sort.Float64s(nums)
	med := len(nums) / 2
	if len(nums)&1 == 1 {
		return nums[med]
	}
	return (nums[med-1] + nums[med]) / 2.0
}
func Mode(nums []float64) float64 {
	sort.Float64s(nums)
	freq := map[float64]int{}
	for i := range nums {
		freq[nums[i]]++
	}
	max := freq[nums[0]]
	mode := nums[0]
	for k, v := range freq {
		// need smallest number
		if v > max {
			max = v
			mode = k
		} else if v == max && mode > k {
			mode = k
		}
	}
	return mode
}

func WeightedMean(nums []float64, weights []float64) float64 {
	wmean := 0.0
	for i := range nums {
		wmean += (nums[i] * weights[i])
	}
	sumOfWeights := 0.0
	for i := range weights {
		sumOfWeights += weights[i]
	}
	return wmean / sumOfWeights
}

func StandardDeviation(nums []float64) float64 {
	mean := Mean(nums)
	sumOfVariance := 0.0
	for i := range nums {
		sumOfVariance += (nums[i] - mean) * (nums[i] - mean)
	}
	return math.Sqrt(sumOfVariance / float64(len(nums)))
}

func InterquartileRange(nums []int) float64 {
	q1, _, q3 := FindQuartiles(nums)
	return q3 - q1
}

func GenerateSeqFromNumAndFreq(nums, freq []int) []int {
	var out []int
	for i := range nums {
		for j := 0; j < freq[i]; j++ {
			out = append(out, nums[i])
		}
	}
	return out
}

func BinomialDistribution(x, n int, p float64) float64 {
	return BinomialCoefficients(n, x) * math.Pow(p, float64(x)) * math.Pow(1-p, float64(n-x))
}
func NegativeBinomialDistribution(x, n int, p float64) float64 {
	return BinomialCoefficients(n-1, x-1) * math.Pow(p, float64(x)) * math.Pow(1-p, float64(n-x))
}
func GeometricDistribution(n int, p float64) float64 {
	return math.Pow(1-p, float64(n-1)) * p
}

/*
(n r) = n!/(r!(n-r)!)
*/
func BinomialCoefficients(n, r int) float64 {
	numerator := 1
	for i := 1; i <= r; i++ {
		numerator *= n - r + i
	}
	denominator := Factorial(r)
	return float64(numerator) / float64(denominator)
}

func Factorial(n int) int {
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}
