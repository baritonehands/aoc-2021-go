package main

import (
	"github.com/baritonehands/aoc-2021-go/utils"
	"maps"
	"slices"
)

type Digit map[rune]bool

func parseDigit(s string) Digit {
	ret := make(Digit)
	for _, c := range s {
		ret[c] = true
	}
	return ret
}

func makeDigit(runes ...rune) Digit {
	return parseDigit(string(runes))
}

func (d Digit) String() string {
	return string(slices.Sorted(maps.Keys(d)))
}

func IsOne(d Digit) bool {
	return len(d) == 2
}

func IsFour(d Digit) bool {
	return len(d) == 4
}

func IsSeven(d Digit) bool {
	return len(d) == 3
}

func IsEight(d Digit) bool {
	return len(d) == 7
}

func IsUnique(d Digit) bool {
	return IsOne(d) || IsFour(d) || IsSeven(d) || IsEight(d)
}

func (d Digit) setDifference(other Digit) Digit {
	return utils.SetDifference(d, other)
}
