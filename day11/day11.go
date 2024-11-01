package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2021-go/utils"
	"iter"
	"maps"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	ret := make([][]int, len(lines))
	for ri, line := range lines {
		row := make([]int, len(line))
		for ci, n := range line {
			row[ci] = int(n - '0')
		}
		ret[ri] = row
	}
	return ret
}

func neighbors(p Pair) iter.Seq[Pair] {
	ret := it.Exhausted[Pair]()
	for ri := p.y - 1; ri <= p.y+1; ri++ {
		for ci := p.x - 1; ci <= p.x+1; ci++ {
			if (ri != p.y || ci != p.x) &&
				(ri >= 0 && ri < 10) &&
				(ci >= 0 && ci < 10) {
				ret = it.Chain(ret, it.Once(Pair{ci, ri}))
			}
		}
	}
	return ret
}

func updateState(state [][]int, updateFn func(x, y int, cell *int)) {
	for ri, row := range state {
		for ci := range row {
			updateFn(ci, ri, &state[ri][ci])
		}
	}
}

func incState(state [][]int, pairs map[Pair]int64) {
	updateState(state, func(x, y int, cell *int) {
		var inc = pairs[Pair{x, y}]
		if pairs == nil {
			inc = 1
		}
		*cell += int(inc)
	})
}

func resetState(state [][]int) {
	updateState(state, func(x, y int, cell *int) {
		if *cell > 9 {
			*cell = 0
		}
	})
}

func filterState(state [][]int, filterFn func(cell int) bool) map[Pair]bool {
	pairs := utils.FlatMap2(slices.All(state), func(ri int, row []int) iter.Seq2[Pair, bool] {
		flashingCols := it.Filter2(slices.All(row), func(ci int, col int) bool {
			return filterFn(col)
		})
		return it.Map2(flashingCols, func(ci int, col int) (Pair, bool) {
			return Pair{ci, ri}, true
		})
	})

	return maps.Collect(pairs)
}

func flashing(state [][]int) map[Pair]bool {
	return filterState(state, func(cell int) bool {
		return cell > 9
	})
}

func stateString(state [][]int) string {
	sb := strings.Builder{}
	it.ForEach(slices.Values(state), func(row []int) {
		sb.WriteString(fmt.Sprintf("%v\n", row))
	})
	return sb.String()
}

func flash(state [][]int, count *int) bool {
	incState(state, nil)
	//fmt.Printf("incremented: \n%v\n", stateString(state))
	flashed := map[Pair]bool{}

	for {
		//fmt.Printf("flashed: %v\n", flashed)
		toFlash := utils.SetDifference(flashing(state), flashed)
		//fmt.Printf("toFlash: %v\n", toFlash)

		if len(toFlash) == 0 {
			// Reset flash
			//println()
			break
		}

		updates := utils.Frequencies(utils.FlatMap(maps.Keys(toFlash), neighbors))
		//fmt.Printf("updates: %v\n", updates)
		incState(state, updates)
		//fmt.Printf("nextState: \n%v\n", stateString(state))
		maps.Insert(flashed, maps.All(toFlash))
		//println()
	}

	// Reset flashed items and set count
	resetState(state)
	*count += len(filterState(state, func(cell int) bool {
		return cell == 0
	}))

	if len(flashed) == 100 {
		return true
	}
	return false
	//println(stateString(state))
}

func part1(input [][]int, steps int) int {
	var flashes int
	state := slices.Collect(it.Map(slices.Values(input), func(row []int) []int {
		ret := make([]int, len(row))
		copy(ret, row)
		return ret
	}))
	fmt.Printf("%v\n", state)
	for step := 1; step < steps; step++ {
		allFlashed := flash(state, &flashes)
		if step == 100 {
			println("part1", flashes)
		}
		if allFlashed {
			println("part2", step)
			break
		}
	}
	return flashes
}

func main() {
	fmt.Printf("%v\n", slices.Collect(neighbors(Pair{2, 0})))
	part1(parseInput(), 1000)
}
