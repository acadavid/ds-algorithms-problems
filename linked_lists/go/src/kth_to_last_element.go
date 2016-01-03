package main

import (
	"fmt"
	. "linked_list"
)

func GetKthToLast(node *SLinkedListNode, n int) *SLinkedListNode {
	head_node := node
	total_nodes := 1
	for node.GetNext() != nil {
		total_nodes += 1
		node = node.GetNext()
	}

	current_node := head_node
	for i := 0; i <= total_nodes - n; i++ {		
		if i == (total_nodes - n) {
			return current_node
		}
		current_node = current_node.GetNext()
	}

	return nil
}

func GetKthToLastFaster(node *SLinkedListNode, n int) *SLinkedListNode {
	fst_pointer := node
	snd_pointer := node
	for i := 0; i < n; i++ {
		snd_pointer = snd_pointer.GetNext()
	}
	for snd_pointer != nil {
		fst_pointer = fst_pointer.GetNext()
		snd_pointer = snd_pointer.GetNext()
	}

	return fst_pointer
}

func main() {
	node := SLinkedListNode{}
	node.SetValue(1)
	node.PushValueToTail(2)
	node.PushValueToTail(3)
	node.PushValueToTail(4)
	node.PushValueToTail(5)
	node.PushValueToTail(6)

	kth_to_last := GetKthToLastFaster(&node, 4)
	fmt.Println(kth_to_last.GetValue())
}
