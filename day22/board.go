package day22

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type place byte

const (
	void  place = ' '
	space place = '.'
	wall  place = '#'
)

func (p place) String() string {
	return string([]byte{'\'', byte(p), '\''})
}

type grid [][]place

type board struct {
	g       grid
	trace   grid
	portals map[state]state
}

type dir byte

const (
	right dir = '>'
	down  dir = 'v'
	left  dir = '<'
	up    dir = '^'
)

func (d dir) String() string { return string(d) }

func (d dir) Value() int {
	i := slices.Index(turns, d)
	if i < 0 {
		panic(fmt.Errorf("invalid dir: %v", d))
	}
	return i
}

func (d dir) delta() (dr, dc int) {
	switch d {
	case right:
		dc = 1
	case down:
		dr = 1
	case left:
		dc = -1
	case up:
		dr = -1
	default:
		panic(fmt.Errorf("invalid dir %v", d))
	}
	return
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

func (d dir) flip() dir {
	switch d {
	case right:
		return left
	case down:
		return up
	case left:
		return right
	case up:
		return down
	default:
		panic(fmt.Errorf("invalid dir %v", d))
	}
}

type state struct {
	r, c int
	d    dir
}

type move int

const (
	turnLeft  move = -1
	turnRight move = -2
)

func (b *board) record(s state) {
	b.trace[s.r][s.c] = place(s.d)
}

func (b *board) traceString() string {
	buf := strings.Builder{}
	for _, row := range b.trace {
		for _, p := range row {
			buf.WriteRune(rune(p))
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func (b *board) move(s state, m move) state {
	b.record(s)
	if m == turnLeft || m == turnRight {
		s.d = s.d.turn(m)
		b.record(s)
		return s
	}
	dr, dc := s.d.delta()
	for i := 0; i < int(m); i++ {
		ns := state{r: s.r + dr, c: s.c + dc, d: s.d}
		if p, ok := b.portals[s]; ok {
			ns = p
			dr, dc = ns.d.delta()
		}
		if ns.r < 0 || ns.c < 0 || ns.r >= len(b.g) || ns.c >= len(b.g[ns.r]) {
			panic(fmt.Errorf("should not step off grid at r=%d,c=%d,d=%v", s.r, s.c, s.d))
		}
		if np := b.g[ns.r][ns.c]; np == wall {
			// stop at a wall
			return s
		} else if np == space {
			// ok
			s = ns
			b.record(s)
		} else {
			panic(fmt.Errorf("should not step into place %v at r=%d,c=%d,d=%v", np, s.r, s.c, s.d))
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
