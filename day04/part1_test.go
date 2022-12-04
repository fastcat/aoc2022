package day04

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample []byte

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	p := parse(sample)
	a.Len(i.ToSlice(p), 6)
	a.Equal(6, i.Count(p))
	overlaps := i.Filter(p, func(p pair) bool {
		return p[0].Contains(p[1]) || p[1].Contains(p[0])
	})
	a.Equal(2, i.Count(overlaps))
}

//go:embed input.txt
var input []byte

func TestPart1(t *testing.T) {
	p := parse(input)
	overlaps := i.Filter(p, func(p pair) bool {
		return p[0].Contains(p[1]) || p[1].Contains(p[0])
	})
	t.Log(i.Count(overlaps))
}

type secrange [2]int

type pair [2]secrange

func parse(in []byte) i.Iterable[pair] {
	return i.Funcer(func() i.Iterator[pair] {
		r := bytes.NewReader(in)
		return i.Func(func() (pair, bool) {
			var p pair
			if n, err := fmt.Fscanf(r, "%d-%d,%d-%d\n", &p[0][0], &p[0][1], &p[1][0], &p[1][1]); err != nil {
				if n == 0 && errors.Is(err, io.EOF) {
					return p, true
				}
				panic(err)
			}
			return p, false
		})
	})
}

func (r secrange) Contains(o secrange) bool {
	return r[0] <= o[0] && r[1] >= o[1]
}
