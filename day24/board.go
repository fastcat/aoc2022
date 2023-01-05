package day24

import (
	"fmt"
	"strings"
)

type pos struct {
	r, c int
}

func (p pos) scale(factor int) pos {
	return pos{p.r * factor, p.c * factor}
}
func (p pos) add(o pos) pos {
	return pos{p.r + o.r, p.c + o.c}
}
func (p pos) modclamp(m pos) pos {
	return pos{modclamp(p.r, m.r), modclamp(p.c, m.c)}
}
func modclamp(v, m int) int {
	return ((v % m) + m) % m
}

type direction int

const (
	left direction = iota
	right
	up
	down
)

func (d direction) offset() pos {
	switch d {
	case left:
		return pos{0, -1}
	case right:
		return pos{0, 1}
	case up:
		return pos{-1, 0}
	case down:
		return pos{1, 0}
	default:
		panic(fmt.Errorf("bad directino %d", d))
	}
}
func (d direction) name() rune {
	switch d {
	case left:
		return leftBlizzard
	case right:
		return rightBlizzard
	case up:
		return upBlizzard
	case down:
		return downBlizzard
	default:
		panic(fmt.Errorf("bad directino %d", d))
	}
}

type board struct {
	dims pos
	l    [4]layer
}

func (b *board) initialPos() pos {
	// player starts just above the board
	return pos{-1, 0}
}
func (b *board) targetPos() pos {
	// player tries to get to just below the bottom right
	return pos{b.dims.r, b.dims.c - 1}
}
func (b *board) initialState() *state {
	return b.customState(b.initialPos(), 0)
}
func (b *board) customState(p pos, minute int) *state {
	return &state{
		b:         b,
		minute:    minute,
		reachable: map[pos]bool{p: true},
	}
}

type layer struct {
	dir      direction
	dims     pos
	occupied [][]bool
}

type layerView struct {
	l      *layer
	offset pos
}

type boardView struct {
	l [4]layerView
}

func (b *board) viewAt(minute int) *boardView {
	var bv boardView
	for i := 0; i < len(b.l); i++ {
		bv.l[i] = b.l[i].viewAt(minute)
	}
	return &bv
}
func (l *layer) viewAt(minute int) layerView {
	return layerView{l, l.dir.offset().scale(-minute)}
}

func (bv *boardView) occupied(p pos) bool {
	for _, lv := range bv.l {
		if lv.occupied(p) {
			return true
		}
	}
	return false
}
func (lv *layerView) occupied(p pos) bool {
	// edges
	if p.r == -1 {
		return p.c != 0
	} else if p.r < -1 || p.c < 0 || p.c >= lv.l.dims.c {
		return true
	} else if p.r == lv.l.dims.r {
		return p.c != lv.l.dims.c-1
	} else if p.r > lv.l.dims.r {
		return true
	}
	// body
	p = p.add(lv.offset).modclamp(lv.l.dims)
	return lv.l.occupied[p.r][p.c]
}
func (bv *boardView) occupancy(p pos) rune {
	if p.r < 0 || p.r >= bv.l[0].l.dims.r || p.c < 0 || p.c >= bv.l[0].l.dims.c {
		if bv.occupied(p) {
			return wall
		}
		return blank
	}
	o := blank
	n := 0
	for _, lv := range bv.l {
		if !lv.occupied(p) {
			continue
		} else if n == 0 {
			n, o = 1, lv.l.dir.name()
		} else {
			n++
		}
	}
	if n < 2 {
		return o
	}
	return '0' + rune(n)
}

func (bv *boardView) String() string {
	dims := bv.l[0].l.dims
	var buf strings.Builder
	var p pos
	for p.r = -1; p.r <= dims.r; p.r++ {
		for p.c = -1; p.c <= dims.c; p.c++ {
			buf.WriteRune(bv.occupancy(p))
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

const (
	blank         = '.'
	wall          = '#'
	leftBlizzard  = '<'
	rightBlizzard = '>'
	upBlizzard    = '^'
	downBlizzard  = 'v'
)
