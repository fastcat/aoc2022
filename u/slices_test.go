package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortDesc(t *testing.T) {
	tests := []struct {
		name    string
		in, out []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"reversed", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"shuffled", []int{3, 1, 4, 2, 5}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := append(make([]int, 0, len(tt.in)), tt.in...)
			if tt.in == nil {
				s = nil
			}
			SortDesc(s)
			assert.Equal(t, tt.out, s)
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name    string
		in, out []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"even", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"odd", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := append(make([]int, 0, len(tt.in)), tt.in...)
			if tt.in == nil {
				s = nil
			}
			Reverse(s)
			assert.Equal(t, tt.out, s)
		})
	}
}
