package main

import "fmt"

//Stack implementation
type Stack struct {
	Arr []int
}

func (s *Stack) Push(i int) {
	s.Arr = append(s.Arr, i)
}

func (s *Stack) Pop() int {
	l := len(s.Arr) - 1
	last := s.Arr[l]
	s.Arr = s.Arr[:l]
	return last
}

//Queue Implementation
type Queue struct {
	Arr []int
}

func (s *Queue) Push(i int) {
	s.Arr = append(s.Arr, i)
}

func (s *Queue) Pop() int {
	first := s.Arr[0]
	s.Arr = s.Arr[1:]
	return first
}

func main() {
	// stack := Stack{}
	// fmt.Println(stack)
	// stack.Push(200)
	// stack.Push(400)
	// stack.Push(500)
	// fmt.Println(stack)
	// fmt.Println(stack.Pop())
	// fmt.Println(stack)
	q := Queue{}
	fmt.Println(q)
	q.Push(200)
	q.Push(400)
	q.Push(500)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q)
}
