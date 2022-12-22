package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	bps := parseMany(sample)
	best := searchMany(bps, 32)
	a.EqualValues([]uint8{56, 62}, best)
}
func TestPart2(t *testing.T) {
	bps := parseMany(input)[:3]
	best := searchMany(bps, 32)
	t.Log(qualityProduct(best))
}
