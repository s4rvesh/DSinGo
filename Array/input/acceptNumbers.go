package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Accept takes in input
func Accept() []int {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" enter elements in array")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var array []int

	strarray := strings.Fields(text)

	for _, i := range strarray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		array = append(array, j)
	}

	return array
}
