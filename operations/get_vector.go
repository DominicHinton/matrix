package operations

func getZeroVector(length int) Vector {
	v := make([]int, length)
	return Vector(v)
}

// GetVectorFromColumn returns a vector at specified column if it exists in the matrix
func (m Matrix) GetVectorFromColumn(column int) (Vector, bool) {
	i, j := m.GetIJ()
	if column >= j {
		return nil, false
	}
	v := getZeroVector(i)
	for k := 0; k < i; k++ {
		v[k] = m[k][column]
	}
	return v, true
}

func (m Matrix) GetVectorFromRow(row int) (Vector, bool) {
	i, _ := m.GetIJ()
	if row >= i {
		return nil, false
	}
	v := Vector(m[row])
	return v, true
}
