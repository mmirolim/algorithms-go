package stats

import (
	"math"
	"sort"
)

var sqrt2Pi = math.Sqrt(2 * math.Pi)
var sqrtPi = math.Sqrt(math.Pi)

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

func PoissonDistribution(lambda float64, k int) float64 {
	return math.Pow(lambda, float64(k)) * math.Exp(-lambda) / float64(Factorial(k))
}

// Expectation of X^2
func PoissonDistributionEX2(x float64) float64 {
	return x + x*x
}

// u = 0, sd = 1
func StandardNormalDistribution(x float64) float64 {
	return math.Exp(-(x*x)/2) / sqrt2Pi
}

func NormalDistribution(x, u, variance float64) float64 {
	return (1 / variance) * StandardNormalDistribution((x-u)/variance)
}

func CumulativeDistributionFunctionOfNormalDistribution(x, u, sd float64) float64 {
	return 0.5 * (1 + Erf((x-u)/(sd*math.Sqrt2)))
}

// Erf fn approximation with elementary functions, maximum error: 1.5×10−7
// https://en.wikipedia.org/wiki/Error_function
// Another approximation which can be used is given by Sergei Winitzki
// using his "global Padé approximations"
func Erf(x float64) float64 {
	// constants
	a1 := 0.254829592
	a2 := -0.284496736
	a3 := 1.421413741
	a4 := -1.453152027
	a5 := 1.061405429
	p := 0.3275911

	// Save the sign of x
	sign := 1
	if x < 0 {
		sign = -1
	}
	x = math.Abs(x)

	// A&S formula 7.1.26
	t := 1.0 / (1.0 + p*x)
	y := 1.0 - (((((a5*t+a4)*t)+a3)*t+a2)*t+a1)*t*math.Exp(-x*x)

	return float64(sign) * y
}

func Covariance(X, Y []float64) float64 {
	if len(X) != len(Y) {
		panic("X and Y number of values different")
	}
	xm := Mean(X)
	ym := Mean(Y)
	cov := 0.0
	for i := range X {
		cov += (X[i] - xm) * (Y[i] - ym)
	}
	return cov / float64(len(X))
}

func PearsonCorrelationCoefficient(X, Y []float64) float64 {
	return Covariance(X, Y) / (StandardDeviation(X) * StandardDeviation(Y))
}

func SpearmansRankCorrelationCoefficient(X, Y []float64) float64 {
	rankx := deriveRank(X)
	ranky := deriveRank(Y)
	return PearsonCorrelationCoefficient(rankx, ranky)
}

func deriveRank(X []float64) []float64 {
	rankx := make([]float64, len(X))
	index := map[float64]int{}
	for i := range X {
		index[X[i]] = i
	}
	sorted := make([]float64, 0, len(index))
	for k := range index {
		sorted = append(sorted, k)
	}
	sort.Float64s(sorted)
	for i := range sorted {
		// rank
		index[sorted[i]] = i + 1
	}
	for i := range X {
		rankx[i] = float64(index[X[i]])
	}
	return rankx
}
