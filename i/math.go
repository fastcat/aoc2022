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

type Addable = u.Addable
