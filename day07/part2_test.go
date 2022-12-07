package day07

import (
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

const totalSpace = 70000000
const minFree = 30000000

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	root := parse(sample)
	best := minToDelete(root)
	a.Equal("d", best.name())
}

func TestPart2(t *testing.T) {
	root := parse(input)
	best := minToDelete(root)
	t.Log(best.name(), best.size())
}

func minToDelete(root *dir) *dir {
	free := totalSpace - root.size()
	needed := minFree - free
	if needed <= 0 {
		return nil
	}
	candidates := i.Filter(
		i.Funcer(root.recursiveDirs),
		func(d *dir) bool { return d.size() >= needed },
	)
	best, _ := i.MinBy(candidates, (*dir).size)
	return best
}
