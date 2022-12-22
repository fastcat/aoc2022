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
			oreBot:      cost{4, 0, 0},
			clayBot:     cost{2, 0, 0},
			obsidianBot: cost{3, 14, 0},
			geodeBot:    cost{2, 0, 7},
		},
		{
			oreBot:      cost{2, 0, 0},
			clayBot:     cost{3, 0, 0},
			obsidianBot: cost{3, 8, 0},
			geodeBot:    cost{3, 0, 12},
		},
	})

	s := initialState()
	for s.minute < 24 {
		s = s.playSimple(bps[0])
		t.Log(s)
	}
	t.Log(s)
	a.Equal(9, s.inv.geodes)
	s = initialState()
	for s.minute < 24 {
		s = s.playSimple(bps[1])
		t.Log(s)
	}
	t.Log(s)
	a.Equal(12, s.inv.geodes)
}
