package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func main() {
	//anonymous functions
	a := func(x int) int {
		return 3 * 14 * x
	}

	func() {
		fmt.Println("This is inside an anonymous function in Go.")
	}()

	fmt.Println("plus(a,b) = ", 10, 20)
	fmt.Println("a(45) = ", a(45))
}
