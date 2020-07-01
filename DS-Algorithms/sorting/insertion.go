package sorting

import "github.com/s4rvesh/DSinGo/DS-Algorithms/output"

//InsertionSort the array
func InsertionSort(array []int) []int {

	for i := 1; i < len(array); i++ {

		key := array[i]
		j := i - 1

		for j >= 0 && array[j] > key {

			array[j+1] = array[j]
			j = j - 1

		}
		array[j+1] = key
	}

	output.Show(array)

	return array
}
