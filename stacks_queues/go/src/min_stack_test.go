package main

import (
	"testing"
	. "stack_queue"
)

func TestMin(t *testing.T) {
	min_stack := MinStack{}
	min_stack.Push(3)
	min_stack.Push(9)
	min_stack.Push(5)
	min_stack.Push(4)
	min_stack.Push(2)
	min_stack.Push(8)
	min_stack.Push(1)
	min_stack.Push(10)

	min := min_stack.Min()
	if min != 1 {
		t.Error("Expecting 1, got ", min_stack.Min())
	}

	min_stack.Pop()
	min_stack.Pop()
	min_stack.Pop()
	min_stack.Pop()

	min = min_stack.Min()
	if min != 3 {
		t.Error("Expecting 3, got ", min_stack.Min())
	}
}
