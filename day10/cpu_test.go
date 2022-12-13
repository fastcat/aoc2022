package day10

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func TestCPU(t *testing.T) {
	a := assert.New(t)
	prog := parseProg("noop\naddx 3\naddx -5\n")
	a.Equal([]instr{noop{}, addx(3), addx(-5)}, prog)
	c := newCPU()
	states := i.ToSliceI(c.run(prog))
	a.Equal(
		[]state{
			{1},
			{1}, {1},
			{4}, {4},
		},
		states,
	)
}
