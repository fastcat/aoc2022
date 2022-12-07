package day06

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed samples.txt
var samples string

func TestPart1Sample(t *testing.T) {
	lines := i.ToSlice(i.Split[rune](i.Runes(samples), []rune{'\n'}))
	expected := []int{7, 5, 6, 10, 11}
	for idx, l := range lines {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			it := i.Slice(l).Iterator()
			p := findMarker(it, 4)
			assert.Equal(t, expected[idx], p)
		})
	}
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	it := i.Runes(strings.TrimSpace(input)).Iterator()
	p := findMarker(it, 4)
	t.Log(p)
}
