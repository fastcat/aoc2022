package day12

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	b := parseBoard(sample)
	p := newPather(b)
	p.search()
	a.Equal(
		[][2]int{
			{0, 0},
			{0, 1},
			{1, 0},
			{2, 0},
			{3, 0},
			{4, 0},
		},
		i.ToSlice(b.starts()),
	)
	p.bestStart()
	a.Equal([2]int{4, 0}, p.b.start)
	a.Equal(29, p.distFromStart())
}

func TestPart2(t *testing.T) {
	b := parseBoard(input)
	p := newPather(b)
	p.search()
	p.bestStart()
	t.Log(p.b.start, p.distFromStart())
}
