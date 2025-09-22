package stablemap

type Maybe[T any] struct {
	Val T
	Is  bool
	Err error
}
