package operations

// NewZeroMatrix returns a zero matrix of dimensions i x j.
func NewZeroMatrix(i, j int) Matrix {
	output := make([][]int, i)
	for row := 0; row < i; row++ {
		line := make([]int, j)
		output[row] = line
	}
	return Matrix(output)
}

// NewMatrix returns a matrix of dimensions i x j.
// If the input is smaller than i multiplied by j then it will be repeated to populate the matrix.
// If the input is larger than i multiplied by j then only the earlier entries that fit into the matrix will populate.
func NewMatrix(i, j int, input []int) Matrix {
	m := NewZeroMatrix(i, j)
	product := i * j
	inputLength := len(input)
	for k := 0; k < product; k++ {
		iPos, jPos, kPos := k/j, k%j, k%inputLength
		m[iPos][jPos] = input[kPos]
	}
	return m
}

// NewIdentityMatrix returns an identity matrix of specified dimension
func NewIdentityMatrix(i int) Matrix {
	m := NewZeroMatrix(i, i)
	for k := 0; k < i; k++ {
		m[k][k] = 1
	}
	return m
}

// NewConstantMatrix returns a matrix where all elements == x
func NewConstantMatrix(i, j, x int) Matrix {
	m := NewZeroMatrix(i, j)
	product := i * j
	for k := 0; k < product; k++ {
		iPos, jPos := k/j, k%j
		m[iPos][jPos] = x
	}
	return m
}

// GetIJ returns the dimensions of a supplied matrix
func (m Matrix) GetIJ() (int, int) {
	i, j := len(m), len(m[0])
	return i, j
}
