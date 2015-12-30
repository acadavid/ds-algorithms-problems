// The brute force way to solve this problem is by having a set of visited
// nodes and then check if the next node has already been visited.

// There is a more efficient solution called Floyd’s cycle-finding algorithm
// (https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare) that
// uses two pointers traveling at 1x and 2x speed each and then checking if they eventually are the same.

package main

import "fmt"

type Node struct {
	next *Node
	data int
}

func (n *Node) AddElementAtEnd(elem int) *Node{
	for(n.next != nil) {
		n = n.next
	}

	new_node := Node{nil, elem}
	n.next = &new_node

	return &new_node
}

func (n *Node) TraverseList() {
	fmt.Println(n.data)
	for(n.next != nil) {
		n = n.next
		fmt.Println(n.data)
	}
}

func (n *Node) PointLastElementTo(point_to *Node) bool{
	for(n.next != nil) {
		n = n.next
	}

	n.next = point_to

	return true
}

func (n *Node) DetectCycle() bool{
	visited := make(map[*Node]bool)
	visited[n] = true
	n = n.next
	for(visited[n] != true && n.next != nil) {
		visited[n] = true
		n = n.next
	}

	if (visited[n] == true) {
		return true
	} else {
		return false
	}
}

func main() {
	head_node := Node{nil, 1}
	head_node.AddElementAtEnd(2)
	trd_node := head_node.AddElementAtEnd(3)
	head_node.AddElementAtEnd(4)
	ffth_node := head_node.AddElementAtEnd(5)
	trd_node.PointLastElementTo(ffth_node)
	fmt.Println("Has cycle?: ")
	fmt.Println(head_node.DetectCycle())
	//head_node.TraverseList()
}
