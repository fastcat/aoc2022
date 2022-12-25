package day21

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_node_Invert(t *testing.T) {
	type test struct {
		in   string
		want int
	}
	tests := []test{
		{
			"root: aaaa + humn\n" +
				"aaaa: 3\n" +
				"humn: 12\n",
			3,
		},
		{
			"root: aaaa * bbbb\n" +
				"aaaa: cccc - dddd\n" +
				"bbbb: eeee / humn\n" +
				"cccc: 12\n" +
				"dddd: 3\n" +
				"eeee: 18\n" +
				"humn: 0\n",
			2,
		},
	}

	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			r := parse(tt.in)
			v := r.InvertRootValue()
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestPart2Sample(t *testing.T) {
	a := assert.New(t)
	r := parse(sample)
	v := r.InvertRootValue()
	a.Equal(301, v)
}

func TestPart2(t *testing.T) {
	r := parse(input)
	v := r.InvertRootValue()
	t.Log(v)
}
