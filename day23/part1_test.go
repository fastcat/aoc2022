package day23

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b := parse(sample)
	a.Equal(27, b.emptyGround())
	b.moveN(moves[:], 0, 10)
	want10 := parse("......#.....\n..........#.\n.#.#..#.....\n.....#......\n..#.....#..#\n#......##...\n....##......\n.#........#.\n...#.#..#...\n............\n...#..#..#..")
	a.Equal(want10.String(), b.String())
	a.Equal(110, b.emptyGround())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	b := parse(input)
	b.moveN(moves[:], 0, 10)
	t.Log(b.emptyGround())
}
