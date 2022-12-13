package i

import (
	"github.com/fastcat/aoc2022/u"
	"golang.org/x/exp/constraints"
)

type slice[T any] []T

func Slice[T any](s []T) Iterable[T] {
	return slice[T](s)
}

func (i slice[T]) Iterator() Iterator[T] {
	return &sliceIterator[T]{i, 0}
}

type sliceIterator[T any] struct {
	s []T
	i int
}

func (i *sliceIterator[T]) Next() (value T, done bool) {
	if i.i >= len(i.s) {
		done = true
		return
	}
	value = i.s[i.i]
	i.i++
	return
}

type revSlice[T any] []T

func RevSlice[T any](s []T) Iterable[T] {
	return revSlice[T](s)
}

func (i revSlice[T]) Iterator() Iterator[T] {
	return &revSliceIterator[T]{i, len(i) - 1}
}

type revSliceIterator[T any] struct {
	s []T
	i int
}

func (i *revSliceIterator[T]) Next() (value T, done bool) {
	if i.i < 0 {
		done = true
		return
	}
	value = i.s[i.i]
	i.i--
	return
}

func ToSlice[T any](in Iterable[T]) []T {
	var out []T
	For(in, func(i T, _ int) { out = append(out, i) })
	return out
}
func ToSliceI[T any](in Iterator[T]) []T {
	var out []T
	ForI(in, func(i T, _ int) { out = append(out, i) })
	return out
}

func Top[T constraints.Ordered](in Iterable[T], n int) []T {
	// TODO: use heapsort instead
	return Reduce(in, make([]T, 0, n+1), func(top []T, item T, _ int) []T {
		top = append(top, item)
		u.SortDesc(top)
		if len(top) > n {
			top = top[0:n]
		}
		return top
	})
}
