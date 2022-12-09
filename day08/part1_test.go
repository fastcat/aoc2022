package day08

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	g := parseGrid(sample)
	a.Equal(
		grid{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
		g,
	)
	a.Equal(21, g.visibleTotal())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	g := parseGrid(input)
	t.Log(g.visibleTotal())
}
