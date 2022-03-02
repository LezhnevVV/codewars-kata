package bitcounting

import "testing"

type addTest struct {
	arg1, expected int
}

var addTests = []addTest{
	{0, 0},
	{4, 1},
	{7, 3},
	{9, 2},
	{10, 2},
}

func TestBitCounting(t *testing.T) {
	for _, test := range addTests {
		if output := CountBits(uint(test.arg1)); output != test.expected {
			t.Errorf("Input was %v, output %v not equal to expected %v", test.arg1, output, test.expected)
		}
	}
}

func BenchmarkBitCounting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountBits(10)
	}
}
