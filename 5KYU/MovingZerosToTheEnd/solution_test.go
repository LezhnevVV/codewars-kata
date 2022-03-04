package movingzerostotheend

import "testing"

type MoveZerosTest struct {
	arr      []int
	expected []int
}

var MoveZerosTests = []MoveZerosTest{
	{
		arr:      []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1},
		expected: []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0},
	},
	{
		arr:      []int{},
		expected: []int{},
	},
	{
		arr:      []int{1, 2, 3, 4, 5, 1, 7, 8, 9, 9},
		expected: []int{1, 2, 3, 4, 5, 1, 7, 8, 9, 9},
	},
	{
		arr:      []int{1, 0, 0},
		expected: []int{1, 0, 0},
	},
}

func TestMoveZeros(t *testing.T) {
	for _, test := range MoveZerosTests {
		output := MoveZeros(test.arr)
		for _, element := range output {
			if output[element] != test.expected[element] {
				t.Errorf("Output array %v not equal to expected %v", output, test.expected)
				break
			}
		}
	}
}
func BenchmarkTribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})
	}
}
