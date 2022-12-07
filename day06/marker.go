package day06

import (
	"errors"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
)

// findMarker returns the number of input runes after which the last l runes are
// unique. If no unique set is found, it returns -1. The smallest value it can
// return on success is l. It will panic if l is less than 1.
func findMarker(it i.Iterator[rune], l int) int {
	if l < 1 {
		panic(errors.New("bad length"))
	}
	c := u.NewCircular[rune](l)
	i := 0
	for {
		v, done := it.Next()
		if done {
			return -1
		}
		i++
		if c.Len() == l {
			c.Pop()
			c.Push(v)
			if isUnique(c.All()) {
				return i
			}
		} else {
			c.Push(v)
		}
	}
}

func isUnique(s []rune) bool {
	var seen int64
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			v := r - 'a'
			m := int64(1) << v
			if seen&m != 0 {
				return false
			}
			seen = seen | m
		}
	}
	return true
}
