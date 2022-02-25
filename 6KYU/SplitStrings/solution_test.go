package splitstrings

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	v := Solution("abc")
	var v_res = []string{"ab", "c_"}
	for i := 0; i < 2; i++ {
		if strings.Compare(v[i], v_res[i]) != 0 {
			t.Error("Something wrong ", v[i], " != ", v_res[i])
		}
	}
}
