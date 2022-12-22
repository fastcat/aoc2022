package day19

type inventory struct {
	ore, clay, obsidian, geodes int
}

func (i *inventory) sub(c cost) {
	i.ore -= c.ore
	i.clay -= c.clay
	i.obsidian -= c.obsidian
}

type state struct {
	minute                                     int
	inv                                        inventory
	oreBots, clayBots, obsidianBots, geodeBots int
}

func initialState() state {
	return state{
		oreBots: 1,
	}
}

func (s state) playSimple(b *blueprint) state {
	s2 := s
	// factory
	if b.geodeBot.canBuild(s.inv) {
		s2.geodeBots++
		s2.inv.sub(b.geodeBot)
	} else if b.obsidianBot.canBuild(s.inv) {
		s2.obsidianBots++
		s2.inv.sub(b.obsidianBot)
	} else if b.clayBot.canBuild(s.inv) {
		s2.clayBots++
		s2.inv.sub(b.clayBot)
	} else if b.oreBot.canBuild(s.inv) {
		s2.oreBots++
		s2.inv.sub(b.oreBot)
	}
	// collection, uses pre-build bot numbers
	s2.inv.ore += s.oreBots
	s2.inv.clay += s.clayBots
	s2.inv.obsidian += s.obsidianBots
	s2.inv.geodes += s.geodeBots
	s2.minute++
	return s2
}
