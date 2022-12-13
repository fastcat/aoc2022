package i

type manyMapper[T, U any] struct {
	in Iterable[T]
	f  func(T, int) Iterable[U]
}

func Many[T, U any](
	in Iterable[T],
	f func(T, int) Iterable[U],
) Iterable[U] {
	return manyMapper[T, U]{in, f}
}
func (m manyMapper[T, U]) Iterator() Iterator[U] {
	return &manyMapIter[T, U]{m.in.Iterator(), nil, m.f, -1}
}

type manyMapIter[T, U any] struct {
	in  Iterator[T]
	out Iterator[U]
	f   func(T, int) Iterable[U]
	i   int
}

func (m *manyMapIter[T, U]) Next() (value U, done bool) {
	for {
		if m.out == nil {
			m.i++
			t, done := m.in.Next()
			if done {
				return value, true
			}
			m.out = m.f(t, m.i).Iterator()
			continue // re-check m.out
		}
		value, done = m.out.Next()
		if done {
			m.out = nil
			continue
		}
		return
	}
}
