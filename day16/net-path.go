package day16

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
	for len(*b) <= k.min {
		(*b) = append((*b), make(atPos, 0, k.pos))
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

type walker struct {
	rn  *reducedNet
	b   atMinute
	lim int
}

func (w *walker) walkFrom(k key) {
	if len(w.b) == 0 {
		w.b = make(atMinute, 1, w.lim+1)
		// minute 0 is nonsense
		w.b[0] = atPos{}
	}
	if k.min >= w.lim {
		// time limit is up
		return
	}
	kb := w.b.get(k)
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
		if nb := w.b.get(k2); nb == nil || nb.total < nt {
			w.b.set(k2, best{k, kb.rate, nt})
			w.walkFrom(k2)
		}
		// if the current valve is closed, we could open it
		if pm := int64(1) << k.pos; k.open&pm == 0 {
			k3 := k2 // keep the min++
			k3.open = k2.open | pm
			if nb := w.b.get(k3); nb == nil || nb.total < nt {
				// rate will go up after we open the valve, but won't count towards
				// the total until the following minute(s)
				w.b.set(k3, best{k, kb.rate + w.rn.rates[k.pos], nt})
				w.walkFrom(k3)
			}
		}
	}
	// we could travel somewhere that has a closed valve (to open it). we should
	// never move twice in a row, that should have been just one move.
	if k.pos < 0 || kb.prev.pos == k.pos {
		var d []int
		if k.pos < 0 {
			d = w.rn.zerodists
		} else {
			d = w.rn.dists[k.pos]
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
			if nb := w.b.get(k4); nb == nil || nb.total < nt {
				w.b.set(k4, best{k, kb.rate, nt})
				w.walkFrom(k4)
			}
		}
	}
}

func (rn *reducedNet) walker(lim int) *walker {
	return &walker{
		rn:  rn,
		lim: lim,
	}
}

func (w *walker) walk() {
	w.walkFrom(key{0, -1, 0})
}

func (w *walker) pathFrom(end key) []best {
	path := make([]best, w.lim)
	for i, k := len(path)-1, end; ; i-- {
		bb := w.b[k.min][k.pos][k.open]
		path[i] = bb
		k = bb.prev
		if k.pos < 0 {
			path = path[i-1:]
			break
		}
	}
	return path
}

func (w *walker) bestPath() (key, []best) {
	// find the best endpoint in the last minute, then walk the path backwards
	// from there
	last := w.b[w.lim]
	var bb best
	bk := key{min: w.lim}
	for pi, p := range last {
		for o, b := range p {
			if b.total > bb.total {
				bb = b
				bk.open = o
				bk.pos = pi
			}
		}
	}
	return bk, w.pathFrom(bk)
}

func (w *walker) bestPair() (key, key, int) {
	end := w.b[w.lim]
	// find the best endpoint for each open mask
	bestByOpen := map[int64]int{}
	for ei, ep := range end {
		for eo, b := range ep {
			if bei, ok := bestByOpen[eo]; !ok {
				bestByOpen[eo] = ei
			} else if end[bei][eo].total < b.total {
				bestByOpen[eo] = ei
			}
		}
	}

	// make the mask of used bits to invert a bitset
	mask := int64(0)
	for i := range w.rn.rates {
		mask |= 1 << i
	}

	bestOpen, bestTotal := int64(0), 0

	// find disjoint pairs, use simplistic mask ordering to deduplicate
	for o, bei := range bestByOpen {
		oo := o ^ mask
		if oo < o {
			continue
		}
		if obei, ok := bestByOpen[oo]; !ok {
			continue
		} else {
			ob := end[bei][o]
			oob := end[obei][oo]
			t := ob.total + oob.total
			if t > bestTotal {
				bestOpen, bestTotal = o, t
			}
		}
	}

	k1 := key{min: w.lim, pos: bestByOpen[bestOpen], open: bestOpen}
	k2 := key{min: w.lim, pos: bestByOpen[bestOpen^mask], open: bestOpen ^ mask}
	return k1, k2, bestTotal
}

func (rn *reducedNet) search(lim int) (key, []best) {
	w := rn.walker(lim)
	w.walk()
	return w.bestPath()
}
