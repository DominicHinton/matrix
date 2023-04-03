package seqoperations

import "math"

// SequentialMapFuncToElements takes fn: a function that returns a result of operation on one element of matrix m,
// x: a second argument for fn and returns new matrix with same operation applied to every element
func (m Matrix[N]) SequentialMapFuncToElements(fn ConstantSequentialOperater[N]) Matrix[N] {
	i, j := m.GetDimensions()
	product := i * j
	output := NewZeroMatrix[N](i, j)
	for k := 0; k < product; k++ {
		iPos, jPos := k/j, k%j
		output[iPos][jPos] = fn(m[iPos][jPos])
	}
	return output
}

// SequentialOneToOne requires two matrices of same dimensions.
// takes fn: a function that is applied element wise to each element Mij and Nij.
// in the resulting matrix : Pij = fn(Mij, Nij)
func (m Matrix[N]) SequentialOneToOne(n Matrix[N], fn OneToOneSequentialOperater[N]) (Matrix[N], error) {
	rows, columns := m.GetDimensions()
	rowsCheck, columnsCheck := n.GetDimensions()
	if (rows != rowsCheck) || (columns != columnsCheck) {
		return nil, errDifferentDimension
	}
	p := NewZeroMatrix[N](rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			p[i][j] = fn(m[i][j], n[i][j])
		}
	}
	return p, nil
}

// SequentialAddToElements adds x to every element in m sequentially and returns resulting matrix
func (m Matrix[N]) SequentialAddToElements(x N) Matrix[N] {
	return m.SequentialMapFuncToElements(func(element N) N { return element + x })
}

// SequentialSubtractFromElements subtracts x from every element in m sequentially and returns resulting matrix
func (m Matrix[N]) SequentialSubtractFromElements(x N) Matrix[N] {
	return m.SequentialMapFuncToElements(func(element N) N { return element - x })
}

// SequentialSubtractElementsFrom subtracts every element in m from x sequentially and returns resulting matrix
func (m Matrix[N]) SequentialSubtractElementsFrom(x N) Matrix[N] {
	return m.SequentialMapFuncToElements(func(element N) N { return x - element })
}

// SequentialMultiplyElementsBy multiplies x by every element in m sequentially and returns resulting matrix
func (m Matrix[N]) SequentialMultiplyElementsBy(x N) Matrix[N] {
	return m.SequentialMapFuncToElements(func(element N) N { return x * element })
}

// SequentialDivideElementsBy divides every element in m by x sequentially and returns resulting matrix
func (m Matrix[N]) SequentialDivideElementsBy(x N) Matrix[N] {
	return m.SequentialMapFuncToElements(func(element N) N { return element / x })
}

// SequentialDivideByElements divides x by each element in m and returns resulting matrix
func (m Matrix[N]) SequentialDivideByElements(x N) Matrix[N] {
	return m.SequentialMapFuncToElements(func(element N) N { return x / element })
}

// SequentialAddMatrices adds two matrices together and returns resulting matrix if the addition is valid
func (m Matrix[N]) SequentialAddMatrices(n Matrix[N]) (Matrix[N], error) {
	out, err := m.SequentialOneToOne(n, Add[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// SequentialSubtractMatrices returns Matrix P such that Pij = Mij - Nij if subtraction is valid
func (m Matrix[N]) SequentialSubtractMatrices(n Matrix[N]) (Matrix[N], error) {
	out, err := m.SequentialOneToOne(n, Subtract[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// SequentialElementWiseMultiply requires two matrices of same dimensions. Returns Matrix P where
// Pij = Mij x Nij
func (m Matrix[N]) SequentialElementWiseMultiply(n Matrix[N]) (Matrix[N], error) {
	out, err := m.SequentialOneToOne(n, Multiply[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// SequentialElementWiseDivide requires two matrices of same dimensions. Returns Matrix P where
// Pij = Mij / Nij
func (m Matrix[N]) SequentialElementWiseDivide(n Matrix[N]) (Matrix[N], error) {
	out, err := m.SequentialOneToOne(n, Divide[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// SequentialMean returns mean of matrix m if m contains elements
func (m Matrix[N]) SequentialMean() (N, bool) {
	rows, columns := m.GetDimensions()
	var total N
	if rows == 0 || columns == 0 {
		return total, false
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			total += m[i][j]
		}
	}
	return total / (N(rows) * N(columns)), true
}

func (m Matrix[N]) SequentialMeanStandardDev() (N, N, bool) {
	mean, ok := m.SequentialMean()
	if !ok {
		return mean, mean, false
	}
	var total N
	rows, columns := m.GetDimensions()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			diff := m[i][j] - mean
			diffSquared := diff * diff
			total += diffSquared
		}
	}
	total = total / ((N(rows) * N(columns)) - 1)
	sqrt := math.Sqrt(float64(total))
	return mean, N(sqrt), true
}

// SequentialMultiplyMatrices returns matrix P = M * N if multiplication is valid
func (m Matrix[N]) SequentialMultiplyMatrices(n Matrix[N]) (Matrix[N], error) {
	rows, columns, ok := m.ValidMultiplication(n)
	if !ok {
		return nil, errMultiplicationValidity
	}
	out := NewZeroMatrix[N](rows, columns)
	for i := 0; i < rows; i++ {
		// Validity of multiplication handled at start of method, booleans from vector ops are discarded
		rowVector, _ := m.GetVectorFromRow(i)
		for j := 0; j < columns; j++ {
			columnVector, _ := n.GetVectorFromColumn(j)
			out[i][j], _ = rowVector.DotProduct(columnVector)
		}
	}
	return out, nil
}

func (m Matrix[N]) SequentialTranspose() Matrix[N] {
	x, y := m.GetDimensions()
	t := NewZeroMatrix[N](y, x)
	entries := x * y
	for k := 0; k < entries; k++ {
		i, j := k%x, k/x
		t[j][i] = m[i][j]
	}
	return t
}

// SameDimensions returns true if m and n have same dimensions, false otherwise
func (m Matrix[N]) SameDimensions(n Matrix[N]) bool {
	mi, mj := m.GetDimensions()
	ni, nj := n.GetDimensions()
	return (mi == ni) && (mj == nj)
}

// ValidMultiplication returns required dimensions of multiplication result and true if valid, false if not valid
func (m Matrix[N]) ValidMultiplication(n Matrix[N]) (int, int, bool) {
	a, b := m.GetDimensions()
	c, d := n.GetDimensions()
	if b != c {
		return -1, -1, false
	}
	return a, d, true
}

func Add[N Number](a, b N) N {
	return a + b
}

func Subtract[N Number](a, b N) N {
	return a - b
}

func ReverseSubtract[N Number](a, b N) N {
	return b - a
}

func Multiply[N Number](a, b N) N {
	return a * b
}

func Divide[N Number](a, b N) N {
	return a / b
}

func ReverseDivide[N Number](a, b N) N {
	return b / a
}
