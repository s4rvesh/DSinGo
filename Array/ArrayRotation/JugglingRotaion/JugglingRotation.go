package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	rotation := 3
	gcd := gcd(len(array), rotation)

	for i := 0; i < gcd; i++ {

		temp := array[i]
		loc := 0

		for j := i; j+gcd < len(array); j += gcd {

			array[j] = array[j+gcd]
			loc = j + gcd

		}
		array[loc] = temp

	}

	fmt.Println(array)
}

func gcd(a, b int) int {

	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
