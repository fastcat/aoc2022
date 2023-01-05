package day24

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_layerView_occupied(t *testing.T) {
	type test struct {
		m    int
		p    pos
		want bool
	}
	type testGroup struct {
		in    string
		tests []test
	}
	groups := []testGroup{
		{
			"" +
				"#.##\n" +
				"#..#\n" +
				"##.#\n",
			[]test{
				{0, pos{-1, -1}, true},
				{0, pos{-1, 0}, false},
				{0, pos{-1, 1}, true},
				{0, pos{-1, 2}, true},
				{0, pos{0, -1}, true},
				{0, pos{0, 0}, false},
				{0, pos{0, 1}, false},
				{0, pos{0, 2}, true},
				{0, pos{1, -1}, true},
				{0, pos{1, 0}, true},
				{0, pos{1, 1}, false},
				{0, pos{1, 2}, true},
			},
		},
		{
			"" +
				"#.##\n" +
				"#>.#\n" +
				"#..#\n" +
				"##.#\n",
			[]test{
				{0, pos{0, 0}, true},
				{0, pos{0, 1}, false},
				{1, pos{0, 0}, false},
				{1, pos{0, 1}, true},
				{2, pos{0, 0}, true},
				{2, pos{0, 1}, false},
			},
		},
		{
			"" +
				"#.###\n" +
				"#<..#\n" +
				"#...#\n" +
				"###.#\n",
			[]test{
				{0, pos{0, 0}, true},
				{0, pos{0, 1}, false},
				{0, pos{0, 2}, false},
				{1, pos{0, 0}, false},
				{1, pos{0, 1}, false},
				{1, pos{0, 2}, true},
				{2, pos{0, 0}, false},
				{2, pos{0, 1}, true},
				{2, pos{0, 2}, false},
			},
		},
		{
			"" +
				"#.###\n" +
				"#^.#\n" +
				"#..#\n" +
				"#..#\n" +
				"###.#\n",
			[]test{
				{0, pos{0, 0}, true},
				{0, pos{1, 0}, false},
				{0, pos{2, 0}, false},
				{1, pos{0, 0}, false},
				{1, pos{1, 0}, false},
				{1, pos{2, 0}, true},
				{2, pos{0, 0}, false},
				{2, pos{1, 0}, true},
				{2, pos{2, 0}, false},
			},
		},
		{
			"" +
				"#.###\n" +
				"#v.#\n" +
				"#..#\n" +
				"###.#\n",
			[]test{
				{0, pos{0, 0}, true},
				{0, pos{1, 0}, false},
				{1, pos{0, 0}, false},
				{1, pos{1, 0}, true},
				{2, pos{0, 0}, true},
				{2, pos{1, 0}, false},
			},
		},
	}
	for gi, tg := range groups {
		b := parse(tg.in)
		for _, tt := range tg.tests {
			t.Run(fmt.Sprintf("%d@%d@r%dc%d", gi, tt.m, tt.p.r, tt.p.c), func(t *testing.T) {
				bv := b.viewAt(tt.m)
				got := bv.occupied(tt.p)
				assert.Equal(t, tt.want, got)
			})
		}
	}

}
