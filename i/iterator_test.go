package i

import (
	"github.com/stretchr/testify/assert"
)

func AssertIterator[T any](
	a *assert.Assertions,
	expected []T,
	it Iterator[T],
) bool {
	ret := true
	// for loop overshoots  intentionally so that we check the return for an
	// extra call to Next() after it was already done
	for i := 0; i <= len(expected)+1; i++ {
		value, done := it.Next()
		if a.Equal(i >= len(expected), done) {
			if done {
				ret = ret && a.Zero(value)
			} else if i < len(expected) {
				ret = ret && a.Equal(expected[i], value)
			}
		} else {
			ret = false
		}
	}
	return ret
}
