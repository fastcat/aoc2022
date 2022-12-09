package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	g := parseGrid(sample)
	best, bestr, bestc := g.bestScore()
	a.Equal(8, best)
	a.Equal(3, bestr)
	a.Equal(2, bestc)
}

func TestPart2(t *testing.T) {
	g := parseGrid(input)
	best, bestr, bestc := g.bestScore()
	t.Log(best, bestr, bestc)
}
