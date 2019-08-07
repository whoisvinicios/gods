package sort

import "gods/utils/arrays"

/*
 * Selection Sort
 *
 * In-place algorithm
 * O(n^2) time complexity - quadratic
 * Doesn't requires as much swapping as bubble sort
 * Unstable Algorithm
 *
 * Example:
 * It will take 10.000 steps to sort 100 items, 1.000.000
 * steps to sort 1.000 items, and so on.
 */

func SelectionSort(array []int) {
	for lastUnsortedIndex := len(array) - 1; lastUnsortedIndex > 0; lastUnsortedIndex-- {
		var largest int
		for index := 0; index <= lastUnsortedIndex; index++ {
			if array[index] > array[largest] {
				largest = index
			}
		}
		arrays.Swap(array, largest, lastUnsortedIndex)
	}
}
