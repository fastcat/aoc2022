package day01

import (
	"testing"

	"github.com/fastcat/aoc2022/u"
	"github.com/stretchr/testify/require"
)

func TestPart2Sample(t *testing.T) {
	r := require.New(t)
	parsed := parse(sample)
	e := elves(parsed)
	top3 := u.Top(e, 3)
	r.Equal([]int{24000, 11000, 10000}, top3)
}

func TestPart2(t *testing.T) {
	parsed := parse(input)
	e := elves(parsed)
	top3 := u.Top(e, 3)
	sum := u.Sum(top3)
	t.Log(sum)
}
