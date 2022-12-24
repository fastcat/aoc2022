package day21

import "fmt"

type op func(a, b int) int

type node struct {
	name       string
	value      int
	resolved   bool
	op         op
	childNames [2]string
	children   [2]*node
}

func (n *node) Value() int {
	if !n.resolved {
		n.value, n.resolved = n.op(n.children[0].Value(), n.children[1].Value()), true
	}
	return n.value
}

var ops = map[rune]op{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int {
		if a%b != 0 {
			panic(fmt.Errorf("not int div %d/%d", a, b))
		}
		return a / b
	},
}
