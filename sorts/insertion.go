package sorts

/*
time complexity: O(n^2)

space complexity: O(1)
*/
func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
}
