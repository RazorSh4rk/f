package f

func Map[FROM any, TO any](f F[FROM], fn func(FROM) TO) F[TO] {
	_mapped := make([]TO, len(f.Val))
	for idx, el := range f.Val {
		_mapped[idx] = fn(el)
	}
	return F[TO]{
		Val: _mapped,
	}
}

func Fold[FROM any, TO any](f F[FROM], r TO, fn func(TO, FROM) TO) TO {
	_r := r
	for _, el := range f.Val {
		_r = fn(_r, el)
	}
	return _r
}
