package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircular(t *testing.T) {
	a := assert.New(t)

	c := NewCircular[int](10)
	e := &Circular[int]{s: make([]int, 10)}
	a.Equal(e, c)
	a.Empty(c.All())

	c.Push(1)
	e.s[0], e.len = 1, 1
	a.Equal(e, c)
	a.Equal([]int{1}, c.All())

	v := c.Pop()
	a.Equal(1, v)
	e.s[0], e.head, e.len = 0, 1, 0
	a.Equal(e, c)
	a.Empty(c.All())

	c.PushAll(1, 2, 3, 4, 5)
	copy(e.s[1:], []int{1, 2, 3, 4, 5})
	e.len = 5
	a.Equal(e, c)
	a.Equal([]int{1, 2, 3, 4, 5}, c.All())
	for i := 0; i < 5; i++ {
		a.Equal(i+1, c.Pop())
	}
	copy(e.s[1:], []int{0, 0, 0, 0, 0})
	e.head, e.len = 6, 0
	a.Equal(e, c)

	c.PushAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	copy(e.s, []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4})
	e.len = 10
	a.Equal(e, c)
	a.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, c.All())
}
