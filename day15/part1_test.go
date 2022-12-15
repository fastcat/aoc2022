package day15

import (
	_ "embed"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b := newBoard()
	b.parseSensors(sample)
	a.Len(b.sp, 14)
	a.Len(b.bp, 14)
	a.Len(b.d, 14)
	n := 0
	for x := -5; x < 30; x++ {
		got := b.isBeaconExcludedSlow(pos{x, 10})
		if x >= -2 && x <= 24 && x != 2 /* becaon at 2,10 */ {
			a.True(got, "expect true at x=%d", x)
			n++
		} else {
			a.False(got, "expect false at x=%d", x)
		}
	}
	a.Equal(26, n)

	ei := b.beaconExcludedInRowBlanks(10)
	a.Equal(
		[]interval{
			{-2, 1},
			{3, 24},
		},
		ei,
	)
	a.Equal(26, i.Sum(i.Map(i.Slice(ei), i.NoIndex(interval.size))))
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	b := newBoard()
	b.parseSensors(input)
	ei := b.beaconExcludedInRowBlanks(2000000)
	n := i.Sum(i.Map(i.Slice(ei), i.NoIndex(interval.size)))
	t.Log(n)
}
