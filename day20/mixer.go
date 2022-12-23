package day20

import (
	"github.com/fastcat/aoc2022/i"
	"golang.org/x/exp/slices"
)

type mixer struct {
	// pos[i] contains indexes in orig, so orig[pos[i]] is the value in the mixed
	// slice at index i
	pos []int
	// orig is the
	orig []int
}

func (m *mixer) reset() {
	m.pos = i.ToSlice(i.Range(0, len(m.orig), 1))
}

func (m *mixer) mixOne(idx int) {
	l := len(m.orig)
	l1 := l - 1
	v := m.orig[idx]
	p := slices.Index(m.pos, idx)
	// wrapping is weird, one left of 0 is not max, it's max-1, and one right of
	// max is not 0, it's one. If you look at it from the point of view of
	// swapping elements in a circular view of the list, the swap would rotate the
	// whole list a bit. to compensate for not doing that.
	// 0 1 2 3 4 5 6 | 0 1 2 3 4 5 6 // initial
	// 1 0 2 3 4 5 6 | 1 0 2 3 4 5 6 // 1 shifts left 1
	// 6 0 2 3 4 5 1 | 6 0 2 3 4 5 1 // 1 shifts left 2, but now 6 has moved
	// 0 2 3 4 5 1 6 | 0 2 3 4 5 1 6 // 1 shifts left 2, corrected
	// 0 2 3 4 1 5 6 | 0 2 3 4 1 5 6 // 1 shifts left 3
	// 0 2 3 1 4 5 6 | 0 2 3 1 4 5 6 // 1 shifts left 4
	// 0 2 1 3 4 5 6 | 0 2 1 3 4 5 6 // 1 shifts left 5
	// 0 1 2 3 4 5 6 | 0 1 2 3 4 5 6 // 1 shifts left 6
	// 1 0 2 3 4 5 6 | 1 0 2 3 4 5 6 // 1 shifts left 7, looks odd but ok
	// 6 0 2 3 4 5 1 | 6 0 2 3 4 5 1 // 1 shifts left 8, but needs correction again
	// 0 2 3 4 5 1 6 | 0 2 3 4 5 1 6 // 1 shifts left 8, corrected
	// a similar thing happens when wrapping to the right, we need to add an extra
	// move to compensate for not shifting the whole slice
	dest := (p + v) % l1
	if dest < 0 {
		dest += l1
	}
	// adapt equivalent behavior to match test case setups
	if v < 0 && dest == 0 {
		dest = l1
	} else if v > l1 && dest == 0 {
		dest = l1
	}
	if dest < p {
		// copy up
		copy(m.pos[dest+1:p+1], m.pos[dest:p])
	} else if dest > p {
		// copy down
		copy(m.pos[p:dest], m.pos[p+1:dest+1])
	}
	m.pos[dest] = idx
}

func (m *mixer) mixAll() {
	for i := range m.orig {
		m.mixOne(i)
	}
}
func (m *mixer) mixAllRepeat(n int) {
	for i := 0; i < n; i++ {
		m.mixAll()
	}
}

func (m *mixer) All() []int {
	v := make([]int, len(m.orig))
	for i, p := range m.pos {
		v[i] = m.orig[p]
	}
	return v
}

func (m *mixer) IndexOf(v int) int {
	return slices.Index(m.pos, slices.Index(m.orig, v))
}

func (m *mixer) At(i int) int {
	i = i % len(m.pos)
	return m.orig[m.pos[i]]
}

func (m *mixer) OffsetSum(v int, pp ...int) int {
	idx := m.IndexOf(v)
	return i.Sum(i.Map(i.Slice(pp), func(p, _ int) int { return m.At(idx + p) }))
}

func (m *mixer) Mult(v int) {
	for i, o := range m.orig {
		m.orig[i] = o * v
	}
}
