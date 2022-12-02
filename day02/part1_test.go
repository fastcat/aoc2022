package day02

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRPSVs(t *testing.T) {
	type test struct {
		a, b RPS
		o    Outcome
	}
	tests := []test{
		{Rock, Paper, Loss},
		{Rock, Scissors, Win},
		{Rock, Rock, Draw},
		{Paper, Scissors, Loss},
		{Scissors, Rock, Loss},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s vs %s", tt.a, tt.b), func(t *testing.T) {
			assert.Equal(t, tt.o, tt.a.vs(tt.b))
		})
	}
}

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	r := require.New(t)
	results := i.Map(parse1(sample), play1)
	r.Equal(
		[]round{
			{Rock, Paper, Win, 8},
			{Paper, Rock, Loss, 1},
			{Scissors, Scissors, Draw, 6},
		},
		i.ToSlice(results),
	)
	total := i.Sum(i.Map(results, func(r round) int { return r.score }))
	r.Equal(15, total)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	results := i.Map(parse1(input), play1)
	total := i.Sum(i.Map(results, func(r round) int { return r.score }))
	t.Log(total)
}

func parse1(in string) i.Iterable[[2]RPS] {
	lines := i.Split[rune](i.Runes(in), []rune{'\n'})
	parsed := i.Map(
		lines,
		func(l []rune) [2]RPS {
			if len(l) != 3 || l[1] != ' ' {
				panic(fmt.Errorf("invalid line: %q", string(l)))
			}
			return [2]RPS{parseRPS(l[0]), parseRPS(l[2])}
		},
	)
	return parsed
}

type round struct {
	them, me RPS
	o        Outcome
	score    int
}

func play1(themMe [2]RPS) round {
	ret := round{them: themMe[0], me: themMe[1]}
	ret.o = ret.me.vs(ret.them)
	ret.score = ret.me.score() + ret.o.score()
	return ret
}
