package day23

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	b := parse(sample)
	a.Equal(27, b.emptyGround())
	n := b.moveN(moves[:], 0, 21)
	a.Equal(20, n+1)
	want20 := parse(".......#......\n....#......#..\n..#.....#.....\n......#.......\n...#....#.#..#\n#.............\n....#.....#...\n..#.....#.....\n....#.#....#..\n.........#....\n....#......#..\n.......#......")
	a.Equal(want20.String(), b.String())
}

func TestPart2(t *testing.T) {
	b := parse(input)
	m := 0
	for {
		n := b.moveN(moves[:], 0, len(moves))
		if n == 0 {
			break
		}
		m += n
	}
	t.Log(m + 1)
}
