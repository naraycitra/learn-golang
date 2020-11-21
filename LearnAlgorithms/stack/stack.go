package main

import "fmt"

// Stack represent stack slice int
type Stack struct {
	datas []int
}

// Push add items to the end of stack
func (s *Stack) Push(i int) {
	s.datas = append(s.datas, i)
}

// Pop remove the last item and return it
func (s *Stack) Pop() int {
	l := len(s.datas) - 1
	toPop := s.datas[l]
	s.datas = s.datas[:l]
	return toPop
}

func main() {
	var myStack Stack
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack)
	fmt.Println("Pop:", myStack.Pop())
	fmt.Println(myStack)
}
