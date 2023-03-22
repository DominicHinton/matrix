package operations

import "errors"

var (
	// errNonSquare              = errors.New("i and j values are not equal, this matrix should be square")
	errDifferentDimesion      = errors.New("matrices must be of same dimension")
	errMultiplicationValidity = errors.New("matrices of these dimensions cannot be multiplied in this order")
)
