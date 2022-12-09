package day08

import (
	"errors"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
)

type grid []row
type row []int

func parseGrid(in string) grid {
	lines := i.Split[rune](i.Runes(in), []rune{'\n'})
	rows := i.Map(lines, func(l []rune, idx int) row {
		row := make(row, len(l))
		for i, r := range l {
			if r >= '0' && r <= '9' {
				row[i] = int(r - '0')
			}
		}
		return row
	})
	return grid(i.ToSlice(rows))
}

func (g grid) visibleTotal() int {
	visMap := make([][]bool, len(g))
	n := 0
	vis := func(r, c int) {
		if !visMap[r][c] {
			visMap[r][c] = true
			n++
		}
	}
	for r, row := range g {
		visMap[r] = make([]bool, len(row))
		i.For(visibleIndexes(i.Slice(row)), func(c, _ int) {
			vis(r, c)
		})
		i.For(visibleIndexes(i.RevSlice(row)), func(revc, _ int) {
			vis(r, len(row)-1-revc)
		})
	}
	nc := len(g[0])
	ri := i.Range(0, len(g), 1)
	rir := i.Range(len(g)-1, -1, -1)
	for c := 0; c < nc; c++ {
		v := func(r, _ int) int { return g[r][c] }
		i.For(visibleIndexes(i.Map(ri, v)), func(r, _ int) { vis(r, c) })
		i.For(visibleIndexes(i.Map(rir, v)), func(revr, _ int) { vis(len(g)-1-revr, c) })
	}
	return n
}

func (g grid) bestScore() (best, bestr, bestc int) {
	scores := [4][][]int{}
	// rows left to right
	scores[0] = i.ToSlice(i.Map(
		i.Slice(g),
		func(row row, idx int) []int {
			return i.ToSlice(visibleScores(i.Slice(row)))
		},
	))
	// rows right to left
	scores[1] = i.ToSlice(i.Map(
		i.Slice(g),
		func(row row, idx int) []int {
			s := i.ToSlice(visibleScores(i.RevSlice(row)))
			u.Reverse(s)
			return s
		},
	))
	// columns
	scores[2] = make([][]int, len(g))
	scores[3] = make([][]int, len(g))
	for r, row := range g {
		scores[2][r] = make([]int, len(row))
		scores[3][r] = make([]int, len(row))
	}
	nc := len(g[0])
	ri := i.Range(0, len(g), 1)
	rir := i.Range(len(g)-1, -1, -1)
	for c := 0; c < nc; c++ {
		v := func(r, _ int) int { return g[r][c] }
		// top to bottom
		i.For(
			visibleScores(i.Map(ri, v)),
			func(s, r int) {
				scores[2][r][c] = s
			},
		)
		// bottom to top
		i.For(
			visibleScores(i.Map(rir, v)),
			func(s, revr int) {
				scores[3][len(g)-1-revr][c] = s
			},
		)
	}
	best, bestr, bestc = 0, -1, -1
	for r := range g {
		for c := range g[r] {
			s := scores[0][r][c] * scores[1][r][c] * scores[2][r][c] * scores[3][r][c]
			if s > best {
				best, bestr, bestc = s, r, c
			}
		}
	}
	return best, bestr, bestc
}

func visibleIndexes(in i.Iterable[int]) i.Iterable[int] {
	return i.Funcer(func() i.Iterator[int] {
		max, idx := 0, 0
		it := in.Iterator()
		return i.Func(func() (int, bool) {
			for {
				v, done := it.Next()
				if done {
					return 0, true
				}
				idx++
				if idx == 1 || v > max {
					max = v
					return idx - 1, false
				}
			}
		})
	})
}

func visibleScores(in i.Iterable[int]) i.Iterable[int] {
	return i.Funcer(func() i.Iterator[int] {
		// only digits 1-9 are possible
		last := make([]int, 10)
		return i.MapI(in.Iterator(), func(v, idx int) int {
			if v < 0 || v > 9 {
				panic(errors.New("digit out of rante"))
			}
			stop := i.Max(i.Slice(last[v:]))
			score := idx - stop
			last[v] = idx
			return score
		})
	})
}
