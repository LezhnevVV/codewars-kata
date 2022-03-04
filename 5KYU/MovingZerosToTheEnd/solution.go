package movingzerostotheend

func MoveZeros(arr []int) []int {
	sortingSlice := make([]int, 0, len(arr))
	var counter int
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			sortingSlice = append(sortingSlice, arr[i])
			counter = counter + 1
		}
	}
	for i := counter; i < len(arr); i++ {
		sortingSlice = append(sortingSlice, 0)
	}
	return sortingSlice
}
