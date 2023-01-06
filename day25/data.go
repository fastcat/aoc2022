package day25

var dec = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var enc = map[int]rune{
	2:  '2',
	1:  '1',
	0:  '0',
	-1: '-',
	-2: '=',
}
