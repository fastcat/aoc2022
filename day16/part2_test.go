package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	n := parseNet(sample)
	r := n.reduce()
	w := r.walker(26)
	w.walk()
	k1, k2, bt := w.bestPair()
	t.Log(k1, k2, bt)
	a.Equal(1707, bt)
	reportPath(t, r, k1, w.pathFrom(k1))
	reportPath(t, r, k2, w.pathFrom(k2))
}

func TestPart2(t *testing.T) {
	n := parseNet(input)
	r := n.reduce()
	w := r.walker(26)
	w.walk()
	k1, k2, bt := w.bestPair()
	t.Log(k1, k2, bt)
	require.NotZero(t, bt)
	reportPath(t, r, k1, w.pathFrom(k1))
	reportPath(t, r, k2, w.pathFrom(k2))
}
