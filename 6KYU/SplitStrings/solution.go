package splitstrings

func Solution(str string) []string {
	if len(str)%2 == 1 {
		str += "_"
	}
	result := make([]string, len(str)/2)
	for i := 0; i < cap(result); i++ {
		result[i] = str[i*2 : i*2+2]
	}
	return result
}
