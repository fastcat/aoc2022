package day21

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	r := parse(sample)
	a.Equal("root", r.name)
	v := r.Value()
	a.Equal(152, v)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	r := parse(input)
	v := r.Value()
	t.Log(v)
}
