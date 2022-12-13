package day10

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
	prog := parseProg(sample)
	// i.ForI(newCPU().trace(prog), func(tt trace, n int) {
	// 	t.Logf("n=%d clk=%d pc=%d instr=%#v, c.x=%d s.x=%d\n", n, tt.clock, tt.pc, tt.instr, tt.cpu.state.x, tt.state.x)
	// })
	signals := i.ToSliceI(skip(newCPU().run(prog), 20, 40))
	a.Equal(
		[]signal{
			{state{21}, 20},
			{state{19}, 60},
			{state{18}, 100},
			{state{21}, 140},
			{state{16}, 180},
			{state{18}, 220},
		},
		signals,
	)
	total := i.ReduceI(skip(newCPU().run(prog), 20, 40), 0, sumSignals)
	a.Equal(13140, total)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	prog := parseProg(input)
	total := i.ReduceI(skip(newCPU().run(prog), 20, 40), 0, sumSignals)
	t.Log(total)
}

type signal struct {
	state state
	n     int
}

func skip(in i.Iterator[state], first, skip int) i.Iterator[signal] {
	n := 0
	return i.Func(func() (signal, bool) {
		for {
			n++
			s, done := in.Next()
			if done {
				return signal{s, n}, true
			}
			if n == first {
				return signal{s, n}, false
			} else if (n-first)%skip == 0 {
				return signal{s, n}, false
			}
		}
	})
}

func sumSignals(sum int, signal signal, _ int) int {
	return sum + signal.state.x*signal.n
}
