package sorting

import (
	"testing"
)

// EqualSlices compare two slices to each other. Returns a Boolean
func EqualSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for indx, elem := range a {
		if elem != b[indx] {
			return false
		}
	}

	return true

}

// TestMergeSort tests MergeSort function
func TestMergeSort(t *testing.T) {

	initial := []int{2, 1, 45, 4, 9, 7, 6, 89, 5, 10, 34, 3, 8}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 34, 45, 89}

	mergesorted := MergeSort(initial, 0, len(initial))

	if !(EqualSlices(mergesorted, sorted)) {
		t.Error("MergeSort failed to sort")
	}

}
