package sort

import (
"gods/utils/arrays"
"testing"
)

var selectionTestArray = []struct {
	array  []int
	expect []int
}{
	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		[]int{0, 1, 4, 2, 6, 5, -22},
		[]int{-22, 0, 1, 2, 4, 5, 6},
	},
	{
		[]int{1, 45, -1, -45, 2, 200, 4},
		[]int{-45, -1, 1, 2, 4, 45, 200},
	},
}

func TestSelectionSort(t *testing.T) {
	for _, array := range selectionTestArray {
		SelectionSort(array.array)
		if !arrays.Equal(array.array, array.expect) {
			t.Error("exp:", array.expect, "got:", array.array)
		}
	}
}

