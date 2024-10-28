package main

import (
	_ "embed"
	"errors"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	lineLen := len(lines[0])
	gamma, epsilon := msbAndLsb(lines)

	var gammaBin int
	var epsilonBin int
	for i := 0; i < lineLen; i++ {
		if gamma[i] > epsilon[i] {
			gammaBin |= 1 << (lineLen - i - 1)
		} else {
			epsilonBin |= 1 << (lineLen - i - 1)
		}
	}
	println(gammaBin * epsilonBin)

	oxygen, _ := rating(lines, '1')
	co2, _ := rating(lines, '0')

	println(oxygen * co2)
}

func msbAndLsb(lines []string) ([]int, []int) {
	lineLen := len(lines[0])
	msb := make([]int, lineLen)
	lsb := make([]int, lineLen)

	for _, line := range lines {
		for i, ch := range line {
			if ch == '0' {
				lsb[i]++
			} else {
				msb[i]++
			}
		}
	}
	return msb, lsb
}

func rating(lines []string, precendence uint8) (int64, error) {
	var exclude map[int]bool
	lineLen := len(lines[0])
	var filteredLines = lines
	for i := 0; i < lineLen; i++ {
		msb, lsb := msbAndLsb(filteredLines)

		var excludeCh uint8
		if precendence == '1' {
			excludeCh = '0'
		} else {
			excludeCh = '1'
		}

		if msb[i] > lsb[i] {
			exclude = excludeIndexes(lines, exclude, i, excludeCh)
		} else if lsb[i] > msb[i] {
			exclude = excludeIndexes(lines, exclude, i, precendence)
		} else {
			exclude = excludeIndexes(lines, exclude, i, excludeCh)
		}
		filteredLines = make([]string, 0, len(lines)-len(exclude))
		for i, line := range lines {
			if !exclude[i] {
				filteredLines = append(filteredLines, line)
			}
		}

		if len(filteredLines) == 1 {
			return strconv.ParseInt(filteredLines[0], 2, 64)
		}
	}

	return 0, errors.New("No rating found")
}

func excludeIndexes(lines []string, exclude map[int]bool, idx int, ch uint8) map[int]bool {
	ret := map[int]bool{}
	for i, line := range lines {
		if exclude[i] == true {
			ret[i] = true
		} else {
			if line[idx] == ch {
				ret[i] = true
			}
		}
	}
	return ret
}
