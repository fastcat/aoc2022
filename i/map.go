package i

type mapper[T, U any] struct {
	in Iterable[T]
	f  func(T, int) U
}

func Map[T, U any](in Iterable[T], f func(T, int) U) Iterable[U] {
	return mapper[T, U]{in, f}
}
func (m mapper[T, U]) Iterator() Iterator[U] {
	return &mapIter[T, U]{m.in.Iterator(), m.f, 0}
}

type mapIter[T, U any] struct {
	in Iterator[T]
	f  func(T, int) U
	i  int
}

func (m *mapIter[T, U]) Next() (value U, done bool) {
	var v T
	v, done = m.in.Next()
	if !done {
		value = m.f(v, m.i)
		m.i++
	}
	return
}

func NoIndex[T, U any](f func(T) U) func(T, int) U {
	return func(t T, i int) U {
		return f(t)
	}
}
