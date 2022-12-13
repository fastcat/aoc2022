package i

import "unicode/utf8"

func Runes(s string) Iterable[rune] { return runes(s) }

type runes string

// var _ Iterable[rune] = Runes("")

func (r runes) Iterator() Iterator[rune] {
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
	return Map(in, func(i T, _ int) string { return string(i) })
}

type stringable interface {
	~[]rune | ~[]byte
}
