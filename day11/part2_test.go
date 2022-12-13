package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	g := parseGame(sample, nil)
	g.useGCD()
	g.rounds(1)
	a.Equal([]int{2, 4, 3, 6}, g.inspections())
	g.rounds(20 - 1)
	a.Equal([]int{99, 97, 8, 103}, g.inspections())
	g.rounds(1000 - 20)
	a.Equal([]int{5204, 4792, 199, 5192}, g.inspections())
	g.rounds(2000 - 1000)
	a.Equal([]int{10419, 9577, 392, 10391}, g.inspections())
	g.rounds(3000 - 2000)
	a.Equal([]int{15638, 14358, 587, 15593}, g.inspections())
	g.rounds(4000 - 3000)
	a.Equal([]int{20858, 19138, 780, 20797}, g.inspections())
	g.rounds(5000 - 4000)
	g.rounds(6000 - 5000)
	g.rounds(7000 - 6000)
	g.rounds(8000 - 7000)
	g.rounds(9000 - 8000)
	a.Equal([]int{46945, 43051, 1746, 46807}, g.inspections())
	g.rounds(10000 - 9000)
	a.Equal([]int{52166, 47830, 1938, 52013}, g.inspections())

	g = parseGame(sample, nil)
	g.useGCD()
	g.rounds(10000)
	a.Equal([]int{52166, 47830, 1938, 52013}, g.inspections())
	a.Equal(2713310158, g.business())
}

func TestPart2(t *testing.T) {
	g := parseGame(input, nil)
	g.useGCD()
	g.rounds(10000)
	t.Log(g.business())
}
