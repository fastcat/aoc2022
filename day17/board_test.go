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
	a := assert.New(t)
	var b board
	jets := parse(sample)
	var pos boardPos
	for i := 0; i < 5; i++ {
		pos = b.play(shapes, jets, pos)
		// t.Log("\n" + b.String())
		// t.Log(pos)
	}
	a.Equal(
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
	// msb->lsb is top->down
	expectedLastRows := uint64(
		0b_0110000_0110000_0010000_0010100_0010100_0011111_0011100_0001000_0111100,
	)
	// t.Logf("%063b", expectedLastRows)
	if !a.Equal(
		expectedLastRows,
		b.lastRows(),
	) {
		t.Logf("%063b", b.lastRows())
	}
	if !a.Equal(
		expectedLastRows,
		pos.lastRows,
	) {
		t.Logf("%063b", pos.lastRows)
	}
}

func Test_board_playLoop(t *testing.T) {
	a := assert.New(t)
	var b board
	jets := parse(sample)
	posList := b.playLoop(shapes, jets, boardPos{}, 5)
	a.Len(posList, 5)
	// t.Log(posList)
	expectedLastRows := uint64(
		0b_0110000_0110000_0010000_0010100_0010100_0011111_0011100_0001000_0111100,
	)
	// t.Logf("%063b", expectedLastRows)
	if !a.Equal(
		expectedLastRows,
		b.lastRows(),
	) {
		t.Logf("%063b", b.lastRows())
	}
	if !a.Equal(
		expectedLastRows,
		posList[4].lastRows,
	) {
		t.Logf("%063b", posList[4].lastRows)
	}

}
