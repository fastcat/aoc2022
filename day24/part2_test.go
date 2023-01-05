package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	b := parse(sample2)
	var s *state
	ip := b.initialPos()
	tp := b.targetPos()
	for s = b.initialState(); !s.reachable[tp]; s = s.next() {
	}
	a.Equal(18, s.minute)
	for s = b.customState(tp, s.minute); !s.reachable[ip]; s = s.next() {
	}
	a.Equal(18+23, s.minute)
	for s = b.customState(ip, s.minute); !s.reachable[tp]; s = s.next() {
	}
	a.Equal(18+23+13, s.minute)
}

func TestPart2(t *testing.T) {
	b := parse(input)
	var s *state
	ip := b.initialPos()
	tp := b.targetPos()
	for s = b.initialState(); !s.reachable[tp]; s = s.next() {
	}
	t.Log(s.minute)
	for s = b.customState(tp, s.minute); !s.reachable[ip]; s = s.next() {
	}
	t.Log(s.minute)
	for s = b.customState(ip, s.minute); !s.reachable[tp]; s = s.next() {
	}
	t.Log(s.minute)
}
