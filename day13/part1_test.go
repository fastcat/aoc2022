package day13

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	lines := i.Split(i.Runes(sample), []rune{'\n'})
	packets := i.Map(lines, func(l []rune, _ int) listItem {
		return parsePacket(i.Slice(l))
	})
	a.Equal(
		[]listItem{
			listOf(1, 1, 3, 1, 1),
			listOf(1, 1, 5, 1, 1),
			listOf(listOf(1), listOf(2, 3, 4)),
			listOf(listOf(1), 4),
			listOf(9),
			listOf(listOf(8, 7, 6)),
			listOf(listOf(4, 4), 4, 4),
			listOf(listOf(4, 4), 4, 4, 4),
			listOf(7, 7, 7, 7),
			listOf(7, 7, 7),
			listOf(),
			listOf(3),
			listOf(listOf(listOf())),
			listOf(listOf()),
			listOf(1, listOf(2, listOf(3, listOf(4, listOf(5, 6, 7)))), 8, 9),
			listOf(1, listOf(2, listOf(3, listOf(4, listOf(5, 6, 0)))), 8, 9),
		},
		i.ToSlice(packets),
	)
	pairs := i.Chunk(packets, 2)
	want := []cmpRes{
		less,
		less,
		greater,
		less,
		greater,
		less,
		greater,
		greater,
	}
	i.For(pairs, func(p []listItem, i int) {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := assert.New(t)
			a.Equal(equal, p[0].cmp(p[0]))
			a.Equal(equal, p[1].cmp(p[1]))
			a.Equal(want[i], p[0].cmp(p[1]))
		})
	})
	a.Equal([]int{1, 2, 4, 6}, i.ToSlice(rightPairs(pairs)))
	a.Equal(13, i.Sum(rightPairs(pairs)))
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	lines := i.Split(i.Runes(input), []rune{'\n'})
	packets := i.Map(lines, func(l []rune, _ int) listItem {
		return parsePacket(i.Slice(l))
	})
	pairs := i.Chunk(packets, 2)
	t.Log(i.Sum(rightPairs(pairs)))
}
