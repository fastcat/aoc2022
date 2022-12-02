package day02

import (
	"fmt"
)

type RPS int8

const (
	Rock     RPS = 0
	Paper    RPS = 1
	Scissors RPS = 2
)

func (r RPS) String() string {
	switch r {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	default:
		return fmt.Sprintf("!(invalid:%d)", r)
	}
}

func parseRPS(in rune) RPS {
	switch in {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	}
	panic(fmt.Errorf("bad input %c", in))
}

func (r RPS) score() int {
	switch r {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic(fmt.Errorf("bad input %c", r))
}

type Outcome int8

const (
	Loss Outcome = -1
	Draw Outcome = 0
	Win  Outcome = 1
)

func parseOutcome(in rune) Outcome {
	switch in {
	case 'X':
		return Loss
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	panic(fmt.Errorf("bad input %c", in))
}

func (o Outcome) String() string {
	switch o {
	case Loss:
		return "loss"
	case Draw:
		return "draw"
	case Win:
		return "win"
	default:
		return fmt.Sprintf("!(invalid:%d)", o)
	}
}

func (r RPS) vs(o RPS) Outcome {
	if r == o {
		return Draw
	} else if (r+1)%3 == o {
		return Loss
	}
	return Win
}

func (r RPS) rev(o Outcome) RPS {
	return RPS((3 + int8(r) + int8(o)) % 3)
}

func (o Outcome) score() int {
	if o <= Loss {
		return 0
	} else if o == Draw {
		return 3
	} else {
		return 6
	}
}
