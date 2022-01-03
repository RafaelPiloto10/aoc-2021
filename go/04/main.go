package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	data   []int
	marked []bool
	didWin bool
}

func (b *Board) call(calledNum int) bool {
	for i, num := range b.data {
		if num == calledNum {
			b.marked[i] = true
			return true
		}
	}

	return false
}

func (b *Board) checkWin() bool {
	// check all cols
	for i := 0; i < 5; i++ {
		isWinningCol := true
		for j := 0; j < 5; j++ {
			isWinningCol = isWinningCol && b.marked[(j*5+i)]
		}
		if isWinningCol {
			return true
		}
	}
	// check all rows
	for i := 0; i < 5; i++ {
		isWinningRow := true
		for j := 0; j < 5; j++ {
			isWinningRow = isWinningRow && b.marked[(i*5+j)]
		}
		if isWinningRow {
			return true
		}
	}
	return false
}

func (b *Board) getSum() int {
	total := 0
	for i, has := range b.marked {
		if !has {
			total += b.data[i]
		}
	}

	return total
}

func main() {
	body, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	data := string(body)

	lines := strings.Split(data, "\n")
	nums := []int{}
	for _, n := range lines[2:] {
		if n == "" {
			continue
		}

		k := strings.Split(n, " ")
		for _, j := range k {
			if j == "" {
				continue
			}
			a, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}
			nums = append(nums, a)
		}
	}

	deck := strings.Split(lines[0], ",")
	boards := parseBoards(nums)

	for _, r := range deck {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		for _, board := range boards {
			if board.didWin {
				continue
			}

			if board.call(num) {
				if board.checkWin() {
					board.didWin = true
					sum := board.getSum()
					fmt.Printf("Winning board: %v with sum %v won on number %v\n", sum*num, sum, num)
				}
			}
		}

	}
}

func getRemainingBoards(boards []*Board) []*Board {
	returnBoards := []*Board{}
	for _, board := range boards {
		if !board.didWin {
			returnBoards = append(returnBoards, board)
		}
	}
	return returnBoards
}

func parseBoards(input []int) []*Board {
	boards := []*Board{}
	for i := 0; i < len(input); {
		row := []int{}
		for j := i; i < j+25; i++ {
			row = append(row, input[i])
		}
		boards = append(boards, NewBoard(row))
	}

	return boards
}

func NewBoard(data []int) *Board {
	board := Board{data: data}
	board.marked = make([]bool, 25)
	return &board
}
