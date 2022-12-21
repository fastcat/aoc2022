package day17

import (
	"strings"
)

const rowWidth = 7

type row byte

func (r row) String() string {
	var buf strings.Builder
	buf.Grow(9)
	buf.WriteRune('|')
	for i := 0; i < rowWidth; i++ {
		if r&(1<<i) == 0 {
			buf.WriteRune('.')
		} else {
			buf.WriteRune('#')
		}
	}
	buf.WriteRune('|')
	return buf.String()
}

type board []row

func (b board) String() string {
	var buf strings.Builder
	for r := len(b) - 1; r >= 0; r-- {
		buf.WriteString(b[r].String())
		buf.WriteRune('\n')
	}
	buf.WriteString("+-------+\n")
	return buf.String()
}

func (b board) clone() board {
	ret := make(board, len(b))
	copy(ret, b)
	return ret
}

func (b *board) placeShape(r, c int, s shape) {
	h := s.height()
	for y := 0; y < h; y++ {
		b.ensureRow(r + y)
		(*b)[r+y] |= row(s.row(y) << c)
	}
}
func (b *board) ensureRow(r int) {
	for len(*b) <= r {
		*b = append(*b, 0)
	}
}

/*
func (b *board) placeRock(r, c int) {
	b.ensureRow(r)
	(*b)[r] |= 1 << c
}
*/

func (b board) collides(r, c int, s shape) bool {
	h := s.height()
	for y := 0; y < h; y++ {
		if len(b) <= r+y {
			// board is empty from here on up
			return false
		}
		if b[r+y]&(s.row(y)<<c) != 0 {
			return true
		}
	}
	return false
}

func (b board) height() int {
	for r := len(b) - 1; r >= 0; r-- {
		if b[r] != 0 {
			return r + 1
		}
	}
	return 0
}

func (b *board) play(shapes []shape, jets []direction, pos boardPos) boardPos {
	c := 2
	r := b.height() + 3
	// debug := func(m string) {
	// 	b2 := b.clone()
	// 	b2.placeShape(r, c, s)
	// 	fmt.Println(m + "\n" + b2.String())
	// }
	s := shapes[pos.shapeNo]
	pos.shapeNo = (pos.shapeNo + 1) % len(shapes)
	for {
		// debug("start round")
		// round part 1: jet
		jet := jets[pos.jetNo]
		pos.jetNo = (pos.jetNo + 1) % len(jets)
		if nc := jet.apply(c); s.canMove(c, jet) && !b.collides(r, nc, s) {
			c = nc
		}
		// debug("after jet")
		if r == 0 || b.collides(r-1, c, s) {
			b.placeShape(r, c, s)
			pos.lastRows = b.lastRows()
			return pos
		} else {
			r--
		}
	}
}

// can fit 9 rows into a uint64
const rowsPerInt = 64 / rowWidth

func (b board) lastRows() uint64 {
	h := b.height()
	var lastRows uint64
	for i := h - 1; i >= h-rowsPerInt && i >= 0; i-- {
		lastRows = (lastRows << rowWidth) | uint64(b[i])
	}
	return lastRows
}

type boardPos struct {
	lastRows uint64
	shapeNo  int
	jetNo    int
}

// playLoop runs turns calls of play, iterating through the shapes. It returns a
// slice of the board positions after each turn. The lastRows of the initial
// position is ignored.
func (b *board) playLoop(
	shapes []shape,
	jets []direction,
	// lastRows is ignored
	pos boardPos,
	turns int,
) []boardPos {
	ret := make([]boardPos, 0, turns)
	for i := 0; i < turns; i++ {
		pos = b.play(shapes, jets, pos)
		ret = append(ret, pos)
	}
	return ret
}

func (b *board) loopFinder(shapes []shape, jets []direction) *loopFinder {
	return &loopFinder{
		b,
		0,
		boardPos{},
		shapes,
		jets,
		make(map[boardPos]int, 1000),
		make(map[boardPos]int, 1000),
	}
}

type loopFinder struct {
	b          *board
	rounds     int
	pos        boardPos
	shapes     []shape
	jets       []direction
	lastRound  map[boardPos]int
	lastHeight map[boardPos]int
}

func (l *loopFinder) run() {
	// update state if we are resuming from a failed verify
	if l.lastRound[l.pos] != l.rounds {
		l.lastHeight[l.pos] = l.b.height()
		l.lastRound[l.pos] = l.rounds
	}
	for {
		l.pos = l.b.play(l.shapes, l.jets, l.pos)
		l.rounds++
		if _, ok := l.lastHeight[l.pos]; ok {
			return
		}
		l.lastHeight[l.pos] = l.b.height()
		l.lastRound[l.pos] = l.rounds
	}
}

func (l *loopFinder) verify() (roundOffset, rounds, heightOffset, height int, valid bool) {
	roundOffset = l.lastRound[l.pos]
	heightOffset = l.lastHeight[l.pos]
	h2, r2 := l.b.height(), l.rounds
	height = h2 - heightOffset
	rounds = r2 - roundOffset
	// clone the board in case we fail verification, we can run again
	b2 := l.b.clone()
	// each loop must increase the height at least one, so playing the height
	// delta more rounds should ensure we can verify the loop
	b2.playLoop(l.shapes, l.jets, l.pos, h2-heightOffset)
	bsub1 := b2[heightOffset:h2]
	bsub2 := b2[h2 : h2+height]
	for i := 0; i < len(bsub1); i++ {
		if bsub1[i] != bsub2[i] {
			valid = false
			return
		}
	}
	valid = true
	return
}

func (l *loopFinder) search() (roundOffset, rounds, heightOffset, height int) {
	for {
		l.run()
		var valid bool
		if roundOffset, rounds, heightOffset, height, valid = l.verify(); valid {
			return
		}
	}
}

func (l *loopFinder) heightAfter(targetRounds int) int {
	roundOffset, loopRounds, heightOffset, loopHeight := l.search()
	// to get to the target number of rounds, we do roundOffset, plus some number
	// of loops length `rounds`, plus a remainder we need to compute the remainder
	// to play it out and find out the height increase.
	loops := (targetRounds - roundOffset) / loopRounds
	remainder := (targetRounds - roundOffset) % loopRounds

	b2 := l.b.clone()
	h1 := b2.height()
	b2.playLoop(l.shapes, l.jets, l.pos, remainder)
	h2 := b2.height()
	heightRemainder := h2 - h1
	finalHeight := heightOffset + loopHeight*loops + heightRemainder
	return finalHeight
}
