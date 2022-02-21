package growthofapopulation

import "testing"

func TestNbYear(t *testing.T) {
	v := NbYear(1500, 5, 100, 5000)
	if v != 15 {
		t.Error("Expect 15, got ", v)
	}

	v = NbYear(1500000, 2.5, 10000, 2000000)
	if v != 10 {
		t.Error("Expect 10, got ", v)
	}

	v = NbYear(1500000, 0.25, 1000, 2000000)
	if v != 94 {
		t.Error("Expect 94, got ", v)
	}

	v = NbYear(1500000, 0.25, -1000, 2000000)
	if v != 151 {
		t.Error("Expect 151, got ", v)
	}
}
