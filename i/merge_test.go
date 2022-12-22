package i

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	type test struct {
		in   []string
		seed func(string) string
		m    MergeFunc[string, string]
		want []string
	}
	ident := func(s string) string { return s }
	tests := []test{
		{
			[]string{"a", " b", " c", "d"},
			ident,
			func(prior, next string) (merged string, merge bool) {
				if next[0] == ' ' {
					return prior + next, true
				}
				return prior, false
			},
			[]string{"a b c", "d"},
		},
		{
			[]string{"a", " b", " c"},
			ident,
			func(prior, next string) (merged string, merge bool) {
				if next[0] == ' ' {
					return prior + next, true
				}
				return prior, false
			},
			[]string{"a b c"},
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, tt.want, ToSlice(Merge(Slice(tt.in), tt.seed, tt.m)))
		})
	}
}
