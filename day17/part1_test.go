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
	j := parse(sample)
	a.Len(j, len(sample)-1)

	var b board
	for i, jo := 0, 0; i < 2022; i++ {
		b.play(shapes[i%len(shapes)], j, &jo)
	}
	a.Equal(3068, b.height())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	j := parse(input)
	var b board
	for i, jo := 0, 0; i < 2022; i++ {
		b.play(shapes[i%len(shapes)], j, &jo)
	}
	t.Log(b.height())
}
