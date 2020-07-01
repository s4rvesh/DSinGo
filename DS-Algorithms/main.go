package main

import (
	"fmt"

	"github.com/s4rvesh/DSinGo/DS-Algorithms/input"
	"github.com/s4rvesh/DSinGo/DS-Algorithms/output"
	"github.com/s4rvesh/DSinGo/DS-Algorithms/searching"
	"github.com/s4rvesh/DSinGo/DS-Algorithms/sorting"
)

func main() {

	fmt.Println("------Array programs------")

	array := input.Accept()
	output.Show(array)
	array = sorting.MergeSort(array)

	//array = sorting.QuickSort(array)
	//array = sorting.InsertionSort(array)
	//array = sorting.Bubble(array)

	var search int
	fmt.Println("enter value to search in array")
	fmt.Scanln(&search)

	elementAt := searching.BinarySearch(array, search)

	fmt.Printf("element found at %v", elementAt)

}
