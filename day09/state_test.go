package day09

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pos_move(t *testing.T) {
	tests := []struct {
		start pos
		dir   rune
		want  pos
	}{
		{pos{0, 0}, 'R', pos{1, 0}},
		{pos{1, 0}, 'L', pos{0, 0}},
		{pos{0, 0}, 'U', pos{0, 1}},
		{pos{0, 1}, 'D', pos{0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v+%c=%v", tt.start, tt.dir, tt.want), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.start.move(tt.dir))
		})
	}
}

func Test_pos_follow(t *testing.T) {
	tests := []struct{ start, head, want pos }{
		{pos{1, 1}, pos{3, 1}, pos{2, 1}},
		{pos{1, 3}, pos{1, 1}, pos{1, 2}},
		{pos{1, 1}, pos{3, 2}, pos{2, 2}},
		{pos{1, 1}, pos{3, 2}, pos{2, 2}},
		{pos{3, 0}, pos{4, 1}, pos{3, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v->%v=%v", tt.start, tt.head, tt.want), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.start.follow(tt.head))
		})
	}
}
