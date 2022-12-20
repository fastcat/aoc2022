package day17

import (
	"fmt"
	"strings"

	"github.com/fastcat/aoc2022/i"
)

type direction bool

const left direction = false
const right direction = true

func (d direction) apply(c int) int {
	if d == left {
		return c - 1
	}
	return c + 1
}

func parse(in string) []direction {
	return i.ToSlice(i.Map(
		i.Runes(strings.TrimSpace(in)),
		func(d rune, _ int) direction {
			switch d {
			case '<':
				return left
			case '>':
				return right
			default:
				panic(fmt.Errorf("invalid direction %c", d))
			}
		},
	))
}
