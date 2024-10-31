package main

import "fmt"

type Pair struct {
	x, y int
}

func (p Pair) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Pair) Neighbors(xMax, yMax int) []Pair {
	ret := make([]Pair, 0, 4)
	if p.x < xMax {
		ret = append(ret, Pair{p.x + 1, p.y})
	}
	if p.y < yMax {
		ret = append(ret, Pair{p.x, p.y + 1})
	}
	if p.x > 0 {
		ret = append(ret, Pair{p.x - 1, p.y})
	}
	if p.y > 0 {
		ret = append(ret, Pair{p.x, p.y - 1})
	}
	return ret
}
