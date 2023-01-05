package day24

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type test struct {
		in   string
		want board
	}
	tests := []test{
		{
			"" +
				"#.##\n" +
				"#..#\n" +
				"##.#\n",
			board{
				dims: pos{1, 2},
				l: [4]layer{
					{
						dir: left, dims: pos{1, 2},
						occupied: [][]bool{{false, false}},
					},
					{
						dir: right, dims: pos{1, 2},
						occupied: [][]bool{{false, false}},
					},
					{
						dir: up, dims: pos{1, 2},
						occupied: [][]bool{{false, false}},
					},
					{
						dir: down, dims: pos{1, 2},
						occupied: [][]bool{{false, false}},
					},
				},
			},
		},
		{
			"" +
				"#.##\n" +
				"#<>#\n" +
				"#^v#\n" +
				"##.#\n",
			board{
				dims: pos{2, 2},
				l: [4]layer{
					{
						dir: left, dims: pos{2, 2},
						occupied: [][]bool{{true, false}, {false, false}},
					},
					{
						dir: right, dims: pos{2, 2},
						occupied: [][]bool{{false, true}, {false, false}},
					},
					{
						dir: up, dims: pos{2, 2},
						occupied: [][]bool{{false, false}, {true, false}},
					},
					{
						dir: down, dims: pos{2, 2},
						occupied: [][]bool{{false, false}, {false, true}},
					},
				},
			},
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := parse(tt.in)
			assert.Equal(t, tt.want, *got)
		})
	}
}
