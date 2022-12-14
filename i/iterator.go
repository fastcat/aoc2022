package i

type Iterator[T any] interface {
	// Next returns either a value and false, or zero(T) and true
	Next() (value T, done bool)
}

type Streamer[T any] interface {
	// Next returns either a value and nil, or zero(T) and an error. If the
	// iterator completed normally, err is Done.
	Next() (value T, err error)
	Close()
}

type errDone struct{}

func (errDone) Error() string { return "iterator done" }

// Done is returned from Streamer[T].Next() when it finishes streaming normally.
var Done errDone

type PeekIterator[T any] interface {
	Iterator[T]
	Peek() (value T, hasNext bool)
}
