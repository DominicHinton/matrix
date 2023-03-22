package operations

// SequentialMapFuncToElements takes fn: a function that returns a result of operation on one element of matrix m,
// x: a second argument for fn and returns new matrix with same operation applied to every element
func (m Matrix) SequentialMapFuncToElements(fn MatrixConstantSequentialOperation, x int) Matrix {
	i, j := m.GetIJ()
	product := i * j
	output := NewZeroMatrix(i, j)
	for k := 0; k < product; k++ {
		iPos, jPos := k/j, k%j
		output[iPos][jPos] = fn(m[iPos][jPos], x)
	}
	return output
}

// SequentialAddToElements adds x to every element in m sequentially and returns resulting matrix
func (m Matrix) SequentialAddToElements(x int) Matrix {
	return m.SequentialMapFuncToElements(Add, x)
}

// SequentialSubtractFromElements subtracts x from every element in m sequentially and returns resulting matrix
func (m Matrix) SequentialSubtractFromElements(x int) Matrix {
	return m.SequentialMapFuncToElements(Subtract, x)
}

// SequentialSubtractElementsFrom subtracts x to every element in m sequentially and returns resulting matrix
func (m Matrix) SequentialSubtractElementsFrom(x int) Matrix {
	return m.SequentialMapFuncToElements(ReverseSubtract, x)
}

// SequentialMultiplyElementsBy multiplies x by every element in m sequentially and returns resulting matrix
func (m Matrix) SequentialMultiplyElementsBy(x int) Matrix {
	return m.SequentialMapFuncToElements(Multiply, x)
}

// SequentialDivideElementsBy divides every element in m by x sequentially and returns resulting matrix
func (m Matrix) SequentialDivideElementsBy(x int) Matrix {
	return m.SequentialMapFuncToElements(Divide, x)
}

// SequentialDivideByElements divides x by each element in m and returns resulting matrix
func (m Matrix) SequentialDivideByElements(x int) Matrix {
	return m.SequentialMapFuncToElements(ReverseDivide, x)
}

// SequentialAddMatrices adds two matrices together and returns resulting matrix if the addition is valid
func (m Matrix) SequentialAddMatrices(n Matrix) (Matrix, error) {
	ok := m.SameDimensions(n)
	if !ok {
		return nil, errDifferentDimesion
	}
	i, j := m.GetIJ()
	out := NewZeroMatrix(i, j)
	for k := 0; k < i*j; k++ {
		iPos, jPos := k/j, k%j
		out[iPos][jPos] = m[iPos][jPos] + n[iPos][jPos]
	}
	return out, nil
}

// SequentialSubtractMatrices returns Matrix P such that Pij = Mij - Nij if subtraction is valid
func (m Matrix) SequentialSubtractMatrices(n Matrix) (Matrix, error) {
	ok := m.SameDimensions(n)
	if !ok {
		return nil, errDifferentDimesion
	}
	i, j := m.GetIJ()
	out := NewZeroMatrix(i, j)
	for k := 0; k < i*j; k++ {
		iPos, jPos := k/j, k%j
		out[iPos][jPos] = m[iPos][jPos] - n[iPos][jPos]
	}
	return out, nil
}

// SequentialMultiplyMatrices returns matrix P = M * N if multiplication is valid
func (m Matrix) SequentialMultiplyMatrices(n Matrix) (Matrix, error) {
	i, j, ok := m.ValidMultiplication(n)
	if !ok {
		return nil, errMultiplicationValidity
	}
	out := NewZeroMatrix(i, j)
	entries := i * j
	for k := 0; k < entries; k++ {
		iPos, jPos := k/j, k%j
		// Validity of multiplication handled at start of method, booleans from vector ops are discarded
		rowVector, _ := m.GetVectorFromRow(iPos)
		columnVector, _ := n.GetVectorFromColumn(jPos)
		out[iPos][jPos], _ = rowVector.DotProduct(columnVector)
	}
	return out, nil
}

func (m Matrix) SequentialTranspose() Matrix {
	x, y := m.GetIJ()
	t := NewZeroMatrix(y, x)
	entries := x * y
	for k := 0; k < entries; k++ {
		i, j := k%x, k/x
		t[j][i] = m[i][j]
	}
	return t
}

// SameDimensions returns true if m and n have same dimensions, false otherwise
func (m Matrix) SameDimensions(n Matrix) bool {
	mi, mj := m.GetIJ()
	ni, nj := n.GetIJ()
	return (mi == ni) && (mj == nj)
}

// ValidMultiplication returns required dimensions of multiplication result and true if valid, false if not valid
func (m Matrix) ValidMultiplication(n Matrix) (int, int, bool) {
	a, b := m.GetIJ()
	c, d := n.GetIJ()
	if b != c {
		return -1, -1, false
	}
	return a, d, true
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func ReverseSubtract(a, b int) int {
	return b - a
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func ReverseDivide(a, b int) int {
	return b / a
}
