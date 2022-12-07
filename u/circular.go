package u

import "errors"

type Circular[T any] struct {
	s         []T
	head, len int
}

func NewCircular[T any](cap int) *Circular[T] {
	return &Circular[T]{
		s: make([]T, cap),
	}
}

func (c *Circular[T]) Push(v T) {
	if len(c.s) == c.len {
		panic(errors.New("circular buffer full"))
	}
	p := (c.head + c.len) % len(c.s)
	c.s[p] = v
	c.len++
}

func (c *Circular[T]) PushAll(vs ...T) {
	if len(c.s) < c.len+len(vs) {
		panic(errors.New("circular buffer full"))
	}
	// TODO: make this efficient with copy()
	for _, v := range vs {
		p := (c.head + c.len) % len(c.s)
		c.s[p] = v
		c.len++
	}
}

func (c *Circular[T]) Pop() T {
	if c.len == 0 {
		panic(errors.New("circular buffer empty"))
	}
	v := c.s[c.head]
	// clear value for GC
	c.s[c.head] = Zero[T]()
	c.head = (c.head + 1) % len(c.s)
	c.len--
	return v
}

func (c *Circular[T]) Len() int {
	return c.len
}

func (c *Circular[T]) Cap() int {
	return len(c.s)
}

func (c *Circular[T]) All() []T {
	if c.len == 0 {
		return nil
	}
	t := c.head + c.len
	if t <= len(c.s) {
		// no wrap, just return a sub-slice, hide exxtra cap however
		return c.s[c.head:t:t]
	}
	t1 := len(c.s)
	t2 := t % len(c.s)
	return append(c.s[c.head:t1:t1], c.s[0:t2]...)
}
