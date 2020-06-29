package sorting

import (
	"fmt"

	"github.com/s4rvesh/DSinGo/Array/output"
)

//Bubble sort algorithm implementation
func Bubble(array []int) {

	fmt.Println("Sorting using Bubble Sort")

	for i := len(array) - 1; i > 1; i-- {

		for x := 0; x < i; x++ {
			y := x + 1

			if array[x] > array[y] {

				array[x], array[y] = array[y], array[x]
			}

		}

	}

	output.Show(array)

}
