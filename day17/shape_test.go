package day17

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shape_height(t *testing.T) {
	tests := []struct {
		s    shape
		want int
	}{
		{shapes[0], 1},
		{shapes[1], 3},
		{shapes[2], 3},
		{shapes[3], 4},
		{shapes[4], 2},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s.height())
		})
	}
}

func Test_shape_width(t *testing.T) {
	tests := []struct {
		s    shape
		want int
	}{
		{shapes[0], 4},
		{shapes[1], 3},
		{shapes[2], 3},
		{shapes[3], 1},
		{shapes[4], 2},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s.width())
		})
	}
}
