package u

import (
	"errors"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Map[T, U any](in []T, f func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func Sum[T Addable](in []T) T {
	var sum T
	for _, v := range in {
		sum += v
	}
	return sum
}
func SumF[T any, U Addable](in []T, f func(T) U) U {
	var sum U
	for _, v := range in {
		sum += f(v)
	}
	return sum
}

func Max[T constraints.Ordered](in []T) (T, int) {
	if len(in) == 0 {
		panic(errors.New("empty slice"))
	}
	max, maxi := in[0], 0
	for i := 1; i < len(in); i++ {
		if in[i] > max {
			max, maxi = in[i], i
		}
	}
	return max, maxi
}

func Top[T constraints.Ordered](in []T, n int) []T {
	// TODO: use heapsort instead
	top := make([]T, 0, n+1)
	if len(in) <= n {
		top = append(top, in...)
		SortDesc(top)
		return top
	}
	top = append(top, in[0:n]...)
	SortDesc(top)
	for i := n; i < len(in); i++ {
		top = append(top, in[i])
		SortDesc(top)
		top = top[0:n]
	}
	return top
}

func SortDesc[T constraints.Ordered](s []T) {
	slices.SortFunc(s, func(a, b T) bool { return b < a })

}

type Addable interface {
	constraints.Ordered | constraints.Complex
}
