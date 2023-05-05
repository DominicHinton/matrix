package concoperations

import (
	"math"
	"sync"
)

// Dimensions returns the dimensions of a supplied matrix.
func (m Matrix[N]) Dimensions() (int, int) {
	if len(m) == 0 {
		return 0, 0
	}
	return len(m), len(m[0])
}

// IsSquare returns true is matrix is square, false otherwise.
// An empty matrix with dimensions 0x0 is assumed to be square.
// A matrix of empty row vectors is assumed to be not square.
func (m Matrix[N]) IsSquare() bool {
	rows, columns := m.Dimensions()
	return rows == columns
}

// MapFunctionToElements takes fn: a function that returns a result of operation on one element of matrix m,
// x: a second argument for fn and returns new matrix with same operation applied to every element
func (m Matrix[N]) MapFunctionToElements(fn ConstantSequentialOperater[N]) Matrix[N] {

	rows, columns := m.Dimensions()
	output := NewZeroMatrix[N](rows, columns)

	var wg sync.WaitGroup
	wg.Add(rows * columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			a, b := i, j
			go func() {
				defer wg.Done()
				output[a][b] = fn(m[a][b])
			}()
		}
	}
	wg.Wait()

	return output
}

// MapFunctionToElementsInRowInPlace takes a function to be applied to every element of a row
// and a row number and then applies the function in place
func (m Matrix[N]) MapFunctionToElementsInRowInPlace(fn ConstantSequentialOperater[N], row int) error {
	rows, cols := m.Dimensions()
	if row >= rows || row < 0 {
		return errRowColSuppliedOutBounds
	}
	var wg sync.WaitGroup
	wg.Add(cols)
	for j := 0; j < cols; j++ {
		k := j
		go func() {
			defer wg.Done()
			m[row][k] = fn(m[row][k])
		}()
	}
	wg.Wait()
	return nil
}

func (m Matrix[N]) RowScalarMultiply(row int, multiplier N) error {
	err := m.MapFunctionToElementsInRowInPlace(func(element N) N { return multiplier * element }, row)
	return err
}

// ApplyOneToOne requires two matrices of same dimensions.
// takes fn: a function that is applied element wise to each element Mij and Nij.
// in the resulting matrix : Pij = fn(Mij, Nij)
func (m Matrix[N]) ApplyOneToOne(n Matrix[N], fn OneToOneSequentialOperater[N]) (Matrix[N], error) {
	rows, columns := m.Dimensions()
	rowsCheck, columnsCheck := n.Dimensions()
	if (rows != rowsCheck) || (columns != columnsCheck) {
		return nil, errDifferentDimension
	}
	p := NewZeroMatrix[N](rows, columns)
	var wg sync.WaitGroup
	wg.Add(rows * columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			a, b := i, j
			go func() {
				defer wg.Done()
				p[a][b] = fn(m[a][b], n[a][b])
			}()
		}
	}
	wg.Wait()
	return p, nil
}

// AddToElements adds x to every element in m sequentially and returns resulting matrix
func (m Matrix[N]) AddToElements(x N) Matrix[N] {
	return m.MapFunctionToElements(func(element N) N { return element + x })
}

// SubtractFromElements subtracts x from every element in m sequentially and returns resulting matrix
func (m Matrix[N]) SubtractFromElements(x N) Matrix[N] {
	return m.MapFunctionToElements(func(element N) N { return element - x })
}

// SubtractElementsFrom subtracts every element in m from x sequentially and returns resulting matrix
func (m Matrix[N]) SubtractElementsFrom(x N) Matrix[N] {
	return m.MapFunctionToElements(func(element N) N { return x - element })
}

// MultiplyElementsBy multiplies x by every element in m sequentially and returns resulting matrix
func (m Matrix[N]) MultiplyElementsBy(x N) Matrix[N] {
	return m.MapFunctionToElements(func(element N) N { return x * element })
}

// DivideElementsBy divides every element in m by x sequentially and returns resulting matrix
func (m Matrix[N]) DivideElementsBy(x N) Matrix[N] {
	return m.MapFunctionToElements(func(element N) N { return element / x })
}

// DivideByElements divides x by each element in m and returns resulting matrix
func (m Matrix[N]) DivideByElements(x N) Matrix[N] {
	return m.MapFunctionToElements(func(element N) N { return x / element })
}

// AddMatrices adds two matrices together and returns resulting matrix if the addition is valid
func (m Matrix[N]) AddMatrices(n Matrix[N]) (Matrix[N], error) {
	out, err := m.ApplyOneToOne(n, Add[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// SubtractMatrices returns Matrix P such that Pij = Mij - Nij if subtraction is valid
func (m Matrix[N]) SubtractMatrices(n Matrix[N]) (Matrix[N], error) {
	out, err := m.ApplyOneToOne(n, Subtract[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// ElementWiseMultiply requires two matrices of same dimensions. Returns Matrix P where
// Pij = Mij x Nij
func (m Matrix[N]) ElementWiseMultiply(n Matrix[N]) (Matrix[N], error) {
	out, err := m.ApplyOneToOne(n, Multiply[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// ElementWiseDivide requires two matrices of same dimensions. Returns Matrix P where
// Pij = Mij / Nij
func (m Matrix[N]) ElementWiseDivide(n Matrix[N]) (Matrix[N], error) {
	out, err := m.ApplyOneToOne(n, Divide[N])
	if err != nil {
		return Matrix[N]{}, err
	}
	return out, nil
}

// Mean returns float64 mean of matrix m if m contains elements
func (m Matrix[N]) Mean() (float64, bool) {

	rows, columns := m.Dimensions()
	var total float64
	if rows == 0 || columns == 0 {
		return total, false
	}

	mCopy := m.Float64Copy()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			total += mCopy[i][j]
		}
	}

	return total / (float64(rows) * float64(columns)), true
}

// MeanStandardDev returns the mean and standard deviation of
// a matrix if it contains elements
func (m Matrix[N]) MeanStandardDev() (float64, float64, bool) {

	rows, columns := m.Dimensions()
	if rows == 0 || columns == 0 {
		return 0.0, 0.0, false
	}

	// instantiate total to sum elements in matrix
	var total float64
	mCopy := m.Float64Copy()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			total += mCopy[i][j]
		}
	}
	rowsTimesColumns := float64(rows) * float64(columns)
	mean := total / rowsTimesColumns

	// reinstantiate total to 0.0 as it will now be used to some square of dist from mean
	total = 0.0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			diff := mCopy[i][j] - mean
			diffSquared := diff * diff
			total += diffSquared
		}
	}
	variance := total / ((rowsTimesColumns) - 1.0)
	stdev := math.Sqrt(variance)

	return mean, stdev, true
}

// Multiply returns matrix P = M * N if multiplication is valid
func (m Matrix[N]) Multiply(n Matrix[N]) (Matrix[N], error) {

	rows, columns, ok := m.MultiplicationDimensions(n)
	if !ok {
		return nil, errMultiplicationValidity
	}

	out := NewZeroMatrix[N](rows, columns)
	for i := 0; i < rows; i++ {
		// Validity of multiplication handled at start of method, booleans from vector ops are discarded
		rowVector, _ := m.VectorFromRow(i)
		for j := 0; j < columns; j++ {
			columnVector, _ := n.VectorFromColumn(j)
			out[i][j], _ = rowVector.DotProduct(columnVector)
		}
	}

	return out, nil
}

func (m Matrix[N]) Inverse() (Matrix[float64], error) {
	return m.InverseAssumeAnyTypeInput()
}

// InverseAssumeAnyTypeInput returns a float64 matrix representing the inverse of the supplied matrix.
func (m Matrix[N]) InverseAssumeAnyTypeInput() (Matrix[float64], error) {
	matrix := m.Float64Copy()
	return matrix.InverseAssumeFloat64Input()
}

// InverseAssumeFloat64Input returns a float64 matrix representing the inverse of the supplied matrix.
// The method requires a float64 matrix to operate.
func (m Matrix[N]) InverseAssumeFloat64Input() (Matrix[N], error) {

	det, err := m.DeterminantAssumeFloat64Input() // this will catch N != float64

	if (err == errZeroLength) || (err == errNonSquare) || (err == errNotFloat64) {
		return Matrix[N]{}, err
	}
	if err != nil {
		return Matrix[N]{}, errUnexpected
	}
	if det == 0 {
		return Matrix[N]{}, errNoInverse
	}

	// instantiate inverse matrix
	dimension := len(m)
	inverseMatrix := NewIdentityMatrix[N](dimension)

	// Apply Gauss-Jordan elimination to find the inverse
	for i := 0; i < dimension; i++ {
		// Divide row i by matrix[i][i]
		factor := 1.0 / m[i][i]
		for j := 0; j < dimension; j++ {
			m[i][j] *= factor
			inverseMatrix[i][j] *= factor
		}

		// Subtract row i from all other rows to make matrix[i][i] equal to 1
		for j := 0; j < dimension; j++ {
			if j != i {
				factor := m[j][i]
				for k := 0; k < dimension; k++ {
					m[j][k] -= factor * m[i][k]
					inverseMatrix[j][k] -= factor * inverseMatrix[i][k]
				}
			}
		}
	}

	return inverseMatrix, nil
}

// Determinant returns the determinant of a matrix as a float64 value
func (m Matrix[N]) Determinant() (float64, error) {
	return m.DeterminantAssumeAnyTypeInput()
}

// DeterminantAssumeAnyTypeInput returns the determinant of a matrix of
// any number type as a float64 value
func (m Matrix[N]) DeterminantAssumeAnyTypeInput() (float64, error) {
	matrix := m.Float64Copy()
	return matrix.DeterminantAssumeFloat64Input()
}

// DeterminantAssumeFloat64Input returns determinant of a float64 matrix as a float64.
// If matrix is non-square or of length zero, an error is returned.
// Although signature allows any type of Number output, type assertion carried out
// in method assures that the Number output will be of type float64 if no error is returned.
func (m Matrix[N]) DeterminantAssumeFloat64Input() (N, error) {

	checkType := N(0)
	isAssumedInputType := IsFloat64(checkType)
	if !isAssumedInputType {
		return 0.0, errNotFloat64
	}
	// N must be type float64 if this line is reached

	// return errors if determinant does not exist
	isSquare := m.IsSquare()
	if !isSquare {
		return 0.0, errNonSquare
	}

	n := len(m)
	if n == 0 {
		return 0.0, errZeroLength
	}

	if n == 1 {
		return m[0][0], nil
	}

	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0], nil
	}

	det := N(0)
	for i := 0; i < n; i++ {
		submatrix := NewZeroMatrix[N](n-1, n-1)
		for j := 1; j < n; j++ {
			for k := 0; k < n; k++ {
				if k < i {
					submatrix[j-1][k] = m[j][k]
				} else if k > i {
					submatrix[j-1][k-1] = m[j][k]
				}
			}
		}

		sign := N(1)
		if i%2 != 0 {
			sign = -N(1)
		}
		subMatrixDeterminant, _ := submatrix.DeterminantAssumeFloat64Input()
		det += sign * m[0][i] * subMatrixDeterminant
	}
	return det, nil
}

// SwapRows swaps row1 and row 2 of matrix m in situ
func (m Matrix[N]) SwapRows(row1, row2 int) error {
	if row1 < 0 || row2 < 0 {
		return errRowColSuppliedOutBounds
	}
	rows, _ := m.Dimensions()
	if row1 >= rows || row2 >= rows {
		return errRowColSuppliedOutBounds
	}
	m[row1], m[row2] = m[row2], m[row1]
	return nil
}

func (m Matrix[N]) FindPivot(row, column int) int {
	rows, _ := m.Dimensions()
	for i := row; i < rows; i++ {
		if m[i][column] != N(0) {
			return i
		}
	}
	return -1
}

// SubMatrix returns a copied sub matrix
func (m Matrix[N]) SubMatrix(rowMin, colMin, rowMax, colMax int) (Matrix[N], error) {
	rows, columns := m.Dimensions()
	if rowMin < 0 || rowMin > rowMax || rowMax >= rows || colMin < 0 || colMin > colMax || colMax >= columns {
		return Matrix[N]{}, errRowColSuppliedOutBounds
	}
	submatrix := NewZeroMatrix[N](rows, columns)
	var wg sync.WaitGroup

	for i := rowMin; i < rowMax; i++ {
		for j := colMin; j < colMax; j++ {
			wg.Add(1)
			a, b := i, j
			go func() {
				defer wg.Done()
				submatrix[a-rowMin][b-colMin] = m[a][b]
			}()
		}
	}
	wg.Wait()
	return submatrix, nil
}

// Copy returns a copy of supplied matrix
func (m Matrix[N]) Copy() Matrix[N] {
	rows, columns := m.Dimensions()
	copy := NewZeroMatrix[N](rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			copy[i][j] = m[i][j]
		}
	}
	return copy
}

// Float64Copy returns a copy of the provided matrix with all Number N converted to float64
func (m Matrix[N]) Float64Copy() Matrix[float64] {
	rows, columns := m.Dimensions()
	copy := NewZeroMatrix[float64](rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			copy[i][j] = float64(m[i][j])
		}
	}
	return copy
}

// SequentialTranspose returns the transpose of a matrix
func (m Matrix[N]) SequentialTranspose() Matrix[N] {
	x, y := m.Dimensions()
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
	mi, mj := m.Dimensions()
	ni, nj := n.Dimensions()
	return (mi == ni) && (mj == nj)
}

// WithinSigma returns true if m and n have the same dimensions and for each element
// sigma > | Mij - Nij |
func (m Matrix[N]) WithinSigma(n Matrix[N], sigma float64) bool {
	sameDimensions := m.SameDimensions(n)
	if !sameDimensions {
		return false
	}
	m64, n64 := m.Float64Copy(), n.Float64Copy()
	rows, columns := m.Dimensions()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			delta := math.Abs(m64[i][j] - n64[i][j])
			if sigma <= delta {
				return false
			}
		}
	}
	return true
}

// MultiplicationDimensions returns required dimensions of multiplication result and true if valid, false if not valid
func (m Matrix[N]) MultiplicationDimensions(n Matrix[N]) (int, int, bool) {
	a, b := m.Dimensions()
	c, d := n.Dimensions()
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

func IsFloat64[T any](x T) (ok bool) {
	_, ok = any(x).(float64)
	return ok
}
