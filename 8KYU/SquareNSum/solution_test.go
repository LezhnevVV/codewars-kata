package squareSum

import "testing"

func TestSquareSum(t *testing.T) {
	v := SquareSum([]int{1, 2, 2})
	if v != 9 {
		t.Error("Expected 9, got ", v)
	}

	v = SquareSum([]int{0, 3, 5, 2})
	if v != 38 {
		t.Error("Expected 38, got ", v)
	}
}
