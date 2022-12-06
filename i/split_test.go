package i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name             string
		input, separator string
		expected         []string
	}{
		{
			"no separators",
			"hello world", "\n",
			[]string{"hello world"},
		},
		{
			"simple",
			"hello\nworld", "\n",
			[]string{"hello", "world"},
		},
		{
			"trailing single",
			"hello\nworld\n", "\n",
			[]string{"hello", "world"},
		},
		{
			"leading",
			"\nhello", "\n",
			[]string{"hello"},
		},
		{
			"multiple",
			"\n\nhello\n\nworld\n\n", "\n",
			[]string{"hello", "world"},
		},
		{
			"multi-char separator",
			"hello\nworld\n\ngoodbye\nworld\n", "\n\n",
			[]string{"hello\nworld", "goodbye\nworld\n"},
		},
		{
			"multiple multi-char",
			"\n\nhello\n\n\n\nworld\n\n", "\n\n",
			[]string{"hello", "world"},
		},
		{
			"complex multi-char",
			"hello\r\nworld\r\ngoodby\rgoodbye", "\r\n",
			[]string{"hello", "world", "goodby\rgoodbye"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			s := Split(Slice([]byte(tt.input)), []byte(tt.separator))
			it := s.Iterator()
			AssertIterator(a, ToSlice(Map(Slice(tt.expected), func(s string, _ int) []byte { return []byte(s) })), it)
		})
	}
}
