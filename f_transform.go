package f

import (
	"fmt"
)

var (
	ErrNotSlice = fmt.Errorf("not a slice")
)

func Map[FROM any, TO any](f F[FROM], fn func(FROM) TO) F[TO] {
	_r := make([]TO, len(f.Val))
	for idx, el := range f.Val {
		_r[idx] = fn(el)
	}
	return F[TO]{
		Val: _r,
	}
}

func Fold[FROM any, TO any](f F[FROM], r TO, fn func(TO, FROM) TO) TO {
	_r := r
	for _, el := range f.Val {
		_r = fn(_r, el)
	}
	return _r
}

func Flatten[T any](f F[[]T]) Option[F[T]] {
	_len := 0
	for _, arr := range f.Val {
		_len += len(arr)
	}
	_r := make([]T, _len)
	idx := 0
	for _, arr := range f.Val {
		for _, el := range arr {
			_r[idx] = el
			idx++
		}
	}
	return NewOpt(F[T]{
		Val: _r,
	}, nil)
}
