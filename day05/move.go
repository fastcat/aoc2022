package day05

import (
	"fmt"

	"github.com/fastcat/aoc2022/u"
)

type move struct {
	source, dest int
	count        int
}

func parseMove(in string) (parsed move) {
	_, err := fmt.Sscanf(in, "move %d from %d to %d\n", &parsed.count, &parsed.source, &parsed.dest)
	u.PanicIf(err)
	// convert from 1-index to 0-index
	parsed.source--
	parsed.dest--
	return
}
