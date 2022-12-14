package day13

import (
	"fmt"
	"testing"

	"github.com/fastcat/aoc2022/i"
	"github.com/stretchr/testify/assert"
)

func listOf(items ...any) listItem {
	l := make(listItem, 0)
	for _, i := range items {
		switch i := i.(type) {
		case int:
			l = append(l, numberItem(i))
		case numberItem:
			l = append(l, l, i)
		case listItem:
			l = append(l, i)
		case item:
			l = append(l, i)
		default:
			panic(fmt.Errorf("can't handle %T %#[1]v", i))
		}
	}
	return l
}

func Test_parsePacket(t *testing.T) {
	tests := []struct {
		in   string
		want listItem
	}{
		{
			"[1,2,3]",
			listOf(1, 2, 3),
		},
		{
			"[1,[2,3],4]",
			listOf(1, listOf(2, 3), 4),
		},
		{
			"[]",
			listOf(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			assert.NotNil(t, tt.want)
			got := parsePacket(i.Runes(tt.in))
			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)
		})
	}
}
