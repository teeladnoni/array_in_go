package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// 1. Get arguments
	args := os.Args[1:]

	// 2. Validate input length
	if len(args) != 9 {
		printError()
		return
	}

	// 3. Create and Parse the Board
	var board [9][9]rune
	for i, rowStr := range args {
		if len(rowStr) != 9 {
			printError()
			return
		}
		for j, char := range rowStr {
			if (char < '1' || char > '9') && char != '.' {
				printError()
				return
			}
			board[i][j] = char
		}
	}

	// 4. Solve and Print
	if SolveSudoku(&board) {
		printBoard(board)
	} else {
		printError()
	}
}

func printError() {
	errStr := "Error\n"
	for _, r := range errStr {
		z01.PrintRune(r)
	}
}

func printBoard(board [9][9]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(board[i][j])
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
