package sort

/*
 * Insertion Sort
 *
 * In-place algorithm
 * O(n^2) time complexity - quadratic
 * Unstable Algorithm
 *
 * Example:
 * It will take 10.000 steps to sort 100 items, 1.000.000
 * steps to sort 1.000 items, and so on.
 */

func InsertionSort(array []int) {
	for firstUnsortedIndex := 1; firstUnsortedIndex < len(array); firstUnsortedIndex++ {
		newElement := array[firstUnsortedIndex]

		var index int
		for index = firstUnsortedIndex; index > 0 && array[index-1] > newElement; index-- {
			array[index] = array[index-1]
		}
		array[index] = newElement
	}
}
