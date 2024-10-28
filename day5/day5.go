package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

type Vector struct {
	from, to Point
}

func (v *Vector) points() []Point {
	if v.from.x == v.to.x {
		var startY, endY int
		if v.from.y > v.to.y {
			startY = v.to.y
			endY = v.from.y
		} else {
			startY = v.from.y
			endY = v.to.y
		}
		ret := make([]Point, endY-startY+1)
		i := 0
		for dy := startY; dy <= endY; dy++ {
			ret[i] = Point{v.from.x, dy}
			i++
		}
		return ret
	} else if v.from.y == v.to.y {
		var startX, endX int
		if v.from.x > v.to.x {
			startX = v.to.x
			endX = v.from.x
		} else {
			startX = v.from.x
			endX = v.to.x
		}
		ret := make([]Point, endX-startX+1)
		i := 0
		for dx := startX; dx <= endX; dx++ {
			ret[i] = Point{dx, v.from.y}
			i++
		}
		return ret
	} else {
		return []Point{}
	}
}

func parseInput() []Vector {
	lines := strings.Split(input, "\n")
	ret := make([]Vector, len(lines))
	for i, line := range lines {
		entries := strings.Split(line, " -> ")
		from := strings.Split(entries[0], ",")
		to := strings.Split(entries[1], ",")

		fromXInt, _ := strconv.Atoi(from[0])
		fromYInt, _ := strconv.Atoi(from[1])
		toXInt, _ := strconv.Atoi(to[0])
		toYInt, _ := strconv.Atoi(to[1])

		vector := &ret[i]
		vector.from.x = fromXInt
		vector.from.y = fromYInt
		vector.to.x = toXInt
		vector.to.y = toYInt
	}
	return ret
}

func main() {
	vectors := parseInput()
	//fmt.Printf("%v\n\n", vectors)
	//fmt.Printf("%v\n", vectors[3].points())

	freq := make(map[Point]int)
	for _, vector := range vectors {
		for _, point := range vector.points() {
			freq[point]++
		}
	}
	var total int
	for _, count := range freq {
		if count > 1 {
			total++
		}
	}
	fmt.Printf("%v\n", total)
}
