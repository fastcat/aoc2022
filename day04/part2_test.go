package day04

import (
	"fmt"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	p := parse(sample)
	overlaps := i.Filter(p, func(p pair) bool {
		return p[0].Overlaps(p[1])
	})
	a.Equal(
		[]pair{
			{{5, 7}, {7, 9}},
			{{2, 8}, {3, 7}},
			{{6, 6}, {4, 6}},
			{{2, 6}, {4, 8}},
		},
		i.ToSlice(overlaps),
	)
	a.Equal(4, i.Count(overlaps))
}

func TestPart2(t *testing.T) {
	p := parse(input)
	overlaps := i.Filter(p, func(p pair) bool {
		return p[0].Overlaps(p[1])
	})
	t.Log(i.Count(overlaps))
}

func (r secrange) TestOverlaps(t *testing.T) {
	for _, tt := range []struct {
		p pair
		o bool
	}{
		{pair{{1, 1}, {1, 1}}, true},
		{pair{{1, 2}, {1, 1}}, true},
		{pair{{1, 3}, {2, 5}}, true},
		{pair{{1, 4}, {2, 3}}, true},
		{pair{{1, 2}, {3, 4}}, false},
	} {
		t.Run(fmt.Sprintf("%#v", tt.p), func(t *testing.T) {
			assert.Equal(t, tt.o, tt.p[0].Overlaps(tt.p[1]))
			assert.Equal(t, tt.o, tt.p[1].Overlaps(tt.p[0]))
		})
	}
}

func (r secrange) Overlaps(o secrange) bool {
	return r[1] >= o[0] && r[0] <= o[1]
}
