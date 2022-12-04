package i

func Chunk[T any](in Iterable[T], size int) Iterable[[]T] {
	return chunk[T]{in, size}
}

type chunk[T any] struct {
	in   Iterable[T]
	size int
}

func (c chunk[T]) Iterator() Iterator[[]T] {
	return chunkIter[T]{c.in.Iterator(), c.size}
}

type chunkIter[T any] struct {
	in   Iterator[T]
	size int
}

func (c chunkIter[T]) Next() ([]T, bool) {
	next := make([]T, 0, c.size)
	for i := 0; i < c.size; i++ {
		value, done := c.in.Next()
		if done {
			break
		}
		next = append(next, value)
	}
	if len(next) == 0 {
		return nil, true
	}
	return next, false
}
