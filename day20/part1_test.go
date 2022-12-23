package day20

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	m := parse(sample)
	a.Equal(
		&mixer{
			orig: []int{1, 2, -3, 3, -2, 0, 4},
			pos:  []int{0, 1, 2, 3, 4, 5, 6},
		},
		m,
	)

	a.Equal([]int{1, 2, -3, 3, -2, 0, 4}, m.All())
	m.mixOne(0)
	a.Equal([]int{2, 1, -3, 3, -2, 0, 4}, m.All())
	m.mixOne(1)
	a.Equal([]int{1, -3, 2, 3, -2, 0, 4}, m.All())
	m.mixOne(2)
	a.Equal([]int{1, 2, 3, -2, -3, 0, 4}, m.All())
	m.mixOne(3)
	a.Equal([]int{1, 2, -2, -3, 0, 3, 4}, m.All())
	m.mixOne(4)
	a.Equal([]int{1, 2, -3, 0, 3, 4, -2}, m.All())
	m.mixOne(5)
	a.Equal([]int{1, 2, -3, 0, 3, 4, -2}, m.All())
	m.mixOne(6)
	a.Equal([]int{1, 2, -3, 4, 0, 3, -2}, m.All())

	m.reset()
	a.Equal([]int{1, 2, -3, 3, -2, 0, 4}, m.All())
	m.mixAll()
	a.Equal([]int{1, 2, -3, 4, 0, 3, -2}, m.All())

	p := m.IndexOf(0)
	a.Equal(4, p)
	a.Equal(4, m.At(p+1000))
	a.Equal(-3, m.At(p+2000))
	a.Equal(2, m.At(p+3000))
	a.Equal(3, m.OffsetSum(0, 1000, 2000, 3000))
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	m := parse(input)
	m.mixAll()
	t.Log(m.OffsetSum(0, 1000, 2000, 3000))
}
