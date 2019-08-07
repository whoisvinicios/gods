package sort

import (
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
		if !Equal(array.array, array.expect) {
			t.Error("exp:", array.expect, "got:", array.array)
		}
	}
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
