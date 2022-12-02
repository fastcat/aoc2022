package day02

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/fastcat/aoc2022/u"
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
	results := u.Map(parse1(sample), play1)
	r.Equal(
		[]round{
			{Rock, Paper, Win, 8},
			{Paper, Rock, Loss, 1},
			{Scissors, Scissors, Draw, 6},
		},
		results,
	)
	total := u.SumF(results, func(r round) int { return r.score })
	r.Equal(15, total)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	results := u.Map(parse1(input), play1)
	total := u.SumF(results, func(r round) int { return r.score })
	t.Log(total)
}

func parse1(in string) [][2]RPS {
	return u.Map(
		strings.Split(strings.TrimRight(in, "\n"), "\n"),
		func(l string) [2]RPS {
			return [2]RPS{parseRPS(rune(l[0])), parseRPS(rune(l[2]))}
		},
	)
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
