package day11

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fastcat/aoc2022/i"
	"github.com/fastcat/aoc2022/u"
)

type monkey struct {
	items     *u.Circular[int]
	op, post  func(int) int
	testDiv   int
	targets   [2]int // true, false AKA %testDiv==0, else
	inspected int
}

func (m *monkey) step(g game) {
	if m.items.Len() == 0 {
		return
	}
	m.inspected++
	item := m.items.Pop()
	item = m.op(item)
	item = m.post(item)
	if item%m.testDiv == 0 {
		g[m.targets[0]].receive(item)
	} else {
		g[m.targets[1]].receive(item)
	}
}
func (m *monkey) turn(g game) {
	for m.items.Len() != 0 {
		m.step(g)
	}
}

func (m *monkey) receive(item int) {
	if m.items.Full() {
		m.items.GrowCap(m.items.Cap() * 2)
	}
	m.items.Push(item)
}

type game []*monkey

func (g game) setPost(post func(int) int) {
	for _, m := range g {
		m.post = post
	}
}
func (g game) useGCD() {
	// lol not the gcd
	gcd := i.Reduce(
		i.Map(i.Slice(g), func(m *monkey, _ int) int { return m.testDiv }),
		1,
		func(p, d, _ int) int { return p * d },
	)
	g.setPost(func(i int) int {
		return i % gcd
	})
}

func (g game) round() {
	for _, m := range g {
		m.turn(g)
	}
}

func (g game) rounds(n int) {
	for ; n > 0; n-- {
		g.round()
	}
}

func (g game) inspections() []int {
	return i.ToSlice(i.Map(i.Slice(g), func(m *monkey, _ int) int { return m.inspected }))
}

func (g game) business() int {
	if len(g) < 2 {
		panic(errors.New("game too small"))
	}
	s := i.Top(i.Map(i.Slice(g), func(m *monkey, _ int) int { return m.inspected }), 2)
	return s[0] * s[1]
}

func parseGame(in string, post func(int) int) game {
	var g game
	// i.Split discards blank lines
	i.For(i.Chunk(i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})), 6), func(stanza []string, idx int) {
		if len(stanza) != 6 {
			panic(errors.New("bad stanza"))
		}
		var mi int
		if _, err := fmt.Sscanf(stanza[0], "Monkey %d:\n", &mi); err != nil {
			panic(err)
		} else if mi != idx {
			panic(fmt.Errorf("got monkey %d in stanza %d", mi, idx))
		}
		var m monkey
		if !strings.HasPrefix(stanza[1], "  Starting items: ") {
			panic(fmt.Errorf("monkey %d bad stanza second line prefix", idx))
		}
		m.items = u.CircularFrom(
			i.ToSlice(
				i.Map(
					i.Split(
						i.Runes(strings.TrimPrefix(stanza[1], "  Starting items: ")),
						[]rune{',', ' '},
					),
					func(s []rune, _ int) int {
						if n, err := strconv.Atoi(string(s)); err != nil {
							panic(err)
						} else {
							return n
						}
					},
				),
			),
		)
		var op rune
		var operand string
		if _, err := fmt.Sscanf(stanza[2], "  Operation: new = old %c %s\n", &op, &operand); err != nil {
			panic(err)
		}
		m.op = makeOp(op, operand)
		m.post = post
		if _, err := fmt.Sscanf(stanza[3], "  Test: divisible by %d\n", &m.testDiv); err != nil {
			panic(err)
		}
		if _, err := fmt.Sscanf(stanza[4], "    If true: throw to monkey %d", &m.targets[0]); err != nil {
			panic(err)
		}
		if _, err := fmt.Sscanf(stanza[5], "    If false: throw to monkey %d", &m.targets[1]); err != nil {
			panic(err)
		}
		g = append(g, &m)
	})
	for idx, m := range g {
		if m.targets[0] >= len(g) || m.targets[1] >= len(g) {
			panic(fmt.Errorf("monkey %d bad target(s) %v", idx, m.targets))
		}
	}
	return g
}

func makeOp(op rune, operand string) func(int) int {
	var opVal func(v int) int
	if operand == "old" {
		opVal = ident
	} else if opi, err := strconv.Atoi(operand); err != nil {
		panic(err)
	} else {
		opVal = func(int) int { return opi }
	}
	switch op {
	case '*':
		return func(v int) int { return v * opVal(v) }
	case '+':
		return func(v int) int { return v + opVal(v) }
	default:
		panic(fmt.Errorf("unsupported op '%c'", op))
	}
}

func div3(v int) int  { return v / 3 }
func ident(v int) int { return v }
