package util

type Accumulate[T any] func(acc, next T) T
type Iterable[T any] interface {
	// Fold Accumulates value starting w/ initial and applying operation from left to right
	Fold(initial T, operation Accumulate[T]) T

	// Reduce Accumulates value applying operation from left to right
	Reduce(operation Accumulate[T]) T
}

type List[T any] []T

func (l List[T]) Fold(initial T, operation Accumulate[T]) T {
	if len(l) == 0 {
		return initial
	} else {
		return l[1:].Fold(operation(initial, l[0]), operation)
	}
}
func (l List[T]) Reduce(operation Accumulate[T]) T {
	if len(l) == 0 {
		var zero T
		return zero
	} else {
		return l[1:].Fold(l[0], operation)
	}
}
