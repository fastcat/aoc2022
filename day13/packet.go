package day13

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"github.com/fastcat/aoc2022/i"
	"golang.org/x/exp/slices"
)

type cmpRes int

const (
	less    cmpRes = -1
	equal   cmpRes = 0
	greater cmpRes = 1
)

type item interface {
	cmp(item) cmpRes
}

type numberItem int

func (n numberItem) cmp(rhs item) cmpRes {
	if rhs, ok := rhs.(numberItem); ok {
		switch {
		case n < rhs:
			return less
		case n > rhs:
			return greater
		default:
			return equal
		}
	}
	return listItem{n}.cmp(rhs)
}

type listItem []item

func (l listItem) cmp(rhs item) cmpRes {
	var rl listItem
	if _, ok := rhs.(numberItem); ok {
		rl = listItem{rhs}
	} else {
		rl = rhs.(listItem)
	}
	for i, li := range l {
		if i >= len(rl) {
			return greater
		} else if c := li.cmp(rl[i]); c != equal {
			return c
		}
	}
	if len(rl) > len(l) {
		return less
	}
	return equal
}

func parsePacket(in i.Iterable[rune]) listItem {
	it := i.Peeker(in.Iterator())
	l := parseList(it)
	r, done := it.Peek()
	if !done {
		panic(fmt.Errorf("got '%c', expecting eof", r))
	}
	return l
}

func parseList(in i.PeekIterator[rune]) listItem {
	r, done := in.Next()
	if done {
		panic(fmt.Errorf("got eof, expecting ["))
	} else if r != '[' {
		panic(fmt.Errorf("got '%c', expecting [", r))
	}
	l := make(listItem, 0)
	comma := false
	for {
		r, done = in.Peek()
		if done {
			if comma {
				panic(fmt.Errorf("got eof, expecting comma or ]"))
			}
			panic(fmt.Errorf("got eof, expecting item or ]"))
		} else if comma && r == ',' {
			in.Next() // consume the peek'd comma
			comma = false
		} else if r == ']' {
			in.Next() // consume the peek'd ]
			return l
		} else if comma {
			panic(fmt.Errorf("got '%c', expecting comma or ]", r))
		} else if r == '[' {
			l = append(l, parseList(in))
			comma = true
		} else if unicode.IsDigit(r) {
			l = append(l, parseNumber(in))
			comma = true
		} else {
			panic(fmt.Errorf("got '%c', expecting [, ], or digit", r))
		}
	}
}

func parseNumber(in i.PeekIterator[rune]) numberItem {
	s := make([]rune, 0, 2)
	for {
		r, done := in.Peek()
		if done {
			break
		} else if unicode.IsDigit(r) {
			in.Next() // consume what we peeked
			s = append(s, r)
		} else {
			break
		}
	}
	if len(s) == 0 {
		panic(fmt.Errorf("got eof, expecting digit"))
	} else if n, err := strconv.Atoi(string(s)); err != nil {
		panic(err)
	} else {
		return numberItem(n)
	}
}

func rightPairs(in i.Iterable[[]listItem]) i.Iterable[int] {
	return i.Filter(
		i.Map(in, func(p []listItem, idx int) int {
			if p[0].cmp(p[1]) == less {
				return idx + 1
			}
			return 0
		}),
		func(idx int) bool { return idx > 0 },
	)
}

func posOf(haystack []listItem, needle listItem) int {
	idx := slices.IndexFunc(
		haystack,
		func(candidate listItem) bool {
			return needle.cmp(candidate) == equal
		},
	)
	if idx < 0 {
		panic(errors.New("needle not found"))
	}
	return idx + 1
}
