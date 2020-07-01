package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	rotate := 3
	temp := make([]int, rotate)

	for i := 0; i < rotate; i++ {

		temp[i] = array[i]

	}

	i := 0
	for j := i + rotate; j < len(array); j++ {

		array[i] = array[j]
		i++
	}

	for j := 0; j < rotate; j++ {

		array[i] = temp[j]
		i++
	}

	fmt.Println(array)

}
