package main

import "fmt"

func main() {
	var board [9][9]rune

	// Setup: Fill board with dots
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = '.'
		}
	}

	// SCENARIO:
	// We put a '5' in the top-left corner (0,0)
	board[0][0] = '5'

	// TEST:
	// We try to put another '5' in the SAME 3x3 box, at (1, 1).
	// This should be ILLEGAL (False).
	if CanPut(&board, 1, 1, '5') {
		fmt.Println("Check Result: YES, you can put '5' at 1,1 (WRONG!)")
	} else {
		fmt.Println("Check Result: NO, you cannot put '5' at 1,1 (CORRECT!)")
	}
}

func CanPut(board *[9][9]rune, row, col int, digit rune) bool {
	// 1. Check the Row (Horizontal)
	for c := 0; c < 9; c++ {
		// If we find the digit already exists in this row...
		if board[row][c] == digit {
			return false // Forbidden!
		}
	}

	// 2. Check the Column (Vertical)
	for r := 0; r < 9; r++ {
		// If we find the digit already exists in this column...
		if board[r][col] == digit {
			return false // Forbidden!
		}
	}

	// 3. Check the 3x3 Box (The Hard Part)
	// We need to find the top-left corner of the box where 'row, col' lives.
	startRow := row / 3 * 3
	startCol := col / 3 * 3

	// Your Code:
	for r := 0; r <= 2; r++ {
		for c := 0; c <= 2; c++ {
			if board[startRow+r][startCol+c] == digit {
				return false
			}
		}
	}

	return true
}
