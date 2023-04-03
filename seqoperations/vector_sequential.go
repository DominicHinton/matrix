package seqoperations

// DotProduct returns dot product and true if vectors are same length, 0 and false otherwise
func (v Vector[N]) DotProduct(u Vector[N]) (N, bool) {
	length := len(v)
	var total N
	if length != len(u) {
		return total, false
	}
	for i := 0; i < length; i++ {
		total += v[i] * u[i]
	}
	return total, true
}

// SequentialMapFuncToElements takes fn: a function that returns a result of operation on one element of vector v,
// x: a second argument for fn and returns new matrix with same operation applied to every element
func (v Vector[N]) SequentialMapFuncToElements(fn ConstantSequentialOperater[N]) Vector[N] {
	length := len(v)
	output := NewZeroVector[N](length)
	for k := 0; k < length; k++ {
		output[k] = fn(v[k])
	}
	return output
}

// SequentialAddToElements adds x to every element in m sequentially and returns resulting matrix
func (v Vector[N]) SequentialAddToElements(x N) Vector[N] {
	return v.SequentialMapFuncToElements(func(element N) N { return element + x })
}

// SequentialSubtractFromElements subtracts x from every element in m sequentially and returns resulting matrix
func (v Vector[N]) SequentialSubtractFromElements(x N) Vector[N] {
	return v.SequentialMapFuncToElements(func(element N) N { return element - x })
}

// SequentialSubtractElementsFrom subtracts every element in m from x sequentially and returns resulting matrix
func (v Vector[N]) SequentialSubtractElementsFrom(x N) Vector[N] {
	return v.SequentialMapFuncToElements(func(element N) N { return x - element })
}

// SequentialMultiplyElementsBy multiplies x by every element in m sequentially and returns resulting matrix
func (v Vector[N]) SequentialMultiplyElementsBy(x N) Vector[N] {
	return v.SequentialMapFuncToElements(func(element N) N { return x * element })
}

// SequentialDivideElementsBy divides every element in m by x sequentially and returns resulting matrix
func (v Vector[N]) SequentialDivideElementsBy(x N) Vector[N] {
	return v.SequentialMapFuncToElements(func(element N) N { return element / x })
}

// SequentialDivideByElements divides x by each element in m and returns resulting matrix
func (v Vector[N]) SequentialDivideByElements(x N) Vector[N] {
	return v.SequentialMapFuncToElements(func(element N) N { return x / element })
}
