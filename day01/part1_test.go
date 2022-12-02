package day01

import (
	"bytes"
	_ "embed"
	"strconv"
	"testing"

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
		parsed,
	)
	e := elves(parsed)
	m, mi := u.Max(e)
	r.Equal(24000, m)
	r.Equal(3, mi)
}

func TestPart1(t *testing.T) {
	parsed := parse(input)
	e := elves(parsed)
	m, mi := u.Max(e)
	t.Log(m, mi)
}

func parse(data []byte) [][]int {
	return u.Map(
		bytes.Split(data, []byte{'\n', '\n'}),
		func(s []byte) []int {
			return u.Map(
				bytes.Split(bytes.TrimRight(s, "\n"), []byte{'\n'}),
				func(l []byte) int {
					v, err := strconv.Atoi(string(l))
					u.PanicIf(err)
					return v
				},
			)
		},
	)
}

func elves(in [][]int) []int {
	return u.Map(in, u.Sum[int])
}
