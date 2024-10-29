package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Signal [10]Digit

type Output [4]Digit

type InputLine struct {
	signal Signal
	output Output
}

func (i InputLine) String() string {
	parts := []string{fmt.Sprint(i.signal), fmt.Sprint(i.output)}
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

func parseDigit(s string) Digit {
	ret := make(Digit)
	for _, c := range s {
		ret[c] = true
	}
	return ret
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
	fmt.Printf("%v\n", ret)
	return ret
}

func main() {
	inputs := parseInput()

	var part1 int
	for _, inp := range inputs {
		for _, digit := range inp.output {
			if digit.isUnique() {
				part1++
			}
		}
	}
	println(part1)

}
