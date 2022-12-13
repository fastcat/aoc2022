package day10

import "github.com/fastcat/aoc2022/i"

func render(in i.Iterator[state], width int) i.Iterator[rune] {
	const radius = 1
	const lit = '#'
	const dark = '.'
	n := 0
	return i.Func(func() (rune, bool) {
		s, done := in.Next()
		if done {
			return 0, true
		}
		delta := (n % width) - s.x
		n++
		if delta >= -radius && delta <= radius {
			return lit, false
		} else {
			return dark, false
		}
	})
}

func scanLines(in i.Iterator[rune], width int) i.Iterator[string] {
	return i.MapI(
		i.ChunkI(in, width),
		func(l []rune, _ int) string { return string(l) },
	)
}
