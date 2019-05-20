package bitsets

import "fmt"

func FindMissingNumber(in []int, maxNum int) []int {
	set := NewBitSet(maxNum)
	for _, v := range in {
		set.Set(v)
	}

	return set.GetClearBits()
}

type BitSet struct {
	segment int
	size    int
	store   []uint64
}

func NewBitSet(size int) *BitSet {
	s := new(BitSet)
	s.segment = 64
	s.size = size
	s.store = make([]uint64, size/s.segment+1)
	return s
}

// panics
func (s *BitSet) Set(num int) {
	if num > s.size || num < 1 {
		panic(
			fmt.Sprintf("number should be <= size, size = %d, number = %d", s.size, num),
		)
	}
	// get index
	segment := (num - 1) / s.segment
	idx := uint((num - 1) % s.segment)
	s.store[segment] = s.store[segment] | 1<<idx
}

// returns index of zero bits
func (s *BitSet) GetClearBits() []int {
	clearBits := make([]int, 0, s.size)
	for i, val := range s.store {
		for j := uint(0); j < 64; j++ {
			if (i*64 + int(j)) > s.size-1 {
				return clearBits
			}
			if val&(1<<j) == 0 {
				clearBits = append(clearBits, int(j)+i*64+1)
			}
		}
	}
	return clearBits
}
