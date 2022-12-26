package day22

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type place byte

const (
	void  place = ' '
	space place = '.'
	wall  place = '#'
)

type grid [][]place

type board struct {
	g       grid
	portals map[state]state
}

type dir byte

const (
	right dir = '>'
	down  dir = 'v'
	left  dir = '<'
	up    dir = '^'
)

func (d dir) Value() int {
	i := slices.Index(turns, d)
	if i < 0 {
		panic(fmt.Errorf("invalid dir: %v", d))
	}
	return i
}

var turns = []dir{right, down, left, up}

func (d dir) turn(m move) dir {
	v := d.Value()
	switch m {
	case turnLeft:
		v = (v + 3) % 4
	case turnRight:
		v = (v + 1) % 4
	default:
		panic(fmt.Errorf("not a turn: %v", m))
	}
	return turns[v]
}

type state struct {
	r, c int
	d    dir
}

func (s state) Row() int    { return s.r + 1 }
func (s state) Column() int { return s.c + 1 }

type move int

const (
	turnLeft  move = -1
	turnRight move = -2
)

func (b *board) move(s state, m move) state {
	if m == turnLeft || m == turnRight {
		s.d = s.d.turn(m)
		return s
	}
	dr, dc := 0, 0
	switch s.d {
	case right:
		dc = 1
	case down:
		dr = 1
	case left:
		dc = -1
	case up:
		dr = -1
	default:
		panic(fmt.Errorf("invalid dir %v", s.d))
	}
	for i := 0; i < int(m); i++ {
		ns := state{r: s.r + dr, c: s.c + dc, d: s.d}
		if p, ok := b.portals[s]; ok {
			ns = p
		}
		if np := b.g[ns.r][ns.c]; np == wall {
			// stop at a wall
			return s
		} else if np == space {
			// ok
			s = ns
		} else {
			panic(fmt.Errorf("should not step into place %v at r=%d,c=%d", np, ns.r, ns.c))
		}
	}
	return s
}

func (b *board) moves(s state, moves ...move) state {
	for _, m := range moves {
		s = b.move(s, m)
	}
	return s
}

func (b *board) initialState() state {
	s := state{d: right}
	for b.g[s.r][s.c] == void {
		s.c++
	}
	return s
}

func (s state) Value() int {
	return 1000*(s.r+1) + 4*(s.c+1) + s.d.Value()
}
