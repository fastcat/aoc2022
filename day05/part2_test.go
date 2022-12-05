package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	state, moves := parseStateAndMoves(sample)
	for _, m := range moves {
		state = state.Move2(m)
	}
	a.Equal(
		[]stack{
			{'M'},
			{'C'},
			{'P', 'Z', 'N', 'D'},
		},
		state.stacks,
	)
	a.Equal(
		[]rune{'M', 'C', 'D'},
		state.Tops(),
	)
}

func TestPart2(t *testing.T) {
	state, moves := parseStateAndMoves(input)
	for _, m := range moves {
		state = state.Move2(m)
	}
	t.Log(string(state.Tops()))
}
