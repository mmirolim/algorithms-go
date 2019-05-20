package bitsets

import (
	"testing"
)

func TestFindMissingNumber(t *testing.T) {
	given := []int{1, 2, 4, 5, 6, 7, 9, 10, 12, 13, 16}
	expected := []int{3, 8, 11, 14, 15, 17}
	out := FindMissingNumber(given, 17)
	if len(expected) != len(out) {
		t.Errorf("expected len %v, got %v", len(expected), len(out))
	}

	for i := range expected {
		if expected[i] != out[i] {
			t.Errorf("expected  val %v, got %v", expected[i], out[i])
		}
	}
}
