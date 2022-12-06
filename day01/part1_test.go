package day01

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
	"github.com/stretchr/testify/require"
)

//go:embed part1-sample.txt
var sample []byte

//go:embed part1-input.txt
var input []byte

func TestPart1Sample(t *testing.T) {
	r := require.New(t)
	parsed := parse(sample)
	r.Equal(
		[][]int{
			{1000, 2000, 3000},
			{4000},
			{5000, 6000},
			{7000, 8000, 9000},
			{10000},
		},
		i.ToSlice(parsed),
	)
	e := elves(parsed)
	m := i.Max(e)
	r.Equal(24000, m)
	// r.Equal(3, mi)
}

func TestPart1(t *testing.T) {
	parsed := parse(input)
	e := elves(parsed)
	m := i.Max(e)
	t.Log(m)
}

func parse(data []byte) i.Iterable[[]int] {
	d := i.Slice(data)
	spl := i.Split(d, []byte{'\n', '\n'})
	parsed := i.Map(spl, func(in []byte, _ int) []int {
		lines := i.ToStrings(i.Split(i.Slice(in), []byte{'\n'}))
		ints := i.Map(lines, i.NoIndex(i.Muster(strconv.Atoi)))
		return i.ToSlice(ints)
	})
	return parsed
}

func elves(in i.Iterable[[]int]) i.Iterable[int] {
	return i.Map(in, i.NoIndex(u.Sum[int]))
}
