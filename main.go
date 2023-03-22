package main

import (
	"fmt"

	"github.com/DominicHinton/matrix/operations"
)

func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := operations.NewMatrix(3, 4, in)
	fmt.Println(m)

	id := operations.NewIdentityMatrix(4)
	fmt.Println(id)

	three := operations.NewConstantMatrix(2, 5, 3)
	fmt.Println(three)

	added3 := m.SequentialAddToElements(3)
	subtracted4 := m.SequentialSubtractFromElements(4)
	fiveSub := m.SequentialSubtractElementsFrom(5)
	mult6 := m.SequentialMultiplyElementsBy(6)
	div2 := m.SequentialDivideElementsBy(2)
	fourDiv := m.SequentialDivideByElements(4)

	fmt.Println(m, "\nm + 3\n", added3)
	fmt.Println(m, "\nm - 4\n", subtracted4)
	fmt.Println(m, "\n5 - m\n", fiveSub)
	fmt.Println(m, "\nm * 6\n", mult6)
	fmt.Println(m, "\nm / 2\n", div2)
	fmt.Println(m, "\n4 / m\n", fourDiv)

	oneTimes := operations.NewMatrix(2, 3, []int{1, 2, 3, 4, 5, 6})
	tenTimes := operations.NewMatrix(2, 3, []int{10, 20, 30, 40, 50, 60})

	addedMatrices, err := oneTimes.SequentialAddMatrices(tenTimes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(oneTimes, "\n+\n", tenTimes, "\n=\n", addedMatrices)
	}

	subbedMatrices, err := oneTimes.SequentialSubtractMatrices(tenTimes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(oneTimes, "\n-\n", tenTimes, "\n=\n", subbedMatrices)
	}

	v, _ := m.GetVectorFromColumn(2)
	fmt.Println("m:\n", m, "\nv from col 2:\n", v)

	w, _ := m.GetVectorFromRow(2)
	fmt.Println("m:\n", m, "\nv from row 2:\n", w)

	A := operations.NewMatrix(2, 3, []int{2, 3, 4, 1, 0, 0})
	B := operations.NewMatrix(3, 2, []int{0, 1000, 1, 100, 0, 10})
	AB, _ := A.SequentialMultiplyMatrices(B)
	fmt.Println("A:\n", A, "\nB:\n", B, "\nAB:\n", AB)

	Q := operations.NewMatrix(2, 3, []int{1, 2, 3, 0, -6, 7})
	QT := Q.SequentialTranspose()
	fmt.Println("Q:\n", Q, "\nQT:\n", QT)

	fmt.Printf("%s", Q)
}
