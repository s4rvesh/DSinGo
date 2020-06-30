package searching

import (
	"fmt"
)

//BinarySearch implementation
func BinarySearch(array []int, tosearch int) int {

	fmt.Println("Binary Search")

	a := searching(array, 0, (len(array) - 1), tosearch)

	return a
}

func searching(array []int, left, right, tosearch int) int {

	if right >= left {

		mid := left + (right-left)/2

		if array[mid] == tosearch {
			return mid
		}
		if array[mid] > tosearch {
			return searching(array, left, mid-1, tosearch)
		}
		if array[mid] < tosearch {
			return searching(array, mid+1, right, tosearch)
		}
	}

	return 999

}
