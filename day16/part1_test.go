package day16

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed sample.txt
var sample string

func TestPart1Sample(t *testing.T) {
	a := assert.New(t)
	n := parseNet(sample)
	a.Len(n, 10)
	a.Equal("BB", n[1].name)
	a.Equal(21, n[9].rate)
	a.Len(n[3].conns, 3)
	a.Equal([]string{"CC", "AA", "EE"}, n[3].conns)
	a.Equal([]int{2, 0, 4}, n[3].connids)

	r := n.reduce()
	t.Log(r)

	end, path := r.search(30)
	reportPath(t, r, end, path)
}

func reportPath(t *testing.T, r *reducedNet, end key, path []best) {
	for i := 1; i < len(path); i++ {
		var k key
		if i == len(path)-1 {
			k = end
		} else {
			k = path[i+1].prev
		}
		t.Logf("minute %d pos %s rate %d total %d", k.min, r.names[k.pos], path[i].rate, path[i].total)
	}
}

//go:embed input.txt
var input string

func TestPart1(t *testing.T) {
	n := parseNet(input)
	r := n.reduce()
	t.Log(r)
	end, path := r.search(30)
	reportPath(t, r, end, path)
}
