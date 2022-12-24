package day21

import (
	"fmt"
	"unicode"

	"github.com/fastcat/aoc2022/i"
)

// parse returns the root tree
func parse(in string) *node {
	nodes :=
		i.Map(
			i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
			i.NoIndex(parseOne),
		)
	byName := i.ValuesToMap(nodes, func(n *node) string { return n.name })

	root := byName["root"]
	if root == nil {
		panic(fmt.Errorf("no 'root' node"))
	}

	resolveChildren(byName)

	return root
}

func parseOne(in string) *node {
	if len(in) < 6 {
		panic(fmt.Errorf("input too short: %q", in))
	}
	n := node{name: in[:4]}
	if unicode.IsDigit(rune(in[6])) {
		if _, err := fmt.Sscanf(in[6:], "%d\n", &n.value); err != nil {
			panic(err)
		}
		n.resolved = true
	} else {
		var op rune
		if _, err := fmt.Sscanf(in[6:], "%4s %c %4s\n", &n.childNames[0], &op, &n.childNames[1]); err != nil {
			panic(err)
		}
		opf := ops[op]
		if opf == nil {
			panic(fmt.Errorf("invalid op '%c'", op))
		}
		n.op = opf
	}
	return &n
}

func resolveChildren(nodes map[string]*node) {
	for _, n := range nodes {
		for i, cn := range n.childNames {
			if cn == "" {
				continue
			} else if cnn := nodes[cn]; cnn == nil {
				panic(fmt.Errorf("node %s children[%d,%s] not found", n.name, i, cn))
			} else {
				n.children[i] = cnn
			}
		}
	}
}
