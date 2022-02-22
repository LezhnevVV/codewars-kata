package bitcounting

import "testing"

func TestBitCounting(t *testing.T) {
	v := CountBits(0)
	if v != 0 {
		t.Error("Expect 0, got ", v)
	}

	v = CountBits(4)
	if v != 1 {
		t.Error("Expect 1. got ", v)
	}

	v = CountBits(0)
	if v != 0 {
		t.Error("Expect 0, got ", v)
	}

	v = CountBits(7)
	if v != 3 {
		t.Error("Expect 3, got ", v)
	}

	v = CountBits(9)
	if v != 2 {
		t.Error("Expect 2, got ", v)
	}

	v = CountBits(10)
	if v != 2 {
		t.Error("Expect 2, got ", v)
	}
}
