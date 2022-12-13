package day10

import (
	"strings"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

const width = 40

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	prog := parseProg(sample)
	scan := i.ToSliceI(scanLines(render(newCPU().run(prog), width), width))
	a.Equal(
		[]string{
			"##..##..##..##..##..##..##..##..##..##..",
			"###...###...###...###...###...###...###.",
			"####....####....####....####....####....",
			"#####.....#####.....#####.....#####.....",
			"######......######......######......####",
			"#######.......#######.......#######.....",
		},
		scan,
	)
}

func TestPart2(t *testing.T) {
	prog := parseProg(input)
	i.ForI(scanLines(render(newCPU().run(prog), width), width), func(sl string, _ int) {
		t.Log(strings.ReplaceAll(sl, ".", " "))
	})
}
