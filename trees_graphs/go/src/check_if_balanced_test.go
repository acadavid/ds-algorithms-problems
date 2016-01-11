// Check if a binary tree is balanced,
// that is the diff of distance between
// any leaf and the root is at most 1

package main

import (
	"math"
	"testing"
	. "binary_tree"
)

func maxDepth(n *Node) float64 {
	if n == nil {
		return 0
	}

	return 1 + math.Max(maxDepth(n.GetRight()), maxDepth(n.GetLeft()))
}

func minDepth(n *Node) float64 {
	if n == nil {
		return 0
	}

	return 1 + math.Min(minDepth(n.GetRight()), minDepth(n.GetLeft()))
}

func checkIfBalanced(node Node) bool {
	return ((maxDepth(&node) - minDepth(&node)) <= 1)
}

func TestCheckBalanced(t *testing.T) {
	n := Node{Value: 20}
	n.InsertNode(&Node{Value: 25})
	n.InsertNode(&Node{Value: 10})
	n.InsertNode(&Node{Value: 5})

	if checkIfBalanced(n) != true {
		t.Error("Expected true (balanced), got", checkIfBalanced(n))
	}
	
	n.InsertNode(&Node{Value: 4})
	if checkIfBalanced(n) != false {
		t.Error("Expected false (unbalanced), got", checkIfBalanced(n))
	}
}
