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
		ret = ret && a.Equal(i >= len(expected), done)
		if done {
			ret = ret && a.Zero(value)
		} else {
			ret = ret && a.NotZero(value)
		}
		if i < len(expected) {
			ret = ret && a.Equal(expected[i], value)
		}
	}
	return ret
}
