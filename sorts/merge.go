package sorts

import "slices"

// time complexity: O(n*log(n))
// space complexity: O(n)
func MergeSort(array []int) {
	l := len(array)
	if l < 2 {
		return
	}

	leftArray := slices.Clone(array[:l/2])
	rightArray := slices.Clone(array[l/2:])

	MergeSort(leftArray)
	MergeSort(rightArray)

	l, r, j := 0, 0, 0
	for l < len(leftArray) && r < len(rightArray) {
		if leftArray[l] < rightArray[r] {
			array[j] = leftArray[l]
			l++
		} else {
			array[j] = rightArray[r]
			r++
		}
		j++
	}
	for r < len(leftArray) {
		array[j] = rightArray[r]
		r++
		j++
	}
	for l < len(leftArray) {
		array[j] = leftArray[l]
		l++
		j++
	}
}
