package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Get the arguments (The 9 rows provided by the user)
	// We skip index 0 because that is the program name.
	arguments := os.Args[1:]

	// Safety Check: Did the user provide exactly 9 rows?
	if len(arguments) != 9 {
		fmt.Println("Error: Please provide exactly 9 rows.")
		return
	}

	// 2. Create the Empty Grid (9x9)
	var board [9][9]rune

	// 3. FILL THE BOARD
	// Loop 1: Go through each string (Row)
	for i, rowString := range arguments {

		// Safety Check: Is this row exactly 9 characters long?
		if len(rowString) != 9 {
			fmt.Println("Error: A row is the wrong length.")
			return
		}

		// Loop 2: Go through each character (Column) inside that string
		for j, char := range rowString {
			// COPY the character from the string into our grid
			board[i][j] = char
		}
	}

	// 4. VERIFY: Print the board to verify we captured the data correctly
	fmt.Println("Success! Here is what we stored in memory:")
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

// testing : {
// 	go run step1.go "1.3......" "4...5...." "........." "........." "........." "........." "........." "........." "........."
// }
