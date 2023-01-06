package day25

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
	sum := i.Sum(i.Map(i.Lines(sample), i.NoIndex(decode)))
	a.Equal(4890, sum)
	e := encode(sum)
	a.Equal("2=-1=0", e)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	sum := i.Sum(i.Map(i.Lines(input), i.NoIndex(decode)))
	t.Log(sum, encode(sum))
}
