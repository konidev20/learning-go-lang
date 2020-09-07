package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["john"] = 10
	m["prucell"] = 15

	fmt.Println("map = ", m)

	//check if a key,value pair exists
	_, ok := m["jack"]

	//if ok is true, that means the value exists, else the key doesn't exists
	//value,ok
	if ok {
		fmt.Println("Not Present")
	} else {
		fmt.Println("Present")
	}

	//delete
	delete(m, "prucell")
	fmt.Println("m = ", m)

}
