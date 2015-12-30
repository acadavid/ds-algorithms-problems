// Implement a Queue using two stacks.

package main

import "fmt"

type stack []int
func (s stack) Empty() bool { return len(s) == 0 }
func (s stack) Peek() int   { return s[len(s)-1] }
func (s *stack) Put(i int)  { (*s) = append((*s), i) }
func (s *stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

type staqueue struct {
	s1 stack
	s2 stack
}

func (s *staqueue) Push(i int) { s.s1.Put(i) }
func (s *staqueue) Pop() int {
	for len(s.s1) != 1 {
		s.s2.Put(s.s1.Pop())
	}
	val := s.s1.Pop()
	for s.s2.Empty() != true {
		s.s1.Put(s.s2.Pop())
	}
	return val
}

func main() {
	var s stack
	fmt.Println("Stack:")
	fmt.Println("Push 1")
	s.Put(1)
	fmt.Println("Push 2")
	s.Put(2)
	fmt.Println("Push 3")
	s.Put(3)
	fmt.Println("Pop-ing:")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	var sq staqueue
	fmt.Println("Staqueue:")
	fmt.Println("Push 1")
	sq.Push(1)
	fmt.Println("Push 2")
	sq.Push(2)
	fmt.Println("Push 3")
	sq.Push(3)
	fmt.Println("Pop-ing:")
	fmt.Println(sq.Pop())
	fmt.Println(sq.Pop())
	fmt.Println(sq.Pop())
	
}
