package util

// region Iterable

// Accumulate is any function that returns a single value given two values of the same type.
type Accumulate[T any] func(acc, next T) T

// Predicate is a function that implements a conditional check on a single value and returns the boolean result.
type Predicate[T any] func(t T) bool

// Iterable is a collection whose elements can be iterated over.
type Iterable[T any] interface {

	// All returns true if all elements match the given predicate.
	All(pred Predicate[T]) bool

	// Any returns true if at least one element matches the given predicate.
	Any(pred Predicate[T]) bool

	// Filter returns a new list containing all elements of this Iterable that match a given Predicate.
	Filter(pred Predicate[T]) Iterable[T]

	// Fold accumulates value starting w/ initial and applying operation from left to right
	Fold(initial T, operation Accumulate[T]) T

	// Reduce accumulates values applying an operation from left to right
	Reduce(operation Accumulate[T]) T
}

// endregion

// region Pure Functions (Generic)

// Transform is any function that determines a return value given any input of a given type
type Transform[T, R any] func(in T) (out R)

// Map returns a List whose elements are the result of applying a Transform function to each element of a given List.
func Map[T, R any](l List[T], transform Transform[T, R]) List[R] {
	out := make([]R, len(l))
	for i, x := range l {
		out[i] = transform(x)
	}

	return out
}

type KeyValueSelector[T, K, V any] func(in T) (key K, value V)

// Associate returns a Map containing key/value pairs provided by a KeyValueTransform applied to each element of a List.
func Associate[T, K, V any](l List[T], transform KeyValueSelector[T, K, V]) {

}

// endregion

// region List

// List implements the Iterable interface using slices.
type List[T any] []T

// NewList is a convenience function to create a List using trailing arguments
// so that a List can be initialized like this:
//
//	l := NewList(1,2,3,4)
//
// instead of like this:
//
//	l := List[int]([]int{1,2,3,4})
//
// TODO: find out if this is idiomatic or if there's another way to shorten initialization
func NewList[T any](e ...T) List[T] { return e }

func (l List[T]) All(pred Predicate[T]) bool {
	// TODO: short circuit in a for
	return Map(l, func(in T) bool {
		return pred(in)
	}).Reduce(func(x, y bool) bool {
		return x && y
	})
}

func (l List[T]) Any(pred Predicate[T]) bool {
	return false
}

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

func (l List[T]) Filter(pred Predicate[T]) Iterable[T] {
	var f List[T]

	for _, x := range l {
		if pred(x) {
			f = append(f, x)
		}
	}

	return f
}

// endregion
