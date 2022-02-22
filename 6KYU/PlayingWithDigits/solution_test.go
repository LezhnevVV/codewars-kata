package playingwithdigits

import "testing"

func TestPlayingWithDigits(t *testing.T) {
	v := DigPow(89, 1)
	if v != 1 {
		t.Error("Expect 1, got ", v)
	}

	v = DigPow(92, 1)
	if v != -1 {
		t.Error("Expect -1, got ", v)
	}

	v = DigPow(695, 2)
	if v != 2 {
		t.Error("Expect 2, got ", v)
	}

	v = DigPow(46288, 3)
	if v != 51 {
		t.Error("Expect 51, got ", v)
	}
}
