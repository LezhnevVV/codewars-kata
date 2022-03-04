package sumofdigits

import "testing"

type DigitalRootTest struct {
	n        int
	expected int
}

var DigitalRootTests = []DigitalRootTest{
	{
		n:        16,
		expected: 7,
	},
	{
		n:        195,
		expected: 6,
	},
	{
		n:        992,
		expected: 2,
	},
	{
		n:        167346,
		expected: 9,
	},
	{
		n:        0,
		expected: 0,
	},
}

func TestDigitalRoot(t *testing.T) {
	for _, tests := range DigitalRootTests {
		if output := DigitalRoot(tests.n); output != tests.expected {
			t.Errorf("Result %v not equal to expected %v", output, tests.expected)
		}
	}
}

func BenchmarkDigitalRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitalRoot(123123)
	}
}
