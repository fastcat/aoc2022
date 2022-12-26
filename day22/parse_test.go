package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseMoves(t *testing.T) {
	tests := []struct {
		in   string
		want []move
	}{
		{
			"10RL11L12R",
			[]move{10, turnRight, turnLeft, 11, turnLeft, 12, turnRight},
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.want, parseMoves(tt.in))
		})
	}
}
