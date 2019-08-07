package sort

import (
	"gods/utils/arrays"
	"testing"
)

var test = []struct {
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
}

func TestBubbleSort(t *testing.T) {
	for _, array := range test {
		BubbleSort(array.array)
		if !arrays.Equal(array.array, array.expect) {
			t.Error("exp:", array.expect, "got:", array.array)
		}
	}
}
