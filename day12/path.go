package day12

import (
	"fmt"
	"math"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
)

type pather struct {
	b    board
	dist [][]int
	next [][]rune
}

func newPather(b board) *pather {
	p := &pather{b: b}
	p.dist = make([][]int, len(b.elevations))
	p.next = make([][]rune, len(b.elevations))
	for r, row := range b.elevations {
		p.dist[r] = make([]int, len(row))
		for c := range p.dist[r] {
			p.dist[r][c] = math.MaxInt - 1
		}
		p.next[r] = make([]rune, len(row))
	}
	return p
}

func (p *pather) search() {
	p.dist[p.b.end[0]][p.b.end[1]] = 0
	p.next[p.b.end[0]][p.b.end[1]] = 'E'
	q := u.NewCircular[[2]int](8)
	q.PushAll(neighbors(p.b.end)...)
	for q.Len() > 0 {
		pos := q.Pop()
		if !p.b.valid(pos) {
			continue
		}
		e := p.b.at(pos)
		d := p.dist[pos[0]][pos[1]]
		for _, n := range neighbors(pos) {
			if !p.b.valid(n) {
				continue
			}
			nd := p.dist[n[0]][n[1]] + 1
			ne := p.b.at(n)
			if ne <= e+1 && nd < d {
				// this is a better neighbor than what we have so far
				p.dist[pos[0]][pos[1]] = nd
				p.next[pos[0]][pos[1]] = dir(pos, n)
				// re-evaluate all our neighbors
				q.GrowPushAll(neighbors(pos)...)
			}
		}
	}
}

func (p *pather) distFrom(pos [2]int) int {
	return p.dist[pos[0]][pos[1]]
}

func (p *pather) distFromStart() int {
	return p.dist[p.b.start[0]][p.b.start[1]]
}

func (p *pather) bestStart() {
	bestStart := i.Reduce(
		p.b.starts(),
		p.b.start,
		func(best, cur [2]int, _ int) [2]int {
			if p.distFrom(cur) < p.distFrom(best) {
				return cur
			}
			return best
		},
	)
	p.b.start = bestStart
}

func (p *pather) pretty(line func([]rune)) {
	// wipe unused steps from the path
	pp := make([][]rune, len(p.next))
	for r, row := range p.next {
		pp[r] = make([]rune, len(row))
	}
	for pos := p.b.start; pos != p.b.end; pos = next(pos, p.next[pos[0]][pos[1]]) {
		pp[pos[0]][pos[1]] = p.next[pos[0]][pos[1]]
	}
	pp[p.b.end[0]][p.b.end[1]] = 'E'
	for _, row := range pp {
		rr := make([]rune, len(row))
		copy(rr, row)
		for c, n := range rr {
			if n == 0 {
				rr[c] = '.'
			}
		}
		line(rr)
	}
}

func neighbors(pos [2]int) [][2]int {
	return [][2]int{
		{pos[0] - 1, pos[1]},
		{pos[0] + 1, pos[1]},
		{pos[0], pos[1] - 1},
		{pos[0], pos[1] + 1},
	}
}

func dir(p1, p2 [2]int) rune {
	if p1[0] == p2[0] {
		if p1[1]+1 == p2[1] {
			return '>'
		} else if p1[1]-1 == p2[1] {
			return '<'
		}
	} else if p1[1] == p2[1] {
		if p1[0]+1 == p2[0] {
			return 'v'
		} else if p1[0]-1 == p2[0] {
			return '^'
		}
	}
	panic(fmt.Errorf("not neighbors %v %v", p1, p2))
}

func next(pos [2]int, dir rune) [2]int {
	switch dir {
	case '<':
		return [2]int{pos[0], pos[1] - 1}
	case '>':
		return [2]int{pos[0], pos[1] + 1}
	case '^':
		return [2]int{pos[0] - 1, pos[1]}
	case 'v':
		return [2]int{pos[0] + 1, pos[1]}
	}
	panic(fmt.Errorf("bad dir %c", dir))
}
