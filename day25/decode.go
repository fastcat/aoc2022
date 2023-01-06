package day25

import "fmt"

func decode(in string) int {
	v := 0
	for _, c := range in {
		if cv, ok := dec[c]; !ok {
			panic(fmt.Errorf("invalid char '%c' in '%s'", c, in))
		} else {
			v = v*5 + cv
		}
	}
	return v
}
