package i

type funcer[T any] func() Iterator[T]

func Funcer[T any](f func() Iterator[T]) Iterable[T] {
	return funcer[T](f)
}

func (f funcer[T]) Iterator() Iterator[T] {
	return f()
}

type funcIter[T any] func() (T, bool)

func Func[T any](f func() (T, bool)) funcIter[T] {
	return funcIter[T](f)
}

func (f funcIter[T]) Iterator() Iterator[T] {
	return f
}

func (f funcIter[T]) Next() (T, bool) {
	return f()
}
