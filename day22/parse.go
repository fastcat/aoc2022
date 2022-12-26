package day22

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/fastcat/aoc2022/i"
	"golang.org/x/exp/slices"
)

func parse(in string, portals func(*board)) (*board, []move) {
	bstr, mstr, found := strings.Cut(in, "\n\n")
	if !found {
		panic(fmt.Errorf("no stanza separator in input"))
	}
	b := &board{g: parseGrid(bstr)}
	portals(b)
	b.trace = make(grid, len(b.g))
	for r, row := range b.g {
		b.trace[r] = make([]place, len(row))
		copy(b.trace[r], row)
	}
	m := parseMoves(mstr)
	return b, m
}

func parseGrid(bstr string) grid {
	grid := i.ToSlice(i.Split(
		i.Map(i.Runes(bstr), func(r rune, _ int) place { return place(r) }),
		[]place{'\n'},
	))
	// normalize all rows to have the same number of columns, filling the gaps
	// with void
	maxCols := i.Max(i.Map(i.Slice(grid), func(r []place, _ int) int { return len(r) }))
	for r, row := range grid {
		if len(row) < maxCols {
			rr := make([]place, maxCols)
			copy(rr, row)
			for c := len(row); c < maxCols; c++ {
				rr[c] = void
			}
			grid[r] = rr
		}
	}
	return grid
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

func (b *board) buildPortals2s() {
	// every row & column should have 2 portals
	b.portals = make(map[state]state, 2*len(b.g)+2*len(b.g[0]))

	// 12 edges total, 4 pre-connected in the cube net, so a full list here would
	// have 8, but only 6 are needed to compute the sample
	b.connectEdges(
		state{4, 7, up}, left,
		state{3, 8, left}, up,
		4,
	)
	b.connectEdges(
		state{4, 3, up}, left,
		state{0, 8, up}, right,
		4,
	)
	b.connectEdges(
		state{7, 7, down}, left,
		state{8, 8, left}, down,
		4,
	)
	b.connectEdges(
		state{7, 11, right}, up,
		state{8, 12, up}, right,
		4,
	)
	b.connectEdges(
		state{4, 0, left}, down,
		state{11, 15, down}, left,
		4,
	)
	b.connectEdges(
		state{7, 3, down}, left,
		state{11, 8, down}, right,
		4,
	)
}

func (b *board) buildPortals2i() {
	// every row & column should have 2 portals
	b.portals = make(map[state]state, 2*len(b.g)+2*len(b.g[0]))
	const l = 50
	if len(b.g) != l*4 || len(b.g[0]) != l*3 {
		panic(fmt.Errorf("bad grid"))
	}
	/*
		 AB
		 C
		DE
		F
	*/
	// 12 edges total, 5 pre-connected, 7 to add
	b.connectEdges(
		// F left t-b meets A up l-r
		state{0, l, up}, right,
		state{3 * l, 0, left}, down,
		l,
	)
	b.connectEdges(
		// A left t-b meets D left b-t
		state{0, l, left}, down,
		state{3*l - 1, 0, left}, up,
		l,
	)
	b.connectEdges(
		// D up l-r meets C left t-b
		state{2 * l, 0, up}, right,
		state{l, l, left}, down,
		l,
	)
	b.connectEdges(
		// F right t-b meets E down l-r
		state{3 * l, l - 1, right}, down,
		state{3*l - 1, l, down}, right,
		l,
	)
	b.connectEdges(
		// F down l-r meets B up l-r
		state{4*l - 1, 0, down}, right,
		state{0, 2 * l, up}, right,
		l,
	)
	b.connectEdges(
		// B right t-b meets E right b-t
		state{0, 3*l - 1, right}, down,
		state{3*l - 1, 2*l - 1, right}, up,
		l,
	)
	b.connectEdges(
		// C right t-b meets B down l-r
		state{l, 2*l - 1, right}, down,
		state{l - 1, 2 * l, down}, right,
		l,
	)
}

func (b *board) connectEdges(s1 state, s1d dir, s2 state, s2d dir, l int) {
	// s1 and s2 are the first cells in the edges, facing the portal direction (so
	// s1 faces s2 and vice versa). each edge is l cells long (including s1/s2),
	// moving in the respective s1d/s2d direction
	s1dr, s1dc := s1d.delta()
	s2dr, s2dc := s2d.delta()
	p1df := s1.d.flip()
	p2df := s2.d.flip()
	for i, p1, p2 := 0, s1, s2; i < l; i, p1.r, p1.c, p2.r, p2.c = i+1, p1.r+s1dr, p1.c+s1dc, p2.r+s2dr, p2.c+s2dc {
		// each portal lands 180 from its return directin
		b.portals[p1] = state{p2.r, p2.c, p2df}
		b.portals[p2] = state{p1.r, p1.c, p1df}
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
