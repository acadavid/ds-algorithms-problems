package stack_queue

import "math"

type Stack struct {
	data []int
}

func (s *Stack) Pop() int {
	elem := s.data[len(s.data)-1]
	s.data = s.data[0:len(s.data)-1]
	return elem
}

func (s *Stack) Push(elem int) {
	s.data = append(s.data, elem)
}

func (s Stack) Empty() bool {
	return len(s.data) == 0
}

func (s Stack) Length() int {
	return len(s.data)
}

type MinStack struct {
	data []ElemWithMin
}

// Stack with Minimum value access in O(1)

type ElemWithMin struct {
	val int
	min int
	prev_node *ElemWithMin
}

func (elem ElemWithMin) GetValue() int {
	return elem.val
}

func (ms *MinStack) Push(elem int) {
	var min int
	var new_elem ElemWithMin
	
	top := ms.Top()
	if (top != ElemWithMin{}) {
		min = int(math.Min(float64(elem), float64(top.min)))
		new_elem = ElemWithMin{elem, min, &top}
	} else {
		min = elem
		new_elem = ElemWithMin{elem, min, nil}
	}
	
	ms.data = append(ms.data, new_elem)
}

func (ms *MinStack) Pop() ElemWithMin {
	elem := ms.data[len(ms.data)-1]
	ms.data = ms.data[0:len(ms.data)-1]
	return elem
}

func (ms MinStack) Top() ElemWithMin {
	if len(ms.data) == 0 {
		return ElemWithMin{}
	}
	return ms.data[len(ms.data)-1]
}

func (ms MinStack) Min() int {
	return ms.Top().min
}
