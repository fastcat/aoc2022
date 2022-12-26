package day22

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b, m := parse(sample)
	a.Equal(12, len(b.g))
	a.Equal(
		[]move{
			10, turnRight,
			5, turnLeft,
			5, turnRight,
			10, turnLeft,
			4, turnRight,
			5, turnLeft,
			5,
		},
		m,
	)
	a.Equal(b.portals[state{0, 8, left}], state{0, 11, left})
}
