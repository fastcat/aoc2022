package day23

import (
	"strconv"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		in   string
		want []pos
	}{
		{
			"",
			[]pos{},
		},
		{
			"#.#",
			[]pos{{0, 0}, {0, 2}},
		},
		{
			"#.#\n.#.\n",
			[]pos{{0, 0}, {0, 2}, {1, 1}},
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			want := &board{
				m: i.KeysToMap(i.Slice(tt.want), func(pos) bool { return true }),
			}
			assert.Equal(t, want, parse(tt.in))
		})
	}
}

func Test_board_move(t *testing.T) {
	tests := []struct {
		in   string
		m    []move
		mo   int
		want string
	}{
		{
			".....\n..##.\n..#..\n.....\n..##.\n.....",
			moves[:], 0,
			"..##.\n.....\n..#..\n...#.\n..#..\n.....",
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			b := parse(tt.in)
			b.move(tt.m, tt.mo)
			want := parse(tt.want)
			if !assert.Equal(t, want, b) {
				t.Log("\n" + b.String())
				t.Log("\n" + want.String())
			}
		})
	}
}
