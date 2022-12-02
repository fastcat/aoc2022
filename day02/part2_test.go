package day02

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/require"
)

func TestPart2Sample(t *testing.T) {
	r := require.New(t)
	results := i.Map(parse2(sample), play2)
	r.Equal(
		[]round{
			{Rock, Rock, Draw, 4},
			{Paper, Rock, Loss, 1},
			{Scissors, Rock, Win, 7},
		},
		i.ToSlice(results),
	)
	total := i.Sum(i.Map(results, func(r round) int { return r.score }))
	r.Equal(12, total)
}

func TestPart2(t *testing.T) {
	results := i.Map(parse2(input), play2)
	total := i.Sum(i.Map(results, func(r round) int { return r.score }))
	t.Log(total)
}

func parse2(in string) i.Iterable[round] {
	return i.Map(
		i.Split[rune](i.Runes(in), []rune{'\n'}),
		func(l []rune) round {
			return round{
				them: parseRPS(rune(l[0])),
				o:    parseOutcome(rune(l[2])),
			}
		},
	)
}

func play2(r round) round {
	r.me = r.them.rev(r.o)
	r.score = r.me.score() + r.o.score()
	return r
}
