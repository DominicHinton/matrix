package operations

// DotProduct returns dot product and true if vectors are same length, 0 and false otherwise
func (v Vector) DotProduct(u Vector) (int, bool) {
	length := len(v)
	total := 0
	if length != len(u) {
		return total, false
	}
	for i := 0; i < length; i++ {
		total += v[i] * u[i]
	}
	return total, true
}
