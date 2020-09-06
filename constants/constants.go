package main

import (
	"fmt"
	"math"
)

//the const keyword is used to declare constant variables
const s string = "constant"

func main() {
	fmt.Println("s = ", s)

	const n = 500000
	const d = 3e20 / n

	//all numeric computations are performed with an arbitrary precision
	fmt.Println("d = ", d)

	//numeric constants don't have a type until they're given an explicit conversion
	fmt.Println("d (int64) = ", int64(d))

	//a const can be implicitly given a type using context
	fmt.Println("sin(n) = ", math.Sin(n))
	//math.Sin(n) expected float64, so before passing, it is implicitly converted into a float64
}
