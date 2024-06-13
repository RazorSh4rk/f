package f

import (
	"fmt"
	"reflect"
)

type F[T any] struct {
	Val []T
}

func (f F[T]) Is(other F[T]) bool {
	if len(f.Val) != len(other.Val) {
		return false
	}
	for i := 0; i < len(f.Val); i++ {
		if !reflect.DeepEqual(f.Val[i], other.Val[i]) {
			return false
		}
	}
	return true
}

func (f F[T]) String() string {
	return fmt.Sprint(f.Val)
}

func From[T any](v []T) F[T] {
	return F[T]{
		Val: v,
	}
}

func Gen[T any](fn func(int) T, count int) F[T] {
	_vals := make([]T, count)
	for i := 0; i < count; i++ {
		_vals[i] = fn(i)
	}
	return F[T]{
		Val: _vals,
	}
}
