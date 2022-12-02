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

type Addable interface {
	constraints.Ordered | constraints.Complex
}
