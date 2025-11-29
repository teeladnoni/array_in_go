package piscine

func SolveSudoku(board *[9][9]rune) bool {
	// 1. Find Empty
	row, col, isEmpty := findEmptyCell(board)
	if !isEmpty {
		return true
	}

	// 2. Try Digits
	for digit := '1'; digit <= '9'; digit++ {
		if isValid(board, row, col, digit) {
			board[row][col] = digit
			if SolveSudoku(board) {
				return true
			}
			board[row][col] = '.' // Backtrack
		}
	}
	return false
}

func findEmptyCell(board *[9][9]rune) (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				return r, c, true
			}
		}
	}
	return -1, -1, false
}

func isValid(board *[9][9]rune, row, col int, digit rune) bool {
	// Check Row
	for c := 0; c < 9; c++ {
		if board[row][c] == digit {
			return false
		}
	}
	// Check Col
	for r := 0; r < 9; r++ {
		if board[r][col] == digit {
			return false
		}
	}
	// Check Box
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if board[startRow+r][startCol+c] == digit {
				return false
			}
		}
	}
	return true
}
