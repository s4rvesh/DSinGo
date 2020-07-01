package main

import "fmt"

func main() {

	array := []int{4, 5, 6, 7, 8, 9, 1, 2, 3}

	key := 3
	index := search(array, 0, len(array)-1, key)

	if index != -1 {

		fmt.Printf("Element found at %v", index)
	} else {
		fmt.Println("Element NOT FOUND")
	}

}

func search(array []int, l, h, key int) int {

	if l > h {
		return -1
	}

	var mid int
	mid = (l + h) / 2

	if array[mid] == key {
		return mid
	}

	if array[l] < array[mid] {

		if key >= array[l] && key <= array[mid] {

			return search(array, l, mid-1, key)

		}
		return search(array, mid+1, h, key)

	}

	if key >= array[mid] && key <= array[h] {

		return search(array, mid+1, h, key)
	}

	return search(array, l, mid-1, key)

}
