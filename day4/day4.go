package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Board struct {
	numbers   [5][5]int
	numberSet map[int]bool
}

func (board *Board) winningTotal(moves []int, curMove int) (int, error) {
	moveSet := make(map[int]bool)
	for i := 0; i <= curMove; i++ {
		moveSet[moves[i]] = true
	}
	var winner bool
	for row := range board.numbers {
		if board.isRowWinner(moveSet, row) {
			winner = true
			break
		}
	}
	if !winner {
		for col := range board.numbers {
			if board.isColWinner(moveSet, col) {
				winner = true
				break
			}
		}
	}
	if !winner && board.isDiagWinner(moveSet) {
		winner = true
	}

	if winner {
		var score int
		for v := range board.numberSet {
			if _, ok := moveSet[v]; !ok {
				score += v
			}
		}
		return score * moves[curMove], nil
	}
	return -1, errors.New("Game in progress")
}

func (board *Board) isRowWinner(moveSet map[int]bool, row int) bool {
	for _, v := range board.numbers[row] {
		if _, ok := moveSet[v]; !ok {
			return false
		}
	}
	return true
}

func (board *Board) isColWinner(moveSet map[int]bool, col int) bool {
	for _, row := range board.numbers {
		if _, ok := moveSet[row[col]]; !ok {
			return false
		}
	}
	return true
}

func (board *Board) isDiagWinner(moveSet map[int]bool) bool {
	var cnt = 0
	for i := range 5 {
		v := board.numbers[i][i]
		if _, ok := moveSet[v]; ok {
			cnt++
		}
	}
	if cnt == 5 {
		return true
	}

	cnt = 0
	for i := range 5 {
		v := board.numbers[5-i-1][i]
		if _, ok := moveSet[v]; ok {
			cnt++
		}
	}
	if cnt == 5 {
		return true
	}

	return false
}

func parseBoards(lines []string) []Board {
	ret := make([]Board, (len(lines)+1)/6)
	boardIdx := 0
	for li := 1; li < len(lines); li++ {
		var board = &ret[boardIdx]
		board.numberSet = make(map[int]bool)
		for row := range 5 {
			line := lines[li]
			numbers := strings.Fields(line)
			for col, num := range numbers {
				val, _ := strconv.Atoi(num)
				board.numbers[row][col] = val
				board.numberSet[val] = true
			}
			li++
		}
		boardIdx++
	}
	return ret
}

func parseMoves(input string) []int {
	movesStr := strings.Split(input, ",")
	moves := make([]int, len(movesStr))
	for i, m := range movesStr {
		moves[i], _ = strconv.Atoi(m)
	}
	return moves
}

func main() {
	l1, rest, _ := strings.Cut(input, "\n")

	boards := parseBoards(strings.Split(rest, "\n"))

	moves := parseMoves(l1)
	for moveIdx := range moves {
		found := false
		for _, board := range boards {
			total, err := board.winningTotal(moves, moveIdx)
			if err == nil {
				println(total)
				fmt.Printf("%v\n\n", board)
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	winnerSet := make(map[int]bool)
	var lastWinner *Board
	var lastTotal int
	for moveIdx := range moves {
		for boardIdx, board := range boards {
			if _, ok := winnerSet[boardIdx]; !ok {
				total, err := board.winningTotal(moves, moveIdx)
				if err == nil {
					winnerSet[boardIdx] = true
					lastWinner = &board
					lastTotal = total
				}
			}
		}
	}

	fmt.Printf("%v\n\n", lastWinner)
	println(lastTotal)
}
