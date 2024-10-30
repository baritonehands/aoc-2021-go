package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

var brackets = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
	'<': '>',
}

var score = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func isOpen(ch rune) bool {
	_, ok := brackets[ch]
	return ok
}

func isClose(ch rune) bool {
	return !isOpen(ch)
}

func parseLine(line string) (bool, any) {
	open := list.List{}
	for _, ch := range line {
		if isOpen(ch) {
			open.PushBack(ch)
		} else {
			if value, ok := open.Back().Value.(rune); ok && brackets[value] == ch {
				open.Remove(open.Back())
			} else {
				return false, ch
			}
		}
	}
	return true, open
}

func main() {
	var sum int
	for _, line := range strings.Split(input, "\n") {
		success, result := parseLine(line)
		if errorCh, _ := result.(rune); !success {
			sum += score[errorCh]
		}
	}
	println(sum)

	scores := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		_, result := parseLine(line)
		if open, ok := result.(list.List); ok {
			closeBrackets := make([]rune, open.Len())
			for e := open.Back(); e != nil; e = e.Prev() {
				if ch, ok := e.Value.(rune); ok {
					closeBrackets = append(closeBrackets, brackets[ch])
				}
			}

			var score = 0
			for _, ch := range closeBrackets {
				score = (score * 5) + part2Score[ch]
			}
			scores = append(scores, score)
		}
	}
	slices.Sort(scores)
	fmt.Printf("part2: %v\n", scores[len(scores)/2])
}

var part2Score = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
