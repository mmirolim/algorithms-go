package set

import (
	"math/rand"
)

func ReservoirSampling(k int, set []int) []int {
	sampled := make([]int, 0, k)
	for i := 0; i < len(set); i++ {
		if i < k {
			sampled = append(sampled, set[i])
		} else if r := rand.Intn(i + 1); r < k {
			sampled[r] = set[i]
		}
	}

	return sampled
}

func RandomKSample(k int, set []int) []int {
	out := make([]int, 0, k)
	n := float64(len(set))
	for i := 0; i < len(set); i++ {
		p := float64(k) / (n - float64(i))
		rp := rand.Float64()
		if p > rp {
			k--
			out = append(out, set[i])
		}
	}
	return out
}
