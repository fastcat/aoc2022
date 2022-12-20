package day17

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_board_placeShape(t *testing.T) {
	tests := []struct {
		from board
		r    int
		c    int
		s    shape
		want board
	}{
		{
			nil,
			0, 0, shapes[0],
			board{0b1111},
		},
		{
			nil,
			0, 0, shapes[1],
			board{0b010, 0b111, 0b010},
		},
		{
			nil,
			0, 0, shapes[2],
			board{0b111, 0b100, 0b100},
		},
		{
			board{0b101, 0, 0b101},
			0, 0, shapes[1],
			board{0b111, 0b111, 0b111},
		},
		{
			board{0b11, 0b11},
			1, 1, shapes[4],
			board{0b11, 0b111, 0b110},
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			b := tt.from
			b.placeShape(tt.r, tt.c, tt.s)
			if !assert.Equal(t, tt.want, b) {
				t.Log("want:\n" + tt.want.String())
				t.Log("got:\n" + b.String())
			}
		})
	}
}

func Test_board_collides(t *testing.T) {
	tests := []struct {
		b    board
		r    int
		c    int
		s    shape
		want bool
	}{
		{
			board{0b101, 0, 0b101},
			0, 0, shapes[1],
			false,
		},
		{
			board{0b11, 0b11},
			1, 1, shapes[4],
			true,
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.b.collides(tt.r, tt.c, tt.s))
		})
	}
}

func Test_board_play(t *testing.T) {
	var b board
	jets := parse(sample)
	jo := 0
	b.play(shapes[0], jets, &jo)
	// t.Log("\n" + b.String())
	b.play(shapes[1], jets, &jo)
	// t.Log("\n" + b.String())
	b.play(shapes[2], jets, &jo)
	// t.Log("\n" + b.String())
	b.play(shapes[3], jets, &jo)
	// t.Log("\n" + b.String())
	b.play(shapes[4], jets, &jo)
	// t.Log("\n" + b.String())
	assert.Equal(t,
		board{
			0b111100,
			0b1000,
			0b11100,
			0b11111,
			0b10100,
			0b10100,
			0b10000,
			0b110000,
			0b110000,
		},
		b,
	)
}
