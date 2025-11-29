package main

import "fmt"

func main() {
	// A 3x3 Grid (3 Rows, 3 Columns)
	// We call this a "Multidimensional Array"
	var grid [3][3]int

	// Put 999 in the exact center (Row 1, Column 1)
	grid[1][1] = 999

	// TASK: Put 777 in the bottom-right corner (Row 2, Column 2)
	// Write your code here:
	grid[2][2] = 777

	// Let's print the center to prove it works
	fmt.Print("Center is: ")
	fmt.Println(grid[1][1])

	// Let's print the corner to see if you succeeded
	fmt.Print("Bottom Right is: ")
	fmt.Println(grid[2][2])
}
