package day18

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	scan := parseMany(sample)
	a.Len(scan, 13)
	faces := 6*len(scan) - 2*countSharedFaces(scan)
	a.Equal(64, faces)
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	scan := parseMany(input)
	faces := 6*len(scan) - 2*countSharedFaces(scan)
	t.Log(faces)
}
