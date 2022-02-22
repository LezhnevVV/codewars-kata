package bitcounting

import "strconv"

func CountBits(input uint) int {
	var digit = string(strconv.FormatInt(int64(input), 2))
	var result = 0

	for _, eachOther := range digit {
		if eachOther == '1' {
			result++
		}
	}
	return result
}
