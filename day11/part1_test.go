package day11

import (
	_ "embed"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	g := parseGame(sample, div3)
	a.Len(g, 4)
	a.Equal(
		[][]int{
			{79, 98},
			{54, 65, 75, 74},
			{79, 60, 97},
			{74},
		},
		i.ToSlice(i.Map(i.Slice(g), func(m *monkey, _ int) []int { return m.items.All() })),
	)
	a.Equal(
		[][2]int{
			{2, 3},
			{2, 0},
			{1, 3},
			{0, 1},
		},
		i.ToSlice(i.Map(i.Slice(g), func(m *monkey, _ int) [2]int { return m.targets })),
	)
	g[0].step(g)
	a.Equal([]int{98}, g[0].items.All())
	a.Equal([]int{74, 500}, g[3].items.All())
	a.Equal(2*19, g[0].op(2))
	a.Equal(18+6, g[1].op(18))
	a.Equal(7*7, g[2].op(7))
	a.Equal(12+3, g[3].op(12))

	g = parseGame(sample, div3)
	g.rounds(20)
	a.Equal(
		[]int{101, 95, 7, 105},
		g.inspections(),
	)
	a.Equal(10605, g.business())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	g := parseGame(input, div3)
	g.rounds(20)
	t.Log(g.business())
}
