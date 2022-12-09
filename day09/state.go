package day09

import (
	"errors"
	"fmt"
	"io"
)

type pos [2]int

func (p pos) move(dir rune) pos {
	switch dir {
	case 'R':
		p[0]++
	case 'L':
		p[0]--
	case 'U':
		p[1]++
	case 'D':
		p[1]--
	default:
		panic(fmt.Errorf("bad dir %c", dir))
	}
	return p
}

func (p pos) follow(head pos) pos {
	diff := [2]int{
		head[0] - p[0],
		head[1] - p[1],
	}
	if diff[0] >= -1 && diff[0] <= 1 && diff[1] >= -1 && diff[1] <= 1 {
		// overlapping, adjacent, or diagonal => no move
		return p
	}
	if diff[0] == 0 {
		// same x, move 1 unit in y towards head
		p[1] += sign(diff[1])
	} else if diff[1] == 0 {
		// same y, move 1 unit in x towards head
		p[0] += sign(diff[0])
	} else {
		// move diagonally 1 unit towards head
		p[0] += sign(diff[0])
		p[1] += sign(diff[1])
	}
	return p
}

func sign(i int) int {
	switch {
	case i < 0:
		return -1
	case i > 0:
		return 1
	default:
		return 0
	}
}

type move struct {
	dir   rune
	count int
}

func parseMoveList(in io.Reader) []move {
	var ret []move
	for {
		var m move
		if _, err := fmt.Fscanf(in, "%c %d\n", &m.dir, &m.count); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		ret = append(ret, m)
	}
	return ret
}

type state struct {
	head        pos
	tails       []pos
	tailVisited map[pos]bool
}

func NewState(numTails int) *state {
	return &state{
		tails:       make([]pos, numTails),
		tailVisited: map[pos]bool{{}: true},
	}
}

func (s *state) apply(moves []move) {
	if len(s.tails) == 0 || s.tailVisited == nil {
		panic(errors.New("uninitialized state"))
	}
	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			s.head = s.head.move(m.dir)
			h := s.head
			for i, tail := range s.tails {
				tail = tail.follow(h)
				s.tails[i] = tail
				h = tail
			}
			// only track the final tail's visits
			s.tailVisited[h] = true
		}
	}
}
