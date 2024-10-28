package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/baritonehands/aoc-2021-go/utils"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	var pos, depth int
	for _, line := range lines {
		cmd, cntStr := utils.Split2(line)
		cnt, _ := strconv.Atoi(cntStr)

		switch cmd {
		case "forward":
			pos += cnt
		case "down":
			depth += cnt
		case "up":
			depth -= cnt
		case "backward":
			pos -= cnt
		}
	}

	println(pos * depth)

	var aim int
	pos = 0
	depth = 0
	for _, line := range lines {
		cmd, cntStr := utils.Split2(line)
		cnt, _ := strconv.Atoi(cntStr)

		switch cmd {
		case "forward":
			pos += cnt
			depth += aim * cnt
		case "down":
			aim += cnt
		case "up":
			aim -= cnt
		}
	}

	println(pos * depth)
}
