package sorts

func MergeSort(array []int) []int {
	l := len(array)
	if l < 2 {
		return array
	}

	leftArray := MergeSort(array[:l/2])
	rightArray := MergeSort(array[l/2:])

	result := make([]int, 0, len(array))
	l, r := 0, 0
	for l < len(leftArray) && r < len(rightArray) {
		if leftArray[l] < rightArray[r] {
			result = append(result, leftArray[l])
			l++
		} else {
			result = append(result, rightArray[r])
			r++
		}
	}
	result = append(result, leftArray[l:]...)
	result = append(result, rightArray[r:]...)

	return result
}
