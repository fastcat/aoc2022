package day08

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func Test_parseGrid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want grid
	}{
		{
			"1x1",
			"1",
			grid{{1}},
		},
		{
			"3x3",
			"123\n456\n789\n",
			grid{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseGrid(tt.in))
		})
	}
}

func Test_visibleIndexes(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			"incr",
			[]int{1, 2, 3, 4, 5},
			[]int{0, 1, 2, 3, 4},
		},
		{
			"scatter",
			[]int{5, 4, 6, 4, 7, 4},
			[]int{0, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, i.ToSlice(visibleIndexes(i.Slice(tt.in))))
		})
	}
}

func Test_visibleScores(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			"incr",
			[]int{1, 2, 3, 4, 5},
			[]int{0, 1, 2, 3, 4},
		},
		{
			"alter",
			[]int{3, 5, 3, 5, 3},
			[]int{0, 1, 1, 2, 1},
		},
		{
			"scatter2",
			[]int{0, 9, 3, 5, 3, 7},
			[]int{0, 1, 1, 2, 1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, i.ToSlice(visibleScores(i.Slice(tt.in))))
		})
	}
}
