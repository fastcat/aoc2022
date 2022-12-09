package day09

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	moves := parseMoveList(strings.NewReader(sample))
	state := NewState(9)
	state.apply(moves)
	a.Equal(
		[]pos{
			{1, 2},
			{2, 2},
			{3, 2},
			{2, 2},
			{1, 1},
			{0, 0},
			{0, 0},
			{0, 0},
			{0, 0},
		},
		state.tails,
	)
}

//go:embed sample2.txt
var sample2 string

func TestPart2Sample2(t *testing.T) {
	a := assert.New(t)
	moves := parseMoveList(strings.NewReader(sample2))
	state := NewState(9)
	state.apply(moves)
	a.Equal(pos{-11, 15}, state.head)
	a.Equal(
		[]pos{
			{-11, 14},
			{-11, 13},
			{-11, 12},
			{-11, 11},
			{-11, 10},
			{-11, 9},
			{-11, 8},
			{-11, 7},
			{-11, 6},
		},
		state.tails,
	)
	a.Len(state.tailVisited, 36)
}

func TestPart2(t *testing.T) {
	moves := parseMoveList(strings.NewReader(input))
	state := NewState(9)
	state.apply(moves)
	t.Log(len(state.tailVisited))
}
