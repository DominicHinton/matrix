package seqoperations

import (
	e "github.com/DominicHinton/matrix/errors"
	"github.com/DominicHinton/matrix/types"
)

type Number = types.Number
type Float = types.Float
type Matrix[N Number] types.Matrix[N]
type Vector[N Number] types.Vector[N]
type ConstantSequentialOperater[N Number] types.ConstantSequentialOperater[N]
type OneToOneSequentialOperater[N Number] types.OneToOneSequentialOperater[N]

var (
	errDifferentDimension      = e.ErrDifferentDimension
	errMultiplicationValidity  = e.ErrMultiplicationValidity
	errNonSquare               = e.ErrNonSquare
	errNoInverse               = e.ErrNoInverse
	errNotFloat64              = e.ErrNotFloat64
	errRowColSuppliedOutBounds = e.ErrRowColSuppliedOutBounds
	errUnexpected              = e.ErrUnexpected
	errZeroLength              = e.ErrZeroLength
)
