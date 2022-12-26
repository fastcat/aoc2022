package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	b, m := parse(sample, (*board).buildPortals2s)
	s := b.initialState()
	fs := b.moves(s, m...)
	a.Equal(state{4, 6, up}, fs)
	a.Equal(5031, fs.Value())
	// t.Log("\n" + b.traceString())
}

func TestPart2(t *testing.T) {
	b, m := parse(input, (*board).buildPortals2i)
	s := b.initialState()
	fs := b.moves(s, m...)
	t.Log(fs, fs.Value())
	// t.Log("\n" + b.traceString())
}
