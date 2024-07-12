package sorts

// time complexity: O(n^2)
// space complexity: O(1)
func SelectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//find the index of the smallest element
		index := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[index] {
				index = j
			}
		}
		array[i], array[index] = array[index], array[i]
	}
}
