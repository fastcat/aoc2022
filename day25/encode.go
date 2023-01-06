package day25

import (
	"strings"

	"github.com/fastcat/aoc2022/i"
)

func encode(value int) string {
	var q strings.Builder
	for value != 0 {
		d := value % 5
		value /= 5
		if d < 3 {
			q.WriteRune(enc[d])
		} else {
			// 3=>-2,4=>-1
			dd := d - 5
			value++ // "carry"
			q.WriteRune(enc[dd])
		}
	}
	return string(i.ToSlice(i.RevSlice(i.ToSlice(i.Runes(q.String())))))
}
