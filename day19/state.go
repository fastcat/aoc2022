package day19

import "math"

// in 24 minutes, can at best get a geode robot on turn 6, and triangle number
// for 24-6 is 171, so everything fits in bytes.

type inventory [4]uint8

func (i *inventory) sub(c cost) {
	i[ore] -= c[ore]
	i[clay] -= c[clay]
	i[obsidian] -= c[obsidian]
}
func (i *inventory) addMul(i2 inventory, x uint8) {
	i[ore] += i2[ore] * x
	i[clay] += i2[clay] * x
	i[obsidian] += i2[obsidian] * x
	i[geode] += i2[geode] * x
}
func (i inventory) u32() uint32 {
	return uint32(i[ore])<<(8*ore) |
		uint32(i[clay])<<(8*clay) |
		uint32(i[obsidian])<<(8*obsidian) |
		uint32(i[geode])<<(8*geode)
}
func invFromU32(v uint32) inventory {
	return inventory{
		ore:      byte(v >> (ore * 8) & math.MaxUint8),
		clay:     byte(v >> (clay * 8) & math.MaxUint8),
		obsidian: byte(v >> (obsidian * 8) & math.MaxUint8),
		geode:    byte(v >> (geode * 8) & math.MaxUint8),
	}
}

type state struct {
	inv, bots inventory
}

func (s state) timeToBuild(bot cost) int {
	need := bot
	for i := ore; i <= obsidian; i++ {
		if s.inv[i] > need[i] {
			need[i] = 0
		} else {
			need[i] -= s.inv[i]
		}
	}
	max := 0
	for i := ore; i <= obsidian; i++ {
		if need[i] == 0 {
			continue
		}
		if s.bots[i] == 0 {
			// will never be able to build it
			return math.MaxUint8
		}
		if turns := int(need[i]+s.bots[i]-1) / int(s.bots[i]); turns > max {
			max = turns
		}
	}
	return max
}

func (s state) play(b *blueprint, bot idx) state {
	s2 := s
	// factory
	if bot != none {
		s2.bots[bot]++
		s2.inv.sub(b[bot])
	}
	// accumulate, uses bot count pre-factory
	s2.inv.addMul(s.bots, 1)
	return s2
}

func (s state) wait(turns uint8) state {
	s.inv.addMul(s.bots, turns)
	return s
}

func (s state) waitAndBuild(turns uint8, b *blueprint, bot idx) state {
	s.inv.addMul(s.bots, turns)
	s.inv.sub(b[bot])
	s.bots[bot]++
	return s
}

func (s state) u64() uint64 {
	return (uint64(s.inv.u32()) << 32) | uint64(s.bots.u32())
}

func stateFromU64(v uint64) state {
	return state{
		inv:  invFromU32(uint32(v >> 32 & math.MaxUint32)),
		bots: invFromU32(uint32(v & math.MaxUint32)),
	}
}

func initialState() state {
	return state{
		bots: inventory{ore: 1},
	}
}

func (s state) playSimple(b *blueprint) state {
	for bot := geode; bot >= 0; bot-- {
		if b[bot].canBuild(s.inv) {
			return s.play(b, bot)
		}
	}
	return s.play(b, none)
}
