package i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	a := assert.New(t)
	AssertIterator(
		a,
		[]int{0, 1, 2, 3},
		Map(Slice([]int{0, 0, 0, 0}), func(_, i int) int { return i }).Iterator(),
	)
	AssertIterator(
		a,
		[]int{1, 3, 5, 7},
		Map(Slice([]int{0, 1, 2, 3}), func(v, _ int) int { return 1 + 2*v }).Iterator(),
	)
}
