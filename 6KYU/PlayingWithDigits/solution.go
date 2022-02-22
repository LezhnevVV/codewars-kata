package playingwithdigits

import (
	"fmt"
	"math"
)

func DigPow(n, p int) int {
	var result int
	var n_work int = n
	var length = len(fmt.Sprint(n))

	p = p + length - 1

	for i := 1; i <= length; i++ {
		result = result + int(math.Pow((float64(n_work%10)), float64(p)))
		p = p - 1
		n_work = n_work / 10
	}
	if result%n == 0 {
		return result / n
	}
	return -1
}
