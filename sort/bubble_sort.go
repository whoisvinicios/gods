package sort

import "gods/utils/arrays"

/*
 * Bubble Sort
 * O(n^2) time complexity - quadratic
 * Algorithm degrades quickly
 *
 * Example:
 * It will take 10.000 steps to sort 100 items, 1.000.000
 * steps to sort 1.000 items, and so on.
 */

func BubbleSort(array []int) {
	for lastUnsortedIndex := len(array) - 1; lastUnsortedIndex > 0; lastUnsortedIndex-- {
		for index := 0; index < lastUnsortedIndex; index++ {
			if array[index] > array[index+1] {
				arrays.Swap(array, index, index+1)
			}
		}
	}
}
