package day19

import (
	"fmt"
	"unicode"

	"github.com/fastcat/aoc2022/i"
)

type cost struct {
	ore, clay, obsidian int
}

func (c cost) canBuild(i inventory) bool {
	return i.ore >= c.ore && i.clay >= c.clay && i.obsidian >= c.obsidian
}

type blueprint struct {
	oreBot, clayBot, obsidianBot, geodeBot cost
}

func parseOne(in string) *blueprint {
	var n int
	var b blueprint
	if _, err := fmt.Sscanf(
		in,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n",
		&n,
		&b.oreBot.ore,
		&b.clayBot.ore,
		&b.obsidianBot.ore, &b.obsidianBot.clay,
		&b.geodeBot.ore, &b.geodeBot.obsidian,
	); err != nil {
		panic(err)
	}
	return &b
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
			i.NoIndex(parseOne),
		),
	)
}
