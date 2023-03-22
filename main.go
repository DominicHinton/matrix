package main

import (
	"fmt"

	"github.com/DominicHinton/matrix/operations"
)

func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := operations.NewMatrix(3, 4, in)
	fmt.Println(m)

	id, err := operations.NewIdentityMatrix(4, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	three := operations.NewConstantMatrix(2, 5, 3)
	fmt.Println(three)
}
