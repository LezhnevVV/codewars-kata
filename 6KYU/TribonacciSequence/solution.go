package tribonaccisequence

func Tribonacci(signature [3]float64, n int) []float64 {
	result := make([]float64, n)

	if n < 3 {
		result = signature[:n]
		return result
	} else {
		for i := 0; i < 3; i++ {
			result[i] = signature[i]
		}

		for i := 3; i < len(result); i++ {
			result[i] = result[i-1] + result[i-2] + result[i-3]
		}
		return result
	}
}
