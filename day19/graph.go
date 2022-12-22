package day19

import (
	"sync"

	"github.com/fastcat/aoc2022/i"
)

type graph struct {
	b         *blueprint
	runtime   int
	best      []map[state]uint8
	cacheHits uint64
}

func (b *blueprint) newGraph(runtime int) *graph {
	return &graph{
		b:       b,
		runtime: runtime,
	}
}

func (g *graph) search() uint8 {
	s := initialState()
	g.best = make([]map[state]uint8, g.runtime)
	for i := range g.best {
		g.best[i] = make(map[state]uint8)
	}
	best := g.walk(0, s, 0)
	return best
}

func (g *graph) walk(min int, s state, curBest uint8) uint8 {
	if min >= g.runtime {
		return s.inv[geode]
	}
	// su := s.u64()
	if b, ok := g.best[min][s]; ok {
		g.cacheHits++
		return b
	}
	var best uint8
	anyBot := false
	// walk the better paths first so we can prune more aggressively
	for bot := geode; bot >= ore; bot-- {
		// need time to save up and then one more to build the bot
		ttb := s.timeToBuild(g.b[bot]) + 1
		nm := min + ttb
		if nm >= g.runtime {
			// no time left to build this bot
			continue
		}
		anyBot = true
		sb := s.waitAndBuild(uint8(ttb), g.b, bot)
		tl := g.runtime - nm
		if tl < len(triangle) {
			maxG := int(sb.inv[geode]) + int(sb.bots[geode])*tl + triangle[tl]
			if maxG < int(curBest) {
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
	if !anyBot {
		// at least we can accumulate until the end
		sb := s.wait(uint8(g.runtime - min)).inv[geode]
		if sb > best {
			best = sb
		}
	}
	g.best[min][s] = best
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
func qualityProduct(best []uint8) uint64 {
	return i.Product(i.Map(i.Slice(best), func(b uint8, _ int) uint64 { return uint64(b) }))
}

func searchMany(bps []*blueprint, runtime int) []uint8 {
	best := make([]uint8, len(bps))
	wg := sync.WaitGroup{}
	wg.Add(len(bps))
	fb := func(i int, b *blueprint) {
		defer wg.Done()
		g := b.newGraph(runtime)
		best[i] = g.search()
	}
	for i, b := range bps {
		go fb(i, b)
	}
	wg.Wait()
	return best
}

var triangle = [...]int{
	0, 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91, 105, 120, 136, 153, 171,
	190, 210, 231, 253, 276, 300, 325, 351, 378, 406, 435, 465, 496, 528, 561,
	595, 630, 666, 703, 741, 780, 820, 861, 903, 946, 990, 1035, 1081, 1128,
	1176, 1225, 1275, 1326, 1378, 1431,
}
