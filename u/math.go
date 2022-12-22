package u

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

func Max[T constraints.Ordered](values ...T) T {
	var max T
	for i, v := range values {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}
