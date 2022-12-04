package i

func Filter[T any](in Iterable[T], f func(T) bool) Iterable[T] {
	return filter[T]{in, f}
}

type filter[T any] struct {
	in Iterable[T]
	f  func(T) bool
}

func (f filter[T]) Iterator() Iterator[T] {
	return filterIter[T]{f.in.Iterator(), f.f}
}

type filterIter[T any] struct {
	in Iterator[T]
	f  func(T) bool
}

func (f filterIter[T]) Next() (T, bool) {
	for {
		if value, done := f.in.Next(); done {
			return value, true
		} else if f.f(value) {
			return value, false
		}
	}
}
