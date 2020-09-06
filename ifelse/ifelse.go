package main

import "fmt"

func main() {
	x := 20
	if x%2 == 0 {
		fmt.Println("Odd number")
	} else {
		fmt.Println("Even Number")
	}

	//the declared variable y is accessible in all branches
	if y := 3; y < 10 {
		fmt.Println("y is lesser")
	} else if y > 20 {
		fmt.Println("y is great")
	} else {
		fmt.Println("y is equal")
	}
}
