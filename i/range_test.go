package i

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	tests := []struct {
		start, end, step int
		want             []int
	}{
		{
			0, 0, 1,
			nil,
		},
		{
			0, 0, -1,
			nil,
		},
		{
			0, 5, 1,
			[]int{0, 1, 2, 3, 4},
		},
		{
			4, -1, -1,
			[]int{4, 3, 2, 1, 0},
		},
		{
			0, 10, 2,
			[]int{0, 2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v->%v/%v", tt.start, tt.end, tt.step), func(t *testing.T) {
			assert.Equal(t, tt.want, ToSlice(Range(tt.start, tt.end, tt.step)))
		})
	}
}
