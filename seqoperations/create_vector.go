package seqoperations

func NewZeroVector[N Number](length int) Vector[N] {
	v := make(Vector[N], length)
	return v
}

// VectorFromColumn returns a vector at specified column if it exists in the matrix
func (m Matrix[N]) VectorFromColumn(column int) (Vector[N], bool) {
	i, j := m.Dimensions()
	if column >= j {
		return Vector[N]{}, false
	}
	v := NewZeroVector[N](i)
	for k := 0; k < i; k++ {
		v[k] = m[k][column]
	}
	return v, true
}

// getVectorFromRow returns vector at specified row if it exists in the matrix
func (m Matrix[N]) VectorFromRow(row int) (Vector[N], bool) {
	i, _ := m.Dimensions()
	if row >= i {
		return Vector[N]{}, false
	}
	return Vector[N](m[row]), true
}
