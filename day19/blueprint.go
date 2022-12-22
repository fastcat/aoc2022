package day19

import (
	"fmt"
	"unicode"

	"github.com/fastcat/aoc2022/i"
)

type idx int

const (
	ore      idx = 0
	clay     idx = 1
	obsidian idx = 2
	geode    idx = 3
	none     idx = -1
)

type cost [3]uint8

type blueprint [4]cost

func parseOne(in string) (int, *blueprint) {
	var n int
	var b blueprint
	if _, err := fmt.Sscanf(
		in,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n",
		&n,
		&b[ore][ore],
		&b[clay][ore],
		&b[obsidian][ore], &b[obsidian][clay],
		&b[geode][ore], &b[geode][obsidian],
	); err != nil {
		panic(err)
	}
	return n, &b
}

func parseMany(in string) []*blueprint {
	return i.ToSlice(
		i.Map(
			i.Merge(
				i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
				func(s string) string { return s },
				func(prior, next string) (merged string, merge bool) {
					if unicode.IsSpace(rune(next[0])) {
						return prior + next, true
					}
					return prior, false
				},
			),
			func(in string, idx int) *blueprint {
				n, b := parseOne(in)
				if n != idx+1 {
					panic(fmt.Errorf("wrong blueprint no, expect %d got %d", idx+1, n))
				}
				return b
			},
		),
	)
}
