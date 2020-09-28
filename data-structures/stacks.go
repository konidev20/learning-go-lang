package main

//Stack will hold the stack
type Stack struct {
	items []int
}

//Push an item into the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop an item from the stack
func (s *Stack) Pop() int {
	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length]
	return toRemove
}
