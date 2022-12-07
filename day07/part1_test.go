package day07

import (
	_ "embed"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	expected := newDir("/")
	{
		a := newDir("a")
		expected.add(a)
		expected.add(newFile("b.txt", 14848514))
		expected.add(newFile("c.dat", 8504156))
		d := newDir("d")
		expected.add(d)

		e := newDir("e")
		a.add(e)
		a.add(newFile("f", 29116))
		a.add(newFile("g", 2557))
		a.add(newFile("h.lst", 62596))

		e.add(newFile("i", 584))

		d.add(newFile("j", 4060174))
		d.add(newFile("d.log", 8033020))
		d.add(newFile("d.ext", 5626152))
		d.add(newFile("k", 7214296))
	}
	a.EqualValues(48381165, expected.size())
	a.Equal(
		[]string{"i", "f", "g", "h.lst", "b.txt", "c.dat", "j", "d.log", "d.ext", "k"},
		i.ToSlice(i.Map(i.Funcer(expected.recursiveFiles), i.NoIndex((*file).name))),
	)
	a.Equal(
		[]string{"/", "a", "e", "d"},
		i.ToSlice(i.Map(i.Funcer(expected.recursiveDirs), i.NoIndex((*dir).name))),
	)

	root := parse(sample)
	a.EqualValues(48381165, root.size())
	a.Equal(expected, root)

	smallDirs := i.Filter(
		i.Funcer(root.recursiveDirs),
		func(d *dir) bool { return d.size() <= 100_000 },
	)
	a.Equal([]string{"a", "e"}, i.ToSlice(i.Map(smallDirs, i.NoIndex((*dir).name))))
	smallDirsSize := i.Sum(i.Map(smallDirs, i.NoIndex((*dir).size)))
	a.EqualValues(95437, smallDirsSize)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	root := parse(input)
	smallDirs := i.Filter(
		i.Funcer(root.recursiveDirs),
		func(d *dir) bool { return d.size() <= 100_000 },
	)
	smallDirsSize := i.Sum(i.Map(smallDirs, i.NoIndex((*dir).size)))
	t.Log(smallDirsSize)
}
