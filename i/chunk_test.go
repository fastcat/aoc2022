package i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	for _, tt := range []struct {
		name     string
		items    []int
		size     int
		expected [][]int
	}{
		{"empty", nil, 1, nil},
		{"singles", []int{1, 2, 3, 4}, 1, [][]int{{1}, {2}, {3}, {4}}},
		{
			"even",
			[]int{1, 2, 3, 4, 5, 6}, 2,
			[][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			"remainder",
			[]int{1, 2, 3, 4, 5, 6, 7}, 3,
			[][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			s := Chunk(Slice(tt.items), tt.size)
			it := s.Iterator()
			AssertIterator(a, tt.expected, it)
		})
	}
}
