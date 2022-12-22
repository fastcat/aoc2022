package day19

type graph struct {
	b       *blueprint
	runtime int
}

func (g *graph) search() state {
	s := initialState()
	return g.walk(0, s, s, make([]idx, 0, g.runtime))
}

func (g *graph) walk(min int, s, curBest state, path []idx) state {
	if min >= g.runtime {
		return s
	}
	// at least we can wait and just accumulate
	best := s.wait(uint8(g.runtime - min))
	if best.inv[geode] > curBest.inv[geode] {
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
			if maxG < int(best.inv[geode]) {
				// no way this can improve on the best
				continue
			}
		}
		sbw := g.walk(nm, sb, curBest, append(path, bot))
		if sbw.inv[geode] > best.inv[geode] {
			best = sbw
			if best.inv[geode] > curBest.inv[geode] {
				curBest = best
			}
		}
	}
	return best
}

var triangle = [...]int{
	0, 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91, 105, 120, 136, 153, 171, 190, 210, 231, 253,
}
