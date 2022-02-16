package findthesmallest

func SmallestIntegerFinder(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}
	return min
}
