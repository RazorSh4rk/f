package f

import "errors"

var (
	ErrEmptyList = errors.New("empty list")
)

func (f F[T]) Head() Option[T] {
	if len(f.Val) == 0 {
		return NewOptE[T](ErrEmptyList)
	}
	return NewOpt(f.Val[0], nil)
}

func (f F[T]) Last() Option[T] {
	if len(f.Val) == 0 {
		return NewOptE[T](ErrEmptyList)
	}
	return NewOpt(f.Val[len(f.Val)-1], nil)
}

func (f F[T]) Tail() F[T] {
	if len(f.Val) == 0 || len(f.Val) == 1 {
		return F[T]{}
	}
	_tail := make([]T, len(f.Val)-1)
	for i := 1; i < len(f.Val); i++ {
		_tail[i-1] = f.Val[i]
	}
	return F[T]{Val: _tail}
}

func (f F[T]) Reverse() F[T] {
	_rev := make([]T, len(f.Val))
	for i := len(f.Val) - 1; i >= 0; i-- {
		_rev[len(_rev)-1-i] = f.Val[i]
	}
	return F[T]{
		Val: _rev,
	}
}

func (f F[T]) TakeWhile(fn func(T) bool) F[T] {
	_vals := make([]T, 0)
	for _, el := range f.Val {
		if fn(el) {
			_vals = append(_vals, el)
		}
	}
	return F[T]{
		Val: _vals,
	}
}

func (f F[T]) DropWhile(fn func(T) bool) F[T] {
	_vals := make([]T, 0)
	for _, el := range f.Val {
		if !fn(el) {
			_vals = append(_vals, el)
		}
	}
	return F[T]{
		Val: _vals,
	}
}

func (f F[T]) ZipWith(other F[T]) F[T] {
	_vals := make([]T, 0, len(f.Val)+len(other.Val))
	minL := len(other.Val)
	if len(f.Val) < len(other.Val) {
		minL = len(f.Val)
	}
	for i := 0; i < minL; i++ {
		_vals = append(_vals, f.Val[i], other.Val[i])
	}
	if len(f.Val) > len(other.Val) {
		_vals = append(_vals, f.Val[minL:]...)
	} else {
		_vals = append(_vals, other.Val[minL:]...)
	}
	return F[T]{
		Val: _vals,
	}
}

func (f F[T]) Filter(fn func(T) bool) F[T] {
	_filtered := make([]T, 0)
	for _, el := range f.Val {
		if fn(el) {
			_filtered = append(_filtered, el)
		}
	}
	return F[T]{
		Val: _filtered,
	}
}
