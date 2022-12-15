package u

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
