package day13

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	lines := i.Split(i.Runes(sample), []rune{'\n'})
	packets := i.ToSlice(i.Map(lines, func(l []rune, _ int) listItem {
		return parsePacket(i.Slice(l))
	}))
	div1 := listOf(listOf(2))
	div2 := listOf(listOf(6))
	packets = append(packets, div1, div2)
	slices.SortFunc(packets, func(a, b listItem) bool {
		return a.cmp(b) == less
	})
	div1pos, div2pos := posOf(packets, div1), posOf(packets, div2)
	a.Equal(10, div1pos)
	a.Equal(14, div2pos)
	a.Equal(140, div1pos*div2pos)
}

func TestPart2(t *testing.T) {
	lines := i.Split(i.Runes(input), []rune{'\n'})
	packets := i.ToSlice(i.Map(lines, func(l []rune, _ int) listItem {
		return parsePacket(i.Slice(l))
	}))
	div1 := listOf(listOf(2))
	div2 := listOf(listOf(6))
	packets = append(packets, div1, div2)
	slices.SortFunc(packets, func(a, b listItem) bool {
		return a.cmp(b) == less
	})
	div1pos, div2pos := posOf(packets, div1), posOf(packets, div2)
	t.Log(div1pos * div2pos)
}
