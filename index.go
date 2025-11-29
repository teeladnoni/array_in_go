package main

import "fmt"

func main() {
	// 1. Build a row of 3 lockers (Indices: 0, 1, 2)
	var myNumbers [3]int

	// 2. Put numbers inside
	myNumbers[0] = 500
	myNumbers[1] = 20
	myNumbers[2] = 7

	// 3. Print the first locker
	fmt.Print("Locker 0 contains: ")
	fmt.Println(myNumbers[0])

	// 4. Print the last locker
	fmt.Print("Locker 2 contains: ")
	fmt.Println(myNumbers[2])

	fmt.Print("middle Locker contains: ")
	fmt.Println(myNumbers[1])
}
