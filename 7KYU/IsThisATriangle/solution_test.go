package isthisatriangle

import "testing"

func TestIsTriangle(t *testing.T) {
	v := IsTriangle(5, 1, 2)
	if v != false {
		t.Error("Expect false, got ", v)
	}

	v = IsTriangle(4, 2, 3)
	if v != true {
		t.Error("Expect false, got ", v)
	}

	v = IsTriangle(-1, 2, 3)
	if v != false {
		t.Error("Expect false, got ", v)
	}
}
