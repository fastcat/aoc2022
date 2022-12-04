package i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTop(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{5, 4, 3}, Top(Slice([]int{1, 2, 3, 4, 5}), 3))
	a.Equal([]int{5, 4}, Top(Slice([]int{4, 5}), 3))
}
