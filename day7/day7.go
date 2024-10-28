package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	crabStrings := strings.Split(input, ",")

	crabs := make([]int, len(crabStrings))
	for i, s := range crabStrings {
		crab, _ := strconv.Atoi(s)
		crabs[i] = crab
	}

	var min = math.MaxInt64
	var max = math.MinInt64
	for _, v := range crabs {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	var minSum = math.MaxInt64
	for pos := min; pos <= max; pos++ {
		totals := make(map[int]int)
		for crabIdx, crab := range crabs {
			if crab > pos {
				totals[crabIdx] = crabs[crabIdx] - pos
			} else {
				totals[crabIdx] = pos - crabs[crabIdx]
			}
		}
		var sum = 0
		for _, v := range totals {
			sum += (v * (v + 1)) / 2
		}

		if sum < minSum {
			minSum = sum
		}
	}
	println(minSum)
}
