package concoperations

import "sync"

// DotProduct returns dot product and true if vectors are same length, 0 and false otherwise
// No concurrency implemented as this is not beleived to offer performance benefit here.
func (v Vector[N]) DotProduct(u Vector[N]) (N, bool) {
	length := len(v)
	var total N
	if length != len(u) {
		return total, false
	}
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {
		total += v[i] * u[i]
	}
	return total, true
}

// MapFunctionToElements takes fn: a function that returns a result of operation on one element of vector v,
// x: a second argument for fn and returns new matrix with same operation applied to every element
func (v Vector[N]) MapFunctionToElements(fn ConstantSequentialOperater[N]) Vector[N] {
	length := len(v)
	output := NewZeroVector[N](length)
	var wg sync.WaitGroup
	wg.Add(length)
	for i := 0; i < length; i++ {
		k := i
		go func() {
			defer wg.Done()
			output[k] = fn(v[k])
		}()
	}
	wg.Wait()
	return output
}

// AddToElements adds x to every element in m sequentially and returns resulting matrix
func (v Vector[N]) AddToElements(x N) Vector[N] {
	return v.MapFunctionToElements(func(element N) N { return element + x })
}

// SubtractFromElements subtracts x from every element in m sequentially and returns resulting matrix
func (v Vector[N]) SubtractFromElements(x N) Vector[N] {
	return v.MapFunctionToElements(func(element N) N { return element - x })
}

// SubtractElementsFrom subtracts every element in m from x sequentially and returns resulting matrix
func (v Vector[N]) SubtractElementsFrom(x N) Vector[N] {
	return v.MapFunctionToElements(func(element N) N { return x - element })
}

// MultiplyElementsBy multiplies x by every element in m sequentially and returns resulting matrix
func (v Vector[N]) MultiplyElementsBy(x N) Vector[N] {
	return v.MapFunctionToElements(func(element N) N { return x * element })
}

// DivideElementsBy divides every element in m by x sequentially and returns resulting matrix
func (v Vector[N]) DivideElementsBy(x N) Vector[N] {
	return v.MapFunctionToElements(func(element N) N { return element / x })
}

// DivideByElements divides x by each element in m and returns resulting matrix
func (v Vector[N]) DivideByElements(x N) Vector[N] {
	return v.MapFunctionToElements(func(element N) N { return x / element })
}
