package i

func Peeker[T any](in Iterator[T]) PeekIterator[T] {
	if p, ok := in.(PeekIterator[T]); ok {
		return p
	}
	return &peeker[T]{it: in}
}

type peeker[T any] struct {
	it     Iterator[T]
	pv     T
	peeked bool
}

func (p *peeker[T]) Next() (T, bool) {
	if p.peeked {
		p.peeked = false
		return p.pv, false
	}
	return p.it.Next()
}

func (p *peeker[T]) Peek() (T, bool) {
	if p.peeked {
		return p.pv, false
	}
	v, done := p.it.Next()
	if done {
		return v, true
	}
	p.pv, p.peeked = v, true
	return v, false
}
