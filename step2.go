package main

import "fmt"

func main() {
	// 1. Create a dummy board with just ONE dot to test our logic
	var board [9][9]rune

	// Fill it with '1's so it's not empty
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = '1'
		}
	}

	// 2. Put a SINGLE dot at Row 0, Column 2
	board[5][5] = '.'

	// 3. Call our scanner function
	row, col, found := FindFirstDot(&board)

	if found {
		fmt.Printf("Found an empty spot at: Row %d, Col %d\n", row, col)
	} else {
		fmt.Println("The board is full!")
	}
}

// This is the function you need to learn
// It scans the board looking for '.'
func FindFirstDot(board *[9][9]rune) (int, int, bool) {
	// Loop through rows
	for i := 0; i < 9; i++ {
		// Loop through columns
		for j := 0; j < 9; j++ {
			// Check if we found a dot
			if board[i][j] == '.' {
				// STOP! Return the location immediately
				return i, j, true
			}
		}
	}
	// If we finish both loops, no dots were found
	return -1, -1, false
}
