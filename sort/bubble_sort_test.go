package sort

import (
	"gods/utils/arrays"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, array := range arrays.TestArrays {
		if len(arrays.TestArrays) <= 1 {
			t.Error("The array must have more than 1 elements")
		}
		
		BubbleSort(array.Array)
		if !arrays.Equal(array.Array, array.Expect) {
			t.Error("exp:", array.Expect, "got:", array.Array)
		}
	}
}
