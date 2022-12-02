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
			s := Split[byte](Slice([]byte(tt.input)), []byte(tt.separator))
			it := s.Iterator()
			// for loop overshoots  intentionally so that we check the return for an
			// extra call to Next() after it was already done
			for i := 0; i <= len(tt.expected)+1; i++ {
				value, done := it.Next()
				a.Equal(i >= len(tt.expected), done)
				if done {
					a.Nil(value)
				} else {
					a.NotNil(value)
				}
				if i < len(tt.expected) {
					a.Equal(tt.expected[i], string(value))
				}
			}
		})
	}
}
