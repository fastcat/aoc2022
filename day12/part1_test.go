package day12

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b := parseBoard(sample)
	a.Len(b.elevations, 5)
	a.Equal([]int{0, 1, 3, 4, 5, 6, 7, 8}, b.elevations[4])
	a.Equal([2]int{0, 0}, b.start)
	a.Equal([2]int{2, 5}, b.end)

	p := newPather(b)
	p.search()
	if false {
		p.pretty(func(l []rune) { t.Log(string(l)) })
	}
	// var pretty []string
	// p.pretty(func(l []rune) { pretty = append(pretty, string(l)) })
	// a.Equal(
	// 	[]string{
	// 		"v..v<<<<",
	// 		">v.vv<<^",
	// 		".>vv>E^^",
	// 		"..v>>>^^",
	// 		"..>>>>>^",
	// 	},
	// 	pretty,
	// )
	a.Equal(31, p.distFromStart())
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	b := parseBoard(input)
	p := newPather(b)
	p.search()
	if false {
		p.pretty(func(l []rune) { t.Log(string(l)) })
	}
	t.Log(p.distFromStart())
}
