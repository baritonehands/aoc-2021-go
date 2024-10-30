package main

import (
	_ "embed"
	"fmt"
	"github.com/BooleanCat/go-functional/v2/it"
	"github.com/baritonehands/aoc-2021-go/utils"
	"maps"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	dirX = iota
	dirY
)

type Manual struct {
	dots  []Pair
	folds []Fold
}

type Fold struct {
	dir int
	n   int
}

func parseInput() Manual {
	ret := Manual{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "fold along ") {

			line = line[11:]
			parts := strings.Split(line, "=")

			dir := dirX
			if parts[0] == "y" {
				dir = dirY
			}
			n, _ := strconv.Atoi(parts[1])

			ret.folds = append(ret.folds, Fold{dir: dir, n: n})
		} else {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])

			ret.dots = append(ret.dots, Pair{x, y})
		}
	}
	return ret
}

func fold(dots []Pair, f Fold) []Pair {
	xMax, _ := it.Max(it.Map(slices.Values(dots), func(p Pair) int {
		return p.x
	}))
	yMax, _ := it.Max(it.Map(slices.Values(dots), func(p Pair) int {
		return p.y
	}))
	fmt.Printf("(%d, %d)\n", xMax, yMax)

	dotSet := utils.SeqSet(slices.Values(dots))

	ret := it.Exhausted[Pair]()
	for y := range yMax + 1 {
		for x := range xMax + 1 {
			_, found := dotSet[Pair{x, y}]

			if found {
				newPos := Pair{x, y}
				if f.dir == dirY && y > f.n {
					newPos.y = f.n + (f.n - y)
				} else if f.dir == dirX && x > f.n {
					newPos.x = f.n + (f.n - x)
				}
				ret = it.Chain(ret, it.Once(newPos))
			}
		}
	}
	return slices.Collect(maps.Keys(utils.SeqSet(ret)))
}

func printDots(dots []Pair) {
	xMax, _ := it.Max(it.Map(slices.Values(dots), func(p Pair) int {
		return p.x
	}))
	yMax, _ := it.Max(it.Map(slices.Values(dots), func(p Pair) int {
		return p.y
	}))

	dotSet := utils.SeqSet(slices.Values(dots))

	lines := slices.Collect(it.Map(it.Integers(0, yMax+1, 1), func(y int) string {
		marks := slices.Collect(it.Map(it.Integers(0, xMax+1, 1), func(x int) rune {
			if _, found := dotSet[Pair{x, y}]; found {
				return 'X'
			}

			return ' '
		}))
		return string(marks)
	}))
	println(strings.Join(lines, "\n"))
}

func main() {
	manual := parseInput()
	fmt.Printf("%v\n", manual)

	newDots := fold(manual.dots, manual.folds[0])
	println("part1", len(newDots))

	dots := manual.dots
	for _, f := range manual.folds {
		dots = fold(dots, f)
	}
	println("part2", len(dots))
	printDots(dots)
}
