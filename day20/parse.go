package day20

import (
	"strconv"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
)

func parse(in string) *mixer {
	var m mixer
	m.orig = i.ToSlice(
		i.Map(
			i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
			i.NoIndex(u.Muster(strconv.Atoi)),
		),
	)
	m.reset()
	return &m
}
