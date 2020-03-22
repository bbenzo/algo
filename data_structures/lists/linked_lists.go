package main

import "fmt"

type Node struct {
	Data interface{}
	Last *Node
	Next *Node
}

type DoublyLinkedList []Node

func New(data ...interface{}) DoublyLinkedList {
	list := DoublyLinkedList{}
	for _, el := range data {
		list = Append(list, el)
	}
	return list
}

func (list DoublyLinkedList) Head() *Node {
	if len(list) == 0 {
		return nil
	}
	return &list[0]
}

func (list DoublyLinkedList) Tail() *Node {
	if len(list) == 0 {
		return nil
	}
	return &list[len(list) - 1]
}

func Append(list DoublyLinkedList, data interface{}) DoublyLinkedList {
	if len(list) == 0 {
		return DoublyLinkedList{
			Node{
				Data: data,
			},
		}
	}

	last := &list[len(list)-1]
	node := Node{
		Data: data,
		Last: last,
	}

	last.Next = &node
	return append(list, node)
}

func Insert(list DoublyLinkedList, d interface{}, i int) DoublyLinkedList {
	if i < 0 || i > len(list)-1 {
		return list
	}

	if len(list) == 0 {
		list[0] = Node{
			Data: d,
		}
		return list
	}

	node := Node{
		Data: d,
	}

	first := list[:i]
	var second = make([]Node, len(list[i:]))
	copy(second, list[i:])

	last := &first[len(first)-1]
	last.Next = &node
	node.Last = last
	node.Next = &second[0]
	second[0].Last = &node

	return append(first, append([]Node{node}, second...)...)
}

func Remove(list DoublyLinkedList, i int) DoublyLinkedList {
	if i < 0 || i > len(list)-1 {
		return list
	}

	list = append(list[:i], list[i+1:]...)

	if len(list) == 0 {
		return list
	}

	list[i-1].Next = &list[i]

	for j := i; j < len(list); j++ {
		if j != len(list)-1 {
			list[j].Last = &list[j-1]
			list[j].Next = &list[j+1]
		} else {
			list[j].Last = &list[j-1]
			list[j].Next = nil
		}
	}

	return list
}

func main() {
	list := DoublyLinkedList{}
	list = Append(list,23)
	list = Append(list, 8)
	list = Append(list, 12)
	list = Append(list, -6)
	list = Append(list, -11)
	list = Append(list, -8)
	list = Append(list, 42)

	list = Insert(list, 5, 4)

	list = Remove(list, 3)

	for _, el := range list {
		fmt.Printf("[Element: %v] \n", el.Data)

		if el.Last != nil {
			fmt.Printf("-> Last %v \n", el.Last.Data)
		}

		if el.Next != nil {
			fmt.Printf("-> Next %v\n\n", el.Next.Data)
		}
	}
}
