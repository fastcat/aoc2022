package i

type mapper[T, U any] struct {
	in Iterable[T]
	f  func(T) U
}

func Map[T, U any](in Iterable[T], f func(T) U) Iterable[U] {
	return mapper[T, U]{in, f}
}
func (m mapper[T, U]) Iterator() Iterator[U] {
	return mapI[T, U]{m.in.Iterator(), m.f}
}

type mapI[T, U any] struct {
	in Iterator[T]
	f  func(T) U
}

func (m mapI[T, U]) Next() (value U, done bool) {
	var v T
	v, done = m.in.Next()
	if !done {
		value = m.f(v)
	}
	return
}
