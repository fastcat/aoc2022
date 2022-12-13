package i

import (
	"github.com/fastcat/aoc2022/u"
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](in Iterable[T]) T {
	return Reduce(in, u.Zero[T](), func(min, value T, i int) T {
		if i == 0 || value < min {
			return value
		}
		return min
	})
}

func MinBy[T any, U constraints.Ordered](in Iterable[T], f func(T) U) (T, int) {
	minIdx := -1
	var minT T
	var minVal U
	For(in, func(value T, i int) {
		if i == 0 {
			minIdx, minT, minVal = 0, value, f(value)
		} else if u := f(value); u < minVal {
			minIdx, minT, minVal = i, value, u
		}
	})
	return minT, minIdx
}

func Max[T constraints.Ordered](in Iterable[T]) T {
	return Reduce(in, u.Zero[T](), func(max, value T, i int) T {
		if i == 0 || value > max {
			return value
		}
		return max
	})
}

func Sum[T Addable](in Iterable[T]) T {
	return Reduce(in, u.Zero[T](), func(sum, value T, i int) T {
		return sum + value
	})
}
func SumI[T Addable](in Iterator[T]) T {
	return ReduceI(in, u.Zero[T](), func(sum, value T, i int) T {
		return sum + value
	})
}

type Addable = u.Addable
