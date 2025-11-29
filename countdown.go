package main

import "fmt"

func main() {
	// Start the recursion with 5
	CountDown(5)
}

func CountDown(n int) {
	// 1. The Base Case (STOP!)
	if n == 0 {
		return
	}

	// 2. Do the work (Print the number)
	fmt.Println(n)

	// 3. The Recursive Step (Call yourself with a smaller number)
	CountDown(n - 1)
}
