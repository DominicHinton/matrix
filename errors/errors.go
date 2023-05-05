package errors

import "errors"

var (
	ErrDifferentDimension      = errors.New("matrices must be of same dimension")
	ErrMultiplicationValidity  = errors.New("matrices of these dimensions cannot be multiplied in this order")
	ErrNonSquare               = errors.New("i and j values are not equal, this matrix should be square")
	ErrNoInverse               = errors.New("no inverse exists for this matrix")
	ErrNotFloat64              = errors.New("this method's assumption of float64 matrix input was not satisfied")
	ErrRowColSuppliedOutBounds = errors.New("row or column number out of bounds")
	ErrUnexpected              = errors.New("unexpected error occurred")
	ErrZeroLength              = errors.New("matrix has no rows")
)
