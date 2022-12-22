package day19

import (
	_ "embed"
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

	g1 := graph{bps[0], 24}
	b1 := g1.search()
	t.Log(b1)
	a.EqualValues(9, b1.inv[geode])

	g2 := graph{bps[1], 24}
	b2 := g2.search()
	t.Log(b2)
	a.EqualValues(12, b2.inv[geode])
}
