package i

func For[T any](in Iterable[T], f func(T, int)) {
	iter := in.Iterator()
	for i := 0; ; i++ {
		value, done := iter.Next()
		if done {
			break
		}
		f(value, i)
	}
}
