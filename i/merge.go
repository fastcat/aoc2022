package i

import "github.com/fastcat/aoc2022/u"

// MergeFunc handles the iteration step during a Merge iteration. `prior`
// receives either the `seed` value if this is the first item in iteration or
// after a split, or else the prior `merged` return from a call to this function
// (which returned `true` for `merge`). `next` receives the next value from the
// input iterator. If the `next` value should not be merged, it should return
// the any value for `merged` (it will be ignored) and `false` for `merge`,
// otherwise it should merge `next` into `prior` and return that and `true`. It
// will not receive a chance to do the final merge for the final item.
type MergeFunc[T, U any] func(prior U, next T) (merged U, merge bool)

func Merge[T, U any](
	in Iterable[T],
	seed func(T) U,
	m func(
		prior U,
		next T,
	) (merged U, merge bool),
) Iterable[U] {
	return &merger[T, U]{in, seed, m}
}

type merger[T, U any] struct {
	in   Iterable[T]
	seed func(T) U
	m    func(U, T) (U, bool)
}

func (m *merger[T, U]) Iterator() Iterator[U] {
	return &mergeIter[T, U]{m.in.Iterator(), m.seed, nil, false, m.m}
}

type mergeIter[T, U any] struct {
	in   Iterator[T]
	seed func(T) U
	next *T
	done bool
	m    func(U, T) (U, bool)
}

func (m *mergeIter[T, U]) Next() (U, bool) {
	if m.done {
		return u.Zero[U](), true
	}
	var merged U
	first := true
	for {
		var next T
		var inDone bool
		if m.next != nil {
			next, inDone, m.next = *m.next, false, nil
		} else if next, inDone = m.in.Next(); inDone {
			m.done = true
			if first {
				return merged, true
			} else {
				return merged, false
			}
		}
		if first {
			merged = m.seed(next)
			first = false
		} else if nextMerged, merge := m.m(merged, next); !merge {
			m.next = &next
			return nextMerged, false
		} else {
			merged = nextMerged
		}
	}
}
