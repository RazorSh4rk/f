package f

type Option[T any] struct {
	result T
	Err    error
}

func NewOpt[T any](result T, err error) Option[T] {
	return Option[T]{result, err}
}

func NewOptE[T any](err error) Option[T] {
	var empty T
	return Option[T]{empty, err}
}

func NewOptT[T any](result T) Option[T] {
	return Option[T]{result, nil}
}

func (o Option[T]) Ok() bool {
	return o.Err == nil
}

func (o Option[T]) Get() (T, error) {
	if !o.Ok() {
		var empty T
		return empty, o.Err
	}
	return o.result, nil
}

func (o Option[T]) GetOrElse(_else T) T {
	if !o.Ok() {
		return _else
	}
	return o.result
}
