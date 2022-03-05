package sumofdigits

func DigitalRoot(n int) int {
	var output int

	if n/10 != 0 {
		output = (n % 10) + DigitalRoot(n/10)
	} else {
		return n % 10
	}

	if output/10 != 0 {
		output = DigitalRoot(output)
	}
	return output
}
