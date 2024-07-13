package search

import (
	"slices"
)

/*
time complexity: O(log(n))

space complexity: O(1)

Binary search needs sorted array.
*/
func BinarySearch(array []int, target int) (int, error) {
	if !slices.IsSorted(array) {
		return -1, ErrIsNotSorted
	}
	min, max := 0, len(array)-1
	for min <= max {
		mid := min + (max-min)/2 // to avoid overflow
		if target > array[mid] {
			min = mid + 1
		} else if target < array[mid] {
			max = mid - 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}
