package day16

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseNet(t *testing.T) {
	tests := []struct {
		in   string
		want net
	}{
		{
			"Valve AA has flow rate=0; tunnels lead to valves BB\n" +
				"Valve BB has flow rate=1; tunnels lead to valves AA\n",
			net{
				{0, "AA", 0, []string{"BB"}, []int{1}},
				{1, "BB", 1, []string{"AA"}, []int{0}},
			},
		},
		{
			"Valve AA has flow rate=0; tunnels lead to valves BB, CC\n" +
				"Valve BB has flow rate=1; tunnels lead to valves AA, CC\n" +
				"Valve CC has flow rate=2; tunnels lead to valves BB, AA\n",
			net{
				{0, "AA", 0, []string{"BB", "CC"}, []int{1, 2}},
				{1, "BB", 1, []string{"AA", "CC"}, []int{0, 2}},
				{2, "CC", 2, []string{"BB", "AA"}, []int{1, 0}},
			},
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, tt.want, parseNet(tt.in))
		})
	}
}

func Test_net_reduce(t *testing.T) {
	tests := []struct {
		n    net
		want reducedNet
	}{
		{
			net{
				{0, "AA", 0, nil, []int{1}},
				{1, "BB", 1, nil, []int{2}},
				{2, "CC", 2, nil, []int{1}},
			},
			reducedNet{
				[]int{1, 2},
				[][]int{
					{0, 1},
					{1, 0},
				},
				[]int{1, 2},
				[]string{"BB", "CC"},
			},
		},
		{
			net{
				{0, "AA", 0, nil, []int{1, 2, 3}},
				{1, "BB", 1, nil, []int{2, 3}},
				{2, "CC", 2, nil, []int{3}},
				{3, "DD", 3, nil, []int{2, 1}},
			},
			reducedNet{
				[]int{1, 2, 3},
				[][]int{
					{0, 1, 1},
					{2, 0, 1},
					{1, 1, 0},
				},
				[]int{1, 1, 1},
				[]string{"BB", "CC", "DD"},
			},
		},
		{
			net{
				{0, "AA", 0, nil, []int{1}},
				{1, "BB", 1, nil, []int{2}},
				{2, "CC", 0, nil, []int{3}},
				{3, "DD", 3, nil, []int{1}},
			},
			reducedNet{
				[]int{1, 3},
				[][]int{
					{0, 2},
					{1, 0},
				},
				[]int{1, 3},
				[]string{"BB", "DD"},
			},
		},
	}
	for idx, tt := range tests {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			assert.Equal(t, &tt.want, tt.n.reduce())
		})
	}
}

func Test_reducedNet_flow(t *testing.T) {
	rn := &reducedNet{rates: []int{1, 2, 4, 8}}
	for i := 0; i < 16; i++ {
		assert.Equal(t, i, rn.flow(int64(i)))
	}
	rn = &reducedNet{rates: []int{2, 4, 8, 16}}
	for i := 0; i < 16; i++ {
		assert.Equal(t, 2*i, rn.flow(int64(i)))
	}
}
