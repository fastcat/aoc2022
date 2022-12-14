package day14

import (
	"fmt"
	"math"
	"strings"

	"github.com/fastcat/aoc2022/i"
)

type pos struct{ x, y int }
type cell rune

const (
	air  cell = '.'
	wall cell = '#'
	sand cell = 'o'
	src  cell = '+'
)

type board struct {
	source   pos
	cells    map[pos]cell
	min, max pos
	floor    int
}

func newBoard(source pos) *board {
	b := &board{
		source: source,
		cells:  make(map[pos]cell),
		min:    source,
		max:    source,
		floor:  math.MaxInt,
	}
	b.setCell(source, src)
	return b
}

func (b *board) cell(p pos) cell {
	if c, ok := b.cells[p]; ok {
		return c
	}
	if p.y == b.floor {
		return wall
	}
	return air
}
func (b *board) setCell(p pos, c cell) {
	b.cells[p] = c
	if p.x < b.min.x {
		b.min.x = p.x
	} else if p.x > b.max.x {
		b.max.x = p.x
	}
	if p.y < b.min.y {
		b.min.y = p.y
	} else if p.y > b.max.y {
		b.max.y = p.y
	}
}
func (b *board) isVoid(p pos) bool {
	if b.floor < math.MaxInt {
		// if we have a floor, there is no let/right void
		return p.y > b.floor
	}
	// y goes down from the top, void is beyond the left, right, or bottom edges
	return p.x < b.min.x || p.x > b.max.x || p.y > b.max.y
}

func (b *board) loadWalls(in string) {
	i.For(
		i.Split(i.Runes(in), []rune{'\n'}),
		func(l []rune, _ int) {
			i.Reduce(
				i.Map(
					i.ToStrings(i.Split(i.Slice(l), []rune{' ', '-', '>', ' '})),
					func(s string, _ int) pos {
						var p pos
						if _, err := fmt.Sscanf(s, "%d,%d\n", &p.x, &p.y); err != nil {
							panic(err)
						}
						return p
					},
				),
				pos{},
				func(last, next pos, idx int) pos {
					b.setCell(next, wall)
					if idx == 0 {
						return next
					}
					// draw from last to next
					var inc pos
					if next.x > last.x {
						inc.x = 1
					} else if next.x < last.x {
						inc.x = -1
					}
					if next.y > last.y {
						inc.y = 1
					} else if next.y < last.y {
						inc.y = -1
					}
					for p := last; p != next; p.x, p.y = p.x+inc.x, p.y+inc.y {
						b.setCell(p, wall)
					}
					return next
				},
			)
		},
	)
}

func (b *board) addFloor() {
	b.max.y += 2
	b.floor = b.max.y
}

func (b *board) String() string {
	ret := &strings.Builder{}
	// skip the void
	for y := b.min.y; y <= b.max.y; y++ {
		for x := b.min.x; x <= b.max.x; x++ {
			ret.WriteRune(rune(b.cell(pos{x, y})))
		}
		ret.WriteRune('\n')
	}
	return ret.String()
}

func (b *board) addSand() *pos {
	p := b.source
	for {
		p2 := b.walkSand(p)
		if b.isVoid(p2) {
			return nil
		}
		if p2 == p {
			b.setCell(p, sand)
			return &p
		}
		p = p2
	}
}

func (b *board) walkSand(p pos) pos {
	if p2 := (pos{p.x, p.y + 1}); b.cell(p2) == air {
		return p2
	}
	if p2 := (pos{p.x - 1, p.y + 1}); b.cell(p2) == air {
		return p2
	}
	if p2 := (pos{p.x + 1, p.y + 1}); b.cell(p2) == air {
		return p2
	}
	return p
}

func (b *board) fillSand() int {
	for i := 0; ; i++ {
		if p := b.addSand(); p == nil {
			return i
		} else if *p == b.source {
			// we placed this unit of sand, count it
			return i + 1
		}
	}
}
