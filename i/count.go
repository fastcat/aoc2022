package i

func Count[T any](in Iterable[T]) int {
	c := 0
	i := in.Iterator()
	for _, d := i.Next(); !d; _, d = i.Next() {
		c++
	}
	return c
}
