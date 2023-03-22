package operations

// SequentialMapFuncToElements takes fn: a function that returns a result of operation on one element of matrix m,
// x: a second argument for fn and returns new matrix with same operation applied to every element
func SequentialMapFuncToElements(fn MatrixConstantSequentialOperation, m Matrix, x int) Matrix {
	i, j := GetIJ(m)
	product := i * j
	output := NewZeroMatrix(i, j)
	for k := 0; k < product; k++ {
		iPos, jPos := k/j, k%j
		output[iPos][jPos] = fn(m[iPos][jPos], x)
	}
	return output
}

// SequentialAddToElements adds x to every element in m sequentially and returns resulting matrix
func SequentialAddToElements(m Matrix, x int) Matrix {
	return SequentialMapFuncToElements(Add, m, x)
}

func Add(a, b int) int {
	return a + b
}
