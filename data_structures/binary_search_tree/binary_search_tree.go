package binary_search_tree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type BinarySearchTree struct {
	Root *Node
}

func compare(current *Node, newNode *Node) {
	if newNode.Value < current.Value {
		if current.Left != nil {
			compare(current.Left, newNode)
		} else {
			current.Left = newNode
		}
	} else {
		if current.Right != nil {
			compare(current.Right, newNode)
		} else {
			current.Right = newNode
		}
	}
}

func (t *BinarySearchTree) Insert(value int) {
	node := &Node{
		Value: value,
	}

	if t.Root == nil {
		t.Root = node
		return
	}

	compare(t.Root, node)
}
