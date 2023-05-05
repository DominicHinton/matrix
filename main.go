package main

import (
	"fmt"

	conc "github.com/DominicHinton/matrix/concoperations"
)

// TODO - change to unit tests

func main() {

	zeros := conc.NewZeroMatrix[int](3, 4)
	fmt.Println(zeros)

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := conc.NewMatrixFromSlice(3, 4, in)
	fmt.Println(m)

	id := conc.NewIdentityMatrix[int](4)
	fmt.Println(id)

	three := conc.NewConstantMatrix(2, 5, 3)
	fmt.Println(three)

	added3 := m.AddToElements(3)
	subtracted4 := m.SubtractFromElements(4)
	fiveSub := m.SubtractElementsFrom(5)
	mult6 := m.MultiplyElementsBy(6)
	div2 := m.DivideElementsBy(2)
	fourDiv := m.DivideByElements(4)

	fmt.Println(m, "\nm + 3\n", added3)
	fmt.Println(m, "\nm - 4\n", subtracted4)
	fmt.Println(m, "\n5 - m\n", fiveSub)
	fmt.Println(m, "\nm * 6\n", mult6)
	fmt.Println(m, "\nm / 2\n", div2)
	fmt.Println(m, "\n4 / m\n", fourDiv)

	oneTimes := conc.NewMatrixFromSlice(2, 3, []int{1, 2, 3, 4, 5, 6})
	tenTimes := conc.NewMatrixFromSlice(2, 3, []int{10, 20, 30, 40, 50, 60})

	addedMatrices, err := oneTimes.AddMatrices(tenTimes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(oneTimes, "\n+\n", tenTimes, "\n=\n", addedMatrices)
	}

	subbedMatrices, err := oneTimes.SubtractMatrices(tenTimes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(oneTimes, "\n-\n", tenTimes, "\n=\n", subbedMatrices)
	}

	v, _ := m.VectorFromColumn(2)
	fmt.Println("m:\n", m, "\nv from col 2:\n", v)

	w, _ := m.VectorFromRow(2)
	fmt.Println("m:\n", m, "\nv from row 2:\n", w)

	A := conc.NewMatrixFromSlice(2, 3, []int{2, 3, 4, 1, 0, 0})
	B := conc.NewMatrixFromSlice(3, 2, []int{0, 1000, 1, 100, 0, 10})
	AB, _ := A.Multiply(B)
	fmt.Println("A:\n", A, "\nB:\n", B, "\nAB:\n", AB)

	Q := conc.NewMatrixFromSlice(2, 3, []int{1, 2, 3, 0, -6, 7})
	QT := Q.SequentialTranspose()
	fmt.Println("Q:\n", Q, "\nQT:\n", QT)

	fmt.Printf("%v", Q)

	J := conc.NewMatrixFromSlice(2, 3, []int{2, 3, 4, 1, 0, 0})
	fmt.Printf("\n%v\n", J)
	J.FillMatrix(-1)
	fmt.Printf("\n%v\n", J)

	m = conc.NewMatrixFromSlice(3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 10})
	// inverse of m = [[-2/3, -4/3, 1], [-2/3, 11/3, -2,], [1, -2, 1]]
	fmt.Println(m)
	inv, err := m.Inverse()
	if err != nil {
		fmt.Println("no inverse")
	} else {
		fmt.Println(inv)

	}
	det, _ := m.Determinant()
	fmt.Println("det = ", det)
	n := conc.NewMatrixFromSlice(3, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// n has no inverse
	fmt.Println(n)
	inv, err = n.Inverse()
	if err != nil {
		fmt.Println("no inverse")
	} else {
		fmt.Printf("\n&&&&%#v\n", inv)

	}
	det, _ = n.Determinant()
	fmt.Println("det = ", det)

	empty := conc.Matrix[int]{}
	i, j := empty.Dimensions()
	fmt.Println(i, j)
}
