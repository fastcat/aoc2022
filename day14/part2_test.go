package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	b := newBoard(pos{500, 0})
	b.loadWalls(sample)
	b.addFloor()
	// t.Log("\n" + b.String())
	a.Equal(93, b.fillSand())
	// t.Log("\n" + b.String())
}
func TestPart2(t *testing.T) {
	b := newBoard(pos{500, 0})
	b.loadWalls(input)
	b.addFloor()
	// t.Log("\n" + b.String())
	t.Log(b.fillSand())
	// t.Log("\n" + b.String())
}
