package i

func For[T any](in Iterable[T], f func(T, int)) {
	ForI(in.Iterator(), f)
}

func ForI[T any](iter Iterator[T], f func(T, int)) {
	for i := 0; ; i++ {
		value, done := iter.Next()
		if done {
			break
		}
		f(value, i)
	}
}
