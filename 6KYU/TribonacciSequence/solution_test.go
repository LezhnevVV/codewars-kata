package tribonaccisequence

import "testing"

type TribonacciTest struct {
	arg1     [3]float64
	arg2     int
	expected []float64
}

var TribonacciTests = []TribonacciTest{
	{[3]float64{1, 1, 1}, 10, []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}},
	{[3]float64{0, 0, 1}, 10, []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}},
	{[3]float64{1, 0, 0}, 10, []float64{1, 0, 0, 1, 1, 2, 4, 7, 13, 24}},
	{[3]float64{0, 0, 0}, 10, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	{[3]float64{1, 2, 3}, 10, []float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230}},
	{[3]float64{3, 2, 1}, 10, []float64{3, 2, 1, 6, 9, 16, 31, 56, 103, 190}},
	{[3]float64{1, 1, 1}, 1, []float64{1}},
	{[3]float64{300, 200, 100}, 0, []float64{}},
	{[3]float64{0.5, 0.5, 0.5}, 15, []float64{0.5, 0.5, 0.5, 1.5, 2.5, 4.5, 8.5, 15.5, 28.5, 52.5, 96.5, 177.5, 326.5, 600.5, 1104.5}},
}

func TestTribonacci(t *testing.T) {
	for _, test := range TribonacciTests {
		output := Tribonacci(test.arg1, test.arg2)
		if len(output) != len(test.expected) {
			t.Errorf("Input was %v, output %v not equal to expected %v", test.arg1, output, test.expected)
		} else {
			for i := 0; i < len(output); i++ {
				if output[i] != test.expected[i] {
					t.Errorf("Input was %v, output %v not equal to expected %v", test.arg1, output, test.expected)
				}
			}
		}
	}
}

func BenchnarkTribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci([3]float64{1, 1, 1}, 10)
	}
}
