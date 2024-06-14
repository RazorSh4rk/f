package f

func (f F[T]) ForEach(fn func(T)) {
	for _, el := range f.Val {
		fn(el)
	}
}

func (f F[T]) ForAll(fn func(T) bool) bool {
	for _, el := range f.Val {
		if !fn(el) {
			return false
		}
	}
	return true
}

func (f F[T]) Has(fn func(T) bool) bool {
	for _, el := range f.Val {
		if fn(el) {
			return true
		}
	}
	return false
}

func (f F[T]) Find(fn func(T) bool) Option[T] {
	for _, el := range f.Val {
		if fn(el) {
			return NewOptT(el)
		}
	}
	return NewOptE[T](nil)
}
