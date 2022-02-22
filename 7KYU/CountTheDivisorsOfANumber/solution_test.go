package countthedivisorsofanumber

import "testing"

func TestDivisors(t *testing.T) {
	v := Divisors(4)
	if v != 3 {
		t.Error("Expect 3, got ", v)
	}

	v = Divisors(5)
	if v != 2 {
		t.Error("Expect 2, got ", v)
	}

	v = Divisors(12)
	if v != 6 {
		t.Error("Expect 6, got ", v)
	}

	v = Divisors(30)
	if v != 8 {
		t.Error("Expect 8, got ", v)
	}
}
