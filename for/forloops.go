package main

import "fmt"

func main() {

	//variant 1
	i := 1
	for i <= 3 {
		i = i + 1
	}
	fmt.Println(i)

	//variant 2
	for j := 1; j <= 9; j++ {
		fmt.Println(j)
	}

	//variant 3
	for {
		fmt.Println("This loop will execute only once.")
		break
	}

	//variant 2 with continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
