package day17

import "strings"

// bytes lsb->msb are rows bottom to top, bits are columns left to right
type shape uint32

func (s shape) height() int {
	for i := uint32(0); i < 4; i++ {
		if uint32(s)&(0xff<<(8*i)) == 0 {
			return int(i)
		}
	}
	return 4
}

func (s shape) width() int {
	for i := uint32(8); i > 0; i-- {
		if uint32(s)&(0x01010101<<(i-1)) != 0 {
			return int(i)
		}
	}
	return 0
}

func (s shape) String() string {
	var buf strings.Builder
	w, h := s.width(), s.height()
	buf.Grow(h * (w + 1))
	for r := h - 1; r >= 0; r-- {
		row := s.row(r)
		for c := 0; c < w; c++ {
			if row&(1<<c) == 0 {
				buf.WriteRune('.')
			} else {
				buf.WriteRune('@')
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

func (s shape) at(r, c int) bool {
	return uint32(s)&(1<<(c+8*r)) != 0
}
func (s shape) row(r int) row {
	return row((uint32(s) >> (8 * r)) & 0xff)
}

func (s shape) canMove(c int, d direction) bool {
	if d == left {
		return c > 0
	}
	return c+s.width() < rowWidth
}

var shapes = []shape{
	0b1111,
	0b010 + 0b111<<8 + 0b010<<16,
	0b111 + 0b100<<8 + 0b100<<16,
	1 + 1<<8 + 1<<16 + 1<<24,
	0b11 + 0b11<<8,
}
