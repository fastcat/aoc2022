package day19

import (
	_ "embed"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	bps := parseMany(sample)
	a.Len(bps, 2)
	a.Equal(bps, []*blueprint{
		{
			{4, 0, 0},
			{2, 0, 0},
			{3, 14, 0},
			{2, 0, 7},
		},
		{
			{2, 0, 0},
			{3, 0, 0},
			{3, 8, 0},
			{3, 0, 12},
		},
	})

	best := searchMany(bps)

	a.EqualValues([]uint8{9, 12}, best)

	qs := qualitySum(best)
	a.EqualValues(33, qs)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	bps := parseMany(input)
	best := make([]uint8, len(bps))
	wg := sync.WaitGroup{}
	wg.Add(len(bps))
	fb := func(i int, b *blueprint) {
		defer wg.Done()
		g := graph{b, 24}
		best[i] = g.search()
	}
	for i, b := range bps {
		go fb(i, b)
	}
	wg.Wait()
	qs := qualitySum(best)
	t.Log(qs)
}
