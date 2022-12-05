package day05

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
)

type state struct {
	stacks []stack
}

// stack goes bottom to top
type stack []rune

func parseState(in *bufio.Reader) state {
	var s state
	for {
		l, err := in.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		// if we get the column numbers line, we are done
		if l[0] == ' ' && unicode.IsDigit(rune(l[1])) {
			break
		}
		// each column is `[x] ` or `[x]` for the last one
		i.For(i.Chunk[rune](i.Runes(l), 4), func(c []rune, i int) {
			if c[0] == ' ' {
				// assume the rest is whitespace
			} else if c[0] == '[' {
				// an entry
				for len(s.stacks) <= i {
					s.stacks = append(s.stacks, nil)
				}
				s.stacks[i] = append(s.stacks[i], c[1])
			} else {
				panic(fmt.Errorf("bad column %q", string(c)))
			}
		})
	}
	// stacks are now populated, but in reverse order, correct them
	for _, st := range s.stacks {
		u.Reverse(st)
	}
	return s
}

func (s state) Move(m move) state {
	if m.count == 0 {
		return s
	}
	src := s.stacks[m.source]
	dst := s.stacks[m.dest]
	// unused cap in dst might be used elsewhere, make sure we don't change it
	dst = dst[0:len(dst):len(dst)]
	for i := 0; i < m.count; i++ {
		v := src[len(src)-1]
		src = src[0 : len(src)-1]
		dst = append(dst, v)
	}
	stacks := make([]stack, len(s.stacks))
	copy(stacks, s.stacks)
	stacks[m.source] = src
	stacks[m.dest] = dst
	return state{stacks}
}

func (s state) Move2(m move) state {
	if m.count == 0 {
		return s
	}
	src := s.stacks[m.source]
	dst := s.stacks[m.dest]
	// unused cap in dst might be used elsewhere, make sure we don't change it
	dst = dst[0:len(dst):len(dst)]
	dst = append(dst, src[len(src)-m.count:]...)
	src = src[0 : len(src)-m.count]
	stacks := make([]stack, len(s.stacks))
	copy(stacks, s.stacks)
	stacks[m.source] = src
	stacks[m.dest] = dst
	return state{stacks}
}

func (s state) Tops() []rune {
	ret := make([]rune, len(s.stacks))
	for i, s := range s.stacks {
		ret[i] = s[len(s)-1]
	}
	return ret
}
