package linked_list

import "fmt"

type SLinkedListNode struct {
	next *SLinkedListNode
	val int
}

func (n *SLinkedListNode) GetValue() int {
	return n.val
}

func (n *SLinkedListNode) GetNext() *SLinkedListNode {
	return n.next
}

func (n *SLinkedListNode) SetValue(value int) {
	n.val = value
}

func (n *SLinkedListNode) SetNext(new_node *SLinkedListNode) {
	n.next = new_node
}

func (n *SLinkedListNode) TraverseList() {
	fmt.Println(n.val)
	for n.next != nil {	
		n = n.next
		fmt.Println(n.val)
	}
}

func (n *SLinkedListNode) PushValueToTail(k int) *SLinkedListNode {
	for n.next != nil {
		n = n.next
	}

	n_node := SLinkedListNode{nil, k}
	n.next = &n_node
	return &n_node
}

func (n *SLinkedListNode) DeleteFirstValueFound(k int) bool {
	prev_node := n
	for n.val != k {
		prev_node = n
		n = n.next
	}

	if n.val == k {
		prev_node.next = n.next
		return true
	} else {
		return false
	}
}

func (n *SLinkedListNode) PointLastElementTo(point_to *SLinkedListNode) bool {
	for n.next != nil {
		n = n.next
	}
	n.next = point_to

	return true
}


// func (n *SLinkedListNode) Clear() {
	
// }

// func (n *SLinkedListNode) GetLast() *SLinkedListNode {
	
// }

// func (n *SLinkedListNode) GetFirst() *SLinkedListNode {
// 	return *n
// }

// func (n *SLinkedListNode) Size() int {
	
// }
