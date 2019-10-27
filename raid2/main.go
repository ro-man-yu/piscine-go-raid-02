package main

import (
	"fmt"
	"os"
)

func main() {

	if !IsValidArgs(os.Args[1:]) {
		fmt.Println("Error")
		return
	}
	board := parseBoard(os.Args)
	board2 := parseBoard(os.Args)
	if SolveSudoku(&board, false) {
		if SolveSudoku(&board2, true) {
			if CompareBoards(&board, &board2) {
				Display(board)
			} else {
				fmt.Println("Error")
			}
		}
	} else {
		fmt.Println("Error")
	}
}

func IsValidArgs(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for _, arg := range args {
		for _, letter := range arg {
			if (letter < '0' || letter > '9') && letter != '.' {
				return false
			}
		}
	}

	return true
}

func parseBoard(arg []string) [9][9]int {
	var board = [9][9]int{}
	for j := 1; j < 10; j++ {
		for index, letter := range os.Args[j] {
			if letter == '.' {
				letter = 48
			}
			board[j-1][index] = RuneToInt(letter)
		}
	}
	return board
}

func RuneToInt(r rune) int {
	var n = 0
	for i := '0'; i <= '9'; i++ {
		if i == r {
			return n
		}
		n++
	}
	return n
}

func Display(arg [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(arg[i][j])
			if j < 8 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func SolveSudoku(board *[9][9]int, isReverse bool) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				if isReverse {
					for candidate := 9; candidate >= 1; candidate-- {
						board[i][j] = candidate
						if isBoardValid(board) {
							if SolveSudoku(board, true) {
								return true
							}
							board[i][j] = 0
						} else {
							board[i][j] = 0
						}
					}
				} else {
					for candidate := 1; candidate <= 9; candidate++ {
						board[i][j] = candidate
						if isBoardValid(board) {
							if SolveSudoku(board, false) {
								return true
							}
							board[i][j] = 0
						} else {
							board[i][j] = 0
						}
					}
				}

				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func CompareBoards(board1 *[9][9]int, board2 *[9][9]int) bool {

	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			if board1[i][j] != board2[i][j] {
				return false
			}
		}
	}
	return true
}
