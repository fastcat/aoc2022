package day24

import (
	"fmt"
	"strings"
)

type move int

const (
	wait move = iota
	moveLeft
	moveRight
	moveUp
	moveDown
)

func (m move) offset() pos {
	switch m {
	case wait:
		return pos{}
	case moveLeft:
		return left.offset()
	case moveRight:
		return right.offset()
	case moveUp:
		return up.offset()
	case moveDown:
		return down.offset()
	default:
		panic(fmt.Errorf("invalid move %d", m))
	}
}

type state struct {
	b         *board
	minute    int
	reachable map[pos]bool
}

func (s *state) next() *state {
	v := s.b.viewAt(s.minute + 1)
	newReachable := make(map[pos]bool, len(s.reachable))
	for p := range s.reachable {
		for m := wait; m <= moveDown; m++ {
			pp := p.add(m.offset())
			if !v.occupied(pp) {
				newReachable[pp] = true
			}
		}
	}
	return &state{s.b, s.minute + 1, newReachable}
}

func (s *state) boardString() string {
	return s.b.viewAt(s.minute).String()
}

func (s *state) reachableString() string {
	var buf strings.Builder
	// top row
	buf.WriteRune(wall)
	if s.reachable[s.b.initialPos()] {
		buf.WriteRune('E')
	} else {
		buf.WriteRune(blank)
	}
	for i := 0; i < s.b.dims.c; i++ {
		buf.WriteRune(wall)
	}
	buf.WriteRune('\n')
	// interior
	var p pos
	for p.r = 0; p.r < s.b.dims.r; p.r++ {
		buf.WriteRune(wall)
		for p.c = 0; p.c < s.b.dims.c; p.c++ {
			if s.reachable[p] {
				buf.WriteRune('E')
			} else {
				buf.WriteRune(blank)
			}
		}
		buf.WriteRune(wall)
		buf.WriteRune('\n')
	}
	// bottom row
	for i := 0; i < s.b.dims.c; i++ {
		buf.WriteRune(wall)
	}
	if s.reachable[s.b.targetPos()] {
		buf.WriteRune('E')
	} else {
		buf.WriteRune(blank)
	}
	buf.WriteRune(wall)
	buf.WriteRune('\n')
	return buf.String()
}
