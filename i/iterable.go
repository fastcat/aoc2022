package i

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type Streamable[T any] interface {
	Stream() Streamer[T]
}
