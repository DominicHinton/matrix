package seqoperations

// packages

// matrices, vectors, operations

// NewZeroMatrix returns a zero matrix of dimensions i x j.
// If i or j is supplied as a negative, an empty matrix is returned.
func NewZeroMatrix[N Number](i, j int) Matrix[N] {
	if (i < 0) || (j < 0) {
		return Matrix[N]{}
	}
	output := make(Matrix[N], i)
	for row := 0; row < i; row++ {
		line := make([]N, j)
		output[row] = line
	}
	return output
}

// NewMatrix returns a matrix of dimensions i x j.
// If the input is smaller than i multiplied by j then it will be repeated to populate the matrix.
// If the input is larger than i multiplied by j then only the earlier entries that fit into the matrix will populate.

// Consider - matrix pool, to store matrices for later use
func NewMatrixFromSlice[N Number](i, j int, input []N) Matrix[N] {
	if (i < 0) || (j < 0) || (len(input) == 0) {
		return Matrix[N]{}
	}
	m := NewZeroMatrix[N](i, j)
	product := i * j
	inputLength := len(input)
	for k := 0; k < product; k++ {
		iPos, jPos, kPos := k/j, k%j, k%inputLength
		m[iPos][jPos] = input[kPos]
	}
	return m
}

// NewIdentityMatrix returns an identity matrix of specified dimension
func NewIdentityMatrix[N Number](dimension int) Matrix[N] {
	if dimension < 1 {
		return Matrix[N]{}
	}
	m := NewZeroMatrix[N](dimension, dimension)
	one := N(1)
	for k := 0; k < dimension; k++ {
		m[k][k] = one
	}
	return m
}

// NewConstantMatrix returns a matrix where all elements == x
func NewConstantMatrix[N Number](iDim, jDim int, x N) Matrix[N] {
	m := NewZeroMatrix[N](iDim, jDim)
	m.FillMatrix(x)
	return m
}

// FillMatrix fills a matrix in order that every element is the supplied constant
func (m Matrix[N]) FillMatrix(x N) {
	rows, columns := m.Dimensions()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			m[i][j] = x
		}
	}
}
