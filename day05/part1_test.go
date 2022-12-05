package day05

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/fastcat/aoc2022/u"
	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	state, moves := parseStateAndMoves(sample)
	a.Equal(
		[]stack{
			{'Z', 'N'},
			{'M', 'C', 'D'},
			{'P'},
		},
		state.stacks,
	)
	a.Len(moves, 4)
	a.Equal(moves[0], move{1, 0, 1})
	for _, m := range moves {
		state = state.Move(m)
	}
	a.Equal(
		[]stack{
			{'C'},
			{'M'},
			{'P', 'D', 'N', 'Z'},
		},
		state.stacks,
	)
	a.Equal(
		[]rune{'C', 'M', 'Z'},
		state.Tops(),
	)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	state, moves := parseStateAndMoves(input)
	for _, m := range moves {
		state = state.Move(m)
	}
	t.Log(string(state.Tops()))
}

func parseStateAndMoves(in string) (state, []move) {
	r := strings.NewReader(in)
	br := bufio.NewReader(r)
	s := parseState(br)
	l, err := br.ReadString('\n')
	u.PanicIf(err)
	if l != "\n" {
		panic(fmt.Errorf("unexpected line %q", l))
	}
	var m []move
	for {
		l, err := br.ReadString('\n')
		if errors.Is(err, io.EOF) {
			break
		}
		m = append(m, parseMove(l))
	}
	return s, m
}
