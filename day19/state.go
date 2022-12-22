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

func initialState() state {
	return state{
		bots: inventory{ore: 1},
	}
}
