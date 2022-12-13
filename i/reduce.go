package i

func Reduce[T, U any](
	in Iterable[T],
	init U,
	f func(U, T, int) U,
) U {
	out := init
	For(in, func(value T, i int) {
		out = f(out, value, i)
	})
	return out
}

func ReduceI[T, U any](
	in Iterator[T],
	init U,
	f func(U, T, int) U,
) U {
	out := init
	ForI(in, func(value T, i int) {
		out = f(out, value, i)
	})
	return out
}
