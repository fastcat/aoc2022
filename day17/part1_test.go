package day17

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	jets := parse(sample)
	a.Len(jets, len(sample)-1)

	var b board
	var pos boardPos
	for i := 0; i < 2022; i++ {
		pos = b.play(shapes, jets, pos)
	}
	a.Equal(3068, b.height())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	jets := parse(input)
	var b board
	var pos boardPos
	for i := 0; i < 2022; i++ {
		pos = b.play(shapes, jets, pos)
	}
	t.Log(b.height())
}
