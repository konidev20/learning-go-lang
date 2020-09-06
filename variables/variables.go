package main

import "fmt"

func main() {
	//define one or more variable using var keyword
	var a = 10
	fmt.Println("a = ", a)

	//go will usually infer the data type, but we can explicitly mention the data type
	var b int = 20
	fmt.Println("b = ", b)

	//define more than one variable at a time
	var c, d int = 30, 40
	fmt.Println("c,d = ", c, ",", d)

	//variables that have not been initialized take the default value
	var e int
	fmt.Println("e takes the default value e = ", e)

	//:= is used to declare and initializing the variable
	f := "hello"
	//equivalent to var f string = "hello"
	fmt.Println("f = ", f)
}
