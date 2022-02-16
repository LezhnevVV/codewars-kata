package findthesmallest

import "testing"

func TestSmallestIntegerFinder(t *testing.T) {
	v := SmallestIntegerFinder([]int{5, 7, 4, 10})
	if v != 4 {
		t.Error("Expected 4, got ", v)
	}

	v = SmallestIntegerFinder([]int{6, -30, 0, 113, 74})
	if v != -30 {
		t.Error("Expected -30, got ", v)
	}
}
