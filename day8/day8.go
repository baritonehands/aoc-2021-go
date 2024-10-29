package main

import (
	_ "embed"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/BooleanCat/go-functional/v2/it/itx"
	"github.com/baritonehands/aoc-2021-go/utils"
	"iter"
	"maps"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Signal [10]Digit

type Output [4]Digit

func digitString(digits []Digit) string {
	sorted := slices.Sorted(it.Map(slices.Values(digits), func(d Digit) string {
		return d.String()
	}))
	return strings.Join(sorted, " ")
}

func RuneSeq(digits []Digit) iter.Seq[rune] {
	return utils.FlatMap(slices.Values(digits), func(d Digit) iter.Seq[rune] {
		return maps.Keys(d)
	})
}

type InputLine struct {
	signal Signal
	output Output
}

func (i InputLine) String() string {
	parts := []string{digitString(i.signal[:]), digitString(i.output[:])}
	return strings.Join(parts, " | ")
}

type Input []InputLine

func (i Input) String() string {
	return strings.Join(
		slices.Collect(it.Map(slices.Values(i), func(l InputLine) string {
			return l.String()
		})),
		"\n")
}

func (i Input) SignalSeq() iter.Seq[Signal] {
	return it.Map(slices.Values(i), func(i InputLine) Signal {
		return i.signal
	})
}

func parseDigits(digits []string) []Digit {
	return slices.Collect(it.Map(slices.Values(digits), parseDigit))
}

func parseInput() Input {
	lines := strings.Split(input, "\n")
	ret := make(Input, len(lines))
	for i, line := range lines {
		row := strings.Split(line, " | ")
		rowInput := &ret[i]
		rowInput.signal = Signal(parseDigits(strings.Split(row[0], " ")))
		rowInput.output = Output(parseDigits(strings.Split(row[1], " ")))
	}
	//fmt.Printf("%v\n", ret)
	return ret
}

func signalKey(s Signal) map[string]int {
	counts := utils.Frequencies(RuneSeq(s[:]))
	one, _ := it.Find(slices.Values(s[:]), IsOne)
	four, _ := it.Find(slices.Values(s[:]), IsFour)
	seven, _ := it.Find(slices.Values(s[:]), IsSeven)
	eight, _ := it.Find(slices.Values(s[:]), IsEight)
	//fmt.Printf("%v\n", counts)

	six, _ := itx.FromSlice(s[:]).Find(func(digit Digit) bool {
		return len(digit) == 6 && len(one.setDifference(digit)) > 0
	})

	a := rune(seven.setDifference(one).String()[0])
	b, _, _ := it.Find2(maps.All(counts), func(r rune, cnt int) bool {
		return cnt == 6
	})
	c := rune(one.setDifference(six).String()[0])
	e, _, _ := it.Find2(maps.All(counts), func(r rune, cnt int) bool {
		return cnt == 4
	})
	f, _, _ := it.Find2(maps.All(counts), func(r rune, cnt int) bool {
		return cnt == 9
	})
	gExclude := makeDigit(a, b, c, e, f)
	g := rune(eight.setDifference(gExclude).setDifference(four).String()[0])
	d, _ := it.Find(maps.Keys(counts), func(r rune) bool {
		return r != a && r != b && r != c && r != e && r != f && r != g
	})

	zero := makeDigit(a, b, c, e, f, g)
	two := makeDigit(a, c, d, e, g)
	three := makeDigit(a, c, d, f, g)
	five := makeDigit(a, b, d, f, g)
	nine := makeDigit(a, b, c, d, f, g)

	//fmt.Printf("%v\n", []any{zero, one, two, three, four, five, six, seven, eight, nine})
	//fmt.Printf("%v\n", []any{a, b, c, d, e, f, g})

	return map[string]int{
		zero.String():  0,
		one.String():   1,
		two.String():   2,
		three.String(): 3,
		four.String():  4,
		five.String():  5,
		six.String():   6,
		seven.String(): 7,
		eight.String(): 8,
		nine.String():  9,
	}
}

func part2(inputs Input) {
	var sum int
	for _, input := range inputs {
		key := signalKey(input.signal)
		numberStr := string(slices.Collect(it.Map(slices.Values(input.output[:]), func(d Digit) rune {
			return rune(key[d.String()] + '0')
		})))
		number, _ := strconv.Atoi(numberStr)
		sum += number
	}
	println("part2", sum)
}

func main() {
	inputs := parseInput()

	var part1 int
	for _, inp := range inputs {
		for _, digit := range inp.output {
			if IsUnique(digit) {
				part1++
			}
		}
	}
	println("part1", part1)

	part2(inputs)
}
