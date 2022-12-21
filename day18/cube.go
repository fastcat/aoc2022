package day18

import (
	"fmt"

	"github.com/fastcat/aoc2022/i"
)

type cube [3]int

func parseOne(in string) cube {
	var c cube
	if _, err := fmt.Sscanf(in, "%d,%d,%d\n", &c[0], &c[1], &c[2]); err != nil {
		panic(err)
	}
	return c
}

func parseMany(in string) []cube {
	return i.ToSlice(i.Map(
		i.ToStrings(i.Split(i.Runes(in), []rune{'\n'})),
		i.NoIndex(parseOne),
	))
}

func (c cube) adjacentTo(c2 cube) bool {
	// adjacent if 2 coords are equal and one is off by one
	if c[0] == c2[0] && c[1] == c2[1] {
		return c[2] == c2[2]-1 || c[2] == c2[2]+1
	} else if c[1] == c2[1] && c[2] == c2[2] {
		return c[0] == c2[0]-1 || c[0] == c2[0]+1
	} else if c[2] == c2[2] && c[0] == c2[0] {
		return c[1] == c2[1]-1 || c[1] == c2[1]+1
	}
	return false
}

func countSharedFaces(scan []cube) int {
	n := 0
	for i, c1 := range scan {
		for j := i + 1; j < len(scan); j++ {
			if c1.adjacentTo(scan[j]) {
				n++
			}
		}
	}
	return n
}

func enclosedCells(scan []cube) []cube {
	// find the bounding box of the scan, and walk all cells adjacent starting
	// from a cell on the bounding box not part of the scan.
	coords := i.Map(i.Range(0, 3, 1), func(n, _ int) i.Iterable[int] {
		return i.Map(i.Slice(scan), func(c cube, _ int) int { return c[n] })
	})
	mins := i.ToSlice(i.Map(coords, func(axis i.Iterable[int], _ int) int {
		return i.Min(axis)
	}))
	maxes := i.ToSlice(i.Map(coords, func(axis i.Iterable[int], _ int) int {
		return i.Max(axis)
	}))

	inScan := i.ToMap(i.Slice(scan), func(c cube) (cube, bool) { return c, true })
	// start from a corner and walk recursively to find reachable cells
	reachable := make(map[cube]bool)
	var walk func(c cube)
	walk = func(c cube) {
		reachable[c] = true
		for n := 0; n < 3; n++ {
			c2 := c
			if c[n] >= mins[n] {
				c2[n] = c[n] - 1
				if !inScan[c2] && !reachable[c2] {
					walk(c2)
				}
			}
			if c[n] <= maxes[n] {
				c2[n] = c[n] + 1
				if !inScan[c2] && !reachable[c2] {
					walk(c2)
				}
			}
		}
	}
	// walk starting from every edge cell
	for n1 := 0; n1 < 3; n1++ {
		// n1 is an axis to hold at min or max, n2,n3 are the others to walk along
		// that face. we place the faces just outside the bounding box so we know
		// all the cells in each face are reachable and not in the scan.
		n2 := (n1 + 1) % 3
		n3 := (n1 + 2) % 3
		for a := mins[n2] - 1; a <= maxes[n2]+1; a++ {
			for b := mins[n3] - 1; b <= maxes[n3]+1; b++ {
				var c cube
				c[n1] = mins[n1] - 1
				c[n2] = a
				c[n3] = b
				walk(c)
				c[n1] = maxes[n1] + 1
				walk(c)
			}
		}
	}

	// for everything in range, if it's not reachable and it's not in the scan,
	// it's enclosed
	var enclosed []cube
	for x := mins[0]; x <= maxes[0]; x++ {
		for y := mins[1]; y <= maxes[1]; y++ {
			for z := mins[2]; z <= maxes[2]; z++ {
				c := cube{x, y, z}
				if !reachable[c] && !inScan[c] {
					enclosed = append(enclosed, c)
				}
			}
		}
	}

	return enclosed
}
