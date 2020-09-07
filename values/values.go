package main

import "fmt"

func main() {
	//string concatenation
	fmt.Println("Hello" + "World")

	//Integers and floats
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1.0 / 2.0 = ", 1.0/2.0)

	//Boolean
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	a, b := 10, 20
	fmt.Println("a,b =", a, ",", b)
	a, b := swap(a, b)
	fmt.Println("a,b =", a, ",", b)
}

//how to return multiple values in Go
func swap(a int, b int) (int, int) {
	return b, a
}
