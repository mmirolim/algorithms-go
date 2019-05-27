package set

import (
	"testing"
)

func TestReservoirSampling(t *testing.T) {
	data := []struct {
		set []int
		k   int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
	}

	for i, d := range data {
		out := ReservoirSampling(d.k, d.set)
		if len(out) != d.k {
			t.Errorf("case [%v] expected %v, got %v", i, d.k, len(out))
		}

	}
}

func TestRandomKSample(t *testing.T) {
	data := []struct {
		set []int
		k   int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
	}

	for i, d := range data {
		out := RandomKSample(d.k, d.set)
		if len(out) != d.k {
			t.Errorf("case [%v] expected %v, got %v", i, d.k, len(out))
		}
	}
}
