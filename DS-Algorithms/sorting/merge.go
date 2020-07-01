package sorting

import "github.com/s4rvesh/DSinGo/DS-Algorithms/output"

//MergeSort array
func MergeSort(array []int) []int {

	array = sort(array, 0, len(array)-1)

	output.Show(array)
	return array
}

func msort(array []int, left, right int) []int {

	if left < right {

		middle := left + (right-left)/2

		msort(array, left, middle)
		msort(array, middle+1, right)
		merge(array, left, middle, right)

	}

	return array

}

func merge(array []int, left, middle, right int) {

	n1 := middle - left + 1
	n2 := right - middle

	l := make([]int, n1)
	r := make([]int, n2)

	for i := 0; i < n1; i++ {
		l[i] = array[left+i]
	}
	for j := 0; j < n2; j++ {
		r[j] = array[middle+1+j]
	}

	var i, j int
	k := left

	for i < n1 && j < n2 {

		if l[i] <= r[j] {

			array[k] = l[i]
			i++
		} else {
			array[k] = r[j]
			j++
		}
		k++
	}

	for i < n1 {
		array[k] = l[i]
		i++
		k++

	}

	for j < n2 {
		array[k] = r[j]
		j++
		k++
	}

}
