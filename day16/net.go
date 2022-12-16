package day16

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/fastcat/aoc2022/i"
	"golang.org/x/exp/slices"
)

type net []*node

type node struct {
	id      int
	name    string
	rate    int
	conns   []string
	connids []int
}

func parseNet(in string) net {
	net := net(i.ToSlice(
		i.Map(
			i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
			i.NoIndex(parseNode),
		),
	))
	// get AA first
	slices.SortFunc(net, func(a, b *node) bool { return a.name < b.name })
	if net[0].name != "AA" {
		panic(fmt.Errorf("no AA node, got %q instead", net[0].name))
	}
	ids := i.ValuesToMap(i.Range(0, len(net), 1), func(id int) string { return net[id].name })
	for id, n := range net {
		n.id = id
		n.connids = make([]int, len(n.conns))
		for j, c := range n.conns {
			if cid, ok := ids[c]; !ok {
				panic(fmt.Errorf("node %q has conn to invalid node %q", n.name, c))
			} else {
				n.connids[j] = cid
			}
		}
	}
	return net
}

func parseNode(in string) *node {
	var n node
	r := strings.NewReader(in)
	var junk string
	if _, err := fmt.Fscanf(r,
		"Valve %s has flow rate=%d; tunne%s lea%s to valv%s ",
		&n.name, &n.rate, &junk, &junk, &junk,
	); err != nil {
		panic(err)
	}
	for {
		var cn string
		if _, err := fmt.Fscanf(r, "%s", &cn); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		} else {
			cn = strings.TrimSpace(cn)
			cn = strings.TrimSuffix(cn, ",")
			n.conns = append(n.conns, cn)
		}
	}
	return &n
}

type reducedNet struct {
	rates     []int
	dists     [][]int
	zerodists []int
	names     []string
}

func (n net) reduce() *reducedNet {
	rates := i.ToSlice(i.Map(i.Slice(n), func(nn *node, _ int) int { return nn.rate }))
	if rates[0] != 0 {
		panic(fmt.Errorf("reduced net requires AA node to have rate=0"))
	}

	// compute the full pair-wise distances
	dists := make([][]int, len(n))
	// initialize the distances to infinite
	for i := 0; i < len(n); i++ {
		dists[i] = make([]int, len(n))
		for j := 0; j < len(n); j++ {
			dists[i][j] = math.MaxInt
		}
	}
	// walk the connections to compute the distances
	var walk func(start, cur *node, bd int)
	walk = func(start, cur *node, bd int) {
		d := dists[start.id]
		d[start.id] = 0
		cd := bd + 1
		for _, cn := range cur.connids {
			if cd < d[cn] {
				d[cn] = cd
				walk(start, n[cn], cd)
			}
		}
	}
	for _, nn := range n {
		walk(nn, nn, 0)
	}

	// keep only nodes with non-zero rates
	// build a map of new to old index
	m := make([]int, 0, len(n))
	for i := 0; i < len(n); i++ {
		if rates[i] != 0 {
			j := len(m)
			m = append(m, i)
			// move values down
			rates[j] = rates[i]
		}
	}
	rates = rates[:len(m)]
	// collapse the distances matrix
	zerodists := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		zerodists[i] = dists[0][m[i]]
	}
	for i := 0; i < len(m); i++ {
		nd, od := dists[i], dists[m[i]]
		for j := 0; j < len(m); j++ {
			nd[j] = od[m[j]]
		}
		dists[i] = nd[:len(m)]
	}
	dists = dists[:len(m)]
	names := i.ToSlice(i.Map(i.Slice(m), func(oi, _ int) string { return n[oi].name }))

	return &reducedNet{rates, dists, zerodists, names}
}

func (rn *reducedNet) flow(mask int64) int {
	f := 0
	for i, j := 0, int64(1); i < len(rn.rates); i, j = i+1, j<<1 {
		if mask&j != 0 {
			f += rn.rates[i]
		}
	}
	return f
}
