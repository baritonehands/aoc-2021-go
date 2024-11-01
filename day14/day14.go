package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2021-go/utils"
	"iter"
	"maps"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func part1(start string, pairs map[string]string) {
	state := []rune(start)
	var freq map[rune]int64
	for step := 0; step < 10; step++ {
		elems := utils.FlatMap(utils.Partition(state, 2, 1), func(inner iter.Seq[rune]) iter.Seq[rune] {
			innerStr := string(slices.Collect(inner))
			if replacement, found := pairs[innerStr]; found {
				return it.Chain(it.Once(rune(innerStr[0])), it.Once(rune(replacement[0])))
			} else {
				return it.Once(rune(innerStr[0]))
			}
		})

		state = slices.Collect(elems)
		freq := utils.Frequencies(slices.Values(state))
		fmt.Printf("%v\n", freq)
	}

	var min int64 = math.MaxInt64
	var max int64 = math.MinInt64
	for _, cnt := range freq {
		if cnt > max {
			max = cnt
		}
		if cnt < min {
			min = cnt
		}
	}
	println("part1", max-min)
}

type Part2 struct {
	pairFreqs map[string]int64
	elemFreqs map[string]int64
}

func step2(pairs map[string]string, state *Part2) {
	initialPairFreqs := maps.Clone(state.pairFreqs)

	for pair, freq := range initialPairFreqs {
		react2(pairs, state, pair, freq)
	}
}

func react2(pairs map[string]string, ret *Part2, pair string, freq int64) {
	if replacement, ok := pairs[pair]; ok {
		l := pair[0]
		r := pair[1]
		ret.pairFreqs[pair] -= freq
		lrep := fmt.Sprintf("%c%s", rune(l), replacement)
		ret.pairFreqs[lrep] += freq
		repr := fmt.Sprintf("%s%c", replacement, rune(r))
		ret.pairFreqs[repr] += freq
		ret.elemFreqs[replacement] += freq
	}
}

func part2(start string, pairMap map[string]string) {
	pairFreqs := utils.Frequencies(it.Map(utils.Partition([]rune(start), 2, 1), func(seq iter.Seq[rune]) string {
		return string(slices.Collect(seq))
	}))
	elemFreqs := utils.Frequencies(it.Map(slices.Values([]rune(start)), func(ch rune) string {
		return string(ch)
	}))
	part2 := Part2{pairFreqs, elemFreqs}

	for i := 0; i < 40; i++ {
		step2(pairMap, &part2)
		//fmt.Printf("iter %d: %v\n\n", i, part2)
	}
	max, _ := it.Max(maps.Values(part2.elemFreqs))
	min, _ := it.Min(maps.Values(part2.elemFreqs))

	println(max - min)
}

func main() {
	lines := strings.Split(input, "\n")
	start := lines[0]

	rest := lines[2:]
	pairs := make(map[string]string, len(rest))
	for _, line := range rest {
		pairArr := strings.Split(line, " -> ")
		pairs[pairArr[0]] = pairArr[1]
	}

	fmt.Printf("%v\n%v\n", start, pairs)

	part2Pairs := maps.Clone(pairs)
	part1(start, pairs)
	part2(start, part2Pairs)
}
