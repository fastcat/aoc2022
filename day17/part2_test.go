package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	var b board
	jets := parse(sample)
	lf := b.loopFinder(shapes, jets)

	const targetRounds = 1_000_000_000_000
	finalHeight := lf.heightAfter(targetRounds)
	a.Equal(1514285714288, finalHeight)
}

func TestPart2(t *testing.T) {
	var b board
	jets := parse(input)
	lf := b.loopFinder(shapes, jets)

	const targetRounds = 1_000_000_000_000
	finalHeight := lf.heightAfter(targetRounds)
	t.Log(finalHeight)
}
