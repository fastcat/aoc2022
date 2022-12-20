package day17

import (
	"strings"
)

const rowWidth = 7

type row byte

func (r row) String() string {
	var buf strings.Builder
	buf.Grow(9)
	buf.WriteRune('|')
	for i := 0; i < rowWidth; i++ {
		if r&(1<<i) == 0 {
			buf.WriteRune('.')
		} else {
			buf.WriteRune('#')
		}
	}
	buf.WriteRune('|')
	return buf.String()
}

type board []row

func (b board) String() string {
	var buf strings.Builder
	for r := len(b) - 1; r >= 0; r-- {
		buf.WriteString(b[r].String())
		buf.WriteRune('\n')
	}
	buf.WriteString("+-------+\n")
	return buf.String()
}

func (b *board) placeShape(r, c int, s shape) {
	h := s.height()
	for y := 0; y < h; y++ {
		b.ensureRow(r + y)
		(*b)[r+y] |= row(s.row(y) << c)
	}
}
func (b *board) ensureRow(r int) {
	for len(*b) <= r {
		*b = append(*b, 0)
	}
}
func (b *board) placeRock(r, c int) {
	b.ensureRow(r)
	(*b)[r] |= 1 << c
}

func (b board) collides(r, c int, s shape) bool {
	h := s.height()
	for y := 0; y < h; y++ {
		if len(b) <= r+y {
			// board is empty from here on up
			return false
		}
		if b[r+y]&(s.row(y)<<c) != 0 {
			return true
		}
	}
	return false
}

func (b board) height() int {
	for r := len(b) - 1; r >= 0; r-- {
		if b[r] != 0 {
			return r + 1
		}
	}
	return 0
}

func (b *board) play(s shape, jets []direction, jo *int) {
	c := 2
	r := b.height() + 3
	// debug := func(m string) {
	// 	b2 := b.clone()
	// 	b2.placeShape(r, c, s)
	// 	fmt.Println(m + "\n" + b2.String())
	// }
	for {
		// debug("start round")
		// round part 1: jet
		d := jets[*jo]
		*jo = (*jo + 1) % len(jets)
		if nc := d.apply(c); s.canMove(c, d) && !b.collides(r, nc, s) {
			c = nc
		}
		// debug("after jet")
		if r == 0 || b.collides(r-1, c, s) {
			b.placeShape(r, c, s)
			return
		} else {
			r--
		}
	}
}

func (b board) clone() board {
	ret := make(board, len(b))
	copy(ret, b)
	return ret
}
