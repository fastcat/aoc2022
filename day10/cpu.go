package day10

import (
	"fmt"

	"github.com/fastcat/aoc2022/i"
)

type instr interface {
	run(*cpu) i.Iterator[state]
}

type cpu struct {
	clock int
	state state
}

func newCPU() *cpu {
	return &cpu{
		clock: 0,
		state: state{x: 1},
	}
}

type state struct {
	x int
}

type noop struct{}

func (noop) run(c *cpu) i.Iterator[state] {
	done := false
	return i.Func(func() (state, bool) {
		if done {
			return state{}, true
		}
		c.clock++
		done = true
		return c.state, false
	})
}

type addx int

func (a addx) run(c *cpu) i.Iterator[state] {
	n := 0
	return i.Func(func() (state, bool) {
		switch n {
		case 0:
			c.clock++
			n++
			return c.state, false
		case 1:
			c.clock++
			n++
			s := c.state
			// incr only happens _after_ this cycle completes
			c.state.x += int(a)
			return s, false
		default:
			return c.state, true
		}
	})
}

func (c *cpu) run(prog []instr) i.Iterator[state] {
	n := -1
	var it i.Iterator[state]
	return i.Func(func() (state, bool) {
		for {
			if it == nil {
				n++
				if n >= len(prog) {
					return c.state, true
				}
				it = prog[n].run(c)
			}
			s, done := it.Next()
			if done {
				it = nil
				continue
			}
			return s, false
		}
	})
}

func parseProg(in string) []instr {
	lines := i.Split[rune](i.Runes(in), []rune{'\n'})
	return i.ToSlice(i.Map(lines, func(l []rune, _ int) instr {
		ls := string(l)
		if ls == "noop" {
			return noop{}
		}
		var a addx
		if _, err := fmt.Sscanf(ls, "addx %d", &a); err != nil {
			panic(err)
		}
		return a
	}))
}
