package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum = ", sum)

	for index, num := range nums {
		if index == 3 {
			fmt.Println(num) //prints the value at index 3
		}
	}

	m := map[string]int{"hello": 1, "bye": 2}
	for key, value := range m {
		fmt.Println("Key = ", key, " Value = ", value)
	}

	for _, character := range "helloworld" {
		fmt.Println("\n", character)
	}
}
