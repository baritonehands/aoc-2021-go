package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	var last int
	var cur int
	var cnt int
	for i, line := range lines {
		if i == 0 {
			last, _ = strconv.Atoi(line)
		} else {
			cur, _ = strconv.Atoi(line)

			if cur > last {
				cnt++
			}
			last = cur
		}
	}

	fmt.Println(cnt)

	var last3 [3]int
	var li int
	var startWindow bool
	lastSum := -1
	cnt = 0

	for i, line := range lines {
		if i == 0 {
			last, _ := strconv.Atoi(line)
			last3[li] = last
			li++
		} else {
			cur, _ = strconv.Atoi(line)

			last3[li] = cur
			li++
			if li == 3 {
				startWindow = true
				li = 0
			}

			if startWindow {
				sum := 0
				for _, v := range last3 {
					sum += v
				}
				if lastSum != -1 && sum > lastSum {
					cnt++
				}
				lastSum = sum
			}
		}
	}
	println(cnt)
}
