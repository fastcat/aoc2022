package day22

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b, m := parse(sample, (*board).buildPortals1)
	a.Equal(12, len(b.g))
	a.Equal(
		[]move{
			10, turnRight,
			5, turnLeft,
			5, turnRight,
			10, turnLeft,
			4, turnRight,
			5, turnLeft,
			5,
		},
		m,
	)
	a.Equal(b.portals[state{0, 8, left}], state{0, 11, left})
	a.Equal(b.portals[state{7, 4, down}], state{4, 4, down})
	a.Equal(
		state{0, 10, right},
		b.move(state{0, 8, right}, 10),
	)

	s := b.initialState()
	a.Equal(state{0, 8, right}, s)
	fs := b.moves(s, m...)
	a.Equal(state{5, 7, right}, fs)
	a.Equal(6032, fs.Value())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	b, m := parse(input, (*board).buildPortals1)
	s := b.initialState()
	fs := b.moves(s, m...)
	t.Log(fs, fs.Value())
}
