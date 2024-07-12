package sorts

/*
average time complexity: O(n*log(n)).

worst time complexity: O(n^2).

average space complexity: O(log(n)).

worst space complexity: O(n).
*/
func QuickSort(array []int) {
	quickSortHelper(array, 0, len(array)-1)
}

func quickSortHelper(array []int, start, end int) {
	if end <= start {
		return
	}
	pivot := quickSortPartition(array, start, end)
	quickSortHelper(array, start, pivot-1)
	quickSortHelper(array, pivot+1, end)
}
func quickSortPartition(array []int, start, end int) int {
	pivot := array[end]

	i := start - 1

	for ; start < end; start++ {
		if array[start] < pivot {
			i++
			array[start], array[i] = array[i], array[start]
		}
	}
	i++
	array[i], array[end] = array[end], array[i]
	return i
}
