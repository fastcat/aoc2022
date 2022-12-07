package day06

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	lines := i.ToSlice(i.Split[rune](i.Runes(samples), []rune{'\n'}))
	expected := []int{19, 23, 23, 29, 26}
	for idx, l := range lines {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			it := i.Slice(l).Iterator()
			p := findMarker(it, 14)
			assert.Equal(t, expected[idx], p)
		})
	}
}

func TestPart2(t *testing.T) {
	it := i.Runes(strings.TrimSpace(input)).Iterator()
	p := findMarker(it, 14)
	t.Log(p)
}
