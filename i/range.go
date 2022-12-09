package i

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Range emits values from start (inclusive) to end (exclusive) jumping by step.
// Step may be negative, but not zero.
func Range[T constraints.Integer](start, end, step T) Iterable[T] {
	if step == 0 {
		panic(errors.New("zero step"))
	}
	if step > 0 && end < start || step < 0 && end > start {
		panic(errors.New("step vs end-start mismatch"))
	}
	return Funcer(func() Iterator[T] {
		i := start
		if step > 0 {
			return Func(func() (T, bool) {
				if i >= end {
					return 0, true
				}
				ret := i
				i += step
				return ret, false
			})
		} else {
			return Func(func() (T, bool) {
				if i <= end {
					return 0, true
				}
				ret := i
				i += step
				return ret, false
			})
		}
	})
}
