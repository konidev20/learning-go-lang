package main

import (
	"fmt"
)

func main() {
	//Arrays
	var a [5]int //creating an array which can store 5 elements of type int
	fmt.Println(a)

	fmt.Println("Length = ", len(a))

	b := [5]int{1, 2, 3, 4, 5} //another way of declaring and initializing an array
	fmt.Println("b = ", b)

	//creating a 2-dimensional array
	var twoDarray [4][4]int
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			twoDarray[i][j] = i + j
		}
	}
	fmt.Println("twoDarray = ", twoDarray)

	//Slices
	s := make([]string, 10) //creating a string sice of size 3
	fmt.Println("s = ", s)

	s[0] = "abc"
	s[1] = "def"
	s[2] = "ghi"
	s[3] = "jkl"

	fmt.Println("set : ", s)
	fmt.Println("get : ", s[2])

	s = append(s, "mno")
	s = append(s, "pqr", "stu")

	fmt.Println(s)

	//copying slices
	c := make([]string, len(s))
	copy(c, s) //copies s into c
	fmt.Println("c = ", c)

	//accessing a slice
	l := c[:5] //extracting the first 4 elements and assigning it to l
	fmt.Println("l = ", l)

	m := c[3:5] //extract the 3rd and 4th element
	fmt.Println("m  = ", m)
	sli := []string{"s", "l", "i", "c", "e"}
	fmt.Println("sli = ", sli)

	twoDslice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDslice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDslice[i][j] = i + j
		}
	}

	fmt.Println(twoDslice)

}
