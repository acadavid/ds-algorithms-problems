package main

import (
	"testing"
	. "binary_tree"
)

func TestBST(t *testing.T) {
	n := Node{Value: 20}
	n.InsertNode(&Node{Value: 5})
	n.InsertNode(&Node{Value: 10})
	n.InsertNode(&Node{Value: 11})
	n.InsertNode(&Node{Value: 9})
	n.InsertNode(&Node{Value: 7})
	n.InsertNode(&Node{Value: 4})
	n.InsertNode(&Node{Value: 6})

	found_n := n.FindValue(4)

	if found_n.GetValue() != 4 {
		t.Error("Expected 4, got", found_n.GetValue())
	}
}
