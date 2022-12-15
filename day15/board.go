package day15

import (
	"fmt"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
	"golang.org/x/exp/slices"
)

type pos struct{ x, y int }

func (p pos) mhDistTo(o pos) int {
	return u.Abs(p.x-o.x) + u.Abs(p.y-o.y)
}

type interval struct{ low, high int }

func (i interval) size() int {
	return i.high - i.low + 1
}

func (i interval) overlaps(o interval) bool {
	return o.low <= i.high && o.high >= i.low
}
func (i interval) canMerge(o interval) bool {
	// can merge if overlapping or adjacent
	return o.low <= i.high+1 && o.high >= i.low-1
}

func (i interval) clipTo(w interval) interval {
	if i.low < w.low {
		i.low = w.low
	}
	if i.high > w.high {
		i.high = w.high
	}
	return i
}

func (i interval) merge(o interval) interval {
	if o.low < i.low {
		i.low = o.low
	}
	if o.high > i.high {
		i.high = o.high
	}
	return i
}

func (i interval) contains(v int) bool {
	return i.low <= v && v <= i.high
}

func (i interval) splitAt(v int) []interval {
	if v == i.low {
		i.low++
		return []interval{i}
	} else if v == i.high {
		i.high--
		return []interval{i}
	}
	return []interval{
		{i.low, v - 1},
		{v + 1, i.high},
	}
}

func clip(in i.Iterable[interval], window interval) i.Iterable[interval] {
	return i.Map(
		i.Filter(in, window.overlaps),
		func(iv interval, _ int) interval {
			return iv.clipTo(window)
		},
	)
}

type board struct {
	sp []pos
	bp []pos
	d  []int
}

func newBoard() *board {
	return &board{}
}

func (b *board) parseSensors(in string) {
	i.For(
		i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
		func(l string, _ int) {
			var sp, bp pos
			if _, err := fmt.Sscanf(
				l,
				"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n",
				&sp.x, &sp.y, &bp.x, &bp.y,
			); err != nil {
				panic(err)
			}
			b.sp = append(b.sp, sp)
			b.bp = append(b.bp, bp)
			b.d = append(b.d, sp.mhDistTo(bp))
		},
	)
}

func (b *board) isBeaconExcludedSlow(c pos) bool {
	for i, sp := range b.sp {
		if c == b.bp[i] || c == sp {
			return false
		}
		if b.d[i] >= sp.mhDistTo(c) {
			return true
		}
	}
	return false
}

func (b *board) beaconExcludedInRow(y int) []interval {
	var ivs []interval
	for i, sp := range b.sp {
		if bp := b.bp[i]; bp.y == y {
			ivs = append(ivs, interval{bp.x, bp.x})
		}
		dy := u.Abs(sp.y - y)
		dx := b.d[i] - dy
		if dx < 0 {
			continue
		}
		ivs = append(ivs, interval{sp.x - dx, sp.x + dx})
	}
	slices.SortFunc(ivs, func(a, b interval) bool {
		if a.low < b.low {
			return true
		} else if a.low == b.low {
			return a.high < b.high
		}
		return false
	})
	mi := make([]interval, 0, len(ivs))
	for i, iv := range ivs {
		if i == 0 {
			mi = append(mi, iv)
		} else if li := mi[len(mi)-1]; li.canMerge(iv) {
			mi[len(mi)-1] = li.merge(iv)
		} else {
			mi = append(mi, iv)
		}
	}
	return mi
}

func (b *board) beaconExcludedInRowBlanks(y int) []interval {
	mi := b.beaconExcludedInRow(y)
	for _, l := range [][]pos{b.sp, b.bp} {
		for _, p := range l {
			if p.y != y {
				continue
			}
			si := make([]interval, 0, len(mi))
			for _, iv := range mi {
				if iv.contains(p.x) {
					si = append(si, iv.splitAt(p.x)...)
				} else {
					si = append(si, iv)
				}
			}
			mi = si
		}
	}
	return mi
}
