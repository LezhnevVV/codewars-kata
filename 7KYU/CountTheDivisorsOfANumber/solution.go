package countthedivisorsofanumber

func Divisors(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result++
		}
	}
	return result
}
