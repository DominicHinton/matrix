package seqoperations_test

import (
	"testing"

	"github.com/DominicHinton/matrix/seqoperations"
	"github.com/stretchr/testify/assert"
)

/*
ZeroMatrix Instantiation Tests
*/

// Tests for instantiation of unlikely dimensions using int type only
func TestNewZeroMatrixNilSizeInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](0, 0)
	expected := seqoperations.Matrix[int]{}
	assert.Equal(t, expected, m)
}

func TestNewZeroMatrixNegativeSizeInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](-1, -1)
	expected := seqoperations.Matrix[int]{}
	assert.Equal(t, expected, m)
}

func TestNewZeroMatrixNilSizeRowsInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](0, 7)
	expected := seqoperations.Matrix[int]{}
	assert.Equal(t, expected, m)
}

func TestNewZeroMatrixNilSizeColumnsInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](3, 0)
	expected := seqoperations.Matrix[int]{{}, {}, {}}
	assert.Equal(t, expected, m)
}

// Tests for instantiation of ZeroMatrix[int]
func TestNewZeroMatrixOneByOneInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](1, 1)
	z := 0
	expected := seqoperations.Matrix[int]{{z}}
	assert.Equal(t, expected, m)
}

func TestNewZeroMatrixTwoByThreeInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](2, 3)
	z := 0
	expected := seqoperations.Matrix[int]{{z, z, z}, {z, z, z}}
	assert.Equal(t, expected, m)
}

func TestNewZeroMatrixFiveByFourInt(t *testing.T) {
	m := seqoperations.NewZeroMatrix[int](5, 4)
	z := 0
	expected := seqoperations.Matrix[int]{{z, z, z, z}, {z, z, z, z}, {z, z, z, z}, {z, z, z, z}, {z, z, z, z}}
	assert.Equal(t, expected, m)
}

// Tests for instantiation of ZeroMatrix[Number]
func TestNewZeroMatrixOneByOneIntegers(t *testing.T) {

	m8 := seqoperations.NewZeroMatrix[int8](1, 1)
	z8 := int8(0)
	e8 := seqoperations.Matrix[int8]{{z8}}
	assert.Equal(t, e8, m8)

	m16 := seqoperations.NewZeroMatrix[int16](1, 1)
	z16 := int16(0)
	e16 := seqoperations.Matrix[int16]{{z16}}
	assert.Equal(t, e16, m16)

	m32 := seqoperations.NewZeroMatrix[int32](1, 1)
	z32 := int32(0)
	e32 := seqoperations.Matrix[int32]{{z32}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.NewZeroMatrix[int64](1, 1)
	z64 := int64(0)
	e64 := seqoperations.Matrix[int64]{{z64}}
	assert.Equal(t, e64, m64)
}

func TestNewZeroMatrixOneByOneUnsignedIntegers(t *testing.T) {

	mu := seqoperations.NewZeroMatrix[uint](1, 1)
	zu := uint(0)
	eu := seqoperations.Matrix[uint]{{zu}}
	assert.Equal(t, eu, mu)

	m8 := seqoperations.NewZeroMatrix[uint8](1, 1)
	z8 := uint8(0)
	e8 := seqoperations.Matrix[uint8]{{z8}}
	assert.Equal(t, e8, m8)

	m16 := seqoperations.NewZeroMatrix[uint16](1, 1)
	z16 := uint16(0)
	e16 := seqoperations.Matrix[uint16]{{z16}}
	assert.Equal(t, e16, m16)

	m32 := seqoperations.NewZeroMatrix[uint32](1, 1)
	z32 := uint32(0)
	e32 := seqoperations.Matrix[uint32]{{z32}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.NewZeroMatrix[uint64](1, 1)
	z64 := uint64(0)
	e64 := seqoperations.Matrix[uint64]{{z64}}
	assert.Equal(t, e64, m64)
}

func TestNewZeroMatrixOneByOneFloats(t *testing.T) {

	m32 := seqoperations.NewZeroMatrix[float32](1, 1)
	z32 := float32(0)
	e32 := seqoperations.Matrix[float32]{{z32}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.NewZeroMatrix[float64](1, 1)
	z64 := float64(0)
	e64 := seqoperations.Matrix[float64]{{z64}}
	assert.Equal(t, e64, m64)
}

func TestNewZeroMatrixTwoByThreeIntegers(t *testing.T) {

	m8 := seqoperations.NewZeroMatrix[int8](2, 3)
	z8 := int8(0)
	e8 := seqoperations.Matrix[int8]{{z8, z8, z8}, {z8, z8, z8}}
	assert.Equal(t, e8, m8)

	m16 := seqoperations.NewZeroMatrix[int16](2, 3)
	z16 := int16(0)
	e16 := seqoperations.Matrix[int16]{{z16, z16, z16}, {z16, z16, z16}}
	assert.Equal(t, e16, m16)

	m32 := seqoperations.NewZeroMatrix[int32](2, 3)
	z32 := int32(0)
	e32 := seqoperations.Matrix[int32]{{z32, z32, z32}, {z32, z32, z32}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.NewZeroMatrix[int64](2, 3)
	z64 := int64(0)
	e64 := seqoperations.Matrix[int64]{{z64, z64, z64}, {z64, z64, z64}}
	assert.Equal(t, e64, m64)
}

func TestNewZeroMatrixTwoByThreeUnsignedIntegers(t *testing.T) {

	mu := seqoperations.NewZeroMatrix[uint](2, 3)
	zu := uint(0)
	eu := seqoperations.Matrix[uint]{{zu, zu, zu}, {zu, zu, zu}}
	assert.Equal(t, eu, mu)

	m8 := seqoperations.NewZeroMatrix[uint8](2, 3)
	z8 := uint8(0)
	e8 := seqoperations.Matrix[uint8]{{z8, z8, z8}, {z8, z8, z8}}
	assert.Equal(t, e8, m8)

	m16 := seqoperations.NewZeroMatrix[uint16](2, 3)
	z16 := uint16(0)
	e16 := seqoperations.Matrix[uint16]{{z16, z16, z16}, {z16, z16, z16}}
	assert.Equal(t, e16, m16)

	m32 := seqoperations.NewZeroMatrix[uint32](2, 3)
	z32 := uint32(0)
	e32 := seqoperations.Matrix[uint32]{{z32, z32, z32}, {z32, z32, z32}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.NewZeroMatrix[uint64](2, 3)
	z64 := uint64(0)
	e64 := seqoperations.Matrix[uint64]{{z64, z64, z64}, {z64, z64, z64}}
	assert.Equal(t, e64, m64)
}

func TestNewZeroMatrixTwoByThreeFloats(t *testing.T) {

	m32 := seqoperations.NewZeroMatrix[float32](2, 3)
	z32 := float32(0)
	e32 := seqoperations.Matrix[float32]{{z32, z32, z32}, {z32, z32, z32}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.NewZeroMatrix[float64](2, 3)
	z64 := float64(0)
	e64 := seqoperations.Matrix[float64]{{z64, z64, z64}, {z64, z64, z64}}
	assert.Equal(t, e64, m64)
}

/*
NewMatrixFromSlice Instantiation Tests
*/

// Test that all types instantiate a valid 1 x 1 matrix
func TestNewMatrixFromSliceOneByOne(t *testing.T) {
	mInt := seqoperations.NewMatrixFromSlice(1, 1, []int{3})
	eInt := seqoperations.Matrix[int]{{3}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewMatrixFromSlice(1, 1, []int8{3})
	e8Int := seqoperations.Matrix[int8]{{3}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewMatrixFromSlice(1, 1, []int16{3})
	e16Int := seqoperations.Matrix[int16]{{3}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewMatrixFromSlice(1, 1, []int32{3})
	e32Int := seqoperations.Matrix[int32]{{3}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewMatrixFromSlice(1, 1, []int64{3})
	e64Int := seqoperations.Matrix[int64]{{3}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewMatrixFromSlice(1, 1, []uint{3})
	eui := seqoperations.Matrix[uint]{{3}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewMatrixFromSlice(1, 1, []uint8{3})
	e8UInt := seqoperations.Matrix[uint8]{{3}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewMatrixFromSlice(1, 1, []uint16{3})
	e16UInt := seqoperations.Matrix[uint16]{{3}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewMatrixFromSlice(1, 1, []uint32{3})
	e32UInt := seqoperations.Matrix[uint32]{{3}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewMatrixFromSlice(1, 1, []uint64{3})
	e64UInt := seqoperations.Matrix[uint64]{{3}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewMatrixFromSlice(1, 1, []float32{3})
	e32Float := seqoperations.Matrix[float32]{{3}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewMatrixFromSlice(1, 1, []float64{3})
	e64Float := seqoperations.Matrix[float64]{{3}}
	assert.Equal(t, e64Float, m64Float)
}

// Test a 1x1 matrix with empty slice.
func TestNewMatrixFromSliceOneByOneEmpty(t *testing.T) {
	mInt := seqoperations.NewMatrixFromSlice(1, 1, []int{})
	eInt := seqoperations.Matrix[int]{}
	assert.Equal(t, eInt, mInt)
}

// Test that a 2 x 3 matrix with six values passed will instantiate for all types.
func TestNewMatrixFromSliceTwoByThree(t *testing.T) {
	mInt := seqoperations.NewMatrixFromSlice(2, 3, []int{1, 2, 3, 4, 5, 6})
	eInt := seqoperations.Matrix[int]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewMatrixFromSlice(2, 3, []int8{1, 2, 3, 4, 5, 6})
	e8Int := seqoperations.Matrix[int8]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewMatrixFromSlice(2, 3, []int16{1, 2, 3, 4, 5, 6})
	e16Int := seqoperations.Matrix[int16]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewMatrixFromSlice(2, 3, []int32{1, 2, 3, 4, 5, 6})
	e32Int := seqoperations.Matrix[int32]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewMatrixFromSlice(2, 3, []int64{1, 2, 3, 4, 5, 6})
	e64Int := seqoperations.Matrix[int64]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewMatrixFromSlice(2, 3, []uint{1, 2, 3, 4, 5, 6})
	eui := seqoperations.Matrix[uint]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint8{1, 2, 3, 4, 5, 6})
	e8UInt := seqoperations.Matrix[uint8]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint16{1, 2, 3, 4, 5, 6})
	e16UInt := seqoperations.Matrix[uint16]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint32{1, 2, 3, 4, 5, 6})
	e32UInt := seqoperations.Matrix[uint32]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint64{1, 2, 3, 4, 5, 6})
	e64UInt := seqoperations.Matrix[uint64]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewMatrixFromSlice(2, 3, []float32{1, 2, 3, 4, 5, 6})
	e32Float := seqoperations.Matrix[float32]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewMatrixFromSlice(2, 3, []float64{1, 2, 3, 4, 5, 6})
	e64Float := seqoperations.Matrix[float64]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e64Float, m64Float)
}

// Test that a 2 x 3 matrix with four values passed will instantiate for all types.
func TestNewMatrixFromSliceTwoByThreeUnderflow(t *testing.T) {
	mInt := seqoperations.NewMatrixFromSlice(2, 3, []int{1, 2, 3, 4})
	eInt := seqoperations.Matrix[int]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewMatrixFromSlice(2, 3, []int8{1, 2, 3, 4})
	e8Int := seqoperations.Matrix[int8]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewMatrixFromSlice(2, 3, []int16{1, 2, 3, 4})
	e16Int := seqoperations.Matrix[int16]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewMatrixFromSlice(2, 3, []int32{1, 2, 3, 4})
	e32Int := seqoperations.Matrix[int32]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewMatrixFromSlice(2, 3, []int64{1, 2, 3, 4})
	e64Int := seqoperations.Matrix[int64]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewMatrixFromSlice(2, 3, []uint{1, 2, 3, 4})
	eui := seqoperations.Matrix[uint]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint8{1, 2, 3, 4})
	e8UInt := seqoperations.Matrix[uint8]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint16{1, 2, 3, 4})
	e16UInt := seqoperations.Matrix[uint16]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint32{1, 2, 3, 4})
	e32UInt := seqoperations.Matrix[uint32]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint64{1, 2, 3, 4})
	e64UInt := seqoperations.Matrix[uint64]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewMatrixFromSlice(2, 3, []float32{1, 2, 3, 4})
	e32Float := seqoperations.Matrix[float32]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewMatrixFromSlice(2, 3, []float64{1, 2, 3, 4})
	e64Float := seqoperations.Matrix[float64]{{1, 2, 3}, {4, 1, 2}}
	assert.Equal(t, e64Float, m64Float)
}

// Test that a 2 x 3 matrix with nine values passed will instantiate for all types.
func TestNewMatrixFromSliceTwoByThreeOverflow(t *testing.T) {
	mInt := seqoperations.NewMatrixFromSlice(2, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	eInt := seqoperations.Matrix[int]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewMatrixFromSlice(2, 3, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e8Int := seqoperations.Matrix[int8]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewMatrixFromSlice(2, 3, []int16{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e16Int := seqoperations.Matrix[int16]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewMatrixFromSlice(2, 3, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e32Int := seqoperations.Matrix[int32]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewMatrixFromSlice(2, 3, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e64Int := seqoperations.Matrix[int64]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewMatrixFromSlice(2, 3, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9})
	eui := seqoperations.Matrix[uint]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e8UInt := seqoperations.Matrix[uint8]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e16UInt := seqoperations.Matrix[uint16]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e32UInt := seqoperations.Matrix[uint32]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewMatrixFromSlice(2, 3, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e64UInt := seqoperations.Matrix[uint64]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewMatrixFromSlice(2, 3, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e32Float := seqoperations.Matrix[float32]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewMatrixFromSlice(2, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	e64Float := seqoperations.Matrix[float64]{{1, 2, 3}, {4, 5, 6}}
	assert.Equal(t, e64Float, m64Float)
}

/*
NewIdentityMatrix Tests
*/

// Instantiate a 1 x 1 identity matrix for all types
func TestIdentityOrderOne(t *testing.T) {
	mInt := seqoperations.NewIdentityMatrix[int](1)
	eInt := seqoperations.Matrix[int]{{1}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewIdentityMatrix[int8](1)
	e8Int := seqoperations.Matrix[int8]{{1}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewIdentityMatrix[int16](1)
	e16Int := seqoperations.Matrix[int16]{{1}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewIdentityMatrix[int32](1)
	e32Int := seqoperations.Matrix[int32]{{1}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewIdentityMatrix[int64](1)
	e64Int := seqoperations.Matrix[int64]{{1}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewIdentityMatrix[uint](1)
	eui := seqoperations.Matrix[uint]{{1}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewIdentityMatrix[uint8](1)
	e8UInt := seqoperations.Matrix[uint8]{{1}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewIdentityMatrix[uint16](1)
	e16UInt := seqoperations.Matrix[uint16]{{1}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewIdentityMatrix[uint32](1)
	e32UInt := seqoperations.Matrix[uint32]{{1}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewIdentityMatrix[uint64](1)
	e64UInt := seqoperations.Matrix[uint64]{{1}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewIdentityMatrix[float32](1)
	e32Float := seqoperations.Matrix[float32]{{1}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewIdentityMatrix[float64](1)
	e64Float := seqoperations.Matrix[float64]{{1}}
	assert.Equal(t, e64Float, m64Float)
}

// Instantiate a 2 x 2 identity matrix for all types
func TestIdentityOrderTwo(t *testing.T) {
	mInt := seqoperations.NewIdentityMatrix[int](2)
	eInt := seqoperations.Matrix[int]{{1, 0}, {0, 1}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewIdentityMatrix[int8](2)
	e8Int := seqoperations.Matrix[int8]{{1, 0}, {0, 1}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewIdentityMatrix[int16](2)
	e16Int := seqoperations.Matrix[int16]{{1, 0}, {0, 1}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewIdentityMatrix[int32](2)
	e32Int := seqoperations.Matrix[int32]{{1, 0}, {0, 1}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewIdentityMatrix[int64](2)
	e64Int := seqoperations.Matrix[int64]{{1, 0}, {0, 1}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewIdentityMatrix[uint](2)
	eui := seqoperations.Matrix[uint]{{1, 0}, {0, 1}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewIdentityMatrix[uint8](2)
	e8UInt := seqoperations.Matrix[uint8]{{1, 0}, {0, 1}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewIdentityMatrix[uint16](2)
	e16UInt := seqoperations.Matrix[uint16]{{1, 0}, {0, 1}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewIdentityMatrix[uint32](2)
	e32UInt := seqoperations.Matrix[uint32]{{1, 0}, {0, 1}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewIdentityMatrix[uint64](2)
	e64UInt := seqoperations.Matrix[uint64]{{1, 0}, {0, 1}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewIdentityMatrix[float32](2)
	e32Float := seqoperations.Matrix[float32]{{1, 0}, {0, 1}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewIdentityMatrix[float64](2)
	e64Float := seqoperations.Matrix[float64]{{1, 0}, {0, 1}}
	assert.Equal(t, e64Float, m64Float)
}

// Instantiate a 3 x 3 identity matrix for all types
func TestIdentityOrderThree(t *testing.T) {
	mInt := seqoperations.NewIdentityMatrix[int](3)
	eInt := seqoperations.Matrix[int]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, eInt, mInt)

	m8Int := seqoperations.NewIdentityMatrix[int8](3)
	e8Int := seqoperations.Matrix[int8]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e8Int, m8Int)

	m16Int := seqoperations.NewIdentityMatrix[int16](3)
	e16Int := seqoperations.Matrix[int16]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e16Int, m16Int)

	m32Int := seqoperations.NewIdentityMatrix[int32](3)
	e32Int := seqoperations.Matrix[int32]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e32Int, m32Int)

	m64Int := seqoperations.NewIdentityMatrix[int64](3)
	e64Int := seqoperations.Matrix[int64]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e64Int, m64Int)

	mui := seqoperations.NewIdentityMatrix[uint](3)
	eui := seqoperations.Matrix[uint]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, eui, mui)

	m8UInt := seqoperations.NewIdentityMatrix[uint8](3)
	e8UInt := seqoperations.Matrix[uint8]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e8UInt, m8UInt)

	m16UInt := seqoperations.NewIdentityMatrix[uint16](3)
	e16UInt := seqoperations.Matrix[uint16]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e16UInt, m16UInt)

	m32UInt := seqoperations.NewIdentityMatrix[uint32](3)
	e32UInt := seqoperations.Matrix[uint32]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e32UInt, m32UInt)

	m64UInt := seqoperations.NewIdentityMatrix[uint64](3)
	e64UInt := seqoperations.Matrix[uint64]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e64UInt, m64UInt)

	m32Float := seqoperations.NewIdentityMatrix[float32](3)
	e32Float := seqoperations.Matrix[float32]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e32Float, m32Float)

	m64Float := seqoperations.NewIdentityMatrix[float64](3)
	e64Float := seqoperations.Matrix[float64]{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	assert.Equal(t, e64Float, m64Float)
}

// Test that an empty matrix is returned when dimension is less than one.
func TestInvalidDimenionIdentityMatrix(t *testing.T) {
	mZero := seqoperations.NewIdentityMatrix[int](0)
	eZero := seqoperations.Matrix[int]{}
	assert.Equal(t, eZero, mZero)

	mNegative := seqoperations.NewIdentityMatrix[int](0)
	eNegative := seqoperations.Matrix[int]{}
	assert.Equal(t, eNegative, mNegative)
}

/*
Fill Matrix Tests
*/

// Test that an empty matrix fills to an empty matrix
func TestFillEmptyMatrix(t *testing.T) {
	mi := seqoperations.Matrix[int]{}
	mi.FillMatrix(3)
	ei := seqoperations.Matrix[int]{}
	assert.Equal(t, ei, mi)

	m8 := seqoperations.Matrix[int8]{}
	m8.FillMatrix(3)
	e8 := seqoperations.Matrix[int8]{}
	assert.Equal(t, e8, m8)

	m16 := seqoperations.Matrix[int16]{}
	m16.FillMatrix(3)
	e16 := seqoperations.Matrix[int16]{}
	assert.Equal(t, e16, m16)

	m32 := seqoperations.Matrix[int32]{}
	m32.FillMatrix(3)
	e32 := seqoperations.Matrix[int32]{}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.Matrix[int64]{}
	m64.FillMatrix(3)
	e64 := seqoperations.Matrix[int64]{}
	assert.Equal(t, e64, m64)

	mu := seqoperations.Matrix[uint]{}
	mu.FillMatrix(3)
	eu := seqoperations.Matrix[uint]{}
	assert.Equal(t, eu, mu)

	mui8 := seqoperations.Matrix[uint8]{}
	mui8.FillMatrix(3)
	eui8 := seqoperations.Matrix[uint8]{}
	assert.Equal(t, eui8, mui8)

	mui16 := seqoperations.Matrix[uint16]{}
	mui16.FillMatrix(3)
	eui16 := seqoperations.Matrix[uint16]{}
	assert.Equal(t, eui16, mui16)

	mui32 := seqoperations.Matrix[uint32]{}
	mui32.FillMatrix(3)
	eui32 := seqoperations.Matrix[uint32]{}
	assert.Equal(t, eui32, mui32)

	mui64 := seqoperations.Matrix[uint64]{}
	mui64.FillMatrix(3)
	eui64 := seqoperations.Matrix[uint64]{}
	assert.Equal(t, eui64, mui64)

	mf32 := seqoperations.Matrix[float32]{}
	mf32.FillMatrix(3)
	ef32 := seqoperations.Matrix[float32]{}
	assert.Equal(t, ef32, mf32)

	mf64 := seqoperations.Matrix[float64]{}
	mf64.FillMatrix(3)
	ef64 := seqoperations.Matrix[float64]{}
	assert.Equal(t, ef64, mf64)
}

func TestFillOneByOneMatrix(t *testing.T) {
	mi := seqoperations.Matrix[int]{{1}}
	mi.FillMatrix(3)
	ei := seqoperations.Matrix[int]{{3}}
	assert.Equal(t, ei, mi)

}

func TestFillTwoByThreeMatrix(t *testing.T) {
	mi := seqoperations.Matrix[int]{{1, 2, 3}, {4, 5, 6}}
	mi.FillMatrix(3)
	ei := seqoperations.Matrix[int]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, ei, mi)

	m8 := seqoperations.Matrix[int8]{{1, 2, 3}, {4, 5, 6}}
	m8.FillMatrix(3)
	e8 := seqoperations.Matrix[int8]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, e8, m8)

	m16 := seqoperations.Matrix[int16]{{1, 2, 3}, {4, 5, 6}}
	m16.FillMatrix(3)
	e16 := seqoperations.Matrix[int16]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, e16, m16)

	m32 := seqoperations.Matrix[int32]{{1, 2, 3}, {4, 5, 6}}
	m32.FillMatrix(3)
	e32 := seqoperations.Matrix[int32]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, e32, m32)

	m64 := seqoperations.Matrix[int64]{{1, 2, 3}, {4, 5, 6}}
	m64.FillMatrix(3)
	e64 := seqoperations.Matrix[int64]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, e64, m64)

	mu := seqoperations.Matrix[uint]{{1, 2, 3}, {4, 5, 6}}
	mu.FillMatrix(3)
	eu := seqoperations.Matrix[uint]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, eu, mu)

	mui8 := seqoperations.Matrix[uint8]{{1, 2, 3}, {4, 5, 6}}
	mui8.FillMatrix(3)
	eui8 := seqoperations.Matrix[uint8]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, eui8, mui8)

	mui16 := seqoperations.Matrix[uint16]{{1, 2, 3}, {4, 5, 6}}
	mui16.FillMatrix(3)
	eui16 := seqoperations.Matrix[uint16]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, eui16, mui16)

	mui32 := seqoperations.Matrix[uint32]{{1, 2, 3}, {4, 5, 6}}
	mui32.FillMatrix(3)
	eui32 := seqoperations.Matrix[uint32]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, eui32, mui32)

	mui64 := seqoperations.Matrix[uint64]{{1, 2, 3}, {4, 5, 6}}
	mui64.FillMatrix(3)
	eui64 := seqoperations.Matrix[uint64]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, eui64, mui64)

	mf32 := seqoperations.Matrix[float32]{{1, 2, 3}, {4, 5, 6}}
	mf32.FillMatrix(3)
	ef32 := seqoperations.Matrix[float32]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, ef32, mf32)

	mf64 := seqoperations.Matrix[float64]{{1, 2, 3}, {4, 5, 6}}
	mf64.FillMatrix(3)
	ef64 := seqoperations.Matrix[float64]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, ef64, mf64)
}

/*
NewConstantMatrix Tests
*/

func TestEmptyNewConstantMatrix(t *testing.T) {
	m := seqoperations.NewConstantMatrix(0, 0, 3)
	e := seqoperations.Matrix[int]{}
	assert.Equal(t, e, m)
}

func TestOneByOneNewConstantMatrix(t *testing.T) {
	m := seqoperations.NewConstantMatrix(1, 1, 3)
	e := seqoperations.Matrix[int]{{3}}
	assert.Equal(t, e, m)
}

func TestTwoByThreeNewConstantMatrix(t *testing.T) {
	m := seqoperations.NewConstantMatrix(2, 3, 3)
	e := seqoperations.Matrix[int]{{3, 3, 3}, {3, 3, 3}}
	assert.Equal(t, e, m)
}
