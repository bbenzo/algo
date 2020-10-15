package binary_search_tree

import "reflect"

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

func find(parent *Node, node *Node, value int) (*Node, *Node) {
	if node.Value == value {
		return parent, node
	} else if value < node.Value {
		return find(node, node.Left, value)
	} else if value > node.Value {
		return find(node, node.Right, value)
	}
	return parent, nil
}

func largest(parent *Node, node *Node) *Node {
	if node.Right != nil {
		return largest(parent, node.Right)
	}
	return parent
}

func lowest(parent *Node, node *Node) *Node {
	if node.Left != nil {
		return lowest(parent, node.Left)
	}
	return parent
}

func wipePassed(r interface{}){
    v := reflect.ValueOf(&r)
    p := v.Elem()
    p.Set(reflect.Zero(p.Type()))
}

func niller() {
    return
}

func (t *BinarySearchTree) Remove(value int) {
	if t.Root == nil {
		return
	}

	parent, node := find(nil, t.Root, value)
	if node == nil {
		return
	}

	if node.Left != nil {
		largest := largest(node, node.Left)
		largest.Right = nil
	} else if node.Right != nil {
		lowest := lowest(node, node.Right)
		lowest.Left = nil
	} else if parent.Left.Value == value {
		parent.Left = nil
	} else if parent.Right.Value == value {
		parent.Right = nil
	}
}
