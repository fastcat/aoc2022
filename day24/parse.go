package day24

import (
	"fmt"

	"github.com/fastcat/aoc2022/i"
)

func parse(in string) *board {
	var b board
	i.For(
		i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
		func(l string, idx int) {
			if len(l) < 4 {
				panic(fmt.Errorf("can't detect end with less than 2 columns"))
			}
			if idx == 0 {
				// TODO: verify it's ^#\.#+$
				if l[1] != blank {
					panic(fmt.Errorf("top line missing entry"))
				}
				b.dims.c = len(l) - 2 // walls don't count
			} else if l[1] != wall {
				for d := left; d <= down; d++ {
					b.l[d].occupied = append(b.l[d].occupied, make([]bool, b.dims.c))
				}
				for c := 0; c < len(l)-2; c++ {
					v := l[c+1]
					switch v {
					case leftBlizzard:
						b.l[left].occupied[idx-1][c] = true
					case rightBlizzard:
						b.l[right].occupied[idx-1][c] = true
					case upBlizzard:
						b.l[up].occupied[idx-1][c] = true
					case downBlizzard:
						b.l[down].occupied[idx-1][c] = true
					case blank: // no-op
					default:
						panic(fmt.Errorf("invalid cell '%c'", v))
					}
				}
				b.dims.r++
			}
		},
	)
	for d := left; d <= down; d++ {
		b.l[d].dir = d
		b.l[d].dims = b.dims
	}
	return &b
}
