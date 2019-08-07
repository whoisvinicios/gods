package arrays

// Swap two numbers of the given integer array
func Swap(array []int, i, j int) {
	if i == j {
		return
	}

	temporarily := array[i]
	array[i] = array[j]
	array[j] = temporarily
}
