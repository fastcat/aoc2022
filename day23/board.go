package day23

import (
	"fmt"
	"math"
	"strings"

	"github.com/fastcat/aoc2022/i"
)

type pos struct {
	r, c int
}

func (p pos) add(p2 pos) pos {
	return pos{p.r + p2.r, p.c + p2.c}
}

type board struct {
	m map[pos]bool
}

const blank = '.'
const elf = '#'

func (b *board) String() string {
	min, max, _ := b.bounds()
	var buf strings.Builder
	buf.Grow((max.r - min.r + 1) * (max.c - min.c + 2))
	for r := min.r; r <= max.r; r++ {
		for c := min.c; c <= max.c; c++ {
			if b.m[pos{r, c}] {
				buf.WriteRune(elf)
			} else {
				buf.WriteRune(blank)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func parse(in string) *board {
	b := board{m: make(map[pos]bool, len(in))}
	var p pos
	i.For(
		i.Runes(in),
		func(v rune, _ int) {
			if v == elf {
				b.m[p] = true
				p.c++
			} else if v == '\n' {
				p.r, p.c = p.r+1, 0
			} else if v == blank {
				p.c++
			} else {
				panic(fmt.Errorf("invalid input '%c'", v))
			}
		},
	)
	return &b
}

func (b *board) emptyGround() int {
	min, max, elves := b.bounds()
	// empty = area - fill
	area := (max.r - min.r + 1) * (max.c - min.c + 1)
	return area - elves
}

func (b *board) bounds() (min, max pos, elves int) {
	min = pos{math.MaxInt, math.MaxInt}
	max = pos{math.MinInt, math.MinInt}
	for p, v := range b.m {
		if !v {
			panic(fmt.Errorf("false in map"))
		} else {
			elves++
		}
		if p.r < min.r {
			min.r = p.r
		}
		if p.r > max.r {
			max.r = p.r
		}
		if p.c < min.c {
			min.c = p.c
		}
		if p.c > max.c {
			max.c = p.c
		}
	}
	return
}

func (b *board) lonely(p pos) bool {
	for pp := (pos{p.r - 1, p.c}); pp.r <= p.r+1; pp.r++ {
		for pp.c = p.c - 1; pp.c <= p.c+1; pp.c++ {
			if pp == p {
				continue
			} else if b.m[pp] {
				return false
			}
		}
	}
	return true
}

func (b *board) move(mm []move, mo int) bool {
	// count proposed targets
	targets := map[pos]int{}
	props := map[pos]pos{}
	for p, v := range b.m {
		if !v {
			panic(fmt.Errorf("false in map"))
		}
		if b.lonely(p) {
			continue
		}
	MOVE:
		for mi := 0; mi < len(mm); mi++ {
			m := mm[(mo+mi)%len(mm)]
			for _, cp := range m.check {
				if b.m[p.add(cp)] {
					continue MOVE
				}
			}
			pp := p.add(m.dp)
			props[p] = pp
			targets[pp]++
			break
		}
	}
	if len(props) == 0 {
		return false
	}
	n := 0
	for p := range b.m {
		if pp, ok := props[p]; !ok {
			continue
		} else if t := targets[pp]; t < 1 {
			panic(fmt.Errorf("props/targets inconsistency"))
		} else if t > 1 {
			continue
		} else if b.m[pp] {
			panic(fmt.Errorf("target collision"))
		} else {
			delete(b.m, p)
			b.m[pp] = true
			n++
		}
	}
	if n == 0 {
		panic(fmt.Errorf("stuck"))
	}
	return true
}

func (b *board) moveN(mm []move, mo, n int) int {
	for i := 0; i < n; i++ {
		if !b.move(mm, mo+i) {
			return i
		}
	}
	return n
}

type move struct {
	dp    pos
	check [3]pos
}

var moves = [4]move{
	{
		pos{-1, 0},
		[3]pos{
			{-1, -1},
			{-1, 0},
			{-1, 1},
		},
	},
	{
		pos{1, 0},
		[3]pos{
			{1, -1},
			{1, 0},
			{1, 1},
		},
	},
	{
		pos{0, -1},
		[3]pos{
			{-1, -1},
			{0, -1},
			{1, -1},
		},
	},
	{
		pos{0, 1},
		[3]pos{
			{-1, 1},
			{0, 1},
			{1, 1},
		},
	},
}
