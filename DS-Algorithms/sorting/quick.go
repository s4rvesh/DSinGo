package sorting

import "github.com/s4rvesh/DSinGo/DS-Algorithms/output"

//QuickSort array
func QuickSort(array []int) []int {

	array = sort(array, 0, len(array)-1)

	output.Show(array)
	return array
}

func sort(array []int, start, end int) []int {

	if end > start {

		index := partition(array, start, end)

		sort(array, start, index-1)
		sort(array, index+1, end)

	}

	return array
}

func partition(array []int, start, end int) int {

	pivot := array[end]

	i := start - 1

	for j := start; j <= end; j++ {

		if array[j] < pivot {
			i++
			array[j], array[i] = array[i], array[j]
		}

	}

	array[i+1], array[end] = array[end], array[i+1]

	return i + 1

}
