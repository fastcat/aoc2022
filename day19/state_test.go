package day19

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_state_timeToBuild(t *testing.T) {
	tests := []struct {
		s    state
		c    cost
		want int
	}{
		{
			state{},
			cost{1, 0, 0},
			math.MaxUint8,
		},
		{
			state{inventory{1}, inventory{}},
			cost{1, 0, 0},
			0,
		},
		{
			state{inventory{0}, inventory{1}},
			cost{1, 0, 0},
			1,
		},
		{
			state{inventory{2}, inventory{}},
			cost{1, 0, 0},
			0,
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.s.timeToBuild(tt.c))
		})
	}
}
