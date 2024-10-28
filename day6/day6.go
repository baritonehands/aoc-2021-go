package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func generate(gen map[int]int) map[int]int {
	ret := make(map[int]int)
	for days, cnt := range gen {
		switch days {
		case 0:
			ret[6] += cnt
			ret[8] += cnt
		default:
			ret[days-1] += cnt
		}

	}
	return ret
}

func main() {
	fish := strings.Split(input, ",")

	gen := make(map[int]int)
	for _, f := range fish {
		days, _ := strconv.Atoi(f)
		gen[days]++
	}

	var nextGen = gen
	for _ = range 256 {
		nextGen = generate(nextGen)
	}
	fmt.Printf("%v\n", nextGen)

	var total int
	for _, cnt := range nextGen {
		total += cnt
	}
	println(total)
}
