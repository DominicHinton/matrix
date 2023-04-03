package seqoperations

func NewZeroVector[N Number](length int) Vector[N] {
	v := make(Vector[N], length)
	return v
}

// GetVectorFromColumn returns a vector at specified column if it exists in the matrix
func (m Matrix[N]) GetVectorFromColumn(column int) (Vector[N], bool) {
	i, j := m.GetDimensions()
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
func (m Matrix[N]) GetVectorFromRow(row int) (Vector[N], bool) {
	i, _ := m.GetDimensions()
	if row >= i {
		return Vector[N]{}, false
	}
	return Vector[N](m[row]), true
}
