package day03

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	p := parseWide(sample)
	b := findBadge(p)
	a.Equal([]int{18, 52}, i.ToSlice(b))
}

func TestPart2(t *testing.T) {
	p := parseWide(input)
	b := findBadge(p)
	s := i.Sum(b)
	t.Log(s)
}

func findBadge(in i.Iterable[[]rune]) i.Iterable[int] {
	threes := i.Chunk(in, 3)
	return i.Map(threes, func(group [][]rune, _ int) int {
		groupSets := i.Map(i.Slice(group), i.NoIndex(prioritySet))
		common := i.Reduce(groupSets, int64(0), func(common int64, bag int64, i int) int64 {
			if i == 0 {
				return bag
			}
			return common & bag
		})
		return singleBit(common)
	})
}

func parseWide(in string) i.Iterable[[]rune] {
	r := i.Runes(in)
	l := i.Split[rune](r, []rune{'\n'})
	return l
}
