package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("number")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := 10
	switch {
	case t < 5:
		fmt.Println("less than 5")
	default:
		fmt.Println("greater than 5")
	}

	//type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Println("Don't know %T", t)
		}
	}

	whatAmI(true)
	whatAmI(30)
}
