// Implement a Queue using two stacks.

package main

import (
	"testing"
	. "stack_queue"
)

type StaQueue struct {
	s1 Stack
	s2 Stack
}

func (s *StaQueue) Push(i int) {
	s.s1.Push(i)
}

func (s *StaQueue) Pop() int {
	for s.s1.Length() != 1 {
		s.s2.Push(s.s1.Pop())
	}
	val := s.s1.Pop()
	for s.s2.Empty() != true {
		s.s1.Push(s.s2.Pop())
	}

	return val
}

func TestStack(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	val := s.Pop()
	if val != 3 {
		t.Error("Expecting 3, got", val)
	}

	val = s.Pop()
	if val != 2  {
		t.Error("Expecting 2, got", val)
	}
}

func TestStaQueue(t *testing.T) {
	var sq StaQueue

	sq.Push(1)
	sq.Push(2)
	sq.Push(3)

	val := sq.Pop()
	if val != 1 {
		t.Error("Expecting 1, got", val)
	}

	val = sq.Pop()
	if val != 2  {
		t.Error("Expecting 2, got", val)
	}
}
