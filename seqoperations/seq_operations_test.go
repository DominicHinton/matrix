package seqoperations_test

import (
	"errors"
	"testing"

	"github.com/DominicHinton/matrix/seqoperations"
	"github.com/stretchr/testify/assert"
)

var errRowColSuppliedOutBounds = errors.New("row or column number out of bounds")

func TestGetDimensions(t *testing.T) {
	m00 := seqoperations.Matrix[int]{}
	i0e, j0e := 0, 0
	i0a, j0a := m00.Dimensions()
	assert.Equal(t, i0e, i0a)
	assert.Equal(t, j0e, j0a)

	m50 := seqoperations.Matrix[int]{{}, {}, {}, {}, {}}
	i5e, j0e := 5, 0
	i5a, j0a := m50.Dimensions()
	assert.Equal(t, i5e, i5a)
	assert.Equal(t, j0e, j0a)

	m11 := seqoperations.Matrix[int]{{5}}
	i1e, j1e := 1, 1
	i1a, j1a := m11.Dimensions()
	assert.Equal(t, i1e, i1a)
	assert.Equal(t, j1e, j1a)

	m23 := seqoperations.Matrix[int]{{1, 2, 3}, {4, 5, 6}}
	i2e, j3e := 2, 3
	i2a, j3a := m23.Dimensions()
	assert.Equal(t, i2e, i2a)
	assert.Equal(t, j3e, j3a)

}

func TestIsSquareTrueCases(t *testing.T) {
	m00 := seqoperations.Matrix[int]{}
	assert.True(t, m00.IsSquare())

	m11 := seqoperations.Matrix[int]{{6}}
	assert.True(t, m11.IsSquare())

	m22 := seqoperations.Matrix[int]{{1, 2}, {3, 4}}
	assert.True(t, m22.IsSquare())

	m500500 := seqoperations.NewIdentityMatrix[int](500)
	assert.True(t, m500500.IsSquare())
}

func TestIsSquareFalseCases(t *testing.T) {
	m50 := seqoperations.Matrix[int]{{}, {}, {}, {}, {}}
	assert.False(t, m50.IsSquare())

	m12 := seqoperations.Matrix[int]{{1, 2}}
	assert.False(t, m12.IsSquare())

	m34 := seqoperations.Matrix[int]{{1, 1, 2, 2}, {34, 5, 6, 7}, {-1, 0, 7, 8}}
	assert.False(t, m34.IsSquare())
}

/*
Test SequentialMapToElements
*/

func TestSequentialMapToEmpty(t *testing.T) {
	m := seqoperations.Matrix[int]{}
	m.MapFunctionToElements(func(element int) int { return element + 1 })
	e := seqoperations.Matrix[int]{}
	assert.Equal(t, e, m)
}

func TestSequentialMapToOneByOne(t *testing.T) {
	m := seqoperations.Matrix[int]{{3}}
	a := m.MapFunctionToElements(func(element int) int { return element + 1 })
	e := seqoperations.Matrix[int]{{4}}
	assert.Equal(t, e, a)
}

func TestSequentialMapToTwoByThree(t *testing.T) {
	m := seqoperations.Matrix[int]{{-5, 101, 2}, {0, -1, 86}}
	a := m.MapFunctionToElements(func(element int) int { return element + 1 })
	e := seqoperations.Matrix[int]{{-4, 102, 3}, {1, 0, 87}}
	assert.Equal(t, e, a)
}

/*
Test MapFunctionToElementsInRowInPlace
*/

func TestMapFuncToElementsInRowInPlaceFourByThree(t *testing.T) {
	m := seqoperations.Matrix[int]{{-5, 100, 2}, {0, -1, 86}, {-5, 100, 2}, {0, -1, 86}}
	err := m.MapFunctionToElementsInRowInPlace(func(element int) int { return element + 1 }, 1)
	assert.Nil(t, err)
	e := seqoperations.Matrix[int]{{-5, 100, 2}, {1, 0, 87}, {-5, 100, 2}, {0, -1, 86}}
	assert.Equal(t, e, m)
}

func TestMapFuncToElementsInRowInPlaceFourByThreeRowOutOfBounds(t *testing.T) {
	m := seqoperations.Matrix[int]{{-5, 100, 2}, {0, -1, 86}, {-5, 100, 2}, {0, -1, 86}}
	err := m.MapFunctionToElementsInRowInPlace(func(element int) int { return element + 1 }, 4)
	assert.Equal(t, errRowColSuppliedOutBounds, err)
	e := seqoperations.Matrix[int]{{-5, 100, 2}, {0, -1, 86}, {-5, 100, 2}, {0, -1, 86}}
	assert.Equal(t, e, m)
	err = m.MapFunctionToElementsInRowInPlace(func(element int) int { return element + 1 }, -1)
	assert.Equal(t, errRowColSuppliedOutBounds, err)
	assert.Equal(t, e, m)
}
