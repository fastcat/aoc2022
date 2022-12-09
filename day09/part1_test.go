package day09

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	moves := parseMoveList(strings.NewReader(sample))
	a.Equal(
		[]move{
			{'R', 4},
			{'U', 4},
			{'L', 3},
			{'D', 1},
			{'R', 4},
			{'D', 1},
			{'L', 5},
			{'R', 2},
		},
		moves,
	)
	state := NewState(1)
	state.apply(moves)
	a.Equal(pos{2, 2}, state.head)
	a.Equal(pos{1, 2}, state.tails[0])
	a.Len(state.tailVisited, 13)
	visited := maps.Keys(state.tailVisited)
	slices.SortFunc(visited, func(a, b pos) bool {
		if a[1] == b[1] {
			return a[0] < b[0]
		}
		return a[1] < b[1]
	})
	a.Equal(
		[]pos{
			{0, 0}, {1, 0}, {2, 0}, {3, 0},
			{4, 1},
			{1, 2}, {2, 2}, {3, 2}, {4, 2},
			{3, 3}, {4, 3},
			{2, 4}, {3, 4},
		},
		visited,
	)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	moves := parseMoveList(strings.NewReader(input))
	state := NewState(1)
	state.apply(moves)
	t.Log(len(state.tailVisited))
}
