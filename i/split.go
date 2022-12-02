package i

import "errors"

func Split[T comparable](in Iterable[T], sep []T) Iterable[[]T] {
	if len(sep) == 0 {
		panic(errors.New("empty separator"))
	}
	return splitter[T]{in, sep}
}

type splitter[T comparable] struct {
	in  Iterable[T]
	sep []T
}

func (s splitter[T]) Iterator() Iterator[[]T] {
	return splitIter[T]{s.in.Iterator(), s.sep}
}

type splitIter[T comparable] struct {
	in  Iterator[T]
	sep []T
}

func (s splitIter[T]) Next() (value []T, done bool) {
	// start skipping leading separators
	sPos := 0
	for {
		item, done := s.in.Next()
		if done {
			// not done to caller if we accumulated anything in value or a partial
			// separator
			if sPos > 0 {
				value = append(value, s.sep[0:sPos]...)
			}
			return value, done && len(value) == 0
		}
		if item == s.sep[sPos] {
			// start/continuation of separator match
			sPos = (sPos + 1) % len(s.sep)
			if sPos == 0 && len(value) != 0 {
				// we hit a full separator after content, return the content
				return value, false
			}
		} else {
			if sPos > 0 {
				// partial separator match, not a separator, add to value
				value = append(value, s.sep[0:sPos]...)
				sPos = 0
			}
			value = append(value, item)
		}
	}
}
