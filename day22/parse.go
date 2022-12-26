package day22

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/fastcat/aoc2022/i"
	"golang.org/x/exp/slices"
)

func parse(in string) (*board, []move) {
	bstr, mstr, found := strings.Cut(in, "\n\n")
	if !found {
		panic(fmt.Errorf("no stanza separator in input"))
	}
	b := parseBoard1(bstr)
	m := parseMoves(mstr)
	return b, m
}

func parseGrid(bstr string) grid {
	return i.ToSlice(i.Split(
		i.Map(i.Runes(bstr), func(r rune, _ int) place { return place(r) }),
		[]place{'\n'},
	))
}

func parseBoard1(bstr string) *board {
	grid := parseGrid(bstr)
	b := &board{g: grid}
	b.buildPortals1()
	return b
}

func (b *board) buildPortals1() {
	// every row & column should have 2 portals
	b.portals = make(map[state]state, 2*len(b.g)+2*len(b.g[0]))
	// rows
	for r, row := range b.g {
		first := slices.IndexFunc(row, func(p place) bool { return p != void })
		if first < 0 {
			panic(fmt.Errorf("no voids in row %d", r+1))
		}
		last := first + slices.Index(row[first+1:], void)
		if last < first {
			last = len(row) - 1
		}
		// left from the left edge leads you to the right edge
		b.portals[state{r, first, left}] = state{r, last, left}
		// vice versa for right edge
		b.portals[state{r, last, right}] = state{r, first, right}
	}

	// columns
	nCols := len(b.g[0])
	rowIdx := i.Range(0, len(b.g), 1)
	for c := 0; c < nCols; c++ {
		first, last := -1, -1
		i.For(rowIdx, func(r, _ int) {
			if first < 0 && b.g[r][c] != void {
				first = r
			} else if first >= 0 && last < 0 && b.g[r][c] == void {
				last = r - 1
			}
		})
		if first < 0 {
			panic(fmt.Errorf("no voids in column %d", c+1))
		}
		if last < 0 {
			last = len(b.g) - 1
		}

		// up from the top edge leaves you to the bottom edge
		b.portals[state{first, c, up}] = state{last, c, up}
		// vice versa for bottom edge
		b.portals[state{last, c, down}] = state{first, c, down}
	}
}

func parseMoves(mstr string) []move {
	return i.ToSlice(i.Map(
		i.Merge(
			i.Runes(strings.TrimSpace(mstr)),
			func(r rune) string { return string(r) },
			func(prior string, next rune) (string, bool) {
				if !unicode.IsDigit(next) {
					return prior, false
				} else if unicode.IsDigit(rune(prior[0])) {
					return prior + string(next), true
				}
				return prior, false
			},
		),
		func(s string, _ int) move {
			switch s {
			case "L":
				return turnLeft
			case "R":
				return turnRight
			default:
				if v, err := strconv.Atoi(s); err != nil {
					panic(err)
				} else {
					return move(v)
				}
			}
		},
	))
}
