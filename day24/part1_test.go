package day24

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// //go:embed sample.txt
// var sample string

//go:embed sample2.txt
var sample2 string

func logState(w io.Writer, s *state) {
	fmt.Fprintf(w, "at minute %d:\n%s\n%s\n", s.minute, s.boardString(), s.reachableString())
}

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	b := parse(sample2)
	a.Equal(pos{4, 6}, b.dims)
	a.Equal(pos{-1, 0}, b.initialPos())
	tp := b.targetPos()
	a.Equal(pos{4, 5}, tp)
	var s *state
	var buf strings.Builder
	for s = b.initialState(); !s.reachable[tp]; s = s.next() {
		logState(&buf, s)
	}
	logState(&buf, s)
	ok := a.Equal(18, s.minute)
	ok = ok && a.Contains(s.reachable, tp)
	ok = ok && a.True(s.reachable[tp])
	if !ok {
		t.Log(buf.String())
	}
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	b := parse(input)
	tp := b.targetPos()
	var s *state
	var buf strings.Builder
	for s = b.initialState(); !s.reachable[tp]; s = s.next() {
		fmt.Fprintf(&buf, "at minute %d:\n%s\n%s\n", s.minute, s.boardString(), s.reachableString())
	}
	fmt.Fprintf(&buf, "at minute %d:\n%s\n%s", s.minute, s.boardString(), s.reachableString())
	t.Log(s.minute)
	// t.Log(buf.String())
}
