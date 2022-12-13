package day12

import (
	"fmt"

	"github.com/fastcat/aoc2022/i"
)

type board struct {
	elevations [][]int
	start      [2]int
	end        [2]int
}

// this flips the board top to bottom, but that doesn't matter for the problem
func parseBoard(in string) board {
	var b board
	i.For(
		i.Split(i.Runes(in), []rune{'\n'}),
		func(l []rune, r int) {
			b.elevations = append(b.elevations, make([]int, len(l)))
			for c, e := range l {
				switch e {
				case 'S':
					b.elevations[r][c] = 0
					b.start = [2]int{r, c}
				case 'E':
					b.elevations[r][c] = 25
					b.end = [2]int{r, c}
				default:
					if e < 'a' || e > 'z' {
						panic(fmt.Errorf("bad elevation '%c' at %d,%d", e, r, c))
					}
					b.elevations[r][c] = int(e - 'a')
				}
			}
		},
	)
	return b
}

func (b *board) valid(pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(b.elevations) &&
		pos[1] >= 0 && pos[1] < len(b.elevations[pos[0]])
}

func (b *board) at(pos [2]int) int {
	return b.elevations[pos[0]][pos[1]]
}

func (b *board) starts() i.Iterable[[2]int] {
	return i.Filter(
		i.Many(
			i.Range(0, len(b.elevations), 1),
			func(row, _ int) i.Iterable[[2]int] {
				return i.Map(i.Range(0, len(b.elevations[row]), 1), func(col, _ int) [2]int {
					return [2]int{row, col}
				})
			},
		),
		func(p [2]int) bool {
			return b.at(p) == 0
		},
	)
}
