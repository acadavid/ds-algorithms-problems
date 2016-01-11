package binary_tree

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func (n *Node) GetValue() int {
	return n.Value
}

func (n *Node) GetRight() *Node {
	return n.Right
}

func (n *Node) GetLeft() *Node {
	return n.Left
}

func (n *Node) SetRight(nn *Node) *Node {
	n.Right = nn
	return nn
}

func (n *Node) SetLeft(nn *Node) *Node {
	n.Left = nn
	return nn
}

func (n *Node) InsertNode(nn *Node) *Node {
	if nn.GetValue() > n.GetValue() {
		if n.GetRight() == nil {
			n.SetRight(nn)
		} else {
			n.GetRight().InsertNode(nn)
		}
	} else {
		if n.GetLeft() == nil {
			n.SetLeft(nn)
		} else {
			n.GetLeft().InsertNode(nn)
		}
	}

	return nn
}

func (n *Node) FindValue(v int) *Node {
	if n.GetValue() == v {
		return n
	}

	if v > n.GetValue() {
		return n.GetRight().FindValue(v)
	} else {
		return n.GetLeft().FindValue(v)
	}

}
