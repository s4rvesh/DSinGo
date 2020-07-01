package sorting

import "github.com/s4rvesh/DSinGo/DS-Algorithms/output"

//SelectionSort on array
func SelectionSort(array []int) []int {

	for i := 0; i < len(array); i++ {

		min := i

		for j := i + 1; j < len(array); j++ {

			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}

	output.Show(array)

	return array

}
