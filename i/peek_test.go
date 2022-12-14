package i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeeker(t *testing.T) {
	a := assert.New(t)
	r := Range(0, 5, 1)
	a.Equal(ToSlice(r), ToSliceI[int](Peeker(r.Iterator())))
	p := Peeker(r.Iterator())
	v, d := p.Peek()
	a.Equal([]any{0, false}, []any{v, d})
	v, d = p.Peek()
	a.Equal([]any{0, false}, []any{v, d})
	v, d = p.Next()
	a.Equal([]any{0, false}, []any{v, d})
}
