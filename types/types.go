package types

type Number interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8 | float64 | float32
}

type Float interface {
	float64 | float32
}

type Matrix[N Number] [][]N
type Vector[N Number] []N

type ConstantSequentialOperater[N Number] func(N) N
type OneToOneSequentialOperater[N Number] func(N, N) N
