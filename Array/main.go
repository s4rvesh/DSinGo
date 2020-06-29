package main

import (
	"fmt"

	"github.com/s4rvesh/DSinGo/Array/input"
	"github.com/s4rvesh/DSinGo/Array/output"
	"github.com/s4rvesh/DSinGo/Array/sorting"
)

func main() {

	fmt.Println("------Array programs------")

	array := input.Accept()
	output.Show(array)

	sorting.Bubble(array)

}
