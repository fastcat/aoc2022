package day22

import "fmt"

type place byte

const (
	void  place = ' '
	space place = '.'
	wall  place = '#'
)

type grid [][]place

type board struct {
	g       grid
	portals map[state]state
}

type dir byte

const (
	right dir = '>'
	down  dir = 'v'
	left  dir = '<'
	up    dir = '^'
)

func (d dir) Value() int {
	switch d {
	case right:
		return 0
	case down:
		return 1
	case left:
		return 2
	case up:
		return 3
	default:
		panic(fmt.Errorf("invalid direction: %c", d))
	}
}

type state struct {
	r, c int
	d    dir
}

func (s state) Row() int    { return s.r + 1 }
func (s state) Column() int { return s.c + 1 }

type move int

const (
	turnLeft  move = -1
	turnRight move = -2
)
