package concoperations

import (
	"sync"
)

// packages

// matrices, vectors, operations

// NewZeroMatrix returns a zero matrix of dimensions i x j.
// If i or j is supplied as a negative, an empty matrix is returned.
func NewZeroMatrix[N Number](i, j int) Matrix[N] {
	if (i < 0) || (j < 0) {
		return Matrix[N]{}
	}

	var wg sync.WaitGroup
	wg.Add(i)
	m := make(Matrix[N], i)

	for row := 0; row < i; row++ {
		r := row
		go func() {
			defer wg.Done()
			line := make([]N, j)
			m[r] = line
		}()
	}
	wg.Wait()

	return m
}

// NewMatrixFromSlice returns a matrix of dimensions i x j.
// If the input is smaller than i multiplied by j then it will be repeated to populate the matrix.
// If the input is larger than i multiplied by j then only the earlier entries that fit into the matrix will populate.
func NewMatrixFromSlice[N Number](i, j int, input []N) Matrix[N] {
	if (i < 0) || (j < 0) || (len(input) == 0) {
		return Matrix[N]{}
	}
	m := NewZeroMatrix[N](i, j)
	product := i * j
	inputLength := len(input)
	var wg sync.WaitGroup
	wg.Add(product)
	for g := 0; g < product; g++ {
		k := g
		go func() {
			defer wg.Done()
			iPos, jPos, kPos := k/j, k%j, k%inputLength
			m[iPos][jPos] = input[kPos]
		}()
	}
	wg.Wait()
	return m
}

// NewIdentityMatrix returns an identity matrix of specified dimension
func NewIdentityMatrix[N Number](dimension int) Matrix[N] {
	if dimension < 1 {
		return Matrix[N]{}
	}
	m := NewZeroMatrix[N](dimension, dimension)
	one := N(1)
	var wg sync.WaitGroup
	wg.Add(dimension)
	for d := 0; d < dimension; d++ {
		k := d
		go func() {
			defer wg.Done()
			m[k][k] = one
		}()
	}
	wg.Wait()
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
	var wg sync.WaitGroup
	wg.Add(rows * columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			a, b := i, j
			go func() {
				defer wg.Done()
				m[a][b] = x
			}()
		}
	}
	wg.Wait()
}
