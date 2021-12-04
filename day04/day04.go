package day04

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	value  uint
	marked bool
}

type Board [5][5]Field

func Play(drawnNumbers []uint, boards []Board) uint {
	for _, drawn := range drawnNumbers {
		for idx, _ := range boards {
			boards[idx].mark(drawn)

			if boards[idx].checkWin() {
				sum := boards[idx].sumUnmarked()
				return sum * drawn
			}
		}
	}

	return 0
}

func PlayFile(file *os.File) uint {
	return Play(parseFile(file))
}

func PlayLast(drawnNumbers []uint, boards []Board) uint {
	calc := uint(0)
	var alreadyWonBoardIndexes []int

	for _, drawn := range drawnNumbers {
		for idx, _ := range boards {
			if contains(alreadyWonBoardIndexes, idx) {
				continue
			}

			boards[idx].mark(drawn)

			if boards[idx].checkWin() {
				sum := boards[idx].sumUnmarked()
				calc = sum * drawn
				alreadyWonBoardIndexes = append(alreadyWonBoardIndexes, idx)
			}
		}
	}

	return calc
}

func PlayLastFile(file *os.File) uint {
	return PlayLast(parseFile(file))
}

func (board *Board) mark(drawn uint) {
	for x, row := range board {
		for y, field := range row {
			if field.value == drawn {
				board[x][y].marked = true
			}
		}
	}
}

func (board *Board) checkWin() bool {
	// check horizontal
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			field := &board[x][y]
			if field.marked == false {
				break
			}
			if x == 4 {
				return true
			}
		}
	}

	// check vertical
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			field := &board[x][y]
			if field.marked == false {
				break
			}
			if y == 4 {
				return true
			}
		}
	}

	return false
}

func (board *Board) sumUnmarked() uint {
	sum := uint(0)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if board[x][y].marked == false {
				sum += board[x][y].value
			}
		}
	}
	return sum
}

func parseFile(file *os.File) ([]uint, []Board) {
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var boards []Board
	var drawnNumbers []uint

	if scanner.Scan() {
		splitDrawn := strings.Split(scanner.Text(), ",")
		for _, s := range splitDrawn {
			numb, _ := strconv.ParseUint(s, 10, 32)
			drawnNumbers = append(drawnNumbers, uint(numb))
		}
	}

	var currentBoard Board
	var row int
	for scanner.Scan() {
		scan := scanner.Text()
		if scan == "" {
			row = 0
			boards = append(boards, currentBoard)
			continue
		}
		rowOfStrings := strings.Fields(scan)
		for i, e := range rowOfStrings {
			numb, _ := strconv.ParseUint(e, 10, 32)
			currentBoard[i][row] = Field{uint(numb), false}
		}
		row++
	}
	return drawnNumbers, boards
}

func contains(x []int, y int) bool {
	for _, a := range x {
		if a == y {
			return true
		}
	}
	return false
}
