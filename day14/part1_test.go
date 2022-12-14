package day14

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b := newBoard(pos{500, 0})
	b.loadWalls(sample)
	a.Equal(pos{494, 0}, b.min)
	a.Equal(pos{503, 9}, b.max)
	// t.Log("\n" + b.String())
	a.Equal(24, b.fillSand())
	// t.Log("\n" + b.String())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	b := newBoard(pos{500, 0})
	b.loadWalls(input)
	// t.Log("\n" + b.String())
	t.Log(b.fillSand())
	// t.Log("\n" + b.String())
}
