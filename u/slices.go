package u

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sum[T Addable](in []T) T {
	var sum T
	for _, v := range in {
		sum += v
	}
	return sum
}

func SortDesc[T constraints.Ordered](s []T) {
	slices.SortFunc(s, func(a, b T) bool { return b < a })
}

func Reverse[T any](s []T) {
	if len(s) == 0 {
		return
	}
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

type Addable interface {
	constraints.Ordered | constraints.Complex
}
