package day21

import "fmt"

type op func(a, b int) int

type node struct {
	name       string
	value      int
	resolved   bool
	opName     rune
	op         op
	childNames [2]string
	children   [2]*node
	human      bool
}

func (n *node) Value() int {
	if !n.resolved {
		n.value, n.resolved = n.op(n.children[0].Value(), n.children[1].Value()), true
	}
	return n.value
}

func (n *node) FindHuman() bool {
	if n.human {
		return true
	} else if n.name == "humn" {
		n.human = true
	} else {
		for _, c := range n.children {
			if c != nil && c.FindHuman() {
				n.human = true
			}
		}
	}
	return n.human
}

func (n *node) InvertRootValue() int {
	if n.name != "root" {
		panic(fmt.Errorf("can't invert non-root node"))
	}
	if !n.FindHuman() {
		panic(fmt.Errorf("root node has no human child"))
	}
	value, human, _ := n.splitHuman()
	return human.InvertValue(value.Value())
}

func (n *node) splitHuman() (value, human *node, humanFirst bool) {
	if !n.human {
		panic(fmt.Errorf("can't split non-human node"))
	}
	if n.children[0].FindHuman() {
		return n.children[1], n.children[0], true
	} else {
		return n.children[0], n.children[1], false
	}
}

func (n *node) InvertValue(target int) int {
	if !n.human {
		panic(fmt.Errorf("can't invert non-human node"))
	} else if n.name == "humn" {
		return target
	}
	valueNode, humanNode, humanFirst := n.splitHuman()
	operand := valueNode.Value()
	humanTarget := invs[n.opName](target, operand, !humanFirst)
	return humanNode.InvertValue(humanTarget)
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

type opInv = func(target, operand int, operandFirst bool) int

var invs = map[rune]opInv{
	'+': func(target, operand int, _ bool) int { return target - operand },
	'-': func(target, operand int, operandFirst bool) int {
		if operandFirst {
			return operand - target
		} else {
			return operand + target
		}
	},
	'*': func(target, operand int, _ bool) int {
		if target%operand != 0 {
			panic(fmt.Errorf("not int div %d/%d", target, operand))
		}
		return target / operand
	},
	'/': func(target, operand int, operandFirst bool) int {
		if operandFirst {
			if operand%target != 0 {
				panic(fmt.Errorf("not int div %d/%d", operand, target))
			}
			return operand / target
		} else {
			return target * operand
		}
	},
}
