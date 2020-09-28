package main

import "fmt"

func main() {
	//Stack Data Structure
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(100)
	myStack.Push(200)
	fmt.Println(myStack)
	fmt.Println("Popped : ", myStack.Pop())
	fmt.Println(myStack)

	//Queue Data Structure
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(111)
	myQueue.Enqueue(112)
	myQueue.Enqueue(113)
	myQueue.Enqueue(114)
	fmt.Println(myQueue)
	fmt.Println("Dequeued : ", myQueue.Dequeue())
	fmt.Println(myQueue)
}
