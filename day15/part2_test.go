package day15

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	b := newBoard()
	b.parseSensors(sample)
	w := interval{0, 20}
	for y := w.low; y <= w.high; y++ {
		e := i.ToSlice(clip(i.Slice(b.beaconExcludedInRow(y)), w))
		if y == 11 {
			a.Equal([]interval{{0, 13}, {15, w.high}}, e, "expected y=11,x=14 open")
		} else {
			a.Equal([]interval{w}, e, "expected y=%d full", y)
		}
	}
	a.Equal(14*4000000+11, 56000011)
}

func TestPart2(t *testing.T) {
	a := assert.New(t)
	b := newBoard()
	b.parseSensors(input)
	w := interval{0, 4000000}
	for y := w.low; y <= w.high; y++ {
		e := i.ToSlice(clip(i.Slice(b.beaconExcludedInRow(y)), w))
		if len(e) == 1 {
			a.Equal([]interval{w}, e)
		} else if len(e) == 2 {
			a.Equal(e[0].high+2, e[1].low)
			x := e[0].high + 1
			tf := x*4000000 + y
			t.Logf("gap at x=%d,y=%d, tf=%d", x, y, tf)
		} else {
			a.Fail("max len=2")
		}
	}
}
