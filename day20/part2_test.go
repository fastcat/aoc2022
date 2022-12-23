package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const magicMultiplier = 811589153

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	m := parse(sample)
	m.Mult(magicMultiplier)

	a.Equal([]int{811589153, 1623178306, -2434767459, 2434767459, -1623178306, 0, 3246356612}, m.All())
	m.mixAll()
	a.Equal([]int{0, -2434767459, 3246356612, -1623178306, 2434767459, 1623178306, 811589153}, m.All())
	m.mixAll()
	a.Equal([]int{0, 2434767459, 1623178306, 3246356612, -2434767459, -1623178306, 811589153}, m.All())
	m.mixAll()
	a.Equal([]int{0, 811589153, 2434767459, 3246356612, 1623178306, -1623178306, -2434767459}, m.All())
	//...

	m.reset()
	m.mixAllRepeat(10)
	a.Equal([]int{0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459, 811589153}, m.All())
	a.Equal(1623178306, m.OffsetSum(0, 1000, 2000, 3000))
}

func TestPart2(t *testing.T) {
	m := parse(input)
	m.Mult(magicMultiplier)
	m.mixAllRepeat(10)
	t.Log(m.OffsetSum(0, 1000, 2000, 3000))
}
