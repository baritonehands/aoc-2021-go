package main

import (
	"maps"
	"strings"
)

type Digit map[rune]bool

func (d Digit) String() string {
	b := strings.Builder{}
	for r := range maps.Keys(d) {
		b.WriteRune(r)
	}
	return b.String()
}

func (d Digit) isOne() bool {
	return len(d) == 2
}

func (d Digit) isFour() bool {
	return len(d) == 4
}

func (d Digit) isSeven() bool {
	return len(d) == 3
}

func (d Digit) isEight() bool {
	return len(d) == 7
}

func (d Digit) isUnique() bool {
	return d.isOne() || d.isFour() || d.isSeven() || d.isEight()
}

func (d Digit) setDifference(other Digit) Digit {
	var ret Digit
	for c, v := range d {
		_, present := other[c]
		if v && !present {
			ret[c] = true
		}
	}
	return ret
}
