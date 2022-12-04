package day03

import (
	_ "embed"
	"errors"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)

	a.Equal(27, priority('A'))
	a.Equal(28, priority('B'))
	a.Equal('A', unPriority(27))
	a.Equal('B', unPriority(28))

	bags := parse(sample)
	a.Equal(
		[][2]string{
			{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			{"PmmdzqPrV", "vPwwTWBwg"},
			{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
			{"ttgJtRGJ", "QctTZtZT"},
			{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
		},
		i.ToSlice(i.Map(bags, func(in [2][]rune) [2]string { return [2]string{string(in[0]), string(in[1])} })),
	)
	dupes := i.Map(bags, findDupe)
	dupeRunes := i.Map(dupes, unPriority)
	a.Equal(
		[]rune{'p', 'L', 'P', 'v', 't', 's'},
		i.ToSlice(dupeRunes),
	)
	result := int(i.Sum(dupes))
	a.Equal(157, result)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	bags := parse(input)
	dupes := i.Map(bags, findDupe)
	result := int(i.Sum(dupes))
	t.Log(result)
}

func parse(in string) i.Iterable[[2][]rune] {
	r := i.Runes(in)
	l := i.Split[rune](r, []rune{'\n'})
	p := i.Map(l, func(l []rune) [2][]rune {
		return [2][]rune{
			l[0 : len(l)/2],
			l[len(l)/2:],
		}
	})
	return p
}

func findDupe(bag [2][]rune) int {
	set1 := prioritySet(bag[0])
	set2 := prioritySet(bag[1])
	dupes := set1 & set2
	return singleBit(dupes)
}

func singleBit(in int64) int {
	ret := 0
	for i := 1; i <= 52; i++ {
		if in&(1<<i) != 0 {
			if ret != 0 {
				panic(errors.New("two dupes"))
			}
			ret = i
		}
	}
	if ret > 0 {
		return ret
	}
	panic(errors.New("no dupe"))
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return 1 + int(r-'a')
	} else if r >= 'A' && r <= 'Z' {
		return 27 + int(r-'A')
	}
	panic(errors.New("invalid"))
}
func unPriority(p int) rune {
	if p >= 1 && p <= 26 {
		return 'a' + rune(p-1)
	} else if p >= 27 && p <= 52 {
		return 'A' + rune(p-27)
	}
	panic(errors.New("invalid"))
}

func prioritySet(in []rune) int64 {
	mask := int64(0)
	for _, r := range in {
		mask |= 1 << priority(r)
	}
	return mask
}
