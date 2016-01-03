// The brute force way to solve this problem is by having a set of visited
// nodes and then check if the next node has already been visited.

// There is a more efficient solution called Floyd’s cycle-finding algorithm
// (https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare) that
// uses two pointers traveling at 1x and 2x speed each and then checking if they eventually are the same.

package main

import (
	"fmt"
	. "linked_list"
)

func DetectCycle(n *SLinkedListNode) bool{
	visited := make(map[*SLinkedListNode]bool)
	visited[n] = true
	n = n.GetNext()
	for visited[n] != true && n.GetNext() != nil {
		visited[n] = true
		n = n.GetNext()
	}

	if visited[n] == true {
		return true
	} else {
		return false
	}
}

func main() {
	head_node := SLinkedListNode{}
	head_node.SetValue(1)
	head_node.PushValueToTail(2)
	trd_node := head_node.PushValueToTail(3)
	head_node.PushValueToTail(4)
	ffth_node := head_node.PushValueToTail(5)
	trd_node.PointLastElementTo(ffth_node)
	fmt.Println("Has cycle?: ")
	fmt.Println(DetectCycle(&head_node))
}
