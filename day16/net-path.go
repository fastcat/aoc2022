package day16

import "fmt"

type atMinute []atPos
type atPos []withOpen
type withOpen map[int64]best
type best struct {
	prev  key
	rate  int
	total int
}
type key struct {
	min  int
	pos  int
	open int64
}

func (b atMinute) get(k key) *best {
	// special case for starting condition
	if k.pos < 0 && k.min == 0 && k.open == 0 {
		return &best{k, 0, 0}
	}
	if k.min >= len(b) {
		return nil
	}
	return b[k.min].get(k)
}
func (b *atMinute) set(k key, bv best) {
	if len(*b) < k.min {
		if len(*b) < k.min-1 {
			panic(fmt.Errorf("skipping time"))
		}
		(*b) = append((*b), make(atPos, 0, len((*b)[k.min-1])))
	}
	(*b)[k.min].set(k, bv)
}
func (m atPos) get(k key) *best {
	if k.pos >= len(m) {
		return nil
	}
	return m[k.pos].get(k)
}
func (m *atPos) set(k key, bv best) {
	for len(*m) <= k.pos {
		(*m) = append((*m), nil)
	}
	if (*m)[k.pos] == nil {
		(*m)[k.pos] = withOpen{k.open: bv}
	} else {
		(*m)[k.pos][k.open] = bv
	}
}
func (mp withOpen) get(k key) *best {
	if mp == nil {
		return nil
	}
	if b, ok := mp[k.open]; !ok {
		return nil
	} else {
		return &b
	}
}

func (rn *reducedNet) search(lim int) (key, []best) {
	b := make(atMinute, lim+1)
	// minute 0 is nonsense
	b[0] = atPos{}
	var walk func(key)
	walk = func(k key) {
		if k.min >= lim {
			// time limit is up
			return
		}
		kb := b.get(k)
		if kb == nil {
			panic("walk from unregistered start")
		}
		// minute 0 is special, we have a bogus pos and must move, can't wait or do
		// a valve
		if k.pos >= 0 {
			k2 := k
			k2.min++
			nt := kb.total + kb.rate
			// we could do nothing and let things flow
			if nb := b.get(k2); nb == nil || nb.total < nt {
				b.set(k2, best{k, kb.rate, nt})
				walk(k2)
			}
			// if the current valve is closed, we could open it
			if pm := int64(1) << k.pos; k.open&pm == 0 {
				k3 := k2 // keep the min++
				k3.open = k2.open | pm
				if nb := b.get(k3); nb == nil || nb.total < nt {
					// rate will go up after we open the valve, but won't count towards
					// the total until the following minute(s)
					b.set(k3, best{k, kb.rate + rn.rates[k.pos], nt})
					walk(k3)
				}
			}
		}
		// we could travel somewhere that has a closed valve (to open it). we should
		// never move twice in a row, that should have been just one move.
		if k.pos < 0 || kb.prev.pos == k.pos {
			var d []int
			if k.pos < 0 {
				d = rn.zerodists
			} else {
				d = rn.dists[k.pos]
			}
			for np, nd := range d {
				if np == k.pos || k.open&(1<<np) != 0 {
					continue
				}
				k4 := k
				k4.min += nd
				k4.pos = np
				if k4.min > 30 {
					// unreachable
					continue
				}
				// things flow while we move!
				nt := kb.total + kb.rate*nd
				if nb := b.get(k4); nb == nil || nb.total < nt {
					b.set(k4, best{k, kb.rate, nt})
					walk(k4)
				}
			}
		}
	}
	walk(key{0, -1, 0})

	// find the best endpoint in the last minute, then walk the path backwards
	// from there
	last := b[lim]
	var bb best
	bk := key{min: lim}
	for pi, p := range last {
		for o, b := range p {
			if b.total > bb.total {
				bb = b
				bk.open = o
				bk.pos = pi
			}
		}
	}
	path := make([]best, lim)
	for i, k := len(path)-1, bk; ; i-- {
		bb := b[k.min][k.pos][k.open]
		path[i] = bb
		k = bb.prev
		if k.pos < 0 {
			path = path[i-1:]
			break
		}
	}
	return bk, path
}
