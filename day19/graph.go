package day19

import (
	"sync"

	"github.com/fastcat/aoc2022/i"
)

type graph struct {
	b       *blueprint
	runtime int
}

func (g *graph) search() uint8 {
	s := initialState()
	return g.walk(0, s, 0)
}

func (g *graph) walk(min int, s state, curBest uint8) uint8 {
	if min >= g.runtime {
		return s.inv[geode]
	}
	// at least we can wait and just accumulate
	best := s.wait(uint8(g.runtime - min)).inv[geode]
	if best > curBest {
		curBest = best
	}
	// walk the better paths first so we can prune more aggressively
	for bot := geode; bot >= ore; bot-- {
		// need time to save up and then one more to build the bot
		ttb := s.timeToBuild(g.b[bot]) + 1
		nm := min + ttb
		if nm >= g.runtime {
			// no time left to build this bot
			continue
		}
		sb := s.waitAndBuild(uint8(ttb), g.b, bot)
		tl := g.runtime - nm
		if tl < len(triangle) {
			maxG := int(sb.inv[geode]) + int(sb.bots[geode])*tl + triangle[tl]
			if maxG < int(best) {
				// no way this can improve on the best
				continue
			}
		}
		sbw := g.walk(nm, sb, curBest)
		if sbw > best {
			best = sbw
			if best > curBest {
				curBest = best
			}
		}
	}
	return best
}

func quality(n int, best uint8) int {
	return n * int(best)
}
func qualitySum(best []uint8) int {
	return i.Sum(i.Map(i.Slice(best), func(b uint8, i int) int {
		return quality(i+1, b)
	}))
}

func searchMany(bps []*blueprint) []uint8 {
	best := make([]uint8, len(bps))
	wg := sync.WaitGroup{}
	wg.Add(len(bps))
	fb := func(i int, b *blueprint) {
		defer wg.Done()
		g := graph{b, 24}
		best[i] = g.search()
	}
	for i, b := range bps {
		go fb(i, b)
	}
	wg.Wait()
	return best
}

var triangle = [...]int{
	0, 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91, 105, 120, 136, 153, 171, 190, 210, 231, 253,
}
