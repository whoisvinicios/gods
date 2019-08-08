package sort

import (
"gods/utils/arrays"
"testing"
)

func TestSelectionSort(t *testing.T) {
	for _, array := range arrays.TestArrays {
		SelectionSort(array.Array)
		if !arrays.Equal(array.Array, array.Expect) {
			t.Error("exp:", array.Expect, "got:", array.Array)
		}
	}
}

