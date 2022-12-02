package i

import "unicode/utf8"

type Runes string

// var _ Iterable[rune] = Runes("")

func (r Runes) Iterator() Iterator[rune] {
	i := runeIterator(r)
	return &i
}

type runeIterator string

func (i *runeIterator) Next() (rune, bool) {
	r, w := utf8.DecodeRuneInString(string(*i))
	*i = (*i)[w:]
	return r, w == 0
}

func ToStrings[T stringable](in Iterable[T]) Iterable[string] {
	return Map(in, func(i T) string { return string(i) })
}

type stringable interface {
	~[]rune | ~[]byte
}
