package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	pq "github.com/baritonehands/aoc-2021-go/utils/priority_queue"
	"math"
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

func nyDistance(start Pair, end Pair) int {
	return int(math.Abs(float64(end.y-start.y)) +
		math.Abs(float64(end.x-start.x)))
}

func stateString(state [][]int) string {
	sb := strings.Builder{}
	it.ForEach(slices.Values(state), func(row []int) {
		for _, col := range row {
			sb.WriteRune(rune(col + '0'))
		}
		sb.WriteString("\n")
		//sb.WriteString(fmt.Sprintf("%v\n", row))
	})
	return sb.String()
}

func walkPath(cameFrom map[Pair]Pair, current Pair) []Pair {
	ret := []Pair{current}
	for {
		if next, ok := cameFrom[current]; ok {
			current = next
			ret = append(ret, current)
		} else {
			break
		}
	}
	return ret
}

func safestPath(risks map[Pair]int, start Pair, end Pair) []Pair {
	weight := func(pair Pair) int { return risks[pair] }
	fScore := map[Pair]int{start: nyDistance(start, end)}
	fScoreFn := func(pair Pair) int { return fScore[pair] }
	openSet := pq.NewQueue[int, Pair](fScoreFn, start)

	cameFrom := map[Pair]Pair{}
	gScore := map[Pair]int{start: 0}
	//fmt.Printf("%v\n", []any{weight(start), fScore, openSet, cameFrom, gScore})

	for {
		if openSet.Len() == 0 {
			panic("Shouldn't happen")
		}

		current := openSet.Peek()

		if current == end {
			// Walk path
			return walkPath(cameFrom, current)
		} else {
			openSet.Poll()

			// For each neighbor of current
			for _, neighbor := range current.Neighbors(end.x, end.y) {
				g := gScore[current] + weight(neighbor)
				gNeighbor, found := gScore[neighbor]
				if !found || g < gNeighbor {
					fScore[neighbor] = g + nyDistance(neighbor, end)

					cameFrom[neighbor] = current
					gScore[neighbor] = g
					openSet.Append(neighbor)
				}
			}
		}

	}

	panic("Shouldn't happen")
}

func part1(state [][]int) int {
	risks := make(map[Pair]int)
	for ri, row := range state {
		for ci, col := range row {
			risks[Pair{ci, ri}] = col
		}
	}
	//fmt.Printf("%v\n", risks)

	yMax := len(state) - 1
	xMax := len(state[0]) - 1

	path := safestPath(risks, Pair{0, 0}, Pair{xMax, yMax})
	fmt.Printf("%v\n", path)
	return it.Fold2(it.Drop2(slices.Backward(path), 1), func(risk int, idx int, pair Pair) int {
		return risk + risks[pair]
	}, 0)
}

func part2(state [][]int) int {
	wider := make([][]int, len(state))
	for ri, row := range state {
		wider[ri] = make([]int, len(row)*5)
		for mx := range 5 {
			for ci, col := range row {
				nextCol := col + mx
				if nextCol > 9 {
					nextCol -= 9
				}
				wider[ri][mx*len(row)+ci] = nextCol
			}
		}
	}

	taller := make([][]int, len(wider)*5)
	for my := range 5 {
		for ri, row := range wider {
			taller[my*len(state)+ri] = make([]int, len(row))
			for ci, col := range row {
				nextCol := col + my
				if nextCol > 9 {
					nextCol -= 9
				}
				taller[my*len(state)+ri][ci] = nextCol
			}
		}
	}

	fmt.Printf("%v\n", stateString(taller))

	return part1(taller)
}

func main() {
	state := parseInput()
	println("part1", part1(state))

	fmt.Printf("%v\n", stateString(state))
	println("part2", part2(state))
}
